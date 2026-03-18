//-------------------------------------------------------------------------
//
// pgEdge PostgreSQL Docs
//
// Copyright (c) 2026, pgEdge, Inc.
// This software is released under The PostgreSQL License
//
//-------------------------------------------------------------------------

package convert

import (
	"regexp"
	"strings"
)

// markdownToHTML converts simple Markdown content to HTML.
// This is used for HTML table cells where Python-Markdown's md_in_html
// extension cannot process Markdown inside <table> elements.
func markdownToHTML(md string) string {
	md = strings.TrimSpace(md)
	if md == "" {
		return ""
	}

	// Split into paragraphs (blank-line separated)
	paragraphs := splitParagraphs(md)

	var parts []string
	for _, para := range paragraphs {
		para = strings.TrimSpace(para)
		if para == "" {
			continue
		}

		// Check if it's already an HTML block (starts with <)
		if isHTMLBlock(para) {
			parts = append(parts, para)
			continue
		}

		// Check if it's a fenced code block
		if strings.HasPrefix(para, "```") {
			parts = append(parts, convertCodeBlock(para))
			continue
		}

		// Convert inline Markdown to HTML
		html := convertInlineMarkdown(para)
		// Wrap in <p> if we have multiple paragraphs
		if len(paragraphs) > 1 {
			parts = append(parts, "<p>"+html+"</p>")
		} else {
			parts = append(parts, html)
		}
	}

	return strings.Join(parts, "\n")
}

// splitParagraphs splits text on blank lines, preserving code blocks.
func splitParagraphs(text string) []string {
	var paragraphs []string
	var current strings.Builder
	inCodeBlock := false
	lines := strings.Split(text, "\n")

	for i := 0; i < len(lines); i++ {
		line := lines[i]

		if strings.HasPrefix(line, "```") {
			inCodeBlock = !inCodeBlock
		}

		if !inCodeBlock && line == "" {
			if current.Len() > 0 {
				paragraphs = append(paragraphs, current.String())
				current.Reset()
			}
			continue
		}

		if current.Len() > 0 {
			current.WriteString("\n")
		}
		current.WriteString(line)
	}

	if current.Len() > 0 {
		paragraphs = append(paragraphs, current.String())
	}

	return paragraphs
}

// isHTMLBlock checks if text starts with an HTML tag.
func isHTMLBlock(text string) bool {
	trimmed := strings.TrimSpace(text)
	return strings.HasPrefix(trimmed, "<a ") ||
		strings.HasPrefix(trimmed, "<pre>") ||
		strings.HasPrefix(trimmed, "<div") ||
		strings.HasPrefix(trimmed, "<img") ||
		strings.HasPrefix(trimmed, "<sup>") ||
		strings.HasPrefix(trimmed, "<sub>")
}

// convertCodeBlock converts a fenced code block to HTML <pre><code>.
func convertCodeBlock(block string) string {
	lines := strings.Split(block, "\n")
	if len(lines) < 2 {
		return block
	}

	// Extract language from opening fence
	opening := strings.TrimPrefix(lines[0], "```")
	lang := strings.TrimSpace(opening)

	// Find closing fence
	var codeLines []string
	for i := 1; i < len(lines); i++ {
		if strings.HasPrefix(lines[i], "```") {
			break
		}
		codeLines = append(codeLines, lines[i])
	}

	code := strings.Join(codeLines, "\n")
	code = htmlEscapeCode(code)

	if lang != "" {
		return "<pre><code class=\"language-" + lang + "\">" +
			code + "</code></pre>"
	}
	return "<pre><code>" + code + "</code></pre>"
}

var (
	// Order matters: bold before italic, double-backtick before single
	reDoubleBacktick = regexp.MustCompile("``([^`]+)``")
	reSingleBacktick = regexp.MustCompile("`([^`]+)`")
	reBold           = regexp.MustCompile(`\*\*([^\*]+)\*\*`)
	reItalic         = regexp.MustCompile(`\*([^\*]+)\*`)
	reImage          = regexp.MustCompile(`!\[([^\]]*)\]\(([^)]+)\)`)
	reLink           = regexp.MustCompile(`\[([^\]]+)\]\(([^)]+)\)`)
)

// convertInlineMarkdown converts inline Markdown formatting to HTML.
func convertInlineMarkdown(text string) string {
	// Replace double backticks first (`` code with backtick ``)
	text = reDoubleBacktick.ReplaceAllStringFunc(text, func(m string) string {
		inner := reDoubleBacktick.FindStringSubmatch(m)[1]
		return "<code>" + htmlEscapeCode(inner) + "</code>"
	})

	// Replace single backticks
	text = reSingleBacktick.ReplaceAllStringFunc(text, func(m string) string {
		inner := reSingleBacktick.FindStringSubmatch(m)[1]
		return "<code>" + htmlEscapeCode(inner) + "</code>"
	})

	// Replace bold (**text**)
	text = reBold.ReplaceAllString(text, "<strong>$1</strong>")

	// Replace italic (*text*)
	text = reItalic.ReplaceAllString(text, "<em>$1</em>")

	// Replace images ![alt](src) — before links to avoid partial match
	text = reImage.ReplaceAllString(text, `<img src="$2" alt="$1">`)

	// Replace links [text](url)
	text = reLink.ReplaceAllString(text, `<a href="$2">$1</a>`)

	// Convert line breaks within a paragraph to <br>
	text = strings.ReplaceAll(text, "\n", "<br>\n")

	return text
}

// htmlEscapeCode escapes HTML special characters in code content.
func htmlEscapeCode(s string) string {
	s = strings.ReplaceAll(s, "&", "&amp;")
	s = strings.ReplaceAll(s, "<", "&lt;")
	s = strings.ReplaceAll(s, ">", "&gt;")
	return s
}
