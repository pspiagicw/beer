package lexer

import (
	"github.com/pspiagicw/osy/token"
)

type Lexer struct {
	input   string
	curPos  int
	readPos int
	line    int
	column  int
	file    string

	EOF bool
}

func New(input string) *Lexer {
	return &Lexer{
		input:   input,
		curPos:  -1,
		readPos: 0,
		column:  -1,
		line:    0,
	}
}
func (l *Lexer) newToken() *token.Token {
	return &token.Token{
		Value:  "",
		Type:   token.EOF,
		Line:   l.line,
		Column: l.column,
		File:   l.file,
	}
}
func (l *Lexer) advance() {
	if l.readPos >= len(l.input) {
		l.EOF = true
	}
	l.column += 1
	l.curPos = l.readPos
	l.readPos += 1

}
func (l *Lexer) currentChar() string {
	return string(l.input[l.curPos])
}
func (l *Lexer) isWhitespace() bool {
	return l.currentChar() == " " || l.currentChar() == "\t" || l.currentChar() == "\n"
}

func (l *Lexer) whitespace() {
	for !l.EOF && l.isWhitespace() {
		if l.currentChar() == "\n" {
			l.line += 1
			l.column = 0
		} else if l.currentChar() == "\t" {
			l.column += 4
		} else {
			l.column += 1
		}
		l.advance()
	}
}
func (l *Lexer) peek() string {
	if l.readPos >= len(l.input) {
		return ""
	}
	return string(l.input[l.readPos])
}
func (l *Lexer) isDigit(char string) bool {
	return char >= "0" && char <= "9"
}

func (l *Lexer) number() string {
	start := l.curPos

	for !l.EOF && l.isDigit(l.peek()) {
		l.advance()
	}

	return l.input[start : l.curPos+1]
}
func (l *Lexer) Next() *token.Token {

	l.advance()
	l.whitespace()

	t := l.newToken()

	if l.EOF {
		return t
	}

	currentChar := l.currentChar()

	switch currentChar {
	case "+":
		t.Type = token.PLUS
		t.Value = currentChar
	case "-":
		t.Type = token.MINUS
		t.Value = currentChar
	case "*":
		t.Type = token.STAR
		t.Value = currentChar
	case "/":
		t.Type = token.SLASH
		t.Value = currentChar
	case "(":
		t.Type = token.LPAREN
		t.Value = currentChar
	case ")":
		t.Type = token.RPAREN
		t.Value = currentChar
	case "{":
		t.Type = token.LBRACE
		t.Value = currentChar
	case "}":
		t.Type = token.RBRACE
		t.Value = currentChar
	case "[":
		t.Type = token.LSQUARE
		t.Value = currentChar
	case "]":
		t.Type = token.RSQUARE
		t.Value = currentChar
	case "&":
		if l.peek() == "&" {
			l.advance()
			t.Type = token.AND
			t.Value = "&&"
		} else {
			t.Type = token.BITAND
			t.Value = currentChar
		}
	case "|":
		if l.peek() == "|" {
			l.advance()
			t.Type = token.OR
			t.Value = "||"
		} else {
			t.Type = token.BITOR
			t.Value = currentChar
		}
	default:
		if l.isDigit(currentChar) {
			t.Value = l.number()
			t.Type = token.INT
		} else {
			t.Type = token.ILLEGAL
			t.Value = string(currentChar)
		}
	}

	return t
}
