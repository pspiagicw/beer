// Code generated by "stringer -type=TokenType"; DO NOT EDIT.

package token

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[EOF-0]
	_ = x[LPAREN-1]
	_ = x[RPAREN-2]
	_ = x[LBRACE-3]
	_ = x[RBRACE-4]
	_ = x[LSQUARE-5]
	_ = x[RSQUARE-6]
	_ = x[SEMICOLON-7]
	_ = x[ASSIGN-8]
	_ = x[IDENT-9]
	_ = x[PLUS-10]
	_ = x[STAR-11]
	_ = x[MINUS-12]
	_ = x[SLASH-13]
	_ = x[PERIOD-14]
	_ = x[FLOAT-15]
	_ = x[INT-16]
	_ = x[BOOL-17]
	_ = x[STRING-18]
	_ = x[TINT-19]
	_ = x[TSTRING-20]
	_ = x[TBOOL-21]
	_ = x[TFLOAT-22]
	_ = x[GT-23]
	_ = x[GTE-24]
	_ = x[LT-25]
	_ = x[LTE-26]
	_ = x[NEQ-27]
	_ = x[OR-28]
	_ = x[AND-29]
	_ = x[BITAND-30]
	_ = x[BITOR-31]
	_ = x[ILLEGAL-32]
}

const _TokenType_name = "EOFLPARENRPARENLBRACERBRACELSQUARERSQUARESEMICOLONASSIGNIDENTPLUSSTARMINUSSLASHPERIODFLOATINTBOOLSTRINGTINTTSTRINGTBOOLTFLOATGTGTELTLTENEQORANDBITANDBITORILLEGAL"

var _TokenType_index = [...]uint8{0, 3, 9, 15, 21, 27, 34, 41, 50, 56, 61, 65, 69, 74, 79, 85, 90, 93, 97, 103, 107, 114, 119, 125, 127, 130, 132, 135, 138, 140, 143, 149, 154, 161}

func (i TokenType) String() string {
	if i < 0 || i >= TokenType(len(_TokenType_index)-1) {
		return "TokenType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _TokenType_name[_TokenType_index[i]:_TokenType_index[i+1]]
}