//-------------------------------------------------------------------------
//
// pgEdge PostgreSQL Docs
//
// Copyright (c) 2026, pgEdge, Inc.
// This software is released under The PostgreSQL License
//
//-------------------------------------------------------------------------

// Package sgml provides SGML parsing for PostgreSQL DocBook documentation.
package sgml

import (
	"strings"
)

// NodeType identifies the kind of node in the document tree.
type NodeType int

const (
	ElementNode NodeType = iota
	TextNode
	CommentNode
)

// Node represents a single node in the parsed SGML document tree.
type Node struct {
	Type     NodeType
	Tag      string
	Attrs    map[string]string
	Children []*Node
	Text     string
	Parent   *Node
	Line     int
}

// GetAttr returns the value of an attribute, or empty string if not present.
func (n *Node) GetAttr(name string) string {
	if n.Attrs == nil {
		return ""
	}
	return n.Attrs[name]
}

// FindChildren returns all direct child elements with the given tag name.
func (n *Node) FindChildren(tag string) []*Node {
	var result []*Node
	for _, c := range n.Children {
		if c.Type == ElementNode && c.Tag == tag {
			result = append(result, c)
		}
	}
	return result
}

// FindChild returns the first direct child element with the given tag name.
func (n *Node) FindChild(tag string) *Node {
	for _, c := range n.Children {
		if c.Type == ElementNode && c.Tag == tag {
			return c
		}
	}
	return nil
}

// FindDescendants returns all descendant elements with the given tag name.
func (n *Node) FindDescendants(tag string) []*Node {
	var result []*Node
	var walk func(*Node)
	walk = func(node *Node) {
		for _, c := range node.Children {
			if c.Type == ElementNode && c.Tag == tag {
				result = append(result, c)
			}
			walk(c)
		}
	}
	walk(n)
	return result
}

// TextContent returns the concatenated text content of this node
// and all descendants.
func (n *Node) TextContent() string {
	if n.Type == TextNode {
		return n.Text
	}
	var b strings.Builder
	for _, c := range n.Children {
		b.WriteString(c.TextContent())
	}
	return b.String()
}

// AppendChild adds a child node and sets its parent.
func (n *Node) AppendChild(child *Node) {
	child.Parent = n
	n.Children = append(n.Children, child)
}

// RemoveSections removes sections/chapters whose title matches
// any of the given headings from the document tree.
func RemoveSections(root *Node, headings []string) {
	skip := make(map[string]bool)
	for _, h := range headings {
		skip[strings.TrimSpace(strings.ToLower(h))] = true
	}

	var walk func(n *Node)
	walk = func(n *Node) {
		var kept []*Node
		for _, child := range n.Children {
			if child.Type == ElementNode {
				// Check if this is a section-like element
				switch child.Tag {
				case "chapter", "section", "appendix",
					"sect1", "refsect1", "simplesect":
					title := child.FindChild("title")
					if title != nil {
						text := strings.TrimSpace(
							strings.ToLower(
								title.TextContent()))
						if skip[text] {
							continue // remove
						}
					}
				}
			}
			walk(child)
			kept = append(kept, child)
		}
		n.Children = kept
	}
	walk(root)
}
