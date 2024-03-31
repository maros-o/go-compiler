package main

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"github.com/fatih/color"
	"go-compiler/parser"
)

type stackListener struct {
	*parser.BaseStackListener
}

func eval(input string) {
	is := antlr.NewInputStream(input)

	lexer := parser.NewStackLexer(is)
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(&colorErrorListener{})
	color.Blue("tokens:\n")
	for {
		t := lexer.NextToken()
		if t.GetTokenType() == antlr.TokenEOF {
			break
		}
		color.Blue("%s (%q)\n", lexer.SymbolicNames[t.GetTokenType()], t.GetText())
		if parser.StackLexerSEM == t.GetTokenType() {
			fmt.Println()
		}
	}
	lexer.Reset()

	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := parser.NewStackParser(stream)
	p.RemoveErrorListeners()
	p.AddErrorListener(&colorErrorListener{})

	var listener stackListener
	antlr.ParseTreeWalkerDefault.Walk(&listener, p.Program())
}
