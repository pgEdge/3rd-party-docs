//-------------------------------------------------------------------------
//
// pgEdge PostgreSQL Docs
//
// Copyright (c) 2026, pgEdge, Inc.
// This software is released under The PostgreSQL License
//
//-------------------------------------------------------------------------

package rst

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/pgEdge/postgresql-docs/builder/shared"
)

// ConvertContext holds state for the RST-to-Markdown conversion.
type ConvertContext struct {
	SrcDir        string
	OutDir        string
	Version       string
	Copyright     string
	PgAdminSrcDir string
	Verbose       bool

	FileMap          map[string]string
	LabelMap         map[string]labelInfo
	Substitutions    map[string]*Node
	HyperlinkTargets map[string]string // name -> URL
	CurrentFile      string
	Warnings         []string
}

// Converter orchestrates RST-to-Markdown conversion.
type Converter struct {
	ctx *ConvertContext
}

// NewConverter creates a new RST converter.
func NewConverter(
	srcDir, outDir, version, copyright, pgadminSrc string,
	verbose bool,
) *Converter {
	return &Converter{
		ctx: &ConvertContext{
			SrcDir:           srcDir,
			OutDir:           outDir,
			Version:          version,
			Copyright:        copyright,
			PgAdminSrcDir:    pgadminSrc,
			Verbose:          verbose,
			Substitutions:    make(map[string]*Node),
			HyperlinkTargets: make(map[string]string),
		},
	}
}

// Convert runs the full conversion pipeline.
func (c *Converter) Convert() error {
	initDirectiveHandlers()
	ctx := c.ctx

	// Phase 1: Resolve toctree and build file/label maps
	if ctx.Verbose {
		fmt.Println("  RST Phase 1: Resolving toctree...")
	}
	_, fileMap, labelMap, _, warnings := ResolveToctree(ctx.SrcDir)
	ctx.FileMap = fileMap
	ctx.LabelMap = labelMap
	ctx.Warnings = append(ctx.Warnings, warnings...)

	if ctx.Verbose {
		fmt.Printf("    Found %d files, %d labels\n",
			len(fileMap), len(labelMap))
	}

	// Phase 2: Scan all labels for complete cross-reference map
	if ctx.Verbose {
		fmt.Println("  RST Phase 2: Scanning labels...")
	}
	ScanAllLabels(ctx.SrcDir, ctx.FileMap, ctx.LabelMap)
	CollectAllLabels(ctx.SrcDir, ctx.FileMap, ctx.LabelMap)

	if ctx.Verbose {
		fmt.Printf("    Total labels: %d\n", len(ctx.LabelMap))
	}

	// Phase 3: Convert each file
	if ctx.Verbose {
		fmt.Println("  RST Phase 3: Converting files...")
	}
	for rstName, outputPath := range ctx.FileMap {
		if err := c.convertFile(rstName, outputPath); err != nil {
			ctx.Warnings = append(ctx.Warnings,
				fmt.Sprintf("error converting %s: %v", rstName, err))
		}
	}

	// Phase 4: Copy images
	if ctx.Verbose {
		fmt.Println("  RST Phase 4: Copying images...")
	}
	c.copyAllImages()

	return nil
}

// Warnings returns accumulated warnings.
func (c *Converter) Warnings() []string {
	return c.ctx.Warnings
}

// Files returns the list of generated files for nav generation.
func (c *Converter) Files() []*shared.FileEntry {
	_, _, _, fileEntries, _ := ResolveToctree(c.ctx.SrcDir)
	return fileEntries
}

// convertFile parses and converts a single RST file.
func (c *Converter) convertFile(rstName, outputPath string) error {
	ctx := c.ctx
	rstPath := filepath.Join(ctx.SrcDir, rstName+".rst")

	data, err := os.ReadFile(rstPath)
	if err != nil {
		return fmt.Errorf("reading %s: %w", rstPath, err)
	}

	ctx.CurrentFile = outputPath

	// Parse
	root := Parse(string(data))

	// Collect substitution definitions from this file
	collectSubstitutions(root, ctx.Substitutions, ctx.HyperlinkTargets)

	// Convert
	w := shared.NewMarkdownWriter()
	for _, child := range root.Children {
		convertNode(ctx, child, w)
	}

	// Write output
	content := cleanMarkdown(w.String())
	return c.writeFile(outputPath, content)
}

// convertNode converts a single RST node to Markdown.
func convertNode(
	ctx *ConvertContext,
	node *Node,
	w *shared.MarkdownWriter,
) {
	switch node.Type {
	case HeadingNode:
		convertHeading(ctx, node, w)
	case ParagraphNode:
		convertParagraph(ctx, node, w)
	case DirectiveNode:
		convertDirective(ctx, node, w)
	case BulletListNode:
		convertBulletList(ctx, node, w)
	case EnumListNode:
		convertEnumList(ctx, node, w)
	case ListItemNode:
		// Should be handled by parent list
		w.Write(convertInlineCtx(ctx, node.Text))
	case BlockQuoteNode:
		convertBlockQuote(ctx, node, w)
	case LiteralBlockNode:
		convertLiteralBlock(ctx, node, w)
	case LabelNode:
		convertLabel(ctx, node, w)
	case SubstitutionDefNode:
		// Handled during collection phase — skip
	case FieldListNode:
		convertFieldList(ctx, node, w)
	case GridTableNode:
		convertGridTable(ctx, node, w)
	case TransitionNode:
		w.BlankLine()
		w.WriteString("---\n")
	case LineBlockNode:
		convertLineBlock(ctx, node, w)
	case CommentNode:
		// Skip comments
	}
}

// convertHeading writes a Markdown heading.
func convertHeading(
	ctx *ConvertContext,
	node *Node,
	w *shared.MarkdownWriter,
) {
	title := cleanTitle(node.Text)
	title = convertInlineCtx(ctx, title)
	w.Heading(node.Level, title, "")
}

// convertParagraph writes a paragraph with inline conversion.
func convertParagraph(
	ctx *ConvertContext,
	node *Node,
	w *shared.MarkdownWriter,
) {
	w.BlankLine()
	text := convertInlineCtx(ctx, node.Text)
	w.WriteString(text + "\n")
}

// convertDirective dispatches to the appropriate handler.
func convertDirective(
	ctx *ConvertContext,
	node *Node,
	w *shared.MarkdownWriter,
) {
	handler, ok := directiveHandlers[strings.ToLower(node.DirectiveName)]
	if !ok {
		ctx.Warnings = append(ctx.Warnings,
			fmt.Sprintf("unhandled directive: %s", node.DirectiveName))
		// Try to render body as plain text
		if node.Body != "" {
			w.BlankLine()
			w.WriteString(convertInlineCtx(ctx, node.Body) + "\n")
		}
		return
	}
	if err := handler(ctx, node, w); err != nil {
		ctx.Warnings = append(ctx.Warnings,
			fmt.Sprintf("directive %s error: %v",
				node.DirectiveName, err))
	}
}

// convertBulletList writes an unordered list.
func convertBulletList(
	ctx *ConvertContext,
	node *Node,
	w *shared.MarkdownWriter,
) {
	w.BlankLine()
	for _, item := range node.Children {
		w.BlankLine()
		convertListItem(ctx, item, w, "- ", "    ")
	}
}

// convertEnumList writes an ordered list.
func convertEnumList(
	ctx *ConvertContext,
	node *Node,
	w *shared.MarkdownWriter,
) {
	w.BlankLine()
	for i, item := range node.Children {
		w.BlankLine()
		prefix := fmt.Sprintf("%d. ", i+1)
		// MkDocs requires 4-space indent for nested blocks
		indent := "    "
		if len(prefix) > 4 {
			indent = strings.Repeat(" ", len(prefix))
		}
		convertListItem(ctx, item, w, prefix, indent)
	}
}

// convertListItem renders a single list item.  If the item's raw
// text contains embedded RST directives (e.g. code-block, image,
// csv-table inside a bullet item) we sub-parse it so the directives
// are properly converted instead of emitted as literal text.
func convertListItem(
	ctx *ConvertContext,
	item *Node,
	w *shared.MarkdownWriter,
	prefix, indent string,
) {
	text := item.Text

	// Check whether the raw text contains an embedded directive.
	if containsDirective(text) {
		// Sub-parse the item text as a mini-RST document.
		subRoot := Parse(text)
		first := true
		for _, child := range subRoot.Children {
			if first {
				// First block: emit with list prefix
				subW := shared.NewMarkdownWriter()
				convertNode(ctx, child, subW)
				content := strings.TrimSpace(subW.String())
				lines := strings.Split(content, "\n")
				for j, line := range lines {
					if j == 0 {
						w.WriteString(prefix + line + "\n")
					} else if line == "" {
						w.WriteString("\n")
					} else {
						w.WriteString(indent + line + "\n")
					}
				}
				first = false
			} else {
				// Subsequent blocks: indent under the list item
				subW := shared.NewMarkdownWriter()
				convertNode(ctx, child, subW)
				content := strings.TrimSpace(subW.String())
				w.WriteString("\n")
				for _, line := range strings.Split(content, "\n") {
					if line == "" {
						w.WriteString("\n")
					} else {
						w.WriteString(indent + line + "\n")
					}
				}
			}
		}
		return
	}

	// Simple case: no embedded directives — inline conversion only.
	converted := convertInlineCtx(ctx, text)
	lines := strings.Split(converted, "\n")
	for i, line := range lines {
		if i == 0 {
			w.WriteString(prefix + line + "\n")
		} else if line == "" {
			w.WriteString("\n")
		} else {
			w.WriteString(indent + line + "\n")
		}
	}
}

// containsDirective checks whether text contains an RST directive
// pattern (.. name::) on its own line.
func containsDirective(text string) bool {
	for _, line := range strings.Split(text, "\n") {
		trimmed := strings.TrimSpace(line)
		if strings.HasPrefix(trimmed, ".. ") &&
			strings.Contains(trimmed[3:], "::") {
			return true
		}
	}
	return false
}

// convertBlockQuote writes a blockquote.
func convertBlockQuote(
	ctx *ConvertContext,
	node *Node,
	w *shared.MarkdownWriter,
) {
	w.BlankLine()
	text := convertInlineCtx(ctx, node.Text)
	for _, line := range strings.Split(text, "\n") {
		w.WriteString("> " + line + "\n")
	}
}

// convertLiteralBlock writes a fenced code block.
func convertLiteralBlock(
	ctx *ConvertContext,
	node *Node,
	w *shared.MarkdownWriter,
) {
	w.StartCodeBlock("")
	w.WriteString(node.Text)
	w.EndCodeBlock()
}

// convertLabel writes an HTML anchor for a cross-reference label.
func convertLabel(
	ctx *ConvertContext,
	node *Node,
	w *shared.MarkdownWriter,
) {
	w.WriteString(fmt.Sprintf("<a id=\"%s\"></a>\n", node.Label))
}

// convertFieldList writes a field list as a definition list.
func convertFieldList(
	ctx *ConvertContext,
	node *Node,
	w *shared.MarkdownWriter,
) {
	w.BlankLine()
	for _, field := range node.Children {
		w.WriteString("**" + field.FieldName + "**\n")
		w.WriteString(":   " +
			convertInlineCtx(ctx, field.FieldBody) + "\n\n")
	}
}

// convertLineBlock writes a line block preserving line breaks.
func convertLineBlock(
	ctx *ConvertContext,
	node *Node,
	w *shared.MarkdownWriter,
) {
	w.BlankLine()
	for _, line := range strings.Split(node.Text, "\n") {
		w.WriteString(convertInlineCtx(ctx, line) + "<br>\n")
	}
}

// convertInlineCtx is a convenience wrapper for ConvertInline.
func convertInlineCtx(ctx *ConvertContext, text string) string {
	if ctx == nil {
		return text
	}
	return ConvertInline(
		text,
		ctx.LabelMap,
		ctx.FileMap,
		ctx.CurrentFile,
		ctx.Substitutions,
		ctx.HyperlinkTargets,
	)
}

// collectSubstitutions extracts substitution definitions and
// hyperlink targets from a document tree.
func collectSubstitutions(root *Node, subs map[string]*Node, targets map[string]string) {
	for _, child := range root.Children {
		if child.Type == SubstitutionDefNode {
			subs[child.SubstitutionName] = child
		}
		if child.Type == LabelNode && child.Text != "" {
			// Hyperlink target: .. _name: URL
			targets[child.Label] = child.Text
		}
	}
}

// copyImage copies an image from source to output directory.
func (ctx *ConvertContext) copyImage(imgPath string) {
	if ctx.SrcDir == "" || ctx.OutDir == "" {
		return
	}

	srcPath := filepath.Join(ctx.SrcDir, imgPath)
	dstPath := filepath.Join(ctx.OutDir, imgPath)

	// Create destination directory
	if err := os.MkdirAll(filepath.Dir(dstPath), 0755); err != nil {
		ctx.Warnings = append(ctx.Warnings,
			fmt.Sprintf("could not create image dir: %v", err))
		return
	}

	data, err := os.ReadFile(srcPath)
	if err != nil {
		ctx.Warnings = append(ctx.Warnings,
			fmt.Sprintf("could not read image %s: %v", srcPath, err))
		return
	}

	if err := os.WriteFile(dstPath, data, 0644); err != nil {
		ctx.Warnings = append(ctx.Warnings,
			fmt.Sprintf("could not write image %s: %v", dstPath, err))
	}
}

// copyAllImages copies the entire images directory.
func (c *Converter) copyAllImages() {
	ctx := c.ctx
	imgSrcDir := filepath.Join(ctx.SrcDir, "images")
	imgDstDir := filepath.Join(ctx.OutDir, "images")

	if _, err := os.Stat(imgSrcDir); os.IsNotExist(err) {
		return
	}

	if err := os.MkdirAll(imgDstDir, 0755); err != nil {
		ctx.Warnings = append(ctx.Warnings,
			fmt.Sprintf("could not create images dir: %v", err))
		return
	}

	entries, err := os.ReadDir(imgSrcDir)
	if err != nil {
		ctx.Warnings = append(ctx.Warnings,
			fmt.Sprintf("could not read images dir: %v", err))
		return
	}

	count := 0
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		srcPath := filepath.Join(imgSrcDir, entry.Name())
		dstPath := filepath.Join(imgDstDir, entry.Name())

		data, err := os.ReadFile(srcPath)
		if err != nil {
			continue
		}
		if err := os.WriteFile(dstPath, data, 0644); err != nil {
			continue
		}
		count++
	}

	if ctx.Verbose {
		fmt.Printf("    Copied %d images\n", count)
	}
}

// writeFile writes content to an output file.
func (c *Converter) writeFile(relPath, content string) error {
	fullPath := filepath.Join(c.ctx.OutDir, relPath)
	dir := filepath.Dir(fullPath)

	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("creating directory %s: %w", dir, err)
	}

	return os.WriteFile(fullPath, []byte(content), 0644)
}

// cleanMarkdown performs final cleanup on generated Markdown.
func cleanMarkdown(content string) string {
	// Remove trailing whitespace on each line
	lines := strings.Split(content, "\n")
	for i, line := range lines {
		lines[i] = strings.TrimRight(line, " \t")
	}
	content = strings.Join(lines, "\n")

	// Collapse more than 2 consecutive blank lines
	for strings.Contains(content, "\n\n\n\n") {
		content = strings.ReplaceAll(content, "\n\n\n\n", "\n\n\n")
	}

	// Ensure single trailing newline
	content = strings.TrimRight(content, "\n") + "\n"

	return content
}
