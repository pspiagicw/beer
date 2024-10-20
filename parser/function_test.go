package parser

import (
	"testing"

	"github.com/pspiagicw/osy/ast"
	"github.com/pspiagicw/osy/types"
)

func TestFunctionDefinition(t *testing.T) {
	t.Skip()
	input := `fn main() {}`

	ast := &ast.Program{
		Statements: []ast.Statement{
			&ast.FunctionDeclaration{
				ReturnType: types.Void,
				Name:       "main",
				Body: &ast.BlockStatement{
					Statements: []ast.Statement{},
				},
			},
		},
	}

	testParser(t, input, ast)
}
