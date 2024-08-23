package parser

import (
	"github.com/pspiagicw/osy/ast"
	"github.com/pspiagicw/osy/lexer"
)

type Parser struct {
	HasError bool
}

func New(l *lexer.Lexer) *Parser {
	return &Parser{}
}

func (p *Parser) Parse() *ast.Program {
	return nil
}
func (p *Parser) Errors() []error {
	return nil
}
