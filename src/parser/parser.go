package parser

import (
	"fmt"
	"strconv"

	"../ast"
	"../lexer"
	"../token"
)

const (
	_           int = iota
	LOWEST          // Starting condition
	EQUALS          // ==
	LESSGREATER     // > or <
	SUM             // +
	PRODUCT         // *
	PREFIX          // -X or !X
	CALL            // myFunction(X)
)

var precedences = map[token.TokenType]int{
	token.EQUAL:            EQUALS,
	token.DIFFERENT:        EQUALS,
	token.LESS_THAN:        LESSGREATER,
	token.GREATER_THAN:     LESSGREATER,
	token.PLUS:             SUM,
	token.MINUS:            SUM,
	token.SLASH:            PRODUCT,
	token.ASTERISK:         PRODUCT,
	token.LEFT_PARENTHESIS: CALL,
}

// Parser :
type Parser struct {
	l      *lexer.Lexer
	errors []string

	currentToken token.Token
	peekToken    token.Token

	prefixParserFunction map[token.TokenType]prefixParserFunction
	infixParserFunction  map[token.TokenType]infixParserFunction
}

type (
	prefixParserFunction func() ast.Expression
	infixParserFunction  func(ast.Expression) ast.Expression
)

// registerPrefix :
func (p *Parser) registerPrefix(tt token.TokenType, fn prefixParserFunction) {
	p.prefixParserFunction[tt] = fn
}

// registerInfix :
func (p *Parser) registerInfix(tt token.TokenType, fn infixParserFunction) {
	p.infixParserFunction[tt] = fn
}

// nextToken :
func (p *Parser) nextToken() {
	p.currentToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

// currentTokenIs :
func (p *Parser) currentTokenIs(t token.TokenType) bool {
	return p.currentToken.Type == t
}

// peekTokenIs :
func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

// peekPrecedence :
func (p *Parser) peekPrecedence() int {
	if p, ok := precedences[p.peekToken.Type]; ok {
		return p
	}

	return LOWEST
}

// expectPeek :
func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()

		return true
	}

	p.peekErrors(t)

	return false
}

// expectType :
func (p *Parser) expectType() token.Token {
	if p.peekTokenIs(token.INTEGER_KEYWORD) {
		p.nextToken()

		return token.Token{
			Type:    "INTEGER_KEYWORD",
			Literal: "integer",
		}
	}
	if p.peekTokenIs(token.REAL_KEYWORD) {
		p.nextToken()

		return token.Token{
			Type:    "REAL_KEYWORD",
			Literal: "real",
		}
	}

	p.nextToken()

	return token.Token{
		Type:    "ILLEGAL",
		Literal: "",
	}
}

// parseExpression :
func (p *Parser) parseExpression(precedence int) ast.Expression {
	prefix := p.prefixParserFunction[p.currentToken.Type]

	if nil == prefix {
		p.noPrefixParserFnError(p.currentToken.Type)

		return nil
	}

	leftExpression := prefix()

	for !p.peekTokenIs(token.SEMICOLON) && precedence < p.peekPrecedence() {
		infix := p.infixParserFunction[p.peekToken.Type]

		if nil == infix {
			return leftExpression
		}

		p.nextToken()

		leftExpression = infix(leftExpression)
	}

	return leftExpression
}

// parseVarStatement :
func (p *Parser) parseVarStatement() *ast.VarStatement {
	statement := &ast.VarStatement{
		Token: p.currentToken,
	}

	if !p.expectPeek(token.IDENTIFIER) {
		return nil
	}

	statement.Name = &ast.Identifier{
		Token: p.currentToken,
		Value: p.currentToken.Literal,
	}

	if !p.expectPeek(token.COLON) {
		return nil
	}

	statement.Type = p.expectType()

	if "ILLEGAL" == statement.Type.Type {
		return nil
	}

	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	p.nextToken()

	statement.Value = p.parseExpression(LOWEST)

	if !p.expectPeek(token.SEMICOLON) {
		return nil
	}

	return statement
}

// parseConstStatement :
func (p *Parser) parseConstStatement() *ast.ConstStatement {
	statement := &ast.ConstStatement{
		Token: p.currentToken,
	}

	if !p.expectPeek(token.IDENTIFIER) {
		return nil
	}

	statement.Name = &ast.Identifier{
		Token: p.currentToken,
		Value: p.currentToken.Literal,
	}

	if !p.expectPeek(token.COLON) {
		return nil
	}

	statement.Type = p.expectType()

	if "ILLEGAL" == statement.Type.Type {
		return nil
	}

	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	p.nextToken()

	statement.Value = p.parseExpression(LOWEST)

	if !p.expectPeek(token.SEMICOLON) {
		return nil
	}

	return statement
}

// parseIdentifier :
func (p *Parser) parseIdentifier() ast.Expression {
	return &ast.Identifier{
		Token: p.currentToken,
		Value: p.currentToken.Literal,
	}
}

// parseStatement :
func (p *Parser) parseStatement() ast.Statement {
	switch p.currentToken.Type {
	case token.VAR:
		return p.parseVarStatement()
	case token.CONST:
		return p.parseConstStatement()
	default:
		return p.parseExpressionStatement()
	}
}

// parseIntegerLiteral :
func (p *Parser) parseIntegerLiteral() ast.Expression {
	literal := &ast.IntegerLiteral{
		Token: p.currentToken,
	}

	value, err := strconv.ParseInt(p.currentToken.Literal, 0, 64)

	if nil != err {
		message := fmt.Sprintf("could not parse '%q' as integer", p.currentToken.Literal)
		p.errors = append(p.errors, message)

		return nil
	}

	literal.Value = value

	return literal
}

// parseRealLiteral :
func (p *Parser) parseRealLiteral() ast.Expression {
	literal := &ast.RealLiteral{
		Token: p.currentToken,
	}

	value, err := strconv.ParseFloat(p.currentToken.Literal, 64)

	if nil != err {
		message := fmt.Sprintf("could not parse '%q' as real", p.currentToken.Literal)
		p.errors = append(p.errors, message)

		return nil
	}

	literal.Value = value

	return literal
}

// noPrefixParserFnError :
func (p *Parser) noPrefixParserFnError(t token.TokenType) {
	message := fmt.Sprintf("no prefix parse function for '%s' was found", t)
	p.errors = append(p.errors, message)
}

// parseExpressionStatement :
func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	statement := &ast.ExpressionStatement{
		Token: p.currentToken,
	}

	statement.Expression = p.parseExpression(LOWEST)

	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return statement
}

// parsePrefixExpression :
func (p *Parser) parsePrefixExpression() ast.Expression {
	expression := &ast.PrefixExpression{
		Token:    p.currentToken,
		Operator: p.currentToken.Literal,
	}

	p.nextToken()

	expression.Right = p.parseExpression(PREFIX)

	return expression
}

// currentPrecedence :
func (p *Parser) currentPrecedence() int {
	if p, ok := precedences[p.currentToken.Type]; ok {
		return p
	}

	return LOWEST
}

// parseInfixExpression :
func (p *Parser) parseInfixExpression(left ast.Expression) ast.Expression {
	expression := &ast.InfixExpression{
		Token:    p.currentToken,
		Operator: p.currentToken.Literal,
		Left:     left,
	}

	precedence := p.currentPrecedence()

	p.nextToken()

	expression.Right = p.parseExpression(precedence)

	return expression
}

// parseGroupedExpression :
func (p *Parser) parseGroupedExpression() ast.Expression {
	p.nextToken()

	expression := p.parseExpression(LOWEST)

	if !p.expectPeek(token.RIGHT_PARENTHESIS) {
		return nil
	}

	return expression
}

// parseBlockStatement :
func (p *Parser) parseBlockStatement() *ast.BlockStatement {
	block := &ast.BlockStatement{
		Token: p.currentToken,
	}
	block.Statements = []ast.Statement{}

	p.nextToken()

	for !p.currentTokenIs(token.END) && !p.currentTokenIs(token.EOF) {
		statement := p.parseStatement()

		if nil != statement {
			block.Statements = append(block.Statements, statement)
		}

		p.nextToken()
	}

	return block
}

// parseConditionalExpression :
func (p *Parser) parseConditionalExpression() ast.Expression {
	expression := &ast.ConditionalExpression{
		Token: p.currentToken,
	}

	p.nextToken()

	expression.Condition = p.parseExpression(LOWEST)

	if !p.expectPeek(token.THEN) {
		return nil
	}

	expression.Consequence = p.parseBlockStatement()

	if p.peekTokenIs(token.ELSE) {
		p.nextToken()

		expression.Alternative = p.parseBlockStatement()
	}

	return expression
}

// parseParameter :
func (p *Parser) parseParameter() *ast.Identifier {
	p.nextToken()

	identifier := &ast.Identifier{
		Token: p.currentToken,
		Value: p.currentToken.Literal,
	}

	if !p.expectPeek(token.COLON) {
		return nil
	}

	p.nextToken()

	identifier.Type = p.currentToken

	p.nextToken()

	return identifier
}

// parseProcedureParameters :
func (p *Parser) parseProcedureParameters() []*ast.Identifier {
	identifiers := []*ast.Identifier{}

	if p.peekTokenIs(token.RIGHT_PARENTHESIS) {
		p.nextToken()

		return identifiers
	}

	identifier := p.parseParameter()
	identifiers = append(identifiers, identifier)

	for p.currentTokenIs(token.COMMA) {
		identifier := p.parseParameter()
		identifiers = append(identifiers, identifier)
	}

	if !p.currentTokenIs(token.RIGHT_PARENTHESIS) {
		return nil
	}

	return identifiers
}

// parseProcedureLiteral :
func (p *Parser) parseProcedureLiteral() ast.Expression {
	literal := &ast.ProcedureLiteral{
		Token: p.currentToken,
	}

	if !p.expectPeek(token.IDENTIFIER) {
		return nil
	}

	literal.Name = p.currentToken.Literal

	if !p.expectPeek(token.LEFT_PARENTHESIS) {
		return nil
	}

	literal.Parameters = p.parseProcedureParameters()

	if !p.expectPeek(token.BEGIN) {
		return nil
	}

	literal.Body = p.parseBlockStatement()

	return literal
}

// parseCallArguments :
func (p *Parser) parseCallArguments() []ast.Expression {
	arguments := []ast.Expression{}

	if p.peekTokenIs(token.RIGHT_PARENTHESIS) {
		p.nextToken()

		return arguments
	}

	p.nextToken()

	arguments = append(arguments, p.parseExpression(LOWEST))

	for p.peekTokenIs(token.COMMA) {
		p.nextToken()
		p.nextToken()

		arguments = append(arguments, p.parseExpression(LOWEST))
	}

	if !p.expectPeek(token.RIGHT_PARENTHESIS) {
		return nil
	}

	return arguments
}

// parseCallExpression :
func (p *Parser) parseCallExpression(procedure ast.Expression) ast.Expression {
	expression := &ast.CallExpression{
		Token:     p.currentToken,
		Procedure: procedure,
	}
	expression.Arguments = p.parseCallArguments()

	return expression
}

// peekErrors :
func (p *Parser) peekErrors(t token.TokenType) {
	message := fmt.Sprintf("Expected next token to be %s, got '%s' instead", t, p.peekToken.Type)
	p.errors = append(p.errors, message)
}

// Errors :
func (p *Parser) Errors() []string {
	return p.errors
}

// ParseProgram :
func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for !p.currentTokenIs(token.EOF) {
		statement := p.parseStatement()

		if nil != statement {
			program.Statements = append(program.Statements, statement)
		}

		p.nextToken()
	}

	return program
}

// InitializeParser :
func InitializeParser(l *lexer.Lexer) *Parser {
	p := &Parser{
		l:      l,
		errors: []string{},
	}

	// Sets the current and peek tokens
	p.nextToken()
	p.nextToken()

	p.prefixParserFunction = make(map[token.TokenType]prefixParserFunction)
	p.infixParserFunction = make(map[token.TokenType]infixParserFunction)

	p.registerPrefix(token.IDENTIFIER, p.parseIdentifier)
	p.registerPrefix(token.INTEGER, p.parseIntegerLiteral)
	p.registerPrefix(token.REAL, p.parseRealLiteral)
	p.registerPrefix(token.MINUS, p.parsePrefixExpression)
	p.registerPrefix(token.LEFT_PARENTHESIS, p.parseGroupedExpression)
	p.registerPrefix(token.IF, p.parseConditionalExpression)
	p.registerPrefix(token.PROCEDURE, p.parseProcedureLiteral)

	p.registerInfix(token.PLUS, p.parseInfixExpression)
	p.registerInfix(token.MINUS, p.parseInfixExpression)
	p.registerInfix(token.SLASH, p.parseInfixExpression)
	p.registerInfix(token.ASTERISK, p.parseInfixExpression)
	p.registerInfix(token.EQUAL, p.parseInfixExpression)
	p.registerInfix(token.DIFFERENT, p.parseInfixExpression)
	p.registerInfix(token.LESS_THAN, p.parseInfixExpression)
	p.registerInfix(token.GREATER_THAN, p.parseInfixExpression)
	p.registerInfix(token.LEFT_PARENTHESIS, p.parseCallExpression)

	return p
}
