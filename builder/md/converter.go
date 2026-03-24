//-------------------------------------------------------------------------
//
// pgEdge PostgreSQL Docs
//
// Copyright (c) 2026, pgEdge, Inc.
// This software is released under The PostgreSQL License
//
//-------------------------------------------------------------------------

// Package md provides a Markdown-to-Markdown converter that copies
// and optionally splits upstream Markdown documentation for use
// with MkDocs Material.
package md

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/pgEdge/postgresql-docs/builder/shared"
)

// Converter processes upstream Markdown documentation, splitting
// single-file projects by H2 headings and copying multi-file
// projects as-is.
type Converter struct {
	srcDir  string
	outDir  string
	version string
	verbose bool

	files    []*shared.FileEntry
	warnings []string
}

// NewConverter creates a Markdown converter.
func NewConverter(
	srcDir, outDir, version string, verbose bool,
) *Converter {
	return &Converter{
		srcDir:  srcDir,
		outDir:  outDir,
		version: version,
		verbose: verbose,
	}
}

// Files returns the output file entries for nav generation.
func (c *Converter) Files() []*shared.FileEntry { return c.files }

// Warnings returns any warnings generated during conversion.
func (c *Converter) Warnings() []string { return c.warnings }

// Convert processes the source directory.
func (c *Converter) Convert() error {
	mdFiles, err := findMarkdownFiles(c.srcDir)
	if err != nil {
		return fmt.Errorf("scanning source: %w", err)
	}
	if len(mdFiles) == 0 {
		return fmt.Errorf("no markdown files found in %s",
			c.srcDir)
	}

	if err := os.MkdirAll(c.outDir, 0755); err != nil {
		return err
	}

	docFiles := filterDocFiles(mdFiles)
	if c.verbose {
		fmt.Printf("  Found %d doc file(s) out of %d total\n",
			len(docFiles), len(mdFiles))
	}

	if len(docFiles) == 1 {
		return c.splitFile(docFiles[0])
	}
	return c.copyFiles(docFiles)
}

// findMarkdownFiles returns .md file paths relative to dir,
// scanning recursively into subdirectories.
func findMarkdownFiles(dir string) ([]string, error) {
	var files []string
	err := filepath.WalkDir(dir,
		func(path string, d os.DirEntry, err error) error {
			if err != nil {
				return err
			}
			if d.IsDir() {
				return nil
			}
			if strings.HasSuffix(strings.ToLower(d.Name()),
				".md") {
				rel, err := filepath.Rel(dir, path)
				if err != nil {
					return err
				}
				files = append(files, rel)
			}
			return nil
		})
	return files, err
}

// skipDirs lists directory names that should be excluded from
// doc file discovery (test infrastructure, CI, etc.).
var skipDirs = map[string]bool{
	"test": true, "tests": true, "testing": true,
	".github": true, ".ci": true,
}

// filterDocFiles removes non-documentation files.
// Paths may be relative (e.g. "subdir/file.md"); filtering
// is based on the base filename and parent directories.
func filterDocFiles(files []string) []string {
	var result []string
	for _, f := range files {
		base := strings.ToLower(filepath.Base(f))
		if strings.HasPrefix(base, "frag-") {
			continue
		}
		switch base {
		case "changelog.md", "changes.md", "contributing.md",
			"license.md", "code_of_conduct.md":
			continue
		}
		// Skip files inside non-doc directories
		if inSkipDir(f) {
			continue
		}
		result = append(result, f)
	}
	return result
}

// inSkipDir checks if any path component is a skipped directory.
func inSkipDir(relPath string) bool {
	dir := filepath.Dir(relPath)
	for dir != "." && dir != "" {
		base := strings.ToLower(filepath.Base(dir))
		if skipDirs[base] {
			return true
		}
		dir = filepath.Dir(dir)
	}
	return false
}

// ── Single-file splitting ────────────────────────────────────────

// section represents one H2-delimited section of a markdown file.
type section struct {
	title   string // H2 heading text
	slug    string // URL-safe filename stem
	content string // raw content (original heading levels)
}

// splitResult holds the parsed structure of a split markdown file.
type splitResult struct {
	title    string
	intro    string
	sections []section
}

var reATXHeading = regexp.MustCompile(`^(#{1,6})\s+(.+?)(?:\s+#*)?$`)

// splitMarkdown splits markdown content by H2 headings.
func splitMarkdown(content string) splitResult {
	lines := strings.Split(content, "\n")
	var res splitResult
	var introLines []string
	var currentSec *section
	inCodeBlock := false

	for _, line := range lines {
		trimmed := strings.TrimSpace(line)

		// Track fenced code blocks
		if strings.HasPrefix(trimmed, "```") ||
			strings.HasPrefix(trimmed, "~~~") {
			inCodeBlock = !inCodeBlock
		}

		if inCodeBlock {
			if currentSec != nil {
				currentSec.content += line + "\n"
			} else {
				introLines = append(introLines, line)
			}
			continue
		}

		m := reATXHeading.FindStringSubmatch(line)
		if m != nil {
			level := len(m[1])
			text := strings.TrimSpace(m[2])

			if level == 1 && res.title == "" {
				res.title = text
				introLines = append(introLines, line)
				continue
			}

			if level == 2 {
				// Start a new section
				if currentSec != nil {
					res.sections = append(res.sections, *currentSec)
				}
				currentSec = &section{
					title:   text,
					slug:    shared.Slugify(text),
					content: line + "\n",
				}
				continue
			}
		}

		if currentSec != nil {
			currentSec.content += line + "\n"
		} else {
			introLines = append(introLines, line)
		}
	}
	if currentSec != nil {
		res.sections = append(res.sections, *currentSec)
	}

	res.intro = strings.Join(introLines, "\n")
	return res
}

// promoteHeadings reduces all heading levels by one (H2→H1, etc.).
// Lines inside fenced code blocks are left unchanged.
func promoteHeadings(content string) string {
	lines := strings.Split(content, "\n")
	inCodeBlock := false
	for i, line := range lines {
		trimmed := strings.TrimSpace(line)
		if strings.HasPrefix(trimmed, "```") ||
			strings.HasPrefix(trimmed, "~~~") {
			inCodeBlock = !inCodeBlock
			continue
		}
		if !inCodeBlock && strings.HasPrefix(line, "##") {
			lines[i] = line[1:]
		}
	}
	return strings.Join(lines, "\n")
}

// githubAnchor generates the anchor slug that GitHub/MkDocs
// produce for a heading.
func githubAnchor(title string) string {
	s := strings.ToLower(title)
	var b strings.Builder
	for _, r := range s {
		switch {
		case r >= 'a' && r <= 'z', r >= '0' && r <= '9':
			b.WriteRune(r)
		case r == ' ' || r == '-':
			b.WriteRune('-')
		case r == '_':
			b.WriteRune('_')
		}
	}
	return b.String()
}

// buildAnchorMap maps internal anchors to their target files.
// H2-level section anchors point to the section file;
// sub-headings point to section file + anchor.
func buildAnchorMap(sections []section) map[string]string {
	m := make(map[string]string)
	for _, s := range sections {
		// H2 anchor → section file
		anchor := githubAnchor(s.title)
		m[anchor] = s.slug + ".md"

		// Scan for sub-headings (H3+) within the section
		inCode := false
		for _, line := range strings.Split(s.content, "\n") {
			trimmed := strings.TrimSpace(line)
			if strings.HasPrefix(trimmed, "```") ||
				strings.HasPrefix(trimmed, "~~~") {
				inCode = !inCode
				continue
			}
			if inCode {
				continue
			}
			sub := reATXHeading.FindStringSubmatch(line)
			if sub != nil && len(sub[1]) >= 3 {
				subAnchor := githubAnchor(
					strings.TrimSpace(sub[2]))
				// After promotion, the anchor stays the same
				m[subAnchor] = s.slug + ".md#" + subAnchor
			}
		}
	}
	return m
}

var reAnchorLink = regexp.MustCompile(
	`\]\(#([a-z0-9_-]+)\)`)

// rewriteAnchors replaces internal #anchor links with the
// appropriate file paths from the anchor map.
func rewriteAnchors(
	content string, anchorMap map[string]string,
) string {
	return reAnchorLink.ReplaceAllStringFunc(
		content, func(match string) string {
			sub := reAnchorLink.FindStringSubmatch(match)
			if len(sub) < 2 {
				return match
			}
			if target, ok := anchorMap[sub[1]]; ok {
				return "](" + target + ")"
			}
			return match
		})
}

// convertAlerts converts GitHub-flavored alerts to MkDocs
// admonitions.
//
// Input:
//
//	> [!NOTE]
//	> Alert body text
//
// Output:
//
//	!!! note
//	    Alert body text
func convertAlerts(content string) string {
	lines := strings.Split(content, "\n")
	var result []string
	alertTypes := map[string]string{
		"NOTE":      "note",
		"TIP":       "tip",
		"IMPORTANT": "important",
		"WARNING":   "warning",
		"CAUTION":   "danger",
	}
	reAlert := regexp.MustCompile(
		`^>\s*\[!(NOTE|TIP|IMPORTANT|WARNING|CAUTION)\]\s*$`)

	i := 0
	for i < len(lines) {
		m := reAlert.FindStringSubmatch(lines[i])
		if m == nil {
			result = append(result, lines[i])
			i++
			continue
		}
		admonType := alertTypes[m[1]]
		result = append(result, "")
		result = append(result, "!!! "+admonType)
		i++
		// Consume continuation lines starting with >
		for i < len(lines) {
			line := lines[i]
			if strings.HasPrefix(line, "> ") {
				result = append(result,
					"    "+strings.TrimPrefix(line, "> "))
				i++
			} else if line == ">" {
				result = append(result, "")
				i++
			} else {
				break
			}
		}
		result = append(result, "")
	}
	return strings.Join(result, "\n")
}

// splitFile splits a single markdown file by H2 and writes
// the resulting pages.
func (c *Converter) splitFile(filename string) error {
	srcPath := filepath.Join(c.srcDir, filename)
	data, err := os.ReadFile(srcPath)
	if err != nil {
		return err
	}

	content := string(data)
	baseDir := filepath.Dir(c.srcDir)
	content = shared.ResolveSnippets(content, srcPath, baseDir)
	content = convertAlerts(content)
	res := splitMarkdown(content)

	if len(res.sections) == 0 {
		// No H2 sections — just copy as index.md
		if err := c.writeFile("index.md", content); err != nil {
			return err
		}
		title := res.title
		if title == "" {
			title = strings.TrimSuffix(filename,
				filepath.Ext(filename))
		}
		c.files = append(c.files, &shared.FileEntry{
			Path:  "index.md",
			Title: title,
			Order: 0,
		})
		return nil
	}

	anchorMap := buildAnchorMap(res.sections)

	// Write index.md (intro)
	intro := rewriteAnchors(res.intro, anchorMap)
	if err := c.writeFile("index.md", intro); err != nil {
		return err
	}
	title := res.title
	if title == "" {
		title = strings.TrimSuffix(filename,
			filepath.Ext(filename))
	}
	c.files = append(c.files, &shared.FileEntry{
		Path:  "index.md",
		Title: title,
		Order: 0,
	})

	// Write section files
	for i, s := range res.sections {
		body := promoteHeadings(s.content)
		body = rewriteAnchors(body, anchorMap)
		outName := s.slug + ".md"
		if err := c.writeFile(outName, body); err != nil {
			return err
		}
		c.files = append(c.files, &shared.FileEntry{
			Path:  outName,
			Title: s.title,
			Order: i + 1,
		})
	}

	if c.verbose {
		fmt.Printf("  Split into %d pages\n",
			len(res.sections)+1)
	}
	return nil
}

// ── Multi-file copying ───────────────────────────────────────────

// copyFiles copies multiple markdown files, creating an index
// if none exists.
func (c *Converter) copyFiles(files []string) error {
	hasIndex := false
	for _, f := range files {
		lower := strings.ToLower(f)
		// Only top-level README/index counts as the site index
		if lower == "readme.md" || lower == "index.md" {
			hasIndex = true
			break
		}
	}

	for i, f := range files {
		data, err := os.ReadFile(filepath.Join(c.srcDir, f))
		if err != nil {
			return err
		}

		content := string(data)
		srcPath := filepath.Join(c.srcDir, f)
		baseDir := filepath.Dir(c.srcDir)
		content = shared.ResolveSnippets(content, srcPath, baseDir)
		content = convertAlerts(content)

		outName := f
		lower := strings.ToLower(filepath.Base(f))
		if lower == "readme.md" {
			// Rename README.md to index.md, preserving dir
			outName = filepath.Join(
				filepath.Dir(f), "index.md")
		}

		title := extractTitle(content, f)

		if err := c.writeFile(outName, content); err != nil {
			return err
		}
		c.files = append(c.files, &shared.FileEntry{
			Path:  outName,
			Title: title,
			Order: i,
		})
	}

	// Generate index.md if the source had no README/index
	if !hasIndex {
		idx := c.generateIndex()
		if err := c.writeFile("index.md", idx); err != nil {
			return err
		}
		// Prepend index entry
		c.files = append([]*shared.FileEntry{{
			Path:  "index.md",
			Title: c.version,
			Order: -1,
		}}, c.files...)
	}

	if c.verbose {
		fmt.Printf("  Copied %d files\n", len(files))
	}
	return nil
}

// reHTMLH1 matches HTML <h1> tags (possibly multiline) and
// extracts the inner text, stripping nested tags like <b>.
var reHTMLH1 = regexp.MustCompile(
	`(?is)<h1[^>]*>(.*?)</h1>`)
var reHTMLTags = regexp.MustCompile(`<[^>]+>`)

// extractTitle returns the first heading text from content.
// It checks for ATX H1 (#), HTML <h1>, and ATX H2 (##) in
// that priority order, falling back to the filename stem.
func extractTitle(content, filename string) string {
	// Try HTML <h1> first (handles multiline tags like
	// <h1 align="center"><b>Title</b></h1>)
	if m := reHTMLH1.FindStringSubmatch(content); m != nil {
		inner := reHTMLTags.ReplaceAllString(m[1], "")
		inner = strings.Join(strings.Fields(inner), " ")
		if inner != "" {
			return inner
		}
	}

	// Scan for ATX headings, skipping code blocks
	var firstH2 string
	inCodeBlock := false
	scanner := bufio.NewScanner(strings.NewReader(content))
	for scanner.Scan() {
		line := scanner.Text()
		trimmed := strings.TrimSpace(line)
		if strings.HasPrefix(trimmed, "```") ||
			strings.HasPrefix(trimmed, "~~~") {
			inCodeBlock = !inCodeBlock
			continue
		}
		if inCodeBlock {
			continue
		}
		m := reATXHeading.FindStringSubmatch(line)
		if m != nil {
			if len(m[1]) == 1 {
				return strings.TrimSpace(m[2])
			}
			if len(m[1]) == 2 && firstH2 == "" {
				firstH2 = strings.TrimSpace(m[2])
			}
		}
	}

	// Fall back to first H2
	if firstH2 != "" {
		return firstH2
	}

	base := filepath.Base(filename)
	return strings.TrimSuffix(base, filepath.Ext(base))
}

// generateIndex creates a simple index page linking to all
// other pages.
func (c *Converter) generateIndex() string {
	var b strings.Builder
	b.WriteString("# " + c.version + "\n\n")
	for _, f := range c.files {
		b.WriteString(fmt.Sprintf("- [%s](%s)\n", f.Title, f.Path))
	}
	return b.String()
}

// writeFile writes content to a file under the output directory.
func (c *Converter) writeFile(relPath, content string) error {
	outPath := filepath.Join(c.outDir, relPath)
	if err := os.MkdirAll(filepath.Dir(outPath), 0755); err != nil {
		return err
	}
	return os.WriteFile(outPath, []byte(content), 0644)
}
