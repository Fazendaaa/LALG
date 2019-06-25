package parser

import (
	"fmt"
	"strconv"
	"testing"

	"../ast"
	"../lexer"
)

// testVarStatements :
func testVarStatements(t *testing.T, s ast.Statement, name string) bool {
	if "var" != s.TokenLiteral() {
		t.Errorf("s.TokenLiteral not 'var', got=%q", s.TokenLiteral())

		return false
	}

	varStatement, ok := s.(*ast.VarStatement)

	if !ok {
		t.Errorf("s not *ast.VarStatement, got=%T", s)

		return false
	}
	if varStatement.Name.Value != name {
		t.Errorf("varStatement.Name.Value not '%s', got=%s", name, varStatement.Name.Value)

		return false
	}
	if varStatement.Name.TokenLiteral() != name {
		t.Errorf("varStatement.Name.TokenLiteral() not '%s', got=%s", name, varStatement.Name.TokenLiteral())

		return false
	}

	return true
}

// testConstStatements :
func testConstStatements(t *testing.T, s ast.Statement, name string) bool {
	if "const" != s.TokenLiteral() {
		t.Errorf("s.TokenLiteral not 'const', got=%q", s.TokenLiteral())

		return false
	}

	constStatement, ok := s.(*ast.ConstStatement)

	if !ok {
		t.Errorf("s not *ast.ConstStatement, got=%T", s)

		return false
	}
	if constStatement.Name.Value != name {
		t.Errorf("constStatement.Name.Value not '%s', got=%s", name, constStatement.Name.Value)

		return false
	}
	if constStatement.Name.TokenLiteral() != name {
		t.Errorf("constStatement.Name.TokenLiteral() not '%s', got=%s", name, constStatement.Name.TokenLiteral())

		return false
	}

	return true
}

// testIdentifier :
func testIdentifier(t *testing.T, exp ast.Expression, value string) bool {
	identifier, ok := exp.(*ast.Identifier)

	if !ok {
		t.Errorf("exp not *ast.Identifier, got=%T", exp)

		return false
	}

	if identifier.Value != value {
		t.Errorf("identifier.Value not %s, got=%s", value, identifier.Value)

		return false
	}

	if identifier.TokenLiteral() != value {
		t.Errorf("identifier.TokenLiteral() not '%s', got=%T", value, identifier.TokenLiteral())

		return false
	}

	return true
}

// testIntegerLiteral :
func testIntegerLiteral(t *testing.T, il ast.Expression, value int64) bool {
	integer, ok := il.(*ast.IntegerLiteral)

	if !ok {
		t.Errorf("il not *ast.IntegerLiteral, got=%T", il)

		return false
	}

	if integer.Value != value {
		t.Errorf("integer.Value not '%d', got=%d", value, integer.Value)

		return false
	}

	if integer.TokenLiteral() != fmt.Sprintf("%d", value) {
		t.Errorf("integer.TokenLiteral() not '%d', got=%T", value, integer.TokenLiteral())

		return false
	}

	return true
}

// testRealLiteral :
func testRealLiteral(t *testing.T, il ast.Expression, value float64) bool {
	real, ok := il.(*ast.RealLiteral)

	if !ok {
		t.Errorf("il not *ast.RealLiteral, got=%T", il)

		return false
	}

	if real.Value != value {
		t.Errorf("real.Value not '%f', got=%f", value, real.Value)

		return false
	}

	if real.TokenLiteral() != strconv.FormatFloat(value, 'f', -1, 64) {
		t.Errorf("real.TokenLiteral() not '%f', got=%T", value, real.TokenLiteral())

		return false
	}

	return true
}

// testLiteralExpresion :
func testLiteralExpresion(t *testing.T, expression ast.Expression, expected interface{}) bool {
	switch v := expected.(type) {
	case int:
		return testIntegerLiteral(t, expression, int64(v))
	case int64:
		return testIntegerLiteral(t, expression, v)
	case float64:
		return testRealLiteral(t, expression, v)
	case string:
		return testIdentifier(t, expression, v)
	}

	t.Errorf("type of expression not handled, got=%T", expression)

	return false
}

// testInfixExpression :
func testInfixExpression(t *testing.T, expression ast.Expression, left interface{}, operator string, right interface{}) bool {
	operatorExpression, ok := expression.(*ast.InfixExpression)

	if !ok {
		t.Errorf("expression is not ast.InfixExpression, got=%T(%s)", expression, expression)

		return false
	}

	if !testLiteralExpresion(t, operatorExpression.Left, left) {
		return false
	}

	if operatorExpression.Operator != operator {
		t.Errorf("exp.Operator is not '%s', got=%q", operator, operatorExpression.Operator)

		return false
	}

	if !testLiteralExpresion(t, operatorExpression.Right, right) {
		return false
	}

	return true
}

// checkParserErrors :
func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()

	if 0 == len(errors) {
		return
	}

	t.Errorf("parser has %d errors", len(errors))

	for _, messsage := range errors {
		t.Errorf("parser error: %q", messsage)
	}

	t.FailNow()
}

// TestVarStatements :
func TestVarStatements(t *testing.T) {
	tests := []struct {
		input              string
		expectedIdentifier string
		expectedValue      interface{}
	}{
		{
			"var x: integer := 5;",
			"x",
			5,
		},
		{
			"var y: real :=10;",
			"y",
			10,
		},
		{
			"var foo: real := y;",
			"foo",
			"y",
		},
	}

	for _, tt := range tests {
		l := lexer.InitializeLexer(tt.input)
		p := InitializeParser(l)
		program := p.ParseProgram()

		checkParserErrors(t, p)

		if 1 != len(program.Statements) {
			t.Fatalf("program.Statements does not contains %d statements, got=%d\n", 1, len(program.Statements))
		}

		statement := program.Statements[0]

		if !testVarStatements(t, statement, tt.expectedIdentifier) {
			return
		}

		value := statement.(*ast.VarStatement).Value

		if !testLiteralExpresion(t, value, tt.expectedValue) {
			return
		}
	}
}

// TestConstStatements :
func TestConstStatements(t *testing.T) {
	tests := []struct {
		input              string
		expectedIdentifier string
		expectedValue      interface{}
	}{
		{
			"const x: integer := 5;",
			"x",
			5,
		},
		{
			"const y: real :=10;",
			"y",
			10,
		},
		{
			"const foo: real := y;",
			"foo",
			"y",
		},
	}

	for _, tt := range tests {
		l := lexer.InitializeLexer(tt.input)
		p := InitializeParser(l)
		program := p.ParseProgram()

		checkParserErrors(t, p)

		if 1 != len(program.Statements) {
			t.Fatalf("program.Statements does not contains %d statements, got=%d\n", 1, len(program.Statements))
		}

		statement := program.Statements[0]

		if !testConstStatements(t, statement, tt.expectedIdentifier) {
			return
		}

		value := statement.(*ast.ConstStatement).Value

		if !testLiteralExpresion(t, value, tt.expectedValue) {
			return
		}
	}
}

// TestIdentifierExpression :
func TestIdentifierExpression(t *testing.T) {
	input := "foobar"

	l := lexer.InitializeLexer(input)
	p := InitializeParser(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	if 1 != len(program.Statements) {
		t.Fatalf("program has not enough statements, got=%d", len(program.Statements))
	}

	statement, ok := program.Statements[0].(*ast.ExpressionStatement)

	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement, got=%T", program.Statements[0])
	}

	identifier, ok := statement.Expression.(*ast.Identifier)

	if !ok {
		t.Fatalf("Expression not *ast.Identifier, got=%T", statement.Expression)
	}
	if "foobar" != identifier.Value {
		t.Errorf("identifier.Value not '%s', got=%s", "foobar", identifier.Value)
	}
	if "foobar" != identifier.TokenLiteral() {
		t.Errorf("identifier.TokenLiteral() not '%s', got=%s", "foobar", identifier.TokenLiteral())
	}
}

// TestIntegerLiteralExpression :
func TestIntegerLiteralExpression(t *testing.T) {
	input := "5;"

	l := lexer.InitializeLexer(input)
	p := InitializeParser(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	if 1 != len(program.Statements) {
		t.Fatalf("program has not enough statements, got=%d", len(program.Statements))
	}

	statement, ok := program.Statements[0].(*ast.ExpressionStatement)

	if !ok {
		t.Fatalf("program.Statements[0] is not as.ExpressionStatement, got=%T", program.Statements[0])
	}

	literal, ok := statement.Expression.(*ast.IntegerLiteral)

	if !ok {
		t.Fatalf("expression not *ast.IntegralLiteral, got=%T", statement.Expression)
	}

	if 5 != literal.Value {
		t.Errorf("literal.Value not '%d', got=%d", 5, literal.Value)
	}

	if "5" != literal.TokenLiteral() {
		t.Errorf("literal.TokenLiteral not '%s', got=%s", "5", literal.TokenLiteral())
	}
}

// TestRealLiteralExpression :
func TestRealLiteralExpression(t *testing.T) {
	input := "5.5;"

	l := lexer.InitializeLexer(input)
	p := InitializeParser(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	if 1 != len(program.Statements) {
		t.Fatalf("program has not enough statements, got=%d", len(program.Statements))
	}

	statement, ok := program.Statements[0].(*ast.ExpressionStatement)

	if !ok {
		t.Fatalf("program.Statements[0] is not as.ExpressionStatement, got=%T", program.Statements[0])
	}

	literal, ok := statement.Expression.(*ast.RealLiteral)

	if !ok {
		t.Fatalf("expression not *ast.IntegralLiteral, got=%T", statement.Expression)
	}

	if 5.5 != literal.Value {
		t.Errorf("literal.Value not '%b', got=%b", 5.5, literal.Value)
	}

	if "5.5" != literal.TokenLiteral() {
		t.Errorf("literal.TokenLiteral not '%s', got=%s", "5.5", literal.TokenLiteral())
	}
}

// TestParsingPrefixExpressions :
func TestParsingPrefixExpressions(t *testing.T) {
	prefixTests := []struct {
		input    string
		operator string
		value    interface{}
	}{
		{
			"-15",
			"-",
			15,
		},
		{
			"-15.5",
			"-",
			15.5,
		},
	}

	for _, tt := range prefixTests {
		l := lexer.InitializeLexer(tt.input)
		p := InitializeParser(l)
		program := p.ParseProgram()

		checkParserErrors(t, p)

		if 1 != len(program.Statements) {
			t.Fatalf("program.Statements does not contain '%d' statements, got=%d", 1, len(program.Statements))
		}

		statement, ok := program.Statements[0].(*ast.ExpressionStatement)

		if !ok {
			t.Fatalf("program.Statements[0] is not ast.ExpressionStatement, got=%T", program.Statements[0])
		}

		expression, ok := statement.Expression.(*ast.PrefixExpression)

		if !ok {
			t.Fatalf("statement is not ast.PrefixExpression, got=%T", statement.Expression)
		}

		if expression.Operator != tt.operator {
			t.Fatalf("expression.Operator is not '%s', got=%s", tt.operator, expression.Operator)
		}

		if !testLiteralExpresion(t, expression.Right, tt.value) {
			return
		}
	}
}

// TestParsingInfixExpressions :
func TestParsingInfixExpressions(t *testing.T) {
	infixTest := []struct {
		input      string
		leftValue  interface{}
		operator   string
		rightValue interface{}
	}{
		{
			"5 + 5",
			5,
			"+",
			5,
		},
		{
			"5 - 5",
			5,
			"-",
			5,
		},
		{
			"5 * 5",
			5,
			"*",
			5,
		},
		{
			"5 / 5",
			5,
			"/",
			5,
		},
		{
			"5 > 5",
			5,
			">",
			5,
		},
		{
			"5 < 5",
			5,
			"<",
			5,
		},
		{
			"5 == 5",
			5,
			"==",
			5,
		},
		{
			"5 <> 5",
			5,
			"<>",
			5,
		},
		{
			"5.5 + 5.5",
			5.5,
			"+",
			5.5,
		},
		{
			"5.5 - 5.5",
			5.5,
			"-",
			5.5,
		},
		{
			"5.5 * 5.5",
			5.5,
			"*",
			5.5,
		},
		{
			"5.5 / 5.5",
			5.5,
			"/",
			5.5,
		},
		{
			"5.5 > 5.5",
			5.5,
			">",
			5.5,
		},
		{
			"5.5 < 5.5",
			5.5,
			"<",
			5.5,
		},
		{
			"5.5 == 5.5",
			5.5,
			"==",
			5.5,
		},
		{
			"5.5 <> 5.5",
			5.5,
			"<>",
			5.5,
		},
	}

	for _, tt := range infixTest {
		l := lexer.InitializeLexer(tt.input)
		p := InitializeParser(l)
		program := p.ParseProgram()

		checkParserErrors(t, p)

		if 1 != len(program.Statements) {
			t.Fatalf("program.Statements does not contain %d statements, got=%d", 1, len(program.Statements))
		}

		statement, ok := program.Statements[0].(*ast.ExpressionStatement)

		if !ok {
			t.Fatalf("program.Statements[0] is not ast.ExpressionStatement, got=%T", program.Statements[0])
		}

		if !testInfixExpression(t, statement.Expression, tt.leftValue, tt.operator, tt.rightValue) {
			return
		}
	}
}

// TestOperatorPrecedenceParsing  :
func TestOperatorPrecedenceParsing(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			"-a * b",
			"((-a) * b)",
		},
		{
			"-a * b",
			"((-a) * b)",
		},
		{
			"a + b + c",
			"((a + b) + c)",
		},
		{
			"a * b + c",
			"((a * b) + c)",
		},
		{
			"a * b * c",
			"((a * b) * c)",
		},
		{
			"a * b / c",
			"((a * b) / c)",
		},
		{
			"a + b / c",
			"(a + (b / c))",
		},
		{
			"a + b * c + d / e - f",
			"(((a + (b * c)) + (d / e)) - f)",
		},
		{
			"3 + 4; -5.5 * 5",
			"(3 + 4)((-5.5) * 5)",
		},
		{
			"5 > 4 == 3 < 4",
			"((5 > 4) == (3 < 4))",
		},
		{
			"5 < 0.4 <> 3 > 4",
			"((5 < 0.4) <> (3 > 4))",
		},
		{
			"3 + 4 * 5 == 3 * 1 + 4 * 5",
			"((3 + (4 * 5)) == ((3 * 1) + (4 * 5)))",
		},
		{
			"1 + (2 + 3) + 4",
			"((1 + (2 + 3)) + 4)",
		},
		{
			"(5 + 5) * 2",
			"((5 + 5) * 2)",
		},
		{
			"2 / (5 + 5.2)",
			"(2 / (5 + 5.2))",
		},
		{
			"-(5 + 5)",
			"(-(5 + 5))",
		},
		{
			"a + add(b * c) + d",
			"((a + add((b * c))) + d)",
		},
		{
			"add(a, b, 1, 2 * 3, 4 + 5, add(6.999, 7 * 8))",
			"add(a, b, 1, (2 * 3), (4 + 5), add(6.999, (7 * 8)))",
		},
		{
			"add(a + b + c * d / f + g)",
			"add((((a + b) + ((c * d) / f)) + g))",
		},
	}

	for _, tt := range tests {
		l := lexer.InitializeLexer(tt.input)
		p := InitializeParser(l)
		program := p.ParseProgram()

		checkParserErrors(t, p)

		actual := program.String()

		if actual != tt.expected {
			t.Errorf("expected=%q, got=%q", tt.expected, actual)
		}
	}
}

// TestConditionalIfOnlyExpressions :
func TestConditionalIfOnlyExpressions(t *testing.T) {
	input := `if x < y then x end`

	l := lexer.InitializeLexer(input)
	p := InitializeParser(l)
	program := p.ParseProgram()

	checkParserErrors(t, p)

	if 1 != len(program.Statements) {
		t.Fatalf("program.Statements does not contain %d statements, got=%d", 1, len(program.Statements))
	}

	statement, ok := program.Statements[0].(*ast.ExpressionStatement)

	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement, got=%T", program.Statements[0])
	}

	expression, ok := statement.Expression.(*ast.ConditionalExpression)

	if !ok {
		t.Fatalf("statement.Expression is not ast.ConditionalExpression, got=%T", statement.Expression)
	}

	if !testInfixExpression(t, expression.Condition, "x", "<", "y") {
		return
	}

	if 1 != len(expression.Consequence.Statements) {
		t.Errorf("consequence is not 1 statement, got=%d", len(expression.Consequence.Statements))
	}

	consequence, ok := expression.Consequence.Statements[0].(*ast.ExpressionStatement)

	if !ok {
		t.Fatalf("Statements[0] is not set.ExpressionStatement, got%T", expression.Consequence.Statements[0])
	}

	if !testIdentifier(t, consequence.Expression, "x") {
		return
	}

	if nil != expression.Alternative {
		t.Errorf("exp.Alternative was not nil, got=%+v", expression.Alternative)
	}
}

// TestConditionalIfElseExpressions :
func TestConditionalIfElseExpressions(t *testing.T) {
	input := `if x < y then x end else y end`

	l := lexer.InitializeLexer(input)
	p := InitializeParser(l)
	program := p.ParseProgram()

	checkParserErrors(t, p)

	if 1 != len(program.Statements) {
		t.Fatalf("program.Statements does not contain %d statements, got=%d", 1, len(program.Statements))
	}

	statement, ok := program.Statements[0].(*ast.ExpressionStatement)

	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement, got=%T", program.Statements[0])
	}

	expression, ok := statement.Expression.(*ast.ConditionalExpression)

	if !ok {
		t.Fatalf("statement.Expression is not ast.ConditionalExpression, got=%T", statement.Expression)
	}

	if !testInfixExpression(t, expression.Condition, "x", "<", "y") {
		return
	}

	if 1 != len(expression.Consequence.Statements) {
		t.Errorf("consequence is not 1 statement, got=%d", len(expression.Consequence.Statements))
	}

	consequence, ok := expression.Consequence.Statements[0].(*ast.ExpressionStatement)

	if !ok {
		t.Fatalf("Statements[0] is not ast.ExpressionStatement, got%T", expression.Consequence.Statements[0])
	}

	if !testIdentifier(t, consequence.Expression, "x") {
		return
	}

	if nil == expression.Alternative {
		t.Errorf("exp.Alternative was nil")
	}

	if 1 != len(expression.Alternative.Statements) {
		t.Errorf("alternative is not 1 stamentent, got=%d", len(expression.Alternative.Statements))
	}

	alternative, ok := expression.Alternative.Statements[0].(*ast.ExpressionStatement)

	if !ok {
		t.Fatalf("Statements[0] is not ast.ExpressionStatement, got=%T", expression.Alternative.Statements[0])
	}

	if !testIdentifier(t, alternative.Expression, "y") {
		return
	}
}

// TestProcedureLiteral :
func TestProcedureLiteral(t *testing.T) {
	input := `procedure add(x: integer, y: integer) begin x + y; end`

	l := lexer.InitializeLexer(input)
	p := InitializeParser(l)
	program := p.ParseProgram()

	checkParserErrors(t, p)

	if 1 != len(program.Statements) {
		t.Fatalf("program.Statements does not contain %d statements, got=%d", 1, len(program.Statements))
	}

	statement, ok := program.Statements[0].(*ast.ExpressionStatement)

	if !ok {
		t.Fatalf("program.Statement[0] is not ast.ExpressionStatement, got=%T", program.Statements[0])
	}

	function, ok := statement.Expression.(*ast.ProcedureLiteral)

	if !ok {
		t.Fatalf("statement.Expression is not ast.ProcedureLiteral, got=%T", statement.Expression)
	}

	if 2 != len(function.Parameters) {
		t.Fatalf("function literal parameters wrong, want %d, got=%d", 2, len(function.Parameters))
	}

	testLiteralExpresion(t, function.Parameters[0], "x")
	testLiteralExpresion(t, function.Parameters[1], "y")

	if 1 != len(function.Body.Statements) {
		t.Fatalf("function.Body.Statements has not %d statements, got=%d", 1, len(function.Body.Statements))
	}

	bodyStatements, ok := function.Body.Statements[0].(*ast.ExpressionStatement)

	if !ok {
		t.Fatalf("function body statement is not ast.ExpressionStatement, got=%T", function.Body.Statements[0])
	}

	testInfixExpression(t, bodyStatements.Expression, "x", "+", "y")
}

// TestProcedureParametersParsing :
func TestProcedureParametersParsing(t *testing.T) {
	tests := []struct {
		input              string
		expectedParameters []string
	}{
		{
			input:              "procedure add() begin end",
			expectedParameters: []string{},
		},
		{
			input: "procedure add(x: real) begin end",
			expectedParameters: []string{
				"x",
			},
		},
		{
			input: "procedure add(x: real, y: real, z: real) begin end",
			expectedParameters: []string{
				"x",
				"y",
				"z",
			},
		},
	}

	for _, tt := range tests {
		l := lexer.InitializeLexer(tt.input)
		p := InitializeParser(l)
		program := p.ParseProgram()

		checkParserErrors(t, p)

		statement := program.Statements[0].(*ast.ExpressionStatement)
		function := statement.Expression.(*ast.ProcedureLiteral)

		if len(function.Parameters) != len(tt.expectedParameters) {
			t.Errorf("length parameters wrong, want %d, got=%d", len(function.Parameters), len(tt.expectedParameters))
		}

		for i, identifier := range tt.expectedParameters {
			testLiteralExpresion(t, function.Parameters[i], identifier)
		}
	}
}

// TestCallExporessionParsing :
func TestCallExporessionParsing(t *testing.T) {
	input := "add(1, 2 * 3, 4 + 5)"

	l := lexer.InitializeLexer(input)
	p := InitializeParser(l)
	program := p.ParseProgram()

	checkParserErrors(t, p)

	if 1 != len(program.Statements) {
		t.Fatalf("program.Statements does not contain %d statements, got=%d", 1, len(program.Statements))
	}

	statement, ok := program.Statements[0].(*ast.ExpressionStatement)

	if !ok {
		t.Fatalf("statement is not ExpressionStatement, got=%T", program.Statements[0])
	}

	expression, ok := statement.Expression.(*ast.CallExpression)

	if !ok {
		t.Fatalf("statement.Expression is not ast.CallExpression, got=%T", statement.Expression)
	}

	if !testIdentifier(t, expression.Procedure, "add") {
		return
	}

	if 3 != len(expression.Arguments) {
		t.Fatalf("wrong length of arguments, got=%d", len(expression.Arguments))
	}

	testLiteralExpresion(t, expression.Arguments[0], 1)
	testInfixExpression(t, expression.Arguments[1], 2, "*", 3)
	testInfixExpression(t, expression.Arguments[2], 4, "+", 5)
}
