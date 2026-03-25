package lexer

import (
	lexertoken "htlx/cmd/lexer/lexertoken"
	"unicode"
)

func LexBegin(lexer *Lexer) LexFn {
	lexer.SkipWhitespace()
	for {
		ch := lexer.Next()

		if string(ch) == lexertoken.TAB {
			return LexTab(lexer)
		}

		if ch == lexertoken.EOF {
			lexer.Emit(lexertoken.TOKEN_EOF)
		}

		if !unicode.IsSpace(ch) {
			lexer.Decrement()
			return LexElementType(lexer)
		}
	}
}
