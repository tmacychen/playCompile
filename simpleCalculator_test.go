package main

import "testing"

func Test_List(t *testing.T) {
	l := NewList()
	l.Add(1)
	l.Add("abc")
	a := struct {
		abc int
		s   string
	}{
		abc: 10,
		s:   "hello",
	}
	l.Add(a)
	for i, _ := range l.cells {
		t.Logf("l :%v ", l.cells[i])
	}
	println()
	p := l.Pop()
	for i, _ := range l.cells {
		t.Logf("l :%v ", l.cells[i])
	}
	println()
	t.Logf("p :%v\n", p)
}

func Test_ASTNode(t *testing.T) {
	ast := NewNode("add", "+")
	node1 := NewNode("int", "1")
	node2 := NewNode("int", "2")
	node3 := NewNode("=", "=")
	ast.AddChild(node1)
	ast.AddChild(node2)
	node1.AddChild(node3)
	t.Logf("ast's children :%v\n", ast.GetChildren().cells)
	t.Logf("node1's parent':%v\n", node1.GetParent().GetValue())
	t.Logf("node1's children :%v\n", node1.GetChildren().cells)
	t.Logf("node2's children :%v\n", node2.GetChildren().cells)
	t.Logf("node3's parent':%v\n", node3.GetParent().GetValue())
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
