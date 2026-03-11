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
	"github.com/pgEdge/postgresql-docs/builder/sgml"
)

// ElementHandler converts a single SGML element to Markdown.
// It writes to the provided MarkdownWriter and returns any error.
type ElementHandler func(ctx *Context, node *sgml.Node, w *MarkdownWriter) error

// handlerRegistry maps element tag names to their handlers.
var handlerRegistry = map[string]ElementHandler{}

// registerHandler adds a handler to the registry.
func registerHandler(tag string, handler ElementHandler) {
	handlerRegistry[tag] = handler
}

// getHandler returns the handler for a given tag, or nil.
func getHandler(tag string) ElementHandler {
	return handlerRegistry[tag]
}

func init() {
	// Block elements
	registerHandler("para", handlePara)
	registerHandler("simpara", handlePara)
	registerHandler("formalpara", handleFormalPara)
	registerHandler("chapter", handleChapter)
	registerHandler("sect1", handleSection)
	registerHandler("sect2", handleSection)
	registerHandler("sect3", handleSection)
	registerHandler("sect4", handleSection)
	registerHandler("sect5", handleSection)
	registerHandler("section", handleSection)
	registerHandler("simplesect", handleSection)
	registerHandler("appendix", handleChapter)
	registerHandler("preface", handleChapter)
	registerHandler("part", handlePart)
	registerHandler("partintro", handlePartIntro)
	registerHandler("book", handleBook)
	registerHandler("reference", handleReference)
	registerHandler("bookinfo", handleSkip)
	registerHandler("title", handleSkip) // titles handled by parent
	registerHandler("titleabbrev", handleSkip)

	// Lists
	registerHandler("itemizedlist", handleItemizedList)
	registerHandler("orderedlist", handleOrderedList)
	registerHandler("variablelist", handleVariableList)
	registerHandler("simplelist", handleSimpleList)

	// Admonitions
	registerHandler("note", handleAdmonition)
	registerHandler("tip", handleAdmonition)
	registerHandler("warning", handleAdmonition)
	registerHandler("caution", handleAdmonition)
	registerHandler("important", handleAdmonition)

	// Code blocks
	registerHandler("programlisting", handleProgramListing)
	registerHandler("screen", handleScreen)
	registerHandler("literallayout", handleLiteralLayout)
	registerHandler("synopsis", handleSynopsis)

	// Tables
	registerHandler("table", handleTable)
	registerHandler("informaltable", handleTable)

	// Refentry (reference pages)
	registerHandler("refentry", handleRefentry)

	// Refsect sub-sections within refentries
	registerHandler("refsect1", handleSection)
	registerHandler("refsect2", handleSection)
	registerHandler("refsect3", handleSection)

	// Cross-references
	registerHandler("xref", handleXref)
	registerHandler("link", handleLink)
	registerHandler("ulink", handleUlink)
	registerHandler("anchor", handleAnchor)

	// Inline elements
	registerHandler("emphasis", handleEmphasis)
	registerHandler("literal", handleCode)
	registerHandler("command", handleCode)
	registerHandler("function", handleCode)
	registerHandler("type", handleCode)
	registerHandler("varname", handleCode)
	registerHandler("structfield", handleCode)
	registerHandler("structname", handleCode)
	registerHandler("constant", handleCode)
	registerHandler("option", handleCode)
	registerHandler("filename", handleCode)
	registerHandler("envar", handleCode)
	registerHandler("token", handleCode)
	registerHandler("symbol", handleCode)
	registerHandler("systemitem", handleCode)
	registerHandler("prompt", handleCode)
	registerHandler("classname", handleCode)
	registerHandler("errorcode", handleCode)
	registerHandler("errorname", handleCode)
	registerHandler("sgmltag", handleCode)
	registerHandler("replaceable", handleReplaceable)
	registerHandler("optional", handleOptional)
	registerHandler("quote", handleQuote)
	registerHandler("phrase", handlePassthrough)
	registerHandler("foreignphrase", handleEmphasis)
	registerHandler("firstterm", handleEmphasis)
	registerHandler("glossterm", handleEmphasis)
	registerHandler("citetitle", handleEmphasis)
	registerHandler("wordasword", handleEmphasis)
	registerHandler("superscript", handleSuperscript)
	registerHandler("subscript", handleSubscript)
	registerHandler("trademark", handleTrademark)
	registerHandler("productname", handlePassthrough)
	registerHandler("application", handlePassthrough)
	registerHandler("acronym", handlePassthrough)
	registerHandler("abbrev", handlePassthrough)
	registerHandler("keycap", handleCode)
	registerHandler("keycombo", handleKeycombo)
	registerHandler("userinput", handleCode)
	registerHandler("computeroutput", handleCode)
	registerHandler("returnvalue", handleCode)
	registerHandler("parameter", handleCode)

	// Block containers — pass through to children
	registerHandler("abstract", handlePassthrough)
	registerHandler("blockquote", handleBlockquote)
	registerHandler("sidebar", handlePassthrough)
	registerHandler("highlights", handlePassthrough)
	registerHandler("example", handleExample)
	registerHandler("informalexample", handlePassthrough)
	registerHandler("figure", handleFigure)
	registerHandler("informalfigure", handlePassthrough)
	registerHandler("mediaobject", handlePassthrough)
	registerHandler("imageobject", handlePassthrough)
	registerHandler("imagedata", handleImagedata)
	registerHandler("textobject", handlePassthrough)
	registerHandler("caption", handlePassthrough)
	registerHandler("epigraph", handleBlockquote)
	registerHandler("attribution", handleEmphasis)

	// Footnotes
	registerHandler("footnote", handleFootnote)

	// Index terms — omit from output
	registerHandler("indexterm", handleSkip)
	registerHandler("primary", handleSkip)
	registerHandler("secondary", handleSkip)
	registerHandler("see", handleSkip)
	registerHandler("seealso", handleSkip)

	// Command synopsis
	registerHandler("cmdsynopsis", handleCmdSynopsis)
	registerHandler("arg", handleArg)
	registerHandler("group", handleGroup)
	registerHandler("sbr", handleSbr)

	// Procedures
	registerHandler("procedure", handleProcedure)
	registerHandler("step", handleStep)
	registerHandler("substeps", handleSubsteps)

	// Email
	registerHandler("email", handleEmail)

	// Footnote references
	registerHandler("footnoteref", handleFootnoteRef)

	// Glossary
	registerHandler("glosslist", handlePassthrough)
	registerHandler("glossary", handlePassthrough)
	registerHandler("glossdiv", handlePassthrough)
	registerHandler("glossentry", handleGlossEntry)
	registerHandler("glossterm", handleEmphasis)
	registerHandler("glossdef", handlePassthrough)
	registerHandler("glosssee", handleGlossSee)
	registerHandler("glossseealso", handleGlossSeeAlso)

	// Bibliography
	registerHandler("bibliography", handleChapter)
	registerHandler("bibliodiv", handleBibliodiv)
	registerHandler("biblioentry", handleBiblioentry)
	registerHandler("authorgroup", handleSkip) // handled by biblioentry
	registerHandler("author", handleSkip)      // handled by biblioentry
	registerHandler("firstname", handleSkip)   // handled by biblioentry
	registerHandler("surname", handleSkip)     // handled by biblioentry
	registerHandler("pubdate", handleSkip)     // handled by biblioentry
	registerHandler("isbn", handleSkip)        // handled by biblioentry
	registerHandler("subtitle", handleSkip)    // handled by biblioentry
	registerHandler("edition", handleSkip)     // handled by biblioentry
	registerHandler("bibliomisc", handleSkip)
	registerHandler("publisher", handleSkip)

	// Cross-references to external man pages
	registerHandler("citerefentry", handleCiteRefEntry)
	registerHandler("refentrytitle", handleCode)
	registerHandler("manvolnum", handleManvolnum)

	// Misc
	registerHandler("comment", handleSkip)
	registerHandler("remark", handleSkip)
	registerHandler("beginpage", handleSkip)
	registerHandler("co", handleSkip)
	registerHandler("calloutlist", handlePassthrough)
	registerHandler("callout", handlePassthrough)
}
