package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	// For unknown tokens
	ILLEGAL = "ILLEGAL"
	// For parser to know when to stop
	EOF = "EOF"

	// Identifiers and literals
	IDENT = "IDENT"
	INT   = "INT"

	// Operators
	ASSIGN = "="
	PLUS   = "+"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"

	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	DECL     = "DECL"
)
