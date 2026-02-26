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
				if res != nil {
					manifest.Blocks = append(manifest.Blocks, res)
				}

			default:

			}
		}
		p.nextToken()
	}

	return manifest
}

func (p *Parser) parseResource() ast.Block {
	//resource definition must follow pattern resource string string {

	res := &ast.Resource{
		Properties: make(map[string]string),
	}
	prov := p.expectPeek(token.STRING)
	if prov {
		res.Provider = p.curToken.Literal
	} else {
		return nil
	}

	name := p.expectPeek(token.STRING)
	if name {
		res.ResourceName = p.curToken.Literal
	} else {
		return nil
	}

	if !p.expectPeek(token.LBRACE) {
		return nil
	}

	for p.peekToken.Type != token.RBRACE && p.peekToken.Type != token.EOF {
		//properties are ident, assign, string, semicolon
		p.nextToken()
		key := p.curToken.Literal //save property type as key for map
		if !p.expectPeek(token.ASSIGN) {
			return nil
		}
		if !p.expectPeek(token.STRING) {
			return nil
		}
		res.Properties[key] = p.curToken.Literal //load property value into map
		if !p.expectPeek(token.SEMICOLON) {
			return nil
		}
	}
	if !p.expectPeek(token.RBRACE) {
		return nil
	}
	return res

}

func (p *Parser) expectPeek(t token.TokenType) bool {

	if t == p.peekToken.Type {
		p.nextToken()
		return true
	} else {
		return false
	}

}
