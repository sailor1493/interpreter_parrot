package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifier + Literal
	IDENT = "IDENT"
	INT   = "INT"

	// Operator
	ASSIGN = "="
	PLUS   = "+"

	// Separator
	COMMA     = ","
	SEMICOLON = ";"

	// Parantheseses
	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Reserved
	FUNCTION = "FUNCTION"
	LET      = "LET"
)
