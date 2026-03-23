package main

import (
	"fmt"
	"htlx/cmd/lexer/lexer"
	"htlx/cmd/parser"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: <path>")
		return
	}

	file, err := os.ReadFile(os.Args[1])

	if err != nil {
		fmt.Println("Error:", fmt.Errorf("reading file: %w", err))
	}

	// l := lexer.BeginLexing(os.Args[1], string(file))

	// go func() {
	// 	for token := range l.Tokens {
	// 		fmt.Printf("%v \t\"%v\"\n", token.Type.String(), strings.TrimSpace(token.Value))
	// 		if token.Type == lexertoken.TOKEN_EOF {
	// 			close(l.Tokens)
	// 		}
	// 	}
	// }()

	// lexer.LexBegin(l)

	l := lexer.BeginLexing(os.Args[1], string(file))
	go lexer.LexBegin(l)

	r, err := parser.ParseTokens(l.Tokens)

	if err != nil {
		fmt.Print(err)
	}

	fmt.Print(r)
}
