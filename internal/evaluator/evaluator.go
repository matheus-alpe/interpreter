package evaluator

import (
	"github.com/matheus-alpe/interpreter/internal/ast"
	"github.com/matheus-alpe/interpreter/internal/object"
)

// Created to not instantiate every time has a bool value, because only have 2 possibles values.
var (
	TRUE  = &object.Boolean{Value: true}
	FALSE = &object.Boolean{Value: false}
)

func Eval(node ast.Node) object.Object {
	switch node := node.(type) {
	case *ast.Program:
		return evalStatements(node.Statements)
	case *ast.ExpressionStatement:
		return Eval(node.Expression)
	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}
	case *ast.Boolean:
		return nativeBoolToBooleanObjecT(node.Value)
	}

	return nil
}

func evalStatements(stmts []ast.Statement) object.Object {
	var result object.Object

	for _, statement := range stmts {
		result = Eval(statement)
	}

	return result
}

func nativeBoolToBooleanObjecT(input bool) *object.Boolean {
	if input {
		return TRUE
	}
	return FALSE
}
