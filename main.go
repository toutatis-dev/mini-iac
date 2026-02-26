package main

import (
	"fmt"
	"mini-iac/internal/lexer"
	"mini-iac/internal/parser"
)

func main() {
	input := `
	resource "file" "main.go" {
		content = "package main";
	}

	resource "file" "readme.md" {
		content = "# My Project";
	}
	`
	lex := lexer.New(input)
	p := parser.New(lex)

	manifest := p.ParseManifest()

	fmt.Printf("Parsed %d resources:\n", len(manifest.Blocks))
	for i, block := range manifest.Blocks {
		// Since Blocks are the 'Block' interface, we cast to *ast.Resource to print
		fmt.Printf("[%d] %+v\n", i, block)
	}
}
