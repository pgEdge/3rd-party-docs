package mkdocsmode

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestExtractYAMLBlock(t *testing.T) {
	content := `site_name: Test
nav:
  - Home: index.md
  - Guide:
    - Intro: guide/intro.md

markdown_extensions:
  - admonition
  - pymdownx.superfences

plugins:
  - search
`
	nav := extractYAMLBlock(content, "nav")
	if !strings.HasPrefix(nav, "nav:") {
		t.Errorf("nav block should start with nav:, got: %q",
			nav[:20])
	}
	if !strings.Contains(nav, "guide/intro.md") {
		t.Error("nav block should contain guide path")
	}
	if strings.Contains(nav, "markdown_extensions") {
		t.Error("nav block should not contain extensions")
	}

	ext := extractYAMLBlock(content, "markdown_extensions")
	if !strings.Contains(ext, "admonition") {
		t.Error("extensions block should contain admonition")
	}

	missing := extractYAMLBlock(content, "nonexistent")
	if missing != "" {
		t.Error("missing block should return empty string")
	}
}

func TestExtractYAMLList(t *testing.T) {
	content := `markdown_extensions:
  - admonition
  - pymdownx.snippets:
      check_paths: true
  - pymdownx.superfences
plugins:
  - search
`
	items := extractYAMLList(content, "markdown_extensions")
	if len(items) != 3 {
		t.Fatalf("got %d items, want 3: %v", len(items), items)
	}
	if items[0] != "- admonition" {
		t.Errorf("item[0] = %q, want '- admonition'", items[0])
	}
	if !strings.Contains(items[1], "pymdownx.snippets") {
		t.Errorf("item[1] should contain snippets: %q", items[1])
	}
	if !strings.Contains(items[1], "check_paths") {
		t.Error("item[1] should include sub-config")
	}
}

func TestExtractNavPaths(t *testing.T) {
	nav := `nav:
  - Vectorize: 'index.md'
  - Server:
    - API:
      - Table: 'server/api/table.md'
      - Search: 'server/api/search.md'
  - Extension:
    - API:
      - Overview: 'extension/api/index.md'
`
	paths := extractNavPaths(nav)
	want := []string{
		"index.md",
		"server/api/table.md",
		"server/api/search.md",
		"extension/api/index.md",
	}
	if len(paths) != len(want) {
		t.Fatalf("got %d paths, want %d: %v",
			len(paths), len(want), paths)
	}
	for i, p := range want {
		if paths[i] != p {
			t.Errorf("paths[%d] = %q, want %q", i, paths[i], p)
		}
	}
}

func TestExtractNavPathsUnquoted(t *testing.T) {
	nav := `nav:
  - Home: index.md
  - Guide: guide/start.md
`
	paths := extractNavPaths(nav)
	if len(paths) != 2 {
		t.Fatalf("got %d paths, want 2: %v", len(paths), paths)
	}
	if paths[0] != "index.md" {
		t.Errorf("paths[0] = %q", paths[0])
	}
}

func TestItemName(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"- admonition", "admonition"},
		{"- pymdownx.snippets:\n    check_paths: true",
			"pymdownx.snippets"},
		{"- search", "search"},
		{"- toc:\n    permalink: true", "toc"},
	}
	for _, tt := range tests {
		got := itemName(tt.input)
		if got != tt.want {
			t.Errorf("itemName(%q) = %q, want %q",
				tt.input, got, tt.want)
		}
	}
}

func TestMergeListBlock(t *testing.T) {
	content := `markdown_extensions:
  - admonition
  - pymdownx.superfences

plugins:
  - search
`
	newExts := []string{
		"- pymdownx.highlight",
		"- admonition", // duplicate
	}
	result := mergeListBlock(content,
		"markdown_extensions", newExts, nil)

	if !strings.Contains(result, "pymdownx.highlight") {
		t.Error("should add pymdownx.highlight")
	}
	// Count occurrences of admonition
	count := strings.Count(result, "admonition")
	if count != 1 {
		t.Errorf("admonition should appear once, got %d", count)
	}
}

func TestMergeListBlockSkipsPlugins(t *testing.T) {
	content := `plugins:
  - search
`
	newPlugins := []string{
		"- mkdocstrings",
		"- awesome-pages",
	}
	result := mergeListBlock(content,
		"plugins", newPlugins, skipPlugins)

	if strings.Contains(result, "mkdocstrings") {
		t.Error("should skip mkdocstrings")
	}
	if !strings.Contains(result, "awesome-pages") {
		t.Error("should add awesome-pages")
	}
}

func TestMergeListBlockSkipsSnippets(t *testing.T) {
	content := `markdown_extensions:
  - admonition
`
	newExts := []string{
		"- pymdownx.snippets:\n    check_paths: true",
	}
	result := mergeListBlock(content,
		"markdown_extensions", newExts, nil)

	if strings.Contains(result, "snippets") {
		t.Error("should skip pymdownx.snippets")
	}
}

func TestConverterEndToEnd(t *testing.T) {
	// Setup: create a repo-like structure
	repoDir := t.TempDir()
	srcDir := filepath.Join(repoDir, "docs")
	outDir := t.TempDir()

	os.MkdirAll(filepath.Join(srcDir, "api"), 0755)

	// Upstream mkdocs.yml at repo root
	os.WriteFile(filepath.Join(repoDir, "mkdocs.yml"),
		[]byte(`site_name: Test Project
nav:
  - Home: index.md
  - API:
    - Reference: api/ref.md
markdown_extensions:
  - admonition
  - pymdownx.snippets:
      check_paths: true
plugins:
  - search
  - mkdocstrings
`), 0644)

	// README at repo root (for snippet resolution)
	os.WriteFile(filepath.Join(repoDir, "README.md"),
		[]byte("# Test Project\n\nWelcome.\n"), 0644)

	// Docs
	os.WriteFile(filepath.Join(srcDir, "index.md"),
		[]byte("--8<-- \"README.md\"\n"), 0644)
	os.WriteFile(filepath.Join(srcDir, "api", "ref.md"),
		[]byte("# API Reference\n\nAPI docs.\n"), 0644)

	c := NewConverter(srcDir, outDir, "Test v1", false)
	if err := c.Convert(); err != nil {
		t.Fatal(err)
	}

	// Check files were copied
	if _, err := os.Stat(
		filepath.Join(outDir, "api", "ref.md")); err != nil {
		t.Error("api/ref.md should be copied")
	}

	// Check snippets were resolved
	data, err := os.ReadFile(filepath.Join(outDir, "index.md"))
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(string(data), "--8<--") {
		t.Error("snippets should be resolved")
	}
	if !strings.Contains(string(data), "# Test Project") {
		t.Error("resolved content should include README")
	}

	// Check nav was extracted
	if !strings.Contains(c.NavYAML(), "api/ref.md") {
		t.Error("nav should contain api/ref.md")
	}

	// Check extensions extracted (snippets should be present
	// in raw list, filtered during merge)
	hasSnippets := false
	for _, ext := range c.Extensions() {
		if strings.Contains(ext, "snippets") {
			hasSnippets = true
		}
	}
	if !hasSnippets {
		t.Error("extensions should include snippets from upstream")
	}

	// Check plugins
	if len(c.Plugins()) != 2 {
		t.Errorf("got %d plugins, want 2", len(c.Plugins()))
	}

	// Check file entries
	if len(c.Files()) != 2 {
		t.Errorf("got %d file entries, want 2", len(c.Files()))
	}
}

func TestFindMkdocsYML(t *testing.T) {
	// Case 1: mkdocs.yml in parent (repo root)
	repoDir := t.TempDir()
	srcDir := filepath.Join(repoDir, "docs")
	os.MkdirAll(srcDir, 0755)
	os.WriteFile(filepath.Join(repoDir, "mkdocs.yml"),
		[]byte("site_name: Test\n"), 0644)

	c := NewConverter(srcDir, "", "", false)
	path := c.findMkdocsYML()
	if path == "" {
		t.Error("should find mkdocs.yml in parent")
	}

	// Case 2: mkdocs.yml in srcDir
	srcDir2 := t.TempDir()
	os.WriteFile(filepath.Join(srcDir2, "mkdocs.yml"),
		[]byte("site_name: Test\n"), 0644)

	c2 := NewConverter(srcDir2, "", "", false)
	path2 := c2.findMkdocsYML()
	if path2 == "" {
		t.Error("should find mkdocs.yml in srcDir")
	}
}
