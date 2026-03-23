package lexer

import (
	lexertoken "htlx/cmd/lexer/lexertoken"
	"strings"
)

func LexAttributeValue(lexer *Lexer) LexFn {
	for {
		lexer.Next()

		if strings.HasPrefix(lexer.InputToEnd(), lexertoken.WHITESPACE) {
			lexer.Emit(lexertoken.TOKEN_ATTRIBUTE_VALUE)
			return LexAttribute(lexer)
		}

		if strings.HasPrefix(lexer.InputToEnd(), lexertoken.LEFT_DELIMITER) {
			lexer.Emit(lexertoken.TOKEN_ATTRIBUTE_VALUE)
			return LexLeftDelimiter(lexer)
		}
	}
}
