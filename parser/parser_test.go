package parser

import (
	"testing"

	"github.com/gkampitakis/go-snaps/snaps"
	"github.com/pspiagicw/osy/ast"
	"github.com/pspiagicw/osy/lexer"
	"github.com/pspiagicw/osy/types"
	"github.com/stretchr/testify/assert"
)

func TestReturn(t *testing.T) {
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
