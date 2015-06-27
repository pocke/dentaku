package eval

import (
	"fmt"
	"strconv"

	"github.com/pocke/dentaku/ast"
)

// Numer / Decom
type Rational struct {
	Numer int
	Decom int
}

func (r Rational) String() string { return fmt.Sprintf("%d/%d", r.Numer, r.Decom) }
func (r Rational) Float() float64 { return float64(r.Numer) / float64(r.Decom) }
func (r Rational) Int() int       { return r.Numer / r.Decom }

func (r Rational) Reduce() Rational {
	n := gcd(r.Numer, r.Decom)
	return Rational{
		Numer: r.Numer / n,
		Decom: r.Decom / n,
	}
}

func gcd(m, n int) int {
	if n == 0 {
		return m
	}
	return gcd(n, m%n)
}

func Evaluate(a ast.Expression) Rational {
	ret := evaluate(a)
	return ret.Reduce()
}

func evaluate(a ast.Expression) Rational {
	switch e := a.(type) {
	case ast.Literal:
		n, _ := strconv.Atoi(e.Literal)
		return Rational{Numer: n, Decom: 1}
	case ast.BinOpExpr:
		l := evaluate(e.Left)
		r := evaluate(e.Right)
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
