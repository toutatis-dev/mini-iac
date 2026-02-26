package lexer

import (
	"mini-iac/token"
)

type Lexer struct {
	Input       string
	CurrentChar int
	NextChar    int
	Char        byte
}

func New(input string) *Lexer {
	if len(input) == 0 {
		return nil
	}

	l := &Lexer{Input: input}
	l.readChar()

	return l
}

func (l *Lexer) readChar() {
	if l.NextChar >= len(l.Input) {
		l.Char = 0 //Nul
	} else {
		l.Char = l.Input[l.NextChar]
	}

	l.CurrentChar = l.NextChar
	l.NextChar += 1

}

func (l *Lexer) NextToken() token.Token {
	l.skipWhiteSpace()
	var tok token.Token

	switch l.Char {
	case '{':
		tok.Type = token.LBRACE
		tok.Literal = "{"
	case '}':
		tok.Type = token.RBRACE
		tok.Literal = "}"
	case ';':
		tok.Type = token.SEMICOLON
		tok.Literal = ";"
	case '0':
		tok.Type = token.EOF
		tok.Literal = "EOF"
	}

	l.readChar()
	return tok
}

func (l *Lexer) skipWhiteSpace() {
	for l.Char == ' ' || l.Char == '\t' || l.Char == '\n' || l.Char == '\r' {
		l.readChar()
	}
}
