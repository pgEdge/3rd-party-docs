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

// handleXref converts <xref linkend="..."> to a Markdown link.
func handleXref(ctx *Context, node *sgml.Node, w *MarkdownWriter) error {
	linkend := node.GetAttr("linkend")
	if linkend == "" {
		return nil
	}

	if ctx == nil {
		// During title extraction, just emit placeholder
		w.WriteString("[" + linkend + "]")
		return nil
	}

	link, title, ok := ctx.ResolveLink(linkend)
	if !ok {
		ctx.Warn("unresolved xref linkend=%q", linkend)
		w.WriteString(fmt.Sprintf("[%s](#%s)", linkend, linkend))
		return nil
	}

	if title == "" {
		title = linkend
	}

	w.WriteString(fmt.Sprintf("[%s](%s)", title, link))
	return nil
}

// handleLink converts <link linkend="...">text</link> to a Markdown link.
func handleLink(ctx *Context, node *sgml.Node, w *MarkdownWriter) error {
	linkend := node.GetAttr("linkend")
	if linkend == "" {
		// No linkend — just render children
		return convertChildren(ctx, node, w)
	}

	// Render link text from children
	textW := NewMarkdownWriter()
	textW.SetSuppressNewlines(true)
	if err := convertChildren(ctx, node, textW); err != nil {
		return err
	}
	text := strings.TrimSpace(textW.String())

	if ctx == nil {
		w.WriteString("[" + text + "]")
		return nil
	}

	link, title, ok := ctx.ResolveLink(linkend)
	if !ok {
		ctx.Warn("unresolved link linkend=%q", linkend)
		w.WriteString(fmt.Sprintf("[%s](#%s)", text, linkend))
		return nil
	}

	if text == "" {
		text = title
	}
	if text == "" {
		text = linkend
	}

	w.WriteString(fmt.Sprintf("[%s](%s)", text, link))
	return nil
}

// handleUlink converts <ulink url="...">text</ulink> to a Markdown link.
func handleUlink(ctx *Context, node *sgml.Node, w *MarkdownWriter) error {
	url := node.GetAttr("url")
	if url == "" {
		return convertChildren(ctx, node, w)
	}

	textW := NewMarkdownWriter()
	textW.SetSuppressNewlines(true)
	if err := convertChildren(ctx, node, textW); err != nil {
		return err
	}
	text := strings.TrimSpace(textW.String())

	if text == "" {
		text = url
	}

	w.WriteString(fmt.Sprintf("[%s](%s)", text, url))
	return nil
}

// handleAnchor converts <anchor id="..."> to an inline HTML anchor.
// The anchor is emitted by convertNode, so this is a no-op.
func handleAnchor(ctx *Context, node *sgml.Node, w *MarkdownWriter) error {
	return nil
}
