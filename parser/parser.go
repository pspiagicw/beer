package parser

import (
	"github.com/pspiagicw/osy/ast"
	"github.com/pspiagicw/osy/lexer"
	"github.com/pspiagicw/osy/token"
	"github.com/pspiagicw/osy/types"
)

type Parser struct {
	HasError  bool
	curToken  token.Token
	nextToken token.Token
	l         *lexer.Lexer
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l:         l,
		nextToken: *l.Next(),
	}
	p.advance()
	return p
}

func (p *Parser) advance() {
	p.curToken = p.nextToken
	p.nextToken = *p.l.Next()
}
func (p *Parser) parseReturnStatement() ast.Statement {
	r := &ast.ReturnStatement{}

	r.Value = p.parseExpression()

	return r
}
func (p *Parser) parseExpression() ast.Expression {
	return &ast.Constant{Value: "1", Type: types.Integer}
}

func (p *Parser) Parse() *ast.Program {
	program := &ast.Program{}

	statements := []ast.Statement{}

	for p.curToken.Type != token.EOF && p.curToken.Type != token.ILLEGAL {
		statement := p.parseStatement()
		if statement != nil {
			statements = append(statements, statement)
		}
		p.advance()
	}

	program.Statements = statements
	return program
}
func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.RETURN:
		return p.parseReturnStatement()
	}
	return nil
}

func (p *Parser) Errors() []error {
	return nil
}
