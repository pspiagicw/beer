package parser

import (
	"testing"

	"github.com/gkampitakis/go-snaps/snaps"
	"github.com/pspiagicw/osy/ast"
	"github.com/pspiagicw/osy/lexer"
	"github.com/pspiagicw/osy/token"
	"github.com/pspiagicw/osy/types"
	"github.com/stretchr/testify/assert"
)

func TestReturnNestedBooleanComparisonExpressions(t *testing.T) {
	input := `return (3 > 2) && (true == false) || (4 <= 5);`

	ast := &ast.Program{
		Statements: []ast.Statement{
			&ast.ReturnStatement{
				Value: &ast.BinaryExpression{
					Left: &ast.BinaryExpression{
						Left: &ast.BinaryExpression{
							Left: &ast.Constant{
								Value: "3",
								Type:  types.Integer,
							},
							Operator: token.LT,
							Right: &ast.Constant{
								Value: "2",
								Type:  types.Integer,
							},
						},
						Operator: token.AND,
						Right: &ast.BinaryExpression{
							Left: &ast.Constant{
								Value: "true",
								Type:  types.Boolean,
							},
							Operator: token.EQ,
							Right: &ast.Constant{
								Value: "false",
								Type:  types.Boolean,
							},
						},
					},
					Operator: token.AND,
					Right: &ast.BinaryExpression{
						Left: &ast.Constant{
							Value: "4",
							Type:  types.Integer,
						},
						Operator: token.GTE,
						Right: &ast.Constant{
							Value: "5",
							Type:  types.Integer,
						},
					},
				},
			},
		},
	}

	testParser(t, input, ast)
}

func TestReturnBooleanLogicalOperators(t *testing.T) {
	input := `return true && false || true;`

	ast := &ast.Program{
		Statements: []ast.Statement{
			&ast.ReturnStatement{
				Value: &ast.BinaryExpression{
					Left: &ast.BinaryExpression{
						Left: &ast.Constant{
							Value: "true",
							Type:  types.Boolean,
						},
						Operator: token.AND,
						Right: &ast.Constant{
							Value: "false",
							Type:  types.Boolean,
						},
					},
					Operator: token.AND,
					Right: &ast.Constant{
						Value: "true",
						Type:  types.Boolean,
					},
				},
			},
		},
	}

	testParser(t, input, ast)
}

func TestReturnLogicalOperatorsWithComparison(t *testing.T) {
	input := `return (5 > 3) && (4 <= 4) || (2 != 1);`

	ast := &ast.Program{
		Statements: []ast.Statement{
			&ast.ReturnStatement{
				Value: &ast.BinaryExpression{
					Left: &ast.BinaryExpression{
						Left: &ast.BinaryExpression{
							Left: &ast.Constant{
								Value: "5",
								Type:  types.Integer,
							},
							Operator: token.LT,
							Right: &ast.Constant{
								Value: "3",
								Type:  types.Integer,
							},
						},
						Operator: token.AND,
						Right: &ast.BinaryExpression{
							Left: &ast.Constant{
								Value: "4",
								Type:  types.Integer,
							},
							Operator: token.GTE,
							Right: &ast.Constant{
								Value: "4",
								Type:  types.Integer,
							},
						},
					},
					Operator: token.AND,
					Right: &ast.BinaryExpression{
						Left: &ast.Constant{
							Value: "2",
							Type:  types.Integer,
						},
						Operator: token.NEQ,
						Right: &ast.Constant{
							Value: "1",
							Type:  types.Integer,
						},
					},
				},
			},
		},
	}

	testParser(t, input, ast)
}

func TestReturnMixedNumericAndBoolean(t *testing.T) {
	input := `return 2 > 1 == true;`

	ast := &ast.Program{
		Statements: []ast.Statement{
			&ast.ReturnStatement{
				Value: &ast.BinaryExpression{
					Left: &ast.BinaryExpression{
						Left: &ast.Constant{
							Value: "2",
							Type:  types.Integer,
						},
						Operator: token.LT,
						Right: &ast.Constant{
							Value: "1",
							Type:  types.Integer,
						},
					},
					Operator: token.EQ,
					Right: &ast.Constant{
						Value: "true",
						Type:  types.Boolean,
					},
				},
			},
		},
	}

	testParser(t, input, ast)
}

func TestReturnBooleanAndComparisonOperators(t *testing.T) {
	input := `return true == false != true;`

	ast := &ast.Program{
		Statements: []ast.Statement{
			&ast.ReturnStatement{
				Value: &ast.BinaryExpression{
					Left: &ast.BinaryExpression{
						Left: &ast.Constant{
							Value: "true",
							Type:  types.Boolean,
						},
						Operator: token.EQ,
						Right: &ast.Constant{
							Value: "false",
							Type:  types.Boolean,
						},
					},
					Operator: token.NEQ,
					Right: &ast.Constant{
						Value: "true",
						Type:  types.Boolean,
					},
				},
			},
		},
	}

	testParser(t, input, ast)
}

func TestReturnEqualityOperators(t *testing.T) {
	input := `return 5 == 5 != 4 == 3;`

	ast := &ast.Program{
		Statements: []ast.Statement{
			&ast.ReturnStatement{
				Value: &ast.BinaryExpression{
					Left: &ast.BinaryExpression{
						Left: &ast.BinaryExpression{
							Left: &ast.Constant{
								Value: "5",
								Type:  types.Integer,
							},
							Operator: token.EQ,
							Right: &ast.Constant{
								Value: "5",
								Type:  types.Integer,
							},
						},
						Operator: token.NEQ,
						Right: &ast.Constant{
							Value: "4",
							Type:  types.Integer,
						},
					},
					Operator: token.EQ,
					Right: &ast.Constant{
						Value: "3",
						Type:  types.Integer,
					},
				},
			},
		},
	}

	testParser(t, input, ast)
}

func TestReturnComparisonOperators(t *testing.T) {
	input := `return 3 > 2 < 4 <= 5 >= 1;`

	ast := &ast.Program{
		Statements: []ast.Statement{
			&ast.ReturnStatement{
				Value: &ast.BinaryExpression{
					Left: &ast.BinaryExpression{
						Left: &ast.BinaryExpression{
							Left: &ast.Constant{
								Value: "3",
								Type:  types.Integer,
							},
							Operator: token.LT,
							Right: &ast.Constant{
								Value: "2",
								Type:  types.Integer,
							},
						},
						Operator: token.GT,
						Right: &ast.Constant{
							Value: "4",
							Type:  types.Integer,
						},
					},
					Operator: token.GTE,
					Right: &ast.BinaryExpression{
						Left: &ast.Constant{
							Value: "5",
							Type:  types.Integer,
						},
						Operator: token.LTE,
						Right: &ast.Constant{
							Value: "1",
							Type:  types.Integer,
						},
					},
				},
			},
		},
	}

	testParser(t, input, ast)
}

func TestReturnNestedBinaryOperations(t *testing.T) {
	input := `return (1 + 2) * (3 - 4) / 5;`

	ast := &ast.Program{
		Statements: []ast.Statement{
			&ast.ReturnStatement{
				Value: &ast.BinaryExpression{
					Left: &ast.BinaryExpression{
						Left: &ast.BinaryExpression{
							Left: &ast.Constant{
								Value: "1",
								Type:  types.Integer,
							},
							Operator: token.PLUS,
							Right: &ast.Constant{
								Value: "2",
								Type:  types.Integer,
							},
						},
						Operator: token.STAR,
						Right: &ast.BinaryExpression{
							Left: &ast.Constant{
								Value: "3",
								Type:  types.Integer,
							},
							Operator: token.MINUS,
							Right: &ast.Constant{
								Value: "4",
								Type:  types.Integer,
							},
						},
					},
					Operator: token.SLASH,
					Right: &ast.Constant{
						Value: "5",
						Type:  types.Integer,
					},
				},
			},
		},
	}

	testParser(t, input, ast)
}

func TestReturnMultipleOperatorsPrecedence(t *testing.T) {
	input := `return 2 + 3 * 4 - 5 / 2;`

	ast := &ast.Program{
		Statements: []ast.Statement{
			&ast.ReturnStatement{
				Value: &ast.BinaryExpression{
					Left: &ast.BinaryExpression{
						Left: &ast.BinaryExpression{
							Left: &ast.Constant{
								Value: "2",
								Type:  types.Integer,
							},
							Operator: token.PLUS,
							Right: &ast.BinaryExpression{
								Left: &ast.Constant{
									Value: "3",
									Type:  types.Integer,
								},
								Operator: token.STAR,
								Right: &ast.Constant{
									Value: "4",
									Type:  types.Integer,
								},
							},
						},
						Operator: token.MINUS,
						Right: &ast.BinaryExpression{
							Left: &ast.Constant{
								Value: "5",
								Type:  types.Integer,
							},
							Operator: token.SLASH,
							Right: &ast.Constant{
								Value: "2",
								Type:  types.Integer,
							},
						},
					},
				},
			},
		},
	}

	testParser(t, input, ast)
}

func TestReturnParenthesesOverridePrecedence(t *testing.T) {
	input := `return (2 + 3) * 4;`

	ast := &ast.Program{
		Statements: []ast.Statement{
			&ast.ReturnStatement{
				Value: &ast.BinaryExpression{
					Left: &ast.BinaryExpression{
						Left: &ast.Constant{
							Value: "2",
							Type:  types.Integer,
						},
						Operator: token.PLUS,
						Right: &ast.Constant{
							Value: "3",
							Type:  types.Integer,
						},
					},
					Operator: token.STAR,
					Right: &ast.Constant{
						Value: "4",
						Type:  types.Integer,
					},
				},
			},
		},
	}

	testParser(t, input, ast)
}

func TestReturnMultiplicationAndDivisionPrecedence(t *testing.T) {
	input := `return 2 + 3 * 4;`

	ast := &ast.Program{
		Statements: []ast.Statement{
			&ast.ReturnStatement{
				Value: &ast.BinaryExpression{
					Left: &ast.Constant{
						Value: "2",
						Type:  types.Integer,
					},
					Operator: token.PLUS,
					Right: &ast.BinaryExpression{
						Left: &ast.Constant{
							Value: "3",
							Type:  types.Integer,
						},
						Operator: token.STAR,
						Right: &ast.Constant{
							Value: "4",
							Type:  types.Integer,
						},
					},
				},
			},
		},
	}

	testParser(t, input, ast)
}

func TestReturnAdditionAndSubtraction(t *testing.T) {
	input := `return 2 + 3 - 1;`

	ast := &ast.Program{
		Statements: []ast.Statement{
			&ast.ReturnStatement{
				Value: &ast.BinaryExpression{
					Left: &ast.BinaryExpression{
						Left: &ast.Constant{
							Value: "2",
							Type:  types.Integer,
						},
						Operator: token.PLUS,
						Right: &ast.Constant{
							Value: "3",
							Type:  types.Integer,
						},
					},
					Operator: token.MINUS,
					Right: &ast.Constant{
						Value: "1",
						Type:  types.Integer,
					},
				},
			},
		},
	}

	testParser(t, input, ast)
}

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

	assert.Equal(t, expected, program)
	snaps.MatchSnapshot(t, program)
}
