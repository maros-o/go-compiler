package main

import (
	"github.com/antlr4-go/antlr/v4"
	"github.com/fatih/color"
)

type colorErrorListener struct {
	*antlr.DefaultErrorListener
	errorOccurred bool
}

func (l *colorErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	color.Red("Syntax Error: %s", msg)
	color.Red("Line %d:%d", line, column)
	l.errorOccurred = true
}
