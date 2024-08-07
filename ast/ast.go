package ast

import (
	"bytes"
	"strings"

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
	Token       token.Token // the'return' token
	Value       Expression
	ReturnValue Expression
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

// Identifier represents an identifier token in the AST.
// The Token field holds the token.IDENT token, and the Value field holds the identifier value.
type Identifier struct {
	Token token.Token // the token.IDENT token
	Value string
}

// Methods on Identifier to satisfy the Expression interface.
func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
func (i *Identifier) String() string       { return i.Value }

// IntegerLiteral represents an integer literal expression in the AST.
// The Token field holds the token.INT token, and the Value field holds the integer value.
type IntegerLiteral struct {
	Token token.Token // the token.INT token
	Value int64
}

// Methods on IntegerLiteral to satisfy the Expression interface.
func (il *IntegerLiteral) expressionNode()      {}
func (il *IntegerLiteral) TokenLiteral() string { return il.Token.Literal }
func (il *IntegerLiteral) String() string       { return il.Token.Literal }

type StringLiteral struct {
	Token token.Token // the token.STRING token
	Value string
}

// Methods on StringLiteral to satisfy the Expression interface.
func (sl *StringLiteral) expressionNode()      {}
func (sl *StringLiteral) TokenLiteral() string { return sl.Token.Literal }
func (sl *StringLiteral) String() string       { return sl.Token.Literal }

// PrefixExpression represents a prefix expression in the abstract syntax tree.
// It contains the prefix token (e.g. "!", "-"), the operator, and the right-hand expression.
type PrefixExpression struct {
	Token    token.Token // the prefix token, e.g.!, -
	Operator string
	Right    Expression // expression to the right of the operator
}

// Methods on PrefixExpression to satisfy the Expression interface.
func (pe *PrefixExpression) expressionNode()      {}
func (pe *PrefixExpression) TokenLiteral() string { return pe.Token.Literal }
func (pe *PrefixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.String())
	out.WriteString(")")

	return out.String()
}

// InfixExpression represents an infix expression in the AST.
// It contains the operator token, the left and right expressions.
type InfixExpression struct {
	Token    token.Token // the operator token, e.g. +, -, etc.
	Operator string
	Left     Expression // expression to the left of the operator
	Right    Expression // expression to the right of the operator
}

// Methods on InfixExpression to satisfy the Expression interface.
func (ie *InfixExpression) expressionNode()      {}
func (ie *InfixExpression) TokenLiteral() string { return ie.Token.Literal }
func (ie *InfixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(ie.Left.String())
	out.WriteString(" " + ie.Operator + " ")
	out.WriteString(ie.Right.String())
	out.WriteString(")")

	return out.String()
}

// Boolean represents a boolean value in the Monkey programming language.
// It contains a Token, which is the token that represents the boolean value,
// and a Value field that holds the actual boolean value.
type Boolean struct {
	Token token.Token
	Value bool
}

// Methods on Boolean to satisfy the Expression interface.
func (b *Boolean) expressionNode()      {}
func (b *Boolean) TokenLiteral() string { return b.Token.Literal }
func (b *Boolean) String() string       { return b.Token.Literal }

// IfExpression represents an if-else expression in the language.
// It contains the 'if' token, the condition expression, the consequence block,
// and an optional alternative block.
type IfExpression struct {
	Token       token.Token //the 'if' token
	Condition   Expression
	Consequence *BlockStatement
	Alternative *BlockStatement
}

// Methods on IfExpression to satisfy the Expression interface.
func (ie *IfExpression) expressionNode()      {}
func (ie *IfExpression) TokenLiteral() string { return ie.Token.Literal }
func (ie *IfExpression) String() string {
	var out bytes.Buffer

	out.WriteString("if")
	out.WriteString(ie.Condition.String())
	out.WriteString(" ")
	out.WriteString(ie.Consequence.String())
	if ie.Alternative != nil {
		out.WriteString(" else ")
		out.WriteString(ie.Alternative.String())
	}

	return out.String()
}

// BlockStatement represents a block of statements. The block is delimited
// by a pair of curly braces { }.
type BlockStatement struct {
	Token      token.Token //the '{' token
	Statements []Statement
}

func (bs *BlockStatement) statementNode()       {}
func (bs *BlockStatement) TokenLiteral() string { return bs.Token.Literal }
func (bs *BlockStatement) String() string {
	var out bytes.Buffer

	for _, s := range bs.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

// FunctionLiteral represents a function literal expression in the Monkey programming language.
// It contains the 'fn' token, the function parameters, and the function body.
type FunctionLiteral struct {
	Token      token.Token   // the 'fn' token
	Parameters []*Identifier // the function parameters
	Body       *BlockStatement
}

// Methods on FunctionLiteral to satisfy the Expression interface.
func (fl *FunctionLiteral) expressionNode()      {}
func (fl *FunctionLiteral) TokenLiteral() string { return fl.Token.Literal }
func (fl *FunctionLiteral) String() string {
	var out bytes.Buffer

	params := []string{}
	for _, param := range fl.Parameters {
		params = append(params, param.String())
	}

	out.WriteString(fl.Token.Literal)
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") ")
	out.WriteString(fl.Body.String())

	return out.String()
}

// CallExpression represents a function call expression in the AST.
// It contains the function being called, and the arguments passed to it.
type CallExpression struct {
	Token     token.Token // the '(' token
	Function  Expression  // Identifier or FunctionLiteral
	Arguments []Expression
}

// // Methods on callExpression to satisfy the Expression interface.
func (ce *CallExpression) expressionNode()      {}
func (ce *CallExpression) TokenLiteral() string { return ce.Token.Literal }
func (ce *CallExpression) String() string {
	var out bytes.Buffer

	args := []string{}
	for _, a := range ce.Arguments {
		args = append(args, a.String())
	}

	out.WriteString(ce.Function.String())
	out.WriteString("(")
	out.WriteString(strings.Join(args, ", "))
	out.WriteString(")")

	return out.String()
}

// ArrayLiteral represents an array literal in the Monkey programming language.
// It contains the '[' token, the elements of the array, and the ']' token.
type ArrayLiteral struct {
	Token    token.Token // the '[' token
	Elements []Expression
}

func (al *ArrayLiteral) expressionNode()      {}
func (al *ArrayLiteral) TokenLiteral() string { return al.Token.Literal }
func (al *ArrayLiteral) String() string {
	var out bytes.Buffer
	elements := []string{}
	for _, el := range al.Elements {
		elements = append(elements, el.String())
	}
	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")

	return out.String()
}

// IndexExpression represents an expression that accesses an element of an array or map using an index expression.
// The Left expression evaluates to the array or map, and the Index expression evaluates to the index or key to access.
type IndexExpression struct {
	Token token.Token // the '[' token
	Left  Expression
	Index Expression
}

func (ie *IndexExpression) expressionNode()      {}
func (ie *IndexExpression) TokenLiteral() string { return ie.Token.Literal }
func (ie *IndexExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(ie.Left.String())
	out.WriteString("[")
	out.WriteString(ie.Index.String())
	out.WriteString("])")

	return out.String()
}

// HashLiteral represents a hash literal expression in the Monkey programming language.
// It contains the '{' token, a map of key-value pairs, and the '}' token.
type HashLiteral struct {
	Token token.Token // the '{' token
	Pairs map[Expression]Expression
}

func (hl *HashLiteral) expressionNode()      {}
func (hl *HashLiteral) TokenLiteral() string { return hl.Token.Literal }
func (hl *HashLiteral) String() string {
	var out bytes.Buffer

	pairs := []string{}
	for key, value := range hl.Pairs {
		pairs = append(pairs, key.String()+":"+value.String())
	}

	out.WriteString("{")
	out.WriteString(strings.Join(pairs, ", "))
	out.WriteString("}")

	return out.String()
}

// TokenLiteral returns the token literal of the first statement in the program.
// If the program has no statements, it returns an empty string.
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}
