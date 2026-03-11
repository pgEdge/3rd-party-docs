//-------------------------------------------------------------------------
//
// pgEdge PostgreSQL Docs
//
// Copyright (c) 2026, pgEdge, Inc.
// This software is released under The PostgreSQL License
//
//-------------------------------------------------------------------------

package validate

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func createFile(t *testing.T, dir, name, content string) {
	t.Helper()
	path := filepath.Join(dir, name)
	err := os.MkdirAll(filepath.Dir(path), 0755)
	if err != nil {
		t.Fatal(err)
	}
	err = os.WriteFile(path, []byte(content), 0644)
	if err != nil {
		t.Fatal(err)
	}
}

func TestValidateDir_ValidLinks(t *testing.T) {
	dir := t.TempDir()
	createFile(t, dir, "index.md", `# Welcome

See [setup](setup.md) for details.
`)
	createFile(t, dir, "setup.md", `# Setup

Go back to [home](index.md).
`)

	result, err := ValidateDir(dir)
	if err != nil {
		t.Fatal(err)
	}
	if len(result.BrokenLinks) != 0 {
		t.Errorf("expected no broken links, got %v",
			result.BrokenLinks)
	}
	if len(result.MissingAnchors) != 0 {
		t.Errorf("expected no missing anchors, got %v",
			result.MissingAnchors)
	}
}

func TestValidateDir_BrokenFileLink(t *testing.T) {
	dir := t.TempDir()
	createFile(t, dir, "index.md", `# Home

See [missing](nonexistent.md) page.
`)

	result, err := ValidateDir(dir)
	if err != nil {
		t.Fatal(err)
	}
	if len(result.BrokenLinks) != 1 {
		t.Fatalf("expected 1 broken link, got %d",
			len(result.BrokenLinks))
	}
	issue := result.BrokenLinks[0]
	if issue.File != "index.md" {
		t.Errorf("file = %q, want 'index.md'", issue.File)
	}
	if issue.Line != 3 {
		t.Errorf("line = %d, want 3", issue.Line)
	}
}

func TestValidateDir_ValidAnchorLink(t *testing.T) {
	dir := t.TempDir()
	createFile(t, dir, "page.md", `# Introduction

Some text.

## Details

See [intro](#introduction) above.
`)

	result, err := ValidateDir(dir)
	if err != nil {
		t.Fatal(err)
	}
	if len(result.BrokenLinks) != 0 {
		t.Errorf("unexpected broken links: %v", result.BrokenLinks)
	}
	if len(result.MissingAnchors) != 0 {
		t.Errorf("unexpected missing anchors: %v",
			result.MissingAnchors)
	}
}

func TestValidateDir_BrokenAnchorLink(t *testing.T) {
	dir := t.TempDir()
	createFile(t, dir, "page.md", `# Introduction

See [bad link](#nonexistent-section).
`)

	result, err := ValidateDir(dir)
	if err != nil {
		t.Fatal(err)
	}
	if len(result.MissingAnchors) != 1 {
		t.Fatalf("expected 1 missing anchor, got %d",
			len(result.MissingAnchors))
	}
	if result.MissingAnchors[0].Line != 3 {
		t.Errorf("line = %d, want 3",
			result.MissingAnchors[0].Line)
	}
}

func TestValidateDir_ExplicitAnchorID(t *testing.T) {
	dir := t.TempDir()
	createFile(t, dir, "page.md", `# Title { #my-custom-id }

See [link](#my-custom-id).
`)

	result, err := ValidateDir(dir)
	if err != nil {
		t.Fatal(err)
	}
	if len(result.MissingAnchors) != 0 {
		t.Errorf("unexpected missing anchors: %v",
			result.MissingAnchors)
	}
}

func TestValidateDir_HTMLAnchorTag(t *testing.T) {
	dir := t.TempDir()
	createFile(t, dir, "page.md", `# Title

<a id="custom-anchor"></a>

See [link](#custom-anchor).
`)

	result, err := ValidateDir(dir)
	if err != nil {
		t.Fatal(err)
	}
	if len(result.MissingAnchors) != 0 {
		t.Errorf("unexpected missing anchors: %v",
			result.MissingAnchors)
	}
}

func TestValidateDir_CrossFileAnchor(t *testing.T) {
	dir := t.TempDir()
	createFile(t, dir, "a.md", `# Page A

See [section](b.md#details).
`)
	createFile(t, dir, "b.md", `# Page B

## Details

Some content.
`)

	result, err := ValidateDir(dir)
	if err != nil {
		t.Fatal(err)
	}
	if len(result.BrokenLinks) != 0 {
		t.Errorf("unexpected broken links: %v", result.BrokenLinks)
	}
	if len(result.MissingAnchors) != 0 {
		t.Errorf("unexpected missing anchors: %v",
			result.MissingAnchors)
	}
}

func TestValidateDir_CrossFileBrokenAnchor(t *testing.T) {
	dir := t.TempDir()
	createFile(t, dir, "a.md", `# Page A

See [bad](b.md#nonexistent).
`)
	createFile(t, dir, "b.md", `# Page B

Content.
`)

	result, err := ValidateDir(dir)
	if err != nil {
		t.Fatal(err)
	}
	if len(result.MissingAnchors) != 1 {
		t.Fatalf("expected 1 missing anchor, got %d",
			len(result.MissingAnchors))
	}
}

func TestValidateDir_ExternalLinksSkipped(t *testing.T) {
	dir := t.TempDir()
	createFile(t, dir, "page.md", `# Links

- [http](http://example.com)
- [https](https://example.com)
- [mail](mailto:test@example.com)
`)

	result, err := ValidateDir(dir)
	if err != nil {
		t.Fatal(err)
	}
	if len(result.BrokenLinks) != 0 {
		t.Errorf("external links should be skipped, got %v",
			result.BrokenLinks)
	}
}

func TestValidateDir_SubdirectoryLinks(t *testing.T) {
	dir := t.TempDir()
	createFile(t, dir, "index.md", `# Home

See [sub](sub/page.md).
`)
	createFile(t, dir, "sub/page.md", `# Sub Page

Go [back](../index.md).
`)

	result, err := ValidateDir(dir)
	if err != nil {
		t.Fatal(err)
	}
	if len(result.BrokenLinks) != 0 {
		t.Errorf("unexpected broken links: %v", result.BrokenLinks)
	}
}

func TestValidateDir_EmptyDirectory(t *testing.T) {
	dir := t.TempDir()
	result, err := ValidateDir(dir)
	if err != nil {
		t.Fatal(err)
	}
	if len(result.BrokenLinks) != 0 {
		t.Error("expected no broken links for empty dir")
	}
}

func TestValidateDir_NonExistentDir(t *testing.T) {
	_, err := ValidateDir("/nonexistent/path/docs")
	if err == nil {
		t.Error("expected error for nonexistent directory")
	}
}

func TestValidateDir_NonMdFilesIgnored(t *testing.T) {
	dir := t.TempDir()
	createFile(t, dir, "page.md", `# Page

Content.
`)
	createFile(t, dir, "style.css", `body { color: red; }`)
	createFile(t, dir, "image.png", `not really a png`)

	result, err := ValidateDir(dir)
	if err != nil {
		t.Fatal(err)
	}
	if len(result.BrokenLinks) != 0 {
		t.Error("non-md files should be ignored")
	}
}

func TestValidateDir_MultipleLinksOnOneLine(t *testing.T) {
	dir := t.TempDir()
	createFile(t, dir, "page.md", `# Page

See [a](a.md) and [b](missing.md) here.
`)
	createFile(t, dir, "a.md", `# A
`)

	result, err := ValidateDir(dir)
	if err != nil {
		t.Fatal(err)
	}
	if len(result.BrokenLinks) != 1 {
		t.Fatalf("expected 1 broken link, got %d",
			len(result.BrokenLinks))
	}
}

func TestValidateDir_FileOnlyLink(t *testing.T) {
	// Link to file without anchor should be valid if file exists
	dir := t.TempDir()
	createFile(t, dir, "a.md", `# Page A

See [b](b.md).
`)
	createFile(t, dir, "b.md", `# Page B
`)

	result, err := ValidateDir(dir)
	if err != nil {
		t.Fatal(err)
	}
	if len(result.BrokenLinks) != 0 {
		t.Errorf("unexpected broken links: %v", result.BrokenLinks)
	}
	if len(result.MissingAnchors) != 0 {
		t.Errorf("unexpected missing anchors: %v",
			result.MissingAnchors)
	}
}

// --- headingToAnchor tests ---

func TestHeadingToAnchor_PlainText(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"Introduction", "introduction"},
		{"Getting Started", "getting-started"},
		{"Chapter 1", "chapter-1"},
	}
	for _, tt := range tests {
		got := headingToAnchor(tt.input)
		if got != tt.want {
			t.Errorf("headingToAnchor(%q) = %q, want %q",
				tt.input, got, tt.want)
		}
	}
}

func TestHeadingToAnchor_InlineMarkup(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"The `SELECT` Statement", "the-select-statement"},
		{"**Bold** Title", "bold-title"},
		{"*Italic* Text", "italic-text"},
		{"_Underlined_ Word", "underlined-word"},
	}
	for _, tt := range tests {
		got := headingToAnchor(tt.input)
		if got != tt.want {
			t.Errorf("headingToAnchor(%q) = %q, want %q",
				tt.input, got, tt.want)
		}
	}
}

func TestHeadingToAnchor_SpecialChars(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"What's New?", "whats-new"},
		{"C++ Support", "c-support"},
		{"A & B", "a-b"},
		{"Item (deprecated)", "item-deprecated"},
		{"Version 1.2.3", "version-123"},
		{"Hello/World", "helloworld"},
	}
	for _, tt := range tests {
		got := headingToAnchor(tt.input)
		if got != tt.want {
			t.Errorf("headingToAnchor(%q) = %q, want %q",
				tt.input, got, tt.want)
		}
	}
}

func TestHeadingToAnchor_MultipleSpaces(t *testing.T) {
	got := headingToAnchor("Too   Many   Spaces")
	if got != "too-many-spaces" {
		t.Errorf("got %q, want 'too-many-spaces'", got)
	}
}

func TestHeadingToAnchor_LeadingTrailingSpecial(t *testing.T) {
	got := headingToAnchor(" -Hello- ")
	if got != "hello" {
		t.Errorf("got %q, want 'hello'", got)
	}
}

func TestHeadingToAnchor_Empty(t *testing.T) {
	got := headingToAnchor("")
	if got != "" {
		t.Errorf("got %q, want ''", got)
	}
}

func TestHeadingToAnchor_Numbers(t *testing.T) {
	got := headingToAnchor("123")
	if got != "123" {
		t.Errorf("got %q, want '123'", got)
	}
}

func TestHeadingToAnchor_Hyphens(t *testing.T) {
	got := headingToAnchor("A--B")
	if got != "a-b" {
		t.Errorf("got %q, want 'a-b'", got)
	}
}

func TestValidateDir_UnreadableFile(t *testing.T) {
	dir := t.TempDir()
	createFile(t, dir, "good.md", `# Good
`)
	// Create an unreadable .md file to trigger ReadFile error
	unreadable := filepath.Join(dir, "bad.md")
	err := os.WriteFile(unreadable, []byte(`# Bad`), 0000)
	if err != nil {
		t.Fatal(err)
	}

	_, err = ValidateDir(dir)
	if err == nil {
		t.Error("expected error for unreadable file")
	}
}

func TestCheckLinks_WalkError(t *testing.T) {
	dir := t.TempDir()
	createFile(t, dir, "page.md", `# Page

Content.
`)

	// Create a subdirectory with a file, then make it
	// unreadable so the walk callback receives an error.
	sub := filepath.Join(dir, "sub")
	err := os.Mkdir(sub, 0755)
	if err != nil {
		t.Fatal(err)
	}
	createFile(t, dir, "sub/nested.md", `# Nested
`)
	err = os.Chmod(sub, 0000)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		os.Chmod(sub, 0755)
	})

	anchors := map[string]map[string]bool{
		"page.md": {"page": true},
	}
	_, err = checkLinks(dir, anchors)
	if err == nil {
		t.Error("expected error from checkLinks walk")
	}
}

func TestValidateDir_CheckLinksWalkError(t *testing.T) {
	dir := t.TempDir()
	createFile(t, dir, "page.md", `# Page

Content.
`)
	sub := filepath.Join(dir, "sub")
	err := os.Mkdir(sub, 0755)
	if err != nil {
		t.Fatal(err)
	}
	createFile(t, dir, "sub/nested.md", `# Nested
`)

	// Use the test hook to make a subdirectory unreadable
	// after collectAnchors succeeds but before checkLinks
	// runs, ensuring the "checking links" error path in
	// ValidateDir is exercised.
	testHookAfterCollect = func() {
		os.Chmod(sub, 0000)
	}
	t.Cleanup(func() {
		testHookAfterCollect = nil
		os.Chmod(sub, 0755)
	})

	_, err = ValidateDir(dir)
	if err == nil {
		t.Fatal("expected error from ValidateDir")
	}
	if !strings.Contains(err.Error(), "checking links") {
		t.Errorf("expected 'checking links' error, got: %v",
			err)
	}
}

func TestHeadingToAnchor_Underscores(t *testing.T) {
	// Underscores are treated as inline markup (like _italic_)
	// and get stripped by the regex
	got := headingToAnchor("my_func")
	if got != "myfunc" {
		t.Errorf("got %q, want 'myfunc'", got)
	}
}
