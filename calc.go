package main

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"go-compiler/parser"
	"strconv"
)

type calcListener struct {
	*parser.BaseCalcListener

	stack []int
}

func (l *calcListener) push(i int) {
	l.stack = append(l.stack, i)
}

func (l *calcListener) pop() int {
	if len(l.stack) == 0 {
		panic("stack is empty, unable to pop")
	}

	last := l.stack[len(l.stack)-1]
	l.stack = l.stack[:len(l.stack)-1]

	return last
}

func (l *calcListener) ExitMulDiv(c *parser.MulDivContext) {
	right, left := l.pop(), l.pop()

	fmt.Printf("MulDiv: %d %s %d\n", left, c.GetOp().GetText(), right)

	switch c.GetOp().GetTokenType() {
	case parser.CalcParserMUL:
		l.push(left * right)
	case parser.CalcParserDIV:
		l.push(left / right)
	default:
		panic(fmt.Sprintf("unexpected op: %s", c.GetOp().GetText()))
	}
}

func (l *calcListener) ExitAddSub(c *parser.AddSubContext) {
	right, left := l.pop(), l.pop()

	fmt.Printf("AddSub: %d %s %d\n", left, c.GetOp().GetText(), right)

	switch c.GetOp().GetTokenType() {
	case parser.CalcParserADD:
		l.push(left + right)
	case parser.CalcParserSUB:
		l.push(left - right)
	default:
		panic(fmt.Sprintf("unexpected op: %s", c.GetOp().GetText()))
	}
}

func (l *calcListener) ExitNumber(c *parser.NumberContext) {
	i, err := strconv.Atoi(c.GetText())
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("Number: %d\n", i)

	l.push(i)
}

func calc(input string) int {
	is := antlr.NewInputStream(input)

	lexer := parser.NewCalcLexer(is)
	for {
		t := lexer.NextToken()
		if t.GetTokenType() == antlr.TokenEOF {
			break
		}
		fmt.Println(t.String())
	}
	lexer.Reset()

	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := parser.NewCalcParser(stream)

	var listener calcListener
	antlr.ParseTreeWalkerDefault.Walk(&listener, p.Expression())

	return listener.pop()
}
