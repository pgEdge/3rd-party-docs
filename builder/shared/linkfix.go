//-------------------------------------------------------------------------
//
// pgEdge PostgreSQL Docs
//
// Copyright (c) 2026, pgEdge, Inc.
// This software is released under The PostgreSQL License
//
//-------------------------------------------------------------------------

package shared

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// reMdLink matches Markdown links: [text](target)
// It avoids matching image links (![...]).
var reMdLink = regexp.MustCompile(
	`(?:^|[^!])\[([^\]]*)\]\(([^)]+)\)`)

// FixBrokenLinksInDir walks outDir and fixes broken relative
// links in all Markdown files.
func FixBrokenLinksInDir(outDir string) error {
	return filepath.WalkDir(outDir,
		func(path string, d os.DirEntry, err error) error {
			if err != nil || d.IsDir() {
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
			rel, _ := filepath.Rel(outDir, path)
			fixed := fixBrokenLinks(content, rel, outDir)
			if fixed != content {
				return os.WriteFile(
					path, []byte(fixed), 0644)
			}
			return nil
		})
}

// fixBrokenLinks rewrites relative links in markdown content
// whose targets don't exist in the output directory. It:
//   - Strips directory prefixes to find relocated files
//   - Converts absolute /path.html links to relative .md links
//   - Removes links to non-existent targets (keeps the text)
//
// filePath is the path of the current file relative to outDir.
func fixBrokenLinks(
	content, filePath, outDir string,
) string {
	fileDir := filepath.Dir(filepath.Join(outDir, filePath))

	return reMdLink.ReplaceAllStringFunc(
		content, func(match string) string {
			sub := reMdLink.FindStringSubmatch(match)
			if len(sub) < 3 {
				return match
			}
			// Preserve any leading character (non-! char)
			prefix := ""
			if len(match) > 0 && match[0] != '[' {
				prefix = string(match[0])
			}
			text := sub[1]
			target := sub[2]

			// Skip external URLs and anchors
			if strings.HasPrefix(target, "http://") ||
				strings.HasPrefix(target, "https://") ||
				strings.HasPrefix(target, "mailto:") ||
				strings.HasPrefix(target, "#") {
				return match
			}

			// Split target into path and fragment
			path, fragment := target, ""
			if idx := strings.Index(target, "#"); idx >= 0 {
				path = target[:idx]
				fragment = target[idx:]
			}

			// Handle absolute /path.html links
			if strings.HasPrefix(path, "/") {
				newPath := rewriteAbsoluteLink(
					path, filePath, outDir)
				if newPath != "" {
					return prefix + "[" + text + "](" +
						newPath + fragment + ")"
				}
				return match // leave as-is (INFO level)
			}

			// Skip non-file links
			if path == "" {
				return match
			}

			// Check if the target exists
			absTarget := filepath.Join(fileDir, path)
			absTarget = filepath.Clean(absTarget)
			if fileExists(absTarget) {
				return match
			}

			// Handle README.md → index.md rename
			if strings.EqualFold(
				filepath.Base(path), "readme.md") {
				renamed := filepath.Join(
					filepath.Dir(path), "index.md")
				absRenamed := filepath.Join(
					fileDir, renamed)
				absRenamed = filepath.Clean(absRenamed)
				if fileExists(absRenamed) {
					return prefix + "[" + text + "](" +
						renamed + fragment + ")"
				}
			}

			// Try to find by basename in the output
			base := filepath.Base(path)
			if strings.EqualFold(base, "readme.md") {
				base = "index.md"
			}
			found := findFileInDir(outDir, base)
			if found != "" {
				rel, err := filepath.Rel(fileDir, found)
				if err == nil {
					return prefix + "[" + text + "](" +
						rel + fragment + ")"
				}
			}

			// Target not found — strip the link, keep text
			return prefix + text
		})
}

// rewriteAbsoluteLink converts an absolute /path.html link
// to a relative .md link if the corresponding file exists.
func rewriteAbsoluteLink(
	path, filePath, outDir string,
) string {
	// Strip leading /
	p := strings.TrimPrefix(path, "/")

	// Try .html → .md conversion
	if strings.HasSuffix(p, ".html") {
		mdPath := strings.TrimSuffix(p, ".html") + ".md"
		abs := filepath.Join(outDir, mdPath)
		if fileExists(abs) {
			fileDir := filepath.Dir(
				filepath.Join(outDir, filePath))
			rel, err := filepath.Rel(fileDir, abs)
			if err == nil {
				return rel
			}
		}
	}
	return ""
}

// findFileInDir searches outDir recursively for a file with
// the given base name. Returns the first match.
func findFileInDir(dir, baseName string) string {
	var result string
	lower := strings.ToLower(baseName)
	_ = filepath.WalkDir(dir,
		func(path string, d os.DirEntry, err error) error {
			if err != nil || d.IsDir() {
				return nil
			}
			if strings.ToLower(d.Name()) == lower {
				result = path
				return filepath.SkipAll
			}
			return nil
		})
	return result
}

// fileExists returns true if the path exists as a file.
func fileExists(path string) bool {
	info, err := os.Stat(path)
	return err == nil && !info.IsDir()
}
