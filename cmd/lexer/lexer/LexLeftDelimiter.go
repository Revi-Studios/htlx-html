package lexer

import lexertoken "htlx/cmd/lexer/lexertoken"

func LexLeftDelimiter(lexer *Lexer) LexFn {
	lexer.Pos += len(lexertoken.LEFT_DELIMITER)
	lexer.Emit(lexertoken.TOKEN_LEFT_DELIMITER)
	return LexElementTextValue(lexer)
}
