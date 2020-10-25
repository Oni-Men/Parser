package lexer

import (
	ast ".."
)

var (
	input       string
	pos         int
	currentChar byte
	nextChar    byte
	buffer      []byte
)

//Lexer start lexer
func Lexer(str string) *ast.Sequence {
	input = str
	pos = 0
	seq := ast.NewSequence()

	for {
		nextToken(seq)
		if currentChar == 0 {
			break
		}
	}
	return seq
}

//IsLetter Returns true when given char matched to a-zA-Z_
func IsLetter(c byte) bool {
	return ('a' <= c && c <= 'z') || ('A' <= c && c <= 'Z') || c == '_'
}

//IsDigit Return true when given char matched to 0-9
func IsDigit(c byte) bool {
	return '0' <= c && c <= '9'
}

//IsSymbol Return true when given char is symbol
func IsSymbol(c byte) bool {
	return !IsLetter(c) && !IsDigit(c)
}

func nextToken(seq *ast.Sequence) {
	var token *ast.Token
	read()
	switch currentChar {
	case ' ', '\t':
		token = ast.NewToken(ast.BLANK, currentChar)
	case '\n', '\r':
		token = ast.NewToken(ast.NEWLINE, currentChar)
	case 4:
		token = ast.NewToken(ast.EOF, currentChar)
	case '\\':
		token = ast.NewToken(ast.ESCAPE, currentChar)
	case '=':
		skip(1)
		switch nextChar {
		case '=':
			token = ast.NewTokenLiteral(ast.EQUAL, "==")
		default:
			skip(-1)
			token = ast.NewToken(ast.ASSIGN, currentChar)
		}
	case '+':
		token = ast.NewToken(ast.PLUS, currentChar)
	case '-':
		token = ast.NewToken(ast.MINUS, currentChar)
	case '!':
		skip(1)
		switch nextChar {
		case '=':
			token = ast.NewTokenLiteral(ast.NOTEQ, "!=")
		default:
			skip(-1)
		}
	case '>':
		skip(1)
		switch nextChar {
		case '=':
			token = ast.NewTokenLiteral(ast.GTEQ, ">=")
		default:
			token = ast.NewToken(ast.GT, currentChar)
			skip(-1)
		}
	case '<':
		skip(1)
		switch nextChar {
		case '=':
			token = ast.NewTokenLiteral(ast.LTEQ, "<=")
		default:
			token = ast.NewToken(ast.LT, currentChar)
			skip(-1)
		}
	case ',':
		token = ast.NewToken(ast.COMMA, currentChar)
	case '(':
		token = ast.NewToken(ast.LPAREN, currentChar)
	case ')':
		token = ast.NewToken(ast.RPAREN, currentChar)
	case '{':
		token = ast.NewToken(ast.LBRACE, currentChar)
	case '}':
		token = ast.NewToken(ast.RBRACE, currentChar)
	default:
		buffer = append(buffer, currentChar)
	}
	if token != nil {
		if len(buffer) != 0 {
			tokenType := ast.IDENT
			str := string(buffer)
			switch str {
			case "return":
				tokenType = ast.RETURN
			case "func":
				tokenType = ast.FUNC
			case "let":
				tokenType = ast.LET
			default:
				isDigit := true
				for _, c := range str {
					isDigit = isDigit && IsDigit(byte(c))
				}
				if isDigit {
					tokenType = ast.INT
				} else {
					tokenType = ast.IDENT
				}
			}
			literal := ast.NewTokenLiteral(tokenType, str)
			seq.AddToken(literal)
		}
		seq.AddToken(token)
		buffer = buffer[:0]
	}
}

func read() {
	if pos >= len(input) {
		currentChar = 0
	} else {
		currentChar = input[pos]
	}

	if pos+1 >= len(input) {
		nextChar = 0
	} else {
		nextChar = input[pos+1]
	}

	pos++
}

func skip(n int) {
	pos += n
}
