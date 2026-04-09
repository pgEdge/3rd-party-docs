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
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
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
		if err := c.splitFile(docFiles[0]); err != nil {
			return err
		}
	} else {
		if err := c.copyFiles(docFiles); err != nil {
			return err
		}
	}

	// Copy image assets referenced by documentation
	if err := c.copyImages(); err != nil {
		c.warnings = append(c.warnings,
			fmt.Sprintf("copying images: %v", err))
	}

	// Post-process: fix broken relative links
	return shared.FixBrokenLinksInDir(c.outDir)
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
			"license.md", "code_of_conduct.md",
			"code-of-conduct.md", "security.md":
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

		// Scan for all headings within the section content.
		// H1 headings can appear inside a section when the
		// upstream file uses inconsistent heading levels
		// (e.g. mostly H1 with one H2 as the split point).
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
			if sub == nil {
				continue
			}
			level := len(sub[1])
			subAnchor := githubAnchor(
				strings.TrimSpace(sub[2]))
			if level == 2 {
				// Skip — already mapped above as the
				// section title
				continue
			}
			// After promotion, the anchor stays the same
			m[subAnchor] = s.slug + ".md#" + subAnchor
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

// githubEmoji maps commonly used GitHub emoji shortcodes to
// their Unicode equivalents.
var githubEmoji = map[string]string{
	":heavy_check_mark:":         "\u2714\uFE0F",
	":white_check_mark:":         "\u2705",
	":x:":                        "\u274C",
	":warning:":                  "\u26A0\uFE0F",
	":information_source:":       "\u2139\uFE0F",
	":bulb:":                     "\U0001F4A1",
	":memo:":                     "\U0001F4DD",
	":rocket:":                   "\U0001F680",
	":star:":                     "\u2B50",
	":thumbsup:":                 "\U0001F44D",
	":thumbsdown:":               "\U0001F44E",
	":tada:":                     "\U0001F389",
	":construction:":             "\U0001F6A7",
	":lock:":                     "\U0001F512",
	":key:":                      "\U0001F511",
	":hammer:":                   "\U0001F528",
	":gear:":                     "\u2699\uFE0F",
	":link:":                     "\U0001F517",
	":book:":                     "\U0001F4D6",
	":clipboard:":                "\U0001F4CB",
	":chart_with_upwards_trend:": "\U0001F4C8",
}

var reEmoji = regexp.MustCompile(`:([a-z0-9_]+):`)

// convertEmoji replaces GitHub emoji shortcodes like
// :heavy_check_mark: with their Unicode equivalents.
func convertEmoji(content string) string {
	return reEmoji.ReplaceAllStringFunc(content,
		func(match string) string {
			if u, ok := githubEmoji[match]; ok {
				return u
			}
			return match
		})
}

// stripLeadingImages removes image-only lines (including
// linked images) that appear before the first heading. These
// are typically GitHub repo banners/badges that won't render
// correctly in MkDocs because the image isn't copied to docs/.
func stripLeadingImages(content string) string {
	lines := strings.Split(content, "\n")
	var result []string
	pastPreamble := false
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if !pastPreamble {
			// Skip blank lines and image-only lines before
			// the first heading or text content
			if trimmed == "" {
				result = append(result, line)
				continue
			}
			if strings.HasPrefix(trimmed, "![") ||
				strings.HasPrefix(trimmed, "[![") {
				continue
			}
			pastPreamble = true
		}
		result = append(result, line)
	}
	return strings.Join(result, "\n")
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

// ── Docusaurus support ───────────────────────────────────────────

// docFrontmatter holds parsed YAML frontmatter from Docusaurus
// markdown files.
type docFrontmatter struct {
	Title           string
	SidebarPosition int
	HasPosition     bool
}

// parseFrontmatter extracts YAML frontmatter delimited by ---
// lines and returns the parsed metadata plus the remaining
// content with frontmatter stripped.
func parseFrontmatter(content string) (docFrontmatter, string) {
	var fm docFrontmatter
	if !strings.HasPrefix(content, "---\n") {
		return fm, content
	}

	end := strings.Index(content[4:], "\n---")
	if end < 0 {
		return fm, content
	}

	fmBlock := content[4 : 4+end]
	// Skip past \n---\n
	afterClose := 4 + end + 4 // past "---\n"
	if afterClose > len(content) {
		afterClose = len(content)
	}
	rest := content[afterClose:]
	// Strip leading newline after frontmatter close
	rest = strings.TrimPrefix(rest, "\n")

	for _, line := range strings.Split(fmBlock, "\n") {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "title:") {
			val := strings.TrimSpace(
				strings.TrimPrefix(line, "title:"))
			val = strings.Trim(val, "\"'")
			fm.Title = val
		}
		if strings.HasPrefix(line, "sidebar_position:") {
			val := strings.TrimSpace(
				strings.TrimPrefix(line, "sidebar_position:"))
			if n, err := strconv.Atoi(val); err == nil {
				fm.SidebarPosition = n
				fm.HasPosition = true
			}
		}
	}

	return fm, rest
}

// reDocusaurusAdmonition matches Docusaurus admonition openers.
// Supports three title forms:
//
//	:::warning              (no title)
//	:::info[Important]      (bracket title)
//	:::note Auth Methods    (space title)
var reDocusaurusAdmonition = regexp.MustCompile(
	`^:::+(note|tip|warning|info|danger|caution|important)` +
		`(?:\[(.+?)\]|\s+(.+?))?\s*$`)

// reDocusaurusClose matches the closing ::: (3 or more colons).
var reDocusaurusClose = regexp.MustCompile(`^:::+\s*$`)

// convertDocusaurusAdmonitions converts Docusaurus :::type and
// :::type[Title] fenced admonitions to MkDocs !!! admonitions.
// Content that is not indented is auto-indented to 4 spaces.
//
// Input:
//
//	:::warning
//	    Content here.
//	:::
//
// Output:
//
//	!!! warning
//	    Content here.
func convertDocusaurusAdmonitions(content string) string {
	lines := strings.Split(content, "\n")
	var result []string
	inAdmonition := false
	inCodeBlock := false

	for _, line := range lines {
		trimmed := strings.TrimSpace(line)

		// Track fenced code blocks
		if strings.HasPrefix(trimmed, "```") ||
			strings.HasPrefix(trimmed, "~~~") {
			inCodeBlock = !inCodeBlock
		}

		if inCodeBlock {
			result = append(result, line)
			continue
		}

		if inAdmonition {
			if reDocusaurusClose.MatchString(trimmed) {
				inAdmonition = false
				result = append(result, "")
				continue
			}
			// Ensure content is indented for MkDocs
			if trimmed != "" && !strings.HasPrefix(line, "    ") {
				line = "    " + line
			}
			result = append(result, line)
			continue
		}

		m := reDocusaurusAdmonition.FindStringSubmatch(trimmed)
		if m != nil {
			admonType := m[1]
			// Title from brackets [Title] or space Title
			title := m[2]
			if title == "" {
				title = m[3]
			}

			result = append(result, "")
			if title != "" {
				result = append(result,
					fmt.Sprintf("!!! %s \"%s\"",
						admonType, title))
			} else {
				result = append(result, "!!! "+admonType)
			}
			inAdmonition = true
			continue
		}

		result = append(result, line)
	}

	return strings.Join(result, "\n")
}

var reSPDX = regexp.MustCompile(
	`(?m)^<!-- SPDX-License-Identifier: [^\n]+ -->\n?`)

// stripSPDXComment removes SPDX license identifier HTML
// comments that Docusaurus files include after the heading.
func stripSPDXComment(content string) string {
	return reSPDX.ReplaceAllString(content, "")
}

// categoryMeta holds parsed _category_.json data from
// Docusaurus subdirectories.
type categoryMeta struct {
	Label    string `json:"label"`
	Position int    `json:"position"`
}

// readCategoryJSON reads a _category_.json file from a
// directory and returns its metadata, or nil if not found.
func readCategoryJSON(dir string) *categoryMeta {
	data, err := os.ReadFile(
		filepath.Join(dir, "_category_.json"))
	if err != nil {
		return nil
	}
	var meta categoryMeta
	if err := json.Unmarshal(data, &meta); err != nil {
		return nil
	}
	return &meta
}

// docEntry holds intermediate state during multi-file
// processing with frontmatter-based ordering.
type docEntry struct {
	relPath  string
	content  string
	title    string
	order    int
	hasOrder bool
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
	// Strip frontmatter if present (single-file mode
	// doesn't use sidebar_position for ordering)
	if _, stripped := parseFrontmatter(content); stripped != content {
		content = stripped
	}
	content = stripSPDXComment(content)
	baseDir := filepath.Dir(c.srcDir)
	content = shared.ResolveSnippets(content, srcPath, baseDir)
	content = convertDocusaurusAdmonitions(content)
	content = convertAlerts(content)
	content = convertEmoji(content)
	content = stripLeadingImages(content)
	res := splitMarkdown(content)

	if len(res.sections) == 0 {
		// No H2 sections — just copy as index.md
		if err := c.writeFile("index.md", content); err != nil {
			return err
		}
		title := res.title
		if title == "" {
			if c.version != "" {
				title = c.version
			} else {
				title = strings.TrimSuffix(filename,
					filepath.Ext(filename))
			}
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
	title := res.title
	if title == "" {
		// Prefer version label over raw filename stem
		if c.version != "" {
			title = c.version
		} else {
			title = strings.TrimSuffix(filename,
				filepath.Ext(filename))
		}
	}
	// If intro is empty (e.g. only had a banner image that
	// was stripped), generate a title heading
	if strings.TrimSpace(intro) == "" {
		intro = "# " + title + "\n"
	}
	if err := c.writeFile("index.md", intro); err != nil {
		return err
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
// if none exists. When files contain Docusaurus frontmatter
// with sidebar_position, the nav order is derived from it.
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

	// First pass: read files and parse frontmatter
	entries := make([]docEntry, 0, len(files))
	for i, f := range files {
		data, err := os.ReadFile(filepath.Join(c.srcDir, f))
		if err != nil {
			return err
		}

		content := string(data)

		// Parse and strip frontmatter
		fm, stripped := parseFrontmatter(content)
		if fm.Title != "" || fm.HasPosition {
			content = stripped
		}

		// Strip SPDX license comments
		content = stripSPDXComment(content)

		srcPath := filepath.Join(c.srcDir, f)
		baseDir := filepath.Dir(c.srcDir)
		content = shared.ResolveSnippets(content, srcPath, baseDir)
		content = convertDocusaurusAdmonitions(content)
		content = convertAlerts(content)
		content = convertEmoji(content)
		content = stripLeadingImages(content)

		title := fm.Title
		if title == "" {
			title = extractTitle(content, f)
		}

		entry := docEntry{
			relPath: f,
			content: content,
			title:   title,
			order:   i,
		}
		if fm.HasPosition {
			entry.order = fm.SidebarPosition
			entry.hasOrder = true
		}
		entries = append(entries, entry)
	}

	// Sort by frontmatter position when any entry has one
	hasAnyOrder := false
	for _, e := range entries {
		if e.hasOrder {
			hasAnyOrder = true
			break
		}
	}
	if hasAnyOrder {
		sort.Slice(entries, func(i, j int) bool {
			return entries[i].order < entries[j].order
		})
	}

	// Read _category_.json files for subdirectory labels
	catLabels := make(map[string]string)
	for _, e := range entries {
		dir := filepath.Dir(e.relPath)
		if dir == "." || dir == "" {
			continue
		}
		if _, ok := catLabels[dir]; ok {
			continue
		}
		srcSubdir := filepath.Join(c.srcDir, dir)
		if meta := readCategoryJSON(srcSubdir); meta != nil {
			catLabels[dir] = meta.Label
		}
	}

	// Second pass: write files and build file entries
	for i, e := range entries {
		outName := e.relPath
		lower := strings.ToLower(filepath.Base(e.relPath))
		if lower == "readme.md" {
			outName = filepath.Join(
				filepath.Dir(e.relPath), "index.md")
		}

		if err := c.writeFile(outName, e.content); err != nil {
			return err
		}
		c.files = append(c.files, &shared.FileEntry{
			Path:  outName,
			Title: e.title,
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

// imageExts lists file extensions treated as image assets.
var imageExts = map[string]bool{
	".png": true, ".jpg": true, ".jpeg": true,
	".gif": true, ".svg": true, ".webp": true,
	".ico": true,
}

// copyImages copies image files from the source directory tree
// into the output directory, preserving relative paths.
func (c *Converter) copyImages() error {
	return filepath.WalkDir(c.srcDir,
		func(path string, d os.DirEntry, err error) error {
			if err != nil || d.IsDir() {
				return err
			}
			ext := strings.ToLower(filepath.Ext(d.Name()))
			if !imageExts[ext] {
				return nil
			}
			rel, err := filepath.Rel(c.srcDir, path)
			if err != nil {
				return err
			}
			outPath := filepath.Join(c.outDir, rel)
			if err := os.MkdirAll(
				filepath.Dir(outPath), 0755); err != nil {
				return err
			}
			data, err := os.ReadFile(path)
			if err != nil {
				return err
			}
			return os.WriteFile(outPath, data, 0644)
		})
}
