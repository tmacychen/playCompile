package playcompile

import (
	"log"
	"testing"
)

//func Test_ASTNode(t *testing.T) {
//}
//func Test_intDeclare(t *testing.T) {
//}
func Test_intDeclare(t *testing.T) {
	l := NewLexer()
	l.tokenize("int a = 10 + 1")
	l.DumpLexer()
	a := intDeclare(l)
	if a != nil {
		a.DumpNode("")
	} else {
		log.Fatalln("a is nil")
	}
}

func Test_Evaluate(t *testing.T) {
	log.Println("1-3+2")
	Evaluate("1-3+2")
	log.Println("1*3+2")
	Evaluate("1*3+2")
	log.Println("1-3*2")
	Evaluate("1-3*2")
}
