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

	PLUS  //  10
	STAR  // 11
	MINUS // 12
	SLASH // 13
	PERIOD

	FLOAT // 15
	INT
	STRING
	TRUE
	FALSE // 19

	TINT // 20
	TSTRING
	TBOOL
	TFLOAT // 23

	GT  // 24
	GTE // 25
	LT  // 26
	LTE // 27
	EQ
	NEQ
	NEGATE

	OR
	AND
	BITAND
	BITOR

	IF
	ELSE
	FOR
	WHILE
	RETURN
	FN
	LET
	DEF

	ILLEGAL
)
