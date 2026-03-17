package rst

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// --- Parser tests ---

func TestParseHeading_Underline(t *testing.T) {
	root := Parse("Title\n=====\n")
	if len(root.Children) != 1 {
		t.Fatalf("expected 1 child, got %d", len(root.Children))
	}
	h := root.Children[0]
	if h.Type != HeadingNode {
		t.Fatalf("expected HeadingNode, got %d", h.Type)
	}
	if h.Text != "Title" {
		t.Errorf("expected 'Title', got %q", h.Text)
	}
	if h.Level != 1 {
		t.Errorf("expected level 1, got %d", h.Level)
	}
}

func TestParseHeading_Overline(t *testing.T) {
	root := Parse("*****\nTitle\n*****\n")
	if len(root.Children) != 1 {
		t.Fatalf("expected 1 child, got %d", len(root.Children))
	}
	h := root.Children[0]
	if h.Type != HeadingNode {
		t.Fatalf("expected HeadingNode, got %d", h.Type)
	}
	if h.Text != "Title" {
		t.Errorf("expected 'Title', got %q", h.Text)
	}
}

func TestParseHeading_LevelOrder(t *testing.T) {
	root := Parse("H1\n***\n\nH2\n===\n\nH3\n---\n")
	if len(root.Children) != 3 {
		t.Fatalf("expected 3 children, got %d", len(root.Children))
	}
	for i, expected := range []int{1, 2, 3} {
		if root.Children[i].Level != expected {
			t.Errorf("child %d: expected level %d, got %d",
				i, expected, root.Children[i].Level)
		}
	}
}

func TestParseParagraph(t *testing.T) {
	root := Parse("This is a paragraph\nwith two lines.\n")
	if len(root.Children) != 1 {
		t.Fatalf("expected 1 child, got %d", len(root.Children))
	}
	p := root.Children[0]
	if p.Type != ParagraphNode {
		t.Fatalf("expected ParagraphNode, got %d", p.Type)
	}
	if !strings.Contains(p.Text, "paragraph") {
		t.Errorf("expected paragraph text, got %q", p.Text)
	}
}

func TestParseLabel(t *testing.T) {
	root := Parse(".. _my_label:\n\nSome text\n")
	if len(root.Children) < 1 {
		t.Fatalf("expected at least 1 child, got %d",
			len(root.Children))
	}
	l := root.Children[0]
	if l.Type != LabelNode {
		t.Fatalf("expected LabelNode, got %d", l.Type)
	}
	if l.Label != "my_label" {
		t.Errorf("expected 'my_label', got %q", l.Label)
	}
}

func TestParseDirective(t *testing.T) {
	root := Parse(".. image:: images/test.png\n    :alt: Test\n    :align: center\n")
	if len(root.Children) != 1 {
		t.Fatalf("expected 1 child, got %d", len(root.Children))
	}
	d := root.Children[0]
	if d.Type != DirectiveNode {
		t.Fatalf("expected DirectiveNode, got %d", d.Type)
	}
	if d.DirectiveName != "image" {
		t.Errorf("expected 'image', got %q", d.DirectiveName)
	}
	if d.DirectiveArg != "images/test.png" {
		t.Errorf("expected 'images/test.png', got %q", d.DirectiveArg)
	}
	if d.Options["alt"] != "Test" {
		t.Errorf("expected alt 'Test', got %q", d.Options["alt"])
	}
	if d.Options["align"] != "center" {
		t.Errorf("expected align 'center', got %q", d.Options["align"])
	}
}

func TestParseDirectiveWithBody(t *testing.T) {
	rst := ".. code-block:: python\n\n   print('hello')\n   print('world')\n"
	root := Parse(rst)
	if len(root.Children) != 1 {
		t.Fatalf("expected 1 child, got %d", len(root.Children))
	}
	d := root.Children[0]
	if d.DirectiveName != "code-block" {
		t.Errorf("expected 'code-block', got %q", d.DirectiveName)
	}
	if d.DirectiveArg != "python" {
		t.Errorf("expected 'python', got %q", d.DirectiveArg)
	}
	if !strings.Contains(d.Body, "print('hello')") {
		t.Errorf("body missing print('hello'): %q", d.Body)
	}
}

func TestParseBulletList(t *testing.T) {
	rst := "* Item 1\n* Item 2\n* Item 3\n"
	root := Parse(rst)
	if len(root.Children) != 1 {
		t.Fatalf("expected 1 child, got %d", len(root.Children))
	}
	list := root.Children[0]
	if list.Type != BulletListNode {
		t.Fatalf("expected BulletListNode, got %d", list.Type)
	}
	if len(list.Children) != 3 {
		t.Errorf("expected 3 items, got %d", len(list.Children))
	}
}

func TestParseBulletList_Continuation(t *testing.T) {
	rst := "* Item 1 continues\n  on the next line\n* Item 2\n"
	root := Parse(rst)
	list := root.Children[0]
	if len(list.Children) != 2 {
		t.Fatalf("expected 2 items, got %d", len(list.Children))
	}
	if !strings.Contains(list.Children[0].Text, "continues") {
		t.Errorf("first item missing continuation: %q",
			list.Children[0].Text)
	}
}

func TestParseEnumList(t *testing.T) {
	rst := "1. First\n2. Second\n3. Third\n"
	root := Parse(rst)
	if len(root.Children) != 1 {
		t.Fatalf("expected 1 child, got %d", len(root.Children))
	}
	list := root.Children[0]
	if list.Type != EnumListNode {
		t.Fatalf("expected EnumListNode, got %d", list.Type)
	}
	if len(list.Children) != 3 {
		t.Errorf("expected 3 items, got %d", len(list.Children))
	}
}

func TestParseToctree(t *testing.T) {
	rst := ".. toctree::\n   :maxdepth: 2\n\n   getting_started\n   installation\n"
	root := Parse(rst)
	if len(root.Children) != 1 {
		t.Fatalf("expected 1 child, got %d", len(root.Children))
	}
	d := root.Children[0]
	if d.DirectiveName != "toctree" {
		t.Errorf("expected 'toctree', got %q", d.DirectiveName)
	}
	if !strings.Contains(d.Body, "getting_started") {
		t.Errorf("body missing 'getting_started': %q", d.Body)
	}
	if !strings.Contains(d.Body, "installation") {
		t.Errorf("body missing 'installation': %q", d.Body)
	}
}

func TestParseGridTable(t *testing.T) {
	rst := `+-------+-------+
| A     | B     |
+=======+=======+
| 1     | 2     |
+-------+-------+
| 3     | 4     |
+-------+-------+
`
	root := Parse(rst)
	if len(root.Children) != 1 {
		t.Fatalf("expected 1 child, got %d", len(root.Children))
	}
	table := root.Children[0]
	if table.Type != GridTableNode {
		t.Fatalf("expected GridTableNode, got %d", table.Type)
	}
	if !table.TableHeader {
		t.Error("expected table to have header")
	}
	if len(table.TableRows) != 3 {
		t.Errorf("expected 3 rows, got %d", len(table.TableRows))
	}
}

func TestParseLiteralBlock(t *testing.T) {
	rst := "Example::\n\n    code line 1\n    code line 2\n\nAfter.\n"
	root := Parse(rst)
	if len(root.Children) < 2 {
		t.Fatalf("expected at least 2 children, got %d",
			len(root.Children))
	}
	// First should be paragraph "Example:"
	// Second should be literal block
	found := false
	for _, child := range root.Children {
		if child.Type == LiteralBlockNode {
			found = true
			if !strings.Contains(child.Text, "code line 1") {
				t.Errorf("literal block missing content: %q",
					child.Text)
			}
		}
	}
	if !found {
		t.Error("no LiteralBlockNode found")
	}
}

func TestParseSubstitutionDef(t *testing.T) {
	rst := ".. |icon| image:: images/icon.png\n"
	root := Parse(rst)
	if len(root.Children) != 1 {
		t.Fatalf("expected 1 child, got %d", len(root.Children))
	}
	sub := root.Children[0]
	if sub.Type != SubstitutionDefNode {
		t.Fatalf("expected SubstitutionDefNode, got %d", sub.Type)
	}
	if sub.SubstitutionName != "icon" {
		t.Errorf("expected 'icon', got %q", sub.SubstitutionName)
	}
	if sub.DirectiveName != "image" {
		t.Errorf("expected 'image', got %q", sub.DirectiveName)
	}
}

func TestParseAdmonition(t *testing.T) {
	rst := ".. note:: This is a note\n    body.\n"
	root := Parse(rst)
	if len(root.Children) != 1 {
		t.Fatalf("expected 1 child, got %d", len(root.Children))
	}
	d := root.Children[0]
	if d.DirectiveName != "note" {
		t.Errorf("expected 'note', got %q", d.DirectiveName)
	}
}

func TestParseComment(t *testing.T) {
	rst := ".. This is a comment\n   spanning two lines\n\nParagraph.\n"
	root := Parse(rst)
	foundComment := false
	foundPara := false
	for _, child := range root.Children {
		if child.Type == CommentNode {
			foundComment = true
		}
		if child.Type == ParagraphNode {
			foundPara = true
		}
	}
	if !foundComment {
		t.Error("no comment node found")
	}
	if !foundPara {
		t.Error("no paragraph node found after comment")
	}
}

// --- Inline conversion tests ---

func TestConvertInline_Bold(t *testing.T) {
	result := ConvertInline("This is **bold** text", nil, nil, "", nil, nil)
	if result != "This is **bold** text" {
		t.Errorf("unexpected result: %q", result)
	}
}

func TestConvertInline_Italic(t *testing.T) {
	result := ConvertInline("This is *italic* text", nil, nil, "", nil, nil)
	if result != "This is *italic* text" {
		t.Errorf("unexpected result: %q", result)
	}
}

func TestConvertInline_Literal(t *testing.T) {
	result := ConvertInline("Use ``config.py`` file", nil, nil, "", nil, nil)
	if result != "Use `config.py` file" {
		t.Errorf("unexpected result: %q", result)
	}
}

func TestConvertInline_ExternalLink(t *testing.T) {
	result := ConvertInline(
		"`PostgreSQL <https://www.postgresql.org/>`_",
		nil, nil, "", nil, nil)
	if result != "[PostgreSQL](https://www.postgresql.org/)" {
		t.Errorf("unexpected result: %q", result)
	}
}

func TestConvertInline_Ref(t *testing.T) {
	labels := map[string]labelInfo{
		"config_py": {
			File:   "config_py.md",
			Anchor: "config_py",
			Title:  "Configuration",
		},
	}
	result := ConvertInline(
		":ref:`config.py <config_py>`",
		labels, nil, "mfa.md", nil, nil)
	if !strings.Contains(result, "[config.py]") {
		t.Errorf("expected link text 'config.py', got %q", result)
	}
	if !strings.Contains(result, "config_py.md") {
		t.Errorf("expected link to config_py.md, got %q", result)
	}
}

func TestConvertInline_IndexRole(t *testing.T) {
	result := ConvertInline(
		"`Getting Started`:index:", nil, nil, "", nil, nil)
	if result != "Getting Started" {
		t.Errorf("expected 'Getting Started', got %q", result)
	}
}

func TestConvertInline_Substitution(t *testing.T) {
	subs := map[string]*Node{
		"icon": {
			DirectiveName: "image",
			DirectiveArg:  "images/icon.png",
			Options:       map[string]string{"alt": "Icon"},
		},
	}
	result := ConvertInline("Click |icon| here", nil, nil, "", subs, nil)
	if !strings.Contains(result, "<img") {
		t.Errorf("expected img tag, got %q", result)
	}
	if !strings.Contains(result, "icon.png") {
		t.Errorf("expected icon.png, got %q", result)
	}
}

// --- Clean title tests ---

func TestCleanTitle_IndexMarkup(t *testing.T) {
	result := cleanTitle("`Getting Started`:index:")
	if result != "Getting Started" {
		t.Errorf("expected 'Getting Started', got %q", result)
	}
}

func TestCleanTitle_Plain(t *testing.T) {
	result := cleanTitle("Plain Title")
	if result != "Plain Title" {
		t.Errorf("expected 'Plain Title', got %q", result)
	}
}

// --- Integration test with temp directory ---

func TestConvertFile_Simple(t *testing.T) {
	initDirectiveHandlers()

	srcDir := t.TempDir()
	outDir := t.TempDir()

	// Write a simple RST file
	indexRST := `*****
Title
*****

Welcome to the docs.

.. toctree::
   :maxdepth: 2

   page1
`
	page1RST := `.. _page1_label:

*****
Page1
*****

This is page 1 with a :ref:` + "`link <page1_label>`" + `.

.. image:: images/test.png
    :alt: Test Image
`

	os.WriteFile(filepath.Join(srcDir, "index.rst"),
		[]byte(indexRST), 0644)
	os.WriteFile(filepath.Join(srcDir, "page1.rst"),
		[]byte(page1RST), 0644)

	// Create images dir
	os.MkdirAll(filepath.Join(srcDir, "images"), 0755)
	os.WriteFile(filepath.Join(srcDir, "images", "test.png"),
		[]byte("PNG"), 0644)

	converter := NewConverter(srcDir, outDir, "1.0", "", "", false)
	if err := converter.Convert(); err != nil {
		t.Fatalf("conversion failed: %v", err)
	}

	// Check output files exist
	indexMD, err := os.ReadFile(filepath.Join(outDir, "index.md"))
	if err != nil {
		t.Fatalf("index.md not created: %v", err)
	}
	if !strings.Contains(string(indexMD), "# Title") {
		t.Errorf("index.md missing title: %s", indexMD)
	}

	page1MD, err := os.ReadFile(filepath.Join(outDir, "page1.md"))
	if err != nil {
		t.Fatalf("page1.md not created: %v", err)
	}
	if !strings.Contains(string(page1MD), "# Page1") {
		t.Errorf("page1.md missing title: %s", page1MD)
	}
	if !strings.Contains(string(page1MD), "![Test Image]") {
		t.Errorf("page1.md missing image: %s", page1MD)
	}

	// Check image copied
	if _, err := os.Stat(filepath.Join(outDir, "images", "test.png")); os.IsNotExist(err) {
		t.Error("image not copied to output")
	}

	// Check files list
	files := converter.Files()
	if len(files) < 2 {
		t.Errorf("expected at least 2 files, got %d", len(files))
	}
}

func TestResolveLink_SameFile(t *testing.T) {
	result := resolveLink("page.md", "page.md", "anchor")
	if result != "#anchor" {
		t.Errorf("expected '#anchor', got %q", result)
	}
}

func TestResolveLink_DifferentFile(t *testing.T) {
	result := resolveLink("page1.md", "page2.md", "anchor")
	if result != "page2.md#anchor" {
		t.Errorf("expected 'page2.md#anchor', got %q", result)
	}
}

func TestResolveLink_NoAnchor(t *testing.T) {
	result := resolveLink("page1.md", "page2.md", "")
	if result != "page2.md" {
		t.Errorf("expected 'page2.md', got %q", result)
	}
}

func TestParseCSVRow(t *testing.T) {
	row := parseCSVRow(`"Name","Description"`)
	if len(row) != 2 {
		t.Fatalf("expected 2 cells, got %d", len(row))
	}
	if row[0] != "Name" {
		t.Errorf("expected 'Name', got %q", row[0])
	}
	if row[1] != "Description" {
		t.Errorf("expected 'Description', got %q", row[1])
	}
}

func TestParseCSVRow_Escaped(t *testing.T) {
	row := parseCSVRow(`"He said ""hello""",value`)
	if len(row) != 2 {
		t.Fatalf("expected 2 cells, got %d", len(row))
	}
	if row[0] != `He said "hello"` {
		t.Errorf("unexpected first cell: %q", row[0])
	}
}

func TestExtractRowCells(t *testing.T) {
	// Use a real grid table example parsed through the full pipeline
	rst := `+-------+-------+
| hello | world |
+-------+-------+
`
	root := Parse(rst)
	if len(root.Children) != 1 || root.Children[0].Type != GridTableNode {
		t.Fatal("expected GridTableNode")
	}
	table := root.Children[0]
	if len(table.TableRows) < 1 {
		t.Fatal("expected at least 1 row")
	}
	row := table.TableRows[0]
	if len(row) != 2 {
		t.Fatalf("expected 2 cells, got %d", len(row))
	}
	if row[0] != "hello" {
		t.Errorf("expected 'hello', got %q", row[0])
	}
	if row[1] != "world" {
		t.Errorf("expected 'world', got %q", row[1])
	}
}

func TestIsDecorationLine(t *testing.T) {
	tests := []struct {
		line string
		ok   bool
		ch   rune
	}{
		{"=====", true, '='},
		{"*****", true, '*'},
		{"-----", true, '-'},
		{"ab", false, 0},
		{"=", false, 0},
		{"==", false, 0},
		{"===", true, '='},
		{"abc", false, 0},
	}
	for _, tt := range tests {
		ch, ok := isDecorationLine(tt.line)
		if ok != tt.ok {
			t.Errorf("isDecorationLine(%q): got ok=%v, want %v",
				tt.line, ok, tt.ok)
		}
		if ok && ch != tt.ch {
			t.Errorf("isDecorationLine(%q): got ch=%c, want %c",
				tt.line, ch, tt.ch)
		}
	}
}

func TestParseFieldList(t *testing.T) {
	rst := ":Name: John\n:Age: 30\n"
	root := Parse(rst)
	found := false
	for _, child := range root.Children {
		if child.Type == FieldListNode {
			found = true
			if len(child.Children) != 2 {
				t.Errorf("expected 2 fields, got %d",
					len(child.Children))
			}
		}
	}
	if !found {
		t.Error("no FieldListNode found")
	}
}

func TestLineBlock(t *testing.T) {
	rst := "| Line one\n| Line two\n| Line three\n"
	root := Parse(rst)
	found := false
	for _, child := range root.Children {
		if child.Type == LineBlockNode {
			found = true
			if !strings.Contains(child.Text, "Line one") {
				t.Errorf("missing 'Line one': %q", child.Text)
			}
		}
	}
	if !found {
		t.Error("no LineBlockNode found")
	}
}
