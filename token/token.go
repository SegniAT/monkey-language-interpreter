package token

type TokenType string

const (
	ILLEGAL TokenType = "ILLEGAL"
	EOF     TokenType = "EOF"

	// Identifiers + literals
	IDENT  TokenType = "IDENT"
	INT    TokenType = "INT"
	STRING TokenType = "STRING"

	// Operators
	ASSIGN   TokenType = "="
	PLUS     TokenType = "+"
	MINUS    TokenType = "-"
	BANG     TokenType = "!"
	ASTERISK TokenType = "*"
	SLASH    TokenType = "/"
	EQ       TokenType = "=="
	NOT_EQ   TokenType = "!="

	LT TokenType = "<"
	GT TokenType = ">"

	// Delimiters
	COMMA     TokenType = ","
	SEMICOLON TokenType = ";"

	LPAREN   TokenType = "("
	RPAREN   TokenType = ")"
	LBRACE   TokenType = "{"
	RBRACE   TokenType = "}"
	LBRACKET TokenType = "["
	RBRACKET TokenType = "]"
	COLON    TokenType = ":"

	// Keywords
	FUNCTION TokenType = "FUNCTION"
	LET      TokenType = "LET"
	TRUE     TokenType = "TRUE"
	FALSE    TokenType = "FALSE"
	IF       TokenType = "IF"
	ELSE     TokenType = "ELSE"
	RETURN   TokenType = "RETURN"
)

// Position describes 1-indexed grid location of a byte in the input.
type Position struct {
	Line      uint
	Character uint
}

type Range struct {
	Start Position
	End   Position
}

type Token struct {
	Type    TokenType
	Literal string
	Range   Range
}

type DiagnosticSeverity int

const (
	Error       DiagnosticSeverity = 1
	Warning     DiagnosticSeverity = 2
	Information DiagnosticSeverity = 3
	Hint        DiagnosticSeverity = 4
)

func (s DiagnosticSeverity) String() string {
	switch s {
	case Error:
		return "Error"
	case Warning:
		return "Warning"
	case Information:
		return "Information"
	case Hint:
		return "Hint"
	default:
		return "Unknown"
	}
}

type Diagnostic struct {
	Range    Range
	Message  string
	Severity DiagnosticSeverity
}

var Keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := Keywords[ident]; ok {
		return tok
	}
	return IDENT
}
