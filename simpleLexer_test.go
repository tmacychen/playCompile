package playcompile

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
	l = NewLexer()
	fmt.Println("*** 2 + 3 = 5 ***")
	l.tokenize("2 + 3 = 5")
	l.DumpLexer()
}
func Test_GetPeekToken(t *testing.T) {
	l := NewLexer()
	l.tokenize("int a = 10")
	l.Read()
	if token, err := l.GetPeekToken(); err != nil {
		t.Logf("error :%v\n", err)
	} else {
		fmt.Printf("GetPeekTocken :%v\n", token.GetValue())
	}
	// BUG 需要先tokenize才能访问
	// l = NewLexer()
	// fmt.Printf("Pos:%v\n", l.GetTokens().GetPos())
	// if token, err := l.GetPeekToken(); err != nil {
	// 	t.Logf("error :%v\n", err)
	// } else {
	// 	fmt.Printf("GetPeekTocken NewLexer (0) :%v\n", token.GetValue())
	// }

}
