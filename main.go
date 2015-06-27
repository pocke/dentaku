package main

import (
	"fmt"
	"os"

	"github.com/pocke/dentaku/eval"
	"github.com/pocke/dentaku/parser"
)

func main() {
	a := parser.Parse(os.Stdin)
	ret := eval.Evaluate(a)
	fmt.Println(ret)
}
