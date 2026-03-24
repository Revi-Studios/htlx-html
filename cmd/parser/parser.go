package parser

import (
	"errors"
	"htlx/cmd/lexer/lexertoken"
	t "htlx/cmd/lexer/lexertoken"
	"htlx/cmd/parser/model/htlx"
)

// // Parses HTLX into normal HTML
// func Parse(str string) (string, error) {
// 	l := lexer.BeginLexing("", str)
// 	outFile, err := parseTokens(l.Tokens)
// 	return outFile.Value, err
// }

func ParseTokens(tokens chan lexertoken.Token) (htlx.HtlxElement, error) {
	defer close(tokens)

	result := htlx.HtlxElement{Type: "root"}
	var token lexertoken.Token
	token = <-tokens

	for {
		var elem htlx.HtlxElement
		var tabs int

		if token.Type == t.TOKEN_TAB {
			tabs++
			for {
				token = <-tokens

				if token.Type == t.TOKEN_TAB {
					tabs++
				} else {
					break
				}
			}
		}

		if token.Type != t.TOKEN_ELEMENT_TYPE {
			return htlx.HtlxElement{}, errors.New("Invalid starting token")
		}
		elem.Type = token.Value
		token = <-tokens

		// Handle Attributes
		if token.Type == t.TOKEN_ATTRIBUTE_NAME {
			var attributes [][2]t.Token
			for {
				if token.Type != t.TOKEN_ATTRIBUTE_NAME {
					break
				}

				// Create a new attribute (fixed-size array of 3 tokens)
				var attribute [2]t.Token

				// Process attribute name
				attribute[0] = token
				token = <-tokens
				if token.Type != t.TOKEN_EQUAL_SIGN {
					return htlx.HtlxElement{}, errors.New("Expected '=' after attribute name")
				}

				// Process equal sign
				token = <-tokens
				if token.Type != t.TOKEN_ATTRIBUTE_VALUE {
					return htlx.HtlxElement{}, errors.New("Expected attribute value after '='")
				}

				// Process attribute value
				attribute[1] = token
				token = <-tokens

				attributes = append(attributes, attribute)
			}

			for _, attribute := range attributes {
				elem.Attributes = append(elem.Attributes, htlx.HtlxAttribute{Key: attribute[0].Value, Value: attribute[1].Value})
			}
		}

		// Hanlde left delimiter
		if token.Type == t.TOKEN_LEFT_DELIMITER {
			token = <-tokens
		} else {
			return htlx.HtlxElement{}, errors.New("Invalid token type")
		}

		if token.Type == t.TOKEN_ELEMENT_TEXTVALUE {
			elem.Value = token.Value
			token = <-tokens
		}

		result.AppenddChildElement(tabs, &elem)

		if token.Type == t.TOKEN_EOF {
			break
		}
	}
	return result, nil
}
