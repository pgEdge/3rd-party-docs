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
	"os"
	"path/filepath"
	"strings"

	"github.com/pgEdge/postgresql-docs/builder/shared"
)

// ToctreeEntry represents a file in the toctree hierarchy.
type ToctreeEntry struct {
	// RSTName is the RST file basename (without .rst).
	RSTName string
	// Title is the document title.
	Title string
	// OutputPath is the .md file path relative to docs/.
	OutputPath string
	// NavParent is the parent's output directory.
	NavParent string
	// Children are toctree child entries.
	Children []*ToctreeEntry
	// Labels defined in this file.
	Labels []string
}

// ResolveToctree walks the toctree starting from index.rst and builds
// the complete file map, label map, and file entry list.
func ResolveToctree(
	srcDir string,
) (
	entries []*ToctreeEntry,
	fileMap map[string]string,
	labelMap map[string]labelInfo,
	fileEntries []*shared.FileEntry,
	warnings []string,
) {
	fileMap = make(map[string]string)
	labelMap = make(map[string]labelInfo)
	order := 0

	var walk func(rstName, parentPath string, isTopLevel bool) *ToctreeEntry

	walk = func(rstName, parentPath string, isTopLevel bool) *ToctreeEntry {
		rstPath := filepath.Join(srcDir, rstName+".rst")
		data, err := os.ReadFile(rstPath)
		if err != nil {
			warnings = append(warnings,
				"could not read "+rstPath+": "+err.Error())
			return nil
		}

		root := Parse(string(data))

		// Extract title
		title := extractDocTitle(root)
		slug := shared.Slugify(title)
		if slug == "" {
			slug = rstName
		}

		// Determine output path
		var outputPath string
		if rstName == "index" {
			outputPath = "index.md"
		} else {
			outputPath = filepath.Join(parentPath, rstName+".md")
		}

		entry := &ToctreeEntry{
			RSTName:    rstName,
			Title:      title,
			OutputPath: outputPath,
			NavParent:  parentPath,
		}

		fileMap[rstName] = outputPath

		// Register as file entry
		order++
		fileEntries = append(fileEntries, &shared.FileEntry{
			Path:      outputPath,
			Title:     title,
			NavParent: parentPath,
			Order:     order,
		})

		// Collect labels
		for _, child := range root.Children {
			if child.Type == LabelNode {
				entry.Labels = append(entry.Labels, child.Label)
				labelMap[child.Label] = labelInfo{
					File:   outputPath,
					Anchor: child.Label,
					Title:  title,
				}
			}
		}

		// Find toctree directives and walk children.
		// Children of this file go into a subdirectory named
		// after this file's slug, creating the nav hierarchy.
		toctreeEntries := collectToctreeEntries(root)
		childPath := parentPath
		if len(toctreeEntries) > 0 && !isTopLevel {
			childPath = filepath.Join(parentPath, slug)
		}
		for _, childName := range toctreeEntries {
			if _, seen := fileMap[childName]; seen {
				continue
			}
			childEntry := walk(childName, childPath, false)
			if childEntry != nil {
				entry.Children = append(entry.Children, childEntry)
			}
		}

		entries = append(entries, entry)
		return entry
	}

	// index.rst is the top level — its direct toctree children
	// are the main sections; they each get their own subdirectory.
	walk("index", "", true)
	return
}

// collectToctreeEntries returns all toctree entries from a document.
func collectToctreeEntries(root *Node) []string {
	var entries []string
	var visit func(*Node)
	visit = func(n *Node) {
		if n.Type == DirectiveNode && n.DirectiveName == "toctree" {
			for _, line := range strings.Split(n.Body, "\n") {
				name := strings.TrimSpace(line)
				if name == "" {
					continue
				}
				// Strip .rst extension if present
				name = strings.TrimSuffix(name, ".rst")
				entries = append(entries, name)
			}
		}
		for _, child := range n.Children {
			visit(child)
		}
	}
	visit(root)
	return entries
}

// extractDocTitle returns the title of an RST document (first heading).
func extractDocTitle(root *Node) string {
	for _, child := range root.Children {
		if child.Type == HeadingNode {
			return cleanTitle(child.Text)
		}
	}
	return ""
}

// cleanTitle removes RST index markup and inline formatting
// from a title for use in nav entries.
func cleanTitle(text string) string {
	// Strip `Title`:index: pattern
	suffix := "`:index:"
	if idx := strings.Index(text, suffix); idx >= 0 {
		sub := text[:idx]
		start := strings.LastIndex(sub, "`")
		if start >= 0 && start < idx {
			text = text[start+1 : idx]
		}
	}

	// Strip ``literal`` -> literal (double backticks)
	text = strings.ReplaceAll(text, "``", "")

	// Strip remaining single backticks
	text = strings.ReplaceAll(text, "`", "")

	// Strip RST backslash escapes
	for strings.Contains(text, "\\") {
		text = reBackslashEscape.ReplaceAllString(text, "$1")
		if !strings.Contains(text, "\\") {
			break
		}
	}

	return strings.TrimSpace(text)
}

// ScanAllLabels does a second pass over all RST files to collect
// labels that weren't in the toctree path. It reads each file in
// the file map and scans for label definitions, updating the label
// map with the heading title that follows each label.
func ScanAllLabels(
	srcDir string,
	fileMap map[string]string,
	labelMap map[string]labelInfo,
) {
	for rstName, outputPath := range fileMap {
		rstPath := filepath.Join(srcDir, rstName+".rst")
		data, err := os.ReadFile(rstPath)
		if err != nil {
			continue
		}
		root := Parse(string(data))

		// Walk nodes looking for labels followed by headings
		for i, child := range root.Children {
			if child.Type == LabelNode {
				if _, exists := labelMap[child.Label]; exists {
					// Already registered — update with correct title
					title := ""
					if i+1 < len(root.Children) &&
						root.Children[i+1].Type == HeadingNode {
						title = cleanTitle(
							root.Children[i+1].Text)
					}
					if title != "" {
						info := labelMap[child.Label]
						info.Title = title
						labelMap[child.Label] = info
					}
					continue
				}
				title := child.Label
				if i+1 < len(root.Children) &&
					root.Children[i+1].Type == HeadingNode {
					title = cleanTitle(root.Children[i+1].Text)
				}
				labelMap[child.Label] = labelInfo{
					File:   outputPath,
					Anchor: child.Label,
					Title:  title,
				}
			}
		}
	}
}
