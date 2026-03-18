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
	"regexp"
	"strings"

	"github.com/pgEdge/postgresql-docs/builder/shared"
)

// DirectiveHandler converts a directive node to Markdown.
type DirectiveHandler func(
	ctx *ConvertContext,
	node *Node,
	w *shared.MarkdownWriter,
) error

// directiveHandlers maps directive names to their handlers.
// Populated by initDirectiveHandlers to avoid initialization cycles.
var directiveHandlers map[string]DirectiveHandler

func initDirectiveHandlers() {
	if directiveHandlers != nil {
		return
	}
	directiveHandlers = map[string]DirectiveHandler{
		"image":          handleImage,
		"figure":         handleFigure,
		"code-block":     handleCodeBlock,
		"sourcecode":     handleCodeBlock,
		"note":           handleAdmonition,
		"warning":        handleAdmonition,
		"tip":            handleAdmonition,
		"caution":        handleAdmonition,
		"important":      handleAdmonition,
		"danger":         handleAdmonition,
		"hint":           handleAdmonition,
		"attention":      handleAdmonition,
		"admonition":     handleAdmonition,
		"toctree":        handleToctreeDirective,
		"csv-table":      handleCSVTable,
		"table":          handleTableDirective,
		"topic":          handleTopic,
		"youtube":        handleYouTube,
		"literalinclude": handleLiteralInclude,
		"deprecated":     handleDeprecated,
		"versionadded":   handleVersionChanged,
		"versionchanged": handleVersionChanged,
		"seealso":        handleSeeAlso,
		"contents":       handleSkipDirective,
		"raw":            handleRaw,
		"only":           handleOnly,
		"highlight":      handleSkipDirective,
		"index":          handleSkipDirective,

		// Sphinx domain directives (API documentation)
		"class":     handleAPIDef,
		"method":    handleAPIDef,
		"function":  handleAPIDef,
		"attribute": handleAPIDef,
		"data":      handleAPIDef,
		"exception": handleAPIDef,
		"module":    handleAPIDef,
		"decorator": handleAPIDef,

		// Sphinx autodoc directives
		"autoclass":     handleAPIDef,
		"autofunction":  handleAPIDef,
		"automethod":    handleAPIDef,
		"autoattribute": handleAPIDef,
		"automodule":    handleAPIDef,
		"autoexception": handleAPIDef,

		// Sphinx misc directives
		"rubric":         handleRubric,
		"parsed-literal": handleParsedLiteral,
		"sectionauthor":  handleSkipDirective,
		"testsetup":      handleSkipDirective,
		"testcode":       handleSkipDirective,
		"testcleanup":    handleSkipDirective,
		"doctest":        handleDoctest,
		"cssclass":       handleSkipDirective,
		"ifconfig":       handleSkipDirective,
		"todolist":       handleSkipDirective,
		"extension":      handleAPIDef,
		"include":        handleInclude,
		"code":           handleCodeBlock,

		// Container/layout directives
		"container":  handleContainer,
		"title":      handleSkipDirective,
		"list-table": handleListTable,
		"tabs":       handleTabs,
		"tab":        handleTab,
	}
}

// colorSchemeFragment returns the MkDocs Material URL fragment
// for dark/light image switching, or "" if not in a themed container.
func colorSchemeFragment(ctx *ConvertContext) string {
	switch ctx.ColorScheme {
	case "dark":
		return "#only-dark"
	case "light":
		return "#only-light"
	}
	return ""
}

// handleImage converts an image directive to Markdown.
func handleImage(
	ctx *ConvertContext,
	node *Node,
	w *shared.MarkdownWriter,
) error {
	imgPath := strings.TrimPrefix(node.DirectiveArg, "/")
	alt := node.Options["alt"]
	if alt == "" {
		alt = "image"
	}
	target := node.Options["target"]
	frag := colorSchemeFragment(ctx)

	// External URLs — use directly, don't copy
	if strings.HasPrefix(imgPath, "http://") ||
		strings.HasPrefix(imgPath, "https://") {
		if !ctx.PrevWasImage {
			w.BlankLine()
		}
		if target != "" {
			w.WriteString(fmt.Sprintf(
				"[![%s](%s%s)](%s)\n", alt, imgPath, frag, target))
		} else {
			w.WriteString(fmt.Sprintf(
				"![%s](%s%s)\n", alt, imgPath, frag))
		}
		return nil
	}

	// Copy image file and get the output-relative path
	dstRel := ctx.copyImage(imgPath)

	// Compute relative path from current output file to the image
	relImg := relativeImagePath(ctx.CurrentFile, dstRel)

	if !ctx.PrevWasImage {
		w.BlankLine()
	}
	if target != "" {
		w.WriteString(fmt.Sprintf(
			"[![%s](%s%s)](%s)\n", alt, relImg, frag, target))
	} else {
		w.WriteString(fmt.Sprintf(
			"![%s](%s%s)\n", alt, relImg, frag))
	}

	return nil
}

// handleFigure converts a figure directive to Markdown.
func handleFigure(
	ctx *ConvertContext,
	node *Node,
	w *shared.MarkdownWriter,
) error {
	imgPath := strings.TrimPrefix(node.DirectiveArg, "/")
	alt := node.Options["alt"]
	if alt == "" {
		alt = "image"
	}

	dstRel := ctx.copyImage(imgPath)
	relImg := relativeImagePath(ctx.CurrentFile, dstRel)

	w.BlankLine()
	w.WriteString(fmt.Sprintf("![%s](%s)\n", alt, relImg))

	// Caption is in the body
	if node.Body != "" {
		w.BlankLine()
		w.WriteString("*" + convertInlineCtx(ctx, node.Body) + "*\n")
	}

	return nil
}

// handleCodeBlock converts a code-block directive.
func handleCodeBlock(
	ctx *ConvertContext,
	node *Node,
	w *shared.MarkdownWriter,
) error {
	lang := node.DirectiveArg
	w.StartCodeBlock(lang)
	w.WriteString(node.Body)
	w.EndCodeBlock()
	return nil
}

// handleAdmonition converts note/warning/tip etc. to MkDocs admonitions.
func handleAdmonition(
	ctx *ConvertContext,
	node *Node,
	w *shared.MarkdownWriter,
) error {
	kind := node.DirectiveName

	w.BlankLine()

	// Only the generic "admonition" directive uses the arg as a
	// title.  For note/warning/tip etc. the arg is the first line
	// of the body text, not a title.
	if kind == "admonition" && node.DirectiveArg != "" {
		w.WriteString(fmt.Sprintf("!!! %s \"%s\"\n\n",
			kind, node.DirectiveArg))
	} else {
		w.WriteString(fmt.Sprintf("!!! %s\n\n", kind))
	}

	// Build body: for non-"admonition" directives, prepend the arg
	// (first line of text) to the body.
	bodyPrefix := ""
	if kind != "admonition" && node.DirectiveArg != "" {
		bodyPrefix = node.DirectiveArg
	}

	// Convert children (parsed body) with 4-space indent
	content := convertAdmonitionBody(ctx, node)

	if bodyPrefix != "" {
		prefix := convertInlineCtx(ctx, bodyPrefix)
		if content != "" {
			content = prefix + " " + content
		} else {
			content = prefix
		}
	}

	for _, line := range strings.Split(content, "\n") {
		if line == "" {
			w.WriteString("\n")
		} else {
			w.WriteString("    " + line + "\n")
		}
	}
	w.EnsureNewline()
	return nil
}

// convertAdmonitionBody converts the children of an admonition node.
func convertAdmonitionBody(ctx *ConvertContext, node *Node) string {
	if len(node.Children) > 0 {
		subW := shared.NewMarkdownWriter()
		for _, child := range node.Children {
			convertNode(ctx, child, subW)
		}
		return strings.TrimSpace(subW.String())
	}
	// Fallback: use raw body with inline conversion
	if node.Body != "" {
		return convertInlineCtx(ctx, node.Body)
	}
	return ""
}

// handleToctreeDirective is a no-op — toctrees are resolved separately.
func handleToctreeDirective(
	ctx *ConvertContext,
	node *Node,
	w *shared.MarkdownWriter,
) error {
	return nil
}

// handleCSVTable converts a csv-table directive to a Markdown table.
func handleCSVTable(
	ctx *ConvertContext,
	node *Node,
	w *shared.MarkdownWriter,
) error {
	w.BlankLine()

	header := node.Options["header"]
	body := node.Body

	// Parse header
	var headerCells []string
	if header != "" {
		headerCells = parseCSVRow(header)
	}

	// Parse body rows
	var rows [][]string
	if body != "" {
		// CSV table body: each logical row may span multiple lines
		// when quoted values contain newlines
		csvRows := splitCSVRows(body)
		for _, row := range csvRows {
			cells := parseCSVRow(row)
			if len(cells) > 0 {
				rows = append(rows, cells)
			}
		}
	}

	// Determine column count
	numCols := len(headerCells)
	for _, row := range rows {
		if len(row) > numCols {
			numCols = len(row)
		}
	}
	if numCols == 0 {
		return nil
	}

	// Write header
	if len(headerCells) > 0 {
		w.WriteString("|")
		for i := 0; i < numCols; i++ {
			cell := ""
			if i < len(headerCells) {
				cell = convertInlineCtx(ctx, headerCells[i])
			}
			w.WriteString(" " + cell + " |")
		}
		w.WriteString("\n")
	} else {
		// Generate empty header
		w.WriteString("|")
		for i := 0; i < numCols; i++ {
			w.WriteString("   |")
		}
		w.WriteString("\n")
	}

	// Write separator
	w.WriteString("|")
	for i := 0; i < numCols; i++ {
		_ = i
		w.WriteString("---|")
	}
	w.WriteString("\n")

	// Write rows
	for _, row := range rows {
		w.WriteString("|")
		for i := 0; i < numCols; i++ {
			cell := ""
			if i < len(row) {
				cell = convertInlineCtx(ctx, row[i])
			}
			w.WriteString(" " + cell + " |")
		}
		w.WriteString("\n")
	}

	return nil
}

// parseCSVRow splits a CSV row into cells, handling quoting.
func parseCSVRow(line string) []string {
	var cells []string
	var current strings.Builder
	inQuote := false

	for i := 0; i < len(line); i++ {
		ch := line[i]
		if ch == '"' {
			if inQuote {
				// Check for escaped quote
				if i+1 < len(line) && line[i+1] == '"' {
					current.WriteByte('"')
					i++
				} else {
					inQuote = false
				}
			} else {
				inQuote = true
			}
		} else if ch == ',' && !inQuote {
			cells = append(cells, strings.TrimSpace(current.String()))
			current.Reset()
		} else {
			current.WriteByte(ch)
		}
	}
	cells = append(cells, strings.TrimSpace(current.String()))
	return cells
}

// splitCSVRows splits CSV body text into logical rows, handling
// multi-line quoted values.
func splitCSVRows(body string) []string {
	var rows []string
	var current strings.Builder
	inQuote := false

	for _, line := range strings.Split(body, "\n") {
		if current.Len() > 0 {
			current.WriteString(" ")
		}
		current.WriteString(line)

		// Count quotes to track state
		for _, ch := range line {
			if ch == '"' {
				inQuote = !inQuote
			}
		}

		if !inQuote {
			row := strings.TrimSpace(current.String())
			if row != "" {
				rows = append(rows, row)
			}
			current.Reset()
		}
	}

	if current.Len() > 0 {
		row := strings.TrimSpace(current.String())
		if row != "" {
			rows = append(rows, row)
		}
	}

	return rows
}

// handleTableDirective converts a .. table:: directive with a grid
// table body.
func handleTableDirective(
	ctx *ConvertContext,
	node *Node,
	w *shared.MarkdownWriter,
) error {
	title := node.DirectiveArg
	if title != "" {
		w.BlankLine()
		w.WriteString("**" + convertInlineCtx(ctx, title) + "**\n")
	}

	// The body contains a grid table — parse and render it
	if node.Body != "" {
		subRoot := Parse(node.Body)
		for _, child := range subRoot.Children {
			convertNode(ctx, child, w)
		}
	}

	return nil
}

// handleTopic converts a topic directive to a blockquote.
func handleTopic(
	ctx *ConvertContext,
	node *Node,
	w *shared.MarkdownWriter,
) error {
	title := node.DirectiveArg

	w.BlankLine()
	if title != "" {
		w.WriteString("**" + convertInlineCtx(ctx, title) + "**\n")
	}
	w.BlankLine()

	// Render body as blockquote
	content := convertAdmonitionBody(ctx, node)
	for _, line := range strings.Split(content, "\n") {
		w.WriteString("> " + line + "\n")
	}
	w.EnsureNewline()
	return nil
}

// handleYouTube converts a youtube directive to an embedded video.
func handleYouTube(
	ctx *ConvertContext,
	node *Node,
	w *shared.MarkdownWriter,
) error {
	videoID := node.DirectiveArg
	width := node.Options["width"]
	if width == "" {
		width = "560"
	}

	w.BlankLine()
	w.WriteString(fmt.Sprintf(
		"<div style=\"text-align: center;\">\n"+
			"<iframe width=\"%s\" height=\"315\" "+
			"src=\"https://www.youtube.com/embed/%s\" "+
			"frameborder=\"0\" allowfullscreen></iframe>\n"+
			"</div>\n",
		width, videoID))
	return nil
}

// handleLiteralInclude includes a file as a code block.
func handleLiteralInclude(
	ctx *ConvertContext,
	node *Node,
	w *shared.MarkdownWriter,
) error {
	filePath := node.DirectiveArg
	lang := node.Options["language"]
	if lang == "" {
		lang = node.Options["lang"]
	}

	// Try to read from pgAdmin source
	if ctx.PgAdminSrcDir != "" {
		fullPath := filepath.Join(ctx.PgAdminSrcDir, filePath)
		data, err := os.ReadFile(fullPath)
		if err == nil {
			w.StartCodeBlock(lang)
			w.WriteString(string(data))
			w.EndCodeBlock()
			return nil
		}
	}

	// Fallback: show as a reference
	w.BlankLine()
	w.WriteString(fmt.Sprintf(
		"*See source file: `%s`*\n", filePath))
	return nil
}

// handleDeprecated converts a deprecated directive.
func handleDeprecated(
	ctx *ConvertContext,
	node *Node,
	w *shared.MarkdownWriter,
) error {
	w.BlankLine()
	version := node.DirectiveArg
	w.WriteString(fmt.Sprintf(
		"!!! warning \"Deprecated since version %s\"\n\n", version))
	if node.Body != "" {
		for _, line := range strings.Split(
			convertInlineCtx(ctx, node.Body), "\n") {
			if line == "" {
				w.WriteString("\n")
			} else {
				w.WriteString("    " + line + "\n")
			}
		}
	}
	return nil
}

// handleVersionChanged converts versionadded/versionchanged.
func handleVersionChanged(
	ctx *ConvertContext,
	node *Node,
	w *shared.MarkdownWriter,
) error {
	version := node.DirectiveArg
	label := "Changed in version"
	if node.DirectiveName == "versionadded" {
		label = "New in version"
	}
	w.BlankLine()
	w.WriteString(fmt.Sprintf("*%s %s.*", label, version))
	if node.Body != "" {
		w.WriteString(" " + convertInlineCtx(ctx, node.Body))
	}
	w.WriteString("\n")
	return nil
}

// handleSeeAlso converts a seealso directive.
func handleSeeAlso(
	ctx *ConvertContext,
	node *Node,
	w *shared.MarkdownWriter,
) error {
	w.BlankLine()
	w.WriteString("!!! tip \"See Also\"\n\n")

	// Combine arg and body (like admonitions)
	fullText := ""
	if node.DirectiveArg != "" {
		fullText = node.DirectiveArg
	}
	if node.Body != "" {
		if fullText != "" {
			fullText += " " + node.Body
		} else {
			fullText = node.Body
		}
	}

	// Resolve anonymous hyperlinks: `text`__ with .. __: URL
	fullText = resolveAnonymousLinks(fullText)

	content := convertInlineCtx(ctx, fullText)
	for _, line := range strings.Split(content, "\n") {
		if line == "" {
			w.WriteString("\n")
		} else {
			w.WriteString("    " + line + "\n")
		}
	}
	return nil
}

// handleSkipDirective silently skips a directive.
func handleSkipDirective(
	ctx *ConvertContext,
	node *Node,
	w *shared.MarkdownWriter,
) error {
	return nil
}

// reObjectTag matches <object ... data="path" ...> so we can
// convert it to a Markdown image that MkDocs handles correctly.
var reObjectTag = regexp.MustCompile(
	`<object\b[^>]*\bdata="([^"]+)"[^>]*>.*?</object>`)

// handleRaw passes through raw HTML content, rewriting local
// asset paths so they resolve correctly in MkDocs output.
func handleRaw(
	ctx *ConvertContext,
	node *Node,
	w *shared.MarkdownWriter,
) error {
	format := node.DirectiveArg
	if format == "html" && node.Body != "" {
		body := node.Body

		// Convert <object data="..."> tags to Markdown
		// images so MkDocs resolves the paths correctly
		// (raw HTML paths break with use_directory_urls).
		body = reObjectTag.ReplaceAllStringFunc(
			body, func(m string) string {
				sub := reObjectTag.FindStringSubmatch(m)
				origPath := sub[1]
				if strings.HasPrefix(origPath, "http://") ||
					strings.HasPrefix(origPath, "https://") {
					return m
				}
				dstRel := ctx.copyImage(origPath)
				relPath := relativeImagePath(
					ctx.CurrentFile, dstRel)
				frag := colorSchemeFragment(ctx)
				return fmt.Sprintf(
					"![image](%s%s)", relPath, frag)
			})

		w.BlankLine()
		w.WriteString(body + "\n")
	}
	return nil
}

// handleOnly handles the .. only:: directive (conditional content).
func handleOnly(
	ctx *ConvertContext,
	node *Node,
	w *shared.MarkdownWriter,
) error {
	// Include all content regardless of condition
	if node.Body != "" {
		subRoot := Parse(node.Body)
		for _, child := range subRoot.Children {
			convertNode(ctx, child, w)
		}
	}
	return nil
}

// handleAPIDef converts Sphinx domain directives (class, method,
// function, attribute, data, exception, etc.) to a definition-style
// block with the signature in a code span and the body indented.
func handleAPIDef(
	ctx *ConvertContext,
	node *Node,
	w *shared.MarkdownWriter,
) error {
	kind := node.DirectiveName
	sig := node.DirectiveArg

	w.BlankLine()

	// For auto* directives, the arg is the object name
	// For regular directives, the arg is the full signature
	if sig != "" {
		// Strip "auto" prefix for display
		displayKind := strings.TrimPrefix(kind, "auto")
		w.WriteString(fmt.Sprintf("*%s* `%s`\n",
			displayKind, convertInlineCtx(ctx, sig)))
	}

	// Render body content
	if len(node.Children) > 0 {
		subW := shared.NewMarkdownWriter()
		for _, child := range node.Children {
			convertNode(ctx, child, subW)
		}
		body := strings.TrimSpace(subW.String())
		if body != "" {
			w.BlankLine()
			w.WriteString(body + "\n")
		}
	} else if node.Body != "" {
		subRoot := Parse(node.Body)
		subW := shared.NewMarkdownWriter()
		for _, child := range subRoot.Children {
			convertNode(ctx, child, subW)
		}
		body := strings.TrimSpace(subW.String())
		if body != "" {
			w.BlankLine()
			w.WriteString(body + "\n")
		}
	}

	return nil
}

// handleRubric converts a rubric directive to a bold heading.
func handleRubric(
	ctx *ConvertContext,
	node *Node,
	w *shared.MarkdownWriter,
) error {
	w.BlankLine()
	w.WriteString("**" + convertInlineCtx(ctx, node.DirectiveArg) + "**\n")
	return nil
}

// handleParsedLiteral converts a parsed-literal to a code block.
func handleParsedLiteral(
	ctx *ConvertContext,
	node *Node,
	w *shared.MarkdownWriter,
) error {
	w.StartCodeBlock("")
	if node.Body != "" {
		w.WriteString(node.Body)
	}
	w.EndCodeBlock()
	return nil
}

// handleDoctest converts doctest blocks to code blocks.
func handleDoctest(
	ctx *ConvertContext,
	node *Node,
	w *shared.MarkdownWriter,
) error {
	if node.Body != "" {
		w.StartCodeBlock("python")
		w.WriteString(node.Body)
		w.EndCodeBlock()
	}
	return nil
}

// handleInclude is a no-op for include directives (the included
// content is not available without the Sphinx build system).
func handleInclude(
	ctx *ConvertContext,
	node *Node,
	w *shared.MarkdownWriter,
) error {
	// Can't resolve includes without the Sphinx build
	if node.DirectiveArg != "" {
		w.BlankLine()
		w.WriteString(fmt.Sprintf(
			"*See: `%s`*\n", node.DirectiveArg))
	}
	return nil
}

// reAnonymousTarget matches ".. __: URL" anywhere in text.
var reAnonymousTarget = regexp.MustCompile(
	`\.\.\s+__:\s*(https?://\S+)`)

// reAnonymousRef matches "`text`__" (anonymous hyperlink reference).
var reAnonymousRef = regexp.MustCompile("`([^`]+)`__")

// resolveAnonymousLinks replaces RST anonymous hyperlink references
// (`text`__ paired with .. __: URL) with inline Markdown links.
func resolveAnonymousLinks(text string) string {
	// Collect all anonymous targets in order
	targets := reAnonymousTarget.FindAllStringSubmatch(text, -1)
	// Remove the target definitions from the text
	text = reAnonymousTarget.ReplaceAllString(text, "")

	// Replace `text`__ with [text](url) in order
	targetIdx := 0
	text = reAnonymousRef.ReplaceAllStringFunc(text, func(m string) string {
		sub := reAnonymousRef.FindStringSubmatch(m)
		linkText := sub[1]
		if targetIdx < len(targets) {
			url := targets[targetIdx][1]
			targetIdx++
			return "[" + linkText + "](" + url + ")"
		}
		return linkText
	})

	// Clean up extra whitespace from removed targets
	text = strings.TrimSpace(text)
	for strings.Contains(text, "  ") {
		text = strings.ReplaceAll(text, "  ", " ")
	}
	return text
}

// relativeImagePath computes the relative path from the current
// output .md file to an image (whose path is relative to the
// docs root, e.g. "images/foo.png").
// handleContainer passes through the body content of a container.
// Containers with class "img-dark" or "img-light" set the color
// scheme so child images get MkDocs Material #only-dark / #only-light
// fragments appended.
func handleContainer(
	ctx *ConvertContext,
	node *Node,
	w *shared.MarkdownWriter,
) error {
	// Detect dark/light image containers
	arg := strings.TrimSpace(node.DirectiveArg)
	prevScheme := ctx.ColorScheme
	if arg == "img-dark" {
		ctx.ColorScheme = "dark"
	} else if arg == "img-light" {
		ctx.ColorScheme = "light"
	}
	defer func() { ctx.ColorScheme = prevScheme }()

	if len(node.Children) > 0 {
		for _, child := range node.Children {
			convertNode(ctx, child, w)
		}
	} else if node.Body != "" {
		subRoot := Parse(node.Body)
		for _, child := range subRoot.Children {
			convertNode(ctx, child, w)
		}
	}
	return nil
}

// handleListTable converts a list-table directive to a Markdown table.
func handleListTable(
	ctx *ConvertContext,
	node *Node,
	w *shared.MarkdownWriter,
) error {
	title := node.DirectiveArg
	if title != "" {
		w.BlankLine()
		w.WriteString("**" + convertInlineCtx(ctx, title) + "**\n")
	}

	if node.Body == "" {
		return nil
	}

	// Parse the body — list-table body is a bullet list of rows,
	// each row is a sub-list of cells.
	lines := strings.Split(node.Body, "\n")
	var rows [][]string
	var currentRow []string
	var currentCell strings.Builder

	flushCell := func() {
		if currentCell.Len() > 0 {
			currentRow = append(currentRow,
				strings.TrimSpace(currentCell.String()))
			currentCell.Reset()
		}
	}
	flushRow := func() {
		flushCell()
		if len(currentRow) > 0 {
			rows = append(rows, currentRow)
			currentRow = nil
		}
	}

	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if trimmed == "" {
			continue
		}
		indent := countIndent(line)
		if strings.HasPrefix(trimmed, "* -") ||
			(strings.HasPrefix(trimmed, "*") && indent == 0) {
			flushRow()
			rest := strings.TrimPrefix(trimmed, "* - ")
			if rest == trimmed {
				rest = strings.TrimPrefix(trimmed, "* ")
			}
			currentCell.WriteString(rest)
		} else if strings.HasPrefix(trimmed, "- ") && indent >= 2 {
			flushCell()
			currentCell.WriteString(trimmed[2:])
		} else {
			if currentCell.Len() > 0 {
				currentCell.WriteString(" ")
			}
			currentCell.WriteString(trimmed)
		}
	}
	flushRow()

	if len(rows) == 0 {
		return nil
	}

	// Determine column count
	numCols := 0
	for _, row := range rows {
		if len(row) > numCols {
			numCols = len(row)
		}
	}

	// Check if we have a header row
	headerRows := 1
	if h, ok := node.Options["header-rows"]; ok && h == "0" {
		headerRows = 0
	}

	w.BlankLine()
	for i, row := range rows {
		w.WriteString("|")
		for j := 0; j < numCols; j++ {
			cell := ""
			if j < len(row) {
				cell = convertInlineCtx(ctx, row[j])
			}
			w.WriteString(" " + cell + " |")
		}
		w.WriteString("\n")
		if i == headerRows-1 {
			w.WriteString("|")
			for j := 0; j < numCols; j++ {
				_ = j
				w.WriteString("---|")
			}
			w.WriteString("\n")
		}
	}

	return nil
}

// handleTabs converts a tabs directive to sequential sections.
func handleTabs(
	ctx *ConvertContext,
	node *Node,
	w *shared.MarkdownWriter,
) error {
	// Tabs don't have direct MkDocs equivalent — render body
	if node.Body != "" {
		subRoot := Parse(node.Body)
		for _, child := range subRoot.Children {
			convertNode(ctx, child, w)
		}
	}
	for _, child := range node.Children {
		convertNode(ctx, child, w)
	}
	return nil
}

// handleTab converts a single tab to a bold heading + content.
func handleTab(
	ctx *ConvertContext,
	node *Node,
	w *shared.MarkdownWriter,
) error {
	if node.DirectiveArg != "" {
		w.BlankLine()
		w.WriteString("**" +
			convertInlineCtx(ctx, node.DirectiveArg) + "**\n")
	}
	if node.Body != "" {
		subRoot := Parse(node.Body)
		for _, child := range subRoot.Children {
			convertNode(ctx, child, w)
		}
	}
	return nil
}

func relativeImagePath(currentFile, imgPath string) string {
	dir := filepath.Dir(currentFile)
	if dir == "." || dir == "" {
		return imgPath
	}
	rel, err := filepath.Rel(dir, imgPath)
	if err != nil {
		return imgPath
	}
	return rel
}
