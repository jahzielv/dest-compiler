package parser

import (
	"fmt"
	"strconv"

	"github.com/jahzielv/dest-compiler/tokenizer"
)

func shift(s *[]tokenizer.Token) tokenizer.Token {
	tkn := (*s)[0]
	*s = (*s)[1:]
	return tkn
}

func (p *Parser) peek(expectedType string) bool {
	return p.Tokens[0].Type == expectedType
}

func (p *Parser) peekOffset(expectedType string, offset int) bool {
	return p.Tokens[offset].Type == expectedType
}

// Parser parses.
type Parser struct {
	Tokens []tokenizer.Token
}

// Parse the tokens
func (p *Parser) Parse() ASTNode {
	return p.parseDef()
}

func (p *Parser) parseDef() ASTNode {
	p.consume("func_def")
	name := p.consume("identifier").Value
	argNames := p.parseArgNames()
	body := p.parseExpr()
	return DefNode{name, argNames, body}
}

func (p *Parser) parseArgNames() []string {
	argNames := make([]string, 0)
	p.consume("oparen")
	if p.peek("identifier") {
		argNames = append(argNames, p.consume("identifier").Value)
		for p.peek("comma") {
			p.consume("comma")
			argNames = append(argNames, p.consume("identifier").Value)
		}
	}
	p.consume("cparen")
	return argNames
}

func (p *Parser) parseExpr() ASTNode {
	// return p.parseInteger()
	if p.peek("integer") {
		return p.parseInteger()
	} else if p.peek("return_stat") {
		return p.parseReturn()
	} else if p.peek("identifier") && p.peekOffset("oparen", 1) {
		return p.parseCall()
	}
	return p.parseVarRef()
}

func (p *Parser) parseInteger() ASTNode {
	numTok := p.consume("integer")
	num, _ := strconv.Atoi(numTok.Value)
	return IntegerNode{num}
}

func (p *Parser) parseCall() ASTNode {
	name := p.consume("identifier").Value
	argExprs := p.parseArgExprs()
	return CallNode{name, argExprs}
}

func (p *Parser) parseArgExprs() []ASTNode {
	argExprs := make([]ASTNode, 0)
	p.consume("oparen")
	if !p.peek("cparen") {
		argExprs = append(argExprs, p.parseExpr())
		for p.peek("comma") {
			p.consume("comma")
			argExprs = append(argExprs, p.parseExpr())
		}
	}
	p.consume("cparen")
	return argExprs
}

func (p *Parser) parseVarRef() ASTNode {
	return VarRefNode{p.consume("identifier").Value}
}

func (p *Parser) parseReturn() ASTNode {
	p.consume("return_stat")
	return RetNode{p.parseExpr()}
}

func (p *Parser) consume(expectedType string) tokenizer.Token {
	token := shift(&p.Tokens)
	if token.Type != expectedType {
		panic(fmt.Sprintf("expected token of type %s but got %s", expectedType, token.Type))
	} else {
		return token
	}
	// return tokenizer.Token{}
}

// ASTNode is a generic interface for all nodes
type ASTNode interface {
	Value() string
}

// RetNode is a return statement
type RetNode struct {
	RetExpr ASTNode
}

// Value makes this an ASTNode
func (r RetNode) Value() string {
	return "return"
}

// DefNode is a node
type DefNode struct {
	Name     string
	ArgNames []string
	Body     ASTNode
}

// Value makes this an ASTNode
func (d DefNode) Value() string {
	return d.Name
}

// IntegerNode hello
type IntegerNode struct {
	value int
}

// Value makes this an ASTNode
func (i IntegerNode) Value() string {
	return strconv.Itoa(i.value)
}

// CallNode is a func call
type CallNode struct {
	Name     string
	ArgExprs []ASTNode
}

// Value makes this an ASTNode
func (c CallNode) Value() string {
	return c.Name
}

// VarRefNode makes vars exprs
type VarRefNode struct {
	value string
}

// Value makes this an ASTNode
func (v VarRefNode) Value() string {
	return v.value
}
