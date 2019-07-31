package generator

import (
	"fmt"
	"strings"

	"github.com/jahzielv/dest-compiler/parser"
)

// Generate generates target code
func Generate(tree parser.ASTNode) string {
	switch node := tree.(type) {
	case parser.DefNode:
		args := strings.Join(node.ArgNames, ",")
		body := Generate(node.Body)
		code := fmt.Sprintf("function %s(%s) { %s }", node.Name, args, body)
		return code
	case parser.IntegerNode:
		return node.Value()
	case parser.CallNode:
		exprStrs := make([]string, 0)
		for _, expr := range node.ArgExprs {
			exprStrs = append(exprStrs, Generate(expr))
		}
		args := strings.Join(exprStrs, ", ")
		code := fmt.Sprintf("%s(%s);", node.Name, args)
		return code
	case parser.VarRefNode:
		return node.Value()
	case parser.RetNode:
		return fmt.Sprintf("return %s", Generate(node.RetExpr))
	default:
		panic(fmt.Sprintf("Unknown node type %T", tree))

	}
	return ""
}
