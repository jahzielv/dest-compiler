package main

import (
	"fmt"
	"io/ioutil"

	"github.com/jahzielv/dest-compiler/compiler"
)

func main() {
	data, err := ioutil.ReadFile("in.dest")
	// data, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(compiler.Compile(data))

}
