package rst

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/pgEdge/postgresql-docs/builder/shared"
)

// helper to convert RST text to Markdown via the full pipeline.
func rstToMD(t *testing.T, rst string) string {
	t.Helper()
	initDirectiveHandlers()
	root := Parse(rst)
	ctx := &ConvertContext{
		FileMap:       map[string]string{},
		LabelMap:      map[string]labelInfo{},
		Substitutions: map[string]*Node{},
		CurrentFile:   "test.md",
	}
	w := shared.NewMarkdownWriter()
	for _, child := range root.Children {
		convertNode(ctx, child, w)
	}
	return w.String()
}

// --- convertNode dispatch tests ---

func TestConvert_Heading(t *testing.T) {
	md := rstToMD(t, "Title\n=====\n")
	if !strings.Contains(md, "# Title") {
		t.Errorf("expected heading: %q", md)
	}
}

func TestConvert_Paragraph(t *testing.T) {
	md := rstToMD(t, "Hello world.\n")
	if !strings.Contains(md, "Hello world.") {
		t.Errorf("expected paragraph: %q", md)
	}
}

func TestConvert_BulletList(t *testing.T) {
	md := rstToMD(t, "* Item A\n\n* Item B\n")
	if !strings.Contains(md, "- Item A") {
		t.Errorf("missing bullet A: %q", md)
	}
	if !strings.Contains(md, "- Item B") {
		t.Errorf("missing bullet B: %q", md)
	}
}

func TestConvert_EnumList(t *testing.T) {
	md := rstToMD(t, "1. First\n2. Second\n")
	if !strings.Contains(md, "1. First") {
		t.Errorf("missing item 1: %q", md)
	}
	if !strings.Contains(md, "2. Second") {
		t.Errorf("missing item 2: %q", md)
	}
}

func TestConvert_LiteralBlock(t *testing.T) {
	md := rstToMD(t, "Example::\n\n    code here\n")
	if !strings.Contains(md, "```") {
		t.Errorf("missing code fence: %q", md)
	}
	if !strings.Contains(md, "code here") {
		t.Errorf("missing code content: %q", md)
	}
}

func TestConvert_Label(t *testing.T) {
	md := rstToMD(t, ".. _my_label:\n\nText\n")
	if !strings.Contains(md, `<a id="my_label"></a>`) {
		t.Errorf("missing anchor: %q", md)
	}
}

func TestConvert_Transition(t *testing.T) {
	md := rstToMD(t, "Para.\n\n----\n\nAfter.\n")
	if !strings.Contains(md, "---") {
		t.Errorf("missing transition: %q", md)
	}
}

func TestConvert_FieldList(t *testing.T) {
	md := rstToMD(t, ":Name: John\n:Age: 30\n")
	if !strings.Contains(md, "**Name**") {
		t.Errorf("missing field name: %q", md)
	}
	if !strings.Contains(md, "John") {
		t.Errorf("missing field value: %q", md)
	}
}

func TestConvert_LineBlock(t *testing.T) {
	md := rstToMD(t, "| Line one\n| Line two\n")
	if !strings.Contains(md, "Line one<br>") {
		t.Errorf("missing line block: %q", md)
	}
}

func TestConvert_Comment(t *testing.T) {
	md := rstToMD(t, ".. This is a comment\n\nVisible.\n")
	if strings.Contains(md, "comment") {
		t.Errorf("comment should be hidden: %q", md)
	}
	if !strings.Contains(md, "Visible") {
		t.Errorf("missing visible text: %q", md)
	}
}

// --- Directive handler tests ---

func TestDirective_CodeBlock(t *testing.T) {
	md := rstToMD(t, ".. code-block:: python\n\n   print('hi')\n")
	if !strings.Contains(md, "```python") {
		t.Errorf("missing python fence: %q", md)
	}
	if !strings.Contains(md, "print('hi')") {
		t.Errorf("missing code: %q", md)
	}
}

func TestDirective_CodeBlock_TwoSpaceIndent(t *testing.T) {
	md := rstToMD(t,
		".. code-block:: bash\n\n  sudo docker run\n")
	if !strings.Contains(md, "```bash") {
		t.Errorf("missing bash fence: %q", md)
	}
	if !strings.Contains(md, "sudo docker run") {
		t.Errorf("code content missing from block: %q", md)
	}
}

func TestDirective_Note_TwoSpaceIndent(t *testing.T) {
	md := rstToMD(t,
		".. note::\n\n  This only works if no other.\n")
	if !strings.Contains(md, "!!! note") {
		t.Errorf("missing admonition: %q", md)
	}
	if !strings.Contains(md, "This only works") {
		t.Errorf("admonition body missing: %q", md)
	}
}

func TestDirective_Note(t *testing.T) {
	md := rstToMD(t, ".. note::\n\n   Important info.\n")
	if !strings.Contains(md, "!!! note") {
		t.Errorf("missing admonition: %q", md)
	}
	if !strings.Contains(md, "Important info.") {
		t.Errorf("missing body: %q", md)
	}
}

func TestDirective_Note_WithInlineArg(t *testing.T) {
	md := rstToMD(t, ".. note:: This is the body text.\n")
	if !strings.Contains(md, "!!! note") {
		t.Errorf("missing admonition: %q", md)
	}
	if !strings.Contains(md, "This is the body text.") {
		t.Errorf("missing body: %q", md)
	}
	// Should NOT be a title
	if strings.Contains(md, `"This is the body text."`) {
		t.Errorf("arg should be body, not title: %q", md)
	}
}

func TestDirective_Warning(t *testing.T) {
	md := rstToMD(t, ".. warning::\n\n   Be careful.\n")
	if !strings.Contains(md, "!!! warning") {
		t.Errorf("missing warning: %q", md)
	}
}

func TestDirective_Admonition_WithTitle(t *testing.T) {
	md := rstToMD(t, ".. admonition:: Custom Title\n\n   Body text.\n")
	if !strings.Contains(md, `"Custom Title"`) {
		t.Errorf("missing custom title: %q", md)
	}
}

func TestDirective_Image(t *testing.T) {
	md := rstToMD(t, ".. image:: images/test.png\n    :alt: My Image\n")
	if !strings.Contains(md, "![My Image](images/test.png)") {
		t.Errorf("unexpected image: %q", md)
	}
}

func TestDirective_Image_AbsolutePath(t *testing.T) {
	md := rstToMD(t, ".. image:: /images/test.png\n    :alt: Abs\n")
	if strings.Contains(md, "/images/") {
		t.Errorf("leading / should be stripped: %q", md)
	}
	if !strings.Contains(md, "images/test.png") {
		t.Errorf("missing image path: %q", md)
	}
}

func TestDirective_Figure(t *testing.T) {
	md := rstToMD(t, ".. figure:: images/fig.png\n    :alt: Fig\n\n   Caption text.\n")
	if !strings.Contains(md, "![Fig]") {
		t.Errorf("missing figure: %q", md)
	}
	if !strings.Contains(md, "Caption text.") {
		t.Errorf("missing caption: %q", md)
	}
}

func TestDirective_CSVTable(t *testing.T) {
	rst := ".. csv-table::\n   :header: \"Name\", \"Age\"\n\n" +
		"   Alice,30\n   Bob,25\n"
	md := rstToMD(t, rst)
	if !strings.Contains(md, "| Name |") {
		t.Errorf("missing header: %q", md)
	}
	if !strings.Contains(md, "| Alice |") {
		t.Errorf("missing row: %q", md)
	}
}

func TestDirective_Topic(t *testing.T) {
	md := rstToMD(t, ".. topic:: My Topic\n\n   Topic body.\n")
	if !strings.Contains(md, "**My Topic**") {
		t.Errorf("missing topic title: %q", md)
	}
	if !strings.Contains(md, "> Topic body.") {
		t.Errorf("missing blockquote body: %q", md)
	}
}

func TestDirective_YouTube(t *testing.T) {
	md := rstToMD(t, ".. youtube:: abc123\n    :width: 640\n")
	if !strings.Contains(md, "youtube.com/embed/abc123") {
		t.Errorf("missing youtube embed: %q", md)
	}
	if !strings.Contains(md, `width="640"`) {
		t.Errorf("missing width: %q", md)
	}
}

func TestDirective_YouTube_DefaultWidth(t *testing.T) {
	md := rstToMD(t, ".. youtube:: xyz\n")
	if !strings.Contains(md, `width="560"`) {
		t.Errorf("expected default width 560: %q", md)
	}
}

func TestDirective_Deprecated(t *testing.T) {
	md := rstToMD(t, ".. deprecated:: 2.0\n\n   Use new API.\n")
	if !strings.Contains(md, "Deprecated since version 2.0") {
		t.Errorf("missing deprecated: %q", md)
	}
}

func TestDirective_VersionAdded(t *testing.T) {
	md := rstToMD(t, ".. versionadded:: 3.0\n")
	if !strings.Contains(md, "New in version 3.0") {
		t.Errorf("missing versionadded: %q", md)
	}
}

func TestDirective_VersionChanged(t *testing.T) {
	md := rstToMD(t, ".. versionchanged:: 4.0\n\n   Updated behaviour.\n")
	if !strings.Contains(md, "Changed in version 4.0") {
		t.Errorf("missing versionchanged: %q", md)
	}
}

func TestDirective_SeeAlso(t *testing.T) {
	md := rstToMD(t, ".. seealso::\n\n   Other docs.\n")
	if !strings.Contains(md, "See Also") {
		t.Errorf("missing seealso: %q", md)
	}
}

func TestDirective_Raw_HTML(t *testing.T) {
	md := rstToMD(t, ".. raw:: html\n\n   <div>hello</div>\n")
	if !strings.Contains(md, "<div>hello</div>") {
		t.Errorf("missing raw html: %q", md)
	}
}

func TestDirective_Raw_NonHTML(t *testing.T) {
	md := rstToMD(t, ".. raw:: latex\n\n   \\textbf{x}\n")
	if strings.Contains(md, "textbf") {
		t.Errorf("non-html raw should be skipped: %q", md)
	}
}

func TestDirective_Only(t *testing.T) {
	md := rstToMD(t, ".. only:: html\n\n   Visible content.\n")
	if !strings.Contains(md, "Visible content.") {
		t.Errorf("only content should be included: %q", md)
	}
}

func TestDirective_Skip(t *testing.T) {
	md := rstToMD(t, ".. contents::\n   :depth: 2\n\nParagraph.\n")
	if strings.Contains(md, "contents") {
		t.Errorf("contents should be skipped: %q", md)
	}
	if !strings.Contains(md, "Paragraph.") {
		t.Errorf("paragraph after skip missing: %q", md)
	}
}

func TestDirective_Toctree_NoOutput(t *testing.T) {
	md := rstToMD(t, ".. toctree::\n   :maxdepth: 2\n\n   page1\n   page2\n")
	// Toctree should produce no visible output
	trimmed := strings.TrimSpace(md)
	if trimmed != "" {
		t.Errorf("toctree should produce no output: %q", md)
	}
}

func TestDirective_LiteralInclude_Fallback(t *testing.T) {
	md := rstToMD(t, ".. literalinclude:: /some/file.py\n   :language: python\n")
	if !strings.Contains(md, "See source file") {
		t.Errorf("expected fallback: %q", md)
	}
}

func TestDirective_Unhandled(t *testing.T) {
	initDirectiveHandlers()
	root := Parse(".. unknown-directive:: arg\n\n   Body text.\n")
	ctx := &ConvertContext{
		FileMap:       map[string]string{},
		LabelMap:      map[string]labelInfo{},
		Substitutions: map[string]*Node{},
		CurrentFile:   "test.md",
	}
	w := shared.NewMarkdownWriter()
	for _, child := range root.Children {
		convertNode(ctx, child, w)
	}
	if len(ctx.Warnings) == 0 {
		t.Error("expected warning for unhandled directive")
	}
	if !strings.Contains(w.String(), "Body text.") {
		t.Errorf("body should be rendered: %q", w.String())
	}
}

// --- Table rendering tests ---

func TestGridTable_SimpleMarkdown(t *testing.T) {
	rst := "+-----+-----+\n| A   | B   |\n+=====+=====+\n" +
		"| 1   | 2   |\n+-----+-----+\n"
	md := rstToMD(t, rst)
	if !strings.Contains(md, "| A |") {
		t.Errorf("missing header: %q", md)
	}
	if !strings.Contains(md, "| 1 |") {
		t.Errorf("missing data: %q", md)
	}
}

func TestGridTable_WithBullets_HTML(t *testing.T) {
	rst := "+------+---------------------------+\n" +
		"| Icon | Behavior                  |\n" +
		"+======+===========================+\n" +
		"| Save | Menu:                     |\n" +
		"|      |                           |\n" +
		"|      |  * Save file              |\n" +
		"|      |                           |\n" +
		"|      |  * Save as                |\n" +
		"+------+---------------------------+\n"
	md := rstToMD(t, rst)
	if !strings.Contains(md, "<table>") {
		t.Errorf("expected HTML table: %q", md)
	}
	if !strings.Contains(md, "<ul>") {
		t.Errorf("expected bullet list: %q", md)
	}
	if !strings.Contains(md, "<li>Save file</li>") {
		t.Errorf("missing list item: %q", md)
	}
}

func TestGridTable_MergedCells(t *testing.T) {
	rst := "+------+--------+------+\n" +
		"| Icon | Action | Key  |\n" +
		"+======+========+======+\n" +
		"| Edit | Find   | CF   |\n" +
		"|      +--------+------+\n" +
		"|      | Replace| CSF  |\n" +
		"+------+--------+------+\n"
	md := rstToMD(t, rst)
	if !strings.Contains(md, "Find") &&
		!strings.Contains(md, "Replace") {
		t.Errorf("missing sub-rows: %q", md)
	}
}

func TestGridTable_NoHeader(t *testing.T) {
	rst := "+-----+-----+\n| a   | b   |\n+-----+-----+\n" +
		"| c   | d   |\n+-----+-----+\n"
	md := rstToMD(t, rst)
	if !strings.Contains(md, "| a |") {
		t.Errorf("missing data: %q", md)
	}
}

// --- inlineToHTML tests ---

func TestInlineToHTML_Bold(t *testing.T) {
	got := inlineToHTML("**bold**")
	if got != "<strong>bold</strong>" {
		t.Errorf("got %q", got)
	}
}

func TestInlineToHTML_Italic(t *testing.T) {
	got := inlineToHTML("*italic*")
	if got != "<em>italic</em>" {
		t.Errorf("got %q", got)
	}
}

func TestInlineToHTML_Code(t *testing.T) {
	got := inlineToHTML("`code`")
	if got != "<code>code</code>" {
		t.Errorf("got %q", got)
	}
}

func TestInlineToHTML_Link(t *testing.T) {
	got := inlineToHTML("[text](http://example.com)")
	if got != `<a href="http://example.com">text</a>` {
		t.Errorf("got %q", got)
	}
}

func TestInlineToHTML_Mixed(t *testing.T) {
	got := inlineToHTML("Click *Save* or **Cancel**")
	if !strings.Contains(got, "<em>Save</em>") {
		t.Errorf("missing italic: %q", got)
	}
	if !strings.Contains(got, "<strong>Cancel</strong>") {
		t.Errorf("missing bold: %q", got)
	}
}

func TestInlineToHTML_NoMarkup(t *testing.T) {
	got := inlineToHTML("plain text")
	if got != "plain text" {
		t.Errorf("got %q", got)
	}
}

// --- Inline role tests ---

func TestConvertRole_MenuSelection(t *testing.T) {
	result := ConvertInline(":menuselection:`File > Save`",
		nil, nil, "", nil, nil)
	if result != "**File > Save**" {
		t.Errorf("got %q", result)
	}
}

func TestConvertRole_Kbd(t *testing.T) {
	result := ConvertInline(":kbd:`Ctrl+S`", nil, nil, "", nil, nil)
	if result != "`Ctrl+S`" {
		t.Errorf("got %q", result)
	}
}

func TestConvertRole_File(t *testing.T) {
	result := ConvertInline(":file:`config.py`", nil, nil, "", nil, nil)
	if result != "`config.py`" {
		t.Errorf("got %q", result)
	}
}

func TestConvertRole_Code(t *testing.T) {
	result := ConvertInline(":code:`x = 1`", nil, nil, "", nil, nil)
	if result != "`x = 1`" {
		t.Errorf("got %q", result)
	}
}

func TestConvertRole_Abbr(t *testing.T) {
	result := ConvertInline(":abbr:`SQL (Structured Query Language)`",
		nil, nil, "", nil, nil)
	if result != "SQL (Structured Query Language)" {
		t.Errorf("got %q", result)
	}
}

func TestConvertRole_Sup(t *testing.T) {
	result := ConvertInline(":sup:`2`", nil, nil, "", nil, nil)
	if result != "<sup>2</sup>" {
		t.Errorf("got %q", result)
	}
}

func TestConvertRole_Sub(t *testing.T) {
	result := ConvertInline(":sub:`i`", nil, nil, "", nil, nil)
	if result != "<sub>i</sub>" {
		t.Errorf("got %q", result)
	}
}

func TestConvertRole_Unknown(t *testing.T) {
	result := ConvertInline(":custom:`value`", nil, nil, "", nil, nil)
	if result != "`value`" {
		t.Errorf("got %q", result)
	}
}

func TestConvertRole_Doc(t *testing.T) {
	fileMap := map[string]string{
		"coding_standards": "coding_standards.md",
	}
	result := ConvertInline(":doc:`coding_standards`",
		nil, fileMap, "code_review.md", nil, nil)
	if !strings.Contains(result, "coding_standards.md") {
		t.Errorf("expected link: %q", result)
	}
}

func TestConvertRole_Doc_WithTitle(t *testing.T) {
	fileMap := map[string]string{
		"coding_standards": "coding_standards.md",
	}
	result := ConvertInline(
		":doc:`Standards <coding_standards>`",
		nil, fileMap, "test.md", nil, nil)
	if !strings.Contains(result, "[Standards]") {
		t.Errorf("expected title: %q", result)
	}
}

func TestConvertRole_Ref_NoMatch(t *testing.T) {
	result := ConvertInline(":ref:`nonexistent`",
		map[string]labelInfo{}, nil, "", nil, nil)
	if result != "nonexistent" {
		t.Errorf("expected plain text: %q", result)
	}
}

func TestConvertRole_Index(t *testing.T) {
	result := ConvertInline(":index:`Search Term`",
		nil, nil, "", nil, nil)
	if result != "Search Term" {
		t.Errorf("expected stripped: %q", result)
	}
}

// --- List item with embedded directive ---

func TestConvert_ListWithEmbeddedCodeBlock(t *testing.T) {
	rst := "* Step one:\n\n  .. code-block:: bash\n\n     echo hello\n\n* Step two\n"
	md := rstToMD(t, rst)
	if !strings.Contains(md, "```bash") {
		t.Errorf("missing code fence: %q", md)
	}
	if !strings.Contains(md, "echo hello") {
		t.Errorf("missing code: %q", md)
	}
	if !strings.Contains(md, "Step two") {
		t.Errorf("missing second item: %q", md)
	}
}

func TestConvert_ListWithEmbeddedCSVTable(t *testing.T) {
	rst := "* Config:\n\n  .. csv-table::\n     :header: \"Key\",\"Value\"\n\n" +
		"     A,1\n\n* Next item\n"
	md := rstToMD(t, rst)
	if !strings.Contains(md, "| Key |") {
		t.Errorf("missing csv table: %q", md)
	}
}

func TestContainsDirective(t *testing.T) {
	if !containsDirective("text\n.. code-block:: python\n   code") {
		t.Error("should detect directive")
	}
	if containsDirective("just text\nno directives") {
		t.Error("should not detect directive")
	}
}

// --- relativeImagePath ---

func TestRelativeImagePath_SameDir(t *testing.T) {
	got := relativeImagePath("page.md", "images/test.png")
	if got != "images/test.png" {
		t.Errorf("got %q", got)
	}
}

func TestRelativeImagePath_Subdir(t *testing.T) {
	got := relativeImagePath("sub/page.md", "images/test.png")
	if got != "../images/test.png" {
		t.Errorf("got %q", got)
	}
}

func TestRelativeImagePath_DeepSubdir(t *testing.T) {
	got := relativeImagePath("a/b/page.md", "images/test.png")
	if got != "../../images/test.png" {
		t.Errorf("got %q", got)
	}
}

// --- cleanMarkdown ---

func TestCleanMarkdown(t *testing.T) {
	input := "line  \t\n\n\n\n\nend"
	got := cleanMarkdown(input)
	if strings.Contains(got, "  ") {
		t.Errorf("trailing whitespace not removed: %q", got)
	}
	if strings.Contains(got, "\n\n\n\n") {
		t.Errorf("excess blank lines not collapsed: %q", got)
	}
	if !strings.HasSuffix(got, "\n") {
		t.Errorf("should end with newline: %q", got)
	}
}

// --- splitCSVRows ---

func TestSplitCSVRows_Simple(t *testing.T) {
	rows := splitCSVRows("a,b\nc,d")
	if len(rows) != 2 {
		t.Fatalf("expected 2 rows, got %d", len(rows))
	}
}

func TestSplitCSVRows_MultiLine(t *testing.T) {
	rows := splitCSVRows("\"line1\nline2\",b\nc,d")
	if len(rows) != 2 {
		t.Fatalf("expected 2 rows, got %d: %v", len(rows), rows)
	}
}

// --- Table directive ---

func TestDirective_Table(t *testing.T) {
	rst := ".. table::\n   :class: longtable\n\n" +
		"   +-----+-----+\n   | a   | b   |\n   +-----+-----+\n" +
		"   | c   | d   |\n   +-----+-----+\n"
	md := rstToMD(t, rst)
	if !strings.Contains(md, "a") && !strings.Contains(md, "c") {
		t.Errorf("missing table content: %q", md)
	}
}

// --- LiteralInclude with pgadmin-src ---

func TestDirective_LiteralInclude_WithSource(t *testing.T) {
	initDirectiveHandlers()
	srcDir := t.TempDir()
	os.MkdirAll(filepath.Join(srcDir, "web"), 0755)
	os.WriteFile(filepath.Join(srcDir, "web", "config.py"),
		[]byte("# config"), 0644)

	root := Parse(".. literalinclude:: web/config.py\n   :language: python\n")
	ctx := &ConvertContext{
		FileMap:       map[string]string{},
		LabelMap:      map[string]labelInfo{},
		Substitutions: map[string]*Node{},
		CurrentFile:   "test.md",
		PgAdminSrcDir: srcDir,
	}
	w := shared.NewMarkdownWriter()
	for _, child := range root.Children {
		convertNode(ctx, child, w)
	}
	if !strings.Contains(w.String(), "# config") {
		t.Errorf("expected included file: %q", w.String())
	}
}

// --- Warnings ---

func TestConverter_Warnings(t *testing.T) {
	initDirectiveHandlers()
	srcDir := t.TempDir()
	outDir := t.TempDir()
	os.WriteFile(filepath.Join(srcDir, "index.rst"),
		[]byte("Title\n=====\n"), 0644)
	c := NewConverter(srcDir, outDir, "1.0", "", "", nil, false)
	c.Convert()
	// Warnings() should return without panic
	_ = c.Warnings()
}

// --- collapseCell ---

func TestCollapseCell_Plain(t *testing.T) {
	lines := []string{"hello", "world"}
	got := collapseCell(lines)
	if got != "hello world" {
		t.Errorf("got %q", got)
	}
}

func TestCollapseCell_Bullets(t *testing.T) {
	lines := []string{"intro:", "", "* item 1", "", "* item 2"}
	got := collapseCell(lines)
	if !strings.Contains(got, "\n* item 1") {
		t.Errorf("bullets should be on new lines: %q", got)
	}
}

func TestCollapseCell_Empty(t *testing.T) {
	got := collapseCell([]string{"", "", ""})
	if got != "" {
		t.Errorf("expected empty: %q", got)
	}
}

// --- isDirectiveOption ---

func TestIsDirectiveOption(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{":alt: text", true},
		{":max-depth: 2", true},
		{":class: longtable", true},
		{":ref:`target`", false},
		{":doc:`path`", false},
		{"not an option", false},
		{"::", false},
		{":", false},
		{":a:", true},
		{":a:b", false}, // no space after colon
	}
	for _, tt := range tests {
		got := isDirectiveOption(tt.input)
		if got != tt.want {
			t.Errorf("isDirectiveOption(%q) = %v, want %v",
				tt.input, got, tt.want)
		}
	}
}

// --- handleContainer tests ---

func TestDirective_Container_ImgDark(t *testing.T) {
	initDirectiveHandlers()
	rst := ".. container:: img-dark\n\n" +
		"   .. image:: images/arch.png\n      :alt: Arch\n"
	root := Parse(rst)
	ctx := &ConvertContext{
		FileMap:       map[string]string{},
		LabelMap:      map[string]labelInfo{},
		Substitutions: map[string]*Node{},
		CurrentFile:   "test.md",
	}
	w := shared.NewMarkdownWriter()
	for _, child := range root.Children {
		convertNode(ctx, child, w)
	}
	md := w.String()
	if !strings.Contains(md, "#only-dark") {
		t.Errorf("expected #only-dark fragment: %q", md)
	}
	if !strings.Contains(md, "Arch") {
		t.Errorf("expected alt text: %q", md)
	}
}

func TestDirective_Container_ImgLight(t *testing.T) {
	initDirectiveHandlers()
	rst := ".. container:: img-light\n\n" +
		"   .. image:: images/arch.png\n      :alt: Light\n"
	root := Parse(rst)
	ctx := &ConvertContext{
		FileMap:       map[string]string{},
		LabelMap:      map[string]labelInfo{},
		Substitutions: map[string]*Node{},
		CurrentFile:   "test.md",
	}
	w := shared.NewMarkdownWriter()
	for _, child := range root.Children {
		convertNode(ctx, child, w)
	}
	md := w.String()
	if !strings.Contains(md, "#only-light") {
		t.Errorf("expected #only-light fragment: %q", md)
	}
}

func TestDirective_Container_NoClass(t *testing.T) {
	md := rstToMD(t, ".. container::\n\n   Some content.\n")
	if !strings.Contains(md, "Some content.") {
		t.Errorf("expected body content: %q", md)
	}
	if strings.Contains(md, "#only-") {
		t.Errorf("should not have color scheme fragment: %q", md)
	}
}

// --- handleListTable tests ---

func TestDirective_ListTable(t *testing.T) {
	rst := ".. list-table::\n" +
		"   :header-rows: 1\n\n" +
		"   * - Name\n" +
		"     - Age\n" +
		"   * - Alice\n" +
		"     - 30\n"
	md := rstToMD(t, rst)
	if !strings.Contains(md, "| Name |") {
		t.Errorf("missing header Name: %q", md)
	}
	if !strings.Contains(md, "| Age |") {
		t.Errorf("missing header Age: %q", md)
	}
	if !strings.Contains(md, "| Alice |") {
		t.Errorf("missing row Alice: %q", md)
	}
	if !strings.Contains(md, "---|") {
		t.Errorf("missing separator: %q", md)
	}
}

func TestDirective_ListTable_NoHeader(t *testing.T) {
	rst := ".. list-table::\n" +
		"   :header-rows: 0\n\n" +
		"   * - A\n" +
		"     - B\n" +
		"   * - C\n" +
		"     - D\n"
	md := rstToMD(t, rst)
	if !strings.Contains(md, "| A |") {
		t.Errorf("missing cell A: %q", md)
	}
	if !strings.Contains(md, "| C |") {
		t.Errorf("missing cell C: %q", md)
	}
}

func TestDirective_ListTable_WithTitle(t *testing.T) {
	rst := ".. list-table:: My Table\n" +
		"   :header-rows: 1\n\n" +
		"   * - Col1\n" +
		"     - Col2\n"
	md := rstToMD(t, rst)
	if !strings.Contains(md, "**My Table**") {
		t.Errorf("missing title: %q", md)
	}
}

// --- handleTabs and handleTab tests ---

func TestDirective_Tabs(t *testing.T) {
	rst := ".. tabs::\n\n" +
		"   .. tab:: Ubuntu\n\n" +
		"      apt install\n\n" +
		"   .. tab:: macOS\n\n" +
		"      brew install\n"
	md := rstToMD(t, rst)
	if !strings.Contains(md, "**Ubuntu**") {
		t.Errorf("missing Ubuntu tab heading: %q", md)
	}
	if !strings.Contains(md, "**macOS**") {
		t.Errorf("missing macOS tab heading: %q", md)
	}
	if !strings.Contains(md, "apt install") {
		t.Errorf("missing Ubuntu content: %q", md)
	}
	if !strings.Contains(md, "brew install") {
		t.Errorf("missing macOS content: %q", md)
	}
}

func TestDirective_Tab_Standalone(t *testing.T) {
	md := rstToMD(t, ".. tab:: Linux\n\n   Use apt.\n")
	if !strings.Contains(md, "**Linux**") {
		t.Errorf("missing tab heading: %q", md)
	}
	if !strings.Contains(md, "Use apt.") {
		t.Errorf("missing tab body: %q", md)
	}
}

// --- handleRaw with object tag ---

func TestDirective_Raw_ObjectTag(t *testing.T) {
	initDirectiveHandlers()
	rst := ".. container:: img-light\n\n" +
		"   .. raw:: html\n\n" +
		"      <object data=\"../_static/arch.svg\" " +
		"type=\"image/svg+xml\"></object>\n"
	root := Parse(rst)
	ctx := &ConvertContext{
		FileMap:       map[string]string{},
		LabelMap:      map[string]labelInfo{},
		Substitutions: map[string]*Node{},
		CurrentFile:   "test.md",
	}
	w := shared.NewMarkdownWriter()
	for _, child := range root.Children {
		convertNode(ctx, child, w)
	}
	md := w.String()
	if !strings.Contains(md, "![image]") {
		t.Errorf("expected markdown image: %q", md)
	}
	if !strings.Contains(md, "#only-light") {
		t.Errorf("expected light scheme fragment: %q", md)
	}
}

// --- handleAPIDef tests ---

func TestDirective_APIDef_Function(t *testing.T) {
	md := rstToMD(t,
		".. function:: my_func(arg1, arg2)\n\n"+
			"   Description of function.\n")
	if !strings.Contains(md, "*function*") {
		t.Errorf("missing kind label: %q", md)
	}
	if !strings.Contains(md, "`my_func(arg1, arg2)`") {
		t.Errorf("missing signature: %q", md)
	}
	if !strings.Contains(md, "Description of function.") {
		t.Errorf("missing body: %q", md)
	}
}

func TestDirective_APIDef_Class(t *testing.T) {
	md := rstToMD(t,
		".. class:: MyClass\n\n   Class description.\n")
	if !strings.Contains(md, "*class*") {
		t.Errorf("missing kind label: %q", md)
	}
	if !strings.Contains(md, "`MyClass`") {
		t.Errorf("missing class name: %q", md)
	}
	if !strings.Contains(md, "Class description.") {
		t.Errorf("missing body: %q", md)
	}
}

func TestDirective_APIDef_AutoClass(t *testing.T) {
	md := rstToMD(t, ".. autoclass:: MyAutoClass\n")
	if !strings.Contains(md, "*class*") {
		t.Errorf("expected 'class' (auto prefix stripped): %q", md)
	}
	if !strings.Contains(md, "`MyAutoClass`") {
		t.Errorf("missing class name: %q", md)
	}
}

// --- handleRubric tests ---

func TestDirective_Rubric(t *testing.T) {
	md := rstToMD(t, ".. rubric:: My Title\n")
	if !strings.Contains(md, "**My Title**") {
		t.Errorf("expected bold rubric: %q", md)
	}
}

// --- handleParsedLiteral tests ---

func TestDirective_ParsedLiteral(t *testing.T) {
	md := rstToMD(t,
		".. parsed-literal::\n\n   code here\n")
	if !strings.Contains(md, "```") {
		t.Errorf("missing code fence: %q", md)
	}
	if !strings.Contains(md, "code here") {
		t.Errorf("missing code content: %q", md)
	}
}

// --- handleDoctest tests ---

func TestDirective_Doctest(t *testing.T) {
	md := rstToMD(t,
		".. doctest::\n\n   >>> print(1)\n   1\n")
	if !strings.Contains(md, "```python") {
		t.Errorf("missing python fence: %q", md)
	}
	if !strings.Contains(md, ">>> print(1)") {
		t.Errorf("missing doctest content: %q", md)
	}
}

// --- handleInclude tests ---

func TestDirective_Include(t *testing.T) {
	md := rstToMD(t,
		".. include:: path/to/file.rst\n")
	if !strings.Contains(md, "*See: `path/to/file.rst`*") {
		t.Errorf("expected see reference: %q", md)
	}
}

// --- colorSchemeFragment tests ---

func TestColorSchemeFragment_Dark(t *testing.T) {
	ctx := &ConvertContext{ColorScheme: "dark"}
	got := colorSchemeFragment(ctx)
	if got != "#only-dark" {
		t.Errorf("expected '#only-dark', got %q", got)
	}
}

func TestColorSchemeFragment_Light(t *testing.T) {
	ctx := &ConvertContext{ColorScheme: "light"}
	got := colorSchemeFragment(ctx)
	if got != "#only-light" {
		t.Errorf("expected '#only-light', got %q", got)
	}
}

func TestColorSchemeFragment_Empty(t *testing.T) {
	ctx := &ConvertContext{ColorScheme: ""}
	got := colorSchemeFragment(ctx)
	if got != "" {
		t.Errorf("expected empty string, got %q", got)
	}
}

// --- resolveAnonymousLinks tests ---

func TestResolveAnonymousLinks_OnePair(t *testing.T) {
	input := "`Click here`__\n\n.. __: https://example.com"
	got := resolveAnonymousLinks(input)
	if !strings.Contains(got, "[Click here](https://example.com)") {
		t.Errorf("expected resolved link: %q", got)
	}
}

func TestResolveAnonymousLinks_MultiplePairs(t *testing.T) {
	input := "`First`__ and `Second`__\n\n" +
		".. __: https://first.com\n" +
		".. __: https://second.com"
	got := resolveAnonymousLinks(input)
	if !strings.Contains(got, "[First](https://first.com)") {
		t.Errorf("missing first link: %q", got)
	}
	if !strings.Contains(got, "[Second](https://second.com)") {
		t.Errorf("missing second link: %q", got)
	}
}

func TestResolveAnonymousLinks_UnmatchedRef(t *testing.T) {
	input := "`Orphan`__"
	got := resolveAnonymousLinks(input)
	if !strings.Contains(got, "Orphan") {
		t.Errorf("expected fallback text: %q", got)
	}
	if strings.Contains(got, "__") {
		t.Errorf("should not have __ markers: %q", got)
	}
}

// --- ProjectName tests ---

func TestProjectName(t *testing.T) {
	srcDir := t.TempDir()
	outDir := t.TempDir()
	confPy := `# Sphinx config
project = "PostgREST"
version = "14.5"
`
	os.WriteFile(filepath.Join(srcDir, "conf.py"),
		[]byte(confPy), 0644)
	os.WriteFile(filepath.Join(srcDir, "index.rst"),
		[]byte("Title\n=====\n"), 0644)

	c := NewConverter(srcDir, outDir, "1.0", "", "", nil, false)
	got := c.ProjectName()
	if got != "PostgREST" {
		t.Errorf("expected 'PostgREST', got %q", got)
	}
}

func TestProjectName_SingleQuotes(t *testing.T) {
	srcDir := t.TempDir()
	outDir := t.TempDir()
	confPy := "project = 'MyProject'\n"
	os.WriteFile(filepath.Join(srcDir, "conf.py"),
		[]byte(confPy), 0644)

	c := NewConverter(srcDir, outDir, "1.0", "", "", nil, false)
	got := c.ProjectName()
	if got != "MyProject" {
		t.Errorf("expected 'MyProject', got %q", got)
	}
}

func TestProjectName_Missing(t *testing.T) {
	srcDir := t.TempDir()
	outDir := t.TempDir()
	c := NewConverter(srcDir, outDir, "1.0", "", "", nil, false)
	got := c.ProjectName()
	if got != "" {
		t.Errorf("expected empty, got %q", got)
	}
}

// --- convertBlockQuote tests ---

func TestConvert_BlockQuote(t *testing.T) {
	// A block quote in RST is indented text after a paragraph.
	// The parser produces a BlockQuoteNode when it encounters
	// indented text that isn't part of another construct.
	initDirectiveHandlers()
	node := &Node{
		Type: BlockQuoteNode,
		Text: "This is quoted.\nSecond line.",
	}
	ctx := &ConvertContext{
		FileMap:       map[string]string{},
		LabelMap:      map[string]labelInfo{},
		Substitutions: map[string]*Node{},
		CurrentFile:   "test.md",
	}
	w := shared.NewMarkdownWriter()
	convertNode(ctx, node, w)
	md := w.String()
	if !strings.Contains(md, "> This is quoted.") {
		t.Errorf("missing blockquote prefix: %q", md)
	}
	if !strings.Contains(md, "> Second line.") {
		t.Errorf("missing second line: %q", md)
	}
}

// --- prescanDefinitions tests ---

func TestPrescanDefinitions_Substitution(t *testing.T) {
	text := ".. |myvar| replace:: Hello World\n"
	subs := map[string]*Node{}
	targets := map[string]string{}
	prescanDefinitions(text, subs, targets)
	if def, ok := subs["myvar"]; !ok {
		t.Error("expected substitution 'myvar'")
	} else if def.DirectiveArg != "Hello World" {
		t.Errorf("expected 'Hello World', got %q",
			def.DirectiveArg)
	}
}

func TestPrescanDefinitions_HyperlinkTarget(t *testing.T) {
	text := ".. _PostgreSQL: https://www.postgresql.org/\n"
	subs := map[string]*Node{}
	targets := map[string]string{}
	prescanDefinitions(text, subs, targets)
	if url, ok := targets["PostgreSQL"]; !ok {
		t.Error("expected target 'PostgreSQL'")
	} else if url != "https://www.postgresql.org/" {
		t.Errorf("expected URL, got %q", url)
	}
}

func TestPrescanDefinitions_Both(t *testing.T) {
	text := ".. |ver| replace:: 14.5\n" +
		".. _docs: https://docs.example.com\n"
	subs := map[string]*Node{}
	targets := map[string]string{}
	prescanDefinitions(text, subs, targets)
	if _, ok := subs["ver"]; !ok {
		t.Error("expected substitution 'ver'")
	}
	if _, ok := targets["docs"]; !ok {
		t.Error("expected target 'docs'")
	}
}

// --- shouldSkipSection tests ---

func TestShouldSkipSection_Match(t *testing.T) {
	ctx := &ConvertContext{
		SkipSections: []string{"Changelog", "Release Notes"},
	}
	if !ctx.shouldSkipSection("Changelog") {
		t.Error("expected match for 'Changelog'")
	}
	if !ctx.shouldSkipSection("changelog") {
		t.Error("expected case-insensitive match")
	}
	if !ctx.shouldSkipSection("Release Notes") {
		t.Error("expected match for 'Release Notes'")
	}
}

func TestShouldSkipSection_NoMatch(t *testing.T) {
	ctx := &ConvertContext{
		SkipSections: []string{"Changelog"},
	}
	if ctx.shouldSkipSection("Introduction") {
		t.Error("should not match 'Introduction'")
	}
}

func TestShouldSkipSection_EmptySkipList(t *testing.T) {
	ctx := &ConvertContext{}
	if ctx.shouldSkipSection("Anything") {
		t.Error("empty skip list should never match")
	}
}

// --- convertChildren with SkipSections ---

func TestConvertChildren_SkipSection(t *testing.T) {
	initDirectiveHandlers()
	rst := "Intro\n=====\n\nKeep this.\n\n" +
		"Changelog\n=========\n\nSkip this.\n\n" +
		"Usage\n=====\n\nAlso keep.\n"
	root := Parse(rst)
	ctx := &ConvertContext{
		FileMap:       map[string]string{},
		LabelMap:      map[string]labelInfo{},
		Substitutions: map[string]*Node{},
		CurrentFile:   "test.md",
		SkipSections:  []string{"Changelog"},
	}
	w := shared.NewMarkdownWriter()
	convertChildren(ctx, root.Children, w)
	md := w.String()
	if !strings.Contains(md, "Keep this.") {
		t.Errorf("expected intro content: %q", md)
	}
	if strings.Contains(md, "Skip this.") {
		t.Errorf("changelog content should be skipped: %q", md)
	}
	if !strings.Contains(md, "Also keep.") {
		t.Errorf("expected usage content: %q", md)
	}
}

// --- parseDoctestBlock tests ---

func TestParseDoctestBlock(t *testing.T) {
	rst := "Example usage:\n\n   >>> x = 1\n   >>> print(x)\n   1\n"
	md := rstToMD(t, rst)
	if !strings.Contains(md, "```python") {
		t.Errorf("expected python code block: %q", md)
	}
	if !strings.Contains(md, ">>> x = 1") {
		t.Errorf("missing doctest line: %q", md)
	}
}

// --- ConvertInline additional branch tests ---

func TestConvertInline_PEP(t *testing.T) {
	result := ConvertInline(":pep:`484`",
		nil, nil, "", nil, nil)
	if !strings.Contains(result, "[PEP 484]") {
		t.Errorf("expected PEP link text: %q", result)
	}
	if !strings.Contains(result, "peps.python.org/pep-484") {
		t.Errorf("expected PEP URL: %q", result)
	}
}

func TestConvertInline_TildePrefix(t *testing.T) {
	result := ConvertInline("`~module.Class`",
		nil, nil, "", nil, nil)
	if result != "`Class`" {
		t.Errorf("expected shortened name, got %q", result)
	}
}

func TestConvertInline_BangPrefix(t *testing.T) {
	result := ConvertInline("`!name`",
		nil, nil, "", nil, nil)
	if result != "`name`" {
		t.Errorf("expected plain name, got %q", result)
	}
}

func TestConvertInline_BackslashEscape(t *testing.T) {
	result := ConvertInline("\\*not italic\\*",
		nil, nil, "", nil, nil)
	if result != "*not italic*" {
		t.Errorf("expected unescaped text, got %q", result)
	}
}

func TestConvertInline_Ref_WithLabelTitle(t *testing.T) {
	labels := map[string]labelInfo{
		"install": {
			File:   "install.md",
			Anchor: "install",
			Title:  "Installation Guide",
		},
	}
	result := ConvertInline(":ref:`install`",
		labels, nil, "index.md", nil, nil)
	if !strings.Contains(result, "[Installation Guide]") {
		t.Errorf("expected label title: %q", result)
	}
	if !strings.Contains(result, "install.md") {
		t.Errorf("expected link to install.md: %q", result)
	}
}

func TestConvertInline_Substitution_Replace(t *testing.T) {
	subs := map[string]*Node{
		"version": {
			DirectiveName: "replace",
			DirectiveArg:  "14.5",
		},
	}
	result := ConvertInline("Version: |version|",
		nil, nil, "", subs, nil)
	if !strings.Contains(result, "14.5") {
		t.Errorf("expected replaced value: %q", result)
	}
}

func TestConvertInline_RoleTildePrefix(t *testing.T) {
	result := ConvertInline(":class:`~mymod.MyClass`",
		nil, nil, "", nil, nil)
	if result != "`MyClass`" {
		t.Errorf("expected shortened class, got %q", result)
	}
}

func TestConvertInline_RoleBangPrefix(t *testing.T) {
	result := ConvertInline(":func:`!hidden_func`",
		nil, nil, "", nil, nil)
	if result != "`hidden_func`" {
		t.Errorf("expected unsuppressed name, got %q", result)
	}
}

// --- copyImage tests ---

// TestCopyImage_ResolvesViaSourceFile verifies that image paths are
// resolved relative to the RST source file (SourceFile), not the
// output markdown file (CurrentFile).  This matters when toctree
// reorganisation places the output file in a different subdirectory.
func TestCopyImage_ResolvesViaSourceFile(t *testing.T) {
	// Set up temp source and output directories
	srcDir := t.TempDir()
	outDir := t.TempDir()

	// Create source images directory with a test image
	imgDir := filepath.Join(srcDir, "images")
	if err := os.MkdirAll(imgDir, 0755); err != nil {
		t.Fatal(err)
	}
	imgData := []byte("fake-png-data")
	if err := os.WriteFile(
		filepath.Join(imgDir, "test.png"), imgData, 0644); err != nil {
		t.Fatal(err)
	}

	ctx := &ConvertContext{
		SrcDir: srcDir,
		OutDir: outDir,
		// Source RST file is at the root of SrcDir
		SourceFile: "page.rst",
		// Output file is nested in a subdirectory
		CurrentFile: "subdir/page.md",
	}

	// Copy with path relative to the source RST file
	result := ctx.copyImage("images/test.png")

	// The resolved path should be images/test.png (relative to
	// source root), not subdir/images/test.png
	if result != "images/test.png" {
		t.Errorf("expected images/test.png, got %q", result)
	}

	// Verify the image was actually copied
	copied := filepath.Join(outDir, "images", "test.png")
	data, err := os.ReadFile(copied)
	if err != nil {
		t.Fatalf("image not copied to %s: %v", copied, err)
	}
	if string(data) != string(imgData) {
		t.Error("copied image content mismatch")
	}
}

// TestCopyImage_NestedSourceAndOutput verifies correct behaviour
// when both source and output files are in subdirectories.
func TestCopyImage_NestedSourceAndOutput(t *testing.T) {
	srcDir := t.TempDir()
	outDir := t.TempDir()

	// Create _static directory with a test image
	staticDir := filepath.Join(srcDir, "_static")
	if err := os.MkdirAll(staticDir, 0755); err != nil {
		t.Fatal(err)
	}
	imgData := []byte("fake-png-data")
	if err := os.WriteFile(
		filepath.Join(staticDir, "test.png"), imgData, 0644); err != nil {
		t.Fatal(err)
	}

	ctx := &ConvertContext{
		SrcDir:      srcDir,
		OutDir:      outDir,
		SourceFile:  "explanations/install.rst",
		CurrentFile: "explanations/install.md",
	}

	// Image ref: ../_static/test.png (relative to source file)
	result := ctx.copyImage("../_static/test.png")

	// Should resolve to _static/test.png
	if result != "_static/test.png" {
		t.Errorf("expected _static/test.png, got %q", result)
	}

	// Verify copy
	copied := filepath.Join(outDir, "_static", "test.png")
	if _, err := os.Stat(copied); err != nil {
		t.Fatalf("image not copied to %s: %v", copied, err)
	}
}
