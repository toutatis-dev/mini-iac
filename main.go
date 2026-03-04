package main

import (
	"fmt"
	"log"
	"mini-iac/internal/lexer"
	"mini-iac/internal/parser"
	"mini-iac/internal/planner"
	"mini-iac/internal/providers/file"
	"mini-iac/internal/state"
	"os"
)

func main() {

	filePath := os.Args[1:]

	bcontents, err := os.ReadFile(filePath[0])
	if err != nil {
		log.Fatal(err)
	}
	input := string(bcontents)
	lex := lexer.New(input)
	p := parser.New(lex)

	manifest := p.ParseManifest()
	s := state.NewState()
	s.LoadState()

	plan, err := planner.Planner(&s, &manifest, &file.FileProvider{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Parsed %d resources:\n", len(manifest.Blocks))
	for i, block := range manifest.Blocks {
		// Since Blocks are the 'Block' interface, we cast to *ast.Resource to print
		fmt.Printf("[%d] %+v\n", i, block)
	}
}
