//-------------------------------------------------------------------------
//
// pgEdge PostgreSQL Docs
//
// Copyright (c) 2026, pgEdge, Inc.
// This software is released under The PostgreSQL License
//
//-------------------------------------------------------------------------

package backrest

import (
	"strings"
	"testing"

	"github.com/pgEdge/postgresql-docs/builder/sgml"
	"github.com/pgEdge/postgresql-docs/builder/shared"
)

func TestParseXMLString(t *testing.T) {
	xml := `<?xml version="1.0" encoding="UTF-8"?>
<doc title="Test Doc" subtitle="A test">
    <section id="intro">
        <title>Introduction</title>
        <p>Hello <b>world</b>.</p>
    </section>
</doc>`

	doc, err := parseXMLString(xml)
	if err != nil {
		t.Fatalf("parse error: %v", err)
	}
	if doc.Tag != "doc" {
		t.Errorf("expected doc tag, got %s", doc.Tag)
	}
	if doc.GetAttr("title") != "Test Doc" {
		t.Errorf("wrong title: %s", doc.GetAttr("title"))
	}
	sect := doc.FindChild("section")
	if sect == nil {
		t.Fatal("no section found")
	}
	if sect.GetAttr("id") != "intro" {
		t.Errorf("wrong section id: %s", sect.GetAttr("id"))
	}
}

func TestParseXMLSelfClosing(t *testing.T) {
	xml := `<doc title="Test">
    <p>Use <backrest/> with <postgres/>.</p>
</doc>`

	doc, err := parseXMLString(xml)
	if err != nil {
		t.Fatalf("parse error: %v", err)
	}
	p := doc.FindChild("p")
	if p == nil {
		t.Fatal("no p found")
	}
	// Should have backrest and postgres as child elements
	found := false
	for _, c := range p.Children {
		if c.Type == sgml.ElementNode && c.Tag == "backrest" {
			found = true
		}
	}
	if !found {
		t.Error("backrest element not found")
	}
}

func TestStripDoctype(t *testing.T) {
	input := `<?xml version="1.0"?>
<!DOCTYPE doc SYSTEM "doc.dtd">
<doc title="Test"></doc>`
	result := stripDoctype(input)
	if strings.Contains(result, "DOCTYPE") {
		t.Error("DOCTYPE not stripped")
	}
	if !strings.Contains(result, "<doc") {
		t.Error("doc element lost")
	}
}

func TestVariableSubstitution(t *testing.T) {
	vars := map[string]string{
		"project": "pgBackRest",
		"version": "2.57.0",
	}
	tests := []struct {
		input, want string
	}{
		{"{[project]}", "pgBackRest"},
		{"{[project]} v{[version]}", "pgBackRest v2.57.0"},
		{"{[unknown]}", "<unknown>"},
		{"no vars", "no vars"},
	}
	for _, tt := range tests {
		got := substituteVariables(tt.input, vars)
		if got != tt.want {
			t.Errorf("substitute(%q) = %q, want %q",
				tt.input, got, tt.want)
		}
	}
}

func TestChainedVariables(t *testing.T) {
	vars := map[string]string{
		"base-url": "https://example.com",
		"full-url": "{[base-url]}/docs",
	}
	got := substituteVariables("{[full-url]}", vars)
	if got != "https://example.com/docs" {
		t.Errorf("chained substitution failed: %s", got)
	}
}

func TestCollectVariables(t *testing.T) {
	xml := `<doc title="Test">
    <variable-list>
        <variable key="name">pgBackRest</variable>
        <variable key="ver">2.57</variable>
        <variable key="dynamic" eval="y">some perl code</variable>
    </variable-list>
</doc>`

	doc, err := parseXMLString(xml)
	if err != nil {
		t.Fatal(err)
	}
	vars := make(map[string]string)
	collectVariables(doc, vars)

	if vars["name"] != "pgBackRest" {
		t.Errorf("name = %q", vars["name"])
	}
	if vars["ver"] != "2.57" {
		t.Errorf("ver = %q", vars["ver"])
	}
	if _, ok := vars["dynamic"]; ok {
		t.Error("eval variable should be skipped")
	}
}

func TestCollectBlocks(t *testing.T) {
	xml := `<doc title="Test">
    <block-define id="my-block">
        <p>Block content</p>
    </block-define>
</doc>`

	doc, err := parseXMLString(xml)
	if err != nil {
		t.Fatal(err)
	}
	blocks := make(map[string]*sgml.Node)
	collectBlocks(doc, blocks)

	if _, ok := blocks["my-block"]; !ok {
		t.Error("block not collected")
	}
}

func TestInlineElements(t *testing.T) {
	xml := `<doc title="Test">
    <p>Use <backrest/> with <postgres/>. Run <exe/> for help.</p>
</doc>`

	doc, err := parseXMLString(xml)
	if err != nil {
		t.Fatal(err)
	}
	c := NewConverter(".", ".", "", false)
	c.vars["project-exe"] = "pgbackrest"

	w := shared.NewMarkdownWriter()
	p := doc.FindChild("p")
	c.convertInlineContent(p, w)
	result := w.String()

	if !strings.Contains(result, "pgBackRest") {
		t.Errorf("missing pgBackRest: %s", result)
	}
	if !strings.Contains(result, "PostgreSQL") {
		t.Errorf("missing PostgreSQL: %s", result)
	}
	if !strings.Contains(result, "`pgbackrest`") {
		t.Errorf("missing exe: %s", result)
	}
}

func TestInlineFormatting(t *testing.T) {
	xml := `<doc title="Test">
    <p><b>bold</b> and <i>italic</i> and <code>inline</code></p>
</doc>`

	doc, err := parseXMLString(xml)
	if err != nil {
		t.Fatal(err)
	}
	c := NewConverter(".", ".", "", false)
	w := shared.NewMarkdownWriter()
	p := doc.FindChild("p")
	c.convertInlineContent(p, w)
	result := w.String()

	if !strings.Contains(result, "**bold**") {
		t.Errorf("missing bold: %s", result)
	}
	if !strings.Contains(result, "*italic*") {
		t.Errorf("missing italic: %s", result)
	}
	if !strings.Contains(result, "`inline`") {
		t.Errorf("missing code: %s", result)
	}
}

func TestSemanticCodeElements(t *testing.T) {
	xml := `<doc title="Test">
    <p><file>/etc/config</file> <path>/var/lib</path> <cmd>ls</cmd> <host>pg1</host></p>
</doc>`

	doc, err := parseXMLString(xml)
	if err != nil {
		t.Fatal(err)
	}
	c := NewConverter(".", ".", "", false)
	w := shared.NewMarkdownWriter()
	p := doc.FindChild("p")
	c.convertInlineContent(p, w)
	result := w.String()

	for _, want := range []string{
		"`/etc/config`", "`/var/lib`", "`ls`", "`pg1`",
	} {
		if !strings.Contains(result, want) {
			t.Errorf("missing %s in: %s", want, result)
		}
	}
}

func TestLinkConversion(t *testing.T) {
	xml := `<doc title="Test">
    <p><link url="https://example.com">Example</link></p>
</doc>`

	doc, err := parseXMLString(xml)
	if err != nil {
		t.Fatal(err)
	}
	c := NewConverter(".", ".", "", false)
	w := shared.NewMarkdownWriter()
	p := doc.FindChild("p")
	c.convertInlineContent(p, w)
	result := w.String()

	if !strings.Contains(result, "[Example](https://example.com)") {
		t.Errorf("wrong link: %s", result)
	}
}

func TestAdmonition(t *testing.T) {
	xml := `<doc title="Test">
    <admonition type="warning">
        <p>Be careful!</p>
    </admonition>
</doc>`

	doc, err := parseXMLString(xml)
	if err != nil {
		t.Fatal(err)
	}
	c := NewConverter(".", ".", "", false)
	w := shared.NewMarkdownWriter()
	c.convertNode(doc.FindChild("admonition"), w, 1)
	result := w.String()

	if !strings.Contains(result, "!!! warning") {
		t.Errorf("missing admonition: %s", result)
	}
	if !strings.Contains(result, "Be careful!") {
		t.Errorf("missing content: %s", result)
	}
}

func TestCodeBlock(t *testing.T) {
	xml := `<doc title="Test">
    <code-block type="sql" title="Example Query">SELECT 1;</code-block>
</doc>`

	doc, err := parseXMLString(xml)
	if err != nil {
		t.Fatal(err)
	}
	c := NewConverter(".", ".", "", false)
	w := shared.NewMarkdownWriter()
	c.convertNode(doc.FindChild("code-block"), w, 1)
	result := w.String()

	if !strings.Contains(result, "```sql") {
		t.Errorf("missing language: %s", result)
	}
	if !strings.Contains(result, "SELECT 1;") {
		t.Errorf("missing content: %s", result)
	}
	if !strings.Contains(result, "**Example Query**") {
		t.Errorf("missing title: %s", result)
	}
}

func TestTableConversion(t *testing.T) {
	xml := `<doc title="Test">
    <table>
        <table-header>
            <table-column>Name</table-column>
            <table-column>Value</table-column>
        </table-header>
        <table-data>
            <table-row>
                <table-cell>foo</table-cell>
                <table-cell>bar</table-cell>
            </table-row>
        </table-data>
    </table>
</doc>`

	doc, err := parseXMLString(xml)
	if err != nil {
		t.Fatal(err)
	}
	c := NewConverter(".", ".", "", false)
	w := shared.NewMarkdownWriter()
	c.convertNode(doc.FindChild("table"), w, 1)
	result := w.String()

	if !strings.Contains(result, "| Name | Value |") {
		t.Errorf("missing header: %s", result)
	}
	if !strings.Contains(result, "| --- | --- |") {
		t.Errorf("missing separator: %s", result)
	}
	if !strings.Contains(result, "| foo | bar |") {
		t.Errorf("missing data: %s", result)
	}
}

func TestExtractTitle(t *testing.T) {
	xml := `<doc title="Test"><section id="test"><title>My Title</title><p>Content</p></section></doc>`
	doc, err := parseXMLString(xml)
	if err != nil {
		t.Fatal(err)
	}
	sect := doc.FindChild("section")
	if sect == nil {
		t.Fatal("no section")
	}
	title := extractTitle(sect)
	if title != "My Title" {
		t.Errorf("got %q", title)
	}
}

func TestReleaseConversion(t *testing.T) {
	xml := `<doc title="Releases">
    <release-list>
        <release version="2.57.0" date="2025-10-18" title="Test Release">
            <release-core-list>
                <release-bug-list>
                    <release-item><p>Fixed a bug.</p></release-item>
                </release-bug-list>
                <release-feature-list>
                    <release-item><p>Added a feature.</p></release-item>
                </release-feature-list>
            </release-core-list>
        </release>
    </release-list>
</doc>`

	doc, err := parseXMLString(xml)
	if err != nil {
		t.Fatal(err)
	}
	c := NewConverter(".", ".", "", false)
	w := shared.NewMarkdownWriter()
	c.convertRelease(doc, w)
	result := w.String()

	if !strings.Contains(result, "v2.57.0") {
		t.Errorf("missing version: %s", result)
	}
	if !strings.Contains(result, "Bug Fixes") {
		t.Errorf("missing bug section: %s", result)
	}
	if !strings.Contains(result, "Fixed a bug") {
		t.Errorf("missing bug content: %s", result)
	}
	if !strings.Contains(result, "Features") {
		t.Errorf("missing feature section: %s", result)
	}
}
