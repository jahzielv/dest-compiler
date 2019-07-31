# A Toy Compiler, in Go

This is a compiler that I wrote by following along with Gary Bernhardt's [Compiler screencast](https://www.destroyallsoftware.com/screencasts).

I called the language _Dest_ after Destroy All Software, Gary's screencast company.

## Installation

```bash
go get github.com/jahzielv/dest-compiler
```

## Running

```bash
dest-compiler yourDestFile.dest
```

The compiler targets JavaScript. Running the above command will output the compiled JavaScript to stdout.

To execute your Dest code, pipe the output into the Node.js REPL, like so

```bash
dest-compiler yourDestFile.dest | node
```
