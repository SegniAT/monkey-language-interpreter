package lexer

import (
	"github.com/SegniAT/monkey-language-interpreter/token"
)

type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination

	line      uint // current line
	character uint // current column
}

func New(input string) *Lexer {
	l := &Lexer{
		input:     input,
		line:      1,
		character: 0, // we begin at 0 since we call readChar below, which advances our cursor and sets this to 1
	}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.ch == '\n' {
		l.line++
		l.character = 1
	} else {
		l.character++
	}

	if l.readPosition >= len(l.input) {
		l.ch = 0 // ASCII code for the "NUL" character
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	startPosition := token.Position{
		Line:      l.line,
		Character: l.character,
	}

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			l.readChar()
			tok = token.Token{
				Type:    token.EQ,
				Literal: "==",
				Range: token.Range{
					Start: startPosition,
					End:   token.Position{Line: l.line, Character: l.character},
				},
			}
		} else {
			tok = newToken(token.ASSIGN, l.ch, startPosition)
		}
	case '+':
		tok = newToken(token.PLUS, l.ch, startPosition)
	case '-':
		tok = newToken(token.MINUS, l.ch, startPosition)
	case '!':
		if l.peekChar() == '=' {
			l.readChar()
			tok = token.Token{
				Type:    token.NOT_EQ,
				Literal: "!=",
				Range: token.Range{
					Start: startPosition,
					End:   token.Position{Line: l.line, Character: l.character},
				},
			}
		} else {
			tok = newToken(token.BANG, l.ch, startPosition)
		}
	case '/':
		tok = newToken(token.SLASH, l.ch, startPosition)
	case '*':
		tok = newToken(token.ASTERISK, l.ch, startPosition)
	case '<':
		tok = newToken(token.LT, l.ch, startPosition)
	case '>':
		tok = newToken(token.GT, l.ch, startPosition)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch, startPosition)
	case ',':
		tok = newToken(token.COMMA, l.ch, startPosition)
	case '(':
		tok = newToken(token.LPAREN, l.ch, startPosition)
	case ')':
		tok = newToken(token.RPAREN, l.ch, startPosition)
	case '{':
		tok = newToken(token.LBRACE, l.ch, startPosition)
	case '}':
		tok = newToken(token.RBRACE, l.ch, startPosition)
	case '"':
		tok.Type = token.STRING
		tok.Literal = l.readString()
		tok.Range = token.Range{
			Start: startPosition,
			End:   token.Position{Line: l.line, Character: l.character}, // Includes quotes as part of the token length
		}
	case '[':
		tok = newToken(token.LBRACKET, l.ch, startPosition)
	case ']':
		tok = newToken(token.RBRACKET, l.ch, startPosition)
	case ':':
		tok = newToken(token.COLON, l.ch, startPosition)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
		tok.Range = token.Range{
			Start: startPosition,
			End:   startPosition,
		}
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			tok.Range = token.Range{
				Start: startPosition,
				End:   token.Position{Line: l.line, Character: l.character - 1}, // readIdentifier goes 1 byte past the identifier
			}
			return tok
		} else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			tok.Range = token.Range{
				Start: startPosition,
				End:   token.Position{Line: l.line, Character: l.character - 1}, // readNumber goes 1 byte past the identifier
			}
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch, startPosition)
		}
	}

	l.readChar()
	return tok
}

// newToken is a helper function to create single byte tokens.
func newToken(tokenType token.TokenType, ch byte, startEnd token.Position) token.Token {
	return token.Token{
		Type:    tokenType,
		Literal: string(ch),
		Range: token.Range{
			Start: startEnd,
			End:   startEnd,
		},
	}
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}

	return l.input[position:l.position]
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}

	return l.input[position:l.position]
}

func (l *Lexer) readString() string {
	position := l.position + 1

	for {
		l.readChar()

		if l.ch == '"' || l.ch == 0 {
			break
		}
	}

	return l.input[position:l.position]
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}
