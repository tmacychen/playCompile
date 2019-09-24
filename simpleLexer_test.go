package main

import (
	"fmt"
	"testing"
)

func Test_simpleLexer(t *testing.T) {

	fmt.Println("vim-go")
	l := &Lexer{}

	l.tokenize("abc > 10")
	for _, v := range l.tokens {
		fmt.Printf("%v:%v\n", v.GetType(), v.v)
	}
	l = &Lexer{}
	l.tokenize("int a = 19")
	for _, v := range l.tokens {
		fmt.Printf("%v:%v\n", v.GetType(), v.v)
	}
	l1 := &Lexer{}
	for _, v := range l1.tokens {
		fmt.Printf("init :%v:%v\n", v.GetType(), v.v)
	}
	l1.tokenize("2 + 3 = 5")
	for _, v := range l1.tokens {
		fmt.Printf("%v:%v\n", v.GetType(), v.v)
	}
}
