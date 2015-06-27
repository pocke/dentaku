package parser

import (
	"io"
	"text/scanner"

	"github.com/pocke/dentaku/ast"
)

type Lexer struct {
	scanner.Scanner
	result ast.Expression
}

func (l *Lexer) Lex(lval *yySymType) int {
	token := int(l.Scan())
	if token == scanner.Int {
		token = LITERAL
	}
	lval.token = Token{Token: token, Literal: l.TokenText()}
	return token
}

func (l *Lexer) Error(e string) {
	panic(e)
}

func Parse(r io.Reader) ast.Expression {
	l := new(Lexer)
	l.Init(r)
	yyParse(l)
	return l.result
}
