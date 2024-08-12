package lexer

import "github.com/pspiagicw/osy/token"

type Lexer struct {
}

func New(input string) *Lexer {
	return &Lexer{}
}
func (l *Lexer) Next() *token.Token {
	return nil
}
