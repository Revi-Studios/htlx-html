package lexer

import (
	lexertoken "htlx/cmd/lexer/lexertoken"
	"strings"
)

func LexElementTextValue(lexer *Lexer) LexFn {
	for {
		ch := lexer.Next()
		if string(ch) == lexertoken.NEWLINE {
			if strings.TrimSpace(lexer.Input[lexer.Start:lexer.Pos]) != "" {
				lexer.Pos -= len(lexertoken.NEWLINE)
				lexer.Emit(lexertoken.TOKEN_ELEMENT_TEXTVALUE)
			}
			return LexBegin(lexer)
		}
	}
}
