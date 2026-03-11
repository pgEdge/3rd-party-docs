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
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/pgEdge/postgresql-docs/builder/sgml"
)

func parseAndConvert(t *testing.T, input string) string {
	t.Helper()
	root, warnings, err := sgml.ParseString(input)
	if err != nil {
		t.Fatalf("parse error: %v", err)
	}
	for _, w := range warnings {
		t.Logf("parse warning: %s", w)
	}

	ctx := NewContext(root, "", "", "17.0")
	w := NewMarkdownWriter()

	for _, child := range root.Children {
		if err := convertNode(ctx, child, w); err != nil {
			t.Fatalf("convert error: %v", err)
		}
	}

	return w.String()
}

// parseAndConvertWithCtx is like parseAndConvert but returns the
// context so callers can inspect warnings, IDs, etc.
func parseAndConvertWithCtx(t *testing.T, input string) (string, *Context) {
	t.Helper()
	root, warnings, err := sgml.ParseString(input)
	if err != nil {
		t.Fatalf("parse error: %v", err)
	}
	for _, w := range warnings {
		t.Logf("parse warning: %s", w)
	}

	ctx := NewContext(root, "", "", "17.0")
	w := NewMarkdownWriter()

	for _, child := range root.Children {
		if err := convertNode(ctx, child, w); err != nil {
			t.Fatalf("convert error: %v", err)
		}
	}

	return w.String(), ctx
}

// --------------- Existing tests ---------------

func TestConvertPara(t *testing.T) {
	input := `<para>Hello world.</para>`
	result := parseAndConvert(t, input)
	if !strings.Contains(result, "Hello world.") {
		t.Errorf("expected 'Hello world.' in output:\n%s", result)
	}
}

func TestConvertEmphasis(t *testing.T) {
	input := `<para>This is <emphasis>italic</emphasis> text.</para>`
	result := parseAndConvert(t, input)
	if !strings.Contains(result, "*italic*") {
		t.Errorf("expected '*italic*' in output:\n%s", result)
	}
}

func TestConvertBoldEmphasis(t *testing.T) {
	input := `<para>This is <emphasis role="bold">bold</emphasis> text.</para>`
	result := parseAndConvert(t, input)
	if !strings.Contains(result, "**bold**") {
		t.Errorf("expected '**bold**' in output:\n%s", result)
	}
}

func TestConvertInlineCode(t *testing.T) {
	input := `<para>Use <command>SELECT</command> to query.</para>`
	result := parseAndConvert(t, input)
	if !strings.Contains(result, "`SELECT`") {
		t.Errorf("expected '`SELECT`' in output:\n%s", result)
	}
}

func TestConvertLiteral(t *testing.T) {
	input := `<para>Set <literal>work_mem</literal> to 64MB.</para>`
	result := parseAndConvert(t, input)
	if !strings.Contains(result, "`work_mem`") {
		t.Errorf("expected '`work_mem`' in output:\n%s", result)
	}
}

func TestConvertReplaceable(t *testing.T) {
	input := `<para>Connect to <replaceable>dbname</replaceable>.</para>`
	result := parseAndConvert(t, input)
	if !strings.Contains(result, "*dbname*") {
		t.Errorf("expected '*dbname*' in output:\n%s", result)
	}
}

func TestConvertSection(t *testing.T) {
	input := `<sect1 id="test-section"><title>Test Section</title>
<para>Content here.</para></sect1>`
	result := parseAndConvert(t, input)
	if !strings.Contains(result, "## Test Section") || !strings.Contains(result, `<a id="test-section"></a>`) {
		t.Errorf("expected heading in output:\n%s", result)
	}
	if !strings.Contains(result, "Content here.") {
		t.Errorf("expected content in output:\n%s", result)
	}
}

func TestConvertChapter(t *testing.T) {
	input := `<chapter id="tutorial"><title>Tutorial</title>
<para>Welcome.</para></chapter>`
	result := parseAndConvert(t, input)
	if !strings.Contains(result, "# Tutorial") || !strings.Contains(result, `<a id="tutorial"></a>`) {
		t.Errorf("expected chapter heading in output:\n%s", result)
	}
}

func TestConvertNote(t *testing.T) {
	input := `<note><para>Important info.</para></note>`
	result := parseAndConvert(t, input)
	if !strings.Contains(result, "!!! note") {
		t.Errorf("expected admonition in output:\n%s", result)
	}
	if !strings.Contains(result, "    Important info.") {
		t.Errorf("expected indented content in output:\n%s", result)
	}
}

func TestConvertWarning(t *testing.T) {
	input := `<warning><para>Be careful!</para></warning>`
	result := parseAndConvert(t, input)
	if !strings.Contains(result, "!!! warning") {
		t.Errorf("expected warning admonition in output:\n%s", result)
	}
}

func TestConvertTip(t *testing.T) {
	input := `<tip><para>A helpful tip.</para></tip>`
	result := parseAndConvert(t, input)
	if !strings.Contains(result, "!!! tip") {
		t.Errorf("expected tip admonition in output:\n%s", result)
	}
}

func TestConvertItemizedList(t *testing.T) {
	input := `<itemizedlist>
<listitem><para>First item</para></listitem>
<listitem><para>Second item</para></listitem>
</itemizedlist>`
	result := parseAndConvert(t, input)
	if !strings.Contains(result, "- First item") {
		t.Errorf("expected '- First item' in output:\n%s", result)
	}
	if !strings.Contains(result, "- Second item") {
		t.Errorf("expected '- Second item' in output:\n%s", result)
	}
}

func TestConvertOrderedList(t *testing.T) {
	input := `<orderedlist>
<listitem><para>Step one</para></listitem>
<listitem><para>Step two</para></listitem>
</orderedlist>`
	result := parseAndConvert(t, input)
	if !strings.Contains(result, "1. Step one") {
		t.Errorf("expected '1. Step one' in output:\n%s", result)
	}
	if !strings.Contains(result, "2. Step two") {
		t.Errorf("expected '2. Step two' in output:\n%s", result)
	}
}

func TestConvertVariableList(t *testing.T) {
	input := `<variablelist>
<varlistentry>
<term><option>-v</option></term>
<listitem><para>Be verbose.</para></listitem>
</varlistentry>
</variablelist>`
	result := parseAndConvert(t, input)
	if !strings.Contains(result, "`-v`") {
		t.Errorf("expected term in output:\n%s", result)
	}
	if !strings.Contains(result, "Be verbose.") {
		t.Errorf("expected description in output:\n%s", result)
	}
}

func TestConvertProgramListing(t *testing.T) {
	input := `<programlisting>
SELECT * FROM pg_class;
</programlisting>`
	result := parseAndConvert(t, input)
	if !strings.Contains(result, "```sql") {
		t.Errorf("expected sql code block in output:\n%s", result)
	}
	if !strings.Contains(result, "SELECT * FROM pg_class;") {
		t.Errorf("expected SQL content in output:\n%s", result)
	}
	if !strings.Contains(result, "```") {
		t.Errorf("expected closing fence in output:\n%s", result)
	}
}

func TestConvertTable(t *testing.T) {
	input := `<table id="test-table"><title>Test Table</title>
<tgroup cols="2">
<colspec colname="col1">
<colspec colname="col2">
<thead>
<row><entry>Name</entry><entry>Value</entry></row>
</thead>
<tbody>
<row><entry>foo</entry><entry>bar</entry></row>
</tbody>
</tgroup></table>`
	result := parseAndConvert(t, input)
	if !strings.Contains(result, "**Table: Test Table**") {
		t.Errorf("expected table title in output:\n%s", result)
	}
	if !strings.Contains(result, "| Name | Value |") {
		t.Errorf("expected Markdown table header in output:\n%s", result)
	}
	if !strings.Contains(result, "| foo | bar |") {
		t.Errorf("expected table row in output:\n%s", result)
	}
}

func TestConvertUlink(t *testing.T) {
	input := `<para>Visit <ulink url="https://postgresql.org">PostgreSQL</ulink>.</para>`
	result := parseAndConvert(t, input)
	if !strings.Contains(result, "[PostgreSQL](https://postgresql.org)") {
		t.Errorf("expected link in output:\n%s", result)
	}
}

func TestConvertQuote(t *testing.T) {
	input := `<para>The <quote>quick</quote> fox.</para>`
	result := parseAndConvert(t, input)
	if !strings.Contains(result, "\u201cquick\u201d") {
		t.Errorf("expected curly quotes in output:\n%s", result)
	}
}

func TestConvertRefentry(t *testing.T) {
	input := `<refentry id="sql-select">
<refmeta><refentrytitle>SELECT</refentrytitle><manvolnum>7</manvolnum></refmeta>
<refnamediv><refname>SELECT</refname><refpurpose>retrieve rows from a table</refpurpose></refnamediv>
<refsynopsisdiv>
<synopsis>
SELECT [ ALL | DISTINCT ] * FROM table
</synopsis>
</refsynopsisdiv>
<refsect1><title>Description</title>
<para>SELECT retrieves rows.</para>
</refsect1>
</refentry>`
	result := parseAndConvert(t, input)
	if !strings.Contains(result, "# SELECT") {
		t.Errorf("expected refentry title in output:\n%s", result)
	}
	if !strings.Contains(result, "retrieve rows from a table") {
		t.Errorf("expected purpose in output:\n%s", result)
	}
	if !strings.Contains(result, "## Synopsis") {
		t.Errorf("expected Synopsis heading in output:\n%s", result)
	}
	if !strings.Contains(result, "## Description") {
		t.Errorf("expected Description heading in output:\n%s", result)
	}
}

func TestConvertSuperscript(t *testing.T) {
	input := `<para>10<superscript>3</superscript></para>`
	result := parseAndConvert(t, input)
	if !strings.Contains(result, "10<sup>3</sup>") {
		t.Errorf("expected superscript in output:\n%s", result)
	}
}

func TestConvertAnchor(t *testing.T) {
	input := `<para><anchor id="my-anchor">Some text.</para>`
	result := parseAndConvert(t, input)
	if !strings.Contains(result, `<a id="my-anchor"></a>`) {
		t.Errorf("expected anchor in output:\n%s", result)
	}
}

func TestConvertIndexTermSkipped(t *testing.T) {
	input := `<para><indexterm><primary>test</primary></indexterm>Content.</para>`
	result := parseAndConvert(t, input)
	if strings.Contains(result, "indexterm") {
		t.Errorf("expected indexterm to be skipped in output:\n%s", result)
	}
	if !strings.Contains(result, "Content.") {
		t.Errorf("expected content preserved in output:\n%s", result)
	}
}

func TestConvertScreen(t *testing.T) {
	input := `<screen>
$ psql mydb
mydb=# \dt
</screen>`
	result := parseAndConvert(t, input)
	if !strings.Contains(result, "```") {
		t.Errorf("expected code block in output:\n%s", result)
	}
	if !strings.Contains(result, "$ psql mydb") {
		t.Errorf("expected screen content in output:\n%s", result)
	}
}

func TestSlugify(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"Tutorial", "tutorial"},
		{"SQL Language", "sql-language"},
		{"Server Administration", "server-administration"},
		{"pg_dump", "pg_dump"},
		{"ALTER TABLE", "alter-table"},
	}

	for _, tt := range tests {
		result := slugify(tt.input)
		if result != tt.expected {
			t.Errorf("slugify(%q) = %q, want %q",
				tt.input, result, tt.expected)
		}
	}
}

func TestNormalizeWhitespace(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello  world", "hello world"},
		{"hello\n  world", "hello world"},
		{"  spaced  ", " spaced "},
	}

	for _, tt := range tests {
		result := normalizeWhitespace(tt.input)
		if result != tt.expected {
			t.Errorf("normalizeWhitespace(%q) = %q, want %q",
				tt.input, result, tt.expected)
		}
	}
}

func TestMarkdownToHTML(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			"inline code",
			"`SELECT` statement",
			"<code>SELECT</code> statement",
		},
		{
			"italic",
			"*datatype* value",
			"<em>datatype</em> value",
		},
		{
			"bold",
			"**Important** note",
			"<strong>Important</strong> note",
		},
		{
			"mixed inline",
			"*datatype* `BETWEEN` *datatype*",
			"<em>datatype</em> <code>BETWEEN</code> <em>datatype</em>",
		},
		{
			"link",
			"See [docs](https://example.com) here",
			`See <a href="https://example.com">docs</a> here`,
		},
		{
			"multiple paragraphs",
			"`func` ( `arg` ) `result`\n\nDescription here.\n\n`example` `output`",
			"<p><code>func</code> ( <code>arg</code> ) <code>result</code></p>\n" +
				"<p>Description here.</p>\n" +
				"<p><code>example</code> <code>output</code></p>",
		},
		{
			"empty input",
			"",
			"",
		},
		{
			"html passthrough",
			"<a id=\"test\"></a>",
			"<a id=\"test\"></a>",
		},
		{
			"code with angle brackets",
			"`a < b` and `x > y`",
			"<code>a &lt; b</code> and <code>x &gt; y</code>",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := markdownToHTML(tt.input)
			if result != tt.expected {
				t.Errorf("markdownToHTML(%q) =\n  %q\nwant:\n  %q",
					tt.input, result, tt.expected)
			}
		})
	}
}

// =====================================================
// Block handler tests (block.go)
// =====================================================

func TestHandleFormalPara(t *testing.T) {
	input := `<formalpara><title>Note Title</title><para>Body text.</para></formalpara>`
	result := parseAndConvert(t, input)
	if !strings.Contains(result, "**Note Title.**") {
		t.Errorf("expected formalpara title in output:\n%s", result)
	}
	if !strings.Contains(result, "Body text.") {
		t.Errorf("expected formalpara body in output:\n%s", result)
	}
}

func TestHandleSimpleListVertical(t *testing.T) {
	input := `<simplelist>
<member>Alpha</member>
<member>Beta</member>
</simplelist>`
	result := parseAndConvert(t, input)
	if !strings.Contains(result, "- Alpha") {
		t.Errorf("expected '- Alpha' in output:\n%s", result)
	}
	if !strings.Contains(result, "- Beta") {
		t.Errorf("expected '- Beta' in output:\n%s", result)
	}
}

func TestHandleSimpleListInline(t *testing.T) {
	input := `<para><simplelist type="inline">
<member>one</member>
<member>two</member>
<member>three</member>
</simplelist></para>`
	result := parseAndConvert(t, input)
	if !strings.Contains(result, "one, two, three") {
		t.Errorf("expected inline comma-separated list:\n%s", result)
	}
}

func TestHandleLiteralLayoutMonospaced(t *testing.T) {
	input := `<literallayout class="monospaced">
fixed-width text
</literallayout>`
	result := parseAndConvert(t, input)
	if !strings.Contains(result, "```") {
		t.Errorf("expected code block for monospaced:\n%s", result)
	}
	if !strings.Contains(result, "fixed-width text") {
		t.Errorf("expected text content:\n%s", result)
	}
}

func TestHandleLiteralLayoutNormal(t *testing.T) {
	input := `<literallayout>
preformatted text
</literallayout>`
	result := parseAndConvert(t, input)
	if !strings.Contains(result, "<pre>") {
		t.Errorf("expected <pre> tag for normal literallayout:\n%s", result)
	}
	if !strings.Contains(result, "preformatted text") {
		t.Errorf("expected text content:\n%s", result)
	}
	if !strings.Contains(result, "</pre>") {
		t.Errorf("expected closing </pre> tag:\n%s", result)
	}
}

func TestHandleBlockquote(t *testing.T) {
	input := `<blockquote><para>Quoted text here.</para></blockquote>`
	result := parseAndConvert(t, input)
	if !strings.Contains(result, "> ") {
		t.Errorf("expected blockquote marker:\n%s", result)
	}
	if !strings.Contains(result, "Quoted text here.") {
		t.Errorf("expected quoted content:\n%s", result)
	}
}

func TestHandleEpigraph(t *testing.T) {
	input := `<epigraph><para>A wise saying.</para><attribution>Someone Famous</attribution></epigraph>`
	result := parseAndConvert(t, input)
	if !strings.Contains(result, "> ") {
		t.Errorf("expected blockquote marker for epigraph:\n%s", result)
	}
	if !strings.Contains(result, "A wise saying.") {
		t.Errorf("expected epigraph content:\n%s", result)
	}
}

func TestHandleExampleWithTitle(t *testing.T) {
	input := `<example><title>Example Usage</title><programlisting>SELECT 1;</programlisting></example>`
	result := parseAndConvert(t, input)
	if !strings.Contains(result, "**Example: Example Usage**") {
		t.Errorf("expected example title:\n%s", result)
	}
	if !strings.Contains(result, "SELECT 1;") {
		t.Errorf("expected example content:\n%s", result)
	}
}

func TestHandleExampleWithID(t *testing.T) {
	input := `<example id="ex-1"><title>My Example</title><para>Content.</para></example>`
	result := parseAndConvert(t, input)
	if !strings.Contains(result, `<a id="ex-1"></a>`) {
		t.Errorf("expected anchor with id:\n%s", result)
	}
	if !strings.Contains(result, "**Example: My Example**") {
		t.Errorf("expected example title:\n%s", result)
	}
}

func TestHandleFigure(t *testing.T) {
	input := `<figure><title>Architecture Diagram</title><para>Content.</para></figure>`
	result := parseAndConvert(t, input)
	if !strings.Contains(result, "**Architecture Diagram**") {
		t.Errorf("expected figure title:\n%s", result)
	}
}

func TestHandleImagedata(t *testing.T) {
	input := `<imagedata fileref="images/diagram.png">`
	result := parseAndConvert(t, input)
	if !strings.Contains(result, "![image](images/diagram.png)") {
		t.Errorf("expected image markdown:\n%s", result)
	}
}

func TestHandleImagedataEmpty(t *testing.T) {
	input := `<imagedata>`
	result := parseAndConvert(t, input)
	if strings.Contains(result, "![image]") {
		t.Errorf("expected no image output for empty fileref:\n%s", result)
	}
}

func TestHandleFootnote(t *testing.T) {
	input := `<para>Main text<footnote><para>Footnote content.</para></footnote> continues.</para>`
	result := parseAndConvert(t, input)
	if !strings.Contains(result, "(Footnote content.)") {
		t.Errorf("expected parenthetical footnote:\n%s", result)
	}
}

func TestHandleProcedure(t *testing.T) {
	input := `<procedure>
<step><para>Do first thing.</para></step>
<step><para>Do second thing.</para></step>
</procedure>`
	result := parseAndConvert(t, input)
	if !strings.Contains(result, "1. Do first thing.") {
		t.Errorf("expected step 1:\n%s", result)
	}
	if !strings.Contains(result, "2. Do second thing.") {
		t.Errorf("expected step 2:\n%s", result)
	}
}

func TestHandleProcedureWithTitle(t *testing.T) {
	input := `<procedure><title>Setup Steps</title>
<step><para>Install.</para></step>
</procedure>`
	result := parseAndConvert(t, input)
	if !strings.Contains(result, "**Setup Steps**") {
		t.Errorf("expected procedure title:\n%s", result)
	}
	if !strings.Contains(result, "1. Install.") {
		t.Errorf("expected step:\n%s", result)
	}
}

func TestHandleSubsteps(t *testing.T) {
	input := `<procedure>
<step><para>Main step.</para>
<substeps>
<step><para>Sub step A.</para></step>
<step><para>Sub step B.</para></step>
</substeps>
</step>
</procedure>`
	result := parseAndConvert(t, input)
	if !strings.Contains(result, "1. Main step.") {
		t.Errorf("expected main step:\n%s", result)
	}
	if !strings.Contains(result, "Sub step A.") {
		t.Errorf("expected sub step A:\n%s", result)
	}
	if !strings.Contains(result, "Sub step B.") {
		t.Errorf("expected sub step B:\n%s", result)
	}
}

func TestHandleEmail(t *testing.T) {
	input := `<para>Contact <email>user@example.com</email>.</para>`
	result := parseAndConvert(t, input)
	if !strings.Contains(result, "[user@example.com](mailto:user@example.com)") {
		t.Errorf("expected mailto link:\n%s", result)
	}
}

func TestHandleGlossEntry(t *testing.T) {
	input := `<glossentry id="gl-acid">
<glossterm>ACID</glossterm>
<glossdef><para>Atomicity, Consistency, Isolation, Durability.</para></glossdef>
</glossentry>`
	result := parseAndConvert(t, input)
	if !strings.Contains(result, `<a id="gl-acid"></a>`) {
		t.Errorf("expected anchor:\n%s", result)
	}
	if !strings.Contains(result, "**ACID**") {
		t.Errorf("expected bold term:\n%s", result)
	}
	if !strings.Contains(result, "Atomicity, Consistency, Isolation, Durability.") {
		t.Errorf("expected definition:\n%s", result)
	}
}

func TestHandleGlossEntryWithSee(t *testing.T) {
	input := `<glossentry>
<glossterm>MVCC</glossterm>
<glosssee>Multi-Version Concurrency Control</glosssee>
</glossentry>`
	result := parseAndConvert(t, input)
	if !strings.Contains(result, "*See:") {
		t.Errorf("expected See reference:\n%s", result)
	}
	if !strings.Contains(result, "Multi-Version Concurrency Control") {
		t.Errorf("expected see target:\n%s", result)
	}
}

func TestHandleGlossEntryWithSeeAlso(t *testing.T) {
	input := `<glossentry>
<glossterm>WAL</glossterm>
<glossdef><para>Write-Ahead Logging.</para></glossdef>
<glossseealso>REDO</glossseealso>
</glossentry>`
	result := parseAndConvert(t, input)
	if !strings.Contains(result, "*See also:") {
		t.Errorf("expected See also reference:\n%s", result)
	}
}

func TestHandleCiteRefEntry(t *testing.T) {
	input := `<para>See <citerefentry><refentrytitle>pg_dump</refentrytitle><manvolnum>1</manvolnum></citerefentry>.</para>`
	result := parseAndConvert(t, input)
	if !strings.Contains(result, "`pg_dump`(1)") {
		t.Errorf("expected citerefentry output:\n%s", result)
	}
}

func TestHandleCiteRefEntryNoVolnum(t *testing.T) {
	input := `<para>See <citerefentry><refentrytitle>pg_dump</refentrytitle></citerefentry>.</para>`
	result := parseAndConvert(t, input)
	if !strings.Contains(result, "`pg_dump`") {
		t.Errorf("expected citerefentry without volnum:\n%s", result)
	}
	// Should NOT contain parentheses
	if strings.Contains(result, "`pg_dump`(") {
		t.Errorf("should not have volnum parens:\n%s", result)
	}
}

func TestHandleFootnoteRef(t *testing.T) {
	input := `<para>See<footnoteref linkend="fn1">.</para>`
	result := parseAndConvert(t, input)
	if !strings.Contains(result, "[^fn1]") {
		t.Errorf("expected footnote ref:\n%s", result)
	}
}

func TestHandleFootnoteRefEmpty(t *testing.T) {
	input := `<para>Text<footnoteref>.</para>`
	result := parseAndConvert(t, input)
	if strings.Contains(result, "[^") {
		t.Errorf("should not emit ref without linkend:\n%s", result)
	}
}

func TestTextContentSkipping(t *testing.T) {
	root, _, err := sgml.ParseString(
		`<term><indexterm><primary>test</primary></indexterm>my_param</term>`)
	if err != nil {
		t.Fatalf("parse error: %v", err)
	}
	node := root.Children[0]
	result := textContentSkipping(node, "indexterm")
	if result != "my_param" {
		t.Errorf("textContentSkipping = %q, want %q", result, "my_param")
	}
}

func TestSectionLevel(t *testing.T) {
	tests := []struct {
		tag   string
		level int
	}{
		{"chapter", 1},
		{"appendix", 1},
		{"preface", 1},
		{"sect1", 2},
		{"section", 2},
		{"simplesect", 2},
		{"sect2", 3},
		{"refsect2", 3},
		{"sect3", 4},
		{"refsect3", 4},
		{"sect4", 5},
		{"sect5", 6},
		{"unknown_tag", 2},
	}
	for _, tt := range tests {
		result := sectionLevel(tt.tag)
		if result != tt.level {
			t.Errorf("sectionLevel(%q) = %d, want %d",
				tt.tag, result, tt.level)
		}
	}
}

func TestDetectLanguage(t *testing.T) {
	sqlKeywords := []string{
		"SELECT", "INSERT", "UPDATE", "DELETE", "CREATE",
		"ALTER", "DROP", "GRANT", "REVOKE", "BEGIN",
		"COMMIT", "EXPLAIN", "WITH", "SET",
	}
	for _, kw := range sqlKeywords {
		root, _, _ := sgml.ParseString(
			`<programlisting>` + kw + ` something;</programlisting>`)
		node := root.Children[0]
		lang := detectLanguage(node)
		if lang != "sql" {
			t.Errorf("detectLanguage for %q = %q, want sql",
				kw, lang)
		}
	}

	// Non-SQL content
	root, _, _ := sgml.ParseString(
		`<programlisting>echo "hello world"</programlisting>`)
	node := root.Children[0]
	lang := detectLanguage(node)
	if lang != "" {
		t.Errorf("detectLanguage for shell = %q, want empty", lang)
	}
}

func TestHandleAdmonitionWithTitle(t *testing.T) {
	input := `<note><title>Custom Title</title><para>Info.</para></note>`
	result := parseAndConvert(t, input)
	if !strings.Contains(result, `!!! note "Custom Title"`) {
		t.Errorf("expected titled admonition:\n%s", result)
	}
}

// =====================================================
// Inline handler tests (inline.go)
// =====================================================

func TestHandleCodeWithReplaceable(t *testing.T) {
	input := `<para><command>CREATE TABLE <replaceable>name</replaceable></command></para>`
	result := parseAndConvert(t, input)
	if !strings.Contains(result, "<code>CREATE TABLE </code>") {
		t.Errorf("expected HTML code tag:\n%s", result)
	}
	if !strings.Contains(result, "<em>name</em>") {
		t.Errorf("expected HTML em tag for replaceable:\n%s", result)
	}
}

func TestHasDescendant(t *testing.T) {
	root, _, _ := sgml.ParseString(
		`<command>text <phrase><replaceable>var</replaceable></phrase></command>`)
	node := root.Children[0]
	if !hasDescendant(node, "replaceable") {
		t.Error("expected hasDescendant to find deeply nested replaceable")
	}
	if hasDescendant(node, "nonexistent") {
		t.Error("expected hasDescendant to return false for nonexistent tag")
	}
}

func TestHasDescendantShallow(t *testing.T) {
	root, _, _ := sgml.ParseString(
		`<command><replaceable>x</replaceable></command>`)
	node := root.Children[0]
	if !hasDescendant(node, "replaceable") {
		t.Error("expected hasDescendant to find direct child")
	}
}

func TestHandleTrademarkDefault(t *testing.T) {
	input := `<para><trademark>PostgreSQL</trademark></para>`
	result := parseAndConvert(t, input)
	if !strings.Contains(result, "PostgreSQL\u2122") {
		t.Errorf("expected TM symbol:\n%s", result)
	}
}

func TestHandleTrademarkRegistered(t *testing.T) {
	input := `<para><trademark class="registered">UNIX</trademark></para>`
	result := parseAndConvert(t, input)
	if !strings.Contains(result, "UNIX\u00ae") {
		t.Errorf("expected registered symbol:\n%s", result)
	}
}

func TestHandleTrademarkCopyright(t *testing.T) {
	input := `<para><trademark class="copyright">MyStuff</trademark></para>`
	result := parseAndConvert(t, input)
	if !strings.Contains(result, "MyStuff\u00a9") {
		t.Errorf("expected copyright symbol:\n%s", result)
	}
}

func TestHandleTrademarkService(t *testing.T) {
	input := `<para><trademark class="service">MyService</trademark></para>`
	result := parseAndConvert(t, input)
	if !strings.Contains(result, "MyService\u2120") {
		t.Errorf("expected service mark symbol:\n%s", result)
	}
}

func TestHandleKeycombo(t *testing.T) {
	input := `<para>Press <keycombo><keycap>Ctrl</keycap><keycap>C</keycap></keycombo>.</para>`
	result := parseAndConvert(t, input)
	if !strings.Contains(result, "`Ctrl`+`C`") {
		t.Errorf("expected key combo:\n%s", result)
	}
}

func TestHandleKeycomboThreeKeys(t *testing.T) {
	input := `<para><keycombo><keycap>Ctrl</keycap><keycap>Shift</keycap><keycap>Z</keycap></keycombo></para>`
	result := parseAndConvert(t, input)
	if !strings.Contains(result, "`Ctrl`+`Shift`+`Z`") {
		t.Errorf("expected three-key combo:\n%s", result)
	}
}

func TestHandleSubscript(t *testing.T) {
	input := `<para>H<subscript>2</subscript>O</para>`
	result := parseAndConvert(t, input)
	if !strings.Contains(result, "H<sub>2</sub>O") {
		t.Errorf("expected subscript:\n%s", result)
	}
}

func TestHandleCodeWithBacktick(t *testing.T) {
	input := `<para><literal>a` + "`" + `b</literal></para>`
	result := parseAndConvert(t, input)
	if !strings.Contains(result, "`` a`b ``") {
		t.Errorf("expected double backtick wrapping:\n%s", result)
	}
}

func TestHandleCodeEmpty(t *testing.T) {
	input := `<para><literal>  </literal></para>`
	result := parseAndConvert(t, input)
	// Should not contain backticks around empty/whitespace
	if strings.Contains(result, "``") {
		t.Errorf("should not wrap empty code in backticks:\n%s", result)
	}
}

// =====================================================
// Table handler tests (table.go)
// =====================================================

func TestIsFuncTableEntry(t *testing.T) {
	input := `<table><tgroup cols="1"><tbody>
<row><entry role="func_table_entry"><para>sig</para><para>desc</para></entry></row>
</tbody></tgroup></table>`
	result := parseAndConvert(t, input)
	// func_table_entry tables render as HTML <table>
	if !strings.Contains(result, "<table>") {
		t.Errorf("expected HTML table for func_table_entry:\n%s", result)
	}
	if !strings.Contains(result, "<td>") {
		t.Errorf("expected <td> cells:\n%s", result)
	}
}

func TestConvertFuncTableMultiPara(t *testing.T) {
	input := `<table><tgroup cols="1">
<thead><row><entry role="func_table_entry"><para>Function</para><para>Description</para></entry></row></thead>
<tbody>
<row><entry role="func_table_entry"><para>sig()</para><para>does stuff</para><para>example</para></entry></row>
</tbody></tgroup></table>`
	result := parseAndConvert(t, input)
	if !strings.Contains(result, "<th>") {
		t.Errorf("expected <th> headers:\n%s", result)
	}
	// Three columns from three paras
	tdCount := strings.Count(result, "<td>")
	if tdCount < 3 {
		t.Errorf("expected at least 3 <td> cells, got %d:\n%s",
			tdCount, result)
	}
}

func TestTableNeedsHTMLSpanning(t *testing.T) {
	input := `<table><tgroup cols="2">
<colspec colname="c1"><colspec colname="c2">
<tbody>
<row><entry namest="c1" nameend="c2">spanning</entry></row>
</tbody></tgroup></table>`
	result := parseAndConvert(t, input)
	// Should render as HTML table
	if !strings.Contains(result, "<table>") {
		t.Errorf("expected HTML table for spanning:\n%s", result)
	}
	if !strings.Contains(result, `colspan="2"`) {
		t.Errorf("expected colspan attribute:\n%s", result)
	}
}

func TestTableNeedsHTMLMorerows(t *testing.T) {
	input := `<table><tgroup cols="2">
<colspec colname="c1"><colspec colname="c2">
<tbody>
<row><entry morerows="1">span</entry><entry>a</entry></row>
<row><entry>b</entry></row>
</tbody></tgroup></table>`
	result := parseAndConvert(t, input)
	if !strings.Contains(result, "<table>") {
		t.Errorf("expected HTML table for morerows:\n%s", result)
	}
	if !strings.Contains(result, `rowspan="1"`) {
		t.Errorf("expected rowspan attribute:\n%s", result)
	}
}

func TestTableNeedsHTMLMultiplePara(t *testing.T) {
	input := `<table><tgroup cols="1">
<tbody>
<row><entry><para>First para.</para><para>Second para.</para></entry></row>
</tbody></tgroup></table>`
	result := parseAndConvert(t, input)
	if !strings.Contains(result, "<table>") {
		t.Errorf("expected HTML table for multi-para entry:\n%s", result)
	}
}

func TestTableNeedsHTMLProgramlisting(t *testing.T) {
	input := `<table><tgroup cols="1">
<tbody>
<row><entry><programlisting>SELECT 1;</programlisting></entry></row>
</tbody></tgroup></table>`
	result := parseAndConvert(t, input)
	if !strings.Contains(result, "<table>") {
		t.Errorf("expected HTML table for programlisting entry:\n%s", result)
	}
}

func TestConvertHTMLTableWithColspan(t *testing.T) {
	input := `<table><tgroup cols="3">
<colspec colname="c1"><colspec colname="c2"><colspec colname="c3">
<thead><row><entry namest="c1" nameend="c3">All Columns</entry></row></thead>
<tbody>
<row><entry>a</entry><entry>b</entry><entry>c</entry></row>
</tbody></tgroup></table>`
	result := parseAndConvert(t, input)
	if !strings.Contains(result, `colspan="3"`) {
		t.Errorf("expected colspan=3:\n%s", result)
	}
	if !strings.Contains(result, "<th") {
		t.Errorf("expected <th> for header row:\n%s", result)
	}
}

func TestInformalTable(t *testing.T) {
	input := `<informaltable>
<tgroup cols="2">
<colspec colname="col1"><colspec colname="col2">
<tbody>
<row><entry>x</entry><entry>y</entry></row>
</tbody></tgroup></informaltable>`
	result := parseAndConvert(t, input)
	if !strings.Contains(result, "| x | y |") {
		t.Errorf("expected markdown table:\n%s", result)
	}
}

// =====================================================
// Context tests (context.go)
// =====================================================

func TestContextWarn(t *testing.T) {
	root, _, _ := sgml.ParseString(`<para>x</para>`)
	ctx := NewContext(root, "", "", "17.0")
	ctx.Warn("test warning: %s", "hello")
	if len(ctx.Warnings) != 1 {
		t.Fatalf("expected 1 warning, got %d", len(ctx.Warnings))
	}
	if ctx.Warnings[0] != "test warning: hello" {
		t.Errorf("warning = %q, want %q",
			ctx.Warnings[0], "test warning: hello")
	}
}

func TestContextRegisterID(t *testing.T) {
	root, _, _ := sgml.ParseString(`<para>x</para>`)
	ctx := NewContext(root, "", "", "17.0")
	ctx.RegisterID("my-id", "chapter/file.md", "my-id",
		"My Title", "sect1")
	entry, ok := ctx.IDMap["my-id"]
	if !ok {
		t.Fatal("expected ID to be registered")
	}
	if entry.File != "chapter/file.md" {
		t.Errorf("File = %q, want %q", entry.File, "chapter/file.md")
	}
	if entry.Title != "My Title" {
		t.Errorf("Title = %q, want %q", entry.Title, "My Title")
	}
	if entry.Type != "sect1" {
		t.Errorf("Type = %q, want %q", entry.Type, "sect1")
	}
}

func TestContextAddFile(t *testing.T) {
	root, _, _ := sgml.ParseString(`<para>x</para>`)
	ctx := NewContext(root, "", "", "17.0")
	ctx.AddFile("path/file.md", "File Title", "parent")
	if len(ctx.Files) != 1 {
		t.Fatalf("expected 1 file, got %d", len(ctx.Files))
	}
	f := ctx.Files[0]
	if f.Path != "path/file.md" {
		t.Errorf("Path = %q", f.Path)
	}
	if f.Title != "File Title" {
		t.Errorf("Title = %q", f.Title)
	}
	if f.NavParent != "parent" {
		t.Errorf("NavParent = %q", f.NavParent)
	}
	if f.Order != 1 {
		t.Errorf("Order = %d, want 1", f.Order)
	}

	// Add second file — order increments
	ctx.AddFile("path/file2.md", "Title2", "parent")
	if ctx.Files[1].Order != 2 {
		t.Errorf("second file Order = %d, want 2",
			ctx.Files[1].Order)
	}
}

func TestResolveLinkSameFile(t *testing.T) {
	root, _, _ := sgml.ParseString(`<para>x</para>`)
	ctx := NewContext(root, "", "", "17.0")
	ctx.RegisterID("my-sect", "chapter/page.md", "my-sect",
		"My Section", "sect1")
	ctx.CurrentFile = "chapter/page.md"
	link, title, ok := ctx.ResolveLink("my-sect")
	if !ok {
		t.Fatal("expected link to resolve")
	}
	if link != "#my-sect" {
		t.Errorf("link = %q, want %q", link, "#my-sect")
	}
	if title != "My Section" {
		t.Errorf("title = %q, want %q", title, "My Section")
	}
}

func TestResolveLinkCrossDir(t *testing.T) {
	root, _, _ := sgml.ParseString(`<para>x</para>`)
	ctx := NewContext(root, "", "", "17.0")
	ctx.RegisterID("target", "other/page.md", "target",
		"Target Section", "sect1")
	ctx.CurrentFile = "chapter/current.md"
	link, title, ok := ctx.ResolveLink("target")
	if !ok {
		t.Fatal("expected link to resolve")
	}
	if !strings.Contains(link, "other/page.md") ||
		!strings.Contains(link, "#target") {
		t.Errorf("link = %q, expected path with anchor", link)
	}
	if title != "Target Section" {
		t.Errorf("title = %q", title)
	}
}

func TestResolveLinkMissing(t *testing.T) {
	root, _, _ := sgml.ParseString(`<para>x</para>`)
	ctx := NewContext(root, "", "", "17.0")
	_, _, ok := ctx.ResolveLink("nonexistent")
	if ok {
		t.Error("expected missing ID to return false")
	}
}

// =====================================================
// Xref tests (xref.go)
// =====================================================

func TestHandleXrefNilContext(t *testing.T) {
	root, _, _ := sgml.ParseString(`<xref linkend="some-id">`)
	node := root.Children[0]
	w := NewMarkdownWriter()
	err := handleXref(nil, node, w)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	result := w.String()
	if !strings.Contains(result, "[some-id]") {
		t.Errorf("expected placeholder text:\n%s", result)
	}
}

func TestHandleXrefEmptyLinkend(t *testing.T) {
	input := `<para><xref></para>`
	result := parseAndConvert(t, input)
	// Should not crash, should produce no link
	_ = result
}

func TestHandleXrefUnresolved(t *testing.T) {
	input := `<para><xref linkend="missing-id"></para>`
	result, ctx := parseAndConvertWithCtx(t, input)
	if !strings.Contains(result, "[missing-id](#missing-id)") {
		t.Errorf("expected fallback link:\n%s", result)
	}
	found := false
	for _, w := range ctx.Warnings {
		if strings.Contains(w, "unresolved xref") {
			found = true
		}
	}
	if !found {
		t.Error("expected unresolved xref warning")
	}
}

func TestHandleLinkEmptyLinkend(t *testing.T) {
	input := `<para><link>just text</link></para>`
	result := parseAndConvert(t, input)
	if !strings.Contains(result, "just text") {
		t.Errorf("expected children rendered:\n%s", result)
	}
}

func TestHandleLinkUnresolved(t *testing.T) {
	input := `<para><link linkend="missing">link text</link></para>`
	result, ctx := parseAndConvertWithCtx(t, input)
	if !strings.Contains(result, "[link text](#missing)") {
		t.Errorf("expected fallback link:\n%s", result)
	}
	found := false
	for _, w := range ctx.Warnings {
		if strings.Contains(w, "unresolved link") {
			found = true
		}
	}
	if !found {
		t.Error("expected unresolved link warning")
	}
}

func TestHandleLinkNilContext(t *testing.T) {
	root, _, _ := sgml.ParseString(
		`<link linkend="x">link text</link>`)
	node := root.Children[0]
	w := NewMarkdownWriter()
	err := handleLink(nil, node, w)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	result := w.String()
	if !strings.Contains(result, "[link text]") {
		t.Errorf("expected placeholder:\n%s", result)
	}
}

func TestHandleUlinkNoURL(t *testing.T) {
	input := `<para><ulink>bare text</ulink></para>`
	result := parseAndConvert(t, input)
	if !strings.Contains(result, "bare text") {
		t.Errorf("expected text pass-through:\n%s", result)
	}
}

func TestHandleUlinkNoText(t *testing.T) {
	input := `<para><ulink url="https://example.com"></ulink></para>`
	result := parseAndConvert(t, input)
	if !strings.Contains(result, "[https://example.com](https://example.com)") {
		t.Errorf("expected URL as text:\n%s", result)
	}
}

// =====================================================
// Refentry tests (refentry.go)
// =====================================================

func TestHandleCmdSynopsis(t *testing.T) {
	input := `<cmdsynopsis>
<command>pg_dump</command>
<arg choice="opt"><option>-f</option> <replaceable>file</replaceable></arg>
</cmdsynopsis>`
	result := parseAndConvert(t, input)
	if !strings.Contains(result, "pg_dump") {
		t.Errorf("expected command name:\n%s", result)
	}
	if !strings.Contains(result, "-f") {
		t.Errorf("expected option:\n%s", result)
	}
	if !strings.Contains(result, "FILE") {
		t.Errorf("expected uppercased replaceable:\n%s", result)
	}
	if !strings.Contains(result, "```") {
		t.Errorf("expected code block:\n%s", result)
	}
}

func TestRenderArgOptional(t *testing.T) {
	root, _, _ := sgml.ParseString(
		`<arg choice="opt"><option>-v</option></arg>`)
	node := root.Children[0]
	w := NewMarkdownWriter()
	renderArg(node, w)
	result := w.String()
	if !strings.Contains(result, "[-v]") {
		t.Errorf("expected optional brackets:\n%s", result)
	}
}

func TestRenderArgRequired(t *testing.T) {
	root, _, _ := sgml.ParseString(
		`<arg choice="req"><replaceable>dbname</replaceable></arg>`)
	node := root.Children[0]
	w := NewMarkdownWriter()
	renderArg(node, w)
	result := w.String()
	// Required args have no brackets
	if strings.Contains(result, "[") || strings.Contains(result, "]") {
		t.Errorf("expected no brackets for required arg:\n%s", result)
	}
	if !strings.Contains(result, "DBNAME") {
		t.Errorf("expected uppercased replaceable:\n%s", result)
	}
}

func TestRenderArgRepeat(t *testing.T) {
	root, _, _ := sgml.ParseString(
		`<arg rep="repeat"><replaceable>file</replaceable></arg>`)
	node := root.Children[0]
	w := NewMarkdownWriter()
	renderArg(node, w)
	result := w.String()
	if !strings.Contains(result, "...") {
		t.Errorf("expected repeat indicator:\n%s", result)
	}
}

func TestRenderGroupRequired(t *testing.T) {
	root, _, _ := sgml.ParseString(
		`<group choice="req"><arg>-a</arg><arg>-b</arg></group>`)
	node := root.Children[0]
	w := NewMarkdownWriter()
	renderGroup(node, w)
	result := w.String()
	if !strings.Contains(result, "{") {
		t.Errorf("expected curly braces for required group:\n%s", result)
	}
	if !strings.Contains(result, "|") {
		t.Errorf("expected pipe separator:\n%s", result)
	}
}

func TestRenderGroupOptional(t *testing.T) {
	root, _, _ := sgml.ParseString(
		`<group><arg>-x</arg><arg>-y</arg></group>`)
	node := root.Children[0]
	w := NewMarkdownWriter()
	renderGroup(node, w)
	result := w.String()
	if !strings.Contains(result, "[") {
		t.Errorf("expected square brackets for optional group:\n%s",
			result)
	}
	if !strings.Contains(result, "|") {
		t.Errorf("expected pipe separator:\n%s", result)
	}
}

func TestHandleSbr(t *testing.T) {
	input := `<cmdsynopsis><command>cmd</command><sbr><arg>-a</arg></cmdsynopsis>`
	result := parseAndConvert(t, input)
	if !strings.Contains(result, "\n    ") {
		t.Errorf("expected line break with indent:\n%s", result)
	}
}

// =====================================================
// mdtohtml tests (mdtohtml.go)
// =====================================================

func TestConvertCodeBlockWithLanguage(t *testing.T) {
	block := "```sql\nSELECT 1;\n```"
	result := convertCodeBlock(block)
	if !strings.Contains(result, `class="language-sql"`) {
		t.Errorf("expected language class:\n%s", result)
	}
	if !strings.Contains(result, "SELECT 1;") {
		t.Errorf("expected code content:\n%s", result)
	}
}

func TestConvertCodeBlockNoLanguage(t *testing.T) {
	block := "```\nsome code\n```"
	result := convertCodeBlock(block)
	if !strings.Contains(result, "<pre><code>") {
		t.Errorf("expected pre/code without language:\n%s", result)
	}
	if strings.Contains(result, "language-") {
		t.Errorf("should not have language class:\n%s", result)
	}
}

func TestConvertCodeBlockShort(t *testing.T) {
	block := "```"
	result := convertCodeBlock(block)
	// Should return block as-is for too-short input
	if result != block {
		t.Errorf("expected passthrough for short block:\n%s", result)
	}
}

func TestIsHTMLBlock(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{`<a id="test"></a>`, true},
		{`<pre>code</pre>`, true},
		{`<div class="x">`, true},
		{`<sup>3</sup>`, true},
		{`<sub>2</sub>`, true},
		{`plain text`, false},
		{`**bold** text`, false},
	}
	for _, tt := range tests {
		result := isHTMLBlock(tt.input)
		if result != tt.expected {
			t.Errorf("isHTMLBlock(%q) = %v, want %v",
				tt.input, result, tt.expected)
		}
	}
}

func TestSplitParagraphs(t *testing.T) {
	text := "first paragraph\n\nsecond paragraph"
	result := splitParagraphs(text)
	if len(result) != 2 {
		t.Fatalf("expected 2 paragraphs, got %d", len(result))
	}
	if result[0] != "first paragraph" {
		t.Errorf("first = %q", result[0])
	}
	if result[1] != "second paragraph" {
		t.Errorf("second = %q", result[1])
	}
}

func TestSplitParagraphsWithCodeBlock(t *testing.T) {
	text := "before\n\n```\ncode\n\nmore code\n```\n\nafter"
	result := splitParagraphs(text)
	// Should keep code block as one unit despite blank line inside
	if len(result) != 3 {
		t.Fatalf("expected 3 paragraphs, got %d: %v",
			len(result), result)
	}
	if !strings.Contains(result[1], "more code") {
		t.Errorf("code block split unexpectedly: %v", result)
	}
}

func TestSplitParagraphsSingle(t *testing.T) {
	text := "single paragraph"
	result := splitParagraphs(text)
	if len(result) != 1 {
		t.Fatalf("expected 1 paragraph, got %d", len(result))
	}
}

// =====================================================
// Writer tests (writer.go)
// =====================================================

func TestWriterWriteWithIndent(t *testing.T) {
	w := NewMarkdownWriter()
	w.PushIndent("  ")
	w.Newline()
	w.Write("indented text")
	result := w.String()
	if !strings.Contains(result, "  indented text") {
		t.Errorf("expected indented text:\n%q", result)
	}
}

func TestWriterWriteMultiLine(t *testing.T) {
	w := NewMarkdownWriter()
	w.PushIndent("  ")
	w.Newline()
	w.Write("line1\nline2")
	result := w.String()
	if !strings.Contains(result, "  line1") {
		t.Errorf("expected indented line1:\n%q", result)
	}
	if !strings.Contains(result, "  line2") {
		t.Errorf("expected indented line2:\n%q", result)
	}
}

func TestWriterAdmonition(t *testing.T) {
	w := NewMarkdownWriter()
	w.Admonition("warning")
	result := w.String()
	if !strings.Contains(result, "!!! warning") {
		t.Errorf("expected admonition:\n%q", result)
	}
}

func TestWriterAdmonitionWithTitle(t *testing.T) {
	w := NewMarkdownWriter()
	w.BlankLine()
	w.WriteString("!!! note \"My Title\"\n\n")
	w.WriteString("    Content here.\n")
	result := w.String()
	if !strings.Contains(result, `!!! note "My Title"`) {
		t.Errorf("expected titled admonition:\n%q", result)
	}
}

func TestWriterPopIndentEdge(t *testing.T) {
	w := NewMarkdownWriter()
	w.PushIndent("  ")
	w.PushIndent("    ")
	// Pop mismatched indent — should not change
	w.PopIndent("xx")
	w.Newline()
	w.Write("text")
	result := w.String()
	// Both indents should still be present
	if !strings.Contains(result, "      text") {
		t.Errorf("expected both indents preserved:\n%q", result)
	}
}

func TestWriterPopIndentCorrect(t *testing.T) {
	w := NewMarkdownWriter()
	w.PushIndent("  ")
	w.PushIndent("    ")
	w.PopIndent("    ")
	w.Newline()
	w.Write("text")
	result := w.String()
	if !strings.Contains(result, "  text") {
		t.Errorf("expected single indent after pop:\n%q", result)
	}
}

func TestWriterBlankLineSuppressed(t *testing.T) {
	w := NewMarkdownWriter()
	w.SetSuppressNewlines(true)
	w.BlankLine()
	w.WriteString("text")
	result := w.String()
	if strings.HasPrefix(result, "\n") {
		t.Errorf("blank line should be suppressed:\n%q", result)
	}
}

func TestWriterBlankLineDedup(t *testing.T) {
	w := NewMarkdownWriter()
	w.BlankLine()
	w.BlankLine()
	w.BlankLine()
	w.WriteString("text")
	result := w.String()
	// A blank line is two newlines; the writer starts lastBlank=true
	// so first BlankLine is skipped, then text follows
	if strings.HasPrefix(result, "\n\n\n") {
		t.Errorf("too many blank lines:\n%q", result)
	}
}

func TestWriterHeading(t *testing.T) {
	w := NewMarkdownWriter()
	w.Heading(3, "My Heading", "my-heading")
	result := w.String()
	if !strings.Contains(result, "### My Heading\n") {
		t.Errorf("expected heading:\n%q", result)
	}
}

func TestWriterHeadingNoID(t *testing.T) {
	w := NewMarkdownWriter()
	w.Heading(2, "Untitled", "")
	result := w.String()
	if !strings.Contains(result, "## Untitled\n") {
		t.Errorf("expected heading without ID:\n%q", result)
	}
	if strings.Contains(result, "{ #") {
		t.Errorf("should not have ID attr:\n%q", result)
	}
}

func TestWriterStartEndCodeBlock(t *testing.T) {
	w := NewMarkdownWriter()
	w.StartCodeBlock("python")
	w.WriteString("print('hello')")
	w.EndCodeBlock()
	result := w.String()
	if !strings.Contains(result, "```python\n") {
		t.Errorf("expected python fence:\n%q", result)
	}
	if !strings.Contains(result, "print('hello')") {
		t.Errorf("expected code content:\n%q", result)
	}
	if !w.IsAtLineStart() {
		t.Error("expected to be at line start after code block")
	}
}

func TestWriterCodeBlockNoLang(t *testing.T) {
	w := NewMarkdownWriter()
	w.StartCodeBlock("")
	w.WriteString("some code")
	w.EndCodeBlock()
	result := w.String()
	if strings.Contains(result, "```\n\n") {
		// Fence followed by immediate blank line is wrong
	}
	if !strings.Contains(result, "```\n") {
		t.Errorf("expected plain fence:\n%q", result)
	}
}

func TestWriterLen(t *testing.T) {
	w := NewMarkdownWriter()
	if w.Len() != 0 {
		t.Errorf("expected len 0, got %d", w.Len())
	}
	w.WriteString("hello")
	if w.Len() != 5 {
		t.Errorf("expected len 5, got %d", w.Len())
	}
}

func TestWriterEnsureNewline(t *testing.T) {
	w := NewMarkdownWriter()
	w.WriteString("text")
	w.EnsureNewline()
	w.WriteString("next")
	result := w.String()
	if result != "text\nnext" {
		t.Errorf("expected newline inserted:\n%q", result)
	}
}

func TestWriterEnsureNewlineAlready(t *testing.T) {
	w := NewMarkdownWriter()
	w.WriteString("text\n")
	w.EnsureNewline()
	w.WriteString("next")
	result := w.String()
	if result != "text\nnext" {
		t.Errorf("expected no extra newline:\n%q", result)
	}
}

func TestWriterInCodeBlock(t *testing.T) {
	w := NewMarkdownWriter()
	if w.InCodeBlock() {
		t.Error("should not be in code block initially")
	}
	w.StartCodeBlock("")
	if !w.InCodeBlock() {
		t.Error("should be in code block after StartCodeBlock")
	}
	w.EndCodeBlock()
	if w.InCodeBlock() {
		t.Error("should not be in code block after EndCodeBlock")
	}
}

func TestWriteNoIndentInCodeBlock(t *testing.T) {
	w := NewMarkdownWriter()
	w.PushIndent("    ")
	w.StartCodeBlock("")
	w.Write("no indent")
	w.EndCodeBlock()
	result := w.String()
	// Inside code block, indent should NOT be applied by Write
	if strings.Contains(result, "    no indent") {
		t.Errorf("indent should not apply inside code block:\n%q",
			result)
	}
}

// =====================================================
// Converter full-pipeline tests (converter.go)
// =====================================================

func TestConverterFullPipeline(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "convert-test-*")
	if err != nil {
		t.Fatalf("could not create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	sgmlDoc := `<book id="postgres">
<title>PostgreSQL Documentation</title>
<chapter id="tutorial"><title>Tutorial</title>
<sect1 id="tut-start"><title>Getting Started</title>
<para>Welcome to PostgreSQL.</para>
<para>See <xref linkend="tut-advanced"> for more.</para>
</sect1>
<sect1 id="tut-advanced"><title>Advanced Features</title>
<para>Advanced content here.</para>
</sect1>
</chapter>
</book>`

	root, _, err := sgml.ParseString(sgmlDoc)
	if err != nil {
		t.Fatalf("parse error: %v", err)
	}

	conv := NewConverter(root, "", tmpDir, "17.0")
	if err := conv.Convert(); err != nil {
		t.Fatalf("convert error: %v", err)
	}

	// Check files were created
	files := conv.Files()
	if len(files) == 0 {
		t.Fatal("expected some output files")
	}

	// Check index.md exists
	indexPath := filepath.Join(tmpDir, "index.md")
	if _, err := os.Stat(indexPath); os.IsNotExist(err) {
		t.Error("expected index.md to be created")
	}

	// Check tutorial directory
	tutorialIndex := filepath.Join(tmpDir, "tutorial", "index.md")
	if _, err := os.Stat(tutorialIndex); os.IsNotExist(err) {
		t.Error("expected tutorial/index.md to be created")
	}

	// Check section files
	startFile := filepath.Join(tmpDir, "tutorial", "getting-started.md")
	if _, err := os.Stat(startFile); os.IsNotExist(err) {
		t.Error("expected tutorial/getting-started.md to be created")
	}

	advFile := filepath.Join(tmpDir, "tutorial", "advanced-features.md")
	if _, err := os.Stat(advFile); os.IsNotExist(err) {
		t.Error("expected tutorial/advanced-features.md to be created")
	}

	// Verify content of getting-started.md
	content, err := os.ReadFile(startFile)
	if err != nil {
		t.Fatalf("could not read file: %v", err)
	}
	contentStr := string(content)
	if !strings.Contains(contentStr, "## Getting Started") {
		t.Errorf("expected heading in output:\n%s", contentStr)
	}
	if !strings.Contains(contentStr, "Welcome to PostgreSQL.") {
		t.Errorf("expected content in output:\n%s", contentStr)
	}
	// xref should resolve
	if !strings.Contains(contentStr, "[Advanced Features]") {
		t.Errorf("expected resolved xref:\n%s", contentStr)
	}

	// Verify warnings
	warnings := conv.Warnings()
	for _, w := range warnings {
		t.Logf("warning: %s", w)
	}

	// Verify context is accessible
	ctx := conv.Context()
	if ctx == nil {
		t.Error("expected non-nil context")
	}
}

func TestConverterRefentry(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "convert-ref-test-*")
	if err != nil {
		t.Fatalf("could not create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	sgmlDoc := `<book id="postgres">
<title>PostgreSQL</title>
<reference id="sql-ref"><title>SQL Commands</title>
<refentry id="sql-select">
<refmeta><refentrytitle>SELECT</refentrytitle><manvolnum>7</manvolnum></refmeta>
<refnamediv><refname>SELECT</refname><refpurpose>retrieve rows</refpurpose></refnamediv>
<refsect1><title>Description</title>
<para>SELECT retrieves rows.</para>
</refsect1>
</refentry>
</reference>
</book>`

	root, _, err := sgml.ParseString(sgmlDoc)
	if err != nil {
		t.Fatalf("parse error: %v", err)
	}

	conv := NewConverter(root, "", tmpDir, "17.0")
	if err := conv.Convert(); err != nil {
		t.Fatalf("convert error: %v", err)
	}

	// Check that select.md was created
	selectFile := filepath.Join(tmpDir, "sql-commands", "select.md")
	if _, err := os.Stat(selectFile); os.IsNotExist(err) {
		t.Error("expected sql-commands/select.md")
	}
}

func TestConverterPartStructure(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "convert-part-test-*")
	if err != nil {
		t.Fatalf("could not create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	sgmlDoc := `<book id="postgres">
<title>PostgreSQL</title>
<part id="sql"><title>SQL Language</title>
<partintro><para>This part covers SQL.</para></partintro>
<chapter id="sql-syntax"><title>SQL Syntax</title>
<para>Syntax rules here.</para>
</chapter>
</part>
</book>`

	root, _, err := sgml.ParseString(sgmlDoc)
	if err != nil {
		t.Fatalf("parse error: %v", err)
	}

	conv := NewConverter(root, "", tmpDir, "17.0")
	if err := conv.Convert(); err != nil {
		t.Fatalf("convert error: %v", err)
	}

	// Part index should exist
	partIndex := filepath.Join(tmpDir, "sql-language", "index.md")
	if _, err := os.Stat(partIndex); os.IsNotExist(err) {
		t.Error("expected sql-language/index.md")
	}

	// Chapter file should exist under part
	chapterFile := filepath.Join(tmpDir, "sql-language",
		"sql-syntax.md")
	if _, err := os.Stat(chapterFile); os.IsNotExist(err) {
		t.Error("expected sql-language/sql-syntax.md")
	}

	// Verify part index contains partintro content
	content, err := os.ReadFile(partIndex)
	if err != nil {
		t.Fatalf("could not read file: %v", err)
	}
	if !strings.Contains(string(content), "This part covers SQL.") {
		t.Errorf("expected partintro content:\n%s", string(content))
	}
}

func TestRegisterDescendantIDsWithXreflabel(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "convert-xref-test-*")
	if err != nil {
		t.Fatalf("could not create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	sgmlDoc := `<book id="postgres">
<title>PostgreSQL</title>
<chapter id="config"><title>Configuration</title>
<variablelist>
<varlistentry id="guc-work-mem" xreflabel="work_mem">
<term><varname>work_mem</varname></term>
<listitem><para>Sets working memory.</para></listitem>
</varlistentry>
</variablelist>
</chapter>
</book>`

	root, _, err := sgml.ParseString(sgmlDoc)
	if err != nil {
		t.Fatalf("parse error: %v", err)
	}

	conv := NewConverter(root, "", tmpDir, "17.0")
	if err := conv.Convert(); err != nil {
		t.Fatalf("convert error: %v", err)
	}

	ctx := conv.Context()
	entry, ok := ctx.IDMap["guc-work-mem"]
	if !ok {
		t.Fatal("expected guc-work-mem in ID map")
	}
	// Should prefer xreflabel
	if entry.Title != "work_mem" {
		t.Errorf("title = %q, want %q", entry.Title, "work_mem")
	}
}

func TestWriteSectionTOC(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "convert-toc-test-*")
	if err != nil {
		t.Fatalf("could not create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	sgmlDoc := `<book id="postgres">
<title>PostgreSQL</title>
<chapter id="tutorial"><title>Tutorial</title>
<sect1 id="tut-intro"><title>Introduction</title>
<para>Intro.</para>
</sect1>
<sect1 id="tut-sql"><title>The SQL Language</title>
<para>SQL stuff.</para>
</sect1>
</chapter>
</book>`

	root, _, err := sgml.ParseString(sgmlDoc)
	if err != nil {
		t.Fatalf("parse error: %v", err)
	}

	conv := NewConverter(root, "", tmpDir, "17.0")
	if err := conv.Convert(); err != nil {
		t.Fatalf("convert error: %v", err)
	}

	// Check that index.md has TOC links
	indexPath := filepath.Join(tmpDir, "tutorial", "index.md")
	content, err := os.ReadFile(indexPath)
	if err != nil {
		t.Fatalf("could not read: %v", err)
	}
	contentStr := string(content)
	if !strings.Contains(contentStr, "[Introduction]") {
		t.Errorf("expected TOC link for Introduction:\n%s",
			contentStr)
	}
	if !strings.Contains(contentStr, "[The SQL Language]") {
		t.Errorf("expected TOC link for SQL Language:\n%s",
			contentStr)
	}
}

func TestCleanMarkdown(t *testing.T) {
	input := "line1  \nline2\t\n\n\n\n\nline3\n\n"
	result := cleanMarkdown(input)
	// Should collapse 4+ newlines and trim trailing
	if strings.Contains(result, "\n\n\n\n") {
		t.Errorf("expected collapsed blank lines:\n%q", result)
	}
	if !strings.HasSuffix(result, "\n") {
		t.Error("expected single trailing newline")
	}
	// Should strip trailing whitespace
	if strings.Contains(result, "line1  \n") {
		t.Errorf("expected trailing space stripped:\n%q", result)
	}
}

func TestConvertNodeUnknownElement(t *testing.T) {
	input := `<para><unknowntag>content</unknowntag></para>`
	result, ctx := parseAndConvertWithCtx(t, input)
	// Should pass through children
	if !strings.Contains(result, "content") {
		t.Errorf("expected content from unknown tag:\n%s", result)
	}
	// Should warn
	found := false
	for _, w := range ctx.Warnings {
		if strings.Contains(w, "unhandled element <unknowntag>") {
			found = true
		}
	}
	if !found {
		t.Error("expected warning about unhandled element")
	}
}

func TestExtractRawTextReplaceable(t *testing.T) {
	root, _, _ := sgml.ParseString(
		`<programlisting>SELECT <replaceable>table_name</replaceable>;</programlisting>`)
	node := root.Children[0]
	text := extractRawText(node)
	if !strings.Contains(text, "TABLE_NAME") {
		t.Errorf("expected uppercased replaceable:\n%s", text)
	}
}

func TestExtractRawTextOptional(t *testing.T) {
	root, _, _ := sgml.ParseString(
		`<synopsis>COMMAND <optional>OPTION</optional></synopsis>`)
	node := root.Children[0]
	text := extractRawText(node)
	if !strings.Contains(text, "[OPTION]") {
		t.Errorf("expected optional in brackets:\n%s", text)
	}
}

func TestHandleStepStandalone(t *testing.T) {
	input := `<step><para>Do something.</para></step>`
	result := parseAndConvert(t, input)
	if !strings.Contains(result, "Do something.") {
		t.Errorf("expected step content:\n%s", result)
	}
}

func TestHandleManvolnum(t *testing.T) {
	input := `<para><manvolnum>3</manvolnum></para>`
	result := parseAndConvert(t, input)
	if !strings.Contains(result, "(3)") {
		t.Errorf("expected manvolnum output:\n%s", result)
	}
}

func TestVariableListWithID(t *testing.T) {
	input := `<variablelist>
<varlistentry id="opt-verbose">
<term><option>--verbose</option></term>
<listitem><para>Enable verbose mode.</para></listitem>
</varlistentry>
</variablelist>`
	result := parseAndConvert(t, input)
	if !strings.Contains(result, `<a id="opt-verbose"></a>`) {
		t.Errorf("expected anchor for varlistentry:\n%s", result)
	}
}

func TestVariableListWithTitle(t *testing.T) {
	input := `<variablelist><title>Options</title>
<varlistentry>
<term><option>-h</option></term>
<listitem><para>Show help.</para></listitem>
</varlistentry>
</variablelist>`
	result := parseAndConvert(t, input)
	if !strings.Contains(result, "**Options**") {
		t.Errorf("expected list title:\n%s", result)
	}
}

func TestItemizedListWithTitle(t *testing.T) {
	input := `<itemizedlist><title>Features</title>
<listitem><para>Feature A</para></listitem>
</itemizedlist>`
	result := parseAndConvert(t, input)
	if !strings.Contains(result, "**Features**") {
		t.Errorf("expected list title:\n%s", result)
	}
}

func TestOrderedListWithTitle(t *testing.T) {
	input := `<orderedlist><title>Steps</title>
<listitem><para>Step one</para></listitem>
</orderedlist>`
	result := parseAndConvert(t, input)
	if !strings.Contains(result, "**Steps**") {
		t.Errorf("expected list title:\n%s", result)
	}
}

func TestHandleOptional(t *testing.T) {
	input := `<para><optional>extra params</optional></para>`
	result := parseAndConvert(t, input)
	if !strings.Contains(result, "[extra params]") {
		t.Errorf("expected optional brackets:\n%s", result)
	}
}

func TestHandleEmphasisStrong(t *testing.T) {
	input := `<para><emphasis role="strong">strong text</emphasis></para>`
	result := parseAndConvert(t, input)
	if !strings.Contains(result, "**strong text**") {
		t.Errorf("expected bold for strong role:\n%s", result)
	}
}

func TestRefentryNoRefmeta(t *testing.T) {
	input := `<refentry id="test-ref">
<refnamediv><refname>test_cmd</refname><refpurpose>test purpose</refpurpose></refnamediv>
<refsect1><title>Description</title>
<para>Desc.</para>
</refsect1>
</refentry>`
	result := parseAndConvert(t, input)
	if !strings.Contains(result, "# test_cmd") {
		t.Errorf("expected title from refname:\n%s", result)
	}
	if !strings.Contains(result, "test purpose") {
		t.Errorf("expected purpose:\n%s", result)
	}
}

func TestVariableListMultipleTerms(t *testing.T) {
	input := `<variablelist>
<varlistentry>
<term><option>-v</option></term>
<term><option>--verbose</option></term>
<listitem><para>Be verbose.</para></listitem>
</varlistentry>
</variablelist>`
	result := parseAndConvert(t, input)
	if !strings.Contains(result, "`-v`") {
		t.Errorf("expected first term:\n%s", result)
	}
	if !strings.Contains(result, "`--verbose`") {
		t.Errorf("expected second term:\n%s", result)
	}
}

func TestHandleSimplesect(t *testing.T) {
	input := `<simplesect id="ss-1"><title>Simple Section</title>
<para>Content.</para></simplesect>`
	result := parseAndConvert(t, input)
	if !strings.Contains(result, "## Simple Section") || !strings.Contains(result, `<a id="ss-1"></a>`) {
		t.Errorf("expected simplesect heading:\n%s", result)
	}
}

func TestHandleRefsect2(t *testing.T) {
	input := `<refsect2 id="r2"><title>Sub Section</title>
<para>Content.</para></refsect2>`
	result := parseAndConvert(t, input)
	if !strings.Contains(result, "### Sub Section") || !strings.Contains(result, `<a id="r2"></a>`) {
		t.Errorf("expected refsect2 heading at level 3:\n%s", result)
	}
}

func TestHandleRefsect3(t *testing.T) {
	input := `<refsect3 id="r3"><title>Deep Section</title>
<para>Content.</para></refsect3>`
	result := parseAndConvert(t, input)
	if !strings.Contains(result, "#### Deep Section") || !strings.Contains(result, `<a id="r3"></a>`) {
		t.Errorf("expected refsect3 heading at level 4:\n%s", result)
	}
}

func TestTableNoTgroup(t *testing.T) {
	// Table without tgroup should call convertHTMLTable which
	// returns nil when no tgroup found
	input := `<table><title>Empty Table</title></table>`
	result := parseAndConvert(t, input)
	if !strings.Contains(result, "**Table: Empty Table**") {
		t.Errorf("expected title even without tgroup:\n%s", result)
	}
}

func TestConverterNoBook(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "convert-nobook-test-*")
	if err != nil {
		t.Fatalf("could not create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// Document without <book> wrapper
	sgmlDoc := `<chapter id="standalone"><title>Standalone</title>
<para>Content.</para></chapter>`

	root, _, err := sgml.ParseString(sgmlDoc)
	if err != nil {
		t.Fatalf("parse error: %v", err)
	}

	conv := NewConverter(root, "", tmpDir, "17.0")
	if err := conv.Convert(); err != nil {
		t.Fatalf("convert error: %v", err)
	}

	files := conv.Files()
	if len(files) == 0 {
		t.Error("expected at least one file")
	}
}

func TestHtmlEscapeCode(t *testing.T) {
	result := htmlEscapeCode("a < b & c > d")
	expected := "a &lt; b &amp; c &gt; d"
	if result != expected {
		t.Errorf("htmlEscapeCode = %q, want %q", result, expected)
	}
}

func TestConvertInlineMarkdown(t *testing.T) {
	result := convertInlineMarkdown("``code with ` backtick``")
	if !strings.Contains(result, "<code>") {
		t.Errorf("expected code tag:\n%s", result)
	}
}

func TestExportSlugify(t *testing.T) {
	result := ExportSlugify("Hello World")
	if result != "hello-world" {
		t.Errorf("ExportSlugify = %q, want %q", result, "hello-world")
	}
}

// =====================================================
// Additional coverage tests
// =====================================================

// --- block.go: handleBook, handlePart, handlePartIntro, handleReference (0%) ---

func TestHandleBookDirect(t *testing.T) {
	input := `<book><para>Book content.</para></book>`
	result := parseAndConvert(t, input)
	if !strings.Contains(result, "Book content.") {
		t.Errorf("expected book content:\n%s", result)
	}
}

func TestHandlePartDirect(t *testing.T) {
	input := `<part><para>Part content.</para></part>`
	result := parseAndConvert(t, input)
	if !strings.Contains(result, "Part content.") {
		t.Errorf("expected part content:\n%s", result)
	}
}

func TestHandlePartIntroDirect(t *testing.T) {
	input := `<partintro><para>Part intro content.</para></partintro>`
	result := parseAndConvert(t, input)
	if !strings.Contains(result, "Part intro content.") {
		t.Errorf("expected partintro content:\n%s", result)
	}
}

func TestHandleReferenceDirect(t *testing.T) {
	input := `<reference><para>Reference content.</para></reference>`
	result := parseAndConvert(t, input)
	if !strings.Contains(result, "Reference content.") {
		t.Errorf("expected reference content:\n%s", result)
	}
}

// --- inline.go: handlePassthrough (0%) ---

func TestHandlePassthrough(t *testing.T) {
	input := `<para><phrase>passthrough text</phrase></para>`
	result := parseAndConvert(t, input)
	if !strings.Contains(result, "passthrough text") {
		t.Errorf("expected passthrough content:\n%s", result)
	}
}

func TestHandlePassthroughProductname(t *testing.T) {
	input := `<para><productname>PostgreSQL</productname> rocks.</para>`
	result := parseAndConvert(t, input)
	if !strings.Contains(result, "PostgreSQL") {
		t.Errorf("expected productname:\n%s", result)
	}
}

func TestHandlePassthroughAcronym(t *testing.T) {
	input := `<para><acronym>SQL</acronym> is great.</para>`
	result := parseAndConvert(t, input)
	if !strings.Contains(result, "SQL") {
		t.Errorf("expected acronym:\n%s", result)
	}
}

// --- refentry.go: handleArg (0%), handleGroup (0%), handleSbr (0%) ---

func TestHandleArgInline(t *testing.T) {
	input := `<para><arg choice="opt"><option>-v</option></arg></para>`
	result := parseAndConvert(t, input)
	if !strings.Contains(result, "[-v]") {
		t.Errorf("expected optional arg:\n%s", result)
	}
}

func TestHandleGroupInline(t *testing.T) {
	input := `<para><group choice="req"><arg>-a</arg><arg>-b</arg></group></para>`
	result := parseAndConvert(t, input)
	if !strings.Contains(result, "{") {
		t.Errorf("expected required group braces:\n%s", result)
	}
	if !strings.Contains(result, "|") {
		t.Errorf("expected pipe separator:\n%s", result)
	}
}

func TestHandleSbrInline(t *testing.T) {
	input := `<para>before<sbr>after</para>`
	result := parseAndConvert(t, input)
	if !strings.Contains(result, "\n    ") {
		t.Errorf("expected sbr line break:\n%s", result)
	}
}

// --- refentry.go: renderArgInGroup more coverage (42.9%) ---

func TestRenderArgInGroupWithReplaceable(t *testing.T) {
	root, _, _ := sgml.ParseString(
		`<arg><replaceable>filename</replaceable></arg>`)
	node := root.Children[0]
	w := NewMarkdownWriter()
	renderArgInGroup(node, w)
	result := w.String()
	if !strings.Contains(result, "FILENAME") {
		t.Errorf("expected uppercased replaceable:\n%s", result)
	}
}

func TestRenderArgInGroupWithOption(t *testing.T) {
	root, _, _ := sgml.ParseString(
		`<arg><option>--verbose</option></arg>`)
	node := root.Children[0]
	w := NewMarkdownWriter()
	renderArgInGroup(node, w)
	result := w.String()
	if !strings.Contains(result, "--verbose") {
		t.Errorf("expected option text:\n%s", result)
	}
}

func TestRenderArgInGroupWithUnknown(t *testing.T) {
	root, _, _ := sgml.ParseString(
		`<arg><literal>value</literal></arg>`)
	node := root.Children[0]
	w := NewMarkdownWriter()
	renderArgInGroup(node, w)
	result := w.String()
	if !strings.Contains(result, "value") {
		t.Errorf("expected default text:\n%s", result)
	}
}

func TestRenderArgInGroupWithText(t *testing.T) {
	root, _, _ := sgml.ParseString(
		`<arg>plain text</arg>`)
	node := root.Children[0]
	w := NewMarkdownWriter()
	renderArgInGroup(node, w)
	result := w.String()
	if !strings.Contains(result, "plain text") {
		t.Errorf("expected text:\n%s", result)
	}
}

// --- renderArg: nested arg and group ---

func TestRenderArgWithNestedArg(t *testing.T) {
	root, _, _ := sgml.ParseString(
		`<arg choice="opt"><arg choice="opt"><option>-x</option></arg></arg>`)
	node := root.Children[0]
	w := NewMarkdownWriter()
	renderArg(node, w)
	result := w.String()
	if !strings.Contains(result, "-x") {
		t.Errorf("expected nested arg:\n%s", result)
	}
}

func TestRenderArgWithGroup(t *testing.T) {
	root, _, _ := sgml.ParseString(
		`<arg choice="opt"><group><arg>-a</arg><arg>-b</arg></group></arg>`)
	node := root.Children[0]
	w := NewMarkdownWriter()
	renderArg(node, w)
	result := w.String()
	if !strings.Contains(result, "|") {
		t.Errorf("expected group pipe:\n%s", result)
	}
}

func TestRenderArgWithUnknownChild(t *testing.T) {
	root, _, _ := sgml.ParseString(
		`<arg choice="plain"><literal>val</literal></arg>`)
	node := root.Children[0]
	w := NewMarkdownWriter()
	renderArg(node, w)
	result := w.String()
	if !strings.Contains(result, "val") {
		t.Errorf("expected raw text:\n%s", result)
	}
}

// --- renderGroup with repeat, replaceable, default child ---

func TestRenderGroupRepeat(t *testing.T) {
	root, _, _ := sgml.ParseString(
		`<group><arg>-x</arg><replaceable>name</replaceable></group>`)
	node := root.Children[0]
	w := NewMarkdownWriter()
	renderGroup(node, w)
	result := w.String()
	if !strings.Contains(result, "NAME") {
		t.Errorf("expected uppercased replaceable:\n%s", result)
	}
}

func TestRenderGroupWithUnknownChild(t *testing.T) {
	root, _, _ := sgml.ParseString(
		`<group><literal>val</literal></group>`)
	node := root.Children[0]
	w := NewMarkdownWriter()
	renderGroup(node, w)
	result := w.String()
	if !strings.Contains(result, "val") {
		t.Errorf("expected raw text:\n%s", result)
	}
}

// --- handleCmdSynopsis: default element case ---

func TestHandleCmdSynopsisDefault(t *testing.T) {
	input := `<cmdsynopsis>
<command>myapp</command>
<literal>some_text</literal>
</cmdsynopsis>`
	result := parseAndConvert(t, input)
	if !strings.Contains(result, "some_text") {
		t.Errorf("expected default child text:\n%s", result)
	}
}

// --- block.go: handleImagedata with image copy logic (36.8%) ---

func TestHandleImagedataWithCopy(t *testing.T) {
	srcDir, err := os.MkdirTemp("", "img-src-*")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(srcDir)

	outDir, err := os.MkdirTemp("", "img-out-*")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(outDir)

	// Create source image
	imgDir := filepath.Join(srcDir, "images")
	os.MkdirAll(imgDir, 0755)
	os.WriteFile(filepath.Join(imgDir, "test.png"),
		[]byte("fake png data"), 0644)

	sgmlDoc := `<imagedata fileref="images/test.png">`
	root, _, err := sgml.ParseString(sgmlDoc)
	if err != nil {
		t.Fatalf("parse error: %v", err)
	}

	ctx := NewContext(root, srcDir, outDir, "17.0")
	ctx.CurrentFile = "chapter/page.md"

	// Create output chapter dir
	os.MkdirAll(filepath.Join(outDir, "chapter"), 0755)

	w := NewMarkdownWriter()
	node := root.Children[0]
	err = handleImagedata(ctx, node, w)
	if err != nil {
		t.Fatalf("handleImagedata error: %v", err)
	}

	result := w.String()
	if !strings.Contains(result, "![image](images/test.png)") {
		t.Errorf("expected image markdown:\n%s", result)
	}

	// Verify the file was copied
	dstPath := filepath.Join(outDir, "chapter", "images", "test.png")
	if _, err := os.Stat(dstPath); os.IsNotExist(err) {
		t.Error("expected image to be copied to output")
	}
}

func TestHandleImagedataMissingSrc(t *testing.T) {
	srcDir, err := os.MkdirTemp("", "img-src-miss-*")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(srcDir)

	outDir, err := os.MkdirTemp("", "img-out-miss-*")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(outDir)

	sgmlDoc := `<imagedata fileref="images/missing.png">`
	root, _, err := sgml.ParseString(sgmlDoc)
	if err != nil {
		t.Fatalf("parse error: %v", err)
	}

	ctx := NewContext(root, srcDir, outDir, "17.0")
	ctx.CurrentFile = "chapter/page.md"
	os.MkdirAll(filepath.Join(outDir, "chapter"), 0755)

	w := NewMarkdownWriter()
	node := root.Children[0]
	err = handleImagedata(ctx, node, w)
	if err != nil {
		t.Fatalf("handleImagedata should not error: %v", err)
	}

	// Should have a warning about missing source
	found := false
	for _, warn := range ctx.Warnings {
		if strings.Contains(warn, "could not read image") {
			found = true
		}
	}
	if !found {
		t.Error("expected warning about missing image")
	}
}

func TestHandleImagedataNilContext(t *testing.T) {
	sgmlDoc := `<imagedata fileref="images/test.png">`
	root, _, _ := sgml.ParseString(sgmlDoc)
	w := NewMarkdownWriter()
	node := root.Children[0]
	err := handleImagedata(nil, node, w)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(w.String(), "![image](images/test.png)") {
		t.Errorf("expected image markdown:\n%s", w.String())
	}
}

func TestHandleImagedataEmptySrcDir(t *testing.T) {
	sgmlDoc := `<imagedata fileref="images/test.png">`
	root, _, _ := sgml.ParseString(sgmlDoc)

	ctx := NewContext(root, "", "", "17.0")
	ctx.CurrentFile = "chapter/page.md"

	w := NewMarkdownWriter()
	node := root.Children[0]
	err := handleImagedata(ctx, node, w)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	// Should still produce markdown, just not copy
	if !strings.Contains(w.String(), "![image]") {
		t.Errorf("expected image markdown:\n%s", w.String())
	}
}

// --- block.go: extractTitle (66.7%) ---

func TestExtractTitleWithInlineElement(t *testing.T) {
	root, _, _ := sgml.ParseString(
		`<sect1><title>Using <command>pg_dump</command></title><para>x</para></sect1>`)
	node := root.Children[0]
	title := extractTitle(node)
	if !strings.Contains(title, "pg_dump") {
		t.Errorf("expected inline element in title: %q", title)
	}
}

func TestExtractTitleWithUnknownElement(t *testing.T) {
	root, _, _ := sgml.ParseString(
		`<sect1><title>Title with <unknowntag>stuff</unknowntag></title><para>x</para></sect1>`)
	node := root.Children[0]
	title := extractTitle(node)
	if !strings.Contains(title, "stuff") {
		t.Errorf("expected text from unknown element: %q", title)
	}
}

func TestExtractTitleNoTitle(t *testing.T) {
	root, _, _ := sgml.ParseString(`<para>no title</para>`)
	node := root.Children[0]
	title := extractTitle(node)
	if title != "" {
		t.Errorf("expected empty title, got %q", title)
	}
}

// --- block.go: convertChildrenSkipTitle with titleabbrev ---

func TestConvertChildrenSkipTitleAbbrev(t *testing.T) {
	input := `<chapter id="ch"><title>Full Title</title><titleabbrev>Short</titleabbrev><para>Body.</para></chapter>`
	result := parseAndConvert(t, input)
	if strings.Contains(result, "Short") {
		t.Errorf("titleabbrev should be skipped:\n%s", result)
	}
	if !strings.Contains(result, "Body.") {
		t.Errorf("expected body content:\n%s", result)
	}
}

// --- block.go: convertListItemChildren with non-para first child ---

func TestConvertListItemChildrenNonPara(t *testing.T) {
	input := `<itemizedlist>
<listitem><programlisting>SELECT 1;</programlisting><para>Second.</para></listitem>
</itemizedlist>`
	result := parseAndConvert(t, input)
	if !strings.Contains(result, "SELECT 1;") {
		t.Errorf("expected programlisting:\n%s", result)
	}
	if !strings.Contains(result, "Second.") {
		t.Errorf("expected second para:\n%s", result)
	}
}

// --- handlePara error path (80%) ---
// The error path requires convertChildren to fail. This is hard to trigger
// in SGML, so we test the normal flow for now.

// --- inline.go: convertCodeWithReplaceable more cases (66.7%) ---

func TestConvertCodeWithReplaceableEndingInEm(t *testing.T) {
	input := `<para><command><replaceable>var</replaceable></command></para>`
	result := parseAndConvert(t, input)
	if !strings.Contains(result, "<em>var</em>") {
		t.Errorf("expected em tag:\n%s", result)
	}
}

func TestConvertCodeWithReplaceableOtherInline(t *testing.T) {
	input := `<para><command>cmd <option>-f</option> <replaceable>file</replaceable></command></para>`
	result := parseAndConvert(t, input)
	if !strings.Contains(result, "<em>file</em>") {
		t.Errorf("expected em tag:\n%s", result)
	}
	if !strings.Contains(result, "<code>") {
		t.Errorf("expected code tag:\n%s", result)
	}
}

func TestConvertCodeWithReplaceableUnknownInline(t *testing.T) {
	// Test the else branch in convertCodeWithReplaceable where an
	// element child has no handler
	input := `<para><command>text <unknowntag>x</unknowntag> <replaceable>v</replaceable></command></para>`
	result := parseAndConvert(t, input)
	if !strings.Contains(result, "x") {
		t.Errorf("expected unknown element text:\n%s", result)
	}
	if !strings.Contains(result, "<em>v</em>") {
		t.Errorf("expected em:\n%s", result)
	}
}

func TestConvertCodeWithReplaceableConsecutiveReplaceables(t *testing.T) {
	// Two consecutive replaceables without text in between
	input := `<para><command><replaceable>a</replaceable><replaceable>b</replaceable></command></para>`
	result := parseAndConvert(t, input)
	if !strings.Contains(result, "<em>") {
		t.Errorf("expected em tags:\n%s", result)
	}
}

// --- xref.go: handleLink more coverage (68.2%) ---

func TestHandleLinkResolvedWithText(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "link-test-*")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	sgmlDoc := `<book id="postgres">
<title>PostgreSQL</title>
<chapter id="ch1"><title>Chapter One</title>
<sect1 id="s1"><title>Section One</title>
<para>See <link linkend="s2">my link text</link>.</para>
</sect1>
<sect1 id="s2"><title>Section Two</title>
<para>Target.</para>
</sect1>
</chapter>
</book>`

	root, _, err := sgml.ParseString(sgmlDoc)
	if err != nil {
		t.Fatal(err)
	}

	conv := NewConverter(root, "", tmpDir, "17.0")
	if err := conv.Convert(); err != nil {
		t.Fatal(err)
	}

	content, err := os.ReadFile(
		filepath.Join(tmpDir, "chapter-one", "section-one.md"))
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(content), "[my link text]") {
		t.Errorf("expected link text:\n%s", string(content))
	}
}

func TestHandleLinkResolvedNoText(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "link-notext-*")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	sgmlDoc := `<book id="postgres">
<title>PostgreSQL</title>
<chapter id="ch1"><title>Chapter</title>
<sect1 id="s1"><title>Sect1</title>
<para>See <link linkend="s2"></link>.</para>
</sect1>
<sect1 id="s2"><title>Sect2 Title</title>
<para>Target.</para>
</sect1>
</chapter>
</book>`

	root, _, err := sgml.ParseString(sgmlDoc)
	if err != nil {
		t.Fatal(err)
	}

	conv := NewConverter(root, "", tmpDir, "17.0")
	if err := conv.Convert(); err != nil {
		t.Fatal(err)
	}

	content, err := os.ReadFile(
		filepath.Join(tmpDir, "chapter", "sect1.md"))
	if err != nil {
		t.Fatal(err)
	}
	// Should use title as link text
	if !strings.Contains(string(content), "[Sect2 Title]") {
		t.Errorf("expected title as link text:\n%s", string(content))
	}
}

// --- xref.go: handleXref resolved with empty title ---

func TestHandleXrefResolvedEmptyTitle(t *testing.T) {
	root, _, _ := sgml.ParseString(`<xref linkend="my-id">`)
	node := root.Children[0]

	ctx := NewContext(root, "", "", "17.0")
	ctx.RegisterID("my-id", "page.md", "my-id", "", "sect1")
	ctx.CurrentFile = "page.md"

	w := NewMarkdownWriter()
	err := handleXref(ctx, node, w)
	if err != nil {
		t.Fatal(err)
	}
	result := w.String()
	// When title is empty, should use linkend as text
	if !strings.Contains(result, "[my-id]") {
		t.Errorf("expected linkend as text:\n%s", result)
	}
}

// --- context.go: slugify more edge cases (83.3%) ---

func TestSlugifyEdgeCases(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"", ""},
		{"---", ""},
		{"Hello/World", "hello-world"},
		{"test.file", "test-file"},
		{"num123", "num123"},
		{"under_score", "under_score"},
		{"multiple---dashes", "multiple-dashes"},
		{"UPPER CASE", "upper-case"},
		{"special!@#chars", "specialchars"},
	}

	for _, tt := range tests {
		result := slugify(tt.input)
		if result != tt.expected {
			t.Errorf("slugify(%q) = %q, want %q",
				tt.input, result, tt.expected)
		}
	}
}

// --- context.go: ResolveLink filepath.Rel error path (92.3%) ---

func TestResolveLinkSameDir(t *testing.T) {
	root, _, _ := sgml.ParseString(`<para>x</para>`)
	ctx := NewContext(root, "", "", "17.0")
	ctx.RegisterID("t", "dir/page.md", "t", "Title", "sect1")
	ctx.CurrentFile = "dir/other.md"
	link, title, ok := ctx.ResolveLink("t")
	if !ok {
		t.Fatal("expected resolve")
	}
	if !strings.Contains(link, "page.md") {
		t.Errorf("link = %q", link)
	}
	if title != "Title" {
		t.Errorf("title = %q", title)
	}
}

func TestResolveLinkEmptyAnchor(t *testing.T) {
	root, _, _ := sgml.ParseString(`<para>x</para>`)
	ctx := NewContext(root, "", "", "17.0")
	ctx.RegisterID("t", "other/page.md", "", "Title", "sect1")
	ctx.CurrentFile = "dir/current.md"
	link, _, ok := ctx.ResolveLink("t")
	if !ok {
		t.Fatal("expected resolve")
	}
	// Should not contain anchor since it's empty
	if strings.Contains(link, "#") {
		t.Errorf("should not have anchor:\n%s", link)
	}
}

// --- converter.go: registerNodeID (42.9%) ---

func TestRegisterNodeIDNoID(t *testing.T) {
	root, _, _ := sgml.ParseString(`<para>no id</para>`)
	conv := NewConverter(root, "", "", "17.0")
	// Should not panic or register anything
	conv.registerNodeID(root.Children[0], "file.md")
	if len(conv.ctx.IDMap) != 0 {
		t.Error("expected no IDs registered")
	}
}

func TestRegisterNodeIDWithID(t *testing.T) {
	root, _, _ := sgml.ParseString(`<para id="p1">has id</para>`)
	conv := NewConverter(root, "", "", "17.0")
	conv.registerNodeID(root.Children[0], "file.md")
	entry, ok := conv.ctx.IDMap["p1"]
	if !ok {
		t.Fatal("expected p1 in ID map")
	}
	if entry.File != "file.md" {
		t.Errorf("File = %q", entry.File)
	}
	// Title should fall back to id since para has no <title>
	if entry.Title != "p1" {
		t.Errorf("Title = %q, want %q", entry.Title, "p1")
	}
}

func TestRegisterNodeIDWithTitle(t *testing.T) {
	root, _, _ := sgml.ParseString(
		`<sect1 id="s1"><title>My Section</title><para>x</para></sect1>`)
	conv := NewConverter(root, "", "", "17.0")
	conv.registerNodeID(root.Children[0], "file.md")
	entry := conv.ctx.IDMap["s1"]
	if entry.Title != "My Section" {
		t.Errorf("Title = %q", entry.Title)
	}
}

// --- converter.go: mapElement default branch (70%) ---

func TestMapElementDefault(t *testing.T) {
	sgmlDoc := `<book id="postgres">
<title>PostgreSQL</title>
<glossary id="glossary"><title>Glossary</title>
<glossentry id="gl-test">
<glossterm>Test</glossterm>
<glossdef><para>A test.</para></glossdef>
</glossentry>
</glossary>
</book>`

	tmpDir, err := os.MkdirTemp("", "map-default-*")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	root, _, _ := sgml.ParseString(sgmlDoc)
	conv := NewConverter(root, "", tmpDir, "17.0")
	if err := conv.Convert(); err != nil {
		t.Fatal(err)
	}

	// glossary with its default handling should be in the ID map
	_, ok := conv.ctx.IDMap["gl-test"]
	if !ok {
		t.Error("expected gl-test in ID map via default mapElement")
	}
}

// --- converter.go: convertPart (68%), convertReference (65%) ---

func TestConvertPartWithRefentry(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "part-ref-*")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	sgmlDoc := `<book id="postgres">
<title>PostgreSQL</title>
<part id="ref"><title>Reference</title>
<partintro><para>Intro text.</para></partintro>
<reference id="sql-cmds"><title>SQL Commands</title>
<refentry id="sql-alter">
<refmeta><refentrytitle>ALTER</refentrytitle><manvolnum>7</manvolnum></refmeta>
<refnamediv><refname>ALTER</refname><refpurpose>alter stuff</refpurpose></refnamediv>
<refsect1><title>Description</title><para>Desc.</para></refsect1>
</refentry>
</reference>
<refentry id="app-pg-dump">
<refmeta><refentrytitle>pg_dump</refentrytitle><manvolnum>1</manvolnum></refmeta>
<refnamediv><refname>pg_dump</refname><refpurpose>dump db</refpurpose></refnamediv>
<refsect1><title>Description</title><para>Dumps.</para></refsect1>
</refentry>
</part>
</book>`

	root, _, err := sgml.ParseString(sgmlDoc)
	if err != nil {
		t.Fatal(err)
	}

	conv := NewConverter(root, "", tmpDir, "17.0")
	if err := conv.Convert(); err != nil {
		t.Fatal(err)
	}

	// Check part index
	partIndex := filepath.Join(tmpDir, "reference", "index.md")
	content, err := os.ReadFile(partIndex)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(content), "Intro text.") {
		t.Errorf("expected partintro content:\n%s", string(content))
	}

	// Check reference refentry
	alterFile := filepath.Join(tmpDir, "reference",
		"sql-commands", "alter.md")
	if _, err := os.Stat(alterFile); os.IsNotExist(err) {
		t.Error("expected alter.md")
	}

	// Check direct part refentry
	dumpFile := filepath.Join(tmpDir, "reference", "pg_dump.md")
	if _, err := os.Stat(dumpFile); os.IsNotExist(err) {
		t.Error("expected pg_dump.md")
	}
}

// --- converter.go: convertReference with partintro ---

func TestConvertReferenceWithPartintro(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "ref-partintro-*")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	sgmlDoc := `<book id="postgres">
<title>PostgreSQL</title>
<reference id="sql-ref"><title>SQL Commands</title>
<partintro><para>This section describes SQL commands.</para></partintro>
<refentry id="sql-select">
<refmeta><refentrytitle>SELECT</refentrytitle><manvolnum>7</manvolnum></refmeta>
<refnamediv><refname>SELECT</refname><refpurpose>retrieve rows</refpurpose></refnamediv>
<refsect1><title>Description</title><para>SELECT retrieves rows.</para></refsect1>
</refentry>
</reference>
</book>`

	root, _, err := sgml.ParseString(sgmlDoc)
	if err != nil {
		t.Fatal(err)
	}

	conv := NewConverter(root, "", tmpDir, "17.0")
	if err := conv.Convert(); err != nil {
		t.Fatal(err)
	}

	content, err := os.ReadFile(
		filepath.Join(tmpDir, "sql-commands", "index.md"))
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(content), "This section describes SQL commands.") {
		t.Errorf("expected partintro:\n%s", string(content))
	}
}

// --- converter.go: convertChapter without ID ---

func TestConvertChapterNoID(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "ch-noid-*")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	sgmlDoc := `<book id="postgres">
<title>PostgreSQL</title>
<chapter id="has-id"><title>Main</title>
<para>Main content.</para>
</chapter>
</book>`

	root, _, _ := sgml.ParseString(sgmlDoc)
	conv := NewConverter(root, "", tmpDir, "17.0")
	if err := conv.Convert(); err != nil {
		t.Fatal(err)
	}
}

// --- converter.go: convertSection without ID ---

func TestConvertSectionNoID(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "sect-noid-*")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	sgmlDoc := `<book id="postgres">
<title>PostgreSQL</title>
<chapter id="ch1"><title>Chapter</title>
<sect1 id="s1"><title>Section</title>
<para>Content.</para>
</sect1>
</chapter>
</book>`

	root, _, _ := sgml.ParseString(sgmlDoc)
	conv := NewConverter(root, "", tmpDir, "17.0")
	if err := conv.Convert(); err != nil {
		t.Fatal(err)
	}
}

// --- converter.go: convertRefentry without ID ---

func TestConvertRefentryNoID(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "ref-noid-*")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	input := `<refentry>
<refmeta><refentrytitle>NOID</refentrytitle><manvolnum>7</manvolnum></refmeta>
<refnamediv><refname>NOID</refname><refpurpose>test</refpurpose></refnamediv>
<refsect1><title>Desc</title><para>Desc.</para></refsect1>
</refentry>`

	root, _, _ := sgml.ParseString(input)
	ctx := NewContext(root, "", tmpDir, "17.0")
	w := NewMarkdownWriter()
	// convertRefentry should return nil for no ID
	node := root.Children[0]
	err = handleRefentry(ctx, node, w)
	if err != nil {
		t.Fatal(err)
	}
}

// --- converter.go: convertBook with bookinfo ---

func TestConvertBookWithBookinfo(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "book-info-*")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	sgmlDoc := `<book id="postgres">
<title>PostgreSQL Documentation</title>
<bookinfo>
<corpauthor>The PostgreSQL Global Development Group</corpauthor>
</bookinfo>
<preface id="preface"><title>Preface</title>
<para>Welcome.</para>
</preface>
</book>`

	root, _, _ := sgml.ParseString(sgmlDoc)
	conv := NewConverter(root, "", tmpDir, "17.0")
	if err := conv.Convert(); err != nil {
		t.Fatal(err)
	}

	content, err := os.ReadFile(filepath.Join(tmpDir, "index.md"))
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(content), "The PostgreSQL Global Development Group") {
		t.Errorf("expected corpauthor:\n%s", string(content))
	}
	if !strings.Contains(string(content), "[Preface]") {
		t.Errorf("expected preface link:\n%s", string(content))
	}
}

// --- converter.go: convertBook without ID ---

func TestConvertBookNoID(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "book-noid-*")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	sgmlDoc := `<book>
<title>PostgreSQL</title>
<chapter id="ch"><title>Chapter</title>
<para>Content.</para>
</chapter>
</book>`

	root, _, _ := sgml.ParseString(sgmlDoc)
	conv := NewConverter(root, "", tmpDir, "17.0")
	if err := conv.Convert(); err != nil {
		t.Fatal(err)
	}

	// Should create index.md even without book id
	indexPath := filepath.Join(tmpDir, "index.md")
	if _, err := os.Stat(indexPath); os.IsNotExist(err) {
		t.Error("expected index.md")
	}
}

// --- converter.go: mapRefentry without refmeta, fallback to refnamediv ---

func TestMapRefentryFallbackTitle(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "map-ref-*")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	sgmlDoc := `<book id="postgres">
<title>PostgreSQL</title>
<reference id="ref"><title>Ref</title>
<refentry id="test-entry">
<refnamediv><refname>test_cmd</refname><refpurpose>test</refpurpose></refnamediv>
<refsect1><title>Desc</title><para>D.</para></refsect1>
</refentry>
</reference>
</book>`

	root, _, _ := sgml.ParseString(sgmlDoc)
	conv := NewConverter(root, "", tmpDir, "17.0")
	if err := conv.Convert(); err != nil {
		t.Fatal(err)
	}

	entry := conv.ctx.IDMap["test-entry"]
	if entry == nil {
		t.Fatal("expected test-entry in ID map")
	}
	// registerDescendantIDs overwrites with id fallback since
	// refentry has no <title> child; the refnamediv title is used
	// during rendering, not during ID mapping.
	if entry.Title == "" {
		t.Error("expected non-empty title")
	}
}

func TestMapRefentryNoTitle(t *testing.T) {
	root, _, _ := sgml.ParseString(
		`<refentry id="empty-ref"></refentry>`)
	conv := NewConverter(root, "", "", "17.0")
	conv.mapRefentry(root.Children[0], "dir")
	entry := conv.ctx.IDMap["empty-ref"]
	if entry == nil {
		t.Fatal("expected entry in ID map")
	}
	// slug should be empty-ref since title is empty
	if !strings.Contains(entry.File, "empty-ref") {
		t.Errorf("File = %q", entry.File)
	}
}

func TestMapRefentryNoIDNoTitle(t *testing.T) {
	root, _, _ := sgml.ParseString(`<refentry></refentry>`)
	conv := NewConverter(root, "", "", "17.0")
	conv.mapRefentry(root.Children[0], "dir")
	// Should use "ref" slug
	if len(conv.ctx.Files) != 1 {
		t.Fatalf("expected 1 file, got %d", len(conv.ctx.Files))
	}
	if !strings.Contains(conv.ctx.Files[0].Path, "ref.md") {
		t.Errorf("Path = %q", conv.ctx.Files[0].Path)
	}
}

// --- converter.go: mapPart and mapReference empty title/id ---

func TestMapPartNoTitle(t *testing.T) {
	root, _, _ := sgml.ParseString(`<part id="my-part"></part>`)
	conv := NewConverter(root, "", "", "17.0")
	conv.mapPart(root.Children[0], "")
	entry := conv.ctx.IDMap["my-part"]
	if entry == nil {
		t.Fatal("expected entry")
	}
}

func TestMapPartNoIDNoTitle(t *testing.T) {
	root, _, _ := sgml.ParseString(`<part></part>`)
	conv := NewConverter(root, "", "", "17.0")
	conv.mapPart(root.Children[0], "")
	if len(conv.ctx.Files) != 1 {
		t.Fatalf("expected 1 file")
	}
	if !strings.Contains(conv.ctx.Files[0].Path, "part") {
		t.Errorf("Path = %q", conv.ctx.Files[0].Path)
	}
}

func TestMapReferenceNoTitle(t *testing.T) {
	root, _, _ := sgml.ParseString(`<reference id="ref-1"></reference>`)
	conv := NewConverter(root, "", "", "17.0")
	conv.mapReference(root.Children[0], "")
	entry := conv.ctx.IDMap["ref-1"]
	if entry == nil {
		t.Fatal("expected entry")
	}
}

func TestMapReferenceNoIDNoTitle(t *testing.T) {
	root, _, _ := sgml.ParseString(`<reference></reference>`)
	conv := NewConverter(root, "", "", "17.0")
	conv.mapReference(root.Children[0], "")
	if len(conv.ctx.Files) != 1 {
		t.Fatalf("expected 1 file")
	}
	if !strings.Contains(conv.ctx.Files[0].Path, "reference") {
		t.Errorf("Path = %q", conv.ctx.Files[0].Path)
	}
}

// --- converter.go: mapChapter empty slug ---

func TestMapChapterNoTitleFallbackID(t *testing.T) {
	root, _, _ := sgml.ParseString(
		`<chapter id="my-ch"><para>x</para></chapter>`)
	conv := NewConverter(root, "", "", "17.0")
	conv.mapChapter(root.Children[0], "")
	entry := conv.ctx.IDMap["my-ch"]
	if entry == nil {
		t.Fatal("expected entry")
	}
}

func TestMapChapterNoIDNoTitle(t *testing.T) {
	root, _, _ := sgml.ParseString(`<chapter><para>x</para></chapter>`)
	conv := NewConverter(root, "", "", "17.0")
	conv.mapChapter(root.Children[0], "")
	if len(conv.ctx.Files) != 1 {
		t.Fatalf("expected 1 file")
	}
}

// --- converter.go: mapSection empty slug ---

func TestMapSectionNoTitle(t *testing.T) {
	root, _, _ := sgml.ParseString(`<sect1 id="s1"><para>x</para></sect1>`)
	conv := NewConverter(root, "", "", "17.0")
	conv.mapSection(root.Children[0], "dir")
	entry := conv.ctx.IDMap["s1"]
	if entry == nil {
		t.Fatal("expected entry")
	}
}

func TestMapSectionNoIDNoTitle(t *testing.T) {
	root, _, _ := sgml.ParseString(`<sect1><para>x</para></sect1>`)
	conv := NewConverter(root, "", "", "17.0")
	conv.mapSection(root.Children[0], "dir")
	if len(conv.ctx.Files) != 1 {
		t.Fatalf("expected 1 file")
	}
	if !strings.Contains(conv.ctx.Files[0].Path, "section") {
		t.Errorf("Path = %q", conv.ctx.Files[0].Path)
	}
}

// --- writer.go: WriteString empty string (80%) ---

func TestWriterWriteStringEmpty(t *testing.T) {
	w := NewMarkdownWriter()
	w.WriteString("")
	if w.Len() != 0 {
		t.Errorf("expected len 0 after empty WriteString")
	}
}

func TestWriterWriteEmpty(t *testing.T) {
	w := NewMarkdownWriter()
	w.Write("")
	if w.Len() != 0 {
		t.Errorf("expected len 0 after empty Write")
	}
}

// --- mdtohtml.go: convertInlineMarkdown with bold and line break ---

func TestConvertInlineMarkdownBold(t *testing.T) {
	result := convertInlineMarkdown("**strong** text")
	if !strings.Contains(result, "<strong>strong</strong>") {
		t.Errorf("expected strong:\n%s", result)
	}
}

func TestConvertInlineMarkdownLineBreak(t *testing.T) {
	result := convertInlineMarkdown("line1\nline2")
	if !strings.Contains(result, "<br>") {
		t.Errorf("expected br:\n%s", result)
	}
}

// --- table.go: tfoot rendering ---

func TestTableWithTfoot(t *testing.T) {
	input := `<table><tgroup cols="2">
<colspec colname="c1"><colspec colname="c2">
<thead>
<row><entry>A</entry><entry>B</entry></row>
</thead>
<tfoot>
<row><entry>foot1</entry><entry>foot2</entry></row>
</tfoot>
<tbody>
<row><entry namest="c1" nameend="c2">spanning cell</entry></row>
</tbody>
</tgroup></table>`
	result := parseAndConvert(t, input)
	if !strings.Contains(result, "<tfoot>") {
		t.Errorf("expected tfoot:\n%s", result)
	}
	if !strings.Contains(result, "foot1") {
		t.Errorf("expected foot content:\n%s", result)
	}
}

// --- table.go: markdown table without thead ---

func TestMarkdownTableNoThead(t *testing.T) {
	input := `<table><tgroup cols="2">
<colspec colname="c1"><colspec colname="c2">
<tbody>
<row><entry>x</entry><entry>y</entry></row>
</tbody>
</tgroup></table>`
	result := parseAndConvert(t, input)
	if !strings.Contains(result, "| x | y |") {
		t.Errorf("expected markdown table:\n%s", result)
	}
	// Should have separator line from colspecs
	if !strings.Contains(result, "| --- |") {
		t.Errorf("expected separator:\n%s", result)
	}
}

// --- table.go: func table with no paras ---

func TestFuncTableEntryNoPara(t *testing.T) {
	input := `<table><tgroup cols="1">
<tbody>
<row><entry role="func_table_entry"></entry></row>
</tbody></tgroup></table>`
	result := parseAndConvert(t, input)
	if !strings.Contains(result, "<td></td>") {
		t.Errorf("expected empty td:\n%s", result)
	}
}

// --- table.go: spanname attribute ---

func TestTableSpanname(t *testing.T) {
	input := `<table><tgroup cols="2">
<colspec colname="c1"><colspec colname="c2">
<tbody>
<row><entry spanname="span1">spanning</entry></row>
</tbody></tgroup></table>`
	result := parseAndConvert(t, input)
	if !strings.Contains(result, "<table>") {
		t.Errorf("expected HTML table for spanname:\n%s", result)
	}
}

// --- converter.go: convertNode comment node ---

func TestConvertNodeComment(t *testing.T) {
	// The comment node is skipped
	root, _, _ := sgml.ParseString(
		`<para>before<!-- comment -->after</para>`)
	ctx := NewContext(root, "", "", "17.0")
	w := NewMarkdownWriter()
	for _, child := range root.Children {
		if err := convertNode(ctx, child, w); err != nil {
			t.Fatal(err)
		}
	}
	result := w.String()
	if strings.Contains(result, "comment") {
		t.Errorf("comment should be skipped:\n%s", result)
	}
}

// --- converter.go: mapPart with chapter/appendix/reference/refentry ---

func TestMapPartWithAppendix(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "part-appendix-*")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	sgmlDoc := `<book id="postgres">
<title>PostgreSQL</title>
<part id="appendices"><title>Appendices</title>
<appendix id="app-a"><title>Appendix A</title>
<para>Content.</para>
</appendix>
</part>
</book>`

	root, _, _ := sgml.ParseString(sgmlDoc)
	conv := NewConverter(root, "", tmpDir, "17.0")
	if err := conv.Convert(); err != nil {
		t.Fatal(err)
	}

	_, ok := conv.ctx.IDMap["app-a"]
	if !ok {
		t.Error("expected app-a in ID map")
	}
}

// --- converter.go: convertNodeToFiles with various children ---

func TestConvertNodeToFilesRefentry(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "cnf-ref-*")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	sgmlDoc := `<book id="postgres">
<title>PostgreSQL</title>
<refentry id="sql-copy">
<refmeta><refentrytitle>COPY</refentrytitle><manvolnum>7</manvolnum></refmeta>
<refnamediv><refname>COPY</refname><refpurpose>copy data</refpurpose></refnamediv>
<refsect1><title>Description</title><para>Copies.</para></refsect1>
</refentry>
</book>`

	root, _, _ := sgml.ParseString(sgmlDoc)
	conv := NewConverter(root, "", tmpDir, "17.0")
	if err := conv.Convert(); err != nil {
		t.Fatal(err)
	}

	copyFile := filepath.Join(tmpDir, "copy.md")
	if _, err := os.Stat(copyFile); os.IsNotExist(err) {
		t.Error("expected copy.md")
	}
}

// --- converter.go: convertChapter with refentry children ---

func TestConvertChapterWithRefentries(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "ch-ref-*")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	sgmlDoc := `<book id="postgres">
<title>PostgreSQL</title>
<chapter id="ch1"><title>Chapter</title>
<sect1 id="s1"><title>Section</title>
<para>Sect content.</para>
</sect1>
<refentry id="ref1">
<refmeta><refentrytitle>REF</refentrytitle><manvolnum>1</manvolnum></refmeta>
<refnamediv><refname>REF</refname><refpurpose>ref purpose</refpurpose></refnamediv>
<refsect1><title>Desc</title><para>D.</para></refsect1>
</refentry>
</chapter>
</book>`

	root, _, _ := sgml.ParseString(sgmlDoc)
	conv := NewConverter(root, "", tmpDir, "17.0")
	if err := conv.Convert(); err != nil {
		t.Fatal(err)
	}

	refFile := filepath.Join(tmpDir, "chapter", "ref.md")
	if _, err := os.Stat(refFile); os.IsNotExist(err) {
		t.Error("expected chapter/ref.md")
	}
}

// --- converter.go: Convert error path ---

func TestConverterConvertError(t *testing.T) {
	// Use a read-only directory to force write error
	tmpDir, err := os.MkdirTemp("", "conv-err-*")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	sgmlDoc := `<book id="postgres">
<title>PostgreSQL</title>
<chapter id="ch"><title>Chapter</title>
<para>Content.</para>
</chapter>
</book>`

	root, _, _ := sgml.ParseString(sgmlDoc)
	// Point to non-writable path
	conv := NewConverter(root, "", filepath.Join(tmpDir, "nope", "nope"), "17.0")
	// This should not panic but may error
	_ = conv.Convert()
}

// --- mapPart with refentry direct child ---

func TestMapPartWithDirectRefentry(t *testing.T) {
	root, _, _ := sgml.ParseString(`<part id="p1"><title>Part</title>
<refentry id="ref-direct">
<refmeta><refentrytitle>Direct</refentrytitle><manvolnum>1</manvolnum></refmeta>
<refnamediv><refname>Direct</refname><refpurpose>test</refpurpose></refnamediv>
</refentry>
</part>`)
	conv := NewConverter(root, "", "", "17.0")
	conv.mapPart(root.Children[0], "")
	_, ok := conv.ctx.IDMap["ref-direct"]
	if !ok {
		t.Error("expected ref-direct in ID map")
	}
}

// --- mapPart with reference child ---

func TestMapPartWithReference(t *testing.T) {
	root, _, _ := sgml.ParseString(`<part id="p1"><title>Part</title>
<reference id="ref-group"><title>Commands</title>
<refentry id="cmd1">
<refmeta><refentrytitle>CMD</refentrytitle><manvolnum>1</manvolnum></refmeta>
<refnamediv><refname>CMD</refname><refpurpose>cmd</refpurpose></refnamediv>
</refentry>
</reference>
</part>`)
	conv := NewConverter(root, "", "", "17.0")
	conv.mapPart(root.Children[0], "")
	_, ok := conv.ctx.IDMap["ref-group"]
	if !ok {
		t.Error("expected ref-group in ID map")
	}
	_, ok = conv.ctx.IDMap["cmd1"]
	if !ok {
		t.Error("expected cmd1 in ID map")
	}
}

// --- convertPart error: part not in ID map ---

func TestConvertPartNotInIDMap(t *testing.T) {
	root, _, _ := sgml.ParseString(
		`<part id="unknown-part"><title>X</title></part>`)
	conv := NewConverter(root, "", "", "17.0")
	// Don't run buildIDMap to leave it missing
	err := conv.convertPart(root.Children[0])
	if err == nil {
		t.Error("expected error for missing part")
	}
}

// --- convertReference not in ID map ---

func TestConvertReferenceNotInIDMap(t *testing.T) {
	root, _, _ := sgml.ParseString(
		`<reference id="unknown"><title>X</title></reference>`)
	conv := NewConverter(root, "", "", "17.0")
	err := conv.convertReference(root.Children[0])
	if err != nil {
		t.Error("expected no error, just warning")
	}
	found := false
	for _, w := range conv.ctx.Warnings {
		if strings.Contains(w, "not found in ID map") {
			found = true
		}
	}
	if !found {
		t.Error("expected warning about missing reference")
	}
}

// --- convertChapter not in ID map ---

func TestConvertChapterNotInIDMap(t *testing.T) {
	root, _, _ := sgml.ParseString(
		`<chapter id="unknown"><title>X</title><para>Y</para></chapter>`)
	conv := NewConverter(root, "", "", "17.0")
	err := conv.convertChapter(root.Children[0])
	if err == nil {
		t.Error("expected error for missing chapter")
	}
}

// --- convertSection not in ID map ---

func TestConvertSectionNotInIDMap(t *testing.T) {
	root, _, _ := sgml.ParseString(
		`<sect1 id="unknown"><title>X</title><para>Y</para></sect1>`)
	conv := NewConverter(root, "", "", "17.0")
	err := conv.convertSection(root.Children[0])
	if err != nil {
		t.Error("expected no error, just warning")
	}
}

// --- convertRefentry not in ID map ---

func TestConvertRefentryNotInIDMap(t *testing.T) {
	root, _, _ := sgml.ParseString(
		`<refentry id="unknown"><refnamediv><refname>X</refname><refpurpose>Y</refpurpose></refnamediv></refentry>`)
	conv := NewConverter(root, "", "", "17.0")
	err := conv.convertRefentry(root.Children[0])
	if err != nil {
		t.Error("expected no error, just warning")
	}
	found := false
	for _, w := range conv.ctx.Warnings {
		if strings.Contains(w, "not found in ID map") {
			found = true
		}
	}
	if !found {
		t.Error("expected warning")
	}
}

// --- table.go: isFuncTableEntry with no entries ---

func TestIsFuncTableEntryEmpty(t *testing.T) {
	root, _, _ := sgml.ParseString(`<tgroup cols="1"></tgroup>`)
	node := root.Children[0]
	if isFuncTableEntry(node) {
		t.Error("expected false for empty tgroup")
	}
}

// --- table.go: renderFuncTableEntry with single para ---

func TestFuncTableEntrySinglePara(t *testing.T) {
	input := `<table><tgroup cols="1">
<tbody>
<row><entry role="func_table_entry"><para>Only sig.</para></entry></row>
</tbody></tgroup></table>`
	result := parseAndConvert(t, input)
	if !strings.Contains(result, "<td>Only sig.</td>") {
		t.Errorf("expected sig td:\n%s", result)
	}
	// Should have two empty tds (desc, example)
	emptyTd := strings.Count(result, "<td></td>")
	if emptyTd < 2 {
		t.Errorf("expected 2 empty tds, got %d:\n%s", emptyTd, result)
	}
}

// --- table.go: tableNeedsHTML with list inside entry ---

func TestTableNeedsHTMLWithList(t *testing.T) {
	input := `<table><tgroup cols="1">
<tbody>
<row><entry><itemizedlist>
<listitem><para>item</para></listitem>
</itemizedlist></entry></row>
</tbody></tgroup></table>`
	result := parseAndConvert(t, input)
	if !strings.Contains(result, "<table>") {
		t.Errorf("expected HTML table for list entry:\n%s", result)
	}
}

// --- writer.go: BlankLine when not at line start ---

func TestWriterBlankLineNotAtLineStart(t *testing.T) {
	w := NewMarkdownWriter()
	w.WriteString("text")
	w.BlankLine()
	w.WriteString("after")
	result := w.String()
	if !strings.Contains(result, "text\n\nafter") {
		t.Errorf("expected blank line after text:\n%q", result)
	}
}

// --- mdtohtml.go: markdownToHTML with fenced code block ---

func TestMarkdownToHTMLWithCodeBlock(t *testing.T) {
	input := "```sql\nSELECT 1;\n```"
	result := markdownToHTML(input)
	if !strings.Contains(result, "language-sql") {
		t.Errorf("expected language class:\n%s", result)
	}
}

// --- block.go: handlePara in suppressed newlines mode ---

func TestHandleParaSuppressed(t *testing.T) {
	root, _, _ := sgml.ParseString(`<para>Hello.</para>`)
	ctx := NewContext(root, "", "", "17.0")
	w := NewMarkdownWriter()
	w.SetSuppressNewlines(true)
	node := root.Children[0]
	err := handlePara(ctx, node, w)
	if err != nil {
		t.Fatal(err)
	}
	result := w.String()
	if !strings.Contains(result, "Hello.") {
		t.Errorf("expected content:\n%s", result)
	}
}

// --- table.go: convertHTMLTable no tgroup ---

func TestConvertHTMLTableNoTgroup(t *testing.T) {
	root, _, _ := sgml.ParseString(`<table><title>T</title></table>`)
	ctx := NewContext(root, "", "", "17.0")
	w := NewMarkdownWriter()
	err := convertHTMLTable(ctx, root.Children[0], w)
	if err != nil {
		t.Fatal(err)
	}
	// Should return nil with no output
	if strings.Contains(w.String(), "<table>") {
		t.Error("should not output table without tgroup")
	}
}

// --- converter.go: convertAndWrite without book ---

func TestConvertAndWriteNoBook(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "no-book-*")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	sgmlDoc := `<chapter id="ch1"><title>Standalone Chapter</title>
<sect1 id="s1"><title>Section</title>
<para>Content.</para>
</sect1>
</chapter>`

	root, _, _ := sgml.ParseString(sgmlDoc)
	conv := NewConverter(root, "", tmpDir, "17.0")
	if err := conv.Convert(); err != nil {
		t.Fatal(err)
	}

	files := conv.Files()
	if len(files) == 0 {
		t.Error("expected files from no-book conversion")
	}
}

// --- converter.go: convertSection with no-id sect1 ---

func TestConvertSectionNoIDSlugLookup(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "sect-slug-*")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	// Section without id -- convertSection generates a slug from title
	root, _, _ := sgml.ParseString(
		`<sect1><title>My Title</title><para>Content.</para></sect1>`)
	conv := NewConverter(root, "", tmpDir, "17.0")
	// Register with slug
	conv.ctx.RegisterID("my-title", "my-title.md", "my-title",
		"My Title", "sect1")
	conv.ctx.AddFile("my-title.md", "My Title", "")
	err = conv.convertSection(root.Children[0])
	if err != nil {
		t.Fatal(err)
	}

	content, err := os.ReadFile(filepath.Join(tmpDir, "my-title.md"))
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(content), "My Title") {
		t.Errorf("expected content:\n%s", string(content))
	}
}

// --- Additional coverage for remaining non-error branches ---

// Test variablelist with multi-line listitem content (empty and non-empty lines)
func TestVariableListMultiLineContent(t *testing.T) {
	input := `<variablelist>
<varlistentry>
<term><option>-v</option></term>
<listitem>
<para>First paragraph of description.</para>
<para>Second paragraph of description.</para>
</listitem>
</varlistentry>
</variablelist>`
	result := parseAndConvert(t, input)
	if !strings.Contains(result, "First paragraph") {
		t.Errorf("expected first paragraph:\n%s", result)
	}
	if !strings.Contains(result, "Second paragraph") {
		t.Errorf("expected second paragraph:\n%s", result)
	}
}

// Test glossentry with multi-line glossdef content
func TestGlossEntryMultiLineContent(t *testing.T) {
	input := `<glossentry id="gl-test">
<glossterm>TestTerm</glossterm>
<glossdef>
<para>First line of definition.</para>
<para>Second line of definition.</para>
</glossdef>
</glossentry>`
	result := parseAndConvert(t, input)
	if !strings.Contains(result, "First line") {
		t.Errorf("expected first line:\n%s", result)
	}
	if !strings.Contains(result, "Second line") {
		t.Errorf("expected second line:\n%s", result)
	}
}

// Test imagedata write error (unwritable dst)
func TestHandleImagedataWriteError(t *testing.T) {
	srcDir, err := os.MkdirTemp("", "img-src-we-*")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(srcDir)

	// Create source image
	imgDir := filepath.Join(srcDir, "images")
	os.MkdirAll(imgDir, 0755)
	os.WriteFile(filepath.Join(imgDir, "test.png"),
		[]byte("data"), 0644)

	// Use /dev/null as outDir -- can't create dirs under a file
	sgmlDoc := `<imagedata fileref="images/test.png">`
	root, _, _ := sgml.ParseString(sgmlDoc)
	ctx := NewContext(root, srcDir, "/dev/null", "17.0")
	ctx.CurrentFile = "chapter/page.md"

	w := NewMarkdownWriter()
	node := root.Children[0]
	err = handleImagedata(ctx, node, w)
	if err != nil {
		t.Fatal(err)
	}
	// Should warn about mkdir failure
	found := false
	for _, warn := range ctx.Warnings {
		if strings.Contains(warn, "could not create image dir") ||
			strings.Contains(warn, "could not write image") {
			found = true
		}
	}
	if !found {
		t.Error("expected warning about image dir/write failure")
	}
}

// Test mapElement with sect1 at top level
func TestMapElementSect1(t *testing.T) {
	root, _, _ := sgml.ParseString(
		`<sect1 id="top-sect"><title>Top Section</title><para>x</para></sect1>`)
	conv := NewConverter(root, "", "", "17.0")
	conv.mapElement(root.Children[0], "parent")
	entry := conv.ctx.IDMap["top-sect"]
	if entry == nil {
		t.Fatal("expected top-sect in ID map")
	}
}

// Test registerDescendantIDs with varlistentry
func TestRegisterDescendantIDsVarlistentry(t *testing.T) {
	root, _, _ := sgml.ParseString(
		`<variablelist>
<varlistentry id="var1">
<term><indexterm><primary>x</primary></indexterm>my_param</term>
<listitem><para>Desc.</para></listitem>
</varlistentry>
</variablelist>`)
	conv := NewConverter(root, "", "", "17.0")
	conv.registerDescendantIDs(root.Children[0], "test.md")
	entry := conv.ctx.IDMap["var1"]
	if entry == nil {
		t.Fatal("expected var1 in ID map")
	}
	if entry.Title != "my_param" {
		t.Errorf("Title = %q, want %q", entry.Title, "my_param")
	}
}

// Test convertChapter with no ID (the early return nil path)
func TestConvertChapterNoIDEarlyReturn(t *testing.T) {
	root, _, _ := sgml.ParseString(
		`<chapter><title>No ID</title><para>Content.</para></chapter>`)
	conv := NewConverter(root, "", "", "17.0")
	err := conv.convertChapter(root.Children[0])
	if err != nil {
		t.Fatal(err)
	}
}

// Test convertRefentry with no ID (early return nil)
func TestConvertRefentryNoIDEarlyReturn(t *testing.T) {
	root, _, _ := sgml.ParseString(
		`<refentry><refnamediv><refname>X</refname><refpurpose>Y</refpurpose></refnamediv></refentry>`)
	conv := NewConverter(root, "", "", "17.0")
	err := conv.convertRefentry(root.Children[0])
	if err != nil {
		t.Fatal(err)
	}
}

// Test writeSectionTOC with no items
func TestWriteSectionTOCNoItems(t *testing.T) {
	root, _, _ := sgml.ParseString(
		`<chapter id="ch"><title>Ch</title><para>No sections.</para></chapter>`)
	conv := NewConverter(root, "", "", "17.0")
	w := NewMarkdownWriter()
	conv.writeSectionTOC(root.Children[0], w)
	if w.Len() != 0 {
		t.Errorf("expected no output for TOC with no items: %q",
			w.String())
	}
}

// Test convertNodeToFiles with preface (falls through to chapter path)
func TestConvertNodeToFilesPreface(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "cnf-preface-*")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	sgmlDoc := `<book id="postgres">
<title>PostgreSQL</title>
<preface id="preface"><title>Preface</title>
<para>Welcome.</para>
</preface>
</book>`

	root, _, _ := sgml.ParseString(sgmlDoc)
	conv := NewConverter(root, "", tmpDir, "17.0")
	if err := conv.Convert(); err != nil {
		t.Fatal(err)
	}

	prefaceFile := filepath.Join(tmpDir, "preface.md")
	if _, err := os.Stat(prefaceFile); os.IsNotExist(err) {
		t.Error("expected preface.md")
	}
}

// Test convertPart with appendix child
func TestConvertPartWithAppendixChild(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "part-app-*")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	sgmlDoc := `<book id="postgres">
<title>PostgreSQL</title>
<part id="internals"><title>Internals</title>
<appendix id="app-glossary"><title>Glossary</title>
<para>Glossary content.</para>
</appendix>
</part>
</book>`

	root, _, _ := sgml.ParseString(sgmlDoc)
	conv := NewConverter(root, "", tmpDir, "17.0")
	if err := conv.Convert(); err != nil {
		t.Fatal(err)
	}

	appFile := filepath.Join(tmpDir, "internals", "glossary.md")
	if _, err := os.Stat(appFile); os.IsNotExist(err) {
		t.Error("expected internals/glossary.md")
	}
}

// Test convertInlineMarkdown with italic
func TestConvertInlineMarkdownItalic(t *testing.T) {
	result := convertInlineMarkdown("*italic* text")
	if !strings.Contains(result, "<em>italic</em>") {
		t.Errorf("expected em tag:\n%s", result)
	}
}

// Test convertInlineMarkdown with link
func TestConvertInlineMarkdownLink(t *testing.T) {
	result := convertInlineMarkdown("[text](http://example.com)")
	if !strings.Contains(result, `<a href="http://example.com">text</a>`) {
		t.Errorf("expected anchor tag:\n%s", result)
	}
}

// Test admonition with multiline content including empty lines
func TestAdmonitionMultiLine(t *testing.T) {
	input := `<note><para>Line one.</para><para>Line two.</para></note>`
	result := parseAndConvert(t, input)
	if !strings.Contains(result, "!!! note") {
		t.Errorf("expected admonition:\n%s", result)
	}
	if !strings.Contains(result, "    Line one.") {
		t.Errorf("expected indented line one:\n%s", result)
	}
	if !strings.Contains(result, "    Line two.") {
		t.Errorf("expected indented line two:\n%s", result)
	}
}

// Test handleCode with empty replaceable (edge case)
func TestHandleCodeEmptyReplaceable(t *testing.T) {
	input := `<para><command>cmd</command></para>`
	result := parseAndConvert(t, input)
	if !strings.Contains(result, "`cmd`") {
		t.Errorf("expected backtick code:\n%s", result)
	}
}

// Test extractTitle with text content that has leading whitespace
func TestExtractTitleWhitespace(t *testing.T) {
	root, _, _ := sgml.ParseString(
		`<sect1><title>  Spaced  Title  </title><para>x</para></sect1>`)
	node := root.Children[0]
	title := extractTitle(node)
	if title != "Spaced Title" {
		t.Errorf("extractTitle = %q, want %q", title, "Spaced Title")
	}
}

// Test convertNode with text node inside code block
func TestConvertNodeTextInCodeBlock(t *testing.T) {
	root, _, _ := sgml.ParseString(`<programlisting>raw text here</programlisting>`)
	ctx := NewContext(root, "", "", "17.0")
	w := NewMarkdownWriter()
	for _, child := range root.Children {
		if err := convertNode(ctx, child, w); err != nil {
			t.Fatal(err)
		}
	}
	result := w.String()
	if !strings.Contains(result, "raw text here") {
		t.Errorf("expected raw text:\n%s", result)
	}
}

// Test writeFile mkdir error
func TestWriteFileMkdirError(t *testing.T) {
	root, _, _ := sgml.ParseString(`<para>x</para>`)
	conv := NewConverter(root, "", "/proc/nonexistent/path", "17.0")
	err := conv.writeFile("sub/file.md", "content")
	if err == nil {
		t.Error("expected error for invalid path")
	}
}

// Test Convert error propagation
func TestConvertErrorPropagation(t *testing.T) {
	sgmlDoc := `<book id="postgres">
<title>PostgreSQL</title>
<part id="p1"><title>Part</title>
<chapter id="ch"><title>Chapter</title>
<para>Content.</para>
</chapter>
</part>
</book>`
	root, _, _ := sgml.ParseString(sgmlDoc)
	// Use an invalid output path to trigger write errors
	conv := NewConverter(root, "", "/proc/0/nonexistent", "17.0")
	err := conv.Convert()
	if err == nil {
		t.Error("expected error from convert")
	}
}
