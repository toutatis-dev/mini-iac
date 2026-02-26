package parser

import (
	"mini-iac/internal/ast"
	"mini-iac/internal/lexer"
	"mini-iac/internal/token"
)

type Parser struct {
	l         *lexer.Lexer
	curToken  token.Token
	peekToken token.Token
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l: l,
	}

	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) ParseManifest() ast.Manifest {
	manifest := ast.Manifest{
		Blocks: []ast.Block{},
	}
	for {
		if p.curToken.Type == token.EOF {
			break
		}

		if p.curToken.Type == token.KEYWORD {
			switch p.curToken.Literal {
			case "resource":
				res := p.parseResource()
				manifest.Blocks = append(manifest.Blocks, res)
			}
		}
		p.nextToken()
	}

	return manifest
}
