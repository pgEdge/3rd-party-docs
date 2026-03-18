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
	"regexp"
	"strings"

	"github.com/pgEdge/postgresql-docs/builder/sgml"
)

// reVariable matches {[key]} substitution patterns.
var reVariable = regexp.MustCompile(`\{\[([^\]]+)\]\}`)

// collectVariables extracts static variables from <variable-list>
// elements in the tree. Variables with eval="y" or with if
// conditions are skipped unless they match default conditions.
func collectVariables(node *sgml.Node, vars map[string]string) {
	for _, vl := range node.FindDescendants("variable-list") {
		for _, v := range vl.FindChildren("variable") {
			key := v.GetAttr("key")
			if key == "" {
				continue
			}
			// Skip eval variables (require Perl runtime)
			if v.GetAttr("eval") == "y" {
				continue
			}
			// For conditional variables, only take the first
			// (default) value if we don't already have one
			ifCond := v.GetAttr("if")
			if ifCond != "" && ifCond != "default" {
				if _, exists := vars[key]; exists {
					continue
				}
			}
			val := strings.TrimSpace(v.TextContent())
			if val != "" {
				vars[key] = val
			}
		}
	}
}

// substituteVariables replaces {[key]} patterns in text.
// Performs multiple passes to resolve chained references.
// Unresolved variables are replaced with their key name in
// angle brackets (e.g. "<backup-id>") to avoid raw template
// syntax in the output.
func substituteVariables(text string, vars map[string]string) string {
	for i := 0; i < 5; i++ {
		prev := text
		text = reVariable.ReplaceAllStringFunc(text, func(m string) string {
			key := reVariable.FindStringSubmatch(m)[1]
			if val, ok := vars[key]; ok {
				return val
			}
			return m // leave for now
		})
		if text == prev {
			break
		}
	}
	// Replace any remaining unresolved variables
	text = reVariable.ReplaceAllStringFunc(text, func(m string) string {
		key := reVariable.FindStringSubmatch(m)[1]
		return "<" + key + ">"
	})
	return text
}
