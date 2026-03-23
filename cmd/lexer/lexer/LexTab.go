package lexer

import lexertoken "htlx/cmd/lexer/lexertoken"

func LexTab(lexer *Lexer) LexFn {
	lexer.Emit(lexertoken.TOKEN_TAB)
	return LexBegin(lexer)
}
