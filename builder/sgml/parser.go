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
	"strings"
)

// emptyElements lists SGML elements that have no closing tag in
// the PostgreSQL documentation. These are elements where the
// DocBook DTD specifies EMPTY content model.
var emptyElements = map[string]bool{
	"xref":    true,
	"anchor":  true,
	"graphic": true,
	"colspec": true,
	"sbr":     true,
	"co":      true,
	"area":    true,
	"void":    true,
	"varargs": true,
}

// htmlElements lists HTML tags that may appear as literal content
// in DocBook source (e.g., inside <literal> elements). These are
// NOT DocBook elements and should be treated as text, not parsed
// as elements with children.
var htmlElements = map[string]bool{
	"a": true, "b": true, "i": true, "u": true,
	"em": true, "strong": true, "br": true, "hr": true,
	"p": true, "span": true, "div": true, "img": true,
	"font": true, "center": true, "small": true, "big": true,
	"sub": true, "sup": true, "tt": true, "code": true,
	"pre": true, "h1": true, "h2": true, "h3": true,
	"h4": true, "h5": true, "h6": true,
	// Tags that appear in PostgreSQL doc examples (SQL, XML, etc.)
	// but are not DocBook elements — treat as literal text.
	"select": true, "key": true, "document": true,
	"relation": true, "criteria": true, "anonymous": true,
	"utc-05": true,
	// False positives from file paths (/dev/null, /usr/local)
	// where &lt; expanded to < creating </dev> and </usr> tokens.
	"dev": true, "usr": true,
	// C header files and misc tags in code examples
	"version": true, "simd.h": true, "regex.h": true,
	"pwd.h": true,
	// DocBook GUI elements (used in older PG versions)
	"menuchoice": true, "guimenu": true,
	"guimenuitem": true, "guibutton": true,
	// XML/SQL tag names in ECPG and other examples
	"order": true, "query": true, "fetch": true,
	"result": true, "offset": true, "describe": true,
}

// Parser builds a document tree from a token stream.
type Parser struct {
	tokens   []Token
	pos      int
	warnings []string
}

// NewParser creates a parser from a slice of tokens.
func NewParser(tokens []Token) *Parser {
	return &Parser{
		tokens: tokens,
		pos:    0,
	}
}

// Parse builds and returns the root node of the document tree.
func (p *Parser) Parse() (*Node, error) {
	root := &Node{
		Type: ElementNode,
		Tag:  "__root__",
	}

	err := p.parseChildren(root)
	if err != nil {
		return root, err
	}

	return root, nil
}

// Warnings returns any non-fatal warnings accumulated during parsing.
func (p *Parser) Warnings() []string {
	return p.warnings
}

// parseChildren parses child nodes and appends them to parent
// until a closing tag for parent is found (or EOF for root).
func (p *Parser) parseChildren(parent *Node) error {
	for p.pos < len(p.tokens) {
		tok := p.tokens[p.pos]

		switch tok.Type {
		case TokenEOF:
			return nil

		case TokenText:
			p.pos++
			// Skip pure whitespace text nodes between block elements
			// but preserve them otherwise
			text := tok.Text
			if text != "" {
				child := &Node{
					Type: TextNode,
					Text: text,
					Line: tok.Line,
				}
				parent.AppendChild(child)
			}

		case TokenComment:
			p.pos++
			child := &Node{
				Type: CommentNode,
				Text: tok.Text,
				Line: tok.Line,
			}
			parent.AppendChild(child)

		case TokenPI:
			p.pos++ // skip processing instructions

		case TokenTagClose:
			// Check if this closes our parent
			if tok.Tag == parent.Tag {
				p.pos++
				return nil
			}
			// HTML close tags in DocBook content — convert to text
			if htmlElements[tok.Tag] {
				p.pos++
				child := &Node{
					Type: TextNode,
					Text: "</" + tok.Tag + ">",
					Line: tok.Line,
				}
				parent.AppendChild(child)
				continue
			}
			// Mismatched close tag — could be implicit close of
			// an ancestor. Don't consume it; let the parent handle it.
			if p.isAncestor(parent, tok.Tag) {
				return nil
			}
			// Stray close tag with no matching open — skip and warn
			p.warn(tok.Line, "unexpected closing tag </%s> inside <%s>",
				tok.Tag, parent.Tag)
			p.pos++

		case TokenTagOpen:
			p.pos++

			// HTML tags in DocBook content — convert to text
			if htmlElements[tok.Tag] {
				// Reconstruct the tag as text
				text := "<" + tok.Tag
				for k, v := range cleanAttrs(tok.Attrs) {
					text += fmt.Sprintf(` %s="%s"`, k, v)
				}
				text += ">"
				child := &Node{
					Type: TextNode,
					Text: text,
					Line: tok.Line,
				}
				parent.AppendChild(child)
				continue
			}

			child := &Node{
				Type:  ElementNode,
				Tag:   tok.Tag,
				Attrs: cleanAttrs(tok.Attrs),
				Line:  tok.Line,
			}
			parent.AppendChild(child)

			// Check if this is a self-closing tag or known empty element
			selfClose := false
			if tok.Attrs != nil && tok.Attrs["\x00selfclose"] == "1" {
				selfClose = true
			}

			if selfClose || emptyElements[tok.Tag] {
				// Empty element, no children to parse
				continue
			}

			// Parse children of this element
			if err := p.parseChildren(child); err != nil {
				return err
			}
		}
	}

	return nil
}

// isAncestor checks if any ancestor of the current node has the
// given tag name, which would mean a close tag should bubble up.
func (p *Parser) isAncestor(node *Node, tag string) bool {
	for n := node.Parent; n != nil; n = n.Parent {
		if n.Tag == tag {
			return true
		}
	}
	return false
}

// warn adds a parsing warning.
func (p *Parser) warn(line int, format string, args ...any) {
	msg := fmt.Sprintf("line %d: %s", line, fmt.Sprintf(format, args...))
	p.warnings = append(p.warnings, msg)
}

// cleanAttrs removes internal marker attributes.
func cleanAttrs(attrs map[string]string) map[string]string {
	if attrs == nil {
		return nil
	}
	result := make(map[string]string)
	for k, v := range attrs {
		if !strings.HasPrefix(k, "\x00") {
			result[k] = v
		}
	}
	if len(result) == 0 {
		return nil
	}
	return result
}

// ParseString is a convenience function that tokenizes and parses
// an SGML string in one call.
func ParseString(input string) (*Node, []string, error) {
	tokenizer := NewTokenizer(input)
	tokens := tokenizer.Tokenize()
	parser := NewParser(tokens)
	root, err := parser.Parse()
	return root, parser.Warnings(), err
}
