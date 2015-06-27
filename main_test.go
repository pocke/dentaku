package main_test

import (
	"strings"
	"testing"

	"github.com/pocke/dentaku/eval"
	"github.com/pocke/dentaku/parser"
)

func TestDentaku(t *testing.T) {
	assert := func(expr string, expected eval.Rational) {
		r := strings.NewReader(expr)
		a := parser.Parse(r)
		got := eval.Evaluate(a)
		if expected != got {
			t.Errorf("Expected %s, but got %s", expected, got)
		}
	}

	assert("1+1", eval.Rational{Numer: 2, Decom: 1})
	assert("1-2", eval.Rational{Numer: -1, Decom: 1})
	assert("3*2", eval.Rational{Numer: 6, Decom: 1})
	assert("24/2", eval.Rational{Numer: 12, Decom: 1})
	assert("24/16", eval.Rational{Numer: 3, Decom: 2})
	assert("(1+2)*(3-5)", eval.Rational{Numer: -6, Decom: 1})
	assert("1+3/6", eval.Rational{Numer: 3, Decom: 2})
}
