package ast

//TokenType Identifer type of Token
type TokenType string

//Enum
const (
	//Special
	BLANK   = TokenType("blank")
	NEWLINE = TokenType("newline")
	EOF     = TokenType("eof")
	//Marks
	ESCAPE    = TokenType("escape")
	ASSIGN    = TokenType("assign")
	PLUS      = TokenType("plus")
	MINUS     = TokenType("minus")
	COMMA     = TokenType("comma")
	SEMICOLON = TokenType("semicolon")
	LPAREN    = TokenType("lparen")
	RPAREN    = TokenType("rparen")
	LBRACE    = TokenType("lbrace")
	RBRACE    = TokenType("rbrace")
	EQUAL     = TokenType("equal")
	NOTEQ     = TokenType("noteq")
	GT        = TokenType("gt")
	LT        = TokenType("lt")
	GTEQ      = TokenType("gteq")
	LTEQ      = TokenType("lteq")
	//Literals
	IDENT  = TokenType("ident")
	FUNC   = TokenType("func")
	RETURN = TokenType("return")
	LET    = TokenType("let")
	INT    = TokenType("int")
)

//Token a
type Token struct {
	Type    TokenType
	Literal string
}

//NewToken Generate new token instance with given parameters
func NewToken(tokenType TokenType, literal byte) *Token {
	token := &Token{}
	token.Type = tokenType
	token.Literal = string(literal)
	return token
}

//NewTokenLiteral Generate new token instance with given parameters
func NewTokenLiteral(tokenType TokenType, literal string) *Token {
	token := &Token{}
	token.Type = tokenType
	token.Literal = literal
	return token
}
