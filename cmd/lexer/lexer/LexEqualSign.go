package lexer

import lexertoken "htlx/cmd/lexer/lexertoken"

func LexEqualSign(lexer *Lexer) LexFn {
	lexer.Pos += len(lexertoken.EQUAL_SIGN)
	lexer.Emit(lexertoken.TOKEN_EQUAL_SIGN)
	return LexAttributeValue(lexer)
}
