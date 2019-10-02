package main

import (
	"fmt"
)

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

//return children
func (node *ASTNode) AddChild(c *ASTNode) *ASTNode {
	node.children.Add(c)
	c.parent = node
	return c
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

func (node *ASTNode) DumpNode(indent string) {
	fmt.Printf("%v%v\n,%v\n", indent, node.GetType(), node.GetValue())
	for i, v := range node.GetChildren().GetValue() {
		v.DumpNode(indent + "\t")
	}
}

func parse(script string) *ASTNode {
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
func additive(l *Lexer) *ASTNode {
	node1 := multiplicative(l)
	t := l.GetPeekToken()
	if t != nil && node != nil {
		if t.GetType() == "Plus" || t.GetType() == "Minus" {
			token := l.PopToken()
			c := NewNode("Additive", t.GetValue())
			node1.AddChild(c)

		}
	}
	return node
}

//TODO
func multiplicative(l *Lexer) *ASTNode {
	return nil
}

func Evaluate(script string) {
	tree := parse(script)
	DumpNode(tree)
	return
}

func intDeclare(l *Lexer) (root *ASTNode) {
	var n, c *ASTNode
	t := l.GetPeekToken() // 预读
	//匹配到int
	if t != nil && t.getType() == "IntToken" {
		t = l.PopToken() //消耗标识
		n = NewNode("IntToken", t.GetValue())
		root = n

		t = l.GetPeekToken()
		//获取一个标识符
		if t != nil && t.GetType() == "Indentifier" {
			t = l.PopToken()
			c = NewNode("Indentifier", t.GetValue())
			n = n.AddChild(c) //加入一个字节点，返回字节点给n，继续向字节点增加节点

			t = l.GetPeekToken()
			if t != nil && t.GetType() == "ASSIGNMENT" {
				t = l.PopToken()
				c = NewNode("ASSIGNMENT", t.GetValue())
				n = n.AddChild(c)
				child := additive(l)
				if child != nil {
					n.AddChild(child)
				}

			}
		}
	}
}
