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
	"regexp"
	"strings"
)

// Precompiled regexes for inline RST markup.
var (
	// :role:`text` or :role:`title <target>`
	reRole = regexp.MustCompile(
		":([a-zA-Z0-9_-]+):`([^`]+)`")

	// `text`:index: (pgAdmin index pattern)
	reIndexRole = regexp.MustCompile(
		"`([^`]+)`:index:")

	// **bold**
	reBold = regexp.MustCompile(`\*\*([^*]+)\*\*`)

	// *italic* (but not inside **)
	reItalic = regexp.MustCompile(`(?:^|[^*])\*([^*]+)\*(?:[^*]|$)`)

	// ``literal``
	reLiteral = regexp.MustCompile("``([^`]+)``")

	// `text <URL>`_ or `text <URL>`__
	reExternalLink = regexp.MustCompile(
		"`([^<]+)<([^>]+)>`_{1,2}")

	// `text`_ (backtick named reference)
	reBacktickRef = regexp.MustCompile("`([^`]+)`_(?:[^_]|$)")

	// word_ (standalone named reference — word followed by _)
	reStandaloneRef = regexp.MustCompile(`\b([A-Za-z][A-Za-z0-9]+)_\b`)

	// |substitution|_ (substitution with hyperlink)
	reSubstitutionLink = regexp.MustCompile(`\|([a-zA-Z0-9_-]+)\|_`)

	// |substitution|
	reSubstitution = regexp.MustCompile(`\|([a-zA-Z0-9_-]+)\|`)

	// RST backslash escape: \x (any character after backslash)
	reBackslashEscape = regexp.MustCompile(`\\(.)`)
)

// ConvertInline converts RST inline markup to Markdown.
func ConvertInline(
	text string,
	labelMap map[string]labelInfo,
	fileMap map[string]string,
	currentFile string,
	substitutions map[string]*Node,
	hyperlinkTargets map[string]string,
) string {
	if text == "" {
		return ""
	}

	// Step 1: Protect literal spans from further processing
	type placeholder struct {
		key     string
		replace string
	}
	var placeholders []placeholder
	phIdx := 0

	text = reLiteral.ReplaceAllStringFunc(text, func(m string) string {
		sub := reLiteral.FindStringSubmatch(m)
		key := "\x00LIT" + string(rune('A'+phIdx)) + "\x00"
		phIdx++
		placeholders = append(placeholders, placeholder{
			key: key, replace: "`" + sub[1] + "`"})
		return key
	})

	// Step 2: Index entries (strip the :index: wrapper)
	text = reIndexRole.ReplaceAllString(text, "$1")

	// Step 3: Roles
	text = reRole.ReplaceAllStringFunc(text, func(m string) string {
		sub := reRole.FindStringSubmatch(m)
		role := sub[1]
		content := sub[2]
		return convertRole(role, content, labelMap, fileMap,
			currentFile)
	})

	// Step 4: External links `text <URL>`_
	text = reExternalLink.ReplaceAllStringFunc(text, func(m string) string {
		sub := reExternalLink.FindStringSubmatch(m)
		title := strings.TrimSpace(sub[1])
		url := strings.TrimSpace(sub[2])
		return "[" + title + "](" + url + ")"
	})

	// Step 5: |substitution|_ (substitution + hyperlink ref)
	if hyperlinkTargets != nil {
		text = reSubstitutionLink.ReplaceAllStringFunc(text,
			func(m string) string {
				sub := reSubstitutionLink.FindStringSubmatch(m)
				name := sub[1]
				display := name
				if def, ok := substitutions[name]; ok {
					display = def.DirectiveArg
				}
				if url, ok := hyperlinkTargets[name]; ok {
					return "[" + display + "](" + url + ")"
				}
				return display
			})
	}

	// Step 6: Substitutions |name|
	text = reSubstitution.ReplaceAllStringFunc(text, func(m string) string {
		sub := reSubstitution.FindStringSubmatch(m)
		name := sub[1]
		if def, ok := substitutions[name]; ok {
			if def.DirectiveName == "image" {
				alt := def.Options["alt"]
				if alt == "" {
					alt = name
				}
				return "<img src=\"images/" +
					strings.TrimPrefix(def.DirectiveArg, "images/") +
					"\" alt=\"" + alt + "\">"
			}
			return def.DirectiveArg
		}
		return m
	})

	// Step 7: `text`_ backtick named references
	if hyperlinkTargets != nil {
		text = reBacktickRef.ReplaceAllStringFunc(text,
			func(m string) string {
				sub := reBacktickRef.FindStringSubmatch(m)
				name := sub[1]
				if url, ok := hyperlinkTargets[name]; ok {
					return "[" + name + "](" + url + ")"
				}
				// Not a known target — just strip backticks
				return name
			})
	}

	// Step 8: standalone word_ named references
	if hyperlinkTargets != nil {
		text = reStandaloneRef.ReplaceAllStringFunc(text,
			func(m string) string {
				sub := reStandaloneRef.FindStringSubmatch(m)
				name := sub[1]
				if url, ok := hyperlinkTargets[name]; ok {
					return "[" + name + "](" + url + ")"
				}
				return m // leave as-is if not a known target
			})
	}

	// Step 9: Backslash escapes
	text = reBackslashEscape.ReplaceAllString(text, "$1")

	// Restore literal placeholders
	for _, ph := range placeholders {
		text = strings.ReplaceAll(text, ph.key, ph.replace)
	}

	return text
}

// labelInfo holds cross-reference target information.
type labelInfo struct {
	File   string
	Anchor string
	Title  string
}

// convertRole converts an RST role to Markdown.
func convertRole(
	role, content string,
	labelMap map[string]labelInfo,
	fileMap map[string]string,
	currentFile string,
) string {
	switch role {
	case "ref":
		return convertRef(content, labelMap, currentFile)
	case "doc":
		return convertDoc(content, fileMap, currentFile)
	case "index":
		return content
	case "menuselection", "guilabel":
		return "**" + content + "**"
	case "kbd":
		return "`" + content + "`"
	case "file", "command", "program", "envvar", "option",
		"class", "func", "meth", "attr", "exc", "obj",
		"mod", "data", "const", "type", "term",
		"sql", "samp":
		return "`" + content + "`"
	case "code":
		return "`" + content + "`"
	case "pep":
		return "[PEP " + content + "](https://peps.python.org/pep-" +
			strings.TrimLeft(content, "0") + "/)"
	case "abbr":
		return content
	case "sup":
		return "<sup>" + content + "</sup>"
	case "sub":
		return "<sub>" + content + "</sub>"
	default:
		return "`" + content + "`"
	}
}

// convertRef converts a :ref: role to a Markdown link.
func convertRef(
	content string,
	labelMap map[string]labelInfo,
	currentFile string,
) string {
	title := ""
	target := content
	if idx := strings.Index(content, "<"); idx >= 0 {
		title = strings.TrimSpace(content[:idx])
		end := strings.Index(content, ">")
		if end > idx {
			target = content[idx+1 : end]
		}
	}

	if info, ok := labelMap[target]; ok {
		if title == "" {
			title = info.Title
			if title == "" {
				title = target
			}
		}
		link := resolveLink(currentFile, info.File, info.Anchor)
		return "[" + title + "](" + link + ")"
	}

	if title == "" {
		title = target
	}
	return title
}

// convertDoc converts a :doc: role to a Markdown link.
func convertDoc(
	content string,
	fileMap map[string]string,
	currentFile string,
) string {
	title := ""
	target := content
	if idx := strings.Index(content, "<"); idx >= 0 {
		title = strings.TrimSpace(content[:idx])
		end := strings.Index(content, ">")
		if end > idx {
			target = content[idx+1 : end]
		}
	}

	target = strings.TrimPrefix(target, "/")
	target = strings.TrimSuffix(target, ".rst")

	if mdPath, ok := fileMap[target]; ok {
		if title == "" {
			title = target
		}
		link := resolveLink(currentFile, mdPath, "")
		return "[" + title + "](" + link + ")"
	}

	if title == "" {
		title = target
	}
	return title
}

// resolveLink computes a relative link from currentFile to target.
func resolveLink(currentFile, targetFile, anchor string) string {
	if currentFile == targetFile {
		if anchor != "" {
			return "#" + anchor
		}
		return ""
	}

	fromParts := strings.Split(currentFile, "/")
	toParts := strings.Split(targetFile, "/")
	fromParts = fromParts[:len(fromParts)-1]

	common := 0
	for common < len(fromParts) && common < len(toParts) {
		if fromParts[common] == toParts[common] {
			common++
		} else {
			break
		}
	}

	var parts []string
	for i := common; i < len(fromParts); i++ {
		parts = append(parts, "..")
	}
	for i := common; i < len(toParts); i++ {
		parts = append(parts, toParts[i])
	}

	link := strings.Join(parts, "/")
	if link == "" {
		link = targetFile
	}
	if anchor != "" {
		link += "#" + anchor
	}
	return link
}
