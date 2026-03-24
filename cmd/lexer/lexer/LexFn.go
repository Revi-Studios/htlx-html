package lexer

import "htlx/cmd/lexer/lexertoken"

type LexFn func(*Lexer) LexFn

func NewLexer(input string) *Lexer {
	l := &Lexer{
		Input:  input,
		Tokens: make(chan lexertoken.Token, 5),
	}

	return l
}
