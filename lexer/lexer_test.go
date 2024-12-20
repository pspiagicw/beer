package lexer

import (
	"github.com/pspiagicw/osy/token"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRoundParentheses(t *testing.T) {
	input := `()`
	expectedTokens := []token.Token{
		{Type: token.LPAREN, Value: "("},
		{Type: token.RPAREN, Value: ")"},
		{Type: token.EOF, Value: ""},
	}
	testTokens(t, input, expectedTokens)
}

func TestCurlyBraces(t *testing.T) {
	input := `{}`
	expectedTokens := []token.Token{
		{Type: token.LBRACE, Value: "{"},
		{Type: token.RBRACE, Value: "}"},
		{Type: token.EOF, Value: ""},
	}
	testTokens(t, input, expectedTokens)
}

func TestSquareBrackets(t *testing.T) {
	input := `[]`
	expectedTokens := []token.Token{
		{Type: token.LSQUARE, Value: "["},
		{Type: token.RSQUARE, Value: "]"},
		{Type: token.EOF, Value: ""},
	}
	testTokens(t, input, expectedTokens)
}

func TestArithmeticOperators(t *testing.T) {
	input := `+ - * / .`
	expectedTokens := []token.Token{
		{Type: token.PLUS, Value: "+"},
		{Type: token.MINUS, Value: "-"},
		{Type: token.STAR, Value: "*"},
		{Type: token.SLASH, Value: "/"},
		{Type: token.PERIOD, Value: "."},
		{Type: token.EOF, Value: ""},
	}
	testTokens(t, input, expectedTokens)
}

func TestLogicalOperators(t *testing.T) {
	input := `& | && ||`
	expectedTokens := []token.Token{
		{Type: token.BITAND, Value: "&"},
		{Type: token.BITOR, Value: "|"},
		{Type: token.AND, Value: "&&"},
		{Type: token.OR, Value: "||"},
		{Type: token.EOF, Value: ""},
	}
	testTokens(t, input, expectedTokens)
}

func TestIntegerLiteral(t *testing.T) {
	input := `123`
	expectedTokens := []token.Token{
		{Type: token.INT, Value: "123"},
		{Type: token.EOF, Value: ""},
	}
	testTokens(t, input, expectedTokens)
}

func TestFloatingPointLiteral(t *testing.T) {
	input := `3.14`
	expectedTokens := []token.Token{
		{Type: token.FLOAT, Value: "3.14"},
		{Type: token.EOF, Value: ""},
	}
	testTokens(t, input, expectedTokens)
}
func TestMultipleFloatingPointLiterals(t *testing.T) {
	input := `3.14 2.71.23.1`
	expectedTokens := []token.Token{
		{Type: token.FLOAT, Value: "3.14"},
		{Type: token.FLOAT, Value: "2.71"},
		{Type: token.PERIOD, Value: "."},
		{Type: token.FLOAT, Value: "23.1"},
		{Type: token.EOF, Value: ""},
	}
	testTokens(t, input, expectedTokens)
}

func TestBooleanLiteral(t *testing.T) {
	input := `true false`
	expectedTokens := []token.Token{
		{Type: token.TRUE, Value: "true"},
		{Type: token.FALSE, Value: "false"},
		{Type: token.EOF, Value: ""},
	}
	testTokens(t, input, expectedTokens)
}

func TestStringLiterals(t *testing.T) {
	input := `"single quoted" "double quoted"`
	expectedTokens := []token.Token{
		{Type: token.STRING, Value: `single quoted`},
		{Type: token.STRING, Value: `double quoted`},
		{Type: token.EOF, Value: ""},
	}
	testTokens(t, input, expectedTokens)
}

func TestComparisonOperators1(t *testing.T) {
	input := `a < b`
	expectedTokens := []token.Token{
		{Type: token.IDENT, Value: "a"},
		{Type: token.LT, Value: "<"},
		{Type: token.IDENT, Value: "b"},
		{Type: token.EOF, Value: ""},
	}
	testTokens(t, input, expectedTokens)
}

func TestComparisonOperators2(t *testing.T) {
	input := `x > y`
	expectedTokens := []token.Token{
		{Type: token.IDENT, Value: "x"},
		{Type: token.GT, Value: ">"},
		{Type: token.IDENT, Value: "y"},
		{Type: token.EOF, Value: ""},
	}
	testTokens(t, input, expectedTokens)
}

func TestComparisonOperators3(t *testing.T) {
	input := `m <= n`
	expectedTokens := []token.Token{
		{Type: token.IDENT, Value: "m"},
		{Type: token.LTE, Value: "<="},
		{Type: token.IDENT, Value: "n"},
		{Type: token.EOF, Value: ""},
	}
	testTokens(t, input, expectedTokens)
}

func TestComparisonOperators4(t *testing.T) {
	input := `p >= q`
	expectedTokens := []token.Token{
		{Type: token.IDENT, Value: "p"},
		{Type: token.GTE, Value: ">="},
		{Type: token.IDENT, Value: "q"},
		{Type: token.EOF, Value: ""},
	}
	testTokens(t, input, expectedTokens)
}

func TestComparisonOperators5(t *testing.T) {
	input := `r != s r == s`
	expectedTokens := []token.Token{
		{Type: token.IDENT, Value: "r"},
		{Type: token.NEQ, Value: "!="},
		{Type: token.IDENT, Value: "s"},
		{Type: token.IDENT, Value: "r"},
		{Type: token.EQ, Value: "=="},
		{Type: token.IDENT, Value: "s"},
		{Type: token.EOF, Value: ""},
	}
	testTokens(t, input, expectedTokens)
}

func TestComparisonOperatorsCombined(t *testing.T) {
	input := `a < b && c > d || e <= f && g >= h || i != j`
	expectedTokens := []token.Token{
		{Type: token.IDENT, Value: "a"},
		{Type: token.LT, Value: "<"},
		{Type: token.IDENT, Value: "b"},
		{Type: token.AND, Value: "&&"},
		{Type: token.IDENT, Value: "c"},
		{Type: token.GT, Value: ">"},
		{Type: token.IDENT, Value: "d"},
		{Type: token.OR, Value: "||"},
		{Type: token.IDENT, Value: "e"},
		{Type: token.LTE, Value: "<="},
		{Type: token.IDENT, Value: "f"},
		{Type: token.AND, Value: "&&"},
		{Type: token.IDENT, Value: "g"},
		{Type: token.GTE, Value: ">="},
		{Type: token.IDENT, Value: "h"},
		{Type: token.OR, Value: "||"},
		{Type: token.IDENT, Value: "i"},
		{Type: token.NEQ, Value: "!="},
		{Type: token.IDENT, Value: "j"},
		{Type: token.EOF, Value: ""},
	}
	testTokens(t, input, expectedTokens)
}

func TestComparisonOperators(t *testing.T) {
	input := `a < b && c >= 5 || d != 10`
	expectedTokens := []token.Token{
		{Type: token.IDENT, Value: "a"},
		{Type: token.LT, Value: "<"},
		{Type: token.IDENT, Value: "b"},
		{Type: token.AND, Value: "&&"},
		{Type: token.IDENT, Value: "c"},
		{Type: token.GTE, Value: ">="},
		{Type: token.INT, Value: "5"},
		{Type: token.OR, Value: "||"},
		{Type: token.IDENT, Value: "d"},
		{Type: token.NEQ, Value: "!="},
		{Type: token.INT, Value: "10"},
		{Type: token.EOF, Value: ""},
	}
	testTokens(t, input, expectedTokens)
}

func TestSingleIdentifier(t *testing.T) {
	input := `varName`
	expectedTokens := []token.Token{
		{Type: token.IDENT, Value: "varName"},
		{Type: token.EOF, Value: ""},
	}
	testTokens(t, input, expectedTokens)
}

func TestMultipleIdentifiers(t *testing.T) {
	input := `varName anotherVar`
	expectedTokens := []token.Token{
		{Type: token.IDENT, Value: "varName"},
		{Type: token.IDENT, Value: "anotherVar"},
		{Type: token.EOF, Value: ""},
	}
	testTokens(t, input, expectedTokens)
}

func TestTypes(t *testing.T) {
	input := `int string bool float`

	expectedTokens := []token.Token{
		{Type: token.TINT, Value: "int"},
		{Type: token.TSTRING, Value: "string"},
		{Type: token.TBOOL, Value: "bool"},
		{Type: token.TFLOAT, Value: "float"},
	}

	testTokens(t, input, expectedTokens)
}

func TestCombinedElements1(t *testing.T) {
	input := `varName = 123 + 3.14;`
	expectedTokens := []token.Token{
		{Type: token.IDENT, Value: "varName"},
		{Type: token.ASSIGN, Value: "="},
		{Type: token.INT, Value: "123"},
		{Type: token.PLUS, Value: "+"},
		{Type: token.FLOAT, Value: "3.14"},
		{Type: token.SEMICOLON, Value: ";"},
		{Type: token.EOF, Value: ""},
	}
	testTokens(t, input, expectedTokens)
}

func TestCombinedElements2(t *testing.T) {
	input := `if (x > 10) { y = true; }`
	expectedTokens := []token.Token{
		{Type: token.IF, Value: "if"},
		{Type: token.LPAREN, Value: "("},
		{Type: token.IDENT, Value: "x"},
		{Type: token.GT, Value: ">"},
		{Type: token.INT, Value: "10"},
		{Type: token.RPAREN, Value: ")"},
		{Type: token.LBRACE, Value: "{"},
		{Type: token.IDENT, Value: "y"},
		{Type: token.ASSIGN, Value: "="},
		{Type: token.TRUE, Value: "true"},
		{Type: token.SEMICOLON, Value: ";"},
		{Type: token.RBRACE, Value: "}"},
		{Type: token.EOF, Value: ""},
	}
	testTokens(t, input, expectedTokens)
}

func TestCombinedElements3(t *testing.T) {
	input := `result = "hello" + "world";`
	expectedTokens := []token.Token{
		{Type: token.IDENT, Value: "result"},
		{Type: token.ASSIGN, Value: "="},
		{Type: token.STRING, Value: `hello`},
		{Type: token.PLUS, Value: "+"},
		{Type: token.STRING, Value: `world`},
		{Type: token.SEMICOLON, Value: ";"},
		{Type: token.EOF, Value: ""},
	}
	testTokens(t, input, expectedTokens)
}

func TestCombinedElements4(t *testing.T) {
	input := `print("result is " + (x * 2));`
	expectedTokens := []token.Token{
		{Type: token.IDENT, Value: "print"},
		{Type: token.LPAREN, Value: "("},
		{Type: token.STRING, Value: `result is `},
		{Type: token.PLUS, Value: "+"},
		{Type: token.LPAREN, Value: "("},
		{Type: token.IDENT, Value: "x"},
		{Type: token.STAR, Value: "*"},
		{Type: token.INT, Value: "2"},
		{Type: token.RPAREN, Value: ")"},
		{Type: token.RPAREN, Value: ")"},
		{Type: token.SEMICOLON, Value: ";"},
		{Type: token.EOF, Value: ""},
	}
	testTokens(t, input, expectedTokens)
}

func TestDotSymbol(t *testing.T) {
	input := `object.method().field`
	expectedTokens := []token.Token{
		{Type: token.IDENT, Value: "object"},
		{Type: token.PERIOD, Value: "."},
		{Type: token.IDENT, Value: "method"},
		{Type: token.LPAREN, Value: "("},
		{Type: token.RPAREN, Value: ")"},
		{Type: token.PERIOD, Value: "."},
		{Type: token.IDENT, Value: "field"},
		{Type: token.EOF, Value: ""},
	}
	testTokens(t, input, expectedTokens)
}

func TestKeywords(t *testing.T) {
	input := `if else for while return fn let def`
	expectedTokens := []token.Token{
		{Type: token.IF, Value: "if"},
		{Type: token.ELSE, Value: "else"},
		{Type: token.FOR, Value: "for"},
		{Type: token.WHILE, Value: "while"},
		{Type: token.RETURN, Value: "return"},
		{Type: token.FN, Value: "fn"},
		{Type: token.LET, Value: "let"},
		{Type: token.DEF, Value: "def"},
		{Type: token.EOF, Value: ""},
	}
	testTokens(t, input, expectedTokens)
}

func testTokens(t *testing.T, input string, expected []token.Token) {
	t.Helper()

	l := New(input)

	for _, expectedToken := range expected {
		actualToken := l.Next()

		if assert.NotNil(t, actualToken) {
			assert.Equal(t, expectedToken.Type.String(), actualToken.Type.String(), "Token type doesn't match")
			assert.Equal(t, expectedToken.Value, actualToken.Value, "Token value doesn't match")
		}
	}
}
