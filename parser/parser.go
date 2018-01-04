package parser

import (
	"github.com/itsHabib/go-interp/ast"
	"github.com/itsHabib/go-interp/lexer"
	"github.com/itsHabib/go-interp/token"
)

type Parser struct {
	l         *lexer.Lexer
	curToken  token.Token
	peekToken token.Token
}

// Creates a new parser instance with a lexer.Lexer field
// and advances it's tokens with the helper method nextToken
func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	// Read two tokens, so curToken and peekToken are set
	p.nextToken()
	p.nextToken()

	return p
}

// Helper method to get current and next token while parsing
func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

// ParseProgram parses the input and creates the valid AST from the given tokens
func (p *Parser) ParseProgram() *ast.Program {
	return nil
}
