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
