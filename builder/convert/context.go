//-------------------------------------------------------------------------
//
// pgEdge PostgreSQL Docs
//
// Copyright (c) 2026, pgEdge, Inc.
// This software is released under The PostgreSQL License
//
//-------------------------------------------------------------------------

// Package convert transforms a parsed SGML document tree into
// Markdown files suitable for MkDocs Material.
package convert

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/pgEdge/postgresql-docs/builder/sgml"
)

// IDEntry records where a document ID maps to in the output.
type IDEntry struct {
	// File is the output .md file path relative to docs/.
	File string
	// Anchor is the anchor name within the file.
	Anchor string
	// Title is the generated link text for <xref> references.
	Title string
	// Type is the element type (chapter, sect1, refentry, etc.).
	Type string
}

// FileEntry tracks a single output file being generated.
type FileEntry struct {
	// Path relative to docs/ (e.g., "tutorial/start.md").
	Path string
	// Title for the nav entry.
	Title string
	// NavParent is the parent nav path for hierarchy.
	NavParent string
	// Order preserves document order within the nav.
	Order int
}

// Context holds the state for a conversion run.
type Context struct {
	// IDMap maps element IDs to their output locations.
	IDMap map[string]*IDEntry
	// Files tracks all output files in document order.
	Files []*FileEntry
	// Warnings accumulates non-fatal issues.
	Warnings []string
	// CurrentFile is the file currently being written.
	CurrentFile string
	// Version is the PostgreSQL version string.
	Version string
	// SrcDir is the SGML source directory.
	SrcDir string
	// OutDir is the output directory for .md files.
	OutDir string

	// docRoot is the parsed document tree.
	docRoot *sgml.Node
	// fileOrder counter for maintaining document order.
	fileOrder int
}

// NewContext creates a new conversion context.
func NewContext(root *sgml.Node, srcDir, outDir, version string) *Context {
	return &Context{
		IDMap:   make(map[string]*IDEntry),
		docRoot: root,
		SrcDir:  srcDir,
		OutDir:  outDir,
		Version: version,
	}
}

// Warn adds a warning message.
func (ctx *Context) Warn(format string, args ...any) {
	msg := fmt.Sprintf(format, args...)
	ctx.Warnings = append(ctx.Warnings, msg)
}

// RegisterID adds an ID to the map.
func (ctx *Context) RegisterID(id, file, anchor, title, elemType string) {
	ctx.IDMap[id] = &IDEntry{
		File:   file,
		Anchor: anchor,
		Title:  title,
		Type:   elemType,
	}
}

// AddFile registers an output file.
func (ctx *Context) AddFile(path, title, navParent string) {
	ctx.fileOrder++
	ctx.Files = append(ctx.Files, &FileEntry{
		Path:      path,
		Title:     title,
		NavParent: navParent,
		Order:     ctx.fileOrder,
	})
}

// ResolveLink returns the Markdown link path from the current file
// to the given ID target.
func (ctx *Context) ResolveLink(id string) (string, string, bool) {
	entry, ok := ctx.IDMap[id]
	if !ok {
		return "", "", false
	}

	// Calculate relative path from current file to target
	fromDir := filepath.Dir(ctx.CurrentFile)
	relPath, err := filepath.Rel(fromDir, entry.File)
	if err != nil {
		relPath = entry.File
	}

	// Same file — just use anchor
	if entry.File == ctx.CurrentFile {
		return "#" + entry.Anchor, entry.Title, true
	}

	link := relPath
	if entry.Anchor != "" {
		link += "#" + entry.Anchor
	}

	return link, entry.Title, true
}

// slugify converts a string to a URL-friendly slug.
func slugify(s string) string {
	s = strings.ToLower(s)
	var b strings.Builder
	for _, r := range s {
		switch {
		case r >= 'a' && r <= 'z':
			b.WriteRune(r)
		case r >= '0' && r <= '9':
			b.WriteRune(r)
		case r == '-' || r == '_':
			b.WriteRune(r)
		case r == ' ' || r == '/' || r == '.':
			b.WriteRune('-')
		}
	}
	// Collapse multiple hyphens
	result := b.String()
	for strings.Contains(result, "--") {
		result = strings.ReplaceAll(result, "--", "-")
	}
	return strings.Trim(result, "-")
}
