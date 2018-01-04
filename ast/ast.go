package ast

import (
	"github.com/itsHabib/go-interp/token"
)

// Node in the AST
type Node interface {
	// TokenLiteral returns the value of the token its associated with
	TokenLiteral() string
}

// Statement interface that will represent statement nodes in the AST
type Statement interface {
	Node
	statementNode()
}

// Expression interface will represent expression nodes in the AST
type Expression interface {
	Node
	expressionNode()
}

// Program node will be the root node of every AST
type Program struct {
	Statements []Statement
}

// TokenLiteral
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""

}

// LetStatement struct keeps track of its token, the identifier used
// and the expression to the right side of the equal sign
type LetStatement struct {
	Token token.Token // token.LET
	Name  *Identifier
	Value Expression
}

// ReturnStatement is the struct(node) responsible for holding a return value,
// and its expression
type ReturnStatement struct {
	Token       token.Token
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode() {}
func (rs *ReturnStatement) TokenLiteral() string {
	return rs.Token.Literal
}

func (ls *LetStatement) statementNode() {}
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

// Identifier struct represents an identifier node in the AST
type Identifier struct {
	Token token.Token // token.IDENT
	Value string
}

func (i *Identifier) expressionNode() {}
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}
