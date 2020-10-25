package ast

//INode Interface for node
type INode interface {
	TokenLiteral() string
}

//IStatement Interface for Statement
type IStatement interface {
	INode
}

//IExpression Interface for Expression
type IExpression interface {
	INode
}

//Root Root of all nodes
type Root struct {
	Statements []*IStatement
}

//NewRootNode generate and return root
func NewRootNode() *Root {
	root := &Root{}
	root.Statements = make([]*IStatement, 0, 10)
	return root
}

//ToStringRoot Get string of root
func (root *Root) ToStringRoot() string {
	if len(root.Statements) != 0 {
		return (*root.Statements[0]).TokenLiteral()
	}
	return ""
}

type LetStatement struct {
	Root
	Token Token
}
