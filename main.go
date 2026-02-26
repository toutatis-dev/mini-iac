package main

import (
	"fmt"
	"mini-iac/lexer"
)

func main() {
	input := "resource \"file\" \"main.go\" { content = \"package main\"; }"
	lex := lexer.New(input)
	for i := 0; i <= len(input); i++ {
		token := lex.NextToken()
		fmt.Printf("token: %v\n", token)
		if token.Type == "EOF" {
			break
		}
	}
}
