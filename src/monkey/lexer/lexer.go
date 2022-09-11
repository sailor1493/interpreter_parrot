package lexer

type Lexer struct {
	input string 
	position input // Current cursor
	readPosition input // Next character to be searched
 	ch byte // Character currently evaluated
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} 
	else {
		l.ch = l.input(l.readPosition)
	}
	l.position = l.readPosition
	l.readPosition += 1
}