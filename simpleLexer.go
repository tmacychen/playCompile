package main

import (
	"errors"
	"fmt"
	"log"
	"strings"
	"unicode"
)

type DfaState int32

// 状态与token(标记)的类型
const (
	Initial DfaState = iota
	If
	Id
	Id_int1
	Id_int2
	Id_int3
	IntLiteral
	Gt
	Ge
	Plus
	Minus
	Star
	Slash
	SemiColon
	LeftParen
	RightParen
	Assignment
)

type tokenType int32

const (
	Indentifier tokenType = iota
	IntToken
	GT
	GE
	PLUS
	MINUS
	STAR
	SLASH
	SEMICOLON
	LEFTPAREN
	RIGHTPAREN
	ASSIGNMENT
)

//标记结构包括标记类型与标记的值
type Token struct {
	t tokenType
	v string
}

func (t Token) GetType() string {
	switch t.t {
	case Indentifier:
		return "Indentifier"
	case IntToken:
		return "IntToken"
	case GT:
		return "GT"
	case PLUS:
		return "Plus"
	case MINUS:
		return "Minus"
	case STAR:
		return "Star"
	case SLASH:
		return "Slash"
	case SEMICOLON:
		return "SemiColon"
	case LEFTPAREN:
		return "LeftParen"
	case RIGHTPAREN:
		return "RightParen"
	case ASSIGNMENT:
		return "Assignment"
	default:
	}
	return ""
}
func (t Token) GetValue() string {
	return t.v
}

var token = Token{}

type Lexer struct {
	tokens *List //保存已解析出来的token
}

//NewLexer create a Lexer
func NewLexer() *Lexer {
	l := &Lexer{tokens: NewList()}
	return l
}

func (l *Lexer) DumpLexer() {
	for _, v := range l.tokens.GetCells() {
		fmt.Printf("%v,%v\n", v.(Token).GetType(), v.(Token).GetValue())
	}
}

func (l *Lexer) GetPeekToken() (Token, error) {
	if l.tokens.GetPos() < 0 {
		return Token{}, errors.New("pos error")
	}
	return l.tokens.Get(l.tokens.GetPos()).(Token), nil
}

func (l *Lexer) GetTokens() *List {
	return l.tokens
}

//读取字符
func (l *Lexer) Read() Token {
	return l.tokens.Read().(Token)
}

//Unread 回溯字符
func (l *Lexer) UnRead() Token {
	return l.tokens.Unread().(Token)
}

var tokenText strings.Builder

//自动状态机，在不同的状态下跃迁
// 解析传入字符串为各种标记
func (l *Lexer) tokenize(s string) {
	state := Initial
	var ch rune
	for _, ch = range []rune(s) {
		switch state {
		case Initial:
			state = l.initToken(ch)
		case Id:
			if isAlpha(ch) || isDigit(ch) {
				tokenText.WriteRune(ch)
			} else {
				state = l.initToken(ch)
			}
		case Gt:
			if ch == '=' {
				state = Ge
				tokenText.WriteRune(ch)
			} else {
				state = l.initToken(ch)
			}
		case Id_int1:
			if ch == 'n' {
				state = Id_int2
				tokenText.WriteRune(ch)
			} else if isDigit(ch) || isAlpha(ch) {
				state = Id
				tokenText.WriteRune(ch)
			} else {
				state = l.initToken(ch)
			}
		case Id_int2:
			if ch == 't' {
				state = Id_int3
				tokenText.WriteRune(ch)
			} else if isDigit(ch) || isAlpha(ch) {
				state = Id
				tokenText.WriteRune(ch)
			} else {
				state = l.initToken(ch)
			}
		case Id_int3:
			if isBlank(ch) {
				l.initToken(ch)
			} else {
				state = Id
				tokenText.WriteRune(ch)
			}
		case IntLiteral:
			if isDigit(ch) {
				tokenText.WriteRune(ch)
			} else {
				state = l.initToken(ch)
			}
		case Plus:
			fallthrough
		case Minus:
			fallthrough
		case Star:
			fallthrough
		case Slash:
			fallthrough
		case Assignment:
			state = l.initToken(ch)

		default:
		}
	}
	//推送最后一个字符
	if tokenText.Len() > 0 {
		// initToken(ch) 会存在保留上一个token的bug，token与tokenText是全局变量
		token.v = tokenText.String()
		l.tokens.Add(token)
		//创建新的临时区域解析标记
		tokenText.Reset()
	}
}

//initToken是状态机的初始状态.
/*
这个初始状态其实并不做停留，它马上进入其他状态
开始解析的时候，进入初始状态；某个Token解析完毕，也进入初始状态，在这里把Token记下来，然后建立一个新的Token。
*/
func (l *Lexer) initToken(ch rune) DfaState {

	//如果lexer的临时文本区域有内容，就保存到已解析的token组里
	if tokenText.Len() > 0 {
		token.v = tokenText.String()
		l.tokens.Add(token)

		//创建新的临时区域解析标记
		tokenText.Reset()
	}

	//设定初始状态
	newState := Initial
	if isAlpha(ch) {
		if ch == 'i' {
			newState = Id_int1
		} else {
			newState = Id
		}
		token.t = Indentifier
		if _, err := tokenText.WriteRune(ch); err != nil {
			log.Fatal("tokenText.WriteRune:", err)
		}
	} else if isDigit(ch) {
		newState = IntLiteral
		token.t = IntToken
		tokenText.WriteRune(ch)
	} else if ch == '>' {
		newState = Gt
		token.t = GT
		tokenText.WriteRune(ch)
	} else if ch == '+' {
		newState = Plus
		token.t = PLUS
		tokenText.WriteRune(ch)
	} else if ch == '-' {
		newState = Minus
		token.t = MINUS
		tokenText.WriteRune(ch)
	} else if ch == '*' {
		newState = Star
		token.t = STAR
		tokenText.WriteRune(ch)
	} else if ch == '/' {
		newState = Slash
		token.t = SLASH
		tokenText.WriteRune(ch)
	} else if ch == ';' {
		newState = SemiColon
		token.t = SEMICOLON
		tokenText.WriteRune(ch)
	} else if ch == '(' {
		newState = LeftParen
		token.t = LEFTPAREN
		tokenText.WriteRune(ch)
	} else if ch == ')' {
		newState = RightParen
		token.t = RIGHTPAREN
		tokenText.WriteRune(ch)
	} else if ch == '=' {
		newState = Assignment
		token.t = ASSIGNMENT
		tokenText.WriteRune(ch)
	}
	return newState
}

func isAlpha(ch rune) bool {
	return unicode.IsLetter(ch)
}
func isDigit(ch rune) bool {
	return unicode.IsDigit(ch)
}
func isBlank(ch rune) bool {
	return unicode.IsSpace(ch)
}
