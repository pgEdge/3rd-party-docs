package md

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/pgEdge/postgresql-docs/builder/shared"
)

func TestFilterDocFiles(t *testing.T) {
	input := []string{
		"README.md",
		"CHANGELOG.md",
		"CONTRIBUTING.md",
		"LICENSE.md",
		"CODE_OF_CONDUCT.md",
		"config.md",
		"usage.md",
		"frag-config-man.md",
		"frag-usage-man.md",
		"changes.md",
		"test/README.md",
		"tests/guide.md",
		".github/ISSUE_TEMPLATE.md",
	}
	got := filterDocFiles(input)
	want := []string{"README.md", "config.md", "usage.md"}
	if len(got) != len(want) {
		t.Fatalf("got %v, want %v", got, want)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("got[%d] = %q, want %q", i, got[i], want[i])
		}
	}
}

func TestGithubAnchor(t *testing.T) {
	tests := []struct {
		title string
		want  string
	}{
		{"Getting Started", "getting-started"},
		{"Installation & Setup", "installation--setup"},
		{"pgvector", "pgvector"},
		{"Half-Precision Vectors", "half-precision-vectors"},
		{"C++ Example", "c-example"},
		{"What's New in v2", "whats-new-in-v2"},
		{"foo_bar_baz", "foo_bar_baz"},
	}
	for _, tt := range tests {
		got := githubAnchor(tt.title)
		if got != tt.want {
			t.Errorf("githubAnchor(%q) = %q, want %q",
				tt.title, got, tt.want)
		}
	}
}

func TestSplitMarkdown(t *testing.T) {
	content := `# My Project

Intro paragraph here.

Some more intro.

## Installation

Install instructions.

### From Source

Build from source.

## Usage

Usage instructions.

## FAQ

Frequently asked questions.
`
	res := splitMarkdown(content)

	if res.title != "My Project" {
		t.Errorf("title = %q, want %q", res.title, "My Project")
	}

	if !strings.Contains(res.intro, "Intro paragraph") {
		t.Error("intro should contain intro paragraph")
	}
	if strings.Contains(res.intro, "## Installation") {
		t.Error("intro should not contain H2 sections")
	}

	if len(res.sections) != 3 {
		t.Fatalf("got %d sections, want 3", len(res.sections))
	}

	wantTitles := []string{"Installation", "Usage", "FAQ"}
	for i, s := range res.sections {
		if s.title != wantTitles[i] {
			t.Errorf("section[%d].title = %q, want %q",
				i, s.title, wantTitles[i])
		}
	}

	// Installation section should include H3 subsection
	if !strings.Contains(res.sections[0].content, "### From Source") {
		t.Error("Installation should contain subsection")
	}
}

func TestSplitMarkdownNoH2(t *testing.T) {
	content := "# Title\n\nJust a simple doc.\n"
	res := splitMarkdown(content)
	if res.title != "Title" {
		t.Errorf("title = %q, want %q", res.title, "Title")
	}
	if len(res.sections) != 0 {
		t.Errorf("got %d sections, want 0", len(res.sections))
	}
}

func TestSplitMarkdownCodeBlock(t *testing.T) {
	content := "# Title\n\n## Real Section\n\n" +
		"```markdown\n## Not A Section\ncode here\n```\n\n" +
		"After code.\n"
	res := splitMarkdown(content)
	if len(res.sections) != 1 {
		t.Fatalf("got %d sections, want 1", len(res.sections))
	}
	if res.sections[0].title != "Real Section" {
		t.Errorf("title = %q, want %q",
			res.sections[0].title, "Real Section")
	}
	if !strings.Contains(res.sections[0].content,
		"## Not A Section") {
		t.Error("code block content should be preserved")
	}
}

func TestPromoteHeadings(t *testing.T) {
	input := "## Title\n\nText\n\n### Sub\n\n#### Deep\n"
	got := promoteHeadings(input)
	if !strings.Contains(got, "# Title\n") {
		t.Error("H2 should become H1")
	}
	if !strings.Contains(got, "## Sub\n") {
		t.Error("H3 should become H2")
	}
	if !strings.Contains(got, "### Deep\n") {
		t.Error("H4 should become H3")
	}
}

func TestPromoteHeadingsCodeBlock(t *testing.T) {
	input := "## Title\n\n```\n## Not promoted\n```\n"
	got := promoteHeadings(input)
	if !strings.Contains(got, "# Title\n") {
		t.Error("H2 should become H1")
	}
	if !strings.Contains(got, "## Not promoted") {
		t.Error("heading inside code block should not be promoted")
	}
}

func TestRewriteAnchors(t *testing.T) {
	anchorMap := map[string]string{
		"installation": "installation.md",
		"from-source":  "installation.md#from-source",
		"usage":        "usage.md",
	}

	input := "See [install](#installation) and " +
		"[build](#from-source) and [use](#usage) and " +
		"[unknown](#other)."

	got := rewriteAnchors(input, anchorMap)

	if !strings.Contains(got, "](installation.md)") {
		t.Error("should rewrite #installation")
	}
	if !strings.Contains(got,
		"](installation.md#from-source)") {
		t.Error("should rewrite #from-source")
	}
	if !strings.Contains(got, "](usage.md)") {
		t.Error("should rewrite #usage")
	}
	if !strings.Contains(got, "](#other)") {
		t.Error("should preserve unknown anchors")
	}
}

func TestBuildAnchorMap(t *testing.T) {
	sections := []section{
		{
			title: "Installation",
			slug:  "installation",
			content: "## Installation\n\n### From Source\n\n" +
				"Build steps.\n",
		},
		{
			title:   "Usage",
			slug:    "usage",
			content: "## Usage\n\nUse it.\n",
		},
	}
	m := buildAnchorMap(sections)

	if m["installation"] != "installation.md" {
		t.Errorf("installation = %q", m["installation"])
	}
	if m["from-source"] != "installation.md#from-source" {
		t.Errorf("from-source = %q", m["from-source"])
	}
	if m["usage"] != "usage.md" {
		t.Errorf("usage = %q", m["usage"])
	}
}

func TestConvertEmoji(t *testing.T) {
	input := "| Service | Supported |\n" +
		"| [Aiven](https://aiven.io/) | :heavy_check_mark: |\n" +
		"| [Heroku](https://heroku.com/) | :x: |\n" +
		"Normal text with no emoji.\n"
	got := convertEmoji(input)

	if strings.Contains(got, ":heavy_check_mark:") {
		t.Error("should replace :heavy_check_mark:")
	}
	if strings.Contains(got, ":x:") {
		t.Error("should replace :x:")
	}
	if !strings.Contains(got, "\u2714\uFE0F") {
		t.Error("should contain check mark unicode")
	}
	if !strings.Contains(got, "\u274C") {
		t.Error("should contain X unicode")
	}
	if !strings.Contains(got, "Normal text") {
		t.Error("should preserve non-emoji text")
	}
}

func TestStripLeadingImages(t *testing.T) {
	input := "[![Banner](/github-banner.png)](https://example.com/)\n" +
		"\n# Title\n\nSome content.\n" +
		"![inline](img.png)\n"
	got := stripLeadingImages(input)

	if strings.Contains(got, "Banner") {
		t.Error("should strip leading image before heading")
	}
	if !strings.Contains(got, "# Title") {
		t.Error("should preserve heading")
	}
	if !strings.Contains(got, "![inline](img.png)") {
		t.Error("should preserve images after first heading")
	}
}

func TestConvertAlerts(t *testing.T) {
	input := `Before.

> [!NOTE]
> This is a note.
> Second line.

After.

> [!WARNING]
> Be careful.

> Normal blockquote.
`
	got := convertAlerts(input)

	if !strings.Contains(got, "!!! note") {
		t.Error("should convert NOTE alert")
	}
	if !strings.Contains(got, "    This is a note.") {
		t.Error("should indent note body")
	}
	if !strings.Contains(got, "    Second line.") {
		t.Error("should include continuation lines")
	}
	if !strings.Contains(got, "!!! warning") {
		t.Error("should convert WARNING alert")
	}
	if !strings.Contains(got, "> Normal blockquote.") {
		t.Error("should preserve normal blockquotes")
	}
}

func TestConvertAlertsCaution(t *testing.T) {
	input := "> [!CAUTION]\n> Danger zone.\n"
	got := convertAlerts(input)
	if !strings.Contains(got, "!!! danger") {
		t.Error("CAUTION should map to danger")
	}
}

func TestExtractTitle(t *testing.T) {
	tests := []struct {
		content  string
		filename string
		want     string
	}{
		{"# My Title\n\nBody.", "file.md", "My Title"},
		{"No heading here.", "config.md", "config"},
		{"## Only H2\n\nBody.", "readme.md", "Only H2"},
		{
			"<h1 align=\"center\">\n <b>My Project</b>\n</h1>\n\nBody.",
			"readme.md",
			"My Project",
		},
		{
			"<H1>Simple HTML Title</H1>\n\nBody.",
			"file.md",
			"Simple HTML Title",
		},
		{
			"## Server\n\nDescription.\n",
			"index.md",
			"Server",
		},
		{
			"<h1 align=\"center\">\n <b>Title</b>\n</h1>\n\n" +
				"```bash\n# not a heading\ndocker up\n```\n",
			"readme.md",
			"Title",
		},
		{
			"Intro.\n\n```bash\n# fake heading\n```\n\n" +
				"## Real H2\n\nBody.\n",
			"file.md",
			"Real H2",
		},
	}
	for _, tt := range tests {
		got := extractTitle(tt.content, tt.filename)
		if got != tt.want {
			t.Errorf("extractTitle(%q, %q) = %q, want %q",
				tt.content, tt.filename, got, tt.want)
		}
	}
}

func TestConverterSplitFile(t *testing.T) {
	srcDir := t.TempDir()
	outDir := t.TempDir()

	readme := `# Test Project

Introduction.

## Installation

Install steps.

## Usage

Usage info.
`
	if err := os.WriteFile(
		filepath.Join(srcDir, "README.md"),
		[]byte(readme), 0644); err != nil {
		t.Fatal(err)
	}

	c := NewConverter(srcDir, outDir, "Test v1.0", false)
	if err := c.Convert(); err != nil {
		t.Fatal(err)
	}

	files := c.Files()
	if len(files) != 3 {
		t.Fatalf("got %d files, want 3", len(files))
	}

	// Check index.md exists
	indexData, err := os.ReadFile(
		filepath.Join(outDir, "index.md"))
	if err != nil {
		t.Fatal("index.md not created")
	}
	if !strings.Contains(string(indexData), "# Test Project") {
		t.Error("index.md should contain title")
	}

	// Check section files
	instData, err := os.ReadFile(
		filepath.Join(outDir, "installation.md"))
	if err != nil {
		t.Fatal("installation.md not created")
	}
	if !strings.Contains(string(instData), "# Installation") {
		t.Error("installation.md should have promoted H1")
	}

	usageData, err := os.ReadFile(
		filepath.Join(outDir, "usage.md"))
	if err != nil {
		t.Fatal("usage.md not created")
	}
	if !strings.Contains(string(usageData), "# Usage") {
		t.Error("usage.md should have promoted H1")
	}
}

func TestConverterCopyFiles(t *testing.T) {
	srcDir := t.TempDir()
	outDir := t.TempDir()

	os.WriteFile(filepath.Join(srcDir, "config.md"),
		[]byte("# Configuration\n\nConfig docs.\n"), 0644)
	os.WriteFile(filepath.Join(srcDir, "usage.md"),
		[]byte("# Usage\n\nUsage docs.\n"), 0644)
	os.WriteFile(filepath.Join(srcDir, "frag-config-man.md"),
		[]byte("Fragment content.\n"), 0644)

	c := NewConverter(srcDir, outDir, "PgBouncer 1.25", false)
	if err := c.Convert(); err != nil {
		t.Fatal(err)
	}

	files := c.Files()
	// 2 doc files + generated index = 3
	if len(files) != 3 {
		t.Fatalf("got %d files, want 3", len(files))
	}

	// Fragment should not be copied
	if _, err := os.Stat(
		filepath.Join(outDir, "frag-config-man.md")); err == nil {
		t.Error("fragment file should not be copied")
	}

	// Index should be generated
	indexData, err := os.ReadFile(
		filepath.Join(outDir, "index.md"))
	if err != nil {
		t.Fatal("index.md not created")
	}
	if !strings.Contains(string(indexData), "PgBouncer 1.25") {
		t.Error("index should contain project name")
	}
}

func TestConverterCopyFilesWithREADME(t *testing.T) {
	srcDir := t.TempDir()
	outDir := t.TempDir()

	os.WriteFile(filepath.Join(srcDir, "README.md"),
		[]byte("# Home\n\nWelcome.\n"), 0644)
	os.WriteFile(filepath.Join(srcDir, "guide.md"),
		[]byte("# Guide\n\nGuide content.\n"), 0644)

	c := NewConverter(srcDir, outDir, "Test v1", false)
	if err := c.Convert(); err != nil {
		t.Fatal(err)
	}

	// README.md should become index.md
	if _, err := os.Stat(
		filepath.Join(outDir, "index.md")); err != nil {
		t.Error("README.md should be renamed to index.md")
	}

	files := c.Files()
	// README→index + guide = 2 (no generated index needed)
	if len(files) != 2 {
		t.Fatalf("got %d files, want 2", len(files))
	}
}

func TestFindMarkdownFiles(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "readme.md"),
		[]byte("# Hi\n"), 0644)
	os.WriteFile(filepath.Join(dir, "GUIDE.MD"),
		[]byte("# Guide\n"), 0644)
	os.WriteFile(filepath.Join(dir, "Makefile"),
		[]byte("all:\n"), 0644)
	os.Mkdir(filepath.Join(dir, "subdir"), 0755)

	files, err := findMarkdownFiles(dir)
	if err != nil {
		t.Fatal(err)
	}
	if len(files) != 2 {
		t.Fatalf("got %d files, want 2: %v", len(files), files)
	}
}

func TestFindMarkdownFilesRecursive(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "index.md"),
		[]byte("# Home\n"), 0644)
	sub := filepath.Join(dir, "api")
	os.MkdirAll(sub, 0755)
	os.WriteFile(filepath.Join(sub, "search.md"),
		[]byte("# Search\n"), 0644)
	os.WriteFile(filepath.Join(sub, "index.md"),
		[]byte("# API\n"), 0644)
	deep := filepath.Join(dir, "ext", "examples")
	os.MkdirAll(deep, 0755)
	os.WriteFile(filepath.Join(deep, "demo.md"),
		[]byte("# Demo\n"), 0644)

	files, err := findMarkdownFiles(dir)
	if err != nil {
		t.Fatal(err)
	}
	if len(files) != 4 {
		t.Fatalf("got %d files, want 4: %v", len(files), files)
	}

	// Check that paths are relative
	for _, f := range files {
		if filepath.IsAbs(f) {
			t.Errorf("expected relative path, got %q", f)
		}
	}
}

func TestResolveSnippets(t *testing.T) {
	// Create a repo-like structure:
	// repo/README.md
	// repo/docs/index.md (contains snippet include)
	dir := t.TempDir()
	repoRoot := dir
	docsDir := filepath.Join(dir, "docs")
	os.MkdirAll(docsDir, 0755)

	os.WriteFile(filepath.Join(repoRoot, "README.md"),
		[]byte("# My Project\n\nWelcome.\n"), 0644)
	os.WriteFile(filepath.Join(docsDir, "index.md"),
		[]byte("--8<-- \"README.md\"\n"), 0644)

	content := "--8<-- \"README.md\"\n"
	filePath := filepath.Join(docsDir, "index.md")
	got := shared.ResolveSnippets(content, filePath, repoRoot)

	if !strings.Contains(got, "# My Project") {
		t.Errorf("snippet not resolved, got: %q", got)
	}
	if strings.Contains(got, "--8<--") {
		t.Error("snippet directive should be replaced")
	}
}

func TestResolveSnippetsUnresolvable(t *testing.T) {
	content := "--8<-- \"nonexistent.md\"\n"
	got := shared.ResolveSnippets(content, "/tmp/fake.md", "/tmp")
	if !strings.Contains(got, "--8<--") {
		t.Error("unresolvable snippet should be left as-is")
	}
}

func TestConverterRecursiveCopy(t *testing.T) {
	srcDir := t.TempDir()
	outDir := t.TempDir()

	// Create nested doc structure
	os.WriteFile(filepath.Join(srcDir, "index.md"),
		[]byte("# Home\n\nWelcome.\n"), 0644)
	apiDir := filepath.Join(srcDir, "api")
	os.MkdirAll(apiDir, 0755)
	os.WriteFile(filepath.Join(apiDir, "search.md"),
		[]byte("# Search API\n\nSearch docs.\n"), 0644)
	os.WriteFile(filepath.Join(apiDir, "index.md"),
		[]byte("# API Reference\n\nAPI docs.\n"), 0644)

	c := NewConverter(srcDir, outDir, "Test v1", false)
	if err := c.Convert(); err != nil {
		t.Fatal(err)
	}

	files := c.Files()
	if len(files) != 3 {
		t.Fatalf("got %d files, want 3: %v", len(files), files)
	}

	// Check nested files exist in output
	if _, err := os.Stat(
		filepath.Join(outDir, "api", "search.md")); err != nil {
		t.Error("api/search.md not created")
	}
	if _, err := os.Stat(
		filepath.Join(outDir, "api", "index.md")); err != nil {
		t.Error("api/index.md not created")
	}
}

func TestConverterSnippetResolution(t *testing.T) {
	// Simulate pg_vectorize layout:
	// repo/README.md
	// repo/docs/index.md (--8<-- "README.md")
	// repo/docs/guide.md
	repoDir := t.TempDir()
	srcDir := filepath.Join(repoDir, "docs")
	outDir := t.TempDir()

	os.MkdirAll(srcDir, 0755)
	os.WriteFile(filepath.Join(repoDir, "README.md"),
		[]byte("# My Project\n\nProject intro.\n\n"+
			"## Installation\n\nInstall steps.\n"), 0644)
	os.WriteFile(filepath.Join(srcDir, "index.md"),
		[]byte("--8<-- \"README.md\"\n"), 0644)
	os.WriteFile(filepath.Join(srcDir, "guide.md"),
		[]byte("# Guide\n\nGuide content.\n"), 0644)

	c := NewConverter(srcDir, outDir, "Test v1", false)
	if err := c.Convert(); err != nil {
		t.Fatal(err)
	}

	// index.md should have the resolved README content
	data, err := os.ReadFile(filepath.Join(outDir, "index.md"))
	if err != nil {
		t.Fatal("index.md not created")
	}
	if !strings.Contains(string(data), "# My Project") {
		t.Error("snippet should be resolved in index.md")
	}
	if strings.Contains(string(data), "--8<--") {
		t.Error("snippet directive should be replaced")
	}
}
