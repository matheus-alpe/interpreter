package parser

import (
	"testing"

	"github.com/matheus-alpe/interpreter/internal/ast"
	"github.com/matheus-alpe/interpreter/internal/lexer"
)

func TestLetStatements(t *testing.T) {
	input := `
    let x = 5;
    let y = 10;
    let foobar = 838383;
    `

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	if program == nil {
		t.Fatal("ParseProgram() returned nil")
	}

	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d", len(program.Statements))
	}

	tests := []struct {
		expectedIdentifer string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testLetStatements(t, stmt, tt.expectedIdentifer) {
			return
		}
	}
}

func testLetStatements(t *testing.T, s ast.Statement, expectedIdentifer string) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf("s.TokenLiteral not 'let'. got=%q", s.TokenLiteral())
		return false
	}

	letStmt, ok := s.(*ast.LetStatement)
	if !ok {
		t.Errorf("s not *ast.LetStatement. got=%T", s)
		return false
	}

	if letStmt.Name.Value != expectedIdentifer {
		t.Errorf("letStmt.Name.Value not %q. got=%q", expectedIdentifer, letStmt.Name.Value)
		return false
	}

	if letStmt.Name.TokenLiteral() != expectedIdentifer {
		t.Errorf("s.Name not %q. got=%q", expectedIdentifer, letStmt.Name)
		return false
	}

	return true
}
