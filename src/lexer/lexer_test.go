package lexer

import (
	token "meowterpretor/src/tokens"

	"testing"
)

func TestNextToken(t *testing.T) {
	input := `=+(){},;`
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	lexer := New(input)

	for i, tested_token := range tests {
		current_token := lexer.NextToken()

		if current_token.Type != tested_token.expectedType {
			t.Fatalf("tests[%d] - Token type incorrect. Expected=%q, got %q..",
				i, tested_token.expectedType, current_token.Type)
		}

		if current_token.Literal != tested_token.expectedLiteral {
			t.Fatalf("tests[%d] -  Literal incorrect. Expected=%q, got %q..",
				i, tested_token.expectedLiteral, current_token.Literal)
		}
	}
}
