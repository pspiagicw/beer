package token

type TokenType int

type Token struct {
	Type   TokenType
	Value  string
	Line   int
	Column int
	File   string
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

	GT
	GTE
	LT
	LTE
	NEQ

	OR
	AND
)
