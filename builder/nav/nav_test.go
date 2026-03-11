//-------------------------------------------------------------------------
//
// pgEdge PostgreSQL Docs
//
// Copyright (c) 2026, pgEdge, Inc.
// This software is released under The PostgreSQL License
//
//-------------------------------------------------------------------------

package nav

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/pgEdge/postgresql-docs/builder/convert"
)

func TestBuildNav_EmptyFiles(t *testing.T) {
	root := BuildNav(nil)
	if root == nil {
		t.Fatal("expected non-nil root")
	}
	if root.Title != "root" {
		t.Errorf("expected root title 'root', got %q", root.Title)
	}
	if len(root.Children) != 0 {
		t.Errorf("expected no children, got %d", len(root.Children))
	}
}

func TestBuildNav_FlatFiles(t *testing.T) {
	files := []*convert.FileEntry{
		{Path: "intro.md", Title: "Introduction", Order: 1},
		{Path: "setup.md", Title: "Setup", Order: 2},
		{Path: "usage.md", Title: "Usage", Order: 3},
	}
	root := BuildNav(files)
	if len(root.Children) != 3 {
		t.Fatalf("expected 3 children, got %d", len(root.Children))
	}
	expected := []struct {
		title string
		path  string
	}{
		{"Introduction", "intro.md"},
		{"Setup", "setup.md"},
		{"Usage", "usage.md"},
	}
	for i, e := range expected {
		child := root.Children[i]
		if child.Title != e.title {
			t.Errorf("child[%d] title = %q, want %q", i, child.Title, e.title)
		}
		if child.Path != e.path {
			t.Errorf("child[%d] path = %q, want %q", i, child.Path, e.path)
		}
	}
}

func TestBuildNav_NestedStructure(t *testing.T) {
	files := []*convert.FileEntry{
		{Path: "tutorial/index.md", Title: "Tutorial", Order: 1},
		{Path: "tutorial/start.md", Title: "Getting Started", Order: 2},
		{Path: "tutorial/advanced.md", Title: "Advanced", Order: 3},
	}
	root := BuildNav(files)
	if len(root.Children) != 1 {
		t.Fatalf("expected 1 top-level child, got %d", len(root.Children))
	}
	tutorial := root.Children[0]
	if tutorial.Title != "Tutorial" {
		t.Errorf("tutorial title = %q, want 'Tutorial'", tutorial.Title)
	}
	if tutorial.Path != "tutorial/index.md" {
		t.Errorf("tutorial path = %q, want 'tutorial/index.md'",
			tutorial.Path)
	}
	if len(tutorial.Children) != 2 {
		t.Fatalf("expected 2 children under tutorial, got %d",
			len(tutorial.Children))
	}
	if tutorial.Children[0].Title != "Getting Started" {
		t.Errorf("first child title = %q, want 'Getting Started'",
			tutorial.Children[0].Title)
	}
}

func TestBuildNav_DeepNesting(t *testing.T) {
	files := []*convert.FileEntry{
		{Path: "part/chapter/section.md", Title: "Section", Order: 1},
	}
	root := BuildNav(files)
	if len(root.Children) != 1 {
		t.Fatalf("expected 1 child, got %d", len(root.Children))
	}
	part := root.Children[0]
	if part.Title != "Part" {
		t.Errorf("part title = %q, want 'Part'", part.Title)
	}
	if len(part.Children) != 1 {
		t.Fatalf("expected 1 child under part, got %d",
			len(part.Children))
	}
	chapter := part.Children[0]
	if chapter.Title != "Chapter" {
		t.Errorf("chapter title = %q, want 'Chapter'", chapter.Title)
	}
	if len(chapter.Children) != 1 {
		t.Fatalf("expected 1 child under chapter, got %d",
			len(chapter.Children))
	}
	section := chapter.Children[0]
	if section.Title != "Section" {
		t.Errorf("section title = %q, want 'Section'", section.Title)
	}
	if section.Path != "part/chapter/section.md" {
		t.Errorf("section path = %q, want 'part/chapter/section.md'",
			section.Path)
	}
}

func TestBuildNav_IndexSetsParentTitle(t *testing.T) {
	files := []*convert.FileEntry{
		{Path: "admin/index.md", Title: "Administration", Order: 1},
		{Path: "admin/config.md", Title: "Configuration", Order: 2},
	}
	root := BuildNav(files)
	admin := root.Children[0]
	if admin.Title != "Administration" {
		t.Errorf("admin title = %q, want 'Administration'", admin.Title)
	}
	if admin.Path != "admin/index.md" {
		t.Errorf("admin path = %q, want 'admin/index.md'", admin.Path)
	}
}

func TestBuildNav_OrderPreserved(t *testing.T) {
	files := []*convert.FileEntry{
		{Path: "aaa.md", Title: "Third", Order: 3},
		{Path: "bbb.md", Title: "First", Order: 1},
		{Path: "ccc.md", Title: "Second", Order: 2},
	}
	root := BuildNav(files)
	// Order should match insertion order (slice order)
	if root.Children[0].Title != "Third" {
		t.Errorf("first child = %q, want 'Third'", root.Children[0].Title)
	}
	if root.Children[1].Title != "First" {
		t.Errorf("second child = %q, want 'First'", root.Children[1].Title)
	}
}

func TestInsertEntry_RootLevel(t *testing.T) {
	root := &NavEntry{Title: "root"}
	insertEntry(root, "test.md", "Test Page", "")
	if len(root.Children) != 1 {
		t.Fatalf("expected 1 child, got %d", len(root.Children))
	}
	if root.Children[0].Title != "Test Page" {
		t.Errorf("title = %q, want 'Test Page'",
			root.Children[0].Title)
	}
	if root.Children[0].Path != "test.md" {
		t.Errorf("path = %q, want 'test.md'",
			root.Children[0].Path)
	}
}

func TestInsertEntry_IndexFile(t *testing.T) {
	root := &NavEntry{Title: "root"}
	insertEntry(root, "guide/index.md", "User Guide", "")
	if len(root.Children) != 1 {
		t.Fatalf("expected 1 child, got %d", len(root.Children))
	}
	guide := root.Children[0]
	if guide.Title != "User Guide" {
		t.Errorf("title = %q, want 'User Guide'", guide.Title)
	}
	if guide.Path != "guide/index.md" {
		t.Errorf("path = %q, want 'guide/index.md'", guide.Path)
	}
	if len(guide.Children) != 0 {
		t.Errorf("expected no children, got %d", len(guide.Children))
	}
}

func TestInsertEntry_IndexEmptyTitle(t *testing.T) {
	root := &NavEntry{Title: "root"}
	insertEntry(root, "section/index.md", "", "")
	section := root.Children[0]
	// When title is empty, the deslugified dir name should remain
	if section.Title != "Section" {
		t.Errorf("title = %q, want 'Section'", section.Title)
	}
}

func TestInsertEntry_DuplicateDir(t *testing.T) {
	root := &NavEntry{Title: "root"}
	insertEntry(root, "docs/a.md", "Page A", "")
	insertEntry(root, "docs/b.md", "Page B", "")
	if len(root.Children) != 1 {
		t.Fatalf("expected 1 dir child, got %d", len(root.Children))
	}
	docs := root.Children[0]
	if len(docs.Children) != 2 {
		t.Fatalf("expected 2 children under docs, got %d",
			len(docs.Children))
	}
}

func TestGenerateYAML_SimpleLeaf(t *testing.T) {
	root := &NavEntry{
		Title: "root",
		Children: []*NavEntry{
			{Title: "Introduction", Path: "intro.md"},
		},
	}
	yaml := GenerateYAML(root)
	if !strings.Contains(yaml, "nav:") {
		t.Error("expected 'nav:' header")
	}
	if !strings.Contains(yaml, "  - Introduction: intro.md") {
		t.Errorf("unexpected YAML output:\n%s", yaml)
	}
}

func TestGenerateYAML_LeafNoTitle(t *testing.T) {
	root := &NavEntry{
		Title: "root",
		Children: []*NavEntry{
			{Title: "", Path: "intro.md"},
		},
	}
	yaml := GenerateYAML(root)
	if !strings.Contains(yaml, "  - intro.md") {
		t.Errorf("expected bare path entry, got:\n%s", yaml)
	}
}

func TestGenerateYAML_NestedWithIndex(t *testing.T) {
	root := &NavEntry{
		Title: "root",
		Children: []*NavEntry{
			{
				Title: "Tutorial",
				Path:  "tutorial/index.md",
				Children: []*NavEntry{
					{Title: "Start", Path: "tutorial/start.md"},
				},
			},
		},
	}
	yaml := GenerateYAML(root)
	lines := strings.Split(yaml, "\n")
	// Check structure
	found := map[string]bool{}
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		found[trimmed] = true
	}
	if !found["- Tutorial:"] {
		t.Errorf("missing Tutorial section header in:\n%s", yaml)
	}
	if !found["- tutorial/index.md"] {
		t.Errorf("missing index entry in:\n%s", yaml)
	}
	if !found["- Start: tutorial/start.md"] {
		t.Errorf("missing Start entry in:\n%s", yaml)
	}
}

func TestGenerateYAML_QuotedTitle(t *testing.T) {
	root := &NavEntry{
		Title: "root",
		Children: []*NavEntry{
			{Title: "Part: Overview", Path: "part.md"},
		},
	}
	yaml := GenerateYAML(root)
	if !strings.Contains(yaml, "'Part: Overview'") {
		t.Errorf("expected quoted title, got:\n%s", yaml)
	}
}

func TestGenerateYAML_Empty(t *testing.T) {
	root := &NavEntry{Title: "root"}
	yaml := GenerateYAML(root)
	if yaml != "nav:\n" {
		t.Errorf("expected just 'nav:\\n', got %q", yaml)
	}
}

func TestUpdateMkdocsYML_ReplaceNav(t *testing.T) {
	tmpDir := t.TempDir()
	mkdocsPath := filepath.Join(tmpDir, "mkdocs.yml")

	original := `site_name: PostgreSQL Docs
theme:
  name: material

nav:
  - Home: index.md
  - Old: old.md

extra:
  key: value
`
	err := os.WriteFile(mkdocsPath, []byte(original), 0644)
	if err != nil {
		t.Fatal(err)
	}

	newNav := "nav:\n  - Intro: intro.md\n  - Setup: setup.md\n"
	err = UpdateMkdocsYML(mkdocsPath, newNav, "17.2")
	if err != nil {
		t.Fatal(err)
	}

	data, _ := os.ReadFile(mkdocsPath)
	content := string(data)

	if !strings.Contains(content, "site_name: PostgreSQL 17.2") {
		t.Errorf("site_name not updated:\n%s", content)
	}
	if !strings.Contains(content, "- Intro: intro.md") {
		t.Errorf("new nav not inserted:\n%s", content)
	}
	if !strings.Contains(content, "extra:") {
		t.Errorf("extra section lost:\n%s", content)
	}
	if strings.Contains(content, "Old: old.md") {
		t.Error("old nav entry should be removed")
	}
}

func TestUpdateMkdocsYML_NoExistingNav(t *testing.T) {
	tmpDir := t.TempDir()
	mkdocsPath := filepath.Join(tmpDir, "mkdocs.yml")

	original := `site_name: PostgreSQL Docs
theme:
  name: material
`
	err := os.WriteFile(mkdocsPath, []byte(original), 0644)
	if err != nil {
		t.Fatal(err)
	}

	newNav := "nav:\n  - Intro: intro.md\n"
	err = UpdateMkdocsYML(mkdocsPath, newNav, "")
	if err != nil {
		t.Fatal(err)
	}

	data, _ := os.ReadFile(mkdocsPath)
	content := string(data)

	if !strings.Contains(content, "nav:") {
		t.Errorf("nav not appended:\n%s", content)
	}
	if !strings.Contains(content, "- Intro: intro.md") {
		t.Errorf("nav entry missing:\n%s", content)
	}
	// site_name should not change when version is ""
	if !strings.Contains(content, "site_name: PostgreSQL Docs") {
		t.Errorf("site_name should not change:\n%s", content)
	}
}

func TestUpdateMkdocsYML_FileNotFound(t *testing.T) {
	err := UpdateMkdocsYML("/nonexistent/path/mkdocs.yml", "nav:\n", "")
	if err == nil {
		t.Error("expected error for missing file")
	}
}

func TestUpdateMkdocsYML_VersionUpdate(t *testing.T) {
	tmpDir := t.TempDir()
	mkdocsPath := filepath.Join(tmpDir, "mkdocs.yml")

	original := "site_name: PostgreSQL Docs\n"
	os.WriteFile(mkdocsPath, []byte(original), 0644)

	err := UpdateMkdocsYML(mkdocsPath, "nav:\n", "16.1")
	if err != nil {
		t.Fatal(err)
	}

	data, _ := os.ReadFile(mkdocsPath)
	if !strings.Contains(string(data), "site_name: PostgreSQL 16.1") {
		t.Errorf("version not set:\n%s", string(data))
	}
}

func TestYamlQuote_PlainString(t *testing.T) {
	result := yamlQuote("Hello World")
	if result != "Hello World" {
		t.Errorf("expected unquoted, got %q", result)
	}
}

func TestYamlQuote_Colon(t *testing.T) {
	result := yamlQuote("Part: Overview")
	if result != "'Part: Overview'" {
		t.Errorf("expected quoted, got %q", result)
	}
}

func TestYamlQuote_Hash(t *testing.T) {
	result := yamlQuote("Section #1")
	if result != "'Section #1'" {
		t.Errorf("expected quoted, got %q", result)
	}
}

func TestYamlQuote_Braces(t *testing.T) {
	result := yamlQuote("Config {default}")
	if result != "'Config {default}'" {
		t.Errorf("expected quoted, got %q", result)
	}
}

func TestYamlQuote_SingleQuoteEscaping(t *testing.T) {
	result := yamlQuote("It's a test: yes")
	if result != "'It''s a test: yes'" {
		t.Errorf("expected escaped quote, got %q", result)
	}
}

func TestYamlQuote_Ampersand(t *testing.T) {
	result := yamlQuote("A & B")
	if result != "'A & B'" {
		t.Errorf("expected quoted, got %q", result)
	}
}

func TestYamlQuote_Asterisk(t *testing.T) {
	result := yamlQuote("Note *important*")
	if result != "'Note *important*'" {
		t.Errorf("expected quoted, got %q", result)
	}
}

func TestYamlQuote_Percent(t *testing.T) {
	result := yamlQuote("100% Complete")
	if result != "'100% Complete'" {
		t.Errorf("expected quoted, got %q", result)
	}
}

func TestYamlQuote_Backtick(t *testing.T) {
	result := yamlQuote("Use `cmd`")
	if result != "'Use `cmd`'" {
		t.Errorf("expected quoted, got %q", result)
	}
}

func TestYamlQuote_DoubleQuote(t *testing.T) {
	result := yamlQuote(`Say "hello"`)
	expected := `'Say "hello"'`
	if result != expected {
		t.Errorf("expected %q, got %q", expected, result)
	}
}

func TestSlugMatch_Exact(t *testing.T) {
	// "Tutorial" slugifies to "tutorial"
	if !slugMatch("Tutorial", "tutorial") {
		t.Error("expected match for Tutorial/tutorial")
	}
}

func TestSlugMatch_MultiWord(t *testing.T) {
	// "Getting Started" should slugify to "getting-started"
	if !slugMatch("Getting Started", "getting-started") {
		t.Error("expected match for 'Getting Started'/'getting-started'")
	}
}

func TestSlugMatch_NoMatch(t *testing.T) {
	if slugMatch("Tutorial", "admin") {
		t.Error("expected no match for Tutorial/admin")
	}
}

func TestDeslugify_Simple(t *testing.T) {
	result := deslugify("tutorial")
	if result != "Tutorial" {
		t.Errorf("got %q, want 'Tutorial'", result)
	}
}

func TestDeslugify_Hyphens(t *testing.T) {
	result := deslugify("getting-started")
	if result != "Getting Started" {
		t.Errorf("got %q, want 'Getting Started'", result)
	}
}

func TestDeslugify_Underscores(t *testing.T) {
	result := deslugify("my_section")
	if result != "My Section" {
		t.Errorf("got %q, want 'My Section'", result)
	}
}

func TestDeslugify_Mixed(t *testing.T) {
	result := deslugify("my-cool_thing")
	if result != "My Cool Thing" {
		t.Errorf("got %q, want 'My Cool Thing'", result)
	}
}

func TestDeslugify_Empty(t *testing.T) {
	result := deslugify("")
	if result != "" {
		t.Errorf("got %q, want ''", result)
	}
}

func TestDeslugify_SingleChar(t *testing.T) {
	result := deslugify("a")
	if result != "A" {
		t.Errorf("got %q, want 'A'", result)
	}
}

func TestBuildNav_MultipleDirectories(t *testing.T) {
	files := []*convert.FileEntry{
		{Path: "tutorial/start.md", Title: "Start", Order: 1},
		{Path: "admin/config.md", Title: "Config", Order: 2},
		{Path: "tutorial/end.md", Title: "End", Order: 3},
	}
	root := BuildNav(files)
	if len(root.Children) != 2 {
		t.Fatalf("expected 2 top-level dirs, got %d",
			len(root.Children))
	}
	// First dir encountered should be Tutorial
	if root.Children[0].Title != "Tutorial" {
		t.Errorf("first dir = %q, want 'Tutorial'",
			root.Children[0].Title)
	}
	if root.Children[1].Title != "Admin" {
		t.Errorf("second dir = %q, want 'Admin'",
			root.Children[1].Title)
	}
	// Tutorial should have 2 children, admin 1
	if len(root.Children[0].Children) != 2 {
		t.Errorf("tutorial children = %d, want 2",
			len(root.Children[0].Children))
	}
}

func TestGenerateYAML_DeepNesting(t *testing.T) {
	root := &NavEntry{
		Title: "root",
		Children: []*NavEntry{
			{
				Title: "Part I",
				Children: []*NavEntry{
					{
						Title: "Chapter 1",
						Children: []*NavEntry{
							{Title: "Section A",
								Path: "part-i/ch1/a.md"},
						},
					},
				},
			},
		},
	}
	yaml := GenerateYAML(root)
	// Check indentation levels
	if !strings.Contains(yaml, "  - Part I:") {
		t.Errorf("missing Part I at depth 1:\n%s", yaml)
	}
	if !strings.Contains(yaml, "    - Chapter 1:") {
		t.Errorf("missing Chapter 1 at depth 2:\n%s", yaml)
	}
	if !strings.Contains(yaml, "      - Section A: part-i/ch1/a.md") {
		t.Errorf("missing Section A at depth 3:\n%s", yaml)
	}
}
