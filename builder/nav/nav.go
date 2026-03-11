//-------------------------------------------------------------------------
//
// pgEdge PostgreSQL Docs
//
// Copyright (c) 2026, pgEdge, Inc.
// This software is released under The PostgreSQL License
//
//-------------------------------------------------------------------------

// Package nav generates the mkdocs.yml nav section from the
// converted document structure.
package nav

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/pgEdge/postgresql-docs/builder/convert"
)

// NavEntry represents a single entry in the mkdocs nav tree.
type NavEntry struct {
	Title    string
	Path     string // file path for leaf nodes
	Children []*NavEntry
}

// BuildNav constructs a nav tree from the converter's file list.
func BuildNav(files []*convert.FileEntry) *NavEntry {
	root := &NavEntry{Title: "root"}

	for _, f := range files {
		insertEntry(root, f.Path, f.Title, f.NavParent)
	}

	return root
}

// insertEntry inserts a file entry into the nav tree at the
// appropriate position based on its path hierarchy.
func insertEntry(root *NavEntry, filePath, title, navParent string) {
	parts := strings.Split(filepath.Dir(filePath), string(filepath.Separator))
	filename := filepath.Base(filePath)

	// Navigate to the correct position in the tree
	current := root
	for _, part := range parts {
		if part == "." || part == "" {
			continue
		}
		// Find or create child directory node
		var found *NavEntry
		for _, child := range current.Children {
			if slugMatch(child.Title, part) {
				found = child
				break
			}
		}
		if found == nil {
			found = &NavEntry{Title: deslugify(part)}
			current.Children = append(current.Children, found)
		}
		current = found
	}

	// Add the file entry
	if filename == "index.md" {
		// Index file — update the parent node
		current.Path = filePath
		if title != "" {
			current.Title = title
		}
	} else {
		entry := &NavEntry{
			Title: title,
			Path:  filePath,
		}
		current.Children = append(current.Children, entry)
	}
}

// GenerateYAML produces the nav: section for mkdocs.yml.
func GenerateYAML(root *NavEntry) string {
	var b strings.Builder
	b.WriteString("nav:\n")
	for _, child := range root.Children {
		writeNavEntry(&b, child, 1)
	}
	return b.String()
}

// writeNavEntry recursively writes a nav entry as YAML.
func writeNavEntry(b *strings.Builder, entry *NavEntry, depth int) {
	indent := strings.Repeat("  ", depth)

	if len(entry.Children) == 0 {
		// Leaf node
		if entry.Title != "" {
			b.WriteString(fmt.Sprintf("%s- %s: %s\n",
				indent, yamlQuote(entry.Title), entry.Path))
		} else {
			b.WriteString(fmt.Sprintf("%s- %s\n", indent, entry.Path))
		}
		return
	}

	// Branch node
	b.WriteString(fmt.Sprintf("%s- %s:\n", indent, yamlQuote(entry.Title)))

	// Write index page first if it exists
	if entry.Path != "" {
		b.WriteString(fmt.Sprintf("%s  - %s\n",
			indent, entry.Path))
	}

	// Write children
	for _, child := range entry.Children {
		writeNavEntry(b, child, depth+1)
	}
}

// UpdateMkdocsYML reads an existing mkdocs.yml, replaces the nav:
// section and updates the site_name with the version, then writes
// it back.
func UpdateMkdocsYML(mkdocsPath, navYAML, version string) error {
	data, err := os.ReadFile(mkdocsPath)
	if err != nil {
		return fmt.Errorf("reading %s: %w", mkdocsPath, err)
	}

	content := string(data)

	// Update site_name with version
	if version != "" {
		lines := strings.Split(content, "\n")
		for i, line := range lines {
			if strings.HasPrefix(line, "site_name:") {
				lines[i] = "site_name: PostgreSQL " + version
				break
			}
		}
		content = strings.Join(lines, "\n")
	}

	// Find the nav: section and replace it
	navIdx := strings.Index(content, "\nnav:")
	if navIdx == -1 {
		// No existing nav section — append
		content = strings.TrimRight(content, "\n") + "\n\n" + navYAML
	} else {
		// Find the end of the nav section (next top-level key or EOF)
		navStart := navIdx + 1
		navEnd := len(content)

		lines := strings.Split(content[navStart:], "\n")
		lineCount := 0
		for i, line := range lines {
			if i == 0 {
				lineCount++
				continue // skip the "nav:" line itself
			}
			trimmed := strings.TrimSpace(line)
			if trimmed == "" {
				lineCount++
				continue
			}
			// Check if this is a new top-level key (no indentation)
			if len(line) > 0 && line[0] != ' ' && line[0] != '-' {
				navEnd = navStart
				for j := 0; j < lineCount; j++ {
					navEnd += len(lines[j]) + 1
				}
				break
			}
			lineCount++
		}

		content = content[:navStart] + navYAML + "\n" + content[navEnd:]
	}

	return os.WriteFile(mkdocsPath, []byte(content), 0644)
}

// yamlQuote quotes a string for YAML if it contains special characters.
func yamlQuote(s string) string {
	if strings.ContainsAny(s, ":#{}[]|>&*!%@,`\"'") {
		return "'" + strings.ReplaceAll(s, "'", "''") + "'"
	}
	return s
}

// slugMatch checks if a title matches a directory slug.
func slugMatch(title, slug string) bool {
	return convert.ExportSlugify(title) == slug
}

// deslugify converts a slug back to a readable title.
func deslugify(s string) string {
	s = strings.ReplaceAll(s, "-", " ")
	s = strings.ReplaceAll(s, "_", " ")
	// Capitalize first letter of each word
	words := strings.Fields(s)
	for i, w := range words {
		if len(w) > 0 {
			words[i] = strings.ToUpper(w[:1]) + w[1:]
		}
	}
	return strings.Join(words, " ")
}
