package main

import "testing"

func Test_ASTNode(t *testing.T) {
}
func Test_intDeclare(t *testing.T) {
}
func Test_intDeclare(t *testing.T) {
	l := NewLexer().tokens("int a = 1")
	a := intDeclare(l)
	a.DumpNode()
}

func Test_Evaluate(t *testing.T) {
	p := Evaluate("1 + 1 = 2")
	p.DumpNode()
}
