package ast

import (
	"../token"
)

//Node Interface for node
type Node interface {
	TokenLiteral() string
}

//Statement Interface for Statement
type Statement interface {
	Node
	statementNode()
}

//Expression Interface for Expression
type Expression interface {
	Node
	expressionNode()
}

//Program root of all nodes
type Program struct {
	Statements []Statement
}

//NewRootNode generate and return root
func NewProgram() *Program {
	p := &Program{}
	p.Statements = make([]Statement, 0, 10)
	return p
}

//ToStringRoot Get string of root
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}else {
		return ""
	}
}

type LetStatement struct {
	Token token.Token
	Name *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode(){}
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode(){}
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}