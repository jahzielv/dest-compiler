package main

import (
	"compiler/generator"
	"compiler/parser"
	"compiler/tokenizer"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println(err.Error())
	}
	tkzr := tokenizer.Tokenizer{Code: string(data)}
	tkns, err := tkzr.Tokenize()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		prsr := parser.Parser{Tokens: tkns}
		tree := prsr.Parse()
		runtime := "function add(x, y) {return x + y};"
		test := "console.log(f(1, 2));"
		code := []string{runtime, generator.Generate(tree), test}
		fmt.Println(strings.Join(code, "\n"))
	}
}
