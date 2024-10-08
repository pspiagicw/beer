package parser

import (
	"fmt"
	"strings"

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

	errors []error

	prefixParseFunctions map[token.TokenType]func() ast.Expression
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l:                    l,
		nextToken:            *l.Next(),
		prefixParseFunctions: make(map[token.TokenType]func() ast.Expression),
		errors:               []error{},
	}
	p.registerPrefixFn(token.INT, p.parseInteger)
	p.registerPrefixFn(token.FALSE, p.parseBoolean)
	p.registerPrefixFn(token.TRUE, p.parseBoolean)
	p.registerPrefixFn(token.STRING, p.parseString)
	p.registerPrefixFn(token.FLOAT, p.parseFloat)
	p.advance()
	return p
}
func (p *Parser) registerPrefixFn(ttype token.TokenType, fn func() ast.Expression) {
	p.prefixParseFunctions[ttype] = fn
}

func (p *Parser) advance() {
	p.curToken = p.nextToken
	p.nextToken = *p.l.Next()
}
func (p *Parser) registerError(msg string, args ...interface{}) {
	p.HasError = true
	err := fmt.Errorf(msg, args...)
	p.errors = append(p.errors, err)
}
func (p *Parser) expectPeek(tokenType token.TokenType) bool {
	if p.nextToken.Type != tokenType {
		p.registerError("Expected %T got %T", tokenType, p.nextToken.Type)
	}
	p.advance()
	return true
}
func (p *Parser) peekType() token.TokenType {
	return p.nextToken.Type
}
func (p *Parser) curType() token.TokenType {
	return p.curToken.Type
}
func (p *Parser) parseReturnStatement() ast.Statement {
	r := &ast.ReturnStatement{}

	p.advance()

	r.Value = p.parseExpression()

	return r
}
func (p *Parser) parseExpression() ast.Expression {
	prefixFn := p.prefixParseFunctions[p.curToken.Type]

	if prefixFn != nil {
		return prefixFn()
	} else {
		p.registerError("No prefix parse fn found for %s", p.curToken.Type)
	}

	return nil
}

//	func (p *Parser) parseBlockStatement() ast.Statement {
//		if p.curType() != token.LPAREN {
//			p.registerError("Can't parse block statement, current token is not parenthesis")
//			return nil
//		}
//
//	    if p.curToken.Type != token.LBRACE
//
//		b := &ast.BlockStatement{}
//
//		for p.curToken.Type != token.EOF && p.curToken.Type != token.RPAREN {
//		}
//	}
//
//	func (p *Parser) parseFunctionDeclaration() ast.Statement {
//		f := &ast.FunctionDeclaration{}
//
//		if !p.expectPeek(token.IDENT) {
//			return nil
//		}
//
//		f.Name = p.curToken.Value
//
//		p.advance()
//
//		if p.curType() == token.LBRACE {
//			f.ReturnType = types.Void
//		} else {
//			f.ReturnType = p.parseType()
//		}
//
//		f.Body = p.parseBlockStatement()
//
//		return f
//	}
func (p *Parser) parseAssignmentStatement() ast.Statement {
	a := &ast.Assignment{}

	if !p.expectPeek(token.IDENT) {
		return nil
	}

	a.Name = p.curToken.Value

	// Skip over the = sign
	p.advance()
	p.advance()

	a.Value = p.parseExpression()

	return a
}
func (p *Parser) parseType() types.Type {
	switch p.curToken.Type {
	case token.TINT:
		return types.Integer
	case token.TSTRING:
		return types.String
	case token.TBOOL:
		return types.Boolean
	case token.TFLOAT:
		return types.Float
	default:
		p.registerError("No type defined for %s", p.curToken.Type)
		return types.Illegal
	}
}
func (p *Parser) parseDeclarationStatement() ast.Statement {
	d := &ast.VariableDeclaration{}

	if !p.expectPeek(token.IDENT) {
		return nil
	}

	d.Name = p.curToken.Value

	p.advance()

	d.Type = p.parseType()

	return d
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
	case token.LET:
		return p.parseAssignmentStatement()
	case token.DEF:
		return p.parseDeclarationStatement()
	}
	return nil
}

func (p *Parser) Errors() []error {
	return p.errors
}
func (p *Parser) parseInteger() ast.Expression {

	c := &ast.Constant{
		Value: p.curToken.Value,
		Type:  types.Integer,
	}

	p.advance()

	return c
}
func (p *Parser) parseBoolean() ast.Expression {
	b := &ast.Constant{
		Value: p.curToken.Value,
		Type:  types.Boolean,
	}

	p.advance()

	return b
}
func (p *Parser) parseString() ast.Expression {

	val := strings.Trim(p.curToken.Value, "\"")
	s := &ast.Constant{
		Value: val,
		Type:  types.String,
	}

	return s
}
func (p *Parser) parseFloat() ast.Expression {
	f := &ast.Constant{
		Value: p.curToken.Value,
		Type:  types.Float,
	}
	return f
}
