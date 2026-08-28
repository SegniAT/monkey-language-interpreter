package lexer

import (
	"github.com/SegniAT/monkey-language-interpreter/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `let five = 5;
	let ten = 10;
	
	let add = fn(x, y) {
		x + y;
	};
	
	let result = add(five, ten);
	!-/*5;
	5 < 10 > 5;

	if (5 < 10) {
		return true;
	} else {
		return false;
	}

	10 == 10;
	10 != 9;
	"foobar"
	"foo bar"
	[1, 2];
	{"foo": "bar"}
	`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
		expectedStart   token.Position
		expectedEnd     token.Position
	}{
		{token.LET, "let", token.Position{Line: 1, Character: 0}, token.Position{Line: 1, Character: 2}},
		{token.IDENT, "five", token.Position{Line: 1, Character: 4}, token.Position{Line: 1, Character: 7}},
		{token.ASSIGN, "=", token.Position{Line: 1, Character: 9}, token.Position{Line: 1, Character: 9}},
		{token.INT, "5", token.Position{Line: 1, Character: 11}, token.Position{Line: 1, Character: 11}},
		{token.SEMICOLON, ";", token.Position{Line: 1, Character: 12}, token.Position{Line: 1, Character: 12}},

		{token.LET, "let", token.Position{Line: 2, Character: 0}, token.Position{Line: 2, Character: 2}},
		{token.IDENT, "ten", token.Position{Line: 2, Character: 4}, token.Position{Line: 2, Character: 6}},
		{token.ASSIGN, "=", token.Position{Line: 2, Character: 8}, token.Position{Line: 2, Character: 8}},
		{token.INT, "10", token.Position{Line: 2, Character: 10}, token.Position{Line: 2, Character: 11}},
		{token.SEMICOLON, ";", token.Position{Line: 2, Character: 12}, token.Position{Line: 2, Character: 12}},

		{token.LET, "let", token.Position{Line: 4, Character: 0}, token.Position{Line: 4, Character: 2}},
		{token.IDENT, "add", token.Position{Line: 4, Character: 4}, token.Position{Line: 4, Character: 6}},
		{token.ASSIGN, "=", token.Position{Line: 4, Character: 8}, token.Position{Line: 4, Character: 8}},
		{token.FUNCTION, "fn", token.Position{Line: 4, Character: 10}, token.Position{Line: 4, Character: 11}},
		{token.LPAREN, "(", token.Position{Line: 4, Character: 12}, token.Position{Line: 4, Character: 12}},
		{token.IDENT, "x", token.Position{Line: 4, Character: 13}, token.Position{Line: 4, Character: 13}},
		{token.COMMA, ",", token.Position{Line: 4, Character: 14}, token.Position{Line: 4, Character: 14}},
		{token.IDENT, "y", token.Position{Line: 4, Character: 16}, token.Position{Line: 4, Character: 16}},
		{token.RPAREN, ")", token.Position{Line: 4, Character: 17}, token.Position{Line: 4, Character: 17}},
		{token.LBRACE, "{", token.Position{Line: 4, Character: 19}, token.Position{Line: 4, Character: 19}},

		{token.IDENT, "x", token.Position{Line: 5, Character: 1}, token.Position{Line: 5, Character: 1}},
		{token.PLUS, "+", token.Position{Line: 5, Character: 3}, token.Position{Line: 5, Character: 3}},
		{token.IDENT, "y", token.Position{Line: 5, Character: 5}, token.Position{Line: 5, Character: 5}},
		{token.SEMICOLON, ";", token.Position{Line: 5, Character: 6}, token.Position{Line: 5, Character: 6}},

		{token.RBRACE, "}", token.Position{Line: 6, Character: 0}, token.Position{Line: 6, Character: 0}},
		{token.SEMICOLON, ";", token.Position{Line: 6, Character: 1}, token.Position{Line: 6, Character: 1}},

		{token.LET, "let", token.Position{Line: 8, Character: 0}, token.Position{Line: 8, Character: 2}},
		{token.IDENT, "result", token.Position{Line: 8, Character: 4}, token.Position{Line: 8, Character: 9}},
		{token.ASSIGN, "=", token.Position{Line: 8, Character: 11}, token.Position{Line: 8, Character: 11}},
		{token.IDENT, "add", token.Position{Line: 8, Character: 13}, token.Position{Line: 8, Character: 15}},
		{token.LPAREN, "(", token.Position{Line: 8, Character: 16}, token.Position{Line: 8, Character: 16}},
		{token.IDENT, "five", token.Position{Line: 8, Character: 17}, token.Position{Line: 8, Character: 20}},
		{token.COMMA, ",", token.Position{Line: 8, Character: 21}, token.Position{Line: 8, Character: 21}},
		{token.IDENT, "ten", token.Position{Line: 8, Character: 23}, token.Position{Line: 8, Character: 25}},
		{token.RPAREN, ")", token.Position{Line: 8, Character: 26}, token.Position{Line: 8, Character: 26}},
		{token.SEMICOLON, ";", token.Position{Line: 8, Character: 27}, token.Position{Line: 8, Character: 27}},

		{token.BANG, "!", token.Position{Line: 9, Character: 0}, token.Position{Line: 9, Character: 0}},
		{token.MINUS, "-", token.Position{Line: 9, Character: 1}, token.Position{Line: 9, Character: 1}},
		{token.SLASH, "/", token.Position{Line: 9, Character: 2}, token.Position{Line: 9, Character: 2}},
		{token.ASTERISK, "*", token.Position{Line: 9, Character: 3}, token.Position{Line: 9, Character: 3}},
		{token.INT, "5", token.Position{Line: 9, Character: 4}, token.Position{Line: 9, Character: 4}},
		{token.SEMICOLON, ";", token.Position{Line: 9, Character: 5}, token.Position{Line: 9, Character: 5}},

		{token.INT, "5", token.Position{Line: 10, Character: 0}, token.Position{Line: 10, Character: 0}},
		{token.LT, "<", token.Position{Line: 10, Character: 2}, token.Position{Line: 10, Character: 2}},
		{token.INT, "10", token.Position{Line: 10, Character: 4}, token.Position{Line: 10, Character: 5}},
		{token.GT, ">", token.Position{Line: 10, Character: 7}, token.Position{Line: 10, Character: 7}},
		{token.INT, "5", token.Position{Line: 10, Character: 9}, token.Position{Line: 10, Character: 9}},
		{token.SEMICOLON, ";", token.Position{Line: 10, Character: 10}, token.Position{Line: 10, Character: 10}},

		{token.IF, "if", token.Position{Line: 12, Character: 0}, token.Position{Line: 12, Character: 1}},
		{token.LPAREN, "(", token.Position{Line: 12, Character: 3}, token.Position{Line: 12, Character: 3}},
		{token.INT, "5", token.Position{Line: 12, Character: 4}, token.Position{Line: 12, Character: 4}},
		{token.LT, "<", token.Position{Line: 12, Character: 6}, token.Position{Line: 12, Character: 6}},
		{token.INT, "10", token.Position{Line: 12, Character: 8}, token.Position{Line: 12, Character: 9}},
		{token.RPAREN, ")", token.Position{Line: 12, Character: 10}, token.Position{Line: 12, Character: 10}},
		{token.LBRACE, "{", token.Position{Line: 12, Character: 12}, token.Position{Line: 12, Character: 12}},

		{token.RETURN, "return", token.Position{Line: 13, Character: 1}, token.Position{Line: 13, Character: 6}},
		{token.TRUE, "true", token.Position{Line: 13, Character: 8}, token.Position{Line: 13, Character: 11}},
		{token.SEMICOLON, ";", token.Position{Line: 13, Character: 12}, token.Position{Line: 13, Character: 12}},

		{token.RBRACE, "}", token.Position{Line: 14, Character: 0}, token.Position{Line: 14, Character: 0}},
		{token.ELSE, "else", token.Position{Line: 14, Character: 2}, token.Position{Line: 14, Character: 5}},
		{token.LBRACE, "{", token.Position{Line: 14, Character: 7}, token.Position{Line: 14, Character: 7}},

		{token.RETURN, "return", token.Position{Line: 15, Character: 1}, token.Position{Line: 15, Character: 6}},
		{token.FALSE, "false", token.Position{Line: 15, Character: 8}, token.Position{Line: 15, Character: 12}},
		{token.SEMICOLON, ";", token.Position{Line: 15, Character: 13}, token.Position{Line: 15, Character: 13}},

		{token.RBRACE, "}", token.Position{Line: 16, Character: 0}, token.Position{Line: 16, Character: 0}},

		{token.INT, "10", token.Position{Line: 18, Character: 0}, token.Position{Line: 18, Character: 1}},
		{token.EQ, "==", token.Position{Line: 18, Character: 3}, token.Position{Line: 18, Character: 4}},
		{token.INT, "10", token.Position{Line: 18, Character: 6}, token.Position{Line: 18, Character: 7}},
		{token.SEMICOLON, ";", token.Position{Line: 18, Character: 8}, token.Position{Line: 18, Character: 8}},

		{token.INT, "10", token.Position{Line: 19, Character: 0}, token.Position{Line: 19, Character: 1}},
		{token.NOT_EQ, "!=", token.Position{Line: 19, Character: 3}, token.Position{Line: 19, Character: 4}},
		{token.INT, "9", token.Position{Line: 19, Character: 6}, token.Position{Line: 19, Character: 6}},
		{token.SEMICOLON, ";", token.Position{Line: 19, Character: 7}, token.Position{Line: 19, Character: 7}},

		{token.STRING, "foobar", token.Position{Line: 20, Character: 0}, token.Position{Line: 20, Character: 7}},
		{token.STRING, "foo bar", token.Position{Line: 21, Character: 0}, token.Position{Line: 21, Character: 8}},

		{token.LBRACKET, "[", token.Position{Line: 22, Character: 0}, token.Position{Line: 22, Character: 0}},
		{token.INT, "1", token.Position{Line: 22, Character: 1}, token.Position{Line: 22, Character: 1}},
		{token.COMMA, ",", token.Position{Line: 22, Character: 2}, token.Position{Line: 22, Character: 2}},
		{token.INT, "2", token.Position{Line: 22, Character: 4}, token.Position{Line: 22, Character: 4}},
		{token.RBRACKET, "]", token.Position{Line: 22, Character: 5}, token.Position{Line: 22, Character: 5}},
		{token.SEMICOLON, ";", token.Position{Line: 22, Character: 6}, token.Position{Line: 22, Character: 6}},

		{token.LBRACE, "{", token.Position{Line: 23, Character: 0}, token.Position{Line: 23, Character: 0}},
		{token.STRING, "foo", token.Position{Line: 23, Character: 1}, token.Position{Line: 23, Character: 5}},
		{token.COLON, ":", token.Position{Line: 23, Character: 6}, token.Position{Line: 23, Character: 6}},
		{token.STRING, "bar", token.Position{Line: 23, Character: 8}, token.Position{Line: 23, Character: 12}},
		{token.RBRACE, "}", token.Position{Line: 23, Character: 13}, token.Position{Line: 23, Character: 13}},

		{token.EOF, "", token.Position{Line: 24, Character: 0}, token.Position{Line: 24, Character: 0}},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q", i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}
