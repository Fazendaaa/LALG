package ast

import (
	"bytes"
	"strings"

	"../token"
)

// Node :
type Node interface {
	TokenLiteral() string
	String() string
}

// Statement :
type Statement interface {
	Node
	statementNode()
}

// Expression :
type Expression interface {
	Node
	expressionNode()
}

// Program :
type Program struct {
	Statements []Statement
}

// Identifier :
type Identifier struct {
	Token token.Token
	Type  token.Token
	Value string
}

// VarStatement :
type VarStatement struct {
	Token token.Token
	Type  token.Token
	Name  *Identifier
	Value Expression
}

// ConstStatement :
type ConstStatement struct {
	Token token.Token
	Type  token.Token
	Name  *Identifier
	Value Expression
}

// ExpressionStatement :
type ExpressionStatement struct {
	// The first token of the expression
	Token      token.Token
	Expression Expression
}

// IntegerLiteral :
type IntegerLiteral struct {
	Token token.Token
	Value int64
}

// RealLiteral :
type RealLiteral struct {
	Token token.Token
	Value float64
}

// PrefixExpression :
type PrefixExpression struct {
	Token    token.Token
	Operator string
	Right    Expression
}

// InfixExpression :
type InfixExpression struct {
	Token    token.Token
	Left     Expression
	Operator string
	Right    Expression
}

// BlockStatement :
type BlockStatement struct {
	Token      token.Token
	Statements []Statement
}

// ConditionalExpression :
type ConditionalExpression struct {
	Token       token.Token
	Condition   Expression
	Consequence *BlockStatement
	Alternative *BlockStatement
}

// ProcedureLiteral :
type ProcedureLiteral struct {
	Token      token.Token
	Name       string
	Parameters []*Identifier
	Body       *BlockStatement
}

// CallExpression :
type CallExpression struct {
	Token     token.Token
	Procedure Expression
	Arguments []Expression
}

// statementNode :
func (i *Identifier) statementNode() {}

// statementNode :
func (i *Identifier) expressionNode() {}

// String :
func (i *Identifier) String() string {
	return i.Value
}

// TokenLiteral :
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

// String :
func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

// TokenLiteral :
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}

	return ""
}

// String :
func (vs *VarStatement) String() string {
	var out bytes.Buffer

	out.WriteString(vs.TokenLiteral() + " ")
	out.WriteString(vs.Name.String())
	out.WriteString(": ")
	out.WriteString(vs.Type.Literal)
	out.WriteString(" := ")

	if nil != vs.Value {
		out.WriteString(vs.Value.String())
	}

	out.WriteString(";")

	return out.String()
}

// statementNode :
func (vs *VarStatement) statementNode() {}

// TokenLiteral :
func (vs *VarStatement) TokenLiteral() string {
	return vs.Token.Literal
}

// String :
func (cs *ConstStatement) String() string {
	var out bytes.Buffer

	out.WriteString(cs.TokenLiteral() + " ")
	out.WriteString(cs.Name.String())
	out.WriteString(": ")
	out.WriteString(cs.Type.Literal)
	out.WriteString(" := ")

	if nil != cs.Value {
		out.WriteString(cs.Value.String())
	}

	out.WriteString(";")

	return out.String()
}

// statementNode :
func (cs *ConstStatement) statementNode() {}

// TokenLiteral :
func (cs *ConstStatement) TokenLiteral() string {
	return cs.Token.Literal
}

// String :
func (es *ExpressionStatement) String() string {
	if nil != es.Expression {
		return es.Expression.String()
	}

	return ""
}

// statementNode :
func (es *ExpressionStatement) statementNode() {}

// TokenLiteral :
func (es *ExpressionStatement) TokenLiteral() string {
	return es.Token.Literal
}

// expressionNode :
func (il *IntegerLiteral) expressionNode() {}

// TokenLiteral :
func (il *IntegerLiteral) TokenLiteral() string {
	return il.Token.Literal
}

// String :
func (il *IntegerLiteral) String() string {
	return il.Token.Literal
}

// expressionNode :
func (rl *RealLiteral) expressionNode() {}

// TokenLiteral :
func (rl *RealLiteral) TokenLiteral() string {
	return rl.Token.Literal
}

// String :
func (rl *RealLiteral) String() string {
	return rl.Token.Literal
}

// expressionNode :
func (pe *PrefixExpression) expressionNode() {}

// TokenLiteral :
func (pe *PrefixExpression) TokenLiteral() string {
	return pe.Token.Literal
}

// String :
func (pe *PrefixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.String())
	out.WriteString(")")

	return out.String()
}

// expressionNode :
func (ie *InfixExpression) expressionNode() {}

// TokenLiteral :
func (ie *InfixExpression) TokenLiteral() string {
	return ie.Token.Literal
}

// String :
func (ie *InfixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(ie.Left.String())
	out.WriteString(" " + ie.Operator + " ")
	out.WriteString(ie.Right.String())
	out.WriteString(")")

	return out.String()
}

// expressionNode :
func (bs *BlockStatement) expressionNode() {}

// TokenLiteral :
func (bs *BlockStatement) TokenLiteral() string {
	return bs.Token.Literal
}

// String :
func (bs *BlockStatement) String() string {
	var out bytes.Buffer

	for _, statement := range bs.Statements {
		out.WriteString(statement.String())
	}

	return out.String()
}

// expressionNode :
func (ce *ConditionalExpression) expressionNode() {}

// TokenLiteral :
func (ce *ConditionalExpression) TokenLiteral() string {
	return ce.Token.Literal
}

// String :
func (ce *ConditionalExpression) String() string {
	var out bytes.Buffer

	out.WriteString("if")
	out.WriteString(ce.Condition.String())
	out.WriteString(" ")
	out.WriteString(ce.Consequence.String())
	out.WriteString("end ")

	if nil != ce.Alternative {
		out.WriteString("else ")
		out.WriteString(ce.Alternative.String())
		out.WriteString("end ")
	}

	return out.String()
}

// expressionNode :
func (pl *ProcedureLiteral) expressionNode() {}

// TokenLiteral :
func (pl *ProcedureLiteral) TokenLiteral() string {
	return pl.Token.Literal
}

// String :
func (pl *ProcedureLiteral) String() string {
	var out bytes.Buffer

	parameters := []string{}

	for _, p := range pl.Parameters {
		parameters = append(parameters, p.String())
	}

	out.WriteString(pl.TokenLiteral())
	out.WriteString("begin")
	out.WriteString(strings.Join(parameters, ", "))
	out.WriteString("end ")
	out.WriteString(pl.Body.String())

	return out.String()
}

// expressionNode :
func (ce *CallExpression) expressionNode() {}

// TokenLiteral :
func (ce *CallExpression) TokenLiteral() string {
	return ce.Token.Literal
}

// String :
func (ce *CallExpression) String() string {
	var out bytes.Buffer

	arguments := []string{}

	for _, a := range ce.Arguments {
		arguments = append(arguments, a.String())
	}

	out.WriteString(ce.Procedure.String())
	out.WriteString("(")
	out.WriteString(strings.Join(arguments, ", "))
	out.WriteString(")")

	return out.String()
}
