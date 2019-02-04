package ast

import (
	"bytes"
	"monkey/token"
)

// Expression is anything that produces a value
type Expression interface {
	Node
	expressionNode()
}

// ExpressionStatement allows us to have one line expressions
type ExpressionStatement struct {
	Token      token.Token // the first token of the expression
	Expression Expression
}

func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}

func (es *ExpressionStatement) statementNode() {}

// TokenLiteral on ExpressionStatement tokenizes the one line expression
func (es *ExpressionStatement) TokenLiteral() string {
	return es.Token.Literal
}

// LetStatement is a statement for assigning variables
type LetStatement struct {
	Token token.Token // the token.LET statement
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode() {}

// TokenLiteral provides a token representing the LetStatement
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

func (ls *LetStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")

	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}
	out.WriteString(";")
	return out.String()
}

// Identifier is for the right side of LetStatements
type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode() {}

func (i *Identifier) String() string { return i.Value }

// TokenLiteral is a token of the right side of LetStatements
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

// Node is a basic type for building an AST
type Node interface {
	TokenLiteral() string
	String() string
}

// Program is a series of Statements - this will create the AST
type Program struct {
	Statements []Statement
}

// TokenLiteral of a Program is the root of the AST
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

// Statement do not produce a value
type Statement interface {
	Node
	statementNode()
}

// ReturnStatement are for returning values
type ReturnStatement struct {
	Token       token.Token // the 'return' token
	ReturnValue Expression
}

func (rs *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral() + " ")

	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}

	out.WriteString(";")
	return out.String()
}

func (rs *ReturnStatement) statementNode() {}

// TokenLiteral is a token representing a ReturnStatement
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }
