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
	"os"
	"path/filepath"
	"strings"

	"github.com/pgEdge/postgresql-docs/builder/sgml"
)

// handleBook converts the root <book> element.
func handleBook(ctx *Context, node *sgml.Node, w *MarkdownWriter) error {
	return convertChildren(ctx, node, w)
}

// handlePart converts <part> elements (major doc divisions).
func handlePart(ctx *Context, node *sgml.Node, w *MarkdownWriter) error {
	return convertChildren(ctx, node, w)
}

// handlePartIntro converts <partintro> elements.
func handlePartIntro(ctx *Context, node *sgml.Node, w *MarkdownWriter) error {
	return convertChildren(ctx, node, w)
}

// handleReference converts <reference> elements (groups of refentries).
func handleReference(ctx *Context, node *sgml.Node, w *MarkdownWriter) error {
	return convertChildren(ctx, node, w)
}

// handleChapter converts <chapter> or <appendix> elements.
func handleChapter(ctx *Context, node *sgml.Node, w *MarkdownWriter) error {
	id := node.GetAttr("id")
	title := extractTitle(node)

	w.Heading(1, title, id)

	return convertChildrenSkipTitle(ctx, node, w)
}

// handleSection converts <sect1> through <sect5> and <section>.
func handleSection(ctx *Context, node *sgml.Node, w *MarkdownWriter) error {
	id := node.GetAttr("id")
	title := extractTitle(node)

	// Determine heading level from tag name
	level := sectionLevel(node.Tag)

	w.Heading(level, title, id)

	return convertChildrenSkipTitle(ctx, node, w)
}

// sectionLevel returns the Markdown heading level for a section tag.
func sectionLevel(tag string) int {
	switch tag {
	case "chapter", "appendix", "preface", "bibliography":
		return 1
	case "sect1", "section", "simplesect":
		return 2
	case "sect2", "refsect2":
		return 3
	case "sect3", "refsect3":
		return 4
	case "sect4":
		return 5
	case "sect5":
		return 6
	default:
		return 2
	}
}

// handlePara converts <para> and <simpara> elements.
func handlePara(ctx *Context, node *sgml.Node, w *MarkdownWriter) error {
	w.BlankLine()
	if err := convertChildren(ctx, node, w); err != nil {
		return err
	}
	w.EnsureNewline()
	return nil
}

// handleFormalPara converts <formalpara> (para with title).
func handleFormalPara(ctx *Context, node *sgml.Node, w *MarkdownWriter) error {
	title := extractTitle(node)
	w.BlankLine()
	if title != "" {
		w.WriteString("**" + title + ".**\n")
	}
	return convertChildrenSkipTitle(ctx, node, w)
}

// handleItemizedList converts <itemizedlist> to Markdown unordered list.
func handleItemizedList(ctx *Context, node *sgml.Node, w *MarkdownWriter) error {
	w.BlankLine()
	title := extractTitle(node)
	if title != "" {
		w.WriteString("**" + title + "**\n")
		w.BlankLine()
	}
	for _, child := range node.Children {
		if child.Type == sgml.ElementNode && child.Tag == "listitem" {
			w.EnsureNewline()
			w.WriteString("- ")
			w.PushIndent("  ")
			if err := convertListItemChildren(ctx, child, w); err != nil {
				return err
			}
			w.PopIndent("  ")
		}
	}
	w.EnsureNewline()
	return nil
}

// handleOrderedList converts <orderedlist> to Markdown ordered list.
func handleOrderedList(ctx *Context, node *sgml.Node, w *MarkdownWriter) error {
	w.BlankLine()
	title := extractTitle(node)
	if title != "" {
		w.WriteString("**" + title + "**\n")
		w.BlankLine()
	}
	num := 1
	for _, child := range node.Children {
		if child.Type == sgml.ElementNode && child.Tag == "listitem" {
			w.EnsureNewline()
			prefix := fmt.Sprintf("%d. ", num)
			w.WriteString(prefix)
			indent := strings.Repeat(" ", len(prefix))
			w.PushIndent(indent)
			if err := convertListItemChildren(ctx, child, w); err != nil {
				return err
			}
			w.PopIndent(indent)
			num++
		}
	}
	w.EnsureNewline()
	return nil
}

// handleVariableList converts <variablelist> to definition-style list.
func handleVariableList(ctx *Context, node *sgml.Node, w *MarkdownWriter) error {
	w.BlankLine()
	title := extractTitle(node)
	if title != "" {
		w.WriteString("**" + title + "**\n")
		w.BlankLine()
	}
	for _, child := range node.Children {
		if child.Type != sgml.ElementNode || child.Tag != "varlistentry" {
			continue
		}
		id := child.GetAttr("id")
		if id != "" {
			w.WriteString(fmt.Sprintf("<a id=\"%s\"></a>\n", id))
		}

		// Render terms
		terms := child.FindChildren("term")
		w.BlankLine()
		for i, term := range terms {
			// Emit anchor for term IDs (these don't go through convertNode)
			if termID := term.GetAttr("id"); termID != "" {
				w.WriteString(fmt.Sprintf("<a id=\"%s\"></a>\n", termID))
			}
			tw := NewMarkdownWriter()
			tw.SetSuppressNewlines(true)
			if err := convertChildren(ctx, term, tw); err != nil {
				return err
			}
			if i > 0 {
				w.WriteString(", ")
			}
			w.WriteString(strings.TrimSpace(tw.String()))
		}
		w.Newline()

		// Render listitem content indented under the term
		listitem := child.FindChild("listitem")
		if listitem != nil {
			w.PushIndent("    ")
			// Use a sub-writer to capture content and prefix with ':'
			subW := NewMarkdownWriter()
			if err := convertChildren(ctx, listitem, subW); err != nil {
				return err
			}
			content := strings.TrimSpace(subW.String())
			if content != "" {
				w.WriteString(":   ")
				// Write first line, then indent the rest
				lines := strings.Split(content, "\n")
				for i, line := range lines {
					if i == 0 {
						w.WriteString(line + "\n")
					} else {
						if line == "" {
							w.WriteString("\n")
						} else {
							w.WriteString("    " + line + "\n")
						}
					}
				}
			}
			w.PopIndent("    ")
		}
	}
	w.EnsureNewline()
	return nil
}

// handleSimpleList converts <simplelist> elements.
func handleSimpleList(ctx *Context, node *sgml.Node, w *MarkdownWriter) error {
	listType := node.GetAttr("type")
	if listType == "inline" {
		// Inline comma-separated list
		members := node.FindChildren("member")
		for i, m := range members {
			if i > 0 {
				w.WriteString(", ")
			}
			if err := convertChildren(ctx, m, w); err != nil {
				return err
			}
		}
		return nil
	}

	// Default: vertical list
	w.BlankLine()
	for _, m := range node.FindChildren("member") {
		w.WriteString("- ")
		if err := convertChildren(ctx, m, w); err != nil {
			return err
		}
		w.Newline()
	}
	return nil
}

// handleAdmonition converts <note>, <tip>, <warning>, <caution>,
// <important> to MkDocs Material admonitions.
func handleAdmonition(ctx *Context, node *sgml.Node, w *MarkdownWriter) error {
	kind := node.Tag
	title := extractTitle(node)

	w.BlankLine()
	if title != "" {
		w.WriteString(fmt.Sprintf("!!! %s \"%s\"\n\n", kind, title))
	} else {
		w.WriteString(fmt.Sprintf("!!! %s\n\n", kind))
	}

	// Render content with 4-space indentation
	subW := NewMarkdownWriter()
	if err := convertChildrenSkipTitle(ctx, node, subW); err != nil {
		return err
	}

	content := strings.TrimSpace(subW.String())
	for _, line := range strings.Split(content, "\n") {
		if line == "" {
			w.WriteString("\n")
		} else {
			w.WriteString("    " + line + "\n")
		}
	}
	w.EnsureNewline()
	return nil
}

// handleProgramListing converts <programlisting> to a fenced code block.
func handleProgramListing(ctx *Context, node *sgml.Node, w *MarkdownWriter) error {
	lang := node.GetAttr("language")
	if lang == "" {
		lang = detectLanguage(node)
	}
	w.StartCodeBlock(lang)
	// Write raw text content preserving whitespace
	w.WriteString(extractRawText(node))
	w.EndCodeBlock()
	return nil
}

// handleScreen converts <screen> to a fenced code block.
func handleScreen(ctx *Context, node *sgml.Node, w *MarkdownWriter) error {
	w.StartCodeBlock("")
	w.WriteString(extractRawText(node))
	w.EndCodeBlock()
	return nil
}

// handleLiteralLayout converts <literallayout> to preformatted text.
func handleLiteralLayout(ctx *Context, node *sgml.Node, w *MarkdownWriter) error {
	class := node.GetAttr("class")
	if class == "monospaced" {
		w.StartCodeBlock("")
		w.WriteString(extractRawText(node))
		w.EndCodeBlock()
	} else {
		// Normal literallayout — preserve lines but allow inline formatting
		w.BlankLine()
		w.WriteString("<pre>\n")
		w.WriteString(extractRawText(node))
		w.WriteString("\n</pre>\n")
	}
	return nil
}

// handleSynopsis converts <synopsis> to a code block.
func handleSynopsis(ctx *Context, node *sgml.Node, w *MarkdownWriter) error {
	w.StartCodeBlock("")
	w.WriteString(extractRawText(node))
	w.EndCodeBlock()
	return nil
}

// handleBlockquote converts to Markdown blockquote.
func handleBlockquote(ctx *Context, node *sgml.Node, w *MarkdownWriter) error {
	w.BlankLine()

	subW := NewMarkdownWriter()
	if err := convertChildren(ctx, node, subW); err != nil {
		return err
	}

	content := strings.TrimSpace(subW.String())
	for _, line := range strings.Split(content, "\n") {
		w.WriteString("> " + line + "\n")
	}
	w.EnsureNewline()
	return nil
}

// handleExample converts <example> with title.
func handleExample(ctx *Context, node *sgml.Node, w *MarkdownWriter) error {
	title := extractTitle(node)

	if title != "" {
		w.BlankLine()
		w.WriteString("**Example: " + title + "**\n")
	}

	return convertChildrenSkipTitle(ctx, node, w)
}

// handleFigure converts <figure> elements.
func handleFigure(ctx *Context, node *sgml.Node, w *MarkdownWriter) error {
	title := extractTitle(node)
	if title != "" {
		w.BlankLine()
		w.WriteString("**" + title + "**\n")
	}
	return convertChildrenSkipTitle(ctx, node, w)
}

// handleImagedata converts <imagedata> to Markdown image and copies
// the image file from the SGML source to the output directory.
func handleImagedata(ctx *Context, node *sgml.Node, w *MarkdownWriter) error {
	fileref := node.GetAttr("fileref")
	if fileref == "" {
		return nil
	}

	w.BlankLine()
	w.WriteString(fmt.Sprintf("![image](%s)\n", fileref))

	// Copy image file to output directory alongside the current .md file
	if ctx != nil && ctx.SrcDir != "" && ctx.OutDir != "" && ctx.CurrentFile != "" {
		srcPath := filepath.Join(ctx.SrcDir, fileref)
		outDir := filepath.Dir(filepath.Join(ctx.OutDir, ctx.CurrentFile))
		dstPath := filepath.Join(outDir, fileref)

		// Create destination directory
		if err := os.MkdirAll(filepath.Dir(dstPath), 0755); err != nil {
			ctx.Warn("could not create image dir: %v", err)
			return nil
		}

		// Copy the file
		data, err := os.ReadFile(srcPath)
		if err != nil {
			ctx.Warn("could not read image %s: %v", srcPath, err)
			return nil
		}
		if err := os.WriteFile(dstPath, data, 0644); err != nil {
			ctx.Warn("could not write image %s: %v", dstPath, err)
		}
	}

	return nil
}

// handleFootnote converts <footnote> to a parenthetical.
func handleFootnote(ctx *Context, node *sgml.Node, w *MarkdownWriter) error {
	subW := NewMarkdownWriter()
	subW.SetSuppressNewlines(true)
	if err := convertChildren(ctx, node, subW); err != nil {
		return err
	}
	content := strings.TrimSpace(subW.String())
	w.WriteString(" (" + content + ")")
	return nil
}

// handleProcedure converts <procedure> to an ordered list.
func handleProcedure(ctx *Context, node *sgml.Node, w *MarkdownWriter) error {
	title := extractTitle(node)
	w.BlankLine()
	if title != "" {
		w.WriteString("**" + title + "**\n")
		w.BlankLine()
	}
	num := 1
	for _, child := range node.Children {
		if child.Type == sgml.ElementNode && child.Tag == "step" {
			// Emit anchor for step IDs (steps bypass convertNode)
			if stepID := child.GetAttr("id"); stepID != "" {
				w.WriteString(fmt.Sprintf("<a id=\"%s\"></a>\n", stepID))
			}
			w.EnsureNewline()
			prefix := fmt.Sprintf("%d. ", num)
			w.WriteString(prefix)
			indent := strings.Repeat(" ", len(prefix))
			w.PushIndent(indent)
			if err := convertListItemChildren(ctx, child, w); err != nil {
				return err
			}
			w.PopIndent(indent)
			num++
		}
	}
	w.EnsureNewline()
	return nil
}

// handleStep converts a <step> element (rendered as ordered list item
// by the parent procedure handler, but if encountered standalone).
func handleStep(ctx *Context, node *sgml.Node, w *MarkdownWriter) error {
	return convertChildren(ctx, node, w)
}

// handleSubsteps converts <substeps> to a nested ordered list.
func handleSubsteps(ctx *Context, node *sgml.Node, w *MarkdownWriter) error {
	w.BlankLine()
	num := 1
	for _, child := range node.Children {
		if child.Type == sgml.ElementNode && child.Tag == "step" {
			if stepID := child.GetAttr("id"); stepID != "" {
				w.WriteString(fmt.Sprintf("<a id=\"%s\"></a>\n", stepID))
			}
			w.EnsureNewline()
			prefix := fmt.Sprintf("%d. ", num)
			w.WriteString(prefix)
			indent := strings.Repeat(" ", len(prefix))
			w.PushIndent(indent)
			if err := convertListItemChildren(ctx, child, w); err != nil {
				return err
			}
			w.PopIndent(indent)
			num++
		}
	}
	w.EnsureNewline()
	return nil
}

// handleEmail converts <email> to a mailto link.
func handleEmail(ctx *Context, node *sgml.Node, w *MarkdownWriter) error {
	addr := node.TextContent()
	w.WriteString(fmt.Sprintf("[%s](mailto:%s)", addr, addr))
	return nil
}

// handleGlossEntry converts <glossentry> to a definition term.
func handleGlossEntry(ctx *Context, node *sgml.Node, w *MarkdownWriter) error {
	// Render glossterm as bold
	glossterm := node.FindChild("glossterm")
	if glossterm != nil {
		w.BlankLine()
		tw := NewMarkdownWriter()
		tw.SetSuppressNewlines(true)
		if err := convertChildren(ctx, glossterm, tw); err != nil {
			return err
		}
		w.WriteString("**" + strings.TrimSpace(tw.String()) + "**\n")
	}

	// Render glossdef content
	glossdef := node.FindChild("glossdef")
	if glossdef != nil {
		w.PushIndent("    ")
		subW := NewMarkdownWriter()
		if err := convertChildren(ctx, glossdef, subW); err != nil {
			return err
		}
		content := strings.TrimSpace(subW.String())
		if content != "" {
			w.WriteString(":   ")
			lines := strings.Split(content, "\n")
			for i, line := range lines {
				if i == 0 {
					w.WriteString(line + "\n")
				} else if line == "" {
					w.WriteString("\n")
				} else {
					w.WriteString("    " + line + "\n")
				}
			}
		}
		w.PopIndent("    ")
	}

	// Render glosssee/glossseealso
	for _, child := range node.Children {
		if child.Type == sgml.ElementNode {
			if child.Tag == "glosssee" || child.Tag == "glossseealso" {
				if err := convertNode(ctx, child, w); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

// handleGlossSee converts <glosssee> to "See: ..."
func handleGlossSee(ctx *Context, node *sgml.Node, w *MarkdownWriter) error {
	w.WriteString("    *See: ")
	if err := convertChildren(ctx, node, w); err != nil {
		return err
	}
	w.WriteString("*\n")
	return nil
}

// handleGlossSeeAlso converts <glossseealso> to "See also: ..."
func handleGlossSeeAlso(ctx *Context, node *sgml.Node, w *MarkdownWriter) error {
	w.WriteString("    *See also: ")
	if err := convertChildren(ctx, node, w); err != nil {
		return err
	}
	w.WriteString("*\n")
	return nil
}

// handleCiteRefEntry converts <citerefentry> to formatted reference.
func handleCiteRefEntry(ctx *Context, node *sgml.Node, w *MarkdownWriter) error {
	titleNode := node.FindChild("refentrytitle")
	volNode := node.FindChild("manvolnum")
	title := ""
	vol := ""
	if titleNode != nil {
		title = titleNode.TextContent()
	}
	if volNode != nil {
		vol = volNode.TextContent()
	}
	if vol != "" {
		w.WriteString(fmt.Sprintf("`%s`(%s)", title, vol))
	} else {
		w.WriteString(fmt.Sprintf("`%s`", title))
	}
	return nil
}

// handleManvolnum converts <manvolnum> (handled by parent citerefentry).
func handleManvolnum(ctx *Context, node *sgml.Node, w *MarkdownWriter) error {
	w.WriteString("(" + node.TextContent() + ")")
	return nil
}

// handleFootnoteRef converts <footnoteref> to a reference to a footnote.
func handleFootnoteRef(ctx *Context, node *sgml.Node, w *MarkdownWriter) error {
	linkend := node.GetAttr("linkend")
	if linkend != "" {
		w.WriteString(fmt.Sprintf("[^%s]", linkend))
	}
	return nil
}

// convertListItemChildren converts the children of a <listitem>,
// handling the first <para> inline (no blank line before it).
func convertListItemChildren(ctx *Context, node *sgml.Node, w *MarkdownWriter) error {
	first := true
	for _, child := range node.Children {
		if child.Type == sgml.TextNode {
			text := strings.TrimSpace(child.Text)
			if text != "" {
				w.Write(text)
			}
			continue
		}
		if child.Type == sgml.ElementNode {
			if first && child.Tag == "para" {
				// First para in a list item — render inline
				first = false
				if err := convertChildren(ctx, child, w); err != nil {
					return err
				}
				w.Newline()
				continue
			}
			first = false
			if err := convertNode(ctx, child, w); err != nil {
				return err
			}
		}
	}
	return nil
}

// extractTitle returns the text content of the first <title> child.
func extractTitle(node *sgml.Node) string {
	title := node.FindChild("title")
	if title == nil {
		return ""
	}
	// Render title content to get inline formatting
	w := NewMarkdownWriter()
	w.SetSuppressNewlines(true)
	for _, child := range title.Children {
		if child.Type == sgml.TextNode {
			w.WriteString(child.Text)
		} else if child.Type == sgml.ElementNode {
			// For title elements, render inline elements
			handler := getHandler(child.Tag)
			if handler != nil {
				handler(nil, child, w)
			} else {
				w.WriteString(child.TextContent())
			}
		}
	}
	// Normalize whitespace — titles from SGML may contain newlines
	result := strings.TrimSpace(w.String())
	return normalizeWhitespace(result)
}

// textContentSkipping returns the text content of a node, skipping
// any descendant elements with the given tag name.
func textContentSkipping(node *sgml.Node, skipTag string) string {
	var b strings.Builder
	var walk func(*sgml.Node)
	walk = func(n *sgml.Node) {
		if n.Type == sgml.TextNode {
			b.WriteString(n.Text)
			return
		}
		if n.Type == sgml.ElementNode && n.Tag == skipTag {
			return
		}
		for _, child := range n.Children {
			walk(child)
		}
	}
	walk(node)
	return normalizeWhitespace(strings.TrimSpace(b.String()))
}

// convertChildrenSkipTitle converts all children except <title>.
func convertChildrenSkipTitle(ctx *Context, node *sgml.Node, w *MarkdownWriter) error {
	for _, child := range node.Children {
		if child.Type == sgml.ElementNode &&
			(child.Tag == "title" || child.Tag == "titleabbrev") {
			continue
		}
		if err := convertNode(ctx, child, w); err != nil {
			return err
		}
	}
	return nil
}

// extractRawText gets the raw text content of an element,
// stripping tags but preserving whitespace.
func extractRawText(node *sgml.Node) string {
	var b strings.Builder
	var walk func(*sgml.Node)
	walk = func(n *sgml.Node) {
		switch n.Type {
		case sgml.TextNode:
			b.WriteString(n.Text)
		case sgml.ElementNode:
			// For replaceable in code blocks, use uppercase convention
			if n.Tag == "replaceable" {
				text := strings.ToUpper(n.TextContent())
				b.WriteString(text)
				return
			}
			// For optional, wrap in brackets
			if n.Tag == "optional" {
				b.WriteString("[")
				for _, c := range n.Children {
					walk(c)
				}
				b.WriteString("]")
				return
			}
			for _, c := range n.Children {
				walk(c)
			}
		}
	}
	walk(node)
	return strings.TrimRight(b.String(), "\n")
}

// detectLanguage guesses a code block language from content.
func detectLanguage(node *sgml.Node) string {
	text := node.TextContent()
	text = strings.TrimSpace(text)

	if strings.HasPrefix(text, "SELECT") ||
		strings.HasPrefix(text, "INSERT") ||
		strings.HasPrefix(text, "UPDATE") ||
		strings.HasPrefix(text, "DELETE") ||
		strings.HasPrefix(text, "CREATE") ||
		strings.HasPrefix(text, "ALTER") ||
		strings.HasPrefix(text, "DROP") ||
		strings.HasPrefix(text, "GRANT") ||
		strings.HasPrefix(text, "REVOKE") ||
		strings.HasPrefix(text, "BEGIN") ||
		strings.HasPrefix(text, "COMMIT") ||
		strings.HasPrefix(text, "EXPLAIN") ||
		strings.HasPrefix(text, "WITH") ||
		strings.HasPrefix(text, "SET") {
		return "sql"
	}

	return ""
}

// handleBibliodiv converts <bibliodiv> (a section within a bibliography).
func handleBibliodiv(ctx *Context, node *sgml.Node, w *MarkdownWriter) error {
	title := extractTitle(node)
	if title != "" {
		w.Heading(2, title, "")
	}
	return convertChildrenSkipTitle(ctx, node, w)
}

// handleBiblioentry converts <biblioentry> to a formatted reference.
func handleBiblioentry(ctx *Context, node *sgml.Node, w *MarkdownWriter) error {
	w.BlankLine()

	title := extractTitle(node)
	subtitle := ""
	if st := node.FindChild("subtitle"); st != nil {
		subtitle = st.TextContent()
	}

	// Collect authors
	var authors []string
	for _, ag := range node.FindDescendants("author") {
		first := ""
		last := ""
		if fn := ag.FindChild("firstname"); fn != nil {
			first = fn.TextContent()
		}
		if sn := ag.FindChild("surname"); sn != nil {
			last = sn.TextContent()
		}
		if first != "" || last != "" {
			authors = append(authors, strings.TrimSpace(first+" "+last))
		}
	}

	pubdate := ""
	if pd := node.FindChild("pubdate"); pd != nil {
		pubdate = pd.TextContent()
	}

	isbn := ""
	if ib := node.FindChild("isbn"); ib != nil {
		isbn = ib.TextContent()
	}

	// Format: **Title**. *Subtitle*. Authors. Year. ISBN.
	var parts []string
	if title != "" {
		parts = append(parts, "**"+title+"**")
	}
	if subtitle != "" {
		parts = append(parts, "*"+subtitle+"*")
	}
	if len(authors) > 0 {
		parts = append(parts, strings.Join(authors, ", "))
	}
	if pubdate != "" {
		parts = append(parts, pubdate)
	}
	if isbn != "" {
		parts = append(parts, "ISBN "+isbn)
	}

	w.WriteString(strings.Join(parts, ". ") + ".\n")
	return nil
}
