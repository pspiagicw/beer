package parser

import (
	"testing"

	"github.com/gkampitakis/go-snaps/snaps"
	"github.com/kr/pretty"
	"github.com/pspiagicw/osy/ast"
	"github.com/pspiagicw/osy/lexer"
	"github.com/pspiagicw/osy/types"
	"github.com/stretchr/testify/assert"
)

func TestVariableAssignmentBoolean(t *testing.T) {
	input := `let status = false;`

	ast := &ast.Program{
		Statements: []ast.Statement{
			&ast.Assignment{
				Name: "status",
				Value: &ast.Constant{
					Value: "false",
					Type:  types.Boolean,
				},
			},
		},
	}

	testParser(t, input, ast)
}

func TestVariableAssignmentString(t *testing.T) {
	input := `let name = "pspiagicw";`

	ast := &ast.Program{
		Statements: []ast.Statement{
			&ast.Assignment{
				Name: "name",
				Value: &ast.Constant{
					Value: "pspiagicw",
					Type:  types.String,
				},
			},
		},
	}

	testParser(t, input, ast)
}

func TestVariableAssignmentInteger(t *testing.T) {
	input := `let a = 1;`

	ast := &ast.Program{
		Statements: []ast.Statement{
			&ast.Assignment{
				Name: "a",
				Value: &ast.Constant{
					Value: "1",
					Type:  types.Integer,
				},
			},
		},
	}

	testParser(t, input, ast)
}

func TestVariableDeclarationString(t *testing.T) {
	input := `def name string;`

	ast := &ast.Program{
		Statements: []ast.Statement{
			&ast.VariableDeclaration{
				Name: "name",
				Type: types.String,
			},
		},
	}

	testParser(t, input, ast)
}

func TestVariableDeclarationBoolean(t *testing.T) {
	input := `def status bool;`

	ast := &ast.Program{
		Statements: []ast.Statement{
			&ast.VariableDeclaration{
				Name: "status",
				Type: types.Boolean,
			},
		},
	}

	testParser(t, input, ast)
}

func TestVariableDeclarationInteger(t *testing.T) {
	input := `def a int;`

	ast := &ast.Program{
		Statements: []ast.Statement{
			&ast.VariableDeclaration{
				Name: "a",
				Type: types.Integer,
			},
		},
	}

	testParser(t, input, ast)
}

func TestReturnBoolean(t *testing.T) {
	input := `return true;`

	ast := &ast.Program{
		Statements: []ast.Statement{
			&ast.ReturnStatement{
				Value: &ast.Constant{
					Value: "true",
					Type:  types.Boolean,
				},
			},
		},
	}

	testParser(t, input, ast)
}

func TestReturnString(t *testing.T) {
	input := `return "hello";`

	ast := &ast.Program{
		Statements: []ast.Statement{
			&ast.ReturnStatement{
				Value: &ast.Constant{
					Value: "hello",
					Type:  types.String,
				},
			},
		},
	}

	testParser(t, input, ast)
}

func TestReturnFloat(t *testing.T) {
	input := `return 3.14;`

	ast := &ast.Program{
		Statements: []ast.Statement{
			&ast.ReturnStatement{
				Value: &ast.Constant{
					Value: "3.14",
					Type:  types.Float,
				},
			},
		},
	}

	testParser(t, input, ast)
}
func TestReturnInteger(t *testing.T) {
	input := `return 1;`

	ast := &ast.Program{
		Statements: []ast.Statement{
			&ast.ReturnStatement{
				Value: &ast.Constant{
					Value: "1",
					Type:  types.Integer,
				},
			},
		},
	}

	testParser(t, input, ast)
}

func testParser(t *testing.T, input string, expected ast.Node) {
	l := lexer.New(input)
	p := New(l)

	program := p.Parse()

	if p.HasError {
		for _, error := range p.Errors() {
			t.Errorf("PARSER ERROR: %v", error)
		}
		t.Fatalf("Parser Errors!")
	}

	if !assert.Equal(t, expected, program) {
		pretty.Println(program)
	}
	snaps.MatchSnapshot(t, program)
}
