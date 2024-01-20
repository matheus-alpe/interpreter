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
	input := `=+(){},;`

	tests := []TokenTestExpected {
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPARENT, "("},
		{token.RPARENT, "="},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
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
            t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, expected.Type, current.Type)
        }
}
