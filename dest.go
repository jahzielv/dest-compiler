package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/jahzielv/dest-compiler/generator"
	"github.com/jahzielv/dest-compiler/parser"
	"github.com/jahzielv/dest-compiler/tokenizer"
)

// Compile compiles from Dest to JS.
func Compile(src []byte) string {
	tkzr := tokenizer.Tokenizer{Code: string(src)}
	tkns, err := tkzr.Tokenize()
	if err != nil {
		return err.Error()
	}
	prsr := parser.Parser{Tokens: tkns}
	tree := prsr.Parse()
	runtime := "function add(x, y) {return x + y};"
	test := "console.log(f(1, 2));"
	code := []string{runtime, generator.Generate(tree), test}
	return strings.Join(code, "\n")
}

func main() {
	data, err := ioutil.ReadFile("in.dest")
	// data, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(Compile(data))

}
