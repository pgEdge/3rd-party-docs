//-------------------------------------------------------------------------
//
// pgEdge PostgreSQL Docs
//
// Copyright (c) 2026, pgEdge, Inc.
// This software is released under The PostgreSQL License
//
//-------------------------------------------------------------------------

package convert

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/pgEdge/postgresql-docs/builder/sgml"
)

// Converter orchestrates the two-pass SGML-to-Markdown conversion.
type Converter struct {
	ctx *Context
}

// NewConverter creates a converter for the given parsed document.
func NewConverter(root *sgml.Node, srcDir, outDir, version string) *Converter {
	return &Converter{
		ctx: NewContext(root, srcDir, outDir, version),
	}
}

// Convert runs both passes and writes the output files.
func (c *Converter) Convert() error {
	// Pass 1: Build ID map and file structure
	c.buildIDMap(c.ctx.docRoot)

	// Pass 2: Convert and write files
	if err := c.convertAndWrite(); err != nil {
		return fmt.Errorf("conversion pass 2: %w", err)
	}

	return nil
}

// Warnings returns accumulated warnings.
func (c *Converter) Warnings() []string {
	return c.ctx.Warnings
}

// Files returns the list of generated files.
func (c *Converter) Files() []*FileEntry {
	return c.ctx.Files
}

// Context returns the conversion context (for nav generation).
func (c *Converter) Context() *Context {
	return c.ctx
}

// buildIDMap walks the entire document tree (Pass 1), recording
// every element with an id attribute and determining which output
// file it belongs to.
func (c *Converter) buildIDMap(root *sgml.Node) {
	// First, determine the file structure by walking top-level elements
	book := root.FindChild("book")
	if book == nil {
		// Try treating root children directly
		for _, child := range root.Children {
			if child.Type == sgml.ElementNode {
				c.mapElement(child, "")
			}
		}
		return
	}

	bookID := book.GetAttr("id")
	if bookID != "" {
		c.ctx.RegisterID(bookID, "index.md", "", "PostgreSQL Documentation", "book")
	}

	c.mapBookChildren(book, "")
}

// mapBookChildren maps children of the book element to files.
func (c *Converter) mapBookChildren(book *sgml.Node, parentPath string) {
	for _, child := range book.Children {
		if child.Type != sgml.ElementNode {
			continue
		}
		c.mapElement(child, parentPath)
	}
}

// mapElement determines the output file for an element and records
// its ID and all descendant IDs.
func (c *Converter) mapElement(node *sgml.Node, parentPath string) {
	switch node.Tag {
	case "part":
		c.mapPart(node, parentPath)
	case "reference":
		c.mapReference(node, parentPath)
	case "chapter", "appendix", "preface", "bibliography":
		c.mapChapter(node, parentPath)
	case "refentry":
		c.mapRefentry(node, parentPath)
	case "sect1", "section":
		c.mapSection(node, parentPath)
	default:
		// Record ID if present, assigned to parent's file
		c.registerNodeID(node, parentPath)
		// Recurse into children
		for _, child := range node.Children {
			if child.Type == sgml.ElementNode {
				c.mapElement(child, parentPath)
			}
		}
	}
}

// mapPart maps a <part> and its children.
func (c *Converter) mapPart(node *sgml.Node, parentPath string) {
	id := node.GetAttr("id")
	title := extractTitle(node)
	slug := slugify(title)
	if slug == "" && id != "" {
		slug = slugify(id)
	}
	if slug == "" {
		slug = "part"
	}

	partPath := filepath.Join(parentPath, slug)
	filePath := filepath.Join(partPath, "index.md")

	if id != "" {
		c.ctx.RegisterID(id, filePath, id, title, "part")
	}

	c.ctx.AddFile(filePath, title, parentPath)

	// Map part children
	for _, child := range node.Children {
		if child.Type == sgml.ElementNode {
			switch child.Tag {
			case "chapter", "appendix", "bibliography":
				c.mapChapter(child, partPath)
			case "reference":
				c.mapReference(child, partPath)
			case "refentry":
				c.mapRefentry(child, partPath)
			case "partintro":
				c.registerDescendantIDs(child, filePath)
			default:
				c.mapElement(child, partPath)
			}
		}
	}
}

// mapReference maps a <reference> element (like a chapter containing
// refentries, used in the SQL/app command reference sections).
func (c *Converter) mapReference(node *sgml.Node, parentPath string) {
	id := node.GetAttr("id")
	title := extractTitle(node)
	slug := slugify(title)
	if slug == "" && id != "" {
		slug = slugify(id)
	}
	if slug == "" {
		slug = "reference"
	}

	refDir := filepath.Join(parentPath, slug)
	indexFile := filepath.Join(refDir, "index.md")

	if id != "" {
		c.ctx.RegisterID(id, indexFile, id, title, "reference")
	}

	c.ctx.AddFile(indexFile, title, parentPath)

	// Register partintro IDs to index file
	partintro := node.FindChild("partintro")
	if partintro != nil {
		c.registerDescendantIDs(partintro, indexFile)
	}

	// Map refentry children
	for _, child := range node.Children {
		if child.Type == sgml.ElementNode && child.Tag == "refentry" {
			c.mapRefentry(child, refDir)
		}
	}
}

// mapChapter maps a <chapter> or <appendix> to a directory with
// an index.md, and its sect1 children to individual files.
func (c *Converter) mapChapter(node *sgml.Node, parentPath string) {
	id := node.GetAttr("id")
	title := extractTitle(node)
	slug := slugify(title)
	if slug == "" && id != "" {
		slug = slugify(id)
	}
	if slug == "" {
		slug = "chapter"
	}

	chapterDir := filepath.Join(parentPath, slug)
	indexFile := filepath.Join(chapterDir, "index.md")

	if id != "" {
		c.ctx.RegisterID(id, indexFile, id, title, node.Tag)
	}

	// Check if chapter has sect1 children
	sect1s := node.FindChildren("sect1")
	if len(sect1s) == 0 {
		sect1s = node.FindChildren("section")
	}

	if len(sect1s) == 0 {
		// No sections — everything goes in a single file
		chapterFile := filepath.Join(parentPath, slug+".md")
		if id != "" {
			c.ctx.IDMap[id].File = chapterFile
		}
		c.ctx.AddFile(chapterFile, title, parentPath)
		c.registerDescendantIDs(node, chapterFile)
		return
	}

	// Chapter with sections — create directory
	c.ctx.AddFile(indexFile, title, parentPath)

	// Register IDs on non-sect1 children to the index file
	for _, child := range node.Children {
		if child.Type == sgml.ElementNode &&
			child.Tag != "sect1" && child.Tag != "section" {
			c.registerDescendantIDs(child, indexFile)
		}
	}

	// Map sect1 children to individual files
	for _, sect := range sect1s {
		c.mapSection(sect, chapterDir)
	}

	// Map any refentry children
	for _, ref := range node.FindChildren("refentry") {
		c.mapRefentry(ref, chapterDir)
	}
}

// mapSection maps a <sect1>/<section> to its own file within a
// chapter directory.
func (c *Converter) mapSection(node *sgml.Node, parentPath string) {
	id := node.GetAttr("id")
	title := extractTitle(node)
	slug := slugify(title)
	if slug == "" && id != "" {
		slug = slugify(id)
	}
	if slug == "" {
		slug = "section"
	}

	filePath := filepath.Join(parentPath, slug+".md")

	if id != "" {
		c.ctx.RegisterID(id, filePath, id, title, node.Tag)
	}

	c.ctx.AddFile(filePath, title, parentPath)

	// Register all descendant IDs to this file
	c.registerDescendantIDs(node, filePath)
}

// mapRefentry maps a <refentry> to its own file.
func (c *Converter) mapRefentry(node *sgml.Node, parentPath string) {
	id := node.GetAttr("id")
	title := ""

	refmeta := node.FindChild("refmeta")
	if refmeta != nil {
		titleElem := refmeta.FindChild("refentrytitle")
		if titleElem != nil {
			title = titleElem.TextContent()
		}
	}
	if title == "" {
		refnamediv := node.FindChild("refnamediv")
		if refnamediv != nil {
			nameElem := refnamediv.FindChild("refname")
			if nameElem != nil {
				title = nameElem.TextContent()
			}
		}
	}

	slug := slugify(title)
	if slug == "" && id != "" {
		slug = slugify(id)
	}
	if slug == "" {
		slug = "ref"
	}

	filePath := filepath.Join(parentPath, slug+".md")

	if id != "" {
		c.ctx.RegisterID(id, filePath, id, title, "refentry")
	}

	c.ctx.AddFile(filePath, title, parentPath)

	// Register all descendant IDs
	c.registerDescendantIDs(node, filePath)
}

// registerNodeID registers a single node's ID.
func (c *Converter) registerNodeID(node *sgml.Node, file string) {
	id := node.GetAttr("id")
	if id == "" {
		return
	}
	title := extractTitle(node)
	if title == "" {
		title = id
	}
	c.ctx.RegisterID(id, file, id, title, node.Tag)
}

// registerDescendantIDs registers IDs for a node and all its descendants.
func (c *Converter) registerDescendantIDs(node *sgml.Node, file string) {
	var walk func(*sgml.Node)
	walk = func(n *sgml.Node) {
		if n.Type != sgml.ElementNode {
			return
		}
		id := n.GetAttr("id")
		if id != "" {
			// Prefer xreflabel if set (concise cross-reference title)
			title := n.GetAttr("xreflabel")
			if title == "" {
				title = extractTitle(n)
			}
			if title == "" {
				// For varlistentry, extract term text
				// but skip indexterm content
				if n.Tag == "varlistentry" {
					terms := n.FindChildren("term")
					if len(terms) > 0 {
						title = textContentSkipping(
							terms[0], "indexterm")
					}
				}
			}
			if title == "" {
				title = id
			}
			c.ctx.RegisterID(id, file, id, title, n.Tag)
		}
		for _, child := range n.Children {
			walk(child)
		}
	}
	walk(node)
}

// convertAndWrite performs Pass 2: converts each mapped file's
// content to Markdown and writes it.
func (c *Converter) convertAndWrite() error {
	book := c.ctx.docRoot.FindChild("book")
	if book == nil {
		// Try converting root children directly
		return c.convertNodeToFiles(c.ctx.docRoot)
	}

	// Write book-level index.md
	if err := c.convertBook(book); err != nil {
		return err
	}

	return c.convertNodeToFiles(book)
}

// convertBook writes the top-level index.md for the book.
func (c *Converter) convertBook(book *sgml.Node) error {
	id := book.GetAttr("id")
	title := extractTitle(book)

	entry, ok := c.ctx.IDMap[id]
	if !ok {
		// No ID registered — create a default index.md
		entry = &IDEntry{File: "index.md"}
		c.ctx.AddFile("index.md", title, "")
	}

	c.ctx.CurrentFile = entry.File
	w := NewMarkdownWriter()
	w.Heading(1, title, id)

	// Convert bookinfo if present
	bookinfo := book.FindChild("bookinfo")
	if bookinfo != nil {
		// Extract author/copyright info
		corpauthor := bookinfo.FindChild("corpauthor")
		if corpauthor != nil {
			w.BlankLine()
			w.WriteString("*" + corpauthor.TextContent() + "*\n")
		}
	}

	// Add a brief intro listing the parts
	w.BlankLine()
	for _, child := range book.Children {
		if child.Type != sgml.ElementNode {
			continue
		}
		if child.Tag == "part" || child.Tag == "preface" {
			childID := child.GetAttr("id")
			childTitle := extractTitle(child)
			if childID != "" && childTitle != "" {
				link, _, ok := c.ctx.ResolveLink(childID)
				if ok {
					w.WriteString(fmt.Sprintf("- [%s](%s)\n", childTitle, link))
				}
			}
		}
	}

	return c.writeFile(entry.File, w.String())
}

// convertNodeToFiles walks the document tree and writes content
// to the appropriate output files based on the ID map.
func (c *Converter) convertNodeToFiles(node *sgml.Node) error {
	for _, child := range node.Children {
		if child.Type != sgml.ElementNode {
			continue
		}

		switch child.Tag {
		case "part":
			if err := c.convertPart(child); err != nil {
				return err
			}
		case "reference":
			if err := c.convertReference(child); err != nil {
				return err
			}
		case "chapter", "appendix", "preface", "bibliography":
			if err := c.convertChapter(child); err != nil {
				return err
			}
		case "refentry":
			if err := c.convertRefentry(child); err != nil {
				return err
			}
		}
	}
	return nil
}

// convertPart converts a <part> element.
func (c *Converter) convertPart(node *sgml.Node) error {
	id := node.GetAttr("id")
	title := extractTitle(node)

	entry, ok := c.ctx.IDMap[id]
	if !ok {
		return fmt.Errorf("part %q not found in ID map", id)
	}

	// Write part index page
	w := NewMarkdownWriter()
	w.Heading(1, title, id)

	// Convert partintro if present
	partintro := node.FindChild("partintro")
	if partintro != nil {
		c.ctx.CurrentFile = entry.File
		if err := convertChildren(c.ctx, partintro, w); err != nil {
			return err
		}
	}

	if err := c.writeFile(entry.File, w.String()); err != nil {
		return err
	}

	// Convert child chapters and references
	for _, child := range node.Children {
		if child.Type != sgml.ElementNode {
			continue
		}
		switch child.Tag {
		case "chapter", "appendix", "bibliography":
			if err := c.convertChapter(child); err != nil {
				return err
			}
		case "reference":
			if err := c.convertReference(child); err != nil {
				return err
			}
		case "refentry":
			if err := c.convertRefentry(child); err != nil {
				return err
			}
		}
	}

	return nil
}

// convertReference converts a <reference> element (group of refentries).
func (c *Converter) convertReference(node *sgml.Node) error {
	id := node.GetAttr("id")
	title := extractTitle(node)

	entry, ok := c.ctx.IDMap[id]
	if !ok {
		c.ctx.Warn("reference %q not found in ID map", id)
		return nil
	}

	// Write reference index page
	w := NewMarkdownWriter()
	w.Heading(1, title, id)

	// Convert partintro if present
	partintro := node.FindChild("partintro")
	if partintro != nil {
		c.ctx.CurrentFile = entry.File
		if err := convertChildren(c.ctx, partintro, w); err != nil {
			return err
		}
	}

	if err := c.writeFile(entry.File, w.String()); err != nil {
		return err
	}

	// Convert child refentries
	for _, child := range node.Children {
		if child.Type == sgml.ElementNode && child.Tag == "refentry" {
			if err := c.convertRefentry(child); err != nil {
				return err
			}
		}
	}

	return nil
}

// convertChapter converts a <chapter>/<appendix>/<preface> element.
func (c *Converter) convertChapter(node *sgml.Node) error {
	id := node.GetAttr("id")

	if id == "" {
		// Chapter without ID — skip (no output file mapped)
		c.ctx.Warn("skipping chapter without id at line %d", node.Line)
		return nil
	}

	entry, ok := c.ctx.IDMap[id]
	if !ok {
		return fmt.Errorf("chapter %q not found in ID map", id)
	}

	sect1s := node.FindChildren("sect1")
	if len(sect1s) == 0 {
		sect1s = node.FindChildren("section")
	}

	if len(sect1s) == 0 {
		// Single-file chapter
		c.ctx.CurrentFile = entry.File
		w := NewMarkdownWriter()
		if err := convertNode(c.ctx, node, w); err != nil {
			return err
		}
		return c.writeFile(entry.File, w.String())
	}

	// Multi-file chapter: write index with chapter intro
	c.ctx.CurrentFile = entry.File
	w := NewMarkdownWriter()
	title := extractTitle(node)
	w.Heading(1, title, id)

	// Write any content before the first sect1
	for _, child := range node.Children {
		if child.Type == sgml.ElementNode {
			if child.Tag == "title" || child.Tag == "titleabbrev" {
				continue
			}
			if child.Tag == "sect1" || child.Tag == "section" ||
				child.Tag == "refentry" {
				break
			}
			if err := convertNode(c.ctx, child, w); err != nil {
				return err
			}
		}
	}

	// Generate TOC listing child sections
	c.writeSectionTOC(node, w)

	if err := c.writeFile(entry.File, w.String()); err != nil {
		return err
	}

	// Convert each sect1 to its own file
	for _, sect := range sect1s {
		if err := c.convertSection(sect); err != nil {
			return err
		}
	}

	// Convert any refentry children
	for _, ref := range node.FindChildren("refentry") {
		if err := c.convertRefentry(ref); err != nil {
			return err
		}
	}

	return nil
}

// writeSectionTOC generates a table of contents for child sections
// and refentries, linking to their output files.
func (c *Converter) writeSectionTOC(node *sgml.Node, w *MarkdownWriter) {
	var items []struct{ title, id string }

	for _, child := range node.Children {
		if child.Type != sgml.ElementNode {
			continue
		}
		switch child.Tag {
		case "sect1", "section", "refentry":
			childID := child.GetAttr("id")
			childTitle := extractTitle(child)
			if childID != "" && childTitle != "" {
				items = append(items, struct{ title, id string }{
					childTitle, childID})
			}
		}
	}

	if len(items) == 0 {
		return
	}

	w.BlankLine()
	for _, item := range items {
		link, _, ok := c.ctx.ResolveLink(item.id)
		if ok {
			w.WriteString(fmt.Sprintf("- [%s](%s)\n", item.title, link))
		}
	}
}

// convertSection converts a <sect1>/<section> to its own file.
func (c *Converter) convertSection(node *sgml.Node) error {
	id := node.GetAttr("id")
	if id == "" {
		// Generate a temporary ID for lookup
		title := extractTitle(node)
		id = slugify(title)
	}

	entry, ok := c.ctx.IDMap[id]
	if !ok {
		c.ctx.Warn("section %q not found in ID map", id)
		return nil
	}

	c.ctx.CurrentFile = entry.File
	w := NewMarkdownWriter()

	if err := convertNode(c.ctx, node, w); err != nil {
		return err
	}

	return c.writeFile(entry.File, w.String())
}

// convertRefentry converts a <refentry> to its own file.
func (c *Converter) convertRefentry(node *sgml.Node) error {
	id := node.GetAttr("id")
	if id == "" {
		return nil
	}

	entry, ok := c.ctx.IDMap[id]
	if !ok {
		c.ctx.Warn("refentry %q not found in ID map", id)
		return nil
	}

	c.ctx.CurrentFile = entry.File
	w := NewMarkdownWriter()

	if err := convertNode(c.ctx, node, w); err != nil {
		return err
	}

	return c.writeFile(entry.File, w.String())
}

// writeFile writes content to an output file, creating directories
// as needed.
func (c *Converter) writeFile(relPath, content string) error {
	fullPath := filepath.Join(c.ctx.OutDir, relPath)
	dir := filepath.Dir(fullPath)

	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("creating directory %s: %w", dir, err)
	}

	// Clean up the content
	content = cleanMarkdown(content)

	return os.WriteFile(fullPath, []byte(content), 0644)
}

// convertNode converts a single node using the handler registry.
func convertNode(ctx *Context, node *sgml.Node, w *MarkdownWriter) error {
	switch node.Type {
	case sgml.TextNode:
		if w.InCodeBlock() {
			w.WriteString(node.Text)
		} else {
			// Normalise whitespace in regular text
			text := normalizeWhitespace(node.Text)
			w.Write(text)
		}
		return nil

	case sgml.CommentNode:
		return nil // skip comments

	case sgml.ElementNode:
		handler := getHandler(node.Tag)
		if handler != nil {
			return handler(ctx, node, w)
		}
		// Unknown element — warn and pass through children
		if ctx != nil {
			ctx.Warn("unhandled element <%s> at line %d", node.Tag, node.Line)
		}
		return convertChildren(ctx, node, w)
	}

	return nil
}

// convertChildren converts all children of a node.
func convertChildren(ctx *Context, node *sgml.Node, w *MarkdownWriter) error {
	for _, child := range node.Children {
		if err := convertNode(ctx, child, w); err != nil {
			return err
		}
	}
	return nil
}

// normalizeWhitespace collapses runs of whitespace to single spaces.
func normalizeWhitespace(s string) string {
	// Preserve single newlines at start/end, collapse internal whitespace
	var b strings.Builder
	inSpace := false
	for _, r := range s {
		if r == ' ' || r == '\t' || r == '\n' || r == '\r' {
			if !inSpace {
				b.WriteRune(' ')
				inSpace = true
			}
		} else {
			b.WriteRune(r)
			inSpace = false
		}
	}
	return b.String()
}

// cleanMarkdown performs final cleanup on generated Markdown.
func cleanMarkdown(content string) string {
	// Remove trailing whitespace on each line
	lines := strings.Split(content, "\n")
	for i, line := range lines {
		lines[i] = strings.TrimRight(line, " \t")
	}
	content = strings.Join(lines, "\n")

	// Collapse more than 2 consecutive blank lines to 2
	for strings.Contains(content, "\n\n\n\n") {
		content = strings.ReplaceAll(content, "\n\n\n\n", "\n\n\n")
	}

	// Ensure single trailing newline
	content = strings.TrimRight(content, "\n") + "\n"

	return content
}
