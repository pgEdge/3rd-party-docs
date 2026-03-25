//-------------------------------------------------------------------------
//
// pgEdge PostgreSQL Docs
//
// Copyright (c) 2026, pgEdge, Inc.
// This software is released under The PostgreSQL License
//
//-------------------------------------------------------------------------

// Package mkdocsmode provides a converter for upstream projects
// that already have their own mkdocs.yml with curated nav, extensions,
// and plugins. Instead of generating nav from file structure, it
// imports the upstream config directly.
package mkdocsmode

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/pgEdge/postgresql-docs/builder/shared"
)

// Converter imports an existing MkDocs site by copying docs and
// merging the upstream mkdocs.yml config into the skeleton.
type Converter struct {
	srcDir  string // upstream docs directory (src_subdir)
	outDir  string // output directory (./docs)
	version string
	verbose bool

	repoRoot string // parent of srcDir (for snippet resolution)

	navYAML    string   // raw nav: block from upstream
	extensions []string // extension lines from upstream
	plugins    []string // plugin lines from upstream

	files    []*shared.FileEntry
	warnings []string
}

// NewConverter creates an MkDocs mode converter.
func NewConverter(
	srcDir, outDir, version string, verbose bool,
) *Converter {
	return &Converter{
		srcDir:   srcDir,
		outDir:   outDir,
		version:  version,
		verbose:  verbose,
		repoRoot: filepath.Dir(srcDir),
	}
}

// Files returns the output file entries.
func (c *Converter) Files() []*shared.FileEntry { return c.files }

// Warnings returns any warnings generated during conversion.
func (c *Converter) Warnings() []string { return c.warnings }

// NavYAML returns the extracted nav block for injection.
func (c *Converter) NavYAML() string { return c.navYAML }

// Extensions returns upstream markdown_extensions lines.
func (c *Converter) Extensions() []string { return c.extensions }

// Plugins returns upstream plugin lines.
func (c *Converter) Plugins() []string { return c.plugins }

// Convert runs the import pipeline.
func (c *Converter) Convert() error {
	// Find upstream mkdocs.yml
	mkdocsPath := c.findMkdocsYML()
	if mkdocsPath == "" {
		return fmt.Errorf(
			"no mkdocs.yml found in %s or its parent", c.srcDir)
	}
	if c.verbose {
		fmt.Printf("  Using upstream mkdocs.yml: %s\n", mkdocsPath)
	}

	// Parse upstream config
	data, err := os.ReadFile(mkdocsPath)
	if err != nil {
		return fmt.Errorf("reading upstream mkdocs.yml: %w", err)
	}
	content := string(data)

	c.navYAML = extractYAMLBlock(content, "nav")
	if c.navYAML == "" {
		return fmt.Errorf("no nav: section in upstream mkdocs.yml")
	}

	c.extensions = extractYAMLList(content, "markdown_extensions")
	c.plugins = extractYAMLList(content, "plugins")

	// Extract file paths from nav
	navPaths := extractNavPaths(c.navYAML)
	if c.verbose {
		fmt.Printf("  Found %d files in upstream nav\n",
			len(navPaths))
	}

	// Copy all files from srcDir (preserving structure)
	if err := c.copyTree(); err != nil {
		return fmt.Errorf("copying docs: %w", err)
	}

	// Resolve snippets in copied files
	if err := c.resolveAllSnippets(); err != nil {
		return fmt.Errorf("resolving snippets: %w", err)
	}

	// Fix broken relative links
	if err := shared.FixBrokenLinksInDir(c.outDir); err != nil {
		return fmt.Errorf("fixing links: %w", err)
	}

	// Build file entries from nav paths
	for i, p := range navPaths {
		c.files = append(c.files, &shared.FileEntry{
			Path:  p,
			Title: p,
			Order: i,
		})
	}

	if c.verbose {
		fmt.Printf("  Copied %d files\n", len(c.files))
	}
	return nil
}

// findMkdocsYML looks for mkdocs.yml in the source directory
// and then in its parent (repo root).
func (c *Converter) findMkdocsYML() string {
	candidates := []string{
		filepath.Join(c.srcDir, "mkdocs.yml"),
		filepath.Join(c.repoRoot, "mkdocs.yml"),
	}
	for _, p := range candidates {
		if _, err := os.Stat(p); err == nil {
			return p
		}
	}
	return ""
}

// copyTree copies all files from srcDir to outDir, preserving
// the directory structure. Non-doc files (images, etc.) are
// included to avoid broken references.
func (c *Converter) copyTree() error {
	return filepath.WalkDir(c.srcDir,
		func(path string, d os.DirEntry, err error) error {
			if err != nil {
				return err
			}
			rel, err := filepath.Rel(c.srcDir, path)
			if err != nil {
				return err
			}
			outPath := filepath.Join(c.outDir, rel)

			if d.IsDir() {
				return os.MkdirAll(outPath, 0755)
			}

			data, err := os.ReadFile(path)
			if err != nil {
				return err
			}
			if err := os.MkdirAll(
				filepath.Dir(outPath), 0755); err != nil {
				return err
			}
			return os.WriteFile(outPath, data, 0644)
		})
}

// resolveAllSnippets walks the output directory and resolves
// pymdownx.snippets includes in all markdown files.
func (c *Converter) resolveAllSnippets() error {
	return filepath.WalkDir(c.outDir,
		func(path string, d os.DirEntry, err error) error {
			if err != nil {
				return err
			}
			if d.IsDir() {
				return nil
			}
			if !strings.HasSuffix(
				strings.ToLower(d.Name()), ".md") {
				return nil
			}
			data, err := os.ReadFile(path)
			if err != nil {
				return err
			}
			content := string(data)
			if !shared.ReSnippet.MatchString(content) {
				return nil
			}

			// Resolve using the original source path
			rel, _ := filepath.Rel(c.outDir, path)
			srcPath := filepath.Join(c.srcDir, rel)
			resolved := shared.ResolveSnippets(
				content, srcPath, c.repoRoot)
			if resolved != content {
				if c.verbose {
					fmt.Printf("  Resolved snippets in %s\n",
						rel)
				}
				return os.WriteFile(
					path, []byte(resolved), 0644)
			}
			return nil
		})
}

// ── YAML helpers ──────────────────────────────────────────────

// extractYAMLBlock extracts a top-level YAML block by key,
// returning the full block including the key line.
func extractYAMLBlock(content, key string) string {
	lines := strings.Split(content, "\n")
	prefix := key + ":"
	startIdx := -1

	for i, line := range lines {
		if strings.TrimSpace(line) == prefix ||
			strings.HasPrefix(line, prefix) {
			startIdx = i
			break
		}
	}
	if startIdx == -1 {
		return ""
	}

	// Collect the key line and all subsequent indented/blank lines
	endIdx := startIdx + 1
	for endIdx < len(lines) {
		line := lines[endIdx]
		if line == "" || strings.TrimSpace(line) == "" {
			endIdx++
			continue
		}
		// Stop at next top-level key (non-indented, non-blank)
		if len(line) > 0 && line[0] != ' ' && line[0] != '\t' {
			break
		}
		endIdx++
	}

	// Trim trailing blank lines
	for endIdx > startIdx+1 &&
		strings.TrimSpace(lines[endIdx-1]) == "" {
		endIdx--
	}

	return strings.Join(lines[startIdx:endIdx], "\n") + "\n"
}

// extractYAMLList extracts items from a YAML list block,
// returning each item as a string (preserving sub-indentation
// for items with config).
func extractYAMLList(content, key string) []string {
	block := extractYAMLBlock(content, key)
	if block == "" {
		return nil
	}

	lines := strings.Split(block, "\n")
	var items []string
	var current []string

	for _, line := range lines[1:] { // skip the key line
		trimmed := strings.TrimSpace(line)
		if trimmed == "" {
			continue
		}
		if strings.HasPrefix(trimmed, "- ") {
			// New item — flush previous
			if len(current) > 0 {
				items = append(items,
					strings.Join(current, "\n"))
			}
			current = []string{trimmed}
		} else if len(current) > 0 {
			// Continuation of current item (sub-config)
			current = append(current, trimmed)
		}
	}
	if len(current) > 0 {
		items = append(items, strings.Join(current, "\n"))
	}
	return items
}

// reNavPath matches file paths in nav YAML lines.
var reNavPath = regexp.MustCompile(
	`['"]([\w/.-]+\.md)['"]|:\s+([\w/.-]+\.md)\s*$|^-\s+([\w/.-]+\.md)\s*$`)

// extractNavPaths returns all .md file paths from a nav block.
func extractNavPaths(navBlock string) []string {
	var paths []string
	seen := make(map[string]bool)
	for _, line := range strings.Split(navBlock, "\n") {
		trimmed := strings.TrimSpace(line)
		if trimmed == "" || trimmed == "nav:" {
			continue
		}
		m := reNavPath.FindStringSubmatch(trimmed)
		if m != nil {
			// Pick the first non-empty group
			for _, g := range m[1:] {
				if g != "" && !seen[g] {
					paths = append(paths, g)
					seen[g] = true
					break
				}
			}
		}
	}
	return paths
}

// ── MkDocs YML merging ───────────────────────────────────────

// skipPlugins lists plugins to exclude from merging because
// they require source code or infrastructure we don't have.
var skipPlugins = map[string]bool{
	"mkdocstrings": true,
}

// MergeMkdocsYML reads the skeleton mkdocs.yml, replaces the
// nav section with the upstream nav, sets site_name, and
// merges upstream extensions and plugins.
func MergeMkdocsYML(
	mkdocsPath, navYAML, siteName string,
	extensions, plugins []string,
) error {
	data, err := os.ReadFile(mkdocsPath)
	if err != nil {
		return fmt.Errorf("reading %s: %w", mkdocsPath, err)
	}
	content := string(data)

	// Update site_name
	if siteName != "" {
		lines := strings.Split(content, "\n")
		for i, line := range lines {
			if strings.HasPrefix(line, "site_name:") {
				lines[i] = "site_name: " + siteName
				break
			}
		}
		content = strings.Join(lines, "\n")
	}

	// Replace or append nav
	content = replaceOrAppendBlock(content, navYAML)

	// Merge extensions
	content = mergeListBlock(content,
		"markdown_extensions", extensions, nil)

	// Merge plugins (skip mkdocstrings etc.)
	content = mergeListBlock(content,
		"plugins", plugins, skipPlugins)

	return os.WriteFile(mkdocsPath, []byte(content), 0644)
}

// replaceOrAppendBlock replaces the nav: section or appends it.
func replaceOrAppendBlock(content, navYAML string) string {
	navIdx := strings.Index(content, "\nnav:")
	if navIdx == -1 {
		// No existing nav — append
		return strings.TrimRight(content, "\n") +
			"\n\n" + navYAML
	}

	// Find end of existing nav section
	navStart := navIdx + 1
	lines := strings.Split(content[navStart:], "\n")
	lineCount := 0
	for i, line := range lines {
		if i == 0 {
			lineCount++
			continue
		}
		trimmed := strings.TrimSpace(line)
		if trimmed == "" {
			lineCount++
			continue
		}
		if len(line) > 0 && line[0] != ' ' && line[0] != '-' {
			break
		}
		lineCount++
	}
	navEnd := navStart
	for j := 0; j < lineCount; j++ {
		navEnd += len(lines[j]) + 1
	}

	return content[:navStart] + navYAML + "\n" + content[navEnd:]
}

// itemName extracts the base name from a YAML list item.
// e.g., "- pymdownx.snippets:" → "pymdownx.snippets"
//
//	"- admonition" → "admonition"
//	"- search" → "search"
func itemName(item string) string {
	s := strings.TrimSpace(item)
	// Get first line only
	if idx := strings.Index(s, "\n"); idx >= 0 {
		s = s[:idx]
	}
	s = strings.TrimPrefix(s, "- ")
	s = strings.TrimSuffix(s, ":")
	s = strings.TrimSpace(s)
	return s
}

// mergeListBlock adds items to a YAML list block, skipping
// duplicates and items in the skip set.
func mergeListBlock(
	content, key string,
	newItems []string,
	skip map[string]bool,
) string {
	if len(newItems) == 0 {
		return content
	}

	// Find existing items
	existing := extractYAMLList(content, key)
	existingNames := make(map[string]bool)
	for _, item := range existing {
		existingNames[itemName(item)] = true
	}

	// Determine indentation from existing block
	indent := "  "
	lines := strings.Split(content, "\n")
	for _, line := range lines {
		if strings.TrimSpace(line) == key+":" {
			continue
		}
		trimmed := strings.TrimSpace(line)
		if strings.HasPrefix(trimmed, "- ") &&
			strings.Contains(content, key) {
			indent = line[:len(line)-len(strings.TrimLeft(
				line, " \t"))]
			break
		}
	}

	// Build lines to append
	var toAdd []string
	for _, item := range newItems {
		name := itemName(item)
		if existingNames[name] {
			continue
		}
		if skip != nil && skip[name] {
			continue
		}
		// Also skip pymdownx.snippets since we resolve inline
		if name == "pymdownx.snippets" {
			continue
		}
		// Format the item with proper indentation
		itemLines := strings.Split(item, "\n")
		for i, il := range itemLines {
			if i == 0 {
				toAdd = append(toAdd, indent+il)
			} else {
				toAdd = append(toAdd, indent+"  "+il)
			}
		}
	}

	if len(toAdd) == 0 {
		return content
	}

	// Find the end of the block to insert before
	blockEnd := findBlockEnd(content, key)
	if blockEnd == -1 {
		return content
	}

	return content[:blockEnd] +
		strings.Join(toAdd, "\n") + "\n" +
		content[blockEnd:]
}

// findBlockEnd returns the byte offset just past the last item
// in a YAML list block.
func findBlockEnd(content, key string) int {
	lines := strings.Split(content, "\n")
	prefix := key + ":"
	inBlock := false
	lastItemEnd := -1
	offset := 0

	for _, line := range lines {
		lineEnd := offset + len(line) + 1
		trimmed := strings.TrimSpace(line)
		if !inBlock {
			if trimmed == prefix ||
				strings.HasPrefix(trimmed, prefix+" ") {
				inBlock = true
			}
		} else {
			if trimmed == "" {
				offset = lineEnd
				continue
			}
			if len(line) > 0 &&
				line[0] != ' ' && line[0] != '\t' {
				break
			}
			lastItemEnd = lineEnd
		}
		offset = lineEnd
	}
	return lastItemEnd
}
