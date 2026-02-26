package token

type TokenType string

const (
	KEYWORD   TokenType = "Keyword"
	STRING    TokenType = "String"
	LBRACE    TokenType = "{"
	RBRACE    TokenType = "}"
	IDENT     TokenType = "Identifier"
	ASSIGN    TokenType = "="
	SEMICOLON TokenType = ";"
	EOF       TokenType = "EOF"
	ILLEGAL   TokenType = "Illegal"
)

type Token struct {
	Type    TokenType
	Literal string
}
