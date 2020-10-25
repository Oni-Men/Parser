package main

import (
	"fmt"

	"./ast"
	"./ast/lexer"
)

const test = "let a = 2 == (15 + 3);"

func main() {
	tokens := lexer.Lexer(test)
	for _, token := range tokens.TokenList {
		fmt.Println(token.Type, token.Literal)
	}

}

func parseIt(tokens *ast.Sequence) *ast.INode {

}
