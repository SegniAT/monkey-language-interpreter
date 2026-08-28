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
		expectedRange   token.Range
	}{
		{token.LET, "let", token.Range{Start: token.Position{Line: 1, Character: 1}, End: token.Position{Line: 1, Character: 3}}},
		{token.IDENT, "five", token.Range{Start: token.Position{Line: 1, Character: 5}, End: token.Position{Line: 1, Character: 8}}},
		{token.ASSIGN, "=", token.Range{Start: token.Position{Line: 1, Character: 10}, End: token.Position{Line: 1, Character: 10}}},
		{token.INT, "5", token.Range{Start: token.Position{Line: 1, Character: 12}, End: token.Position{Line: 1, Character: 12}}},
		{token.SEMICOLON, ";", token.Range{Start: token.Position{Line: 1, Character: 13}, End: token.Position{Line: 1, Character: 13}}},

		{token.LET, "let", token.Range{Start: token.Position{Line: 2, Character: 2}, End: token.Position{Line: 2, Character: 4}}},
		{token.IDENT, "ten", token.Range{Start: token.Position{Line: 2, Character: 6}, End: token.Position{Line: 2, Character: 8}}},
		{token.ASSIGN, "=", token.Range{Start: token.Position{Line: 2, Character: 10}, End: token.Position{Line: 2, Character: 10}}},
		{token.INT, "10", token.Range{Start: token.Position{Line: 2, Character: 12}, End: token.Position{Line: 2, Character: 13}}},
		{token.SEMICOLON, ";", token.Range{Start: token.Position{Line: 2, Character: 14}, End: token.Position{Line: 2, Character: 14}}},

		{token.LET, "let", token.Range{Start: token.Position{Line: 4, Character: 2}, End: token.Position{Line: 4, Character: 4}}},
		{token.IDENT, "add", token.Range{Start: token.Position{Line: 4, Character: 6}, End: token.Position{Line: 4, Character: 8}}},
		{token.ASSIGN, "=", token.Range{Start: token.Position{Line: 4, Character: 10}, End: token.Position{Line: 4, Character: 10}}},
		{token.FUNCTION, "fn", token.Range{Start: token.Position{Line: 4, Character: 12}, End: token.Position{Line: 4, Character: 13}}},
		{token.LPAREN, "(", token.Range{Start: token.Position{Line: 4, Character: 14}, End: token.Position{Line: 4, Character: 14}}},
		{token.IDENT, "x", token.Range{Start: token.Position{Line: 4, Character: 15}, End: token.Position{Line: 4, Character: 15}}},
		{token.COMMA, ",", token.Range{Start: token.Position{Line: 4, Character: 16}, End: token.Position{Line: 4, Character: 16}}},
		{token.IDENT, "y", token.Range{Start: token.Position{Line: 4, Character: 18}, End: token.Position{Line: 4, Character: 18}}},
		{token.RPAREN, ")", token.Range{Start: token.Position{Line: 4, Character: 19}, End: token.Position{Line: 4, Character: 19}}},
		{token.LBRACE, "{", token.Range{Start: token.Position{Line: 4, Character: 21}, End: token.Position{Line: 4, Character: 21}}},

		{token.IDENT, "x", token.Range{Start: token.Position{Line: 5, Character: 3}, End: token.Position{Line: 5, Character: 3}}},
		{token.PLUS, "+", token.Range{Start: token.Position{Line: 5, Character: 5}, End: token.Position{Line: 5, Character: 5}}},
		{token.IDENT, "y", token.Range{Start: token.Position{Line: 5, Character: 7}, End: token.Position{Line: 5, Character: 7}}},
		{token.SEMICOLON, ";", token.Range{Start: token.Position{Line: 5, Character: 8}, End: token.Position{Line: 5, Character: 8}}},

		{token.RBRACE, "}", token.Range{Start: token.Position{Line: 6, Character: 2}, End: token.Position{Line: 6, Character: 2}}},
		{token.SEMICOLON, ";", token.Range{Start: token.Position{Line: 6, Character: 3}, End: token.Position{Line: 6, Character: 3}}},

		{token.LET, "let", token.Range{Start: token.Position{Line: 8, Character: 2}, End: token.Position{Line: 8, Character: 4}}},
		{token.IDENT, "result", token.Range{Start: token.Position{Line: 8, Character: 6}, End: token.Position{Line: 8, Character: 11}}},
		{token.ASSIGN, "=", token.Range{Start: token.Position{Line: 8, Character: 13}, End: token.Position{Line: 8, Character: 13}}},
		{token.IDENT, "add", token.Range{Start: token.Position{Line: 8, Character: 15}, End: token.Position{Line: 8, Character: 17}}},
		{token.LPAREN, "(", token.Range{Start: token.Position{Line: 8, Character: 18}, End: token.Position{Line: 8, Character: 18}}},
		{token.IDENT, "five", token.Range{Start: token.Position{Line: 8, Character: 19}, End: token.Position{Line: 8, Character: 22}}},
		{token.COMMA, ",", token.Range{Start: token.Position{Line: 8, Character: 23}, End: token.Position{Line: 8, Character: 23}}},
		{token.IDENT, "ten", token.Range{Start: token.Position{Line: 8, Character: 25}, End: token.Position{Line: 8, Character: 27}}},
		{token.RPAREN, ")", token.Range{Start: token.Position{Line: 8, Character: 28}, End: token.Position{Line: 8, Character: 28}}},
		{token.SEMICOLON, ";", token.Range{Start: token.Position{Line: 8, Character: 29}, End: token.Position{Line: 8, Character: 29}}},

		{token.BANG, "!", token.Range{Start: token.Position{Line: 9, Character: 2}, End: token.Position{Line: 9, Character: 2}}},
		{token.MINUS, "-", token.Range{Start: token.Position{Line: 9, Character: 3}, End: token.Position{Line: 9, Character: 3}}},
		{token.SLASH, "/", token.Range{Start: token.Position{Line: 9, Character: 4}, End: token.Position{Line: 9, Character: 4}}},
		{token.ASTERISK, "*", token.Range{Start: token.Position{Line: 9, Character: 5}, End: token.Position{Line: 9, Character: 5}}},
		{token.INT, "5", token.Range{Start: token.Position{Line: 9, Character: 6}, End: token.Position{Line: 9, Character: 6}}},
		{token.SEMICOLON, ";", token.Range{Start: token.Position{Line: 9, Character: 7}, End: token.Position{Line: 9, Character: 7}}},

		{token.INT, "5", token.Range{Start: token.Position{Line: 10, Character: 2}, End: token.Position{Line: 10, Character: 2}}},
		{token.LT, "<", token.Range{Start: token.Position{Line: 10, Character: 4}, End: token.Position{Line: 10, Character: 4}}},
		{token.INT, "10", token.Range{Start: token.Position{Line: 10, Character: 6}, End: token.Position{Line: 10, Character: 7}}},
		{token.GT, ">", token.Range{Start: token.Position{Line: 10, Character: 9}, End: token.Position{Line: 10, Character: 9}}},
		{token.INT, "5", token.Range{Start: token.Position{Line: 10, Character: 11}, End: token.Position{Line: 10, Character: 11}}},
		{token.SEMICOLON, ";", token.Range{Start: token.Position{Line: 10, Character: 12}, End: token.Position{Line: 10, Character: 12}}},

		{token.IF, "if", token.Range{Start: token.Position{Line: 12, Character: 2}, End: token.Position{Line: 12, Character: 3}}},
		{token.LPAREN, "(", token.Range{Start: token.Position{Line: 12, Character: 5}, End: token.Position{Line: 12, Character: 5}}},
		{token.INT, "5", token.Range{Start: token.Position{Line: 12, Character: 6}, End: token.Position{Line: 12, Character: 6}}},
		{token.LT, "<", token.Range{Start: token.Position{Line: 12, Character: 8}, End: token.Position{Line: 12, Character: 8}}},
		{token.INT, "10", token.Range{Start: token.Position{Line: 12, Character: 10}, End: token.Position{Line: 12, Character: 11}}},
		{token.RPAREN, ")", token.Range{Start: token.Position{Line: 12, Character: 12}, End: token.Position{Line: 12, Character: 12}}},
		{token.LBRACE, "{", token.Range{Start: token.Position{Line: 12, Character: 14}, End: token.Position{Line: 12, Character: 14}}},

		{token.RETURN, "return", token.Range{Start: token.Position{Line: 13, Character: 3}, End: token.Position{Line: 13, Character: 8}}},
		{token.TRUE, "true", token.Range{Start: token.Position{Line: 13, Character: 10}, End: token.Position{Line: 13, Character: 13}}},
		{token.SEMICOLON, ";", token.Range{Start: token.Position{Line: 13, Character: 14}, End: token.Position{Line: 13, Character: 14}}},

		{token.RBRACE, "}", token.Range{Start: token.Position{Line: 14, Character: 2}, End: token.Position{Line: 14, Character: 2}}},
		{token.ELSE, "else", token.Range{Start: token.Position{Line: 14, Character: 4}, End: token.Position{Line: 14, Character: 7}}},
		{token.LBRACE, "{", token.Range{Start: token.Position{Line: 14, Character: 9}, End: token.Position{Line: 14, Character: 9}}},

		{token.RETURN, "return", token.Range{Start: token.Position{Line: 15, Character: 3}, End: token.Position{Line: 15, Character: 8}}},
		{token.FALSE, "false", token.Range{Start: token.Position{Line: 15, Character: 10}, End: token.Position{Line: 15, Character: 14}}},
		{token.SEMICOLON, ";", token.Range{Start: token.Position{Line: 15, Character: 15}, End: token.Position{Line: 15, Character: 15}}},

		{token.RBRACE, "}", token.Range{Start: token.Position{Line: 16, Character: 2}, End: token.Position{Line: 16, Character: 2}}},

		{token.INT, "10", token.Range{Start: token.Position{Line: 18, Character: 2}, End: token.Position{Line: 18, Character: 3}}},
		{token.EQ, "==", token.Range{Start: token.Position{Line: 18, Character: 5}, End: token.Position{Line: 18, Character: 6}}},
		{token.INT, "10", token.Range{Start: token.Position{Line: 18, Character: 8}, End: token.Position{Line: 18, Character: 9}}},
		{token.SEMICOLON, ";", token.Range{Start: token.Position{Line: 18, Character: 10}, End: token.Position{Line: 18, Character: 10}}},

		{token.INT, "10", token.Range{Start: token.Position{Line: 19, Character: 2}, End: token.Position{Line: 19, Character: 3}}},
		{token.NOT_EQ, "!=", token.Range{Start: token.Position{Line: 19, Character: 5}, End: token.Position{Line: 19, Character: 6}}},
		{token.INT, "9", token.Range{Start: token.Position{Line: 19, Character: 8}, End: token.Position{Line: 19, Character: 8}}},
		{token.SEMICOLON, ";", token.Range{Start: token.Position{Line: 19, Character: 9}, End: token.Position{Line: 19, Character: 9}}},

		{token.STRING, "foobar", token.Range{Start: token.Position{Line: 20, Character: 2}, End: token.Position{Line: 20, Character: 9}}},
		{token.STRING, "foo bar", token.Range{Start: token.Position{Line: 21, Character: 2}, End: token.Position{Line: 21, Character: 10}}},

		{token.LBRACKET, "[", token.Range{Start: token.Position{Line: 22, Character: 2}, End: token.Position{Line: 22, Character: 2}}},
		{token.INT, "1", token.Range{Start: token.Position{Line: 22, Character: 3}, End: token.Position{Line: 22, Character: 3}}},
		{token.COMMA, ",", token.Range{Start: token.Position{Line: 22, Character: 4}, End: token.Position{Line: 22, Character: 4}}},
		{token.INT, "2", token.Range{Start: token.Position{Line: 22, Character: 6}, End: token.Position{Line: 22, Character: 6}}},
		{token.RBRACKET, "]", token.Range{Start: token.Position{Line: 22, Character: 7}, End: token.Position{Line: 22, Character: 7}}},
		{token.SEMICOLON, ";", token.Range{Start: token.Position{Line: 22, Character: 8}, End: token.Position{Line: 22, Character: 8}}},

		{token.LBRACE, "{", token.Range{Start: token.Position{Line: 23, Character: 2}, End: token.Position{Line: 23, Character: 2}}},
		{token.STRING, "foo", token.Range{Start: token.Position{Line: 23, Character: 3}, End: token.Position{Line: 23, Character: 7}}},
		{token.COLON, ":", token.Range{Start: token.Position{Line: 23, Character: 8}, End: token.Position{Line: 23, Character: 8}}},
		{token.STRING, "bar", token.Range{Start: token.Position{Line: 23, Character: 10}, End: token.Position{Line: 23, Character: 14}}},
		{token.RBRACE, "}", token.Range{Start: token.Position{Line: 23, Character: 15}, End: token.Position{Line: 23, Character: 15}}},

		{token.EOF, "", token.Range{Start: token.Position{Line: 24, Character: 2}, End: token.Position{Line: 24, Character: 2}}},
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

		if tok.Range != tt.expectedRange {
			t.Fatalf("tests[%d] - range wrong. expected=%+v, got=%+v", i, tt.expectedRange, tok.Range)
		}
	}
}
