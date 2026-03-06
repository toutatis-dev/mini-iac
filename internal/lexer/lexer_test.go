package lexer

import (
	"mini-iac/internal/token"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {

	tests := []struct {
		name     string
		input    string
		expected *Lexer
	}{
		{"Valid input", "Hello, World!", &Lexer{Input: "Hello, World!", CurrentChar: 0, NextChar: 1, Char: 'H'}},
		{"No input", "", nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testLexer := New(tt.input)
			assert.Equal(t, tt.expected, testLexer)
		})
	}
}

func TestNextToken(t *testing.T) {

	tests := []struct {
		name     string
		input    string
		expected token.Token
	}{
		{"LBRACE", "{", token.Token{Type: token.LBRACE, Literal: "{"}},
		{"RBRACE", "}", token.Token{Type: token.RBRACE, Literal: "}"}},
		{"SEMICOLON", ";", token.Token{Type: token.SEMICOLON, Literal: ";"}},
		{"ASSIGN", "=", token.Token{Type: token.ASSIGN, Literal: "="}},
		{"STRING", "\"Hello, World!", token.Token{Type: token.STRING, Literal: "Hello, World!"}},
		{"ILLEGAL", "1", token.Token{Type: token.ILLEGAL, Literal: "1"}},
		{"KEYWORD", "resource", token.Token{Type: token.KEYWORD, Literal: "resource"}},
		{"IDENT", "file", token.Token{Type: token.IDENT, Literal: "file"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := New(tt.input)
			tok := l.NextToken()
			assert.Equal(t, tt.expected, tok)
		})
	}

}
