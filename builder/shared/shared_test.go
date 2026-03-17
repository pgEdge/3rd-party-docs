package shared

import (
	"strings"
	"testing"
)

func TestSlugify(t *testing.T) {
	tests := []struct {
		in, want string
	}{
		{"Hello World", "hello-world"},
		{"foo/bar.baz", "foo-bar-baz"},
		{"  spaces  ", "spaces"},
		{"UPPER", "upper"},
		{"a--b", "a-b"},
		{"a___b", "a___b"},
		{"", ""},
		{"123", "123"},
		{"-lead-trail-", "lead-trail"},
	}
	for _, tt := range tests {
		got := Slugify(tt.in)
		if got != tt.want {
			t.Errorf("Slugify(%q) = %q, want %q",
				tt.in, got, tt.want)
		}
	}
}

func TestMarkdownWriter_WriteString(t *testing.T) {
	w := NewMarkdownWriter()
	w.WriteString("hello")
	if w.String() != "hello" {
		t.Errorf("got %q", w.String())
	}
	if w.IsAtLineStart() {
		t.Error("should not be at line start")
	}
}

func TestMarkdownWriter_Write_Indent(t *testing.T) {
	w := NewMarkdownWriter()
	w.PushIndent("  ")
	w.Write("line1\nline2")
	got := w.String()
	if !strings.Contains(got, "  line2") {
		t.Errorf("expected indented line2, got %q", got)
	}
	w.PopIndent("  ")
}

func TestMarkdownWriter_BlankLine(t *testing.T) {
	w := NewMarkdownWriter()
	w.WriteString("text")
	w.BlankLine()
	w.BlankLine() // should not double
	got := w.String()
	if strings.Count(got, "\n") != 2 {
		t.Errorf("expected 2 newlines, got %q", got)
	}
}

func TestMarkdownWriter_BlankLine_Suppressed(t *testing.T) {
	w := NewMarkdownWriter()
	w.SetSuppressNewlines(true)
	w.WriteString("text")
	w.BlankLine()
	if strings.Contains(w.String(), "\n") {
		t.Error("blank line should be suppressed")
	}
}

func TestMarkdownWriter_EnsureNewline(t *testing.T) {
	w := NewMarkdownWriter()
	w.WriteString("text")
	w.EnsureNewline()
	w.EnsureNewline() // should not double
	if w.String() != "text\n" {
		t.Errorf("got %q", w.String())
	}
}

func TestMarkdownWriter_CodeBlock(t *testing.T) {
	w := NewMarkdownWriter()
	w.StartCodeBlock("python")
	w.WriteString("print('hi')")
	w.EndCodeBlock()
	got := w.String()
	if !strings.Contains(got, "```python") {
		t.Errorf("missing python fence: %q", got)
	}
	if !strings.Contains(got, "print('hi')") {
		t.Errorf("missing code: %q", got)
	}
	if w.InCodeBlock() {
		t.Error("should not be in code block after end")
	}
}

func TestMarkdownWriter_CodeBlock_NoLang(t *testing.T) {
	w := NewMarkdownWriter()
	w.StartCodeBlock("")
	w.WriteString("code")
	w.EndCodeBlock()
	got := w.String()
	if !strings.Contains(got, "```\n") {
		t.Errorf("expected bare fence: %q", got)
	}
}

func TestMarkdownWriter_Heading(t *testing.T) {
	w := NewMarkdownWriter()
	w.Heading(2, "Title", "")
	got := w.String()
	if !strings.Contains(got, "## Title") {
		t.Errorf("expected ## heading: %q", got)
	}
}

func TestMarkdownWriter_Admonition(t *testing.T) {
	w := NewMarkdownWriter()
	w.Admonition("note")
	got := w.String()
	if !strings.Contains(got, "!!! note") {
		t.Errorf("expected admonition: %q", got)
	}
}

func TestMarkdownWriter_Len(t *testing.T) {
	w := NewMarkdownWriter()
	w.WriteString("abc")
	if w.Len() != 3 {
		t.Errorf("expected 3, got %d", w.Len())
	}
}

func TestMarkdownWriter_Newline(t *testing.T) {
	w := NewMarkdownWriter()
	w.WriteString("a")
	w.Newline()
	if !w.IsAtLineStart() {
		t.Error("should be at line start after newline")
	}
}

func TestMarkdownWriter_WriteEmpty(t *testing.T) {
	w := NewMarkdownWriter()
	w.WriteString("")
	w.Write("")
	if w.Len() != 0 {
		t.Error("empty writes should produce no output")
	}
}

func TestMarkdownWriter_IndentInCodeBlock(t *testing.T) {
	w := NewMarkdownWriter()
	w.PushIndent("  ")
	w.StartCodeBlock("")
	w.Write("code")
	// Indent should NOT apply inside code blocks
	got := w.String()
	if strings.Contains(got, "  code") {
		t.Errorf("indent should not apply in code: %q", got)
	}
	w.EndCodeBlock()
	w.PopIndent("  ")
}
