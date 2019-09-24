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
