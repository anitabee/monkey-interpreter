package lexer

import (
	"fmt"

	"github.com/anitabee/monkey-interpreter/cmd/token"
)

type Lexer struct {
	input string
	// both position and readPosition will be used to access characters in the input string by using them
	// as indices e.g. l.input[l.position] or l.input[l.readPosition]
	// The reason for this two "pointers" is that we need to be able to peek ahead in the input to see next character.
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0 // if 0 (ASCI for null) then we've reached the end of the input
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition // move the position to the readPosition
	l.readPosition += 1         // move the readPosition to the next character
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	switch l.ch {

	case '=':
		tok = newToken(token.ASSIGN, l.ch)

	case ';':
		tok = newToken(token.SEMICOLON, l.ch)

	case '(':
		tok = newToken(token.LPAREN, l.ch)

	case ')':
		tok = newToken(token.RPAREN, l.ch)

	case ',':
		tok = newToken(token.COMMA, l.ch)

	case '+':
		tok = newToken(token.PLUS, l.ch)

	case '{':
		tok = newToken(token.LBRACE, l.ch)

	case '}':
		tok = newToken(token.RBRACE, l.ch)

	case 0:
		tok.Literal = ""
		tok.Type = token.EOF

	}

	l.readChar()

	return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	fmt.Println("newToken: tokenType =", tokenType, "ch =", string(ch))
	return token.Token{Type: tokenType, Literal: string(ch)}
}
