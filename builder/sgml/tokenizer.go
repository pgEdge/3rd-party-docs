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
	"fmt"
	"strings"
	"unicode"
)

// TokenType identifies the kind of token produced by the tokenizer.
type TokenType int

const (
	TokenTagOpen  TokenType = iota // <element attr="val">
	TokenTagClose                  // </element>
	TokenText                      // character data
	TokenComment                   // <!-- ... -->
	TokenPI                        // <?...?>
	TokenEOF
)

// Token represents a single token from the SGML input.
type Token struct {
	Type  TokenType
	Tag   string            // element name (lowercased)
	Attrs map[string]string // attributes (for TagOpen)
	Text  string            // text content or comment body
	Line  int               // source line number
}

// Tokenizer breaks SGML input into a stream of tokens.
type Tokenizer struct {
	input []rune
	pos   int
	line  int
}

// NewTokenizer creates a tokenizer for the given SGML text.
func NewTokenizer(input string) *Tokenizer {
	return &Tokenizer{
		input: []rune(input),
		pos:   0,
		line:  1,
	}
}

// Next returns the next token from the input.
func (t *Tokenizer) Next() Token {
	if t.pos >= len(t.input) {
		return Token{Type: TokenEOF, Line: t.line}
	}

	if t.input[t.pos] == '<' {
		return t.readTag()
	}

	return t.readText()
}

// Tokenize returns all tokens from the input.
func (t *Tokenizer) Tokenize() []Token {
	var tokens []Token
	for {
		tok := t.Next()
		tokens = append(tokens, tok)
		if tok.Type == TokenEOF {
			break
		}
	}
	return tokens
}

// readText reads character data up to the next '<'.
func (t *Tokenizer) readText() Token {
	start := t.pos
	line := t.line
	for t.pos < len(t.input) && t.input[t.pos] != '<' {
		if t.input[t.pos] == '\n' {
			t.line++
		}
		t.pos++
	}
	text := string(t.input[start:t.pos])
	return Token{Type: TokenText, Text: text, Line: line}
}

// readTag reads a tag, comment, or processing instruction.
func (t *Tokenizer) readTag() Token {
	line := t.line

	// Check for CDATA section: <![CDATA[...]]>
	if t.lookingAt("<![CDATA[") {
		return t.readCDATA(line)
	}

	// Check for comment: <!-- ... -->
	// But NOT if followed immediately by </ which means this is literal
	// text like <literal><!--</literal> in the PG docs.
	if t.lookingAt("<!--") {
		if t.lookingAt("<!--</") {
			// Literal text, not a real comment — return < as text
			t.pos++
			return Token{Type: TokenText, Text: "<", Line: line}
		}
		return t.readComment(line)
	}

	// Check for processing instruction: <?...?>
	// But NOT if the content before ?> contains </ which indicates
	// literal text like <literal><?x</literal>.
	if t.lookingAt("<?") {
		if t.isLiteralPI() {
			// Literal text, not a real PI — return < as text
			t.pos++
			return Token{Type: TokenText, Text: "<", Line: line}
		}
		return t.readPI(line)
	}

	// Check for SGML declarations we should skip: <!DOCTYPE, <!ENTITY, etc.
	// These should have been processed by entity resolution, but skip
	// any that remain. Only treat as a declaration if followed by a
	// valid SGML declaration keyword or '['. Otherwise it's literal
	// text (e.g., "(?<!" from regex docs where &lt; expanded to <).
	if t.lookingAt("<!") && !t.lookingAt("</") {
		if t.isRealDeclaration() {
			return t.readDeclaration(line)
		}
		// Literal text — return the < as text
		t.pos++
		return Token{Type: TokenText, Text: "<", Line: line}
	}

	// Check for closing tag: </element>
	if t.lookingAt("</") {
		return t.readCloseTag(line)
	}

	// Opening tag: <element ...>
	return t.readOpenTag(line)
}

// readCDATA reads <![CDATA[...]]> and returns a text token.
func (t *Tokenizer) readCDATA(line int) Token {
	t.pos += 9 // skip <![CDATA[
	start := t.pos
	for t.pos < len(t.input)-2 {
		if t.input[t.pos] == '\n' {
			t.line++
		}
		if t.input[t.pos] == ']' && t.input[t.pos+1] == ']' && t.input[t.pos+2] == '>' {
			text := string(t.input[start:t.pos])
			t.pos += 3 // skip ]]>
			return Token{Type: TokenText, Text: text, Line: line}
		}
		t.pos++
	}
	// Unterminated CDATA — consume to end
	t.pos = len(t.input)
	return Token{Type: TokenText, Text: string(t.input[start:]), Line: line}
}

// readComment reads <!-- ... --> and returns a comment token.
func (t *Tokenizer) readComment(line int) Token {
	t.pos += 4 // skip <!--
	start := t.pos
	for t.pos < len(t.input)-2 {
		if t.input[t.pos] == '\n' {
			t.line++
		}
		if t.input[t.pos] == '-' && t.input[t.pos+1] == '-' && t.input[t.pos+2] == '>' {
			text := strings.TrimSpace(string(t.input[start:t.pos]))
			t.pos += 3
			return Token{Type: TokenComment, Text: text, Line: line}
		}
		t.pos++
	}
	// Unterminated comment — consume to end
	t.pos = len(t.input)
	return Token{Type: TokenComment, Text: string(t.input[start:]), Line: line}
}

// readPI reads <?...?> and returns a processing instruction token.
func (t *Tokenizer) readPI(line int) Token {
	t.pos += 2 // skip <?
	start := t.pos
	for t.pos < len(t.input)-1 {
		if t.input[t.pos] == '\n' {
			t.line++
		}
		if t.input[t.pos] == '?' && t.input[t.pos+1] == '>' {
			text := string(t.input[start:t.pos])
			t.pos += 2
			return Token{Type: TokenPI, Text: text, Line: line}
		}
		t.pos++
	}
	t.pos = len(t.input)
	return Token{Type: TokenPI, Text: string(t.input[start:]), Line: line}
}

// readDeclaration skips <!...> declarations that weren't processed
// during entity resolution (e.g. stray DOCTYPE fragments).
func (t *Tokenizer) readDeclaration(line int) Token {
	start := t.pos
	depth := 0
	for t.pos < len(t.input) {
		ch := t.input[t.pos]
		if ch == '\n' {
			t.line++
		}
		if ch == '<' {
			depth++
		}
		if ch == '>' {
			depth--
			if depth <= 0 {
				t.pos++
				return Token{
					Type: TokenComment,
					Text: string(t.input[start:t.pos]),
					Line: line,
				}
			}
		}
		t.pos++
	}
	t.pos = len(t.input)
	return Token{Type: TokenComment, Text: string(t.input[start:]), Line: line}
}

// readCloseTag reads </element>.
func (t *Tokenizer) readCloseTag(line int) Token {
	t.pos += 2 // skip </
	t.skipWhitespace()

	name := t.readName()
	t.skipWhitespace()

	// Consume closing >
	if t.pos < len(t.input) && t.input[t.pos] == '>' {
		t.pos++
	}

	return Token{
		Type: TokenTagClose,
		Tag:  strings.ToLower(name),
		Line: line,
	}
}

// readOpenTag reads <element attr="val" ...> or <element ... />.
func (t *Tokenizer) readOpenTag(line int) Token {
	ltPos := t.pos // position of the <
	t.pos++        // skip <

	// Check if there's whitespace after < — real tags don't have
	// whitespace between < and the tag name, so this is literal text.
	if t.pos < len(t.input) && unicode.IsSpace(t.input[t.pos]) {
		// Just return the < as text
		return Token{Type: TokenText, Text: "<", Line: line}
	}

	// If the next char can't start a tag name, treat < as literal text.
	if t.pos < len(t.input) {
		ch := t.input[t.pos]
		if ch == '>' || ch == '-' || (ch >= '0' && ch <= '9') {
			// For <> and <-> operators: consume up to and including >
			// For digits: just return < as text
			if ch == '>' {
				t.pos++ // consume >
				return Token{Type: TokenText, Text: "<>", Line: line}
			}
			if ch == '-' {
				// Scan for > to capture operators like <->
				start := ltPos
				for t.pos < len(t.input) && t.input[t.pos] != '>' && t.input[t.pos] != '<' {
					if t.input[t.pos] == '\n' {
						t.line++
					}
					t.pos++
				}
				if t.pos < len(t.input) && t.input[t.pos] == '>' {
					t.pos++
				}
				return Token{Type: TokenText, Text: string(t.input[start:t.pos]), Line: line}
			}
			// Digit — just return < as text
			return Token{Type: TokenText, Text: "<", Line: line}
		}
	}

	name := t.readName()

	// If name is empty or starts with a non-letter (e.g., ":"
	// from regex bracket expressions like [:alpha:]), treat as text
	if name == "" || !unicode.IsLetter([]rune(name)[0]) {
		// Reset position and return < as text
		t.pos = ltPos + 1
		return Token{Type: TokenText, Text: "<", Line: line}
	}

	attrs := t.readAttributes()

	// Check for self-closing />
	selfClosing := false
	if t.pos < len(t.input) && t.input[t.pos] == '/' {
		selfClosing = true
		t.pos++
	}

	// Consume closing >
	if t.pos < len(t.input) && t.input[t.pos] == '>' {
		t.pos++
	}

	tag := strings.ToLower(name)

	// If self-closing in the source, we still return TagOpen and
	// let the parser handle it as an empty element
	tok := Token{
		Type:  TokenTagOpen,
		Tag:   tag,
		Attrs: attrs,
		Line:  line,
	}

	if selfClosing {
		// Mark it so parser knows no close tag is coming
		if tok.Attrs == nil {
			tok.Attrs = make(map[string]string)
		}
		tok.Attrs["\x00selfclose"] = "1"
	}

	return tok
}

// readName reads an SGML name (letters, digits, hyphens, dots, underscores).
func (t *Tokenizer) readName() string {
	start := t.pos
	for t.pos < len(t.input) {
		ch := t.input[t.pos]
		if unicode.IsLetter(ch) || unicode.IsDigit(ch) ||
			ch == '-' || ch == '.' || ch == '_' || ch == ':' {
			t.pos++
		} else {
			break
		}
	}
	return string(t.input[start:t.pos])
}

// readAttributes reads attribute pairs from a tag.
func (t *Tokenizer) readAttributes() map[string]string {
	attrs := make(map[string]string)
	for {
		t.skipWhitespace()
		if t.pos >= len(t.input) {
			break
		}
		ch := t.input[t.pos]
		if ch == '>' || ch == '/' {
			break
		}

		// Read attribute name
		name := t.readName()
		if name == "" {
			// Skip unexpected character
			t.pos++
			continue
		}
		name = strings.ToLower(name)

		t.skipWhitespace()

		// Check for = and value
		if t.pos < len(t.input) && t.input[t.pos] == '=' {
			t.pos++
			t.skipWhitespace()
			value := t.readAttrValue()
			attrs[name] = value
		} else {
			// Boolean attribute (SGML allows valueless attributes)
			attrs[name] = name
		}
	}
	return attrs
}

// readAttrValue reads a quoted or unquoted attribute value.
func (t *Tokenizer) readAttrValue() string {
	if t.pos >= len(t.input) {
		return ""
	}

	quote := t.input[t.pos]
	if quote == '"' || quote == '\'' {
		t.pos++
		start := t.pos
		for t.pos < len(t.input) && t.input[t.pos] != quote {
			if t.input[t.pos] == '\n' {
				t.line++
			}
			t.pos++
		}
		val := string(t.input[start:t.pos])
		if t.pos < len(t.input) {
			t.pos++ // skip closing quote
		}
		return val
	}

	// Unquoted value (SGML allows this)
	start := t.pos
	for t.pos < len(t.input) &&
		!unicode.IsSpace(t.input[t.pos]) &&
		t.input[t.pos] != '>' {
		t.pos++
	}
	return string(t.input[start:t.pos])
}

// skipWhitespace advances past whitespace, tracking line numbers.
func (t *Tokenizer) skipWhitespace() {
	for t.pos < len(t.input) && unicode.IsSpace(t.input[t.pos]) {
		if t.input[t.pos] == '\n' {
			t.line++
		}
		t.pos++
	}
}

// isLiteralPI checks if what looks like a PI (<?...) is actually
// literal text inside element content (e.g., <literal><?x</literal>).
// Returns true if the content between <? and the next > contains </.
func (t *Tokenizer) isLiteralPI() bool {
	// Scan ahead from <? to find either ?> or </
	for i := t.pos + 2; i < len(t.input)-1; i++ {
		if t.input[i] == '?' && t.input[i+1] == '>' {
			return false // found proper ?> ending — it's a real PI
		}
		if t.input[i] == '<' && t.input[i+1] == '/' {
			return true // found </ before ?> — it's literal text
		}
		if t.input[i] == '\n' {
			// Real PIs don't span many lines in DocBook; if we hit
			// a close tag first, it's literal text
			continue
		}
	}
	return false
}

// isRealDeclaration checks if <! at the current position is a real
// SGML declaration (<!DOCTYPE, <!ENTITY, <!NOTATION, etc.) vs literal
// text like "(?<!" from regex documentation where &lt; was expanded.
func (t *Tokenizer) isRealDeclaration() bool {
	// Must start with <!
	if t.pos+2 >= len(t.input) {
		return false
	}
	ch := t.input[t.pos+2]
	// Real declarations: <!D(OCTYPE), <!E(NTITY), <!N(OTATION),
	// <!A(TTLIST), <![ (marked section)
	return ch == 'D' || ch == 'd' ||
		ch == 'E' || ch == 'e' ||
		ch == 'N' || ch == 'n' ||
		ch == 'A' || ch == 'a' ||
		ch == '['
}

// lookingAt checks if the input at current position starts with s.
func (t *Tokenizer) lookingAt(s string) bool {
	rs := []rune(s)
	if t.pos+len(rs) > len(t.input) {
		return false
	}
	for i, r := range rs {
		if t.input[t.pos+i] != r {
			return false
		}
	}
	return true
}

// String returns a human-readable representation of the token.
func (tok Token) String() string {
	switch tok.Type {
	case TokenTagOpen:
		if len(tok.Attrs) > 0 {
			return fmt.Sprintf("<%s ...> (line %d)", tok.Tag, tok.Line)
		}
		return fmt.Sprintf("<%s> (line %d)", tok.Tag, tok.Line)
	case TokenTagClose:
		return fmt.Sprintf("</%s> (line %d)", tok.Tag, tok.Line)
	case TokenText:
		text := tok.Text
		if len(text) > 40 {
			text = text[:40] + "..."
		}
		return fmt.Sprintf("TEXT(%q) (line %d)", text, tok.Line)
	case TokenComment:
		return fmt.Sprintf("COMMENT (line %d)", tok.Line)
	case TokenPI:
		return fmt.Sprintf("PI (line %d)", tok.Line)
	case TokenEOF:
		return "EOF"
	}
	return "UNKNOWN"
}
