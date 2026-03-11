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
	"fmt"
	"strings"

	"github.com/pgEdge/postgresql-docs/builder/sgml"
)

// handleRefentry converts <refentry> to a consistently structured
// Markdown reference page.
func handleRefentry(ctx *Context, node *sgml.Node, w *MarkdownWriter) error {
	id := node.GetAttr("id")

	// Extract metadata
	refmeta := node.FindChild("refmeta")
	refnamediv := node.FindChild("refnamediv")

	title := ""
	if refmeta != nil {
		titleElem := refmeta.FindChild("refentrytitle")
		if titleElem != nil {
			title = titleElem.TextContent()
		}
	}
	if title == "" && refnamediv != nil {
		nameElem := refnamediv.FindChild("refname")
		if nameElem != nil {
			title = nameElem.TextContent()
		}
	}

	purpose := ""
	if refnamediv != nil {
		purposeElem := refnamediv.FindChild("refpurpose")
		if purposeElem != nil {
			pw := NewMarkdownWriter()
			pw.SetSuppressNewlines(true)
			convertChildren(ctx, purposeElem, pw)
			purpose = strings.TrimSpace(pw.String())
		}
	}

	// Page title
	w.Heading(1, title, id)

	// Purpose line
	if purpose != "" {
		w.BlankLine()
		w.WriteString(purpose + "\n")
	}

	// Synopsis
	synopsisDiv := node.FindChild("refsynopsisdiv")
	if synopsisDiv != nil {
		w.Heading(2, "Synopsis", "")
		if err := convertChildrenSkipTitle(ctx, synopsisDiv, w); err != nil {
			return err
		}
	}

	// Remaining refsect1 sections
	for _, sect := range node.FindChildren("refsect1") {
		sectID := sect.GetAttr("id")
		sectTitle := extractTitle(sect)

		if sectID != "" {
			w.WriteString(fmt.Sprintf("<a id=\"%s\"></a>\n", sectID))
		}
		w.Heading(2, sectTitle, sectID)

		if err := convertChildrenSkipTitle(ctx, sect, w); err != nil {
			return err
		}
	}

	return nil
}

// handleCmdSynopsis converts <cmdsynopsis> to a code-block representation.
func handleCmdSynopsis(ctx *Context, node *sgml.Node, w *MarkdownWriter) error {
	w.StartCodeBlock("")

	for _, child := range node.Children {
		switch child.Type {
		case sgml.TextNode:
			text := strings.TrimSpace(child.Text)
			if text != "" {
				w.WriteString(text)
			}
		case sgml.ElementNode:
			switch child.Tag {
			case "command":
				w.WriteString(child.TextContent())
			case "arg":
				renderArg(child, w)
			case "group":
				renderGroup(child, w)
			case "sbr":
				w.WriteString("\n    ")
			default:
				w.WriteString(extractRawText(child))
			}
		}
	}

	w.Newline()
	w.EndCodeBlock()
	return nil
}

// handleArg converts <arg> elements (used in inline contexts).
func handleArg(ctx *Context, node *sgml.Node, w *MarkdownWriter) error {
	renderArg(node, w)
	return nil
}

// handleGroup converts <group> elements (used in inline contexts).
func handleGroup(ctx *Context, node *sgml.Node, w *MarkdownWriter) error {
	renderGroup(node, w)
	return nil
}

// handleSbr converts <sbr> (synopsis line break).
func handleSbr(ctx *Context, node *sgml.Node, w *MarkdownWriter) error {
	w.WriteString("\n    ")
	return nil
}

// renderArg renders an <arg> element in synopsis context.
func renderArg(node *sgml.Node, w *MarkdownWriter) {
	choice := node.GetAttr("choice")
	rep := node.GetAttr("rep")

	// Opening bracket based on choice
	switch choice {
	case "plain", "req":
		// No brackets for required/plain
	case "opt", "":
		w.WriteString(" [")
	}

	// Render children
	for _, child := range node.Children {
		switch child.Type {
		case sgml.TextNode:
			w.WriteString(child.Text)
		case sgml.ElementNode:
			switch child.Tag {
			case "option":
				w.WriteString(child.TextContent())
			case "replaceable":
				w.WriteString(strings.ToUpper(child.TextContent()))
			case "arg":
				renderArg(child, w)
			case "group":
				renderGroup(child, w)
			default:
				w.WriteString(extractRawText(child))
			}
		}
	}

	// Repeat indicator
	if rep == "repeat" {
		w.WriteString("...")
	}

	// Closing bracket
	switch choice {
	case "plain", "req":
		// No brackets
	case "opt", "":
		w.WriteString("]")
	}
}

// renderGroup renders a <group> element (alternatives) in synopsis context.
func renderGroup(node *sgml.Node, w *MarkdownWriter) {
	choice := node.GetAttr("choice")

	switch choice {
	case "plain", "req":
		w.WriteString(" {")
	default:
		w.WriteString(" [")
	}

	first := true
	for _, child := range node.Children {
		if child.Type == sgml.ElementNode {
			if !first {
				w.WriteString(" | ")
			}
			first = false
			switch child.Tag {
			case "arg":
				renderArgInGroup(child, w)
			case "option":
				w.WriteString(child.TextContent())
			case "replaceable":
				w.WriteString(strings.ToUpper(child.TextContent()))
			default:
				w.WriteString(extractRawText(child))
			}
		}
	}

	switch choice {
	case "plain", "req":
		w.WriteString("}")
	default:
		w.WriteString("]")
	}
}

// renderArgInGroup renders an <arg> inside a <group> (no extra brackets).
func renderArgInGroup(node *sgml.Node, w *MarkdownWriter) {
	for _, child := range node.Children {
		switch child.Type {
		case sgml.TextNode:
			w.WriteString(child.Text)
		case sgml.ElementNode:
			switch child.Tag {
			case "option":
				w.WriteString(child.TextContent())
			case "replaceable":
				w.WriteString(strings.ToUpper(child.TextContent()))
			default:
				w.WriteString(extractRawText(child))
			}
		}
	}
}
