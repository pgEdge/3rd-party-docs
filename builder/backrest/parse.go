//-------------------------------------------------------------------------
//
// pgEdge PostgreSQL Docs
//
// Copyright (c) 2026, pgEdge, Inc.
// This software is released under The PostgreSQL License
//
//-------------------------------------------------------------------------

package backrest

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/pgEdge/postgresql-docs/builder/sgml"
)

// reEntity matches <!ENTITY name SYSTEM "path"> declarations.
var reEntity = regexp.MustCompile(
	`<!ENTITY\s+(\S+)\s+SYSTEM\s+"([^"]+)"[^>]*>`)

// reEntityRef matches &name; entity references.
var reEntityRef = regexp.MustCompile(`&(\w[\w.-]*);`)

// resolveEntities resolves external SYSTEM entities in XML content.
// It reads entity declarations from the internal DTD subset,
// loads the referenced files, and replaces entity references.
func resolveEntities(content, baseDir string) string {
	// Extract entity declarations
	entities := make(map[string]string)
	for _, m := range reEntity.FindAllStringSubmatch(content, -1) {
		name, path := m[1], m[2]
		fullPath := filepath.Join(baseDir, path)
		data, err := os.ReadFile(fullPath)
		if err != nil {
			continue
		}
		entities[name] = string(data)
	}

	if len(entities) == 0 {
		return content
	}

	// Replace entity references (multiple passes for nesting)
	for i := 0; i < 3; i++ {
		prev := content
		content = reEntityRef.ReplaceAllStringFunc(content,
			func(m string) string {
				name := reEntityRef.FindStringSubmatch(m)[1]
				if val, ok := entities[name]; ok {
					return val
				}
				return m
			})
		if content == prev {
			break
		}
	}

	return content
}

// parseXMLFile reads a pgBackRest XML file and returns a node tree.
// It uses Go's encoding/xml to avoid the SGML parser's htmlElements
// filter which conflicts with pgBackRest element names.
// External entities (<!ENTITY name SYSTEM "path">) are resolved
// relative to the file's directory.
func parseXMLFile(path string) (*sgml.Node, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("reading %s: %w", path, err)
	}
	content := string(data)
	dir := filepath.Dir(path)

	// Resolve external entities
	content = resolveEntities(content, dir)

	return parseXMLString(content)
}

// parseXMLString parses an XML string into an sgml.Node tree.
func parseXMLString(content string) (*sgml.Node, error) {
	// Strip the DOCTYPE declaration since encoding/xml chokes on it
	content = stripDoctype(content)

	dec := xml.NewDecoder(strings.NewReader(content))
	dec.Strict = false
	dec.Entity = xml.HTMLEntity

	root := &sgml.Node{Type: sgml.ElementNode, Tag: "root"}
	stack := []*sgml.Node{root}

	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("parsing XML: %w", err)
		}

		parent := stack[len(stack)-1]

		switch t := tok.(type) {
		case xml.StartElement:
			node := &sgml.Node{
				Type:  sgml.ElementNode,
				Tag:   t.Name.Local,
				Attrs: make(map[string]string),
			}
			for _, a := range t.Attr {
				node.Attrs[a.Name.Local] = a.Value
			}
			parent.AppendChild(node)
			stack = append(stack, node)

		case xml.EndElement:
			if len(stack) > 1 {
				stack = stack[:len(stack)-1]
			}

		case xml.CharData:
			text := string(t)
			if text != "" {
				parent.AppendChild(&sgml.Node{
					Type: sgml.TextNode,
					Text: text,
				})
			}

		case xml.Comment:
			// skip comments

		case xml.ProcInst:
			// skip processing instructions
		}
	}

	// Return the <doc> element if present, otherwise root
	if doc := root.FindChild("doc"); doc != nil {
		return doc, nil
	}
	return root, nil
}

// stripDoctype removes <!DOCTYPE ...> from XML content.
func stripDoctype(s string) string {
	for {
		idx := strings.Index(s, "<!DOCTYPE")
		if idx == -1 {
			break
		}
		end := strings.Index(s[idx:], ">")
		if end == -1 {
			break
		}
		s = s[:idx] + s[idx+end+1:]
	}
	return s
}
