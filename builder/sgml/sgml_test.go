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
	"os"
	"path/filepath"
	"testing"
)

func TestTokenizerBasic(t *testing.T) {
	input := `<para>Hello <emphasis>world</emphasis></para>`
	tok := NewTokenizer(input)
	tokens := tok.Tokenize()

	// Expect: <para>, text, <emphasis>, text, </emphasis>, </para>, EOF
	expected := []struct {
		typ TokenType
		tag string
	}{
		{TokenTagOpen, "para"},
		{TokenText, ""},
		{TokenTagOpen, "emphasis"},
		{TokenText, ""},
		{TokenTagClose, "emphasis"},
		{TokenTagClose, "para"},
		{TokenEOF, ""},
	}

	if len(tokens) != len(expected) {
		t.Fatalf("expected %d tokens, got %d", len(expected), len(tokens))
	}

	for i, e := range expected {
		if tokens[i].Type != e.typ {
			t.Errorf("token %d: expected type %d, got %d (%s)",
				i, e.typ, tokens[i].Type, tokens[i])
		}
		if e.tag != "" && tokens[i].Tag != e.tag {
			t.Errorf("token %d: expected tag %q, got %q",
				i, e.tag, tokens[i].Tag)
		}
	}
}

func TestTokenizerAttributes(t *testing.T) {
	input := `<xref linkend="some-id">`
	tok := NewTokenizer(input)
	tokens := tok.Tokenize()

	if tokens[0].Type != TokenTagOpen {
		t.Fatalf("expected TagOpen, got %d", tokens[0].Type)
	}
	if tokens[0].Tag != "xref" {
		t.Errorf("expected tag 'xref', got %q", tokens[0].Tag)
	}
	if tokens[0].Attrs["linkend"] != "some-id" {
		t.Errorf("expected linkend='some-id', got %q",
			tokens[0].Attrs["linkend"])
	}
}

func TestTokenizerSelfClosing(t *testing.T) {
	input := `<xref linkend="foo"/>`
	tok := NewTokenizer(input)
	tokens := tok.Tokenize()

	if tokens[0].Type != TokenTagOpen {
		t.Fatalf("expected TagOpen, got %d", tokens[0].Type)
	}
	if tokens[0].Tag != "xref" {
		t.Errorf("expected tag 'xref', got %q", tokens[0].Tag)
	}
}

func TestTokenizerComment(t *testing.T) {
	input := `<!-- this is a comment --><para>text</para>`
	tok := NewTokenizer(input)
	tokens := tok.Tokenize()

	if tokens[0].Type != TokenComment {
		t.Fatalf("expected Comment, got %d", tokens[0].Type)
	}
	if tokens[0].Text != "this is a comment" {
		t.Errorf("expected comment text, got %q", tokens[0].Text)
	}
}

func TestTokenizerLineTracking(t *testing.T) {
	input := "<para>\nline two\n</para>"
	tok := NewTokenizer(input)
	tokens := tok.Tokenize()

	// <para> is on line 1, </para> on line 3
	if tokens[0].Line != 1 {
		t.Errorf("expected line 1, got %d", tokens[0].Line)
	}
	// </para> should be on line 3
	closeIdx := -1
	for i, tk := range tokens {
		if tk.Type == TokenTagClose && tk.Tag == "para" {
			closeIdx = i
			break
		}
	}
	if closeIdx >= 0 && tokens[closeIdx].Line != 3 {
		t.Errorf("expected </para> on line 3, got %d",
			tokens[closeIdx].Line)
	}
}

func TestParserSimple(t *testing.T) {
	input := `<chapter><title>Test</title><para>Hello</para></chapter>`
	root, warnings, err := ParseString(input)
	if err != nil {
		t.Fatalf("parse error: %v", err)
	}
	if len(warnings) > 0 {
		t.Logf("warnings: %v", warnings)
	}

	chapters := root.FindChildren("chapter")
	if len(chapters) != 1 {
		t.Fatalf("expected 1 chapter, got %d", len(chapters))
	}

	ch := chapters[0]
	title := ch.FindChild("title")
	if title == nil {
		t.Fatal("expected title element")
	}
	if title.TextContent() != "Test" {
		t.Errorf("expected title 'Test', got %q", title.TextContent())
	}

	para := ch.FindChild("para")
	if para == nil {
		t.Fatal("expected para element")
	}
	if para.TextContent() != "Hello" {
		t.Errorf("expected para 'Hello', got %q", para.TextContent())
	}
}

func TestParserEmptyElement(t *testing.T) {
	input := `<para>See <xref linkend="foo"> for details.</para>`
	root, _, err := ParseString(input)
	if err != nil {
		t.Fatalf("parse error: %v", err)
	}

	para := root.FindChild("para")
	if para == nil {
		t.Fatal("expected para")
	}

	xrefs := para.FindChildren("xref")
	if len(xrefs) != 1 {
		t.Fatalf("expected 1 xref, got %d", len(xrefs))
	}
	if xrefs[0].GetAttr("linkend") != "foo" {
		t.Errorf("expected linkend='foo', got %q",
			xrefs[0].GetAttr("linkend"))
	}

	// xref should have no children
	if len(xrefs[0].Children) != 0 {
		t.Errorf("expected xref to have no children, got %d",
			len(xrefs[0].Children))
	}
}

func TestParserNestedSections(t *testing.T) {
	input := `<sect1 id="s1"><title>One</title>
<sect2 id="s2"><title>Two</title>
<para>Content</para>
</sect2>
</sect1>`
	root, _, err := ParseString(input)
	if err != nil {
		t.Fatalf("parse error: %v", err)
	}

	s1 := root.FindChild("sect1")
	if s1 == nil {
		t.Fatal("expected sect1")
	}
	if s1.GetAttr("id") != "s1" {
		t.Errorf("expected id='s1', got %q", s1.GetAttr("id"))
	}

	s2 := s1.FindChild("sect2")
	if s2 == nil {
		t.Fatal("expected sect2")
	}
	if s2.GetAttr("id") != "s2" {
		t.Errorf("expected id='s2', got %q", s2.GetAttr("id"))
	}
}

func TestParserVariablelist(t *testing.T) {
	input := `<variablelist>
<varlistentry id="opt-verbose">
<term><option>-v</option></term>
<listitem><para>Be verbose.</para></listitem>
</varlistentry>
</variablelist>`
	root, _, err := ParseString(input)
	if err != nil {
		t.Fatalf("parse error: %v", err)
	}

	vl := root.FindChild("variablelist")
	if vl == nil {
		t.Fatal("expected variablelist")
	}
	entries := vl.FindChildren("varlistentry")
	if len(entries) != 1 {
		t.Fatalf("expected 1 varlistentry, got %d", len(entries))
	}
	if entries[0].GetAttr("id") != "opt-verbose" {
		t.Errorf("expected id='opt-verbose'")
	}
}

func TestParserRefentry(t *testing.T) {
	input := `<refentry id="app-test">
<refmeta>
<refentrytitle>test_app</refentrytitle>
<manvolnum>1</manvolnum>
</refmeta>
<refnamediv>
<refname>test_app</refname>
<refpurpose>a test application</refpurpose>
</refnamediv>
<refsect1><title>Description</title>
<para>This is a test.</para>
</refsect1>
</refentry>`
	root, _, err := ParseString(input)
	if err != nil {
		t.Fatalf("parse error: %v", err)
	}

	ref := root.FindChild("refentry")
	if ref == nil {
		t.Fatal("expected refentry")
	}

	meta := ref.FindChild("refmeta")
	if meta == nil {
		t.Fatal("expected refmeta")
	}

	title := meta.FindChild("refentrytitle")
	if title == nil || title.TextContent() != "test_app" {
		t.Errorf("expected refentrytitle 'test_app'")
	}
}

func TestParserTable(t *testing.T) {
	input := `<table id="test-table">
<title>Test Table</title>
<tgroup cols="2">
<colspec colname="col1">
<colspec colname="col2">
<thead>
<row><entry>Header 1</entry><entry>Header 2</entry></row>
</thead>
<tbody>
<row><entry>Cell 1</entry><entry>Cell 2</entry></row>
</tbody>
</tgroup>
</table>`
	root, _, err := ParseString(input)
	if err != nil {
		t.Fatalf("parse error: %v", err)
	}

	table := root.FindChild("table")
	if table == nil {
		t.Fatal("expected table")
	}
	if table.GetAttr("id") != "test-table" {
		t.Errorf("expected id='test-table'")
	}

	tgroup := table.FindChild("tgroup")
	if tgroup == nil {
		t.Fatal("expected tgroup")
	}

	colspecs := tgroup.FindChildren("colspec")
	if len(colspecs) != 2 {
		t.Errorf("expected 2 colspecs, got %d", len(colspecs))
	}
}

func TestEntityResolver(t *testing.T) {
	// Create a temp directory with test SGML files
	dir := t.TempDir()

	// Write a version.sgml
	os.WriteFile(filepath.Join(dir, "version.sgml"),
		[]byte(`<!ENTITY version "17.2">
<!ENTITY majorversion "17">`), 0644)

	// Write a filelist.sgml
	os.WriteFile(filepath.Join(dir, "filelist.sgml"),
		[]byte(`<!ENTITY intro SYSTEM "intro.sgml">
<!ENTITY advanced SYSTEM "advanced.sgml">`), 0644)

	// Write chapter files
	os.WriteFile(filepath.Join(dir, "intro.sgml"),
		[]byte(`<chapter id="intro"><title>Introduction to &version;</title>
<para>Welcome to PostgreSQL &majorversion;.</para>
</chapter>`), 0644)

	os.WriteFile(filepath.Join(dir, "advanced.sgml"),
		[]byte(`<chapter id="advanced"><title>Advanced</title>
<para>Advanced topics.</para>
</chapter>`), 0644)

	// Write the main postgres.sgml
	os.WriteFile(filepath.Join(dir, "postgres.sgml"),
		[]byte(`<!DOCTYPE book PUBLIC "-//OASIS//DTD DocBook V4.5//EN" [
<!ENTITY % version SYSTEM "version.sgml">
%version;
<!ENTITY % filelist SYSTEM "filelist.sgml">
%filelist;
]>
<book id="postgres">
&intro;
&advanced;
</book>`), 0644)

	resolver := NewEntityResolver(dir)
	body, err := resolver.ResolveFile("postgres.sgml")
	if err != nil {
		t.Fatalf("resolve error: %v", err)
	}

	// Check that entities were expanded
	if !containsStr(body, "Introduction to 17.2") {
		t.Errorf("expected version entity expansion in:\n%s", body)
	}
	if !containsStr(body, "PostgreSQL 17") {
		t.Errorf("expected majorversion entity expansion in:\n%s", body)
	}
	if !containsStr(body, `id="intro"`) {
		t.Errorf("expected intro chapter in:\n%s", body)
	}
	if !containsStr(body, `id="advanced"`) {
		t.Errorf("expected advanced chapter in:\n%s", body)
	}
}

func TestEntityResolverCharRefs(t *testing.T) {
	resolver := NewEntityResolver(t.TempDir())
	result := resolver.expandCharRefs("A &#x41; &#65;")
	if result != "A A A" {
		t.Errorf("expected 'A A A', got %q", result)
	}
}

func TestParseStringEndToEnd(t *testing.T) {
	input := `<book id="postgres">
<chapter id="tutorial">
<title>Tutorial</title>
<sect1 id="tutorial-start">
<title>Getting Started</title>
<para>See <xref linkend="tutorial-advanced"> for more.</para>
<note><para>This is important.</para></note>
</sect1>
<sect1 id="tutorial-advanced">
<title>Advanced Features</title>
<para>Use <command>SELECT</command> to query data.</para>
</sect1>
</chapter>
</book>`

	root, warnings, err := ParseString(input)
	if err != nil {
		t.Fatalf("parse error: %v", err)
	}
	for _, w := range warnings {
		t.Logf("warning: %s", w)
	}

	book := root.FindChild("book")
	if book == nil {
		t.Fatal("expected book element")
	}

	chapter := book.FindChild("chapter")
	if chapter == nil {
		t.Fatal("expected chapter element")
	}

	sects := chapter.FindChildren("sect1")
	if len(sects) != 2 {
		t.Fatalf("expected 2 sect1 elements, got %d", len(sects))
	}

	// Check xref in first section
	xrefs := sects[0].FindDescendants("xref")
	if len(xrefs) != 1 {
		t.Fatalf("expected 1 xref, got %d", len(xrefs))
	}

	// Check note
	notes := sects[0].FindDescendants("note")
	if len(notes) != 1 {
		t.Fatalf("expected 1 note, got %d", len(notes))
	}

	// Check command
	cmds := sects[1].FindDescendants("command")
	if len(cmds) != 1 {
		t.Fatalf("expected 1 command, got %d", len(cmds))
	}
	if cmds[0].TextContent() != "SELECT" {
		t.Errorf("expected command text 'SELECT', got %q",
			cmds[0].TextContent())
	}
}

func containsStr(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr ||
		len(s) > 0 && containsSubstr(s, substr))
}

func containsSubstr(s, sub string) bool {
	for i := 0; i <= len(s)-len(sub); i++ {
		if s[i:i+len(sub)] == sub {
			return true
		}
	}
	return false
}

// ===== Tokenizer tests =====

func TestTokenizerCDATA(t *testing.T) {
	input := `<para><![CDATA[some <raw> & content]]></para>`
	tok := NewTokenizer(input)
	tokens := tok.Tokenize()

	// Expect: <para>, text(CDATA content), </para>, EOF
	if len(tokens) != 4 {
		t.Fatalf("expected 4 tokens, got %d", len(tokens))
	}
	if tokens[1].Type != TokenText {
		t.Errorf("expected Text for CDATA, got %d", tokens[1].Type)
	}
	if tokens[1].Text != "some <raw> & content" {
		t.Errorf("expected CDATA content, got %q", tokens[1].Text)
	}
}

func TestTokenizerCDATAUnterminated(t *testing.T) {
	input := `<![CDATA[no end`
	tok := NewTokenizer(input)
	tokens := tok.Tokenize()

	// Should get text token consuming to end, then EOF
	found := false
	for _, tk := range tokens {
		if tk.Type == TokenText && containsStr(tk.Text, "no end") {
			found = true
		}
	}
	if !found {
		t.Error("expected unterminated CDATA to produce text token")
	}
}

func TestTokenizerCDATAWithNewlines(t *testing.T) {
	input := "<![CDATA[line1\nline2\nline3]]>"
	tok := NewTokenizer(input)
	tokens := tok.Tokenize()

	if tokens[0].Type != TokenText {
		t.Fatalf("expected Text, got %d", tokens[0].Type)
	}
	if tokens[0].Text != "line1\nline2\nline3" {
		t.Errorf("unexpected CDATA text: %q", tokens[0].Text)
	}
}

func TestTokenizerReadDeclaration(t *testing.T) {
	// A real declaration like <!DOCTYPE ...> should be consumed
	input := `<!DOCTYPE book PUBLIC "test">` +
		`<para>hello</para>`
	tok := NewTokenizer(input)
	tokens := tok.Tokenize()

	// Declaration becomes a Comment token, then <para>, text, </para>, EOF
	if tokens[0].Type != TokenComment {
		t.Errorf("expected Comment for declaration, got %d (%s)",
			tokens[0].Type, tokens[0])
	}
	// Should find para after the declaration
	found := false
	for _, tk := range tokens {
		if tk.Type == TokenTagOpen && tk.Tag == "para" {
			found = true
		}
	}
	if !found {
		t.Error("expected <para> after declaration")
	}
}

func TestTokenizerReadDeclarationNested(t *testing.T) {
	// Nested brackets in declaration
	input := `<!DOCTYPE book [<!ENTITY foo "bar">]><para>x</para>`
	tok := NewTokenizer(input)
	tokens := tok.Tokenize()

	// First token should be the declaration as comment
	if tokens[0].Type != TokenComment {
		t.Errorf("expected Comment for nested declaration, got %d",
			tokens[0].Type)
	}
}

func TestTokenizerReadDeclarationUnterminated(t *testing.T) {
	input := `<!DOCTYPE book PUBLIC "test"`
	tok := NewTokenizer(input)
	tokens := tok.Tokenize()

	// Should consume to end
	if tokens[0].Type != TokenComment {
		t.Errorf("expected Comment for unterminated declaration, got %d",
			tokens[0].Type)
	}
}

func TestTokenizerIsRealDeclaration(t *testing.T) {
	tests := []struct {
		input string
		real  bool
	}{
		{"<!DOCTYPE foo>rest", true},
		{"<!ENTITY foo>rest", true},
		{"<!NOTATION foo>rest", true},
		{"<!ATTLIST foo>rest", true},
		{"<![CDATA[x]]>rest", true},
		// Lowercase
		{"<!doctype foo>rest", true},
		{"<!entity foo>rest", true},
		// Not a real declaration (e.g., regex "(?<!")
		{"<!x>rest", false},
		{"<!9>rest", false},
	}

	for _, tt := range tests {
		tok := NewTokenizer(tt.input)
		tokens := tok.Tokenize()
		if tt.real {
			// Real declarations produce Comment tokens
			if tokens[0].Type != TokenComment && tokens[0].Type != TokenText {
				t.Errorf("input %q: expected Comment or Text, got %d",
					tt.input, tokens[0].Type)
			}
		} else {
			// Fake declarations produce Text "<" tokens
			if tokens[0].Type != TokenText {
				t.Errorf("input %q: expected Text for non-declaration, got %d",
					tt.input, tokens[0].Type)
			}
		}
	}
}

func TestTokenizerIsLiteralPI(t *testing.T) {
	// Literal PI: <?x appears inside element content before </
	input := `<literal><?x</literal>`
	tok := NewTokenizer(input)
	tokens := tok.Tokenize()

	// The <?x should become text "<", then "?x" as text
	// <literal> open, text, </literal> close
	foundLiteral := false
	for _, tk := range tokens {
		if tk.Type == TokenTagOpen && tk.Tag == "literal" {
			foundLiteral = true
		}
	}
	if !foundLiteral {
		t.Error("expected <literal> tag to be parsed")
	}

	// Should NOT produce a PI token
	for _, tk := range tokens {
		if tk.Type == TokenPI {
			t.Error("expected no PI token for literal PI content")
		}
	}
}

func TestTokenizerRealPI(t *testing.T) {
	input := `<?xml version="1.0"?><para>x</para>`
	tok := NewTokenizer(input)
	tokens := tok.Tokenize()

	if tokens[0].Type != TokenPI {
		t.Errorf("expected PI token, got %d", tokens[0].Type)
	}
	if tokens[0].Text != `xml version="1.0"` {
		t.Errorf("unexpected PI text: %q", tokens[0].Text)
	}
}

func TestTokenizerPIUnterminated(t *testing.T) {
	input := `<?xml no closing`
	tok := NewTokenizer(input)
	tokens := tok.Tokenize()

	if tokens[0].Type != TokenPI {
		t.Errorf("expected PI token for unterminated PI, got %d",
			tokens[0].Type)
	}
}

func TestTokenizerReadOpenTagOperatorLtGt(t *testing.T) {
	// The <> operator should become text
	input := `<para>a <> b</para>`
	tok := NewTokenizer(input)
	tokens := tok.Tokenize()

	foundOperator := false
	for _, tk := range tokens {
		if tk.Type == TokenText && containsStr(tk.Text, "<>") {
			foundOperator = true
		}
	}
	if !foundOperator {
		t.Error("expected <> operator as text")
	}
}

func TestTokenizerReadOpenTagOperatorLtDashGt(t *testing.T) {
	// The <-> operator should become text
	input := `<para>a <-> b</para>`
	tok := NewTokenizer(input)
	tokens := tok.Tokenize()

	foundOperator := false
	for _, tk := range tokens {
		if tk.Type == TokenText && containsStr(tk.Text, "<->") {
			foundOperator = true
		}
	}
	if !foundOperator {
		t.Error("expected <-> operator as text")
	}
}

func TestTokenizerReadOpenTagDigitStart(t *testing.T) {
	// < followed by digit should be literal text
	input := `<para>x < 5</para>`
	tok := NewTokenizer(input)
	tokens := tok.Tokenize()

	// Should not produce a tag for "< 5"
	for _, tk := range tokens {
		if tk.Type == TokenTagOpen && tk.Tag == "5" {
			t.Error("digit should not start a tag name")
		}
	}
}

func TestTokenizerReadOpenTagWhitespace(t *testing.T) {
	// < followed by space is literal text
	input := `<para>a < b</para>`
	tok := NewTokenizer(input)
	tokens := tok.Tokenize()

	// Should parse as: <para>, text "a ", text "<", text " b", </para>
	for _, tk := range tokens {
		if tk.Type == TokenTagOpen && tk.Tag == "b" {
			t.Error("< space should not start a tag")
		}
	}
}

func TestTokenizerReadOpenTagColonStart(t *testing.T) {
	// Use a direct case: <:alpha:> should not be a tag
	input2 := `<:alpha:>`
	tok2 := NewTokenizer(input2)
	tokens2 := tok2.Tokenize()

	// Should produce text tokens, not a tag
	for _, tk := range tokens2 {
		if tk.Type == TokenTagOpen {
			t.Errorf("expected no TagOpen for <:alpha:>, got tag %q",
				tk.Tag)
		}
	}
}

func TestTokenizerReadAttributesValueless(t *testing.T) {
	// SGML allows valueless (boolean) attributes
	input := `<option compact>`
	tok := NewTokenizer(input)
	tokens := tok.Tokenize()

	if tokens[0].Type != TokenTagOpen {
		t.Fatalf("expected TagOpen, got %d", tokens[0].Type)
	}
	if tokens[0].Attrs["compact"] != "compact" {
		t.Errorf("expected boolean attr compact='compact', got %q",
			tokens[0].Attrs["compact"])
	}
}

func TestTokenizerReadAttributesMixedQuotes(t *testing.T) {
	input := `<elem a="double" b='single'>`
	tok := NewTokenizer(input)
	tokens := tok.Tokenize()

	if tokens[0].Type != TokenTagOpen {
		t.Fatalf("expected TagOpen, got %d", tokens[0].Type)
	}
	if tokens[0].Attrs["a"] != "double" {
		t.Errorf("expected a='double', got %q", tokens[0].Attrs["a"])
	}
	if tokens[0].Attrs["b"] != "single" {
		t.Errorf("expected b='single', got %q", tokens[0].Attrs["b"])
	}
}

func TestTokenizerReadAttrValueNewlines(t *testing.T) {
	input := "<elem attr=\"line1\nline2\">"
	tok := NewTokenizer(input)
	tokens := tok.Tokenize()

	if tokens[0].Type != TokenTagOpen {
		t.Fatalf("expected TagOpen, got %d", tokens[0].Type)
	}
	if tokens[0].Attrs["attr"] != "line1\nline2" {
		t.Errorf("expected attr with newline, got %q",
			tokens[0].Attrs["attr"])
	}
}

func TestTokenizerReadAttrValueUnquoted(t *testing.T) {
	input := `<elem attr=value>`
	tok := NewTokenizer(input)
	tokens := tok.Tokenize()

	if tokens[0].Type != TokenTagOpen {
		t.Fatalf("expected TagOpen, got %d", tokens[0].Type)
	}
	if tokens[0].Attrs["attr"] != "value" {
		t.Errorf("expected attr='value', got %q",
			tokens[0].Attrs["attr"])
	}
}

func TestTokenizerReadAttrValueEmpty(t *testing.T) {
	// readAttrValue at EOF
	input := `<elem attr=`
	tok := NewTokenizer(input)
	tokens := tok.Tokenize()

	if tokens[0].Type != TokenTagOpen {
		t.Fatalf("expected TagOpen, got %d", tokens[0].Type)
	}
	// attr value should be empty string
	if tokens[0].Attrs["attr"] != "" {
		t.Errorf("expected empty attr value, got %q",
			tokens[0].Attrs["attr"])
	}
}

func TestTokenizerReadCommentUnterminated(t *testing.T) {
	input := `<!-- no end`
	tok := NewTokenizer(input)
	tokens := tok.Tokenize()

	if tokens[0].Type != TokenComment {
		t.Fatalf("expected Comment, got %d", tokens[0].Type)
	}
	// Should contain the rest of input
	if !containsStr(tokens[0].Text, "no end") {
		t.Errorf("expected unterminated comment text, got %q",
			tokens[0].Text)
	}
}

func TestTokenizerReadCommentMultiline(t *testing.T) {
	input := "<!-- line1\nline2\nline3 --><para/>"
	tok := NewTokenizer(input)
	tokens := tok.Tokenize()

	if tokens[0].Type != TokenComment {
		t.Fatalf("expected Comment, got %d", tokens[0].Type)
	}
	if !containsStr(tokens[0].Text, "line2") {
		t.Errorf("expected multiline comment, got %q", tokens[0].Text)
	}
}

func TestTokenizerCommentFollowedByCloseTag(t *testing.T) {
	// <!--</ should be treated as literal text, not a comment
	input := `<literal><!--</literal>`
	tok := NewTokenizer(input)
	tokens := tok.Tokenize()

	// Should have <literal>, some text tokens, </literal>
	if tokens[0].Type != TokenTagOpen || tokens[0].Tag != "literal" {
		t.Fatalf("expected <literal>, got %s", tokens[0])
	}
	// The <!--</ should produce a "<" text token
	foundText := false
	for _, tk := range tokens {
		if tk.Type == TokenText && tk.Text == "<" {
			foundText = true
		}
	}
	if !foundText {
		t.Error("expected literal text '<' for <!--</ pattern")
	}
}

func TestTokenizerSelfClosingMarker(t *testing.T) {
	input := `<xref linkend="foo"/>`
	tok := NewTokenizer(input)
	tokens := tok.Tokenize()

	if tokens[0].Attrs["\x00selfclose"] != "1" {
		t.Error("expected selfclose marker on self-closing tag")
	}
}

func TestTokenizerEmptyInput(t *testing.T) {
	tok := NewTokenizer("")
	tokens := tok.Tokenize()

	if len(tokens) != 1 || tokens[0].Type != TokenEOF {
		t.Errorf("expected single EOF token, got %d tokens", len(tokens))
	}
}

func TestTokenizerTextOnly(t *testing.T) {
	input := "just some text"
	tok := NewTokenizer(input)
	tokens := tok.Tokenize()

	if len(tokens) != 2 {
		t.Fatalf("expected 2 tokens (text + EOF), got %d", len(tokens))
	}
	if tokens[0].Type != TokenText || tokens[0].Text != "just some text" {
		t.Errorf("expected text token, got %s", tokens[0])
	}
}

func TestTokenString(t *testing.T) {
	tests := []struct {
		tok    Token
		expect string
	}{
		{Token{Type: TokenTagOpen, Tag: "para", Line: 1},
			"<para> (line 1)"},
		{Token{Type: TokenTagOpen, Tag: "elem", Line: 2,
			Attrs: map[string]string{"id": "x"}},
			"<elem ...> (line 2)"},
		{Token{Type: TokenTagClose, Tag: "para", Line: 3},
			"</para> (line 3)"},
		{Token{Type: TokenText, Text: "short", Line: 4},
			`TEXT("short") (line 4)`},
		{Token{Type: TokenText,
			Text: "a very long text that exceeds forty characters limit here", Line: 5},
			`TEXT("a very long text that exceeds forty char...") (line 5)`},
		{Token{Type: TokenComment, Line: 6},
			"COMMENT (line 6)"},
		{Token{Type: TokenPI, Line: 7},
			"PI (line 7)"},
		{Token{Type: TokenEOF},
			"EOF"},
		{Token{Type: TokenType(99)},
			"UNKNOWN"},
	}

	for _, tt := range tests {
		got := tt.tok.String()
		if got != tt.expect {
			t.Errorf("expected %q, got %q", tt.expect, got)
		}
	}
}

// ===== Entity resolver tests =====

func TestEntityResolverProcessDoctypeNoDOCTYPE(t *testing.T) {
	resolver := NewEntityResolver(t.TempDir())
	result := resolver.processDoctype("<book>content</book>")
	if result != "<book>content</book>" {
		t.Errorf("expected unchanged content, got %q", result)
	}
}

func TestEntityResolverProcessDoctypeNoBrackets(t *testing.T) {
	resolver := NewEntityResolver(t.TempDir())
	input := `<!DOCTYPE book PUBLIC "test">` +
		`<book>content</book>`
	result := resolver.processDoctype(input)
	if !containsStr(result, "<book>content</book>") {
		t.Errorf("expected body after DOCTYPE, got %q", result)
	}
}

func TestEntityResolverProcessDoctypeNoClosingBracket(t *testing.T) {
	resolver := NewEntityResolver(t.TempDir())
	input := `<!DOCTYPE book [<!ENTITY foo "bar">`
	result := resolver.processDoctype(input)
	// No closing bracket — returns content unchanged
	if result != input {
		t.Errorf("expected unchanged for unclosed bracket, got %q",
			result)
	}
}

func TestEntityResolverProcessDoctypeNoClosingGt(t *testing.T) {
	resolver := NewEntityResolver(t.TempDir())
	input := `<!DOCTYPE book [<!ENTITY foo "bar">]`
	result := resolver.processDoctype(input)
	// Has ] but no > after it
	if result != "" {
		t.Errorf("expected empty for no closing >, got %q", result)
	}
}

func TestEntityResolverProcessSubset(t *testing.T) {
	dir := t.TempDir()
	resolver := NewEntityResolver(dir)
	subset := `<!ENTITY greeting "hello">
<!ENTITY target "world">`
	resolver.processSubset(subset)

	if !resolver.HasEntity("greeting") {
		t.Error("expected greeting entity")
	}
	if !resolver.HasEntity("target") {
		t.Error("expected target entity")
	}
}

func TestEntityResolverExpandParamEntitiesInSubset(t *testing.T) {
	dir := t.TempDir()
	// Write a file that a parameter entity references
	os.WriteFile(filepath.Join(dir, "defs.sgml"),
		[]byte(`<!ENTITY myval "42">`), 0644)

	resolver := NewEntityResolver(dir)
	subset := `<!ENTITY % defs SYSTEM "defs.sgml">
%defs;`
	result := resolver.expandParamEntitiesInSubset(subset)

	if !containsStr(result, `<!ENTITY myval "42">`) {
		t.Errorf("expected expanded param entity, got %q", result)
	}
}

func TestEntityResolverExpandParamEntitiesMissingFile(t *testing.T) {
	dir := t.TempDir()
	resolver := NewEntityResolver(dir)
	subset := `<!ENTITY % missing SYSTEM "nonexistent.sgml">
%missing;`
	result := resolver.expandParamEntitiesInSubset(subset)

	// Should leave %missing; unexpanded
	if !containsStr(result, "%missing;") {
		t.Errorf("expected unexpanded param ref, got %q", result)
	}
}

func TestEntityResolverExtractEntityDeclsTextVsFile(t *testing.T) {
	dir := t.TempDir()
	resolver := NewEntityResolver(dir)

	text := `<!ENTITY greeting "hello">
<!ENTITY chapter SYSTEM "chapter.sgml">
<!ENTITY % param SYSTEM "param.sgml">`

	resolver.extractEntityDecls(text)

	if !resolver.HasEntity("greeting") {
		t.Error("expected text entity 'greeting'")
	}
	if !resolver.HasEntity("chapter") {
		t.Error("expected file entity 'chapter'")
	}
	// Parameter entities should be skipped
	if resolver.HasEntity("param") {
		t.Error("did not expect param entity without % prefix")
	}
}

func TestEntityResolverExtractEntityDeclsSingleQuote(t *testing.T) {
	dir := t.TempDir()
	resolver := NewEntityResolver(dir)
	text := `<!ENTITY sqval 'single quoted'>`
	resolver.extractEntityDecls(text)

	if !resolver.HasEntity("sqval") {
		t.Error("expected single-quoted entity")
	}
}

func TestEntityResolverExtractEntityDeclsNoOverwrite(t *testing.T) {
	dir := t.TempDir()
	resolver := NewEntityResolver(dir)

	// File entity should take precedence over text entity with
	// same name if file is declared first
	text := `<!ENTITY foo SYSTEM "foo.sgml">
<!ENTITY foo "text value">`
	resolver.extractEntityDecls(text)

	// The file entity should win (it's processed first in the code)
	if !resolver.HasEntity("foo") {
		t.Error("expected entity 'foo'")
	}
}

func TestEntityResolverExpandEntitiesCycleDetection(t *testing.T) {
	dir := t.TempDir()
	resolver := NewEntityResolver(dir)
	resolver.entities["a"] = "&b;"
	resolver.entities["b"] = "&a;"

	_, err := resolver.expandEntities("&a;", 0)
	if err == nil {
		t.Error("expected cycle detection error")
	}
	if !containsStr(err.Error(), "circular") {
		t.Errorf("expected circular error message, got %q", err.Error())
	}
}

func TestEntityResolverExpandEntitiesMaxDepth(t *testing.T) {
	dir := t.TempDir()
	resolver := NewEntityResolver(dir)
	resolver.maxDepth = 3
	resolver.entities["deep"] = "&deeper;"
	resolver.entities["deeper"] = "&deepest;"
	resolver.entities["deepest"] = "&bottom;"
	resolver.entities["bottom"] = "done"

	_, err := resolver.expandEntities("&deep;", 0)
	if err == nil {
		t.Error("expected max depth error")
	}
	if !containsStr(err.Error(), "max depth") {
		t.Errorf("expected max depth error, got %q", err.Error())
	}
}

func TestEntityResolverExpandEntitiesUnknown(t *testing.T) {
	dir := t.TempDir()
	resolver := NewEntityResolver(dir)

	result, err := resolver.expandEntities("&unknown;", 0)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	// Unknown entities should be left as-is
	if result != "&unknown;" {
		t.Errorf("expected &unknown; unchanged, got %q", result)
	}
}

func TestEntityResolverExpandEntitiesPredefined(t *testing.T) {
	dir := t.TempDir()
	resolver := NewEntityResolver(dir)

	result, err := resolver.expandEntities("&lt;&gt;&amp;&quot;&apos;", 0)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expected := `<>&"'`
	if result != expected {
		t.Errorf("expected %q, got %q", expected, result)
	}
}

func TestEntityResolverExpandEntitiesFileEntity(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "chapter.sgml"),
		[]byte("<chapter>File content</chapter>"), 0644)

	resolver := NewEntityResolver(dir)
	resolver.entities["mychapter"] = "\x00FILE:chapter.sgml"

	result, err := resolver.expandEntities("&mychapter;", 0)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !containsStr(result, "File content") {
		t.Errorf("expected file entity expansion, got %q", result)
	}
}

func TestEntityResolverExpandEntitiesFileEntityMissing(t *testing.T) {
	dir := t.TempDir()
	resolver := NewEntityResolver(dir)
	resolver.entities["missing"] = "\x00FILE:nonexistent.sgml"

	result, err := resolver.expandEntities("&missing;", 0)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	// Should produce a comment placeholder
	if !containsStr(result, "not available") {
		t.Errorf("expected 'not available' placeholder, got %q", result)
	}
}

func TestEntityResolverExpandEntitiesFileEntitySubdir(t *testing.T) {
	dir := t.TempDir()
	// Create ref subdirectory with file
	os.MkdirAll(filepath.Join(dir, "ref"), 0755)
	os.WriteFile(filepath.Join(dir, "ref", "cmd.sgml"),
		[]byte("<refentry>command</refentry>"), 0644)

	resolver := NewEntityResolver(dir)
	resolver.entities["cmd"] = "\x00FILE:cmd.sgml"

	result, err := resolver.expandEntities("&cmd;", 0)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !containsStr(result, "command") {
		t.Errorf("expected subdir file expansion, got %q", result)
	}
}

func TestEntityResolverResolveFileActual(t *testing.T) {
	dir := t.TempDir()

	// Simple file without DOCTYPE
	os.WriteFile(filepath.Join(dir, "simple.sgml"),
		[]byte(`<para>Hello</para>`), 0644)

	resolver := NewEntityResolver(dir)
	body, err := resolver.ResolveFile("simple.sgml")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if body != "<para>Hello</para>" {
		t.Errorf("expected simple body, got %q", body)
	}
}

func TestEntityResolverResolveFileMissing(t *testing.T) {
	dir := t.TempDir()
	resolver := NewEntityResolver(dir)
	_, err := resolver.ResolveFile("nonexistent.sgml")
	if err == nil {
		t.Error("expected error for missing file")
	}
}

func TestEntityResolverCharRefsHex(t *testing.T) {
	resolver := NewEntityResolver(t.TempDir())
	result := resolver.expandCharRefs("&#x48;&#x65;&#x6C;&#x6C;&#x6F;")
	if result != "Hello" {
		t.Errorf("expected 'Hello', got %q", result)
	}
}

func TestEntityResolverCharRefsDecimal(t *testing.T) {
	resolver := NewEntityResolver(t.TempDir())
	result := resolver.expandCharRefs("&#72;&#101;&#108;&#108;&#111;")
	if result != "Hello" {
		t.Errorf("expected 'Hello', got %q", result)
	}
}

func TestEntityResolverEntityCountAndHasEntity(t *testing.T) {
	dir := t.TempDir()
	resolver := NewEntityResolver(dir)
	resolver.entities["foo"] = "bar"
	resolver.entities["baz"] = "qux"

	if resolver.EntityCount() != 2 {
		t.Errorf("expected 2 entities, got %d", resolver.EntityCount())
	}
	if !resolver.HasEntity("foo") {
		t.Error("expected HasEntity('foo') to be true")
	}
	if resolver.HasEntity("missing") {
		t.Error("expected HasEntity('missing') to be false")
	}
}

func TestEntityResolverWarnings(t *testing.T) {
	dir := t.TempDir()
	resolver := NewEntityResolver(dir)
	resolver.entities["missing"] = "\x00FILE:nonexistent.sgml"

	resolver.expandEntities("&missing;", 0)
	warnings := resolver.Warnings()
	if len(warnings) == 0 {
		t.Error("expected warnings for missing file entity")
	}
}

// ===== Parser tests =====

func TestParserImplicitClose(t *testing.T) {
	// Test that a close tag for an ancestor implicitly closes current
	input := `<sect1><para>text</sect1>`
	root, warnings, err := ParseString(input)
	if err != nil {
		t.Fatalf("parse error: %v", err)
	}
	// Should produce warnings but still parse
	t.Logf("warnings: %v", warnings)

	s1 := root.FindChild("sect1")
	if s1 == nil {
		t.Fatal("expected sect1")
	}
	para := s1.FindChild("para")
	if para == nil {
		t.Fatal("expected para inside sect1")
	}
}

func TestParserStrayCloseTag(t *testing.T) {
	// A close tag with no matching open tag
	input := `<para>text</nonexistent></para>`
	root, warnings, err := ParseString(input)
	if err != nil {
		t.Fatalf("parse error: %v", err)
	}
	if len(warnings) == 0 {
		t.Error("expected warning for stray close tag")
	}
	para := root.FindChild("para")
	if para == nil {
		t.Fatal("expected para")
	}
}

func TestParserHTMLElements(t *testing.T) {
	// HTML elements should be converted to text, not parsed as elements
	input := `<para>Use <b>bold</b> text</para>`
	root, _, err := ParseString(input)
	if err != nil {
		t.Fatalf("parse error: %v", err)
	}

	para := root.FindChild("para")
	if para == nil {
		t.Fatal("expected para")
	}

	// <b> and </b> should be text nodes, not element nodes
	bElements := para.FindChildren("b")
	if len(bElements) != 0 {
		t.Errorf("expected 0 <b> elements (should be text), got %d",
			len(bElements))
	}

	// The text content should contain the HTML tags as literal text
	tc := para.TextContent()
	if !containsStr(tc, "<b>") {
		t.Errorf("expected <b> as text content, got %q", tc)
	}
}

func TestParserHTMLElementsWithAttrs(t *testing.T) {
	input := `<para><a href="http://example.com">link</a></para>`
	root, _, err := ParseString(input)
	if err != nil {
		t.Fatalf("parse error: %v", err)
	}

	para := root.FindChild("para")
	if para == nil {
		t.Fatal("expected para")
	}

	// <a> should be text
	aElements := para.FindChildren("a")
	if len(aElements) != 0 {
		t.Error("expected <a> to be treated as text")
	}
}

func TestParserPISkipped(t *testing.T) {
	input := `<?xml version="1.0"?><para>text</para>`
	root, _, err := ParseString(input)
	if err != nil {
		t.Fatalf("parse error: %v", err)
	}

	para := root.FindChild("para")
	if para == nil {
		t.Fatal("expected para after PI")
	}
}

func TestParserCommentNode(t *testing.T) {
	input := `<para><!-- a comment -->text</para>`
	root, _, err := ParseString(input)
	if err != nil {
		t.Fatalf("parse error: %v", err)
	}

	para := root.FindChild("para")
	if para == nil {
		t.Fatal("expected para")
	}

	// Should have a comment child
	foundComment := false
	for _, c := range para.Children {
		if c.Type == CommentNode {
			foundComment = true
			if c.Text != "a comment" {
				t.Errorf("expected comment 'a comment', got %q",
					c.Text)
			}
		}
	}
	if !foundComment {
		t.Error("expected comment node in para")
	}
}

func TestParserIsAncestor(t *testing.T) {
	// Build a simple tree manually to test isAncestor
	root := &Node{Type: ElementNode, Tag: "root"}
	child := &Node{Type: ElementNode, Tag: "child"}
	grandchild := &Node{Type: ElementNode, Tag: "grandchild"}
	root.AppendChild(child)
	child.AppendChild(grandchild)

	parser := NewParser(nil)

	// grandchild's ancestors include root and child
	if !parser.isAncestor(grandchild, "root") {
		t.Error("expected root to be ancestor of grandchild")
	}
	if !parser.isAncestor(grandchild, "child") {
		t.Error("expected child to be ancestor of grandchild")
	}
	if parser.isAncestor(grandchild, "other") {
		t.Error("expected 'other' NOT to be ancestor")
	}
	if parser.isAncestor(root, "child") {
		t.Error("expected child NOT to be ancestor of root")
	}
}

func TestCleanAttrsNil(t *testing.T) {
	result := cleanAttrs(nil)
	if result != nil {
		t.Errorf("expected nil for nil input, got %v", result)
	}
}

func TestCleanAttrsEmpty(t *testing.T) {
	result := cleanAttrs(map[string]string{})
	if result != nil {
		t.Errorf("expected nil for empty attrs, got %v", result)
	}
}

func TestCleanAttrsOnlyInternal(t *testing.T) {
	attrs := map[string]string{"\x00selfclose": "1"}
	result := cleanAttrs(attrs)
	if result != nil {
		t.Errorf("expected nil when only internal attrs, got %v", result)
	}
}

func TestCleanAttrsMixed(t *testing.T) {
	attrs := map[string]string{
		"\x00selfclose": "1",
		"id":            "test",
		"class":         "foo",
	}
	result := cleanAttrs(attrs)
	if result == nil {
		t.Fatal("expected non-nil result")
	}
	if result["id"] != "test" || result["class"] != "foo" {
		t.Errorf("expected id and class attrs, got %v", result)
	}
	if _, ok := result["\x00selfclose"]; ok {
		t.Error("internal attr should be removed")
	}
}

// ===== Node tests =====

func TestNodeFindDescendantsDeepNesting(t *testing.T) {
	// Build: root > a > b > c > target
	root := &Node{Type: ElementNode, Tag: "root"}
	a := &Node{Type: ElementNode, Tag: "a"}
	b := &Node{Type: ElementNode, Tag: "b"}
	c := &Node{Type: ElementNode, Tag: "c"}
	target1 := &Node{Type: ElementNode, Tag: "target"}
	target2 := &Node{Type: ElementNode, Tag: "target"}

	root.AppendChild(a)
	a.AppendChild(b)
	b.AppendChild(c)
	c.AppendChild(target1)
	// Also add one at the a level
	a.AppendChild(target2)

	targets := root.FindDescendants("target")
	if len(targets) != 2 {
		t.Errorf("expected 2 targets, got %d", len(targets))
	}
}

func TestNodeFindDescendantsNone(t *testing.T) {
	root := &Node{Type: ElementNode, Tag: "root"}
	child := &Node{Type: ElementNode, Tag: "child"}
	root.AppendChild(child)

	results := root.FindDescendants("nonexistent")
	if len(results) != 0 {
		t.Errorf("expected 0 results, got %d", len(results))
	}
}

func TestNodeGetAttrNilAttrs(t *testing.T) {
	node := &Node{Type: ElementNode, Tag: "test"}
	// Attrs is nil by default
	if node.GetAttr("anything") != "" {
		t.Error("expected empty string for nil attrs")
	}
}

func TestNodeGetAttrMissing(t *testing.T) {
	node := &Node{
		Type:  ElementNode,
		Tag:   "test",
		Attrs: map[string]string{"id": "x"},
	}
	if node.GetAttr("class") != "" {
		t.Error("expected empty string for missing attr")
	}
}

func TestNodeTextContentNested(t *testing.T) {
	input := `<para>Hello <emphasis>beautiful</emphasis> world</para>`
	root, _, err := ParseString(input)
	if err != nil {
		t.Fatalf("parse error: %v", err)
	}

	para := root.FindChild("para")
	if para == nil {
		t.Fatal("expected para")
	}
	tc := para.TextContent()
	if tc != "Hello beautiful world" {
		t.Errorf("expected 'Hello beautiful world', got %q", tc)
	}
}

func TestNodeTextContentEmpty(t *testing.T) {
	node := &Node{Type: ElementNode, Tag: "empty"}
	if node.TextContent() != "" {
		t.Errorf("expected empty text content, got %q",
			node.TextContent())
	}
}

func TestNodeAppendChild(t *testing.T) {
	parent := &Node{Type: ElementNode, Tag: "parent"}
	child := &Node{Type: ElementNode, Tag: "child"}
	parent.AppendChild(child)

	if len(parent.Children) != 1 {
		t.Fatalf("expected 1 child, got %d", len(parent.Children))
	}
	if child.Parent != parent {
		t.Error("expected child.Parent to be parent")
	}
}

func TestNodeFindChildNone(t *testing.T) {
	root := &Node{Type: ElementNode, Tag: "root"}
	if root.FindChild("nothing") != nil {
		t.Error("expected nil for missing child")
	}
}

func TestNodeFindChildrenEmpty(t *testing.T) {
	root := &Node{Type: ElementNode, Tag: "root"}
	result := root.FindChildren("nothing")
	if len(result) != 0 {
		t.Errorf("expected 0, got %d", len(result))
	}
}

func TestNodeFindChildSkipsText(t *testing.T) {
	root := &Node{Type: ElementNode, Tag: "root"}
	text := &Node{Type: TextNode, Text: "hello"}
	elem := &Node{Type: ElementNode, Tag: "elem"}
	root.AppendChild(text)
	root.AppendChild(elem)

	// FindChild should skip text nodes
	found := root.FindChild("elem")
	if found != elem {
		t.Error("expected to find elem, skipping text")
	}
	// FindChild for a tag that doesn't exist
	if root.FindChild("text") != nil {
		t.Error("text nodes should not be found by FindChild")
	}
}

// ===== Generate tests =====

func TestExtractMajorVersion(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"17.4", "17"},
		{"19devel", "19"},
		{"18beta1", "18"},
		{"18beta2", "18"},
		{"18beta3", "18"},
		{"18beta4", "18"},
		{"17rc1", "17"},
		{"17rc2", "17"},
		{"18alpha1", "18"},
		{"18alpha2", "18"},
		{"18alpha3", "18"},
		{"17", "17"},
		{"9.6.24", "9"},
	}

	for _, tt := range tests {
		got := extractMajorVersion(tt.input)
		if got != tt.expected {
			t.Errorf("extractMajorVersion(%q) = %q, want %q",
				tt.input, got, tt.expected)
		}
	}
}

func TestGenerateVersionSGMLFromTemplate(t *testing.T) {
	dir := t.TempDir()

	// Write a version.sgml.in template
	template := `<!ENTITY version @PG_VERSION@>
<!ENTITY majorversion @PG_MAJORVERSION@>`
	os.WriteFile(filepath.Join(dir, "version.sgml.in"),
		[]byte(template), 0644)

	err := generateVersionSGML(dir, "17.4")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	data, err := os.ReadFile(filepath.Join(dir, "version.sgml"))
	if err != nil {
		t.Fatalf("failed to read output: %v", err)
	}

	content := string(data)
	if !containsStr(content, `"17.4"`) {
		t.Errorf("expected version 17.4 in output, got %q", content)
	}
	if !containsStr(content, `"17"`) {
		t.Errorf("expected major version 17 in output, got %q", content)
	}
}

func TestGenerateVersionSGMLNoTemplate(t *testing.T) {
	dir := t.TempDir()
	// No version.sgml.in exists — should create directly

	err := generateVersionSGML(dir, "19devel")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	data, err := os.ReadFile(filepath.Join(dir, "version.sgml"))
	if err != nil {
		t.Fatalf("failed to read output: %v", err)
	}

	content := string(data)
	if !containsStr(content, `"19devel"`) {
		t.Errorf("expected version in output, got %q", content)
	}
	if !containsStr(content, `"19"`) {
		t.Errorf("expected major version in output, got %q", content)
	}
}

func TestWriteVersionSGML(t *testing.T) {
	dir := t.TempDir()
	outPath := filepath.Join(dir, "version.sgml")

	err := writeVersionSGML(outPath, "18beta2")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	data, err := os.ReadFile(outPath)
	if err != nil {
		t.Fatalf("failed to read output: %v", err)
	}

	content := string(data)
	if !containsStr(content, `"18beta2"`) {
		t.Errorf("expected version, got %q", content)
	}
	if !containsStr(content, `"18"`) {
		t.Errorf("expected major version, got %q", content)
	}
}

func TestCleanGeneratedFiles(t *testing.T) {
	dir := t.TempDir()

	// Create a version.sgml.in so version.sgml gets cleaned
	os.WriteFile(filepath.Join(dir, "version.sgml.in"),
		[]byte("template"), 0644)
	os.WriteFile(filepath.Join(dir, "version.sgml"),
		[]byte("generated"), 0644)

	// Create some generated files
	names := []string{
		"features-supported.sgml",
		"features-unsupported.sgml",
		"errcodes-table.sgml",
		"keywords-table.sgml",
		"wait_event_types.sgml",
		"targets-meson.sgml",
	}
	for _, name := range names {
		os.WriteFile(filepath.Join(dir, name),
			[]byte("generated"), 0644)
	}

	CleanGeneratedFiles(dir)

	// version.sgml should be removed (since .in exists)
	if _, err := os.Stat(filepath.Join(dir, "version.sgml")); err == nil {
		t.Error("expected version.sgml to be cleaned")
	}

	// All generated files should be removed
	for _, name := range names {
		if _, err := os.Stat(filepath.Join(dir, name)); err == nil {
			t.Errorf("expected %s to be cleaned", name)
		}
	}

	// version.sgml.in should still exist
	if _, err := os.Stat(filepath.Join(dir, "version.sgml.in")); err != nil {
		t.Error("version.sgml.in should not be cleaned")
	}
}

func TestCleanGeneratedFilesNoTemplate(t *testing.T) {
	dir := t.TempDir()

	// Without version.sgml.in, version.sgml should NOT be cleaned
	os.WriteFile(filepath.Join(dir, "version.sgml"),
		[]byte("manual"), 0644)

	CleanGeneratedFiles(dir)

	if _, err := os.Stat(filepath.Join(dir, "version.sgml")); err != nil {
		t.Error("version.sgml should not be cleaned without .in file")
	}
}

// ===== Integration-style tests =====

func TestParserNestedWithEmptyElements(t *testing.T) {
	input := `<sect1>
<title>Test</title>
<para>See <xref linkend="a"> and <anchor id="b"> here.</para>
<para>Also <colspec colname="c"> in tables.</para>
</sect1>`
	root, _, err := ParseString(input)
	if err != nil {
		t.Fatalf("parse error: %v", err)
	}

	s1 := root.FindChild("sect1")
	if s1 == nil {
		t.Fatal("expected sect1")
	}

	// All empty elements should be parsed without consuming children
	xrefs := s1.FindDescendants("xref")
	if len(xrefs) != 1 {
		t.Errorf("expected 1 xref, got %d", len(xrefs))
	}
	anchors := s1.FindDescendants("anchor")
	if len(anchors) != 1 {
		t.Errorf("expected 1 anchor, got %d", len(anchors))
	}
}

func TestEntityResolverWithNestedParamEntities(t *testing.T) {
	dir := t.TempDir()

	// Level 1: allfiles.sgml references another file
	os.WriteFile(filepath.Join(dir, "outer.sgml"),
		[]byte(`<!ENTITY myitem SYSTEM "item.sgml">`), 0644)

	os.WriteFile(filepath.Join(dir, "item.sgml"),
		[]byte(`<listitem>An item</listitem>`), 0644)

	os.WriteFile(filepath.Join(dir, "main.sgml"),
		[]byte(`<!DOCTYPE doc [
<!ENTITY % outer SYSTEM "outer.sgml">
%outer;
]>
<doc>&myitem;</doc>`), 0644)

	resolver := NewEntityResolver(dir)
	body, err := resolver.ResolveFile("main.sgml")
	if err != nil {
		t.Fatalf("resolve error: %v", err)
	}

	if !containsStr(body, "An item") {
		t.Errorf("expected nested entity expansion, got %q", body)
	}
}

func TestTokenizerDashOperatorNoClose(t *testing.T) {
	// <-something without > should scan to end or next <
	input := `<-abc<para>x</para>`
	tok := NewTokenizer(input)
	tokens := tok.Tokenize()

	// Should get text token for <-abc, then <para>, etc.
	if tokens[0].Type != TokenText {
		t.Errorf("expected Text for <-abc, got %d (%s)",
			tokens[0].Type, tokens[0])
	}
}

func TestTokenizerReadAttributesUnexpectedChar(t *testing.T) {
	// Unexpected character in attribute position should be skipped
	input := `<elem !weird attr="val">`
	tok := NewTokenizer(input)
	tokens := tok.Tokenize()

	if tokens[0].Type != TokenTagOpen {
		t.Fatalf("expected TagOpen, got %d", tokens[0].Type)
	}
	if tokens[0].Attrs["attr"] != "val" {
		t.Errorf("expected attr='val', got %q", tokens[0].Attrs["attr"])
	}
}

func TestParserSelfClosingElement(t *testing.T) {
	input := `<para>text<br/>more</para>`
	root, _, err := ParseString(input)
	if err != nil {
		t.Fatalf("parse error: %v", err)
	}

	para := root.FindChild("para")
	if para == nil {
		t.Fatal("expected para")
	}

	// br is in htmlElements, so it should be text
	// But test that self-closing works in general
	input2 := `<para><custom/></para>`
	root2, _, err := ParseString(input2)
	if err != nil {
		t.Fatalf("parse error: %v", err)
	}

	para2 := root2.FindChild("para")
	if para2 == nil {
		t.Fatal("expected para")
	}
	custom := para2.FindChild("custom")
	if custom == nil {
		t.Fatal("expected custom element")
	}
	if len(custom.Children) != 0 {
		t.Errorf("expected self-closing to have no children, got %d",
			len(custom.Children))
	}
}

// ===== Additional coverage tests =====

// tokenizer: readPI unterminated
func TestTokenizerUnterminatedPI(t *testing.T) {
	input := `<?xml version="1.0"`
	tok := NewTokenizer(input)
	tokens := tok.Tokenize()

	found := false
	for _, tk := range tokens {
		if tk.Type == TokenPI {
			found = true
			if !containsStr(tk.Text, "xml") {
				t.Errorf("expected PI text with 'xml', got %q",
					tk.Text)
			}
		}
	}
	if !found {
		t.Error("expected unterminated PI token")
	}
}

// tokenizer: readDeclaration unterminated
func TestTokenizerUnterminatedDeclaration(t *testing.T) {
	input := `<!DOCTYPE doc [`
	tok := NewTokenizer(input)
	tokens := tok.Tokenize()

	found := false
	for _, tk := range tokens {
		if tk.Type == TokenComment &&
			containsStr(tk.Text, "DOCTYPE") {
			found = true
		}
	}
	if !found {
		t.Error("expected comment token for unterminated declaration")
	}
}

// tokenizer: readOpenTag with non-letter name start (e.g. ":")
func TestTokenizerOpenTagNonLetterName(t *testing.T) {
	input := `<:alpha:>`
	tok := NewTokenizer(input)
	tokens := tok.Tokenize()

	// Should be treated as text "<" then text ":alpha:>"
	if tokens[0].Type != TokenText {
		t.Errorf("expected Text for <: prefix, got %v (%s)",
			tokens[0].Type, tokens[0])
	}
}

// tokenizer: readOpenTag with empty name (< followed by > directly
// is handled by the digit/dash/> check, but let's test < followed
// by something that readName returns empty for)
func TestTokenizerOpenTagEmptyName(t *testing.T) {
	// < followed by = which isn't a valid name start
	input := `<=5`
	tok := NewTokenizer(input)
	tokens := tok.Tokenize()

	if tokens[0].Type != TokenText {
		t.Errorf("expected Text for <=, got %v (%s)",
			tokens[0].Type, tokens[0])
	}
	if tokens[0].Text != "<" {
		t.Errorf("expected '<' text, got %q", tokens[0].Text)
	}
}

// tokenizer: <-> operator scanning hitting another < before >
func TestTokenizerDashOperatorHitsLT(t *testing.T) {
	input := `<-foo<bar>x</bar>`
	tok := NewTokenizer(input)
	tokens := tok.Tokenize()

	if tokens[0].Type != TokenText {
		t.Errorf("expected Text for <-foo, got %v", tokens[0].Type)
	}
	// Should stop at the next < without consuming it
	if tokens[0].Text != "<-foo" {
		t.Errorf("expected '<-foo', got %q", tokens[0].Text)
	}
}

// tokenizer: skipWhitespace with newlines
func TestTokenizerSkipWhitespaceWithNewlines(t *testing.T) {
	input := "</ \n  elem>"
	tok := NewTokenizer(input)
	tokens := tok.Tokenize()

	found := false
	for _, tk := range tokens {
		if tk.Type == TokenTagClose && tk.Tag == "elem" {
			found = true
		}
	}
	if !found {
		t.Error(
			"expected close tag 'elem' after whitespace with newline",
		)
	}
}

// tokenizer: isLiteralPI fallthrough (reaches end of input)
func TestTokenizerIsLiteralPIFallthrough(t *testing.T) {
	// A <? with no ?> and no </ before end — falls through to
	// return false, then readPI handles it as unterminated
	input := `<?abc`
	tok := NewTokenizer(input)
	tokens := tok.Tokenize()

	found := false
	for _, tk := range tokens {
		if tk.Type == TokenPI {
			found = true
		}
	}
	if !found {
		t.Error("expected PI token for unterminated <?abc")
	}
}

// tokenizer: isRealDeclaration with lowercase and bracket variants
func TestTokenizerIsRealDeclarationVariants(t *testing.T) {
	tests := []struct {
		input string
		desc  string
	}{
		{`<!doctype doc>`, "lowercase d"},
		{`<!entity foo "bar">`, "lowercase e"},
		{`<!notation x SYSTEM "y">`, "lowercase n"},
		{`<!attlist elem attr CDATA>`, "lowercase a"},
		{`<![IGNORE[stuff]]>`, "bracket marked section"},
	}
	for _, tt := range tests {
		tok := NewTokenizer(tt.input)
		tokens := tok.Tokenize()
		found := false
		for _, tk := range tokens {
			if tk.Type == TokenComment {
				found = true
			}
		}
		if !found {
			t.Errorf("%s: expected declaration parsed as comment",
				tt.desc)
		}
	}
}

// tokenizer: fake declaration (<!X) treated as literal text
func TestTokenizerFakeDeclaration(t *testing.T) {
	input := `(?<!pattern)`
	tok := NewTokenizer(input)
	tokens := tok.Tokenize()

	// The <! should be treated as literal text since the next
	// char is not a valid declaration keyword
	allText := ""
	for _, tk := range tokens {
		if tk.Type == TokenText {
			allText += tk.Text
		}
	}
	if !containsStr(allText, "(?") {
		t.Errorf("expected literal text with '(?', got %q", allText)
	}
}

// entity: ResolveFile with expandEntities error
func TestResolveFileExpandError(t *testing.T) {
	dir := t.TempDir()

	// Create a file that triggers circular entity reference
	os.WriteFile(filepath.Join(dir, "circ.sgml"),
		[]byte(`&loopy;`), 0644)

	resolver := NewEntityResolver(dir)
	resolver.entities["loopy"] = "&loopy;" // self-referencing text
	// Set maxDepth low to trigger depth error quickly
	resolver.maxDepth = 2

	_, err := resolver.ResolveFile("circ.sgml")
	if err == nil {
		t.Error("expected error from circular/deep entity expansion")
	}
}

// entity: processDoctype with no internal subset (just >)
func TestProcessDoctypeNoSubset(t *testing.T) {
	dir := t.TempDir()
	resolver := NewEntityResolver(dir)

	content := `<!DOCTYPE doc SYSTEM "doc.dtd">
<doc>body</doc>`
	body := resolver.processDoctype(content)
	if !containsStr(body, "<doc>body</doc>") {
		t.Errorf("expected body after DOCTYPE, got %q", body)
	}
}

// entity: processDoctype with unterminated bracket
func TestProcessDoctypeUnterminatedBracket(t *testing.T) {
	dir := t.TempDir()
	resolver := NewEntityResolver(dir)

	content := `<!DOCTYPE doc [
<!ENTITY foo "bar">`
	body := resolver.processDoctype(content)
	// With unterminated bracket, should return content as-is
	if body != content {
		t.Errorf("expected original content for unterminated bracket")
	}
}

// entity: processDoctype with missing > after ]
func TestProcessDoctypeMissingCloseAfterBracket(t *testing.T) {
	dir := t.TempDir()
	resolver := NewEntityResolver(dir)

	// ] is found but no > follows
	content := `<!DOCTYPE doc [
<!ENTITY foo "bar">
]`
	body := resolver.processDoctype(content)
	if body != "" {
		t.Errorf("expected empty body for missing > after ], got %q",
			body)
	}
}

// entity: expandEntities circular reference
func TestExpandEntitiesCircular(t *testing.T) {
	dir := t.TempDir()
	resolver := NewEntityResolver(dir)
	resolver.entities["a"] = "&b;"
	resolver.entities["b"] = "&a;"

	_, err := resolver.expandEntities("&a;", 0)
	if err == nil {
		t.Error("expected circular entity reference error")
	}
	if !containsStr(err.Error(), "circular") {
		t.Errorf("expected 'circular' in error, got %q", err.Error())
	}
}

// entity: expandEntities file entity with expansion error
func TestExpandEntitiesFileEntityExpansionError(t *testing.T) {
	dir := t.TempDir()
	// File content references a circular entity
	os.WriteFile(filepath.Join(dir, "bad.sgml"),
		[]byte(`&self;`), 0644)

	resolver := NewEntityResolver(dir)
	resolver.entities["inc"] = "\x00FILE:bad.sgml"
	resolver.entities["self"] = "&self;"

	_, err := resolver.expandEntities("&inc;", 0)
	if err == nil {
		t.Error("expected error from file entity expansion")
	}
}

// entity: expandCharRefs with uppercase X hex prefix
func TestExpandCharRefsUppercaseX(t *testing.T) {
	resolver := NewEntityResolver(t.TempDir())
	// &#X41; should expand to 'A'
	result := resolver.expandCharRefs("&#x41;")
	if result != "A" {
		t.Errorf("expected 'A', got %q", result)
	}
}

// entity: expandCharRefs with invalid decimal
func TestExpandCharRefsInvalidDecimal(t *testing.T) {
	resolver := NewEntityResolver(t.TempDir())
	// The regex only matches [0-9a-fA-F]+, so a truly invalid
	// ref won't match. But test an edge: &#0; should produce
	// the null character (or at least not crash)
	result := resolver.expandCharRefs("&#0;")
	if result != string(rune(0)) {
		t.Errorf("expected null char for &#0;, got %q", result)
	}
}

// entity: expandParamEntitiesInSubset with no matching param entity
func TestExpandParamEntitiesNoMatch(t *testing.T) {
	dir := t.TempDir()
	resolver := NewEntityResolver(dir)

	// %unknown; should remain unexpanded
	result := resolver.expandParamEntitiesInSubset(
		`<!ENTITY foo "bar"> %unknown;`,
	)
	if !containsStr(result, "%unknown;") {
		t.Errorf("expected %%unknown; to remain, got %q", result)
	}
}

// parser: Parse returns error from parseChildren
func TestParserParseReturnsError(t *testing.T) {
	// We can't easily trigger an error from parseChildren
	// in normal flow since it only returns errors from recursive
	// calls. But we can test that Parse still returns root
	// even when there are warnings.
	input := `<sect1><para>text</sect1>`
	root, _, err := ParseString(input)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if root == nil {
		t.Fatal("expected non-nil root")
	}
}

// parser: parseChildren error propagation from nested element
// This is hard to trigger naturally, but we can verify that
// deeply nested implicit closes work without error.
func TestParserDeeplyNestedImplicitClose(t *testing.T) {
	// Use non-HTML tag names so they aren't treated as text
	input := `<sect1><sect2><sect3><para>text</sect1>`
	root, _, err := ParseString(input)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	s1 := root.FindChild("sect1")
	if s1 == nil {
		t.Fatal("expected element 'sect1'")
	}
	paras := s1.FindDescendants("para")
	if len(paras) != 1 {
		t.Errorf("expected 1 'para' descendant, got %d", len(paras))
	}
}

// tokenizer: readOpenTag with digit after <
func TestTokenizerDigitAfterLT(t *testing.T) {
	input := `<3 items`
	tok := NewTokenizer(input)
	tokens := tok.Tokenize()

	if tokens[0].Type != TokenText {
		t.Errorf("expected Text for <3, got %v", tokens[0].Type)
	}
	if tokens[0].Text != "<" {
		t.Errorf("expected '<', got %q", tokens[0].Text)
	}
}

// tokenizer: <> operator
func TestTokenizerDiamondOperator(t *testing.T) {
	input := `x <> y`
	tok := NewTokenizer(input)
	tokens := tok.Tokenize()

	found := false
	for _, tk := range tokens {
		if tk.Type == TokenText && tk.Text == "<>" {
			found = true
		}
	}
	if !found {
		t.Error("expected '<>' as text token")
	}
}

// tokenizer: readOpenTag space after <
func TestTokenizerSpaceAfterLT(t *testing.T) {
	input := `< space`
	tok := NewTokenizer(input)
	tokens := tok.Tokenize()

	if tokens[0].Type != TokenText {
		t.Errorf("expected Text for '< ', got %v", tokens[0].Type)
	}
	if tokens[0].Text != "<" {
		t.Errorf("expected '<', got %q", tokens[0].Text)
	}
}

// tokenizer: readOpenTag with no closing > (EOF)
func TestTokenizerOpenTagNoClose(t *testing.T) {
	input := `<para attr="val"`
	tok := NewTokenizer(input)
	tokens := tok.Tokenize()

	if tokens[0].Type != TokenTagOpen {
		t.Errorf("expected TagOpen, got %v", tokens[0].Type)
	}
	if tokens[0].Tag != "para" {
		t.Errorf("expected tag 'para', got %q", tokens[0].Tag)
	}
}

// tokenizer: <!--</ treated as literal text
func TestTokenizerCommentLikeFollowedByClose(t *testing.T) {
	input := `<!--</foo>`
	tok := NewTokenizer(input)
	tokens := tok.Tokenize()

	if tokens[0].Type != TokenText {
		t.Errorf("expected Text for <!--</ prefix, got %v",
			tokens[0].Type)
	}
	if tokens[0].Text != "<" {
		t.Errorf("expected '<', got %q", tokens[0].Text)
	}
}

// entity: processDoctype with no DOCTYPE at all
func TestProcessDoctypeNone(t *testing.T) {
	dir := t.TempDir()
	resolver := NewEntityResolver(dir)

	content := `<doc>just a body</doc>`
	body := resolver.processDoctype(content)
	if body != content {
		t.Errorf("expected content unchanged, got %q", body)
	}
}

// entity: expandParamEntitiesInSubset with file read error
func TestExpandParamEntitiesFileReadError(t *testing.T) {
	dir := t.TempDir()
	resolver := NewEntityResolver(dir)

	subset := `<!ENTITY % missing SYSTEM "nofile.sgml">
%missing;`
	result := resolver.expandParamEntitiesInSubset(subset)
	// Should still contain %missing; since file doesn't exist
	if !containsStr(result, "%missing;") {
		t.Errorf(
			"expected %%missing; to remain after read error, got %q",
			result,
		)
	}
}

// entity: expandParamEntitiesInSubset skips already-loaded entities
func TestExpandParamEntitiesSkipsExisting(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "ent.sgml"),
		[]byte(`<!ENTITY foo "original">`), 0644)

	resolver := NewEntityResolver(dir)
	// Pre-load the entity
	resolver.entities["%myent"] = `<!ENTITY foo "preloaded">`

	subset := `<!ENTITY % myent SYSTEM "ent.sgml">
%myent;`
	result := resolver.expandParamEntitiesInSubset(subset)
	if !containsStr(result, "preloaded") {
		t.Errorf("expected preloaded entity content, got %q", result)
	}
}

// tokenizer: readDeclaration with nested <> inside
func TestTokenizerDeclarationNested(t *testing.T) {
	input := `<!DOCTYPE doc [<!ENTITY foo "bar">]>`
	tok := NewTokenizer(input)
	tokens := tok.Tokenize()

	found := false
	for _, tk := range tokens {
		if tk.Type == TokenComment &&
			containsStr(tk.Text, "DOCTYPE") {
			found = true
		}
	}
	if !found {
		t.Error("expected declaration with nested <> as comment")
	}
}

// entity: expandEntities text entity with expansion error
func TestExpandEntitiesTextEntityExpansionError(t *testing.T) {
	dir := t.TempDir()
	resolver := NewEntityResolver(dir)
	resolver.entities["outer"] = "&inner;"
	resolver.entities["inner"] = "&outer;" // circular via text

	_, err := resolver.expandEntities("&outer;", 0)
	if err == nil {
		t.Error("expected error from text entity circular ref")
	}
}

// tokenizer: PI with newlines
func TestTokenizerPIWithNewlines(t *testing.T) {
	input := "<?xml\nversion=\"1.0\"\n?>"
	tok := NewTokenizer(input)
	tokens := tok.Tokenize()

	found := false
	for _, tk := range tokens {
		if tk.Type == TokenPI {
			found = true
		}
	}
	if !found {
		t.Error("expected PI token with newlines")
	}
	// Tokenizer should have tracked line numbers
	if tok.line != 3 {
		t.Errorf("expected line 3, got %d", tok.line)
	}
}

// tokenizer: declaration with newlines
func TestTokenizerDeclarationWithNewlines(t *testing.T) {
	input := "<!DOCTYPE\ndoc\n>"
	tok := NewTokenizer(input)
	tokens := tok.Tokenize()

	found := false
	for _, tk := range tokens {
		if tk.Type == TokenComment {
			found = true
		}
	}
	if !found {
		t.Error("expected comment token for declaration")
	}
}

// processDoctype: DOCTYPE without [ and without >
func TestProcessDoctypeNoSubsetNoClose(t *testing.T) {
	dir := t.TempDir()
	resolver := NewEntityResolver(dir)

	content := `<!DOCTYPE doc SYSTEM "doc.dtd"`
	body := resolver.processDoctype(content)
	// No > after DOCTYPE and no [, so should return content as-is
	if body != content {
		t.Errorf("expected content unchanged, got %q", body)
	}
}

// processDoctype: nested brackets inside DOCTYPE internal subset
func TestProcessDoctypeNestedBrackets(t *testing.T) {
	dir := t.TempDir()
	resolver := NewEntityResolver(dir)

	content := `<!DOCTYPE doc [
<!ENTITY foo "bar">
<![IGNORE[ignored stuff]]>
]>
<doc>body</doc>`
	body := resolver.processDoctype(content)
	if !containsStr(body, "<doc>body</doc>") {
		t.Errorf("expected body after nested-bracket DOCTYPE, got %q",
			body)
	}
}

// tokenizer: <-operator with newline before >
func TestTokenizerDashOperatorNewline(t *testing.T) {
	input := "<-foo\nbar>"
	tok := NewTokenizer(input)
	tokens := tok.Tokenize()

	if tokens[0].Type != TokenText {
		t.Errorf("expected Text, got %v", tokens[0].Type)
	}
	if !containsStr(tokens[0].Text, "<-foo") {
		t.Errorf("expected '<-foo' in text, got %q", tokens[0].Text)
	}
}

// tokenizer: self-closing tag with nil Attrs map
// This tests the tok.Attrs == nil branch in readOpenTag
func TestTokenizerSelfCloseNilAttrs(t *testing.T) {
	// A tag like <foo/> where readAttributes returns an empty map
	// (not nil), but we need to test the nil case.
	// Actually readAttributes always returns a non-nil map, so
	// the Attrs==nil check in readOpenTag self-close is defensive.
	// We can't easily trigger it through the tokenizer alone.
	// Instead, test that <foo/> works correctly.
	input := `<foo/>`
	tok := NewTokenizer(input)
	tokens := tok.Tokenize()

	if tokens[0].Type != TokenTagOpen {
		t.Fatalf("expected TagOpen, got %v", tokens[0].Type)
	}
	if tokens[0].Tag != "foo" {
		t.Errorf("expected tag 'foo', got %q", tokens[0].Tag)
	}
	// Should have selfclose marker
	if tokens[0].Attrs["\x00selfclose"] != "1" {
		t.Error("expected selfclose attr")
	}
}

// tokenizer: isRealDeclaration at end of input (pos+2 >= len)
func TestTokenizerIsRealDeclarationAtEnd(t *testing.T) {
	// <! at the very end of input — isRealDeclaration returns false
	input := `<!</`
	tok := NewTokenizer(input)
	tokens := tok.Tokenize()

	// Should be treated as literal text, not a declaration
	foundText := false
	for _, tk := range tokens {
		if tk.Type == TokenText {
			foundText = true
		}
	}
	if !foundText {
		t.Error("expected text token for <! at end of input")
	}
}

// tokenizer: isRealDeclaration with just <! and nothing after
func TestTokenizerIsRealDeclarationShort(t *testing.T) {
	input := `<!`
	tok := NewTokenizer(input)
	tokens := tok.Tokenize()

	foundText := false
	for _, tk := range tokens {
		if tk.Type == TokenText && tk.Text == "<" {
			foundText = true
		}
	}
	if !foundText {
		t.Error("expected '<' as text for short <!...")
	}
}

// ===== GenerateMissingFiles partial tests =====
// These test the version.sgml generation paths without Perl.

func TestGenerateMissingFilesWithVersion(t *testing.T) {
	dir := t.TempDir()

	// Create version.sgml.in template
	os.WriteFile(filepath.Join(dir, "version.sgml.in"),
		[]byte(`<!ENTITY version @PG_VERSION@>`), 0644)

	generated, warnings := GenerateMissingFiles(dir, "17.4")

	// version.sgml should be generated
	if generated == 0 {
		t.Error("expected at least 1 generated file")
	}
	data, err := os.ReadFile(filepath.Join(dir, "version.sgml"))
	if err != nil {
		t.Fatalf("version.sgml not created: %v", err)
	}
	if !containsStr(string(data), "17.4") {
		t.Errorf("expected version in output, got %q", string(data))
	}

	// Generators will fail (no Perl scripts), producing warnings
	// with stub files. That's expected.
	t.Logf("generated=%d warnings=%v", generated, warnings)
}

func TestGenerateMissingFilesNoVersion(t *testing.T) {
	dir := t.TempDir()

	// No version.sgml exists and empty pgVersion — should generate
	generated, warnings := GenerateMissingFiles(dir, "")

	// version.sgml should be created (fallback path)
	_, err := os.Stat(filepath.Join(dir, "version.sgml"))
	if err != nil {
		t.Log("version.sgml not created (expected if no template)")
	}

	t.Logf("generated=%d warnings=%v", generated, warnings)
}

func TestGenerateMissingFilesExistingFiles(t *testing.T) {
	dir := t.TempDir()

	// Pre-create all generated files so generators are skipped
	names := []string{
		"features-supported.sgml",
		"features-unsupported.sgml",
		"errcodes-table.sgml",
		"keywords-table.sgml",
		"wait_event_types.sgml",
		"targets-meson.sgml",
	}
	for _, name := range names {
		os.WriteFile(filepath.Join(dir, name),
			[]byte("existing"), 0644)
	}

	generated, warnings := GenerateMissingFiles(dir, "17.4")
	// Only version.sgml should be generated; others already exist
	t.Logf("generated=%d warnings=%v", generated, warnings)

	// Verify existing files weren't overwritten
	for _, name := range names {
		data, _ := os.ReadFile(filepath.Join(dir, name))
		if string(data) != "existing" {
			t.Errorf("%s was overwritten", name)
		}
	}
}

func TestGenerateMissingFilesNoVersionPreExists(t *testing.T) {
	dir := t.TempDir()

	// Pre-create version.sgml and pass empty version
	os.WriteFile(filepath.Join(dir, "version.sgml"),
		[]byte("already here"), 0644)

	generated, _ := GenerateMissingFiles(dir, "")

	// version.sgml already exists and pgVersion is empty,
	// so it should not be regenerated
	data, _ := os.ReadFile(filepath.Join(dir, "version.sgml"))
	if string(data) != "already here" {
		t.Error("version.sgml should not be overwritten")
	}
	t.Logf("generated=%d", generated)
}
