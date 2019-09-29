package main

import (
	"container/list"
	"fmt"
)

type List struct {
	cells []interface{}
}

func NewList() *List {
	return &List{}
}

func (l *List) Add(a interface{}) {
	l.cells = append(l.cells, a)
}
func (l *List) Get(p int) interface{} {
	if p >= len(l.cells) {
		return nil
	}
	return l.cells[p]
}

func (l *List) Pop() interface{} {
	r := l.cells[len(l.cells)-1]
	l.cells = l.cells[0 : len(l.cells)-1]
	return r
}

type ASTNode struct {
	parent    *ASTNode
	children  *List
	node_type string
	value     string
}

func NewNode(t string, v string) *ASTNode {
	return &ASTNode{
		node_type: t,
		value:     v,
		children:  NewList(),
	}
}

func (node *ASTNode) AddChild(c *ASTNode) {
	node.children.Add(c)
	c.parent = node
}

func (node *ASTNode) GetParent() *ASTNode {
	return node.parent
}

func (node *ASTNode) GetChildren() *List {
	return node.children
}

func (node *ASTNode) GetType() string {
	return node.node_type
}
func (node *ASTNode) GetValue() string {
	return node.value
}

func (node *ASTNode) DumpNode() {
	fmt.Printf("%v\n%v\n%v\n%v\n", node.GetParent(), node.GetType(), node.GetValue(), node.GetChildren())
}

func Parse(script string) *ASTNode {
	lexer := NewLexer()
	lexer.tokenize(script)
	return astRoot(lexer)
}
func astRoot(l *Lexer) *ASTNode {
	root := NewNode("Program", "Calculator")
	child := additive(l.tokens)
	if child != nil {
		root.AddChild(child)
	}
	return root
}

//TODO
func additive(tokens *list.List) *ASTNode {
	return nil
}

//TODO
func multiplicative(tokens *list.List) *ASTNode {
	return nil
}

func evaluate(node *ASTNode, ident string) (res int) {
	return
}
