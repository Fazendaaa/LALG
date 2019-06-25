package lexer

import (
	"testing"

	"../token"
)

// TestNextToken :
func TestNextToken(t *testing.T) {
	input := `var five: integer := 5;const ten: integer := 10;

procedure testing(x: real, y: integer);
5 + 10;
10 - 5;

if 5 <= 10 then
	begin
		5 / 5;
		10.5 * 10.5;
	end
else
	10.7;

10 == 10;
10 <> 9;

{just a comment}

var foo: integer := 10
foo := 5
`

	test := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.VAR, "var"},
		{token.IDENTIFIER, "five"},
		{token.COLON, ":"},
		{token.INTEGER_KEYWORD, "integer"},
		{token.ASSIGN, ":="},
		{token.INTEGER, "5"},
		{token.SEMICOLON, ";"},
		{token.CONST, "const"},
		{token.IDENTIFIER, "ten"},
		{token.COLON, ":"},
		{token.INTEGER_KEYWORD, "integer"},
		{token.ASSIGN, ":="},
		{token.INTEGER, "10"},
		{token.SEMICOLON, ";"},
		{token.PROCEDURE, "procedure"},
		{token.IDENTIFIER, "testing"},
		{token.LEFT_PARENTHESIS, "("},
		{token.IDENTIFIER, "x"},
		{token.COLON, ":"},
		{token.REAL_KEYWORD, "real"},
		{token.COMMA, ","},
		{token.IDENTIFIER, "y"},
		{token.COLON, ":"},
		{token.INTEGER_KEYWORD, "integer"},
		{token.RIGHT_PARENTHESIS, ")"},
		{token.SEMICOLON, ";"},
		{token.INTEGER, "5"},
		{token.PLUS, "+"},
		{token.INTEGER, "10"},
		{token.SEMICOLON, ";"},
		{token.INTEGER, "10"},
		{token.MINUS, "-"},
		{token.INTEGER, "5"},
		{token.SEMICOLON, ";"},
		{token.IF, "if"},
		{token.INTEGER, "5"},
		{token.LESS_THAN_EQUAL, "<="},
		{token.INTEGER, "10"},
		{token.THEN, "then"},
		{token.BEGIN, "begin"},
		{token.INTEGER, "5"},
		{token.SLASH, "/"},
		{token.INTEGER, "5"},
		{token.SEMICOLON, ";"},
		{token.REAL, "10.5"},
		{token.ASTERISK, "*"},
		{token.REAL, "10.5"},
		{token.SEMICOLON, ";"},
		{token.END, "end"},
		{token.ELSE, "else"},
		{token.REAL, "10.7"},
		{token.SEMICOLON, ";"},
		{token.INTEGER, "10"},
		{token.EQUAL, "=="},
		{token.INTEGER, "10"},
		{token.SEMICOLON, ";"},
		{token.INTEGER, "10"},
		{token.DIFFERENT, "<>"},
		{token.INTEGER, "9"},
		{token.SEMICOLON, ";"},
		{token.LEFT_BRACES, "{"},
		{token.IDENTIFIER, "just"},
		{token.IDENTIFIER, "a"},
		{token.IDENTIFIER, "comment"},
		{token.RIGHT_BRACES, "}"},
		{token.VAR, "var"},
		{token.IDENTIFIER, "foo"},
		{token.COLON, ":"},
		{token.INTEGER_KEYWORD, "integer"},
		{token.ASSIGN, ":="},
		{token.INTEGER, "10"},
		{token.IDENTIFIER, "foo"},
		{token.ASSIGN, ":="},
		{token.INTEGER, "5"},
		{token.EOF, ""},
	}

	l := InitializeLexer(input)

	for i, tt := range test {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - token type wrong\n\texpected=%q, got=%q", i, tt.expectedLiteral, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong\n\texpected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}
