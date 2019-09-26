package main

import (
	"fmt"
	"testing"
)

func Test_simpleLexer(t *testing.T) {
	l := NewLexer()

	fmt.Println("*** abc > 10 ***")
	l.tokenize("abc > 10")
	l.DumpLexer()
	fmt.Println("*** int a = 19 ***")
	l.tokenize("int a = 19")
	l.DumpLexer()
	fmt.Println("*** 2 + 3 = 5 ***")
	l.tokenize("2 + 3 = 5")
	l.DumpLexer()
}
func Test_GetPeekToken(t *testing.T) {
	l := NewLexer()
	l.tokenize("int a = 10")
	fmt.Printf("GetPeekTocken :%v == int\n", l.GetPeekTocken().v)
	a := l.PopToken()
	fmt.Printf("pop token should be (int) :%v\n", a.v)
	fmt.Printf("Now GetPeekTocken should be a :%v\n", l.GetPeekTocken().v)

}
