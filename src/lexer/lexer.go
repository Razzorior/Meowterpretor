package lexer

import token "meowterpretor/src/tokens"

type Lexer struct {
	input           string
	inputPosition   int
	readingPosition int
	currentChar     byte
}

func New(input string) *Lexer {
	lexer := &Lexer{input: input}
	lexer.readChar()
	return lexer
}

func (lexer *Lexer) NextToken() token.Token {
	var resulting_token token.Token

	switch lexer.currentChar {
	case '=':
		resulting_token = newToken(token.ASSIGN, lexer.currentChar)
	case ';':
		resulting_token = newToken(token.SEMICOLON, lexer.currentChar)
	case ',':
		resulting_token = newToken(token.COMMA, lexer.currentChar)
	case '+':
		resulting_token = newToken(token.PLUS, lexer.currentChar)
	case '(':
		resulting_token = newToken(token.LPAREN, lexer.currentChar)
	case ')':
		resulting_token = newToken(token.RPAREN, lexer.currentChar)
	case '{':
		resulting_token = newToken(token.LBRACE, lexer.currentChar)
	case '}':
		resulting_token = newToken(token.RBRACE, lexer.currentChar)
	case 0:
		resulting_token.Literal = ""
		resulting_token.Type = token.EOF
	}

	lexer.readChar()
	return resulting_token
}

// Only supports ASCII for now
// TODO: Modify to support Unicode
func (lexer *Lexer) readChar() {
	if lexer.readingPosition >= len(lexer.input) {
		lexer.currentChar = 0
	} else {
		lexer.currentChar = lexer.input[lexer.readingPosition]
	}
	lexer.inputPosition = lexer.readingPosition
	lexer.readingPosition++
}

func newToken(tokenType token.TokenType, char byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(char)}
}
