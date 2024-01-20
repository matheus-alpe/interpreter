package lexer

import (
	"testing"

	"github.com/matheus-alpe/interpreter/internal/token"
)

type TokenTestExpected struct {
	Type    token.TokenType
	Literal string
}

func TestSimpleNextToken(t *testing.T) {
	input := `=+(){},;!-/*><`

	tests := []TokenTestExpected{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPARENT, "("},
		{token.RPARENT, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.BANG, "!"},
		{token.MINUS, "-"},
		{token.SLASH, "/"},
		{token.ASTERISK, "*"},
		{token.GT, ">"},
		{token.LT, "<"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()
		checkToken(t, i, &tok, &tt)
	}
}

func TestComplexNextToken(t *testing.T) {
	input := `let five = 5;
	let ten = 10;

	let add = fn(x, y) {
		x + y;
	};

	let result = add(five, ten);
	`

	tests := []TokenTestExpected{
		{token.LET, "let"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LPARENT, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPARENT, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LPARENT, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.RPARENT, ")"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()
		checkToken(t, i, &tok, &tt)
	}
}

func checkToken(t *testing.T, i int, current *token.Token, expected *TokenTestExpected) {
	t.Helper()
	if current.Type != expected.Type {
		t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q", i, expected.Type, current.Type)
	}

	if current.Literal != expected.Literal {
		t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, expected.Literal, current.Literal)
	}
}
