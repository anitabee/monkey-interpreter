package lexer

import (
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

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {

	case '=':
		tok = newToken(token.ASSIGN, l.ch)
		// todo: abstarct behaviour in a method called readTwoCharToken
		if l.peekChar() == '=' {
			ch := l.ch // we save the current character in a local variable before reading the next character, 
			// in this way we don't lose the current character and can safely advance the lexer so it leaves
			// the NextToken() with l.position and l.readPosition in correct state.
			l.readChar()
			tok = token.Token{Type: token.EQ, Literal: string(ch) + string(l.ch)}
		} else {
			tok = newToken(token.ASSIGN, l.ch)
		}

	case '+':
		tok = newToken(token.PLUS, l.ch)

	case '-':
		tok = newToken(token.MINUS, l.ch)

	case '!':
		if l.peekChar() == '=' {
			ch := l.ch 
			l.readChar()
			tok = token.Token{Type: token.NOT_EQ, Literal: string(ch) + string(l.ch)}
		} else {
			tok = newToken(token.BANG, l.ch)
		}
	case '/':
		tok = newToken(token.SLASH, l.ch)

	case '*':
		tok = newToken(token.ASTERISK, l.ch)

	case '<':
		tok = newToken(token.LT, l.ch)

	case '>':
		tok = newToken(token.GT, l.ch)

	case ';':
		tok = newToken(token.SEMICOLON, l.ch)

	case ',':
		tok = newToken(token.COMMA, l.ch)

	case '{':
		tok = newToken(token.LBRACE, l.ch)

	case '}':
		tok = newToken(token.RBRACE, l.ch)

	case '(':
		tok = newToken(token.LPAREN, l.ch)

	case ')':
		tok = newToken(token.RPAREN, l.ch)

	case 0:
		tok.Literal = ""
		tok.Type = token.EOF

	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			// the return statement of tok here is necessary beceause when calling readIdentifier()
			// we call readChar() which moves the position to the next character.
			// So we don't need to call readChar() after the switch statement again.
			return tok
		}
		if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok

		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}

	}

	l.readChar()

	return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	// fmt.Printf("newToken: %v, %v\n", tokenType, string(ch))
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func isLetter(ch byte) bool {
	// ASCII values for a-z, A-Z and _
	// a-z: 97-122
	// A-Z: 65-90
	// _: 95
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	// ASCII values for 0-9
	// 0: 48
	// 9: 57
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) readIdentifier() string {
	postion := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[postion:l.position]
}

func (l *Lexer) readNumber() string {
	position := l.position
	// we support only integers for now
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}
