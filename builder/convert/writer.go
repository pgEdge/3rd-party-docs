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
	"strings"
)

// MarkdownWriter accumulates Markdown output with awareness of
// block structure, indentation, and whitespace management.
type MarkdownWriter struct {
	buf    strings.Builder
	indent string
	// atLineStart tracks whether we're at the beginning of a line
	atLineStart bool
	// lastBlank tracks whether the last thing written was a blank line
	lastBlank bool
	// inCodeBlock tracks whether we're inside a fenced code block
	inCodeBlock bool
	// suppressNewlines prevents blank line insertion (for inline contexts)
	suppressNewlines bool
}

// NewMarkdownWriter creates a new writer.
func NewMarkdownWriter() *MarkdownWriter {
	return &MarkdownWriter{
		atLineStart: true,
		lastBlank:   true,
	}
}

// String returns the accumulated Markdown content.
func (w *MarkdownWriter) String() string {
	return w.buf.String()
}

// WriteString writes raw text to the output.
func (w *MarkdownWriter) WriteString(s string) {
	if s == "" {
		return
	}
	w.buf.WriteString(s)
	w.atLineStart = strings.HasSuffix(s, "\n")
	w.lastBlank = false
}

// Write writes text, applying current indentation at line starts.
func (w *MarkdownWriter) Write(s string) {
	if s == "" {
		return
	}
	lines := strings.Split(s, "\n")
	for i, line := range lines {
		if i > 0 {
			w.buf.WriteString("\n")
			w.atLineStart = true
		}
		if line != "" {
			if w.atLineStart && w.indent != "" && !w.inCodeBlock {
				w.buf.WriteString(w.indent)
			}
			w.buf.WriteString(line)
			w.atLineStart = false
			w.lastBlank = false
		}
	}
}

// Newline writes a single newline.
func (w *MarkdownWriter) Newline() {
	w.buf.WriteString("\n")
	w.atLineStart = true
}

// BlankLine ensures there's a blank line (paragraph separator).
// Won't produce multiple consecutive blank lines.
func (w *MarkdownWriter) BlankLine() {
	if w.suppressNewlines {
		return
	}
	if w.lastBlank {
		return
	}
	if !w.atLineStart {
		w.buf.WriteString("\n")
	}
	w.buf.WriteString("\n")
	w.atLineStart = true
	w.lastBlank = true
}

// EnsureNewline makes sure we're at the start of a new line.
func (w *MarkdownWriter) EnsureNewline() {
	if !w.atLineStart {
		w.buf.WriteString("\n")
		w.atLineStart = true
	}
}

// PushIndent adds indentation for nested content.
func (w *MarkdownWriter) PushIndent(indent string) {
	w.indent += indent
}

// PopIndent removes the most recently pushed indentation.
func (w *MarkdownWriter) PopIndent(indent string) {
	if strings.HasSuffix(w.indent, indent) {
		w.indent = w.indent[:len(w.indent)-len(indent)]
	}
}

// StartCodeBlock begins a fenced code block.
func (w *MarkdownWriter) StartCodeBlock(lang string) {
	w.BlankLine()
	if lang != "" {
		w.WriteString("```" + lang + "\n")
	} else {
		w.WriteString("```\n")
	}
	w.inCodeBlock = true
	w.atLineStart = true
}

// EndCodeBlock ends a fenced code block.
func (w *MarkdownWriter) EndCodeBlock() {
	w.EnsureNewline()
	w.WriteString("```\n")
	w.inCodeBlock = false
	w.atLineStart = true
	w.lastBlank = false
}

// Heading writes a Markdown heading. The id parameter is accepted
// for backward compatibility but ignored — anchor IDs are emitted
// globally by convertNode via <a id=""> tags.
func (w *MarkdownWriter) Heading(level int, text, _ string) {
	w.BlankLine()
	prefix := strings.Repeat("#", level)
	w.WriteString(prefix + " " + text + "\n")
	w.atLineStart = true
	w.lastBlank = false
}

// Admonition writes an mkdocs-material admonition block.
func (w *MarkdownWriter) Admonition(kind string) {
	w.BlankLine()
	w.WriteString("!!! " + kind + "\n\n")
	w.atLineStart = true
	w.lastBlank = true
}

// IsAtLineStart returns whether the writer is at the start of a line.
func (w *MarkdownWriter) IsAtLineStart() bool {
	return w.atLineStart
}

// SetSuppressNewlines controls blank line suppression.
func (w *MarkdownWriter) SetSuppressNewlines(suppress bool) {
	w.suppressNewlines = suppress
}

// InCodeBlock returns whether we're in a code block.
func (w *MarkdownWriter) InCodeBlock() bool {
	return w.inCodeBlock
}

// Len returns the current length of the buffer.
func (w *MarkdownWriter) Len() int {
	return w.buf.Len()
}
