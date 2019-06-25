package token

// TokenType : this will work as a PoC only, needs to change it to an int or a byte later on
type TokenType string

// Token : stores the information token related; later on add a line and column to it to make an easier debug later on
type Token struct {
	Type    TokenType
	Literal string
}

const (
	EOF     = "EOF"
	ILLEGAL = "ILLEGAL"

	VAR   = "VAR"
	CONST = "CONST"

	REAL_KEYWORD    = "REAL_KEYWORD"
	INTEGER_KEYWORD = "INTEGER_KEYWORD"

	INTEGER    = "INTEGER"
	REAL       = "REAL"
	IDENTIFIER = "IDENTIFIER"

	FOR       = "FOR"
	PROCEDURE = "PROCEDURE"
	BEGIN     = "BEGIN"
	END       = "END"

	IF   = "IF"
	THEN = "THEN"
	ELSE = "ELSE"

	ASSIGN   = ":="
	PLUS     = "+"
	MINUS    = "-"
	SLASH    = "/"
	ASTERISK = "*"

	LESS_THAN          = "<"
	GREATER_THAN       = ">"
	LESS_THAN_EQUAL    = "<="
	GREATER_THAN_EQUAl = ">="
	EQUAL              = "=="
	DIFFERENT          = "<>"

	COMMA             = ","
	COLON             = ":"
	SEMICOLON         = ";"
	LEFT_PARENTHESIS  = "("
	RIGHT_PARENTHESIS = ")"
	RIGHT_BRACES      = "{"
	LEFT_BRACES       = "}"
)

var keywords = map[string]TokenType{
	"if":        IF,
	"var":       VAR,
	"for":       FOR,
	"end":       END,
	"else":      ELSE,
	"then":      THEN,
	"+":         PLUS,
	"const":     CONST,
	"begin":     BEGIN,
	"==":        EQUAL,
	":":         COLON,
	"-":         MINUS,
	"/":         SLASH,
	":=":        ASSIGN,
	"*":         ASTERISK,
	"procedure": PROCEDURE,
	"<":         LESS_THAN,
	">":         GREATER_THAN,
	"real":      REAL_KEYWORD,
	"integer":   INTEGER_KEYWORD,
	"<=":        LESS_THAN_EQUAL,
	">=":        GREATER_THAN_EQUAl,
}

// LookupIdentifier :
func LookupIdentifier(identification string) TokenType {
	if tok, ok := keywords[identification]; ok {
		return tok
	}

	return IDENTIFIER
}
