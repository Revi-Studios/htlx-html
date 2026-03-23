package lexer

import (
	"fmt"
	lexertoken "htlx/cmd/lexer/lexertoken"
	"unicode/utf8"
)

type Lexer struct {
	Name   string
	Input  string
	Tokens chan lexertoken.Token

	Start int
	Pos   int
	Width int
}

// Puts a token onto the token channel. The value of this token is read from the input based on the current lexer position.
func (this *Lexer) Emit(tokenType lexertoken.TokenType) {
	this.Tokens <- lexertoken.Token{Type: tokenType, Value: this.Input[this.Start:this.Pos]}
	this.Start = this.Pos
}

// Increments the position
func (this *Lexer) Increment() {
	this.Pos++
	if this.Pos >= utf8.RuneCountInString(this.Input) {
		this.Emit(lexertoken.TOKEN_EOF)
	}
}

// Decrements the position
func (this *Lexer) Decrement() {
	this.Pos--
	if this.Pos >= utf8.RuneCountInString(this.Input) {
		this.Emit(lexertoken.TOKEN_EOF)
	}
}

// Return a slice of the input from the current lexer position to the end of the input string.
func (this *Lexer) InputToEnd() string {
	return this.Input[this.Pos:]
}

// Skips whitespace until we get something meaningful.
func (this *Lexer) SkipWhitespace() {
	for {
		ch := this.Next()

		if ch == lexertoken.EOF {
			this.Emit(lexertoken.TOKEN_EOF)
			break
		}

		if ch != ' ' {
			this.Decrement()
			this.Start = this.Pos
			break
		}
	}
}

func (this *Lexer) SkipToCh(ch rune) {
	for {
		rune := this.Next()

		if rune == ch {
			this.Decrement()
			break
		}

		if rune == lexertoken.EOF {
			this.Emit(lexertoken.TOKEN_EOF)
			break
		}
	}
}

func (this *Lexer) Next() rune {
	if this.Pos >= utf8.RuneCountInString(this.Input) {
		this.Width = 0
		return lexertoken.EOF
	}

	result, width := utf8.DecodeRuneInString(this.Input[this.Pos:])

	this.Width = width
	this.Pos += this.Width
	return result
}

// Returns a token with error information.
func (this *Lexer) Errorf(format string, args ...any) LexFn {
	this.Tokens <- lexertoken.Token{
		Type:  lexertoken.TOKEN_ERROR,
		Value: fmt.Sprintf(format, args...),
	}

	return nil
}

func (this *Lexer) CurrentCh() rune {
	result, width := utf8.DecodeRuneInString(this.Input[this.Pos : this.Pos+1])

	this.Width = width
	this.Pos += this.Width
	return result
}

// func (this *Lexer) NextToken() lexertoken.Token {
// 	return <-this.Tokens
// }
