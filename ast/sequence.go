package ast

//Sequence manage tokens
type Sequence struct {
	TokenList []*Token
	pos       int
}

//NewSequence Generate and return Sequence object
func NewSequence() *Sequence {
	seq := &Sequence{}
	seq.TokenList = make([]*Token, 0, 5)
	return seq
}

//AddToken Add token to last of sequence
func (seq *Sequence) AddToken(token *Token) {
	if len(seq.TokenList) == cap(seq.TokenList) {
		seq.TokenList = append(seq.TokenList, make([]*Token, 0, 5)...)
	}
	seq.TokenList = append(seq.TokenList, token)
}

//RemoveToken Remove token from last of sequence
func (seq *Sequence) RemoveToken() {
	seq.TokenList = seq.TokenList[:seq.pos-1]
}
