package main

import (
	"fmt"
	"os"

	"github.com/pocke/dentaku/parser"
)

func main() {
	a := parser.Parse(os.Stdin)
	fmt.Println(a)
}
