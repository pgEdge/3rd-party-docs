package shared

import (
	"os"
	"path/filepath"
	"testing"
)

func TestFixBrokenLinks_ExistingTarget(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "index.md"), []byte("# Hi"), 0644)
	os.WriteFile(filepath.Join(dir, "other.md"), []byte("x"), 0644)

	content := "See [Other](other.md) for details."
	got := fixBrokenLinks(content, "index.md", dir)
	if got != content {
		t.Errorf("should not change valid link:\ngot:  %q\nwant: %q",
			got, content)
	}
}

func TestFixBrokenLinks_BrokenStripped(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "index.md"), []byte("# Hi"), 0644)

	content := "See [Guide](missing.md) for details."
	got := fixBrokenLinks(content, "index.md", dir)
	want := "See Guide for details."
	if got != want {
		t.Errorf("should strip broken link:\ngot:  %q\nwant: %q",
			got, want)
	}
}

func TestFixBrokenLinks_RelocatedFile(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "index.md"), []byte("# Hi"), 0644)
	os.WriteFile(filepath.Join(dir, "guide.md"), []byte("x"), 0644)

	// Link uses docs/ prefix but file is at root level
	content := "See [Guide](docs/guide.md#intro) for details."
	got := fixBrokenLinks(content, "index.md", dir)
	want := "See [Guide](guide.md#intro) for details."
	if got != want {
		t.Errorf("should rewrite relocated file:\ngot:  %q\nwant: %q",
			got, want)
	}
}

func TestFixBrokenLinks_READMEtoIndex(t *testing.T) {
	dir := t.TempDir()
	os.MkdirAll(filepath.Join(dir, "sub"), 0755)
	os.WriteFile(filepath.Join(dir, "index.md"), []byte("# Hi"), 0644)
	os.WriteFile(
		filepath.Join(dir, "sub", "index.md"), []byte("x"), 0644)

	content := "See [Sub](sub/README.md) page."
	got := fixBrokenLinks(content, "index.md", dir)
	want := "See [Sub](sub/index.md) page."
	if got != want {
		t.Errorf("should rewrite README→index:\ngot:  %q\nwant: %q",
			got, want)
	}
}

func TestFixBrokenLinks_AbsoluteHTML(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "config.md"), []byte("x"), 0644)
	os.WriteFile(filepath.Join(dir, "usage.md"), []byte("x"), 0644)

	content := "See [config](/config.html#option) page."
	got := fixBrokenLinks(content, "usage.md", dir)
	want := "See [config](config.md#option) page."
	if got != want {
		t.Errorf("should rewrite absolute HTML:\ngot:  %q\nwant: %q",
			got, want)
	}
}

func TestFixBrokenLinks_ExternalSkipped(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "index.md"), []byte("# Hi"), 0644)

	content := "Visit [site](https://example.com) now."
	got := fixBrokenLinks(content, "index.md", dir)
	if got != content {
		t.Errorf("should not modify external links:\ngot:  %q", got)
	}
}

func TestFixBrokenLinks_ImageSkipped(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "index.md"), []byte("# Hi"), 0644)

	content := "![logo](missing.png)"
	got := fixBrokenLinks(content, "index.md", dir)
	if got != content {
		t.Errorf("should not modify image links:\ngot:  %q", got)
	}
}

func TestFixBrokenLinks_AnchorOnly(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "index.md"), []byte("# Hi"), 0644)

	content := "See [section](#details) here."
	got := fixBrokenLinks(content, "index.md", dir)
	if got != content {
		t.Errorf("should not modify anchor links:\ngot:  %q", got)
	}
}

func TestFixBrokenLinksInDir(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "index.md"),
		[]byte("See [Guide](docs/guide.md) here."), 0644)
	os.WriteFile(filepath.Join(dir, "guide.md"),
		[]byte("# Guide"), 0644)

	err := FixBrokenLinksInDir(dir)
	if err != nil {
		t.Fatal(err)
	}

	data, _ := os.ReadFile(filepath.Join(dir, "index.md"))
	got := string(data)
	want := "See [Guide](guide.md) here."
	if got != want {
		t.Errorf("FixBrokenLinksInDir:\ngot:  %q\nwant: %q",
			got, want)
	}
}
