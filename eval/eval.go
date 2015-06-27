package eval

import (
	"strconv"

	"github.com/pocke/dentaku/ast"
)

// Numer / Decom
type Rational struct {
	Numer int
	Decom int
}

func Evaluate(a ast.Expression) Rational {
	switch e := a.(type) {
	case ast.Literal:
		n, _ := strconv.Atoi(e.Literal)
		return Rational{Numer: n, Decom: 1}
	case ast.BinOpExpr:
		l := Evaluate(e.Left)
		r := Evaluate(e.Right)
		switch e.Operator {
		case '+':
			return Rational{
				Numer: l.Numer*r.Decom + r.Numer*l.Decom,
				Decom: l.Decom * r.Decom,
			}
		case '-':
			return Rational{
				Numer: l.Numer*r.Decom - r.Numer*l.Decom,
				Decom: l.Decom * r.Decom,
			}
		case '*':
			return Rational{
				Numer: l.Numer * r.Numer,
				Decom: l.Decom * r.Decom,
			}
		case '/':
			return Rational{
				Numer: l.Numer * r.Decom,
				Decom: l.Decom * r.Numer,
			}
		default:
			panic("Error")
		}
	default:
		panic("Error")
	}
}
