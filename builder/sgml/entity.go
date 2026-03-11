//-------------------------------------------------------------------------
//
// pgEdge PostgreSQL Docs
//
// Copyright (c) 2026, pgEdge, Inc.
// This software is released under The PostgreSQL License
//
//-------------------------------------------------------------------------

package sgml

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// EntityResolver handles SGML entity declaration and expansion.
type EntityResolver struct {
	// baseDir is the directory containing the SGML source files.
	baseDir string
	// entities maps entity names to their replacement text or file content.
	entities map[string]string
	// expanding tracks entities currently being expanded (cycle detection).
	expanding map[string]bool
	// maxDepth limits recursive entity expansion.
	maxDepth int
	// warnings accumulates non-fatal issues.
	warnings []string
}

// NewEntityResolver creates a resolver rooted at the given directory.
func NewEntityResolver(baseDir string) *EntityResolver {
	return &EntityResolver{
		baseDir:   baseDir,
		entities:  make(map[string]string),
		expanding: make(map[string]bool),
		maxDepth:  50,
	}
}

// Predefined character entities that would normally come from the DTD.
var predefinedEntities = map[string]string{
	"lt":     "<",
	"gt":     ">",
	"amp":    "&",
	"quot":   "\"",
	"apos":   "'",
	"nbsp":   "\u00a0",
	"mdash":  "\u2014",
	"ndash":  "\u2013",
	"lsquo":  "\u2018",
	"rsquo":  "\u2019",
	"ldquo":  "\u201c",
	"rdquo":  "\u201d",
	"bull":   "\u2022",
	"copy":   "\u00a9",
	"reg":    "\u00ae",
	"trade":  "\u2122",
	"hellip": "\u2026",
	"pi":     "\u03c0",
}

var (
	// Matches <!ENTITY name SYSTEM "file.sgml">
	reEntityFile = regexp.MustCompile(
		`<!ENTITY\s+(%?\s*[\w.-]+)\s+SYSTEM\s+"([^"]+)"\s*>`)
	// Matches <!ENTITY name "value">
	reEntityText = regexp.MustCompile(
		`<!ENTITY\s+([\w.-]+)\s+"([^"]*?)"\s*>`)
	// Matches single-quoted text entities: <!ENTITY name 'value'>
	reEntityTextSQ = regexp.MustCompile(
		`<!ENTITY\s+([\w.-]+)\s+'([^']*?)'\s*>`)
	// Matches &entity; references in text
	reEntityRef = regexp.MustCompile(`&([\w.-]+);`)
	// Matches %entity; parameter entity references
	reParamRef = regexp.MustCompile(`%([\w.-]+);`)
	// Matches numeric character references &#NN; or &#xHH;
	reCharRef = regexp.MustCompile(`&#(x?[0-9a-fA-F]+);`)
)

// ResolveFile reads the given SGML file, processes its DOCTYPE
// declarations, resolves all entities, and returns the document body
// with all entities expanded.
func (r *EntityResolver) ResolveFile(filename string) (string, error) {
	path := filepath.Join(r.baseDir, filename)
	data, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("reading %s: %w", filename, err)
	}
	content := string(data)

	// Extract and process DOCTYPE internal subset if present
	body := r.processDoctype(content)

	// Expand all entity references in the body
	expanded, err := r.expandEntities(body, 0)
	if err != nil {
		return "", fmt.Errorf("expanding entities in %s: %w", filename, err)
	}

	return expanded, nil
}

// processDoctype extracts entity declarations from the DOCTYPE
// internal subset and returns the document body.
func (r *EntityResolver) processDoctype(content string) string {
	// Find DOCTYPE declaration
	dtStart := strings.Index(content, "<!DOCTYPE")
	if dtStart == -1 {
		return content
	}

	// Find the internal subset between [ and ]
	bracketStart := strings.Index(content[dtStart:], "[")
	if bracketStart == -1 {
		// No internal subset; find end of DOCTYPE
		dtEnd := strings.Index(content[dtStart:], ">")
		if dtEnd == -1 {
			return content
		}
		return content[dtStart+dtEnd+1:]
	}
	bracketStart += dtStart

	// Find matching closing bracket, accounting for nesting
	depth := 1
	bracketEnd := -1
	for i := bracketStart + 1; i < len(content); i++ {
		switch content[i] {
		case '[':
			depth++
		case ']':
			depth--
			if depth == 0 {
				bracketEnd = i
			}
		}
		if bracketEnd != -1 {
			break
		}
	}
	if bracketEnd == -1 {
		return content
	}

	subset := content[bracketStart+1 : bracketEnd]

	// Process the internal subset: expand parameter entities first,
	// then extract entity declarations
	r.processSubset(subset)

	// Find the > that closes the DOCTYPE
	dtEnd := strings.Index(content[bracketEnd:], ">")
	if dtEnd == -1 {
		return ""
	}

	return strings.TrimSpace(content[bracketEnd+dtEnd+1:])
}

// processSubset extracts entity declarations from a DOCTYPE subset,
// expanding parameter entity references along the way.
func (r *EntityResolver) processSubset(subset string) {
	// First pass: extract parameter entity declarations and expand them
	expanded := r.expandParamEntitiesInSubset(subset)

	// Second pass: extract general entity declarations
	r.extractEntityDecls(expanded)
}

// expandParamEntitiesInSubset handles %entity; references within
// the DOCTYPE subset. These expand to file contents which may
// contain more entity declarations. Performs multiple passes to
// handle nested parameter entity references (e.g., filelist.sgml
// includes allfiles.sgml which defines refentry entities).
func (r *EntityResolver) expandParamEntitiesInSubset(subset string) string {
	result := subset

	for pass := 0; pass < 10; pass++ {
		// Extract parameter entity declarations from current text
		for _, m := range reEntityFile.FindAllStringSubmatch(result, -1) {
			name := strings.TrimSpace(m[1])
			file := m[2]
			if strings.HasPrefix(name, "%") {
				name = strings.TrimSpace(name[1:])
				if _, exists := r.entities["%"+name]; exists {
					continue // already loaded
				}
				// Read the referenced file
				path := filepath.Join(r.baseDir, file)
				data, err := os.ReadFile(path)
				if err != nil {
					continue
				}
				r.entities["%"+name] = string(data)
			}
		}

		// Expand parameter entity references
		changed := false
		expanded := reParamRef.ReplaceAllStringFunc(result, func(ref string) string {
			m := reParamRef.FindStringSubmatch(ref)
			if m == nil {
				return ref
			}
			name := m[1]
			if val, ok := r.entities["%"+name]; ok {
				changed = true
				return val
			}
			return ref
		})

		result = expanded
		if !changed {
			break // no more expansions possible
		}
	}

	return result
}

// extractEntityDecls pulls <!ENTITY ...> declarations from text.
func (r *EntityResolver) extractEntityDecls(text string) {
	// File entities (non-parameter)
	for _, m := range reEntityFile.FindAllStringSubmatch(text, -1) {
		name := strings.TrimSpace(m[1])
		if strings.HasPrefix(name, "%") {
			continue // skip parameter entities, already handled
		}
		file := m[2]
		r.entities[name] = "\x00FILE:" + file
	}

	// Text entities (double-quoted)
	for _, m := range reEntityText.FindAllStringSubmatch(text, -1) {
		name := m[1]
		value := m[2]
		// Don't overwrite file entities with text entity matches
		// from the SYSTEM keyword line
		if _, exists := r.entities[name]; !exists {
			r.entities[name] = value
		}
	}

	// Text entities (single-quoted)
	for _, m := range reEntityTextSQ.FindAllStringSubmatch(text, -1) {
		name := m[1]
		value := m[2]
		if _, exists := r.entities[name]; !exists {
			r.entities[name] = value
		}
	}
}

// expandEntities recursively expands &entity; references in text.
func (r *EntityResolver) expandEntities(text string, depth int) (string, error) {
	if depth > r.maxDepth {
		return text, fmt.Errorf("entity expansion exceeded max depth %d", r.maxDepth)
	}

	// Expand character references first
	text = r.expandCharRefs(text)

	// Expand named entity references
	var lastErr error
	result := reEntityRef.ReplaceAllStringFunc(text, func(ref string) string {
		m := reEntityRef.FindStringSubmatch(ref)
		if m == nil {
			return ref
		}
		name := m[1]

		// Check predefined entities
		if val, ok := predefinedEntities[name]; ok {
			return val
		}

		// Check our declared entities
		val, ok := r.entities[name]
		if !ok {
			// Unknown entity — leave as-is and warn
			return ref
		}

		// Cycle detection
		if r.expanding[name] {
			lastErr = fmt.Errorf("circular entity reference: %s", name)
			return ref
		}

		// File entity: read and expand the file
		if strings.HasPrefix(val, "\x00FILE:") {
			filename := val[6:]
			path := filepath.Join(r.baseDir, filename)
			data, err := os.ReadFile(path)
			if err != nil {
				// Try common subdirectories (entities declared in
				// ref/allfiles.sgml and func/allfiles.sgml use paths
				// relative to their own directory)
				found := false
				for _, subdir := range []string{"ref", "func"} {
					altPath := filepath.Join(r.baseDir, subdir, filename)
					altData, altErr := os.ReadFile(altPath)
					if altErr == nil {
						data = altData
						found = true
						break
					}
				}
				if !found {
					r.warnings = append(r.warnings,
						fmt.Sprintf("missing entity file %s for &%s; (skipping)", filename, name))
					return "<!-- entity " + name + " not available -->"
				}
			}

			r.expanding[name] = true
			expanded, err := r.expandEntities(string(data), depth+1)
			delete(r.expanding, name)
			if err != nil {
				lastErr = err
				return string(data)
			}
			return expanded
		}

		// Text entity: expand the value
		r.expanding[name] = true
		expanded, err := r.expandEntities(val, depth+1)
		delete(r.expanding, name)
		if err != nil {
			lastErr = err
			return val
		}
		return expanded
	})

	return result, lastErr
}

// expandCharRefs expands numeric character references like &#123;
// and &#x1F;
func (r *EntityResolver) expandCharRefs(text string) string {
	return reCharRef.ReplaceAllStringFunc(text, func(ref string) string {
		m := reCharRef.FindStringSubmatch(ref)
		if m == nil {
			return ref
		}
		numStr := m[1]
		var codepoint int64
		if strings.HasPrefix(numStr, "x") || strings.HasPrefix(numStr, "X") {
			_, err := fmt.Sscanf(numStr[1:], "%x", &codepoint)
			if err != nil {
				return ref
			}
		} else {
			_, err := fmt.Sscanf(numStr, "%d", &codepoint)
			if err != nil {
				return ref
			}
		}
		return string(rune(codepoint))
	})
}

// EntityCount returns the number of resolved entities (for testing).
func (r *EntityResolver) EntityCount() int {
	return len(r.entities)
}

// HasEntity checks if an entity is defined (for testing).
func (r *EntityResolver) HasEntity(name string) bool {
	_, ok := r.entities[name]
	return ok
}

// Warnings returns non-fatal warnings from entity resolution.
func (r *EntityResolver) Warnings() []string {
	return r.warnings
}
