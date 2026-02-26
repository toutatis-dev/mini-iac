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
	case '=':
		tok.Type = token.ASSIGN
		tok.Literal = "="
	case '"':
		tok.Type = token.STRING
		tok.Literal = l.readString()
	case 0:
		tok.Type = token.EOF
		tok.Literal = "EOF"
	default:
		if isLetter(l.Char) {
			ident := l.readIdentifier()

			tok.Type = token.LookupIdent(ident)
			tok.Literal = ident
			return tok
		} else {
			tok.Type = token.ILLEGAL
			tok.Literal = string(l.Char)
		}

	}

	l.readChar()
	return tok
}

func (l *Lexer) skipWhiteSpace() {
	for l.Char == ' ' || l.Char == '\t' || l.Char == '\n' || l.Char == '\r' {
		l.readChar()
	}
}

func (l *Lexer) readIdentifier() string {

	initPos := l.CurrentChar

	for isLetter(l.Char) {
		l.readChar()
	}

	return l.Input[initPos:l.CurrentChar]

}

func isLetter(char byte) bool {
	return 'a' <= char && char <= 'z' || 'A' <= char && char <= 'Z' || char == '_'
}

func (l *Lexer) readString() string {
	firstLetter := l.NextChar

	for {
		l.readChar()
		if l.Char == '"' || l.Char == 0 {
			break
		}

	}
	return l.Input[firstLetter:l.CurrentChar]

}
