package token

import "fmt"

//go:generate stringer -type=TokenType

type TokenType int

type Token struct {
	Type   TokenType
	Value  string
	Line   int
	Column int
	File   string
}

func (t Token) String() string {
	return fmt.Sprintf("Token{Type: %s, Value: %s, Line: %d, Column: %d, File: %s}", t.Type, t.Value, t.Line, t.Column, t.File)
}

const (
	EOF TokenType = iota
	LPAREN
	RPAREN
	LBRACE
	RBRACE
	LSQUARE
	RSQUARE

	SEMICOLON
	ASSIGN

	IDENT

	PLUS
	STAR
	MINUS
	SLASH
	PERIOD

	FLOAT
	INT
	BOOL
	STRING

	TINT
	TSTRING
	TBOOL
	TFLOAT

	GT
	GTE
	LT
	LTE
	NEQ

	OR
	AND
	BITAND
	BITOR

	ILLEGAL
)
