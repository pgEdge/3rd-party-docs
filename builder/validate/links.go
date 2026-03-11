//-------------------------------------------------------------------------
//
// pgEdge PostgreSQL Docs
//
// Copyright (c) 2026, pgEdge, Inc.
// This software is released under The PostgreSQL License
//
//-------------------------------------------------------------------------

// Package validate checks generated Markdown files for broken
// links and missing anchors.
package validate

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// Result holds validation findings.
type Result struct {
	BrokenLinks    []Issue
	MissingAnchors []Issue
}

// Issue describes a single validation problem.
type Issue struct {
	File    string
	Line    int
	Message string
}

// reMarkdownLink matches [text](path#anchor) links.
var reMarkdownLink = regexp.MustCompile(`\[([^\]]*)\]\(([^)]+)\)`)

// reAnchorID matches <a id="..."></a> and { #id } heading attributes.
var reAnchorID = regexp.MustCompile(
	`(?:<a\s+id="([^"]+)"|#\s*([\w-]+)\s*})`)

// reHeading matches Markdown headings.
var reHeading = regexp.MustCompile(
	`^#{1,6}\s+(.+?)(?:\s*\{[^}]*\})?\s*$`)

// testHookAfterCollect is called between anchor collection and
// link checking during tests. It is nil in production.
var testHookAfterCollect func()

// ValidateDir checks all .md files in a directory for link
// integrity.
func ValidateDir(docsDir string) (*Result, error) {
	anchors, err := collectAnchors(docsDir)
	if err != nil {
		return nil, fmt.Errorf("walking docs dir: %w", err)
	}

	if testHookAfterCollect != nil {
		testHookAfterCollect()
	}

	result, err := checkLinks(docsDir, anchors)
	if err != nil {
		return nil, fmt.Errorf("checking links: %w", err)
	}

	return result, nil
}

// collectAnchors walks docsDir and returns a map of relative file
// paths to their sets of defined anchor IDs.
func collectAnchors(
	docsDir string,
) (map[string]map[string]bool, error) {
	anchors := make(map[string]map[string]bool)

	err := filepath.Walk(docsDir,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() ||
				!strings.HasSuffix(path, ".md") {
				return nil
			}

			relPath, _ := filepath.Rel(docsDir, path)
			data, err := os.ReadFile(path)
			if err != nil {
				return err
			}
			content := string(data)

			fileAnchors := make(map[string]bool)
			for _, m := range reAnchorID.FindAllStringSubmatch(
				content, -1) {
				if m[1] != "" {
					fileAnchors[m[1]] = true
				}
				if m[2] != "" {
					fileAnchors[m[2]] = true
				}
			}

			for _, line := range strings.Split(content, "\n") {
				m := reHeading.FindStringSubmatch(line)
				if m != nil {
					heading := m[1]
					anchor := headingToAnchor(heading)
					fileAnchors[anchor] = true
				}
			}

			anchors[relPath] = fileAnchors
			return nil
		})
	if err != nil {
		return nil, err
	}

	return anchors, nil
}

// checkLinks walks docsDir and validates all internal Markdown
// links against the provided anchor map.
func checkLinks(
	docsDir string,
	anchors map[string]map[string]bool,
) (*Result, error) {
	result := &Result{}

	err := filepath.Walk(docsDir,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() ||
				!strings.HasSuffix(path, ".md") {
				return nil
			}

			relPath, _ := filepath.Rel(docsDir, path)
			data, err := os.ReadFile(path)
			if err != nil {
				return err
			}
			content := string(data)

			lines := strings.Split(content, "\n")
			for lineNum, line := range lines {
				for _, m := range reMarkdownLink.FindAllStringSubmatch(
					line, -1) {
					href := m[2]

					if strings.HasPrefix(href, "http://") ||
						strings.HasPrefix(href, "https://") ||
						strings.HasPrefix(href, "mailto:") {
						continue
					}

					targetFile := ""
					targetAnchor := ""
					if strings.HasPrefix(href, "#") {
						targetFile = relPath
						targetAnchor = href[1:]
					} else {
						parts := strings.SplitN(href, "#", 2)
						resolved := filepath.Join(
							filepath.Dir(relPath), parts[0])
						targetFile = filepath.Clean(resolved)
						if len(parts) > 1 {
							targetAnchor = parts[1]
						}
					}

					if _, ok := anchors[targetFile]; !ok {
						result.BrokenLinks = append(
							result.BrokenLinks, Issue{
								File: relPath,
								Line: lineNum + 1,
								Message: fmt.Sprintf(
									"broken link to %q"+
										" (file not found)",
									href),
							})
						continue
					}

					if targetAnchor != "" {
						if !anchors[targetFile][targetAnchor] {
							result.MissingAnchors = append(
								result.MissingAnchors, Issue{
									File: relPath,
									Line: lineNum + 1,
									Message: fmt.Sprintf(
										"missing anchor"+
											" %q in %s",
										targetAnchor,
										targetFile),
								})
						}
					}
				}
			}
			return nil
		})
	if err != nil {
		return nil, err
	}

	return result, nil
}

// headingToAnchor converts a heading to its auto-generated anchor
// ID.
func headingToAnchor(heading string) string {
	// Remove inline markup
	heading = regexp.MustCompile(`[*_\x60]`).ReplaceAllString(
		heading, "")
	heading = strings.ToLower(heading)

	var b strings.Builder
	for _, r := range heading {
		switch {
		case r >= 'a' && r <= 'z':
			b.WriteRune(r)
		case r >= '0' && r <= '9':
			b.WriteRune(r)
		case r == '-' || r == '_':
			b.WriteRune(r)
		case r == ' ':
			b.WriteRune('-')
		}
	}

	result := b.String()
	for strings.Contains(result, "--") {
		result = strings.ReplaceAll(result, "--", "-")
	}
	return strings.Trim(result, "-")
}
