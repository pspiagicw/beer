package ast

import (
	"fmt"
	"strings"

	"github.com/pspiagicw/osy/types"
)

type Node interface {
	String() string
	Pretty() string
}

type Expression interface {
	Node
	expressionNode()
}

type Statement interface {
	Node
	statementNode()
}

type Program struct {
	Statements []Statement
}

func (p Program) statementNode() {}
func (p Program) String() string {
	return fmt.Sprintf("Program with %d statements", len(p.Statements))
}
func (p Program) Pretty() string {
	var buffer strings.Builder

	for _, statement := range p.Statements {
		buffer.WriteString(statement.Pretty())
		buffer.WriteString("\n")
	}

	return buffer.String()
}

type ReturnStatement struct {
	Value Expression
}

func (r ReturnStatement) statementNode() {}
func (r ReturnStatement) String() string {
	return fmt.Sprintf(`Return statement;`)
}
func (r ReturnStatement) Pretty() string {
	return fmt.Sprintf("return %s;", r.Value.Pretty())
}

type Constant struct {
	Value string
	Type  types.Type
}

func (c Constant) expressionNode() {}
func (c Constant) String() string {
	return fmt.Sprintf("Value %s of type %s", c.Value, c.Type)
}
func (c Constant) Pretty() string {
	return c.Value
}

type BlockStatement struct {
	Statements []Statement
}

func (b BlockStatement) String() string {
	return fmt.Sprintf("Block with %d statements", len(b.Statements))
}
func (b BlockStatement) Pretty() string {
	// TODO: Write a pretty printer for BlockStatement

	return "\n"
}

type FunctionDeclaration struct {
	ReturnType types.Type
	Name       string
	Body       *BlockStatement
}

func (f FunctionDeclaration) statementNode() {}
func (f FunctionDeclaration) String() string {
	return fmt.Sprintf("Function %s with return type %T", f.Name, f.ReturnType)
}
func (f FunctionDeclaration) Pretty() string {
	var buffer strings.Builder

	// TOOD: Complete the pretty printer for the FunctionDeclaration

	buffer.WriteString(f.Body.Pretty())

	return buffer.String()
}

type Assignment struct {
	Name  string
	Value Expression
}

func (a Assignment) statementNode() {}
func (a Assignment) String() string {
	return fmt.Sprintf("Assignment for %s", a.Name)
}
func (a Assignment) Pretty() string {
	// TODO: Pretty print assignment statements.

	return "pretty assignment"
}

type VariableDeclaration struct {
	Name string
	Type types.Type
}

func (v VariableDeclaration) statementNode() {}
func (v VariableDeclaration) String() string {
	return fmt.Sprintf("declared %s of type %T", v.Name, v.Type)
}
func (v VariableDeclaration) Pretty() string {
	// TOOD: Pretty print Variable Declaration
	return "pretty variable declaration"
}
