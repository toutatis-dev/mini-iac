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

var Keywords = map[string]TokenType{
	"resource": KEYWORD,
}

type Token struct {
	Type    TokenType
	Literal string
}

func LookupIdent(ident string) TokenType {
	if tok, ok := Keywords[ident]; ok {
		return tok
	} else {
		return IDENT
	}
}
