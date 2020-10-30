package ast

import (
	"../token"
)

//Sequence manage tokens
type Sequence struct {
	TokenList []*token.Token
	pos       int
}

//NewSequence Generate and return Sequence object
func NewSequence() *Sequence {
	seq := &Sequence{}
	seq.TokenList = make([]*token.Token, 0, 5)
	return seq
}

//AddToken Add token to last of sequence
func (seq *Sequence) AddToken(t *token.Token) {
	if len(seq.TokenList) == cap(seq.TokenList) {
		seq.TokenList = append(seq.TokenList, make([]*token.Token, 0, 5)...)
	}
	seq.TokenList = append(seq.TokenList, t)
}

//RemoveToken Remove token from last of sequence
func (seq *Sequence) RemoveToken() {
	seq.TokenList = seq.TokenList[:seq.pos-1]
}
