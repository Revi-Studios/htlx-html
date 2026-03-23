package lexertoken

type TokenType int

const (
	TOKEN_ERROR TokenType = iota
	TOKEN_EOF

	TOKEN_LEFT_DELIMITER
	TOKEN_EQUAL_SIGN
	TOKEN_TAB
	TOKEN_NEWLINE
	TOKEN_WHITESPACE

	TOKEN_ELEMENT_TYPE
	TOKEN_ELEMENT_TEXTVALUE
	TOKEN_ATTRIBUTE_NAME
	TOKEN_ATTRIBUTE_VALUE
)

const EOF rune = 0

const LEFT_DELIMITER string = ">"
const EQUAL_SIGN string = "="
const WHITESPACE string = " "
const TAB string = "\t"
const NEWLINE string = "\n"

func (this TokenType) String() string {
	switch this {
	case TOKEN_ERROR:
		return "TOKEN_ERROR"
	case TOKEN_EOF:
		return "TOKEN_EOF"
	case TOKEN_LEFT_DELIMITER:
		return "TOKEN_LEFT_DELIMITER"
	case TOKEN_EQUAL_SIGN:
		return "TOKEN_EQUAL_SIGN"
	case TOKEN_TAB:
		return "TOKEN_TAB"
	case TOKEN_NEWLINE:
		return "TOKEN_NEWLINE"
	case TOKEN_WHITESPACE:
		return "TOKEN_WHITESPACE"
	case TOKEN_ELEMENT_TYPE:
		return "TOKEN_ELEMENT_TYPE"
	case TOKEN_ELEMENT_TEXTVALUE:
		return "TOKEN_ELEMENT_TEXTVALUE"
	case TOKEN_ATTRIBUTE_NAME:
		return "TOKEN_ATTRIBUTE_NAME"
	case TOKEN_ATTRIBUTE_VALUE:
		return "TOKEN_ATTRIBUTE_VALUE"
	default:
		return "TOKEN_NOT_IDENTIFIED"
	}
}
