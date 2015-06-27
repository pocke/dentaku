package ast

type Expression interface{}

type Literal struct {
	Literal string
}

type MinusOpExpr struct {
	Operator rune
	Right    Expression
}

type BinOpExpr struct {
	Left     Expression
	Operator rune
	Right    Expression
}
