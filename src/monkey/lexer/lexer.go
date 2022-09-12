package lexer

import "monkey/token"

type Lexer struct {
	input        string
	position     int  // Current cursor
	readPosition int  // Next character to be searched
	ch           byte // Character currently evaluated
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()
	switch l.ch {

	// Operator
	case '=':
		tok = l.peekCharAndMakeToken('=', token.EQ, token.ASSIGN)
	case '+':
		tok = l.peekCharAndMakeToken('=', token.ADDASSIGN, token.PLUS)
	case '-':
		tok = l.peekCharAndMakeToken('=', token.MINUSASSIGN, token.MINUS)
	case '!':
		tok = l.peekCharAndMakeToken('=', token.NEQ, token.BANG)
	case '/':
		tok = l.peekCharAndMakeToken('=', token.DIVASSIGN, token.SLASH)
	case '*':
		tok = l.peekCharAndMakeToken('=', token.MULASSIGN, token.ASTERISK)

		// Comparator
	case '<':
		tok = l.peekCharAndMakeToken('=', token.LE, token.LT)
	case '>':
		tok = l.peekCharAndMakeToken('=', token.GE, token.GT)

		// Separators
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)

		// Parantheses
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)

		// Error check
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tok
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

func (l *Lexer) peekCharAndMakeToken(crit byte, wt_type token.TokenType, wo_type token.TokenType) token.Token {
	if l.peekChar() == crit {
		literal := l.buildDoubleCharacterLiteral()
		return token.Token{Type: wt_type, Literal: literal}
	} else {
		return newToken(wo_type, l.ch)
	}
}

func (l *Lexer) buildDoubleCharacterLiteral() string {
	ch := l.ch
	l.readChar()
	return string(ch) + string(l.ch)
}
