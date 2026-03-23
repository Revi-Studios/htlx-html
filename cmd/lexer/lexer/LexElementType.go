package lexer

import (
	lexertoken "htlx/cmd/lexer/lexertoken"
	"strings"
)

func LexElementType(lexer *Lexer) LexFn {
	lexer.SkipWhitespace()
	for {
		lexer.Next()

		if strings.HasPrefix(lexer.InputToEnd(), lexertoken.WHITESPACE) {
			lexer.Emit(lexertoken.TOKEN_ELEMENT_TYPE)
			return LexAttribute(lexer)
		}

		if strings.HasPrefix(lexer.InputToEnd(), lexertoken.LEFT_DELIMITER) {
			lexer.Emit(lexertoken.TOKEN_ELEMENT_TYPE)
			return LexLeftDelimiter(lexer)
		}
	}
}
