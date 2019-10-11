package playcompile

import (
	"fmt"
	"log"
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
	fmt.Printf("%v%v %v\n", indent, node.GetType(), node.GetValue())
	for _, v := range node.GetChildren().GetCells() {
		v.(*ASTNode).DumpNode(indent + "\t")
	}
}

func parse(script string) *ASTNode {
	lexer := NewLexer()
	lexer.tokenize(script)
	return astRoot(lexer)
}
func astRoot(l *Lexer) *ASTNode {
	root := NewNode("Program", "Calculator")
	child := additive_a(l)
	if child != nil {
		root.AddChild(child)
	}
	return root
}

//加法表达式
func additive(l *Lexer) *ASTNode {
	var node *ASTNode
	node1 := multiplicative(l)
	node = node1
	t, err := l.GetPeekToken()
	if err == nil && node1 != nil {
		if t.GetType() == "Plus" || t.GetType() == "Minus" {
			token := l.Read()
			node2 := additive(l)
			if node2 != nil {
				node = NewNode("Additive", token.GetValue())
				node.AddChild(node1)
				node.AddChild(node2)
			} else {
				log.Fatalf("invalid additive expression,expecting the correct right part")
			}

		}
	}
	return node
}

//解决左递归问题
func additive_a(l *Lexer) *ASTNode {
	var node *ASTNode
	node1 := multiplicative(l)
	node = node1
	token, err := l.GetPeekToken()
	for err == nil && token.GetType() == "Plus" || token.GetType() == "Minus" {
		token = l.Read()
		node2 := multiplicative(l)
		node = NewNode("Additive", token.GetValue())
		node.AddChild(node1)
		node.AddChild(node2)
		node1 = node
		token, err = l.GetPeekToken()
	}
	return node

}

//乘法表达式
func multiplicative(l *Lexer) *ASTNode {
	node1 := primary(l)
	var node *ASTNode
	node = node1
	t, err := l.GetPeekToken()
	if err == nil && node1 != nil {
		if t.GetType() == "Star" || t.GetType() == "Slash" {
			token := l.Read()
			node2 := primary(l)
			if node2 != nil {
				node = NewNode("Muliplicative", token.GetValue())
				node.AddChild(node1)
				node.AddChild(node2)
			} else {
				log.Fatalf("invalid multiplicative expression, expecting the right part.")
			}
		}
	}
	return node
}

//基础表达式
func primary(l *Lexer) *ASTNode {
	var node *ASTNode
	t, err := l.GetPeekToken()
	//log.Printf("primary get peek token %v \t %v\n", t.GetValue(), t.GetType())
	if err == nil {
		if t.GetType() == "IntToken" {
			token := l.Read()
			node = NewNode("IntLiteral", token.GetValue())
		} else if t.GetType() == "Identifier" {
			token := l.Read()
			node = NewNode("Identifier", token.GetValue())
		} else if t.GetType() == "LeftParen" {
			l.Read()
			node = additive(l)
			if node != nil {
				t, err := l.GetPeekToken()
				if err == nil && t.GetType() == "RightParen" {
					l.Read()
				} else {
					log.Fatalln("expecting the right parenthesis")
				}
			} else {
				log.Fatalln("expecting an additive expression in parenthesis")
			}
		}
	}
	return node
}

//Evaluate 计算表达式的值
func Evaluate(script string) {
	tree := parse(script)
	tree.DumpNode("")
	return
}

func intDeclare(l *Lexer) *ASTNode {
	var n, c *ASTNode
	t, err := l.GetPeekToken() // 预读
	//	log.Printf("get peek token  token %v \t %v\n", t.GetValue(), t.GetType())
	//匹配到int
	if err == nil && t.GetType() == "IntKey" {
		t = l.Read() //消耗标识
		//	log.Printf("read a token %v\n", t.GetValue())
		t, err = l.GetPeekToken()
		//获取一个标识符
		if err == nil && t.GetType() == "Indentifier" {
			token := l.Read()
			n = NewNode("IntDeclaration", token.GetValue())
			t, err = l.GetPeekToken()
			//	log.Printf("get peek token  token %v \t %v\n", t.GetValue(), t.GetType())
			if err == nil && t.GetType() == "Assignment" {
				l.Read()
				c = additive(l)
				if c != nil {
					n.AddChild(c)
				}
			}
		} else {
			l.UnRead()
			log.Printf("unread a token\n")
			return nil
		}
	}
	return n
}
