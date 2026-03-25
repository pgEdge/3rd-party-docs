//-------------------------------------------------------------------------
//
// pgEdge PostgreSQL Docs
//
// Copyright (c) 2026, pgEdge, Inc.
// This software is released under The PostgreSQL License
//
//-------------------------------------------------------------------------

package backrest

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/pgEdge/postgresql-docs/builder/sgml"
	"github.com/pgEdge/postgresql-docs/builder/shared"
)

// Converter handles pgBackRest custom XML to Markdown conversion.
type Converter struct {
	srcDir      string
	outDir      string
	version     string
	verbose     bool
	vars        map[string]string
	blocks      map[string]*sgml.Node
	idMap       map[string]*shared.IDEntry
	pageMap     map[string]string // source key → output path
	currentFile string            // path of file being written
	files       []*shared.FileEntry
	warnings    []string
}

// NewConverter creates a new pgBackRest converter.
func NewConverter(srcDir, outDir, version string, verbose bool) *Converter {
	return &Converter{
		srcDir:  srcDir,
		outDir:  outDir,
		version: version,
		verbose: verbose,
		vars:    make(map[string]string),
		blocks:  make(map[string]*sgml.Node),
		idMap:   make(map[string]*shared.IDEntry),
		pageMap: make(map[string]string),
	}
}

// Warnings returns accumulated warnings.
func (c *Converter) Warnings() []string { return c.warnings }

// Files returns the output file list for nav generation.
func (c *Converter) Files() []*shared.FileEntry { return c.files }

// ProjectName returns the project name for site_name.
func (c *Converter) ProjectName() string { return "pgBackRest" }

// Convert runs the full conversion pipeline.
func (c *Converter) Convert() error {
	// Set default variables
	c.vars["project"] = "pgBackRest"
	c.vars["project-exe"] = "pgbackrest"
	c.vars["postgres"] = "PostgreSQL"
	c.vars["dash"] = "-"
	if c.version != "" {
		c.vars["version"] = c.version
		c.vars["version-stable"] = c.version
	}

	// Discover source files
	sources := c.discoverSources()
	if len(sources) == 0 {
		return fmt.Errorf("no XML source files found in %s", c.srcDir)
	}

	// Phase 1: Parse all source files and collect variables/blocks
	docs := make(map[string]*sgml.Node)
	for key, path := range sources {
		if c.verbose {
			fmt.Printf("  Parsing %s (%s)\n", key, path)
		}
		doc, err := parseXMLFile(path)
		if err != nil {
			c.warn("failed to parse %s: %v", key, err)
			continue
		}
		docs[key] = doc
		collectVariables(doc, c.vars)
		collectBlocks(doc, c.blocks)
	}

	if len(docs) == 0 {
		return fmt.Errorf("no documents parsed successfully")
	}

	// Phase 2: Build ID map and page map
	c.buildMaps(docs)

	// Add aliases for known page references
	if _, ok := c.pageMap["user-guide-index"]; !ok {
		if ug, ok := c.pageMap["user-guide"]; ok {
			c.pageMap["user-guide-index"] = ug
		}
	}

	// Phase 3: Convert each document
	if err := os.MkdirAll(c.outDir, 0755); err != nil {
		return fmt.Errorf("creating output dir: %w", err)
	}

	for key, doc := range docs {
		if err := c.convertDocument(key, doc); err != nil {
			c.warn("converting %s: %v", key, err)
		}
	}

	return nil
}

// discoverSources finds available XML files to convert.
func (c *Converter) discoverSources() map[string]string {
	sources := make(map[string]string)

	// Hand-written docs in xml/
	xmlDir := filepath.Join(c.srcDir, "xml")
	for _, name := range []string{
		"index", "user-guide", "faq", "metric",
		"coding", "contributing", "documentation",
	} {
		path := filepath.Join(xmlDir, name+".xml")
		if _, err := os.Stat(path); err == nil {
			sources[name] = path
		}
	}

	// Auto-generated docs in output/xml/
	outXMLDir := filepath.Join(c.srcDir, "output", "xml")
	for _, name := range []string{"command", "configuration"} {
		path := filepath.Join(outXMLDir, name+".xml")
		if _, err := os.Stat(path); err == nil {
			sources[name] = path
		}
	}

	// Release notes
	releasePath := filepath.Join(xmlDir, "release.xml")
	if _, err := os.Stat(releasePath); err == nil {
		sources["release"] = releasePath
	}

	return sources
}

// docOrder defines the preferred order of documents in the nav.
var docOrder = []string{
	"index", "user-guide", "faq", "metric",
	"command", "configuration",
	"coding", "contributing", "documentation", "release",
}

// buildMaps creates the ID map and page map from parsed documents.
func (c *Converter) buildMaps(docs map[string]*sgml.Node) {
	order := 0
	// Process in fixed order for consistent nav
	for _, key := range docOrder {
		doc, ok := docs[key]
		if !ok {
			continue
		}
		sections := doc.FindChildren("section")

		// Determine if this document should be single-page
		singlePage := len(sections) <= 1 ||
			key == "faq" || key == "metric" ||
			key == "coding" || key == "contributing" ||
			key == "documentation" || key == "release" ||
			key == "index"

		if singlePage {
			// Single-page document
			slug := shared.Slugify(key)
			outPath := slug + ".md"
			if key == "index" {
				outPath = "index.md"
			}
			c.pageMap[key] = outPath

			title := c.docTitle(doc)
			if key == "index" {
				title = c.vars["project"]
			}
			c.files = append(c.files, &shared.FileEntry{
				Path:  outPath,
				Title: title,
				Order: order,
			})
			order++

			// Map section IDs
			c.mapSectionIDs(doc, outPath)
		} else {
			// Multi-page: split by top-level sections
			dirSlug := shared.Slugify(key)
			indexPath := dirSlug + "/index.md"
			c.pageMap[key] = indexPath

			title := doc.GetAttr("title")
			title = substituteVariables(title, c.vars)
			c.files = append(c.files, &shared.FileEntry{
				Path:  indexPath,
				Title: title,
				Order: order,
			})
			order++

			for _, sect := range sections {
				sectID := sect.GetAttr("id")
				if sectID == "" {
					continue
				}
				sectSlug := shared.Slugify(sectID)
				sectPath := dirSlug + "/" + sectSlug + ".md"
				sectTitle := extractTitle(sect)
				sectTitle = substituteVariables(sectTitle, c.vars)

				c.files = append(c.files, &shared.FileEntry{
					Path:      sectPath,
					Title:     sectTitle,
					NavParent: dirSlug,
					Order:     order,
				})
				order++

				c.idMap[sectID] = &shared.IDEntry{
					File:  sectPath,
					Title: sectTitle,
				}
				c.mapSectionIDs(sect, sectPath)
			}
		}
	}
}

// mapSectionIDs recursively maps section IDs to output locations.
func (c *Converter) mapSectionIDs(node *sgml.Node, filePath string) {
	for _, sect := range node.FindDescendants("section") {
		id := sect.GetAttr("id")
		if id != "" {
			title := extractTitle(sect)
			c.idMap[id] = &shared.IDEntry{
				File:   filePath,
				Anchor: id,
				Title:  substituteVariables(title, c.vars),
			}
		}
	}
}

// convertDocument converts a single parsed document to Markdown.
func (c *Converter) convertDocument(key string, doc *sgml.Node) error {
	sections := doc.FindChildren("section")

	singlePage := len(sections) <= 1 ||
		key == "faq" || key == "metric" ||
		key == "coding" || key == "contributing" ||
		key == "documentation" || key == "release" ||
		key == "index"

	if singlePage {
		return c.convertSinglePage(key, doc)
	}
	return c.convertMultiPage(key, doc)
}

// convertSinglePage renders a document as a single Markdown file.
func (c *Converter) convertSinglePage(key string, doc *sgml.Node) error {
	outPath := c.pageMap[key]
	c.currentFile = outPath
	w := shared.NewMarkdownWriter()

	title := doc.GetAttr("title")
	title = substituteVariables(title, c.vars)
	if title != "" {
		w.Heading(1, title, "")
	}

	subtitle := doc.GetAttr("subtitle")
	subtitle = substituteVariables(subtitle, c.vars)
	if subtitle != "" {
		w.BlankLine()
		w.Write("*" + subtitle + "*")
		w.BlankLine()
	}

	// Handle release.xml specially
	if key == "release" {
		c.convertRelease(doc, w)
	} else {
		c.convertChildren(doc, w, 1)
	}

	return c.writeFile(outPath, w.String())
}

// convertMultiPage splits a document into multiple Markdown files.
func (c *Converter) convertMultiPage(key string, doc *sgml.Node) error {
	dirSlug := shared.Slugify(key)

	// Write index page with description/intro
	indexW := shared.NewMarkdownWriter()
	title := doc.GetAttr("title")
	title = substituteVariables(title, c.vars)
	if title != "" {
		indexW.Heading(1, title, "")
	}

	// Render description
	if desc := doc.FindChild("description"); desc != nil {
		indexW.BlankLine()
		text := substituteVariables(desc.TextContent(), c.vars)
		indexW.Write(text)
		indexW.BlankLine()
	}

	// Render non-section content (intro, etc.)
	for _, child := range doc.Children {
		if child.Type != sgml.ElementNode {
			continue
		}
		if child.Tag == "section" || child.Tag == "variable-list" ||
			child.Tag == "description" || child.Tag == "block-define" ||
			child.Tag == "host-define" || child.Tag == "cleanup" {
			continue
		}
		c.convertNode(child, indexW, 1)
	}

	indexPath := dirSlug + "/index.md"
	c.currentFile = indexPath
	if err := c.writeFile(indexPath, indexW.String()); err != nil {
		return err
	}

	// Write each section as a separate file
	for _, sect := range doc.FindChildren("section") {
		sectID := sect.GetAttr("id")
		if sectID == "" {
			continue
		}
		sectSlug := shared.Slugify(sectID)
		sectPath := dirSlug + "/" + sectSlug + ".md"
		c.currentFile = sectPath

		sectW := shared.NewMarkdownWriter()
		sectTitle := extractTitle(sect)
		sectTitle = substituteVariables(sectTitle, c.vars)
		if sectTitle != "" {
			sectW.Heading(1, sectTitle, "")
		}

		// Add anchor for the section ID
		if sectID != "" {
			sectW.WriteString(fmt.Sprintf(
				"<a name=\"%s\"></a>\n", sectID))
		}

		c.convertSectionChildren(sect, sectW, 1)

		if err := c.writeFile(sectPath, sectW.String()); err != nil {
			return err
		}
	}

	return nil
}

// convertRelease handles release.xml with its custom structure.
func (c *Converter) convertRelease(doc *sgml.Node, w *shared.MarkdownWriter) {
	releaseList := doc.FindChild("release-list")
	if releaseList == nil {
		c.convertChildren(doc, w, 1)
		return
	}

	for _, rel := range releaseList.FindChildren("release") {
		version := rel.GetAttr("version")
		date := rel.GetAttr("date")
		relTitle := rel.GetAttr("title")

		heading := fmt.Sprintf("v%s", version)
		if relTitle != "" {
			heading += " — " + relTitle
		}
		w.Heading(2, heading, "")

		if date != "" && date != "XXXX-XX-XX" {
			w.BlankLine()
			w.Write("*Released: " + date + "*")
			w.BlankLine()
		}

		// Core changes
		c.convertReleaseSection(rel.FindChild("release-core-list"),
			"Core", w)
		// Doc changes
		c.convertReleaseSection(rel.FindChild("release-doc-list"),
			"Documentation", w)
		// Test changes
		c.convertReleaseSection(rel.FindChild("release-test-list"),
			"Test", w)
	}
}

// convertReleaseSection renders a release section (bugs, features, etc.).
func (c *Converter) convertReleaseSection(
	node *sgml.Node, label string, w *shared.MarkdownWriter,
) {
	if node == nil {
		return
	}

	subsections := []struct {
		tag   string
		title string
	}{
		{"release-bug-list", "Bug Fixes"},
		{"release-feature-list", "Features"},
		{"release-improvement-list", "Improvements"},
		{"release-development-list", "Development"},
	}

	for _, sub := range subsections {
		list := node.FindChild(sub.tag)
		if list == nil {
			continue
		}
		w.Heading(3, label+" "+sub.title, "")

		for _, item := range list.FindChildren("release-item") {
			w.BlankLine()
			w.Write("- ")
			for _, p := range item.FindChildren("p") {
				c.convertInlineContent(p, w)
				w.Write(" ")
			}
			w.Newline()
		}
	}
}

// convertChildren converts all children of a node.
func (c *Converter) convertChildren(
	node *sgml.Node, w *shared.MarkdownWriter, depth int,
) {
	for _, child := range node.Children {
		c.convertNode(child, w, depth)
	}
}

// convertSectionChildren converts children of a section,
// excluding the title (which is handled by the caller).
func (c *Converter) convertSectionChildren(
	node *sgml.Node, w *shared.MarkdownWriter, depth int,
) {
	for _, child := range node.Children {
		if child.Type == sgml.ElementNode && child.Tag == "title" {
			continue
		}
		c.convertNode(child, w, depth)
	}
}

// convertNode dispatches conversion for a single node.
func (c *Converter) convertNode(
	node *sgml.Node, w *shared.MarkdownWriter, depth int,
) {
	if node.Type == sgml.TextNode {
		text := substituteVariables(node.Text, c.vars)
		// Skip whitespace-only text nodes (XML indentation)
		if strings.TrimSpace(text) == "" {
			return
		}
		w.Write(text)
		return
	}

	if node.Type != sgml.ElementNode {
		return
	}

	switch node.Tag {
	case "section":
		c.handleSection(node, w, depth)
	case "title":
		// handled by parent
	case "p":
		c.handleParagraph(node, w)
	case "text":
		c.convertChildren(node, w, depth)
	case "list":
		c.handleList(node, w, depth)
	case "admonition":
		c.handleAdmonition(node, w, depth)
	case "code-block":
		c.handleCodeBlock(node, w)
	case "table":
		c.handleTable(node, w)
	case "execute-list":
		c.handleExecuteList(node, w)
	case "backrest-config":
		c.handleBackrestConfig(node, w)
	case "postgres-config":
		c.handlePostgresConfig(node, w)
	case "host-add", "host-define", "cleanup":
		// skip infrastructure
	case "variable-list":
		// already processed
	case "block-define":
		// already collected
	case "block":
		c.handleBlock(node, w, depth)
	case "description":
		text := substituteVariables(node.TextContent(), c.vars)
		w.BlankLine()
		w.Write(text)
		w.BlankLine()
	case "intro":
		c.convertChildren(node, w, depth)

	// Command/config document elements
	case "operation":
		c.handleOperation(node, w)
	case "config":
		c.handleConfig(node, w)
	case "command-list":
		c.handleCommandList(node, w)
	case "config-section-list":
		c.handleConfigSectionList(node, w)

	// Release elements handled in convertRelease
	case "release-list", "contributor-list":
		// handled specially

	// Self-closing brand elements
	case "backrest":
		w.Write("pgBackRest")
	case "postgres":
		w.Write("PostgreSQL")
	case "exe":
		w.Write("`" + c.vars["project-exe"] + "`")

	// Inline formatting/semantic elements
	case "b", "i", "bi", "code", "quote", "proper", "id",
		"file", "path", "cmd", "host", "user", "setting",
		"br-option", "br-setting", "pg-option", "pg-setting",
		"link", "br", "option-description", "cmd-description",
		"summary":
		c.convertInlineElement(node, w)

	// Inline elements
	default:
		c.convertInlineContent(node, w)
	}
}

// handleSection converts a <section> to a Markdown heading.
func (c *Converter) handleSection(
	node *sgml.Node, w *shared.MarkdownWriter, depth int,
) {
	title := extractTitle(node)
	title = substituteVariables(title, c.vars)
	id := node.GetAttr("id")

	level := depth + 1
	if level > 6 {
		level = 6
	}
	w.Heading(level, title, "")

	if id != "" {
		w.WriteString(fmt.Sprintf("<a name=\"%s\"></a>\n", id))
	}

	c.convertSectionChildren(node, w, depth+1)
}

// handleParagraph converts <p> to a Markdown paragraph.
func (c *Converter) handleParagraph(
	node *sgml.Node, w *shared.MarkdownWriter,
) {
	w.BlankLine()
	c.convertInlineContent(node, w)
	w.BlankLine()
}

// handleList converts <list> to Markdown bullet list.
func (c *Converter) handleList(
	node *sgml.Node, w *shared.MarkdownWriter, depth int,
) {
	w.BlankLine()
	for _, item := range node.FindChildren("list-item") {
		w.Write("- ")
		w.PushIndent("  ")
		c.convertChildren(item, w, depth)
		w.PopIndent("  ")
		w.EnsureNewline()
	}
}

// handleAdmonition converts <admonition> to MkDocs admonition.
func (c *Converter) handleAdmonition(
	node *sgml.Node, w *shared.MarkdownWriter, depth int,
) {
	kind := node.GetAttr("type")
	if kind == "" {
		kind = "note"
	}
	w.Admonition(kind)
	w.PushIndent("    ")
	c.convertChildren(node, w, depth)
	w.PopIndent("    ")
}

// handleCodeBlock converts <code-block> to fenced code block.
func (c *Converter) handleCodeBlock(
	node *sgml.Node, w *shared.MarkdownWriter,
) {
	lang := node.GetAttr("type")
	title := node.GetAttr("title")

	if title != "" {
		w.BlankLine()
		w.Write("**" + substituteVariables(title, c.vars) + "**")
	}

	w.StartCodeBlock(lang)
	text := substituteVariables(node.TextContent(), c.vars)
	w.WriteString(text)
	w.EndCodeBlock()
}

// handleTable converts <table> to Markdown table.
func (c *Converter) handleTable(
	node *sgml.Node, w *shared.MarkdownWriter,
) {
	// Table title
	if title := node.FindChild("title"); title != nil {
		w.BlankLine()
		text := substituteVariables(title.TextContent(), c.vars)
		w.Write("**" + strings.TrimSpace(text) + "**")
		w.BlankLine()
	}

	header := node.FindChild("table-header")
	data := node.FindChild("table-data")
	if data == nil {
		return
	}

	w.BlankLine()

	// Determine column count
	numCols := 0
	if header != nil {
		numCols = len(header.FindChildren("table-column"))
	} else if rows := data.FindChildren("table-row"); len(rows) > 0 {
		numCols = len(rows[0].FindChildren("table-cell"))
	}

	// Header row
	if header != nil {
		w.WriteString("|")
		for _, col := range header.FindChildren("table-column") {
			text := c.inlineToString(col)
			w.WriteString(" " + text + " |")
		}
		w.Newline()
	} else {
		w.WriteString("|")
		for i := 0; i < numCols; i++ {
			w.WriteString("  |")
		}
		w.Newline()
	}

	// Separator
	w.WriteString("|")
	for i := 0; i < numCols; i++ {
		w.WriteString(" --- |")
	}
	w.Newline()

	// Data rows
	for _, row := range data.FindChildren("table-row") {
		w.WriteString("|")
		for _, cell := range row.FindChildren("table-cell") {
			text := c.inlineToString(cell)
			text = strings.ReplaceAll(text, "\n", " ")
			w.WriteString(" " + text + " |")
		}
		w.Newline()
	}
	w.BlankLine()
}

// handleExecuteList renders <execute-list> as a code block.
func (c *Converter) handleExecuteList(
	node *sgml.Node, w *shared.MarkdownWriter,
) {
	title := extractTitle(node)
	title = substituteVariables(title, c.vars)
	show := node.GetAttr("show")

	// Skip hidden execute lists
	if show == "n" {
		return
	}

	if title != "" {
		w.BlankLine()
		w.Write("**" + title + "**")
	}

	for _, exec := range node.FindChildren("execute") {
		execShow := exec.GetAttr("show")
		if execShow == "n" {
			continue
		}
		cmd := exec.FindChild("exe-cmd")
		if cmd == nil {
			continue
		}

		cmdText := substituteVariables(
			strings.TrimSpace(cmd.TextContent()), c.vars)
		if cmdText == "" {
			continue
		}

		w.StartCodeBlock("bash")
		w.WriteString(cmdText + "\n")
		w.EndCodeBlock()

		// Show output if present and visible
		output := exec.FindChild("exe-output")
		if output != nil && exec.GetAttr("output") != "n" {
			w.StartCodeBlock("")
			w.WriteString(strings.TrimSpace(
				output.TextContent()) + "\n")
			w.EndCodeBlock()
		}
	}
}

// handleBackrestConfig renders a pgBackRest config block.
func (c *Converter) handleBackrestConfig(
	node *sgml.Node, w *shared.MarkdownWriter,
) {
	show := node.GetAttr("show")
	if show == "n" {
		return
	}

	title := extractTitle(node)
	title = substituteVariables(title, c.vars)
	if title != "" {
		w.BlankLine()
		w.Write("**" + title + "**")
	}

	w.StartCodeBlock("ini")
	for _, opt := range node.FindChildren("backrest-config-option") {
		section := opt.GetAttr("section")
		key := opt.GetAttr("key")
		val := strings.TrimSpace(opt.TextContent())
		val = substituteVariables(val, c.vars)

		if section != "" {
			w.WriteString("[" + section + "]\n")
		}
		w.WriteString(key + "=" + val + "\n")
	}
	w.EndCodeBlock()
}

// handlePostgresConfig renders a PostgreSQL config block.
func (c *Converter) handlePostgresConfig(
	node *sgml.Node, w *shared.MarkdownWriter,
) {
	show := node.GetAttr("show")
	if show == "n" {
		return
	}

	title := extractTitle(node)
	title = substituteVariables(title, c.vars)
	if title != "" {
		w.BlankLine()
		w.Write("**" + title + "**")
	}

	w.StartCodeBlock("ini")
	for _, opt := range node.FindChildren("postgres-config-option") {
		key := opt.GetAttr("key")
		val := strings.TrimSpace(opt.TextContent())
		val = substituteVariables(val, c.vars)
		w.WriteString(key + " = " + val + "\n")
	}
	w.EndCodeBlock()
}

// handleBlock inlines a previously-defined block.
func (c *Converter) handleBlock(
	node *sgml.Node, w *shared.MarkdownWriter, depth int,
) {
	id := node.GetAttr("id")
	block, ok := c.blocks[id]
	if !ok {
		c.warn("block %q not found", id)
		return
	}
	c.convertChildren(block, w, depth)
}

// handleOperation converts a command operation document.
func (c *Converter) handleOperation(
	node *sgml.Node, w *shared.MarkdownWriter,
) {
	title := node.GetAttr("title")
	title = substituteVariables(title, c.vars)
	w.Heading(1, title, "")

	if desc := node.FindChild("description"); desc != nil {
		w.BlankLine()
		w.Write(substituteVariables(desc.TextContent(), c.vars))
		w.BlankLine()
	}

	if text := node.FindChild("text"); text != nil {
		c.convertChildren(text, w, 1)
	}

	// General options
	if gen := node.FindChild("operation-general"); gen != nil {
		genTitle := gen.GetAttr("title")
		genTitle = substituteVariables(genTitle, c.vars)
		w.Heading(2, genTitle, "")
		c.convertOptionList(gen.FindChild("option-list"), w)
	}

	// Command list
	if cmdList := node.FindChild("command-list"); cmdList != nil {
		c.handleCommandList(cmdList, w)
	}
}

// handleCommandList converts a list of commands.
func (c *Converter) handleCommandList(
	node *sgml.Node, w *shared.MarkdownWriter,
) {
	title := node.GetAttr("title")
	title = substituteVariables(title, c.vars)
	if title != "" {
		w.Heading(2, title, "")
	}

	if text := node.FindChild("text"); text != nil {
		c.convertChildren(text, w, 2)
	}

	for _, cmd := range node.FindChildren("command") {
		name := cmd.GetAttr("name")
		w.Heading(3, name, "")

		if summary := cmd.FindChild("summary"); summary != nil {
			w.BlankLine()
			c.convertInlineContent(summary, w)
			w.BlankLine()
		}

		if text := cmd.FindChild("text"); text != nil {
			c.convertChildren(text, w, 3)
		}

		if optList := cmd.FindChild("option-list"); optList != nil {
			c.convertOptionList(optList, w)
		}
	}
}

// handleConfig converts a configuration document.
func (c *Converter) handleConfig(
	node *sgml.Node, w *shared.MarkdownWriter,
) {
	title := node.GetAttr("title")
	title = substituteVariables(title, c.vars)
	w.Heading(1, title, "")

	if desc := node.FindChild("description"); desc != nil {
		w.BlankLine()
		w.Write(substituteVariables(desc.TextContent(), c.vars))
		w.BlankLine()
	}

	if text := node.FindChild("text"); text != nil {
		c.convertChildren(text, w, 1)
	}

	if secList := node.FindChild("config-section-list"); secList != nil {
		c.handleConfigSectionList(secList, w)
	}
}

// handleConfigSectionList converts configuration sections.
func (c *Converter) handleConfigSectionList(
	node *sgml.Node, w *shared.MarkdownWriter,
) {
	title := node.GetAttr("title")
	title = substituteVariables(title, c.vars)
	if title != "" {
		w.Heading(2, title, "")
	}

	for _, sect := range node.FindChildren("config-section") {
		name := sect.GetAttr("name")
		w.Heading(3, name, "")

		if text := sect.FindChild("text"); text != nil {
			c.convertChildren(text, w, 3)
		}

		if keyList := sect.FindChild("config-key-list"); keyList != nil {
			c.convertConfigKeyList(keyList, w)
		}
	}
}

// convertConfigKeyList converts a list of config keys.
func (c *Converter) convertConfigKeyList(
	node *sgml.Node, w *shared.MarkdownWriter,
) {
	for _, key := range node.FindChildren("config-key") {
		name := key.GetAttr("name")
		w.Heading(4, name, "")

		if summary := key.FindChild("summary"); summary != nil {
			w.BlankLine()
			c.convertInlineContent(summary, w)
			w.BlankLine()
		}

		if text := key.FindChild("text"); text != nil {
			c.convertChildren(text, w, 4)
		}

		if def := key.FindChild("default"); def != nil {
			defText := strings.TrimSpace(def.TextContent())
			if defText != "" {
				w.BlankLine()
				w.Write("**Default:** `" + defText + "`")
				w.BlankLine()
			}
		}

		if allow := key.FindChild("allow"); allow != nil {
			allowText := strings.TrimSpace(allow.TextContent())
			if allowText != "" {
				w.BlankLine()
				w.Write("**Allowed:** " + allowText)
				w.BlankLine()
			}
		}

		examples := key.FindChildren("example")
		if len(examples) > 0 {
			w.BlankLine()
			w.Write("**Example" +
				pluralS(len(examples)) + ":** ")
			for i, ex := range examples {
				if i > 0 {
					w.Write(", ")
				}
				w.Write("`" + strings.TrimSpace(
					ex.TextContent()) + "`")
			}
			w.BlankLine()
		}
	}
}

// convertOptionList converts an <option-list>.
func (c *Converter) convertOptionList(
	node *sgml.Node, w *shared.MarkdownWriter,
) {
	if node == nil {
		return
	}
	for _, opt := range node.FindChildren("option") {
		name := opt.GetAttr("name")
		w.Heading(4, name, "")

		if summary := opt.FindChild("summary"); summary != nil {
			w.BlankLine()
			c.convertInlineContent(summary, w)
			w.BlankLine()
		}

		if text := opt.FindChild("text"); text != nil {
			c.convertChildren(text, w, 4)
		}

		examples := opt.FindChildren("example")
		if len(examples) > 0 {
			w.BlankLine()
			w.Write("**Example" +
				pluralS(len(examples)) + ":** ")
			for i, ex := range examples {
				if i > 0 {
					w.Write(", ")
				}
				w.Write("`" + strings.TrimSpace(
					ex.TextContent()) + "`")
			}
			w.BlankLine()
		}
	}
}

// convertInlineElement converts a single inline element node.
func (c *Converter) convertInlineElement(
	node *sgml.Node, w *shared.MarkdownWriter,
) {
	switch node.Tag {
	case "b":
		w.Write("**")
		c.convertInlineContent(node, w)
		w.Write("**")
	case "i":
		w.Write("*")
		c.convertInlineContent(node, w)
		w.Write("*")
	case "bi":
		w.Write("***")
		c.convertInlineContent(node, w)
		w.Write("***")
	case "code":
		text := substituteVariables(node.TextContent(), c.vars)
		w.Write("`" + text + "`")
	case "quote":
		w.Write("\"")
		c.convertInlineContent(node, w)
		w.Write("\"")
	case "proper":
		c.convertInlineContent(node, w)
	case "id":
		text := substituteVariables(node.TextContent(), c.vars)
		w.Write("`" + text + "`")
	case "file", "path", "cmd", "host", "user",
		"setting", "br-option", "br-setting",
		"pg-option", "pg-setting":
		text := substituteVariables(node.TextContent(), c.vars)
		w.Write("`" + text + "`")
	case "link":
		c.handleLink(node, w)
	case "br":
		w.Newline()
	case "option-description", "cmd-description":
		key := node.GetAttr("key")
		if key != "" {
			w.Write("`" + key + "`")
		}
	case "summary":
		c.convertInlineContent(node, w)
	default:
		c.convertInlineContent(node, w)
	}
}

// convertInlineContent converts the inline content of a node.
func (c *Converter) convertInlineContent(
	node *sgml.Node, w *shared.MarkdownWriter,
) {
	for _, child := range node.Children {
		if child.Type == sgml.TextNode {
			text := substituteVariables(child.Text, c.vars)
			// Normalize whitespace in inline context
			text = collapseWhitespace(text)
			if text != "" {
				w.Write(text)
			}
			continue
		}
		if child.Type != sgml.ElementNode {
			continue
		}

		switch child.Tag {
		// Self-closing brand elements
		case "backrest":
			w.Write("pgBackRest")
		case "postgres":
			w.Write("PostgreSQL")
		case "exe":
			w.Write("`" + c.vars["project-exe"] + "`")

		// Formatting
		case "b":
			w.Write("**")
			c.convertInlineContent(child, w)
			w.Write("**")
		case "i":
			w.Write("*")
			c.convertInlineContent(child, w)
			w.Write("*")
		case "bi":
			w.Write("***")
			c.convertInlineContent(child, w)
			w.Write("***")
		case "code":
			text := substituteVariables(
				child.TextContent(), c.vars)
			w.Write("`" + text + "`")
		case "quote":
			w.Write("\"")
			c.convertInlineContent(child, w)
			w.Write("\"")
		case "proper":
			c.convertInlineContent(child, w)
		case "id":
			text := substituteVariables(
				child.TextContent(), c.vars)
			w.Write("`" + text + "`")

		// Semantic code elements
		case "file", "path", "cmd", "host", "user",
			"setting", "br-option", "br-setting",
			"pg-option", "pg-setting":
			text := substituteVariables(
				child.TextContent(), c.vars)
			w.Write("`" + text + "`")

		// Links
		case "link":
			c.handleLink(child, w)

		// Line break
		case "br":
			w.Newline()

		// Nested block elements that can appear inline
		case "list":
			c.handleList(child, w, 1)
		case "code-block":
			c.handleCodeBlock(child, w)
		case "p":
			c.handleParagraph(child, w)

		// Option/command descriptions (auto-generated refs)
		case "option-description", "cmd-description":
			// These reference auto-generated content;
			// skip in static conversion
			key := child.GetAttr("key")
			if key != "" {
				w.Write("`" + key + "`")
			}

		default:
			// Unknown element — render children
			c.convertInlineContent(child, w)
		}
	}
}

// handleLink converts <link> to Markdown link.
func (c *Converter) handleLink(
	node *sgml.Node, w *shared.MarkdownWriter,
) {
	url := substituteVariables(node.GetAttr("url"), c.vars)
	page := substituteVariables(node.GetAttr("page"), c.vars)
	section := substituteVariables(node.GetAttr("section"), c.vars)

	text := c.inlineToString(node)
	text = substituteVariables(text, c.vars)

	if url != "" {
		// Convert relative .html links to pgbackrest.org
		if !strings.Contains(url, "://") &&
			!strings.HasPrefix(url, "/") &&
			(strings.HasSuffix(url, ".html") ||
				strings.Contains(url, ".html#")) {
			url = "https://pgbackrest.org/" + url
		}
		w.Write("[" + text + "](" + url + ")")
		return
	}

	if page != "" {
		target := c.resolvePageLink(page, section)
		if text == "" {
			text = page
		}
		w.Write("[" + text + "](" + target + ")")
		return
	}

	if section != "" {
		target := c.resolveSectionLink(section)
		if text == "" {
			text = section
		}
		w.Write("[" + text + "](" + target + ")")
		return
	}

	// No link target — just emit text
	w.Write(text)
}

// relativeLink computes a relative link from the current file
// to the target file path, appending the fragment if non-empty.
func (c *Converter) relativeLink(target, fragment string) string {
	if target == "" {
		if fragment != "" {
			return "#" + fragment
		}
		return ""
	}
	curDir := filepath.Dir(c.currentFile)
	rel, err := filepath.Rel(curDir, target)
	if err != nil {
		rel = target
	}
	if fragment != "" {
		return rel + "#" + fragment
	}
	return rel
}

// resolvePageLink resolves a page reference to an output path.
func (c *Converter) resolvePageLink(page, section string) string {
	target := c.pageMap[page]
	if target == "" {
		// Try as an external pgbackrest.org link
		target = "https://pgbackrest.org/" + page + ".html"
		if section != "" {
			target += "#" + strings.ReplaceAll(section, "/", "-")
		}
		return target
	}
	if section != "" {
		// Section may use / separators (e.g. "quickstart/configure-stanza")
		// Try to find the subsection as a separate page first
		parts := strings.SplitN(section, "/", 2)
		if len(parts) == 2 {
			// Look up the subsection ID in the idMap
			if entry, ok := c.idMap[parts[1]]; ok {
				return c.relativeLink(entry.File, parts[1])
			}
			if entry, ok := c.idMap[parts[0]]; ok {
				return c.relativeLink(entry.File, parts[1])
			}
		}
		// Try the full section as an anchor
		flat := strings.ReplaceAll(section, "/", "-")
		if entry, ok := c.idMap[flat]; ok {
			return c.relativeLink(entry.File, flat)
		}
		target += "#" + strings.ReplaceAll(section, "/", "-")
	}
	return c.relativeLink(target, "")
}

// resolveSectionLink resolves a section-only link.
// Sections may use / path format (e.g. "/quickstart/perform-restore").
func (c *Converter) resolveSectionLink(section string) string {
	// Strip leading /
	section = strings.TrimPrefix(section, "/")

	// Try exact match first
	if entry, ok := c.idMap[section]; ok {
		if entry.File != "" {
			return c.relativeLink(entry.File, section)
		}
		return "#" + section
	}

	// Try splitting on / — "quickstart/perform-restore" means
	// section "perform-restore" within page "quickstart"
	parts := strings.SplitN(section, "/", 2)
	if len(parts) == 2 {
		if entry, ok := c.idMap[parts[1]]; ok {
			return c.relativeLink(entry.File, parts[1])
		}
		// Look up the parent section as a page
		if entry, ok := c.idMap[parts[0]]; ok {
			return c.relativeLink(entry.File, parts[1])
		}
	}

	// Fall back to same-page anchor
	return "#" + strings.ReplaceAll(section, "/", "-")
}

// inlineToString renders inline content to a plain string.
func (c *Converter) inlineToString(node *sgml.Node) string {
	w := shared.NewMarkdownWriter()
	c.convertInlineContent(node, w)
	return strings.TrimSpace(w.String())
}

// collectBlocks collects <block-define> elements into the block map.
func collectBlocks(node *sgml.Node, blocks map[string]*sgml.Node) {
	for _, bd := range node.FindDescendants("block-define") {
		id := bd.GetAttr("id")
		if id != "" {
			blocks[id] = bd
		}
	}
}

// docTitle returns a descriptive title for a <doc> element.
// When the title is just the project name, it falls back to
// the subtitle to give each page a distinct nav label.
func (c *Converter) docTitle(doc *sgml.Node) string {
	title := substituteVariables(doc.GetAttr("title"), c.vars)
	subtitle := substituteVariables(doc.GetAttr("subtitle"), c.vars)

	projectName := c.vars["project"]
	if title == projectName && subtitle != "" {
		return subtitle
	}
	return title
}

// extractTitle returns the text content of a node's <title> child.
func extractTitle(node *sgml.Node) string {
	if title := node.FindChild("title"); title != nil {
		return strings.TrimSpace(title.TextContent())
	}
	return ""
}

// writeFile writes content to an output file.
func (c *Converter) writeFile(relPath, content string) error {
	fullPath := filepath.Join(c.outDir, relPath)
	dir := filepath.Dir(fullPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("creating dir %s: %w", dir, err)
	}
	content = strings.TrimSpace(content) + "\n"
	return os.WriteFile(fullPath, []byte(content), 0644)
}

// warn adds a warning message.
func (c *Converter) warn(format string, args ...interface{}) {
	c.warnings = append(c.warnings,
		fmt.Sprintf(format, args...))
}

// collapseWhitespace normalizes runs of whitespace to a single
// space, preserving leading/trailing space if originally present.
func collapseWhitespace(s string) string {
	if s == "" {
		return ""
	}
	fields := strings.Fields(s)
	if len(fields) == 0 {
		// Pure whitespace — preserve as single space
		return " "
	}
	result := strings.Join(fields, " ")
	// Preserve a leading space if original had one
	if s[0] == ' ' || s[0] == '\t' || s[0] == '\n' {
		result = " " + result
	}
	// Preserve a trailing space if original had one
	last := s[len(s)-1]
	if last == ' ' || last == '\t' || last == '\n' {
		result = result + " "
	}
	return result
}

// pluralS returns "s" if n != 1.
func pluralS(n int) string {
	if n == 1 {
		return ""
	}
	return "s"
}
