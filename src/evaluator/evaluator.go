package evaluator

import (
	"fmt"

	"../ast"
	"../object"
)

var (
	NULL = &object.Null{}
	TRUE = &object.Boolean{
		Value: true,
	}
	FALSE = &object.Boolean{
		Value: false,
	}
)

// newError :
func newError(format string, a ...interface{}) *object.Error {
	return &object.Error{
		Message: fmt.Sprintf(format, a...),
	}
}

// evalProgram :
func evalProgram(statements []ast.Statement) object.Object {
	var result object.Object

	for _, statement := range statements {
		result = Eval(statement)

		switch result := result.(type) {
		case *object.ReturnValue:
			return result.Value
		case *object.Error:
			return result
		}
	}

	return result
}

// nativeBoolToBooleanObject :
func nativeBoolToBooleanObject(input bool) object.Object {
	if input {
		return TRUE
	}

	return FALSE
}

// evalBangOperatorExpression :
func evalBangOperatorExpression(right object.Object) object.Object {
	switch right {
	case TRUE:
		return FALSE
	case FALSE:
		return TRUE
	case NULL:
		return TRUE
	default:
		return FALSE
	}
}

// evalMinusPrefixOperatorExpression :
func evalMinusPrefixOperatorExpression(right object.Object) object.Object {
	if right.Type() != object.INTEGER_OBJECT {
		return newError("unknown operator: -%s", right.Type())
	}

	value := right.(*object.Integer).Value

	return &object.Integer{
		Value: -value,
	}
}

// evalPrefixExpression :
func evalPrefixExpression(operator string, right object.Object) object.Object {
	switch operator {
	case "!":
		return evalBangOperatorExpression(right)
	case "-":
		return evalMinusPrefixOperatorExpression(right)
	default:
		return newError("unknown operator: %s%s", operator, right.Type())
	}
}

// evalIntgerInfixExpression :
func evalIntgerInfixExpression(operator string, left, right object.Object) object.Object {
	leftValue := left.(*object.Integer).Value
	rightValue := right.(*object.Integer).Value

	switch operator {
	case "+":
		return &object.Integer{
			Value: leftValue + rightValue,
		}
	case "-":
		return &object.Integer{
			Value: leftValue - rightValue,
		}
	case "*":
		return &object.Integer{
			Value: leftValue * rightValue,
		}
	case "/":
		return &object.Integer{
			Value: leftValue / rightValue,
		}
	case "<":
		return nativeBoolToBooleanObject(leftValue < rightValue)
	case ">":
		return nativeBoolToBooleanObject(leftValue > rightValue)
	case "==":
		return nativeBoolToBooleanObject(leftValue == rightValue)
	case "!=":
		return nativeBoolToBooleanObject(leftValue != rightValue)
	default:
		return newError("unknown operator: %s %s %s", left.Type(), operator, right.Type())
	}
}

// evalInfixExpression :
func evalInfixExpression(operator string, left, right object.Object) object.Object {
	switch {
	case left.Type() == object.INTEGER_OBJECT && right.Type() == object.INTEGER_OBJECT:
		return evalIntgerInfixExpression(operator, left, right)
	case "==" == operator:
		return nativeBoolToBooleanObject(left == right)
	case "!=" == operator:
		return nativeBoolToBooleanObject(left != right)
	case left.Type() != right.Type():
		return newError("type mismatch: %s %s %s", left.Type(), operator, right.Type())
	default:
		return newError("unknown operator: %s %s %s", left.Type(), operator, right.Type())
	}
}

// isTruthy :
func isTruthy(obj object.Object) bool {
	switch obj {
	case NULL:
		return false
	case TRUE:
		return true
	case FALSE:
		return false
	default:
		return true
	}
}

// evalConditionalExpression :
func evalConditionalExpression(ce *ast.ConditionalExpression) object.Object {
	condition := Eval(ce.Condition)

	if isError(condition) {
		return condition
	}
	if isTruthy(condition) {
		return Eval(ce.Consequence)
	} else if nil != ce.Alternative {
		return Eval(ce.Alternative)
	}

	return NULL
}

// evalBlockStatement :
func evalBlockStatement(block *ast.BlockStatement) object.Object {
	var result object.Object

	for _, statement := range block.Statements {
		result = Eval(statement)

		if nil != result {
			resultType := result.Type()

			if resultType == object.RETURN_VALUE_OBJECT || resultType == object.ERROR_OBJECT {
				return result
			}
		}
	}

	return result
}

// isError :
func isError(obj object.Object) bool {
	if nil != obj {
		return object.ERROR_OBJECT == obj.Type()
	}

	return false
}

// Eval :
func Eval(node ast.Node) object.Object {
	switch node := node.(type) {
	case *ast.Program:
		return evalProgram(node.Statements)
	case *ast.ExpressionStatement:
		return Eval(node.Expression)
	case *ast.IntegerLiteral:
		return &object.Integer{
			Value: node.Value,
		}
	case *ast.Boolean:
		return nativeBoolToBooleanObject(node.Value)
	case *ast.PrefixExpression:
		right := Eval(node.Right)

		if isError(right) {
			return right
		}

		return evalPrefixExpression(node.Operator, right)
	case *ast.InfixExpression:
		left := Eval(node.Left)

		if isError(left) {
			return left
		}

		right := Eval(node.Right)

		if isError(right) {
			return right
		}

		return evalInfixExpression(node.Operator, left, right)
	case *ast.BlockStatement:
		return evalBlockStatement(node)
	case *ast.ConditionalExpression:
		return evalConditionalExpression(node)
	case *ast.ReturnStatement:
		value := Eval(node.ReturnValue)

		if isError(value) {
			return value
		}

		return &object.ReturnValue{
			Value: value,
		}
	}

	return nil
}
