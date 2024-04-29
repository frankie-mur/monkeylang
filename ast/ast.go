package ast

import (
	"bytes"

	"github.com/frankie-mur/monkeylang/token"
)

// Node is an interface that represents a node in the abstract syntax tree (AST).
// The TokenLiteral method returns the literal representation of the token
// associated with the node.
type Node interface {
	TokenLiteral() string
	String() string
}

// Statement is an interface that represents a statement in the abstract syntax tree.
// Statements are the building blocks of a program, and can include things like
// variable declarations, function calls, control flow statements, and more.
// The statementNode method is used to satisfy the Node interface, and is a
// marker method that indicates that a type implements the Statement interface.
type Statement interface {
	Node
	statementNode()
}

// Expression is an interface that represents an expression in the abstract syntax tree.
// Expressions can be evaluated to produce a value.
// The expressionNode method is a marker method that identifies a type as an Expression.
type Expression interface {
	Node
	expressionNode()
}

// Program is the root node of an abstract syntax tree.
// It represents a complete Monkey program and holds a slice of Statement nodes.
type Program struct {
	Statements []Statement
}

// String returns a string representation of the program's statements.
// It iterates through the program's statements and calls the String()
// method on each one, concatenating the results into a single string.
func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

// LetStatement represents a let statement in the Monkey programming language.
// It consists of a token representing the 'let' keyword, an Identifier for the
// variable name, and an Expression for the assigned value.
type LetStatement struct {
	Token token.Token // the token.LET token
	Name  *Identifier
	Value Expression
}

// Methods on LetStatement to satisfy the Statement interface.
func (l *LetStatement) statementNode()       {}
func (l *LetStatement) TokenLiteral() string { return l.Token.Literal }
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

// ReturnStatement represents the return statement in the language.
// It holds the 'return' token and the expression to be returned.
type ReturnStatement struct {
	Token token.Token // the'return' token
	Value Expression
}

// Methods on ReturnStatement to satisfy the Statement interface.
func (r *ReturnStatement) statementNode()       {}
func (r *ReturnStatement) TokenLiteral() string { return r.Token.Literal }
func (rs *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral() + " ")

	if rs.Value != nil {
		out.WriteString(rs.Value.String())
	}

	out.WriteString(";")

	return out.String()
}

// ExpressionStatement represents an expression statement in the AST.
// An expression statement is a standalone expression that is evaluated for its side effects.
type ExpressionStatement struct {
	Token      token.Token // the first token of the expression
	Expression Expression
}

// Methods on ExpressionStatement to satisfy the Statement interface.
func (e *ExpressionStatement) statementNode()       {}
func (e *ExpressionStatement) TokenLiteral() string { return e.Token.Literal }
func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}

type Identifier struct {
	Token token.Token // the token.IDENT token
	Value string
}

// Methods on Identifier to satisfy the Expression interface.
func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
func (i *Identifier) String() string       { return i.Value }

// TokenLiteral returns the token literal of the first statement in the program.
// If the program has no statements, it returns an empty string.
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}
