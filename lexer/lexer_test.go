package lexer

import (
	"testing"

	"github.com/kaleabAlemayehu/monkey/token"
)

func TestNextToken(t *testing.T) {
	input := `let five = 5;let ten = 10;let add = fn(x, y) {x + y;};let result = add(five, ten);`
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}
	l := New(input)

	for index, testToken := range tests {
		token := l.NextToken()
		if token.Type != testToken.expectedType {
			t.Fatalf("tests[%d] - tokenType wrong. expected=%q, got=%q", index, testToken.expectedType, token.Type)
		}
		if token.Literal != testToken.expectedLiteral {
			t.Fatalf("tests[%d] - token Literal wrong, expected=%q, got=%q", index, testToken.expectedLiteral, token.Literal)
		}
	}
}