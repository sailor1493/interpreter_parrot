package lexer

import (
	"monkey/token"
	"testing"
)

type TokenTestcase struct {
	expectedType    token.TokenType
	expectedLiteral string
}

func evaulateTestcases(lex *Lexer, tcs []TokenTestcase, t *testing.T) {
	for i, tt := range tcs {
		tok := lex.NextToken()
		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q", i, tt.expectedType, tok.Type)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}

func TestSingleCharacterTokens(t *testing.T) {
	input := `=+(){},;-!/*<>`

	tests := []TokenTestcase{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.MINUS, "-"},
		{token.BANG, "!"},
		{token.SLASH, "/"},
		{token.ASTERISK, "*"},
		{token.LT, "<"},
		{token.GT, ">"},
		{token.EOF, ""},
	}

	l := New(input)
	evaulateTestcases(l, tests, t)
}

func TestNextTokenRealworldCase1(t *testing.T) {
	input := `let five = 5;
	let ten = 10;
	let add = fn(x, y) {
		x + y;
	};
	let result = add(five, ten);
	`

	tests := []TokenTestcase{
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
	evaulateTestcases(l, tests, t)
}

func TestKeywords(t *testing.T) {
	input := `
	if (5 < 10) {
		return true;
	} else {
		return false;
	}
	True False`

	tests := []TokenTestcase{
		{token.IF, "if"},
		{token.LPAREN, "("},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.TRUE, "true"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.ELSE, "else"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.FALSE, "false"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.TRUE, "True"},
		{token.FALSE, "False"},
		{token.EOF, ""},
	}

	lex := New(input)
	evaulateTestcases(lex, tests, t)
}

func TestDoubleCharacterTokens(t *testing.T) {
	input := `== != <= >= += -= *= /=`
	tests := []TokenTestcase{
		{token.EQ, "=="},
		{token.NEQ, "!="},
		{token.LE, "<="},
		{token.GE, ">="},
		{token.ADDASSIGN, "+="},
		{token.MINUSASSIGN, "-="},
		{token.MULASSIGN, "*="},
		{token.DIVASSIGN, "/="},
	}

	lex := New(input)
	evaulateTestcases(lex, tests, t)
}
