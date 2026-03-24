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

// ReSnippet matches pymdownx.snippets include directives like
// --8<-- "filename" (with optional leading whitespace).
var ReSnippet = regexp.MustCompile(
	`^\s*--8<--\s+"([^"]+)"\s*$`)

// ResolveSnippets replaces pymdownx.snippets include lines with
// the referenced file's content. It searches for the included
// file relative to filePath's directory, then relative to
// baseDir (typically the repo root / parent of src_subdir).
func ResolveSnippets(
	content, filePath, baseDir string,
) string {
	lines := strings.Split(content, "\n")
	var result []string
	changed := false
	for _, line := range lines {
		m := ReSnippet.FindStringSubmatch(line)
		if m == nil {
			result = append(result, line)
			continue
		}
		ref := m[1]
		// Try relative to the file's directory first
		candidates := []string{
			filepath.Join(filepath.Dir(filePath), ref),
		}
		// Then try relative to baseDir (repo root)
		if baseDir != "" {
			candidates = append(candidates,
				filepath.Join(baseDir, ref))
		}
		var data []byte
		for _, cand := range candidates {
			var err error
			data, err = os.ReadFile(cand)
			if err == nil {
				break
			}
		}
		if data != nil {
			// Insert file content (trim trailing newline to
			// avoid double blank lines)
			result = append(result,
				strings.TrimRight(string(data), "\n"))
			changed = true
		} else {
			// Leave the directive as-is if we can't resolve
			result = append(result, line)
		}
	}
	if !changed {
		return content
	}
	return strings.Join(result, "\n")
}
