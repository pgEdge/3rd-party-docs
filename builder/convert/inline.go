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

	"github.com/pgEdge/postgresql-docs/builder/sgml"
)

// handleEmphasis converts <emphasis> to italic or bold Markdown.
func handleEmphasis(ctx *Context, node *sgml.Node, w *MarkdownWriter) error {
	role := node.GetAttr("role")
	marker := "*"
	if role == "bold" || role == "strong" {
		marker = "**"
	}
	w.WriteString(marker)
	if err := convertChildren(ctx, node, w); err != nil {
		return err
	}
	w.WriteString(marker)
	return nil
}

// handleCode converts inline code elements to backtick spans.
// When the element contains <replaceable> children, uses HTML <code>
// tags instead so that italic formatting renders correctly inside.
func handleCode(ctx *Context, node *sgml.Node, w *MarkdownWriter) error {
	text := node.TextContent()
	// Avoid empty code spans
	if strings.TrimSpace(text) == "" {
		return convertChildren(ctx, node, w)
	}

	// If the code element contains <replaceable>, use HTML tags so
	// that italic formatting renders inside the code span.
	// Outputs: <code>text </code><em>replaceable</em><code> text</code>
	if hasDescendant(node, "replaceable") {
		return convertCodeWithReplaceable(ctx, node, w)
	}

	// If text contains backticks, use double backticks
	if strings.Contains(text, "`") {
		w.WriteString("`` ")
		if err := convertChildren(ctx, node, w); err != nil {
			return err
		}
		w.WriteString(" ``")
	} else {
		w.WriteString("`")
		if err := convertChildren(ctx, node, w); err != nil {
			return err
		}
		w.WriteString("`")
	}
	return nil
}

// hasDescendant checks if a node has any descendant with the given tag.
func hasDescendant(node *sgml.Node, tag string) bool {
	for _, child := range node.Children {
		if child.Type == sgml.ElementNode {
			if child.Tag == tag {
				return true
			}
			if hasDescendant(child, tag) {
				return true
			}
		}
	}
	return false
}

// handleReplaceable converts <replaceable> to italic (placeholder style).
// Uses <em> HTML tags instead of * markers so it works inside <code> too.
func handleReplaceable(ctx *Context, node *sgml.Node, w *MarkdownWriter) error {
	w.WriteString("*")
	if err := convertChildren(ctx, node, w); err != nil {
		return err
	}
	w.WriteString("*")
	return nil
}

// convertCodeWithReplaceable renders a code element that contains
// <replaceable> children. Outputs HTML: <code>text </code><em>var</em><code> text</code>
func convertCodeWithReplaceable(ctx *Context, node *sgml.Node, w *MarkdownWriter) error {
	// Build segments: alternate between code and em spans
	var parts []string
	inCode := true
	var current strings.Builder

	for _, child := range node.Children {
		if child.Type == sgml.TextNode {
			if !inCode {
				// Close em, start code
				parts = append(parts, "<em>"+current.String()+"</em>")
				current.Reset()
				inCode = true
			}
			current.WriteString(child.Text)
		} else if child.Type == sgml.ElementNode && child.Tag == "replaceable" {
			if inCode {
				// Close code segment
				codeText := current.String()
				if codeText != "" {
					parts = append(parts, "<code>"+codeText+"</code>")
				}
				current.Reset()
				inCode = false
			}
			replW := NewMarkdownWriter()
			convertChildren(ctx, child, replW)
			current.WriteString(replW.String())
		} else if child.Type == sgml.ElementNode {
			// Other inline elements — render into current segment
			handler := getHandler(child.Tag)
			if handler != nil {
				segW := NewMarkdownWriter()
				handler(ctx, child, segW)
				current.WriteString(segW.String())
			} else {
				current.WriteString(child.TextContent())
			}
		}
	}

	// Flush remaining segment
	if current.Len() > 0 {
		if inCode {
			parts = append(parts, "<code>"+current.String()+"</code>")
		} else {
			parts = append(parts, "<em>"+current.String()+"</em>")
		}
	}

	w.WriteString(strings.Join(parts, ""))
	return nil
}

// handleOptional converts <optional> to bracket notation.
func handleOptional(ctx *Context, node *sgml.Node, w *MarkdownWriter) error {
	w.WriteString("[")
	if err := convertChildren(ctx, node, w); err != nil {
		return err
	}
	w.WriteString("]")
	return nil
}

// handleQuote converts <quote> to curly double quotes.
func handleQuote(ctx *Context, node *sgml.Node, w *MarkdownWriter) error {
	w.WriteString("\u201c")
	if err := convertChildren(ctx, node, w); err != nil {
		return err
	}
	w.WriteString("\u201d")
	return nil
}

// handleSuperscript converts <superscript> to HTML.
func handleSuperscript(ctx *Context, node *sgml.Node, w *MarkdownWriter) error {
	w.WriteString("<sup>")
	if err := convertChildren(ctx, node, w); err != nil {
		return err
	}
	w.WriteString("</sup>")
	return nil
}

// handleSubscript converts <subscript> to HTML.
func handleSubscript(ctx *Context, node *sgml.Node, w *MarkdownWriter) error {
	w.WriteString("<sub>")
	if err := convertChildren(ctx, node, w); err != nil {
		return err
	}
	w.WriteString("</sub>")
	return nil
}

// handleTrademark converts <trademark> by appending the TM symbol.
func handleTrademark(ctx *Context, node *sgml.Node, w *MarkdownWriter) error {
	if err := convertChildren(ctx, node, w); err != nil {
		return err
	}
	class := node.GetAttr("class")
	switch class {
	case "registered":
		w.WriteString("\u00ae")
	case "copyright":
		w.WriteString("\u00a9")
	case "service":
		w.WriteString("\u2120")
	default:
		w.WriteString("\u2122")
	}
	return nil
}

// handleKeycombo converts <keycombo> to key+key format.
func handleKeycombo(ctx *Context, node *sgml.Node, w *MarkdownWriter) error {
	first := true
	for _, child := range node.Children {
		if child.Type == sgml.ElementNode {
			if !first {
				w.WriteString("+")
			}
			if err := convertNode(ctx, child, w); err != nil {
				return err
			}
			first = false
		}
	}
	return nil
}

// handlePassthrough renders children without any wrapping.
func handlePassthrough(ctx *Context, node *sgml.Node, w *MarkdownWriter) error {
	return convertChildren(ctx, node, w)
}

// handleSkip silently skips the element and its children.
func handleSkip(ctx *Context, node *sgml.Node, w *MarkdownWriter) error {
	return nil
}
