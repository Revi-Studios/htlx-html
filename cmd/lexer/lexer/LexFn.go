package lexer

import "htlx/cmd/lexer/lexertoken"

type LexFn func(*Lexer) LexFn

func BeginLexing(name, input string) *Lexer {
	l := &Lexer{
		Name:   name,
		Input:  input,
		Tokens: make(chan lexertoken.Token, 5),
	}

	return l
}
