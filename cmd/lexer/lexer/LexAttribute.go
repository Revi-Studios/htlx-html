package lexer

import (
	errors "htlx/cmd/lexer/error"
	"htlx/cmd/lexer/lexertoken"
	"strings"
	"unicode"
)

func LexAttribute(lexer *Lexer) LexFn {
	lexer.SkipWhitespace()
	for {
		ch := lexer.Next()

		if ch == lexertoken.EOF {
			return lexer.Errorf(errors.LEXER_ERROR_UNEXPECTED_EOF)
		}

		if string(ch) == lexertoken.LEFT_DELIMITER {
			return LexLeftDelimiter(lexer)
		}

		if unicode.IsSpace(ch) {
			return lexer.Errorf(errors.LEXER_ERROR_UNEXPECTED_ATTRIBUTE_WITHOUT_VALUE)
		}

		if strings.HasPrefix(lexer.InputToEnd(), lexertoken.EQUAL_SIGN) {
			lexer.Emit(lexertoken.TOKEN_ATTRIBUTE_NAME)
			return LexEqualSign(lexer)
		}
	}
}
