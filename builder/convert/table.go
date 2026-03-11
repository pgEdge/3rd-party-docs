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

// handleTable converts <table> and <informaltable> to Markdown or HTML.
func handleTable(ctx *Context, node *sgml.Node, w *MarkdownWriter) error {
	title := extractTitle(node)

	w.BlankLine()

	if title != "" {
		w.WriteString("**Table: " + title + "**\n")
		w.BlankLine()
	}

	tgroup := node.FindChild("tgroup")
	if tgroup == nil {
		// Try direct thead/tbody (HTML-style table in DocBook)
		return convertHTMLTable(ctx, node, w)
	}

	// Check for func_table_entry pattern (single-column tables where
	// each entry has multiple <para> elements that should be columns).
	if isFuncTableEntry(tgroup) {
		return convertFuncTable(ctx, tgroup, w)
	}

	// Check if table needs HTML fallback (spans, complex content)
	if tableNeedsHTML(tgroup) {
		return convertHTMLTable(ctx, node, w)
	}

	return convertMarkdownTable(ctx, tgroup, w)
}

// isFuncTableEntry checks if a table uses the func_table_entry pattern:
// single-column tables where each entry has role="func_table_entry" and
// multiple <para> elements that should be rendered as separate columns.
func isFuncTableEntry(tgroup *sgml.Node) bool {
	entries := tgroup.FindDescendants("entry")
	if len(entries) == 0 {
		return false
	}
	for _, entry := range entries {
		if entry.GetAttr("role") == "func_table_entry" {
			return true
		}
	}
	return false
}

// convertFuncTable renders a func_table_entry table as a multi-column
// HTML table. Each entry's <para> children become separate cells.
func convertFuncTable(ctx *Context, tgroup *sgml.Node, w *MarkdownWriter) error {
	thead := tgroup.FindChild("thead")
	tbody := tgroup.FindChild("tbody")

	w.WriteString("<table>\n")

	// Render header — split paras into columns
	if thead != nil {
		w.WriteString("<thead>\n")
		for _, row := range thead.FindChildren("row") {
			w.WriteString("<tr>\n")
			for _, entry := range row.FindChildren("entry") {
				paras := entry.FindChildren("para")
				for _, para := range paras {
					cellW := NewMarkdownWriter()
					convertChildren(ctx, para, cellW)
					content := markdownToHTML(cellW.String())
					w.WriteString("<th>" + content + "</th>\n")
				}
			}
			w.WriteString("</tr>\n")
		}
		w.WriteString("</thead>\n")
	}

	// Render body — split paras into columns
	if tbody != nil {
		w.WriteString("<tbody>\n")
		for _, row := range tbody.FindChildren("row") {
			// Check for entry IDs to add to the row
			rowAttrs := ""
			for _, entry := range row.FindChildren("entry") {
				if entryID := entry.GetAttr("id"); entryID != "" {
					rowAttrs = fmt.Sprintf(` id="%s"`, entryID)
					break
				}
			}
			w.WriteString(fmt.Sprintf("<tr%s>\n", rowAttrs))
			for _, entry := range row.FindChildren("entry") {
				renderFuncTableEntry(ctx, entry, w)
			}
			w.WriteString("</tr>\n")
		}
		w.WriteString("</tbody>\n")
	}

	w.WriteString("</table>\n")
	return nil
}

// renderFuncTableEntry renders a func_table_entry as multiple <td> cells.
// The first <para role="func_signature"> becomes the signature column.
// The second <para> becomes the description column.
// Remaining <para> elements (examples) are joined in the third column.
func renderFuncTableEntry(ctx *Context, entry *sgml.Node, w *MarkdownWriter) {
	paras := entry.FindChildren("para")
	if len(paras) == 0 {
		w.WriteString("<td></td>\n")
		return
	}

	// Column 1: Signature (first para, typically role="func_signature")
	sigW := NewMarkdownWriter()
	convertChildren(ctx, paras[0], sigW)
	w.WriteString("<td>" + markdownToHTML(sigW.String()) + "</td>\n")

	// Column 2: Description (second para if it exists)
	if len(paras) > 1 {
		descW := NewMarkdownWriter()
		convertChildren(ctx, paras[1], descW)
		w.WriteString("<td>" + markdownToHTML(descW.String()) + "</td>\n")
	} else {
		w.WriteString("<td></td>\n")
	}

	// Column 3: Examples (remaining paras joined)
	if len(paras) > 2 {
		var exParts []string
		for _, p := range paras[2:] {
			exW := NewMarkdownWriter()
			convertChildren(ctx, p, exW)
			html := markdownToHTML(exW.String())
			if html != "" {
				exParts = append(exParts, html)
			}
		}
		w.WriteString("<td>" + strings.Join(exParts, "<br>") +
			"</td>\n")
	} else {
		w.WriteString("<td></td>\n")
	}
}

// tableNeedsHTML checks if a table has features Markdown can't handle.
func tableNeedsHTML(tgroup *sgml.Node) bool {
	entries := tgroup.FindDescendants("entry")
	for _, entry := range entries {
		// Check for spanning
		if entry.GetAttr("morerows") != "" {
			return true
		}
		if entry.GetAttr("namest") != "" || entry.GetAttr("nameend") != "" {
			return true
		}
		if entry.GetAttr("spanname") != "" {
			return true
		}

		// Check for complex content (multiple paragraphs, lists, etc.)
		paras := entry.FindChildren("para")
		if len(paras) > 1 {
			return true
		}
		if len(entry.FindChildren("itemizedlist")) > 0 ||
			len(entry.FindChildren("orderedlist")) > 0 ||
			len(entry.FindChildren("programlisting")) > 0 {
			return true
		}
	}
	return false
}

// convertMarkdownTable renders a simple table as a Markdown pipe table.
func convertMarkdownTable(ctx *Context, tgroup *sgml.Node, w *MarkdownWriter) error {
	thead := tgroup.FindChild("thead")
	tbody := tgroup.FindChild("tbody")

	// Render header
	if thead != nil {
		rows := thead.FindChildren("row")
		if len(rows) > 0 {
			if err := renderMarkdownRow(ctx, rows[0], w); err != nil {
				return err
			}
			// Separator
			cells := rows[0].FindChildren("entry")
			w.WriteString("|")
			for range cells {
				w.WriteString(" --- |")
			}
			w.Newline()
		}
	} else {
		// No header — check colspec count for column count
		colspecs := tgroup.FindChildren("colspec")
		if len(colspecs) > 0 {
			w.WriteString("|")
			for range colspecs {
				w.WriteString("  |")
			}
			w.Newline()
			w.WriteString("|")
			for range colspecs {
				w.WriteString(" --- |")
			}
			w.Newline()
		}
	}

	// Render body rows
	if tbody != nil {
		for _, row := range tbody.FindChildren("row") {
			if err := renderMarkdownRow(ctx, row, w); err != nil {
				return err
			}
		}
	}

	w.EnsureNewline()
	return nil
}

// renderMarkdownRow renders a single table row as Markdown.
func renderMarkdownRow(ctx *Context, row *sgml.Node, w *MarkdownWriter) error {
	entries := row.FindChildren("entry")
	w.WriteString("|")
	for _, entry := range entries {
		cellW := NewMarkdownWriter()
		cellW.SetSuppressNewlines(true)
		if err := convertChildren(ctx, entry, cellW); err != nil {
			return err
		}
		text := strings.TrimSpace(cellW.String())
		// Replace newlines with spaces in cell content
		text = strings.ReplaceAll(text, "\n", " ")
		w.WriteString(" " + text + " |")
	}
	w.Newline()
	return nil
}

// convertHTMLTable renders a table as HTML for complex cases.
func convertHTMLTable(ctx *Context, node *sgml.Node, w *MarkdownWriter) error {
	tgroup := node.FindChild("tgroup")
	if tgroup == nil {
		return nil
	}

	// Build colspec name→index map for span resolution
	colspecs := tgroup.FindChildren("colspec")
	colNames := make(map[string]int)
	for i, cs := range colspecs {
		name := cs.GetAttr("colname")
		if name != "" {
			colNames[name] = i
		}
	}

	w.WriteString("<table>\n")

	// Render thead
	thead := tgroup.FindChild("thead")
	if thead != nil {
		w.WriteString("<thead>\n")
		for _, row := range thead.FindChildren("row") {
			renderHTMLRow(ctx, row, colNames, w, "th")
		}
		w.WriteString("</thead>\n")
	}

	// Render tbody
	tbody := tgroup.FindChild("tbody")
	if tbody != nil {
		w.WriteString("<tbody>\n")
		for _, row := range tbody.FindChildren("row") {
			renderHTMLRow(ctx, row, colNames, w, "td")
		}
		w.WriteString("</tbody>\n")
	}

	// Render tfoot
	tfoot := tgroup.FindChild("tfoot")
	if tfoot != nil {
		w.WriteString("<tfoot>\n")
		for _, row := range tfoot.FindChildren("row") {
			renderHTMLRow(ctx, row, colNames, w, "td")
		}
		w.WriteString("</tfoot>\n")
	}

	w.WriteString("</table>\n")
	return nil
}

// renderHTMLRow renders a single table row as HTML.
// Cell content is converted from Markdown to HTML since Python-Markdown
// cannot process Markdown inside <table> elements.
func renderHTMLRow(ctx *Context, row *sgml.Node, colNames map[string]int, w *MarkdownWriter, cellTag string) {
	w.WriteString("<tr>\n")
	for _, entry := range row.FindChildren("entry") {
		attrs := ""

		// Emit anchor for entry IDs (e.g. func_table_entry references)
		if entryID := entry.GetAttr("id"); entryID != "" {
			attrs += fmt.Sprintf(` id="%s"`, entryID)
		}

		// Handle column spanning
		namest := entry.GetAttr("namest")
		nameend := entry.GetAttr("nameend")
		if namest != "" && nameend != "" {
			startCol, ok1 := colNames[namest]
			endCol, ok2 := colNames[nameend]
			if ok1 && ok2 {
				span := endCol - startCol + 1
				if span > 1 {
					attrs += fmt.Sprintf(` colspan="%d"`, span)
				}
			}
		}

		// Handle row spanning
		morerows := entry.GetAttr("morerows")
		if morerows != "" {
			// Validate morerows is a number to prevent attribute injection
			valid := true
			for _, ch := range morerows {
				if ch < '0' || ch > '9' {
					valid = false
					break
				}
			}
			if valid {
				attrs += fmt.Sprintf(` rowspan="%s"`, morerows)
			}
		}

		// Render cell content as Markdown, then convert to HTML.
		cellW := NewMarkdownWriter()
		convertChildren(ctx, entry, cellW)
		content := markdownToHTML(cellW.String())

		w.WriteString(fmt.Sprintf("<%s%s>%s</%s>\n",
			cellTag, attrs, content, cellTag))
	}
	w.WriteString("</tr>\n")
}
