package parser

import (
	"fmt"

	"github.com/frankie-mur/monkeylang/ast"
	"github.com/frankie-mur/monkeylang/lexer"
	"github.com/frankie-mur/monkeylang/token"
)

// Parser is a struct that holds the lexer and the current and peek tokens.
// It is used to parse the input tokens and generate an AST representation of the program.
type Parser struct {
	l *lexer.Lexer

	errors []string

	curToken  token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l:      l,
		errors: []string{},
	}

	// Read two tokens, so curToken and peekToken are both set
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) Errors() []string {
	return p.errors
}

// nextToken advances the parser to the next token in the input stream.
func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

// ParseProgram parses the input tokens and returns an AST representation of the program.
// It loops through the tokens, parsing each statement and appending it to the program's
// list of statements. The loop continues until the end of the file is reached.
func (p *Parser) ParseProgram() *ast.Program {
	//Declare the root node
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	// Loop until we have reached the end of the file
	// Each iteration we parse a statement and append it to the program
	for p.curToken.Type != token.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}
	return program
}

// parseStatement parses the current token and returns an AST Statement.
func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()
	default:
		return nil
	}
}

// parseLetStatement parses a let statement, which declares a new variable
// with a name and an initial value. It returns an ast.LetStatement node.
func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: p.curToken}

	if !p.expectedPeek(token.IDENT) {
		return nil
	}

	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	if !p.expectedPeek(token.ASSIGN) {
		return nil
	}

	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

// curTokenIs returns true if the current token is of the given type.
func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

// peekTokenIs checks if the next token is of the specified type.
// It returns true if the next token is of the specified type, false otherwise.
func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

// expectedPeek checks if the next token is of the expected type. If it is, it consumes the token
// and returns true. Otherwise, it appends an error message to the parser's errors and returns false.
func (p *Parser) expectedPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		p.peekError(t)
		return false
	}
}

// peekError appends an error message to the parser's errors slice when the next token
// is not the expected type. The error message includes the expected token type and
// the actual next token type.
func (p *Parser) peekError(t token.TokenType) {
	msg := "expected next token to be %s, got %s instead"
	msg = fmt.Sprintf(msg, t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}
