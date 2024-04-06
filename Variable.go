package main

import (
	"go-compiler/parser"
	"strconv"
)

type Variable struct {
	typeTokenIndex int
	value          interface{}
}

func (v *Variable) GetValueString() string {
	if v.typeTokenIndex == parser.StackLexerINT {
		return strconv.Itoa(v.value.(int))
	} else if v.typeTokenIndex == parser.StackLexerFLOAT {
		return strconv.FormatFloat(v.value.(float64), 'f', -1, 64)
	} else if v.typeTokenIndex == parser.StackLexerSTRING {
		return v.value.(string)
	} else if v.typeTokenIndex == parser.StackLexerBOOL {
		return strconv.FormatBool(v.value.(bool))
	} else {
		panic("Unknown value type")
	}
}

func (v *Variable) GetTypeString() string {
	if v.typeTokenIndex == parser.StackLexerINT {
		return "int"
	} else if v.typeTokenIndex == parser.StackLexerFLOAT {
		return "float"
	} else if v.typeTokenIndex == parser.StackLexerSTRING {
		return "string"
	} else if v.typeTokenIndex == parser.StackLexerBOOL {
		return "bool"
	} else {
		panic("Unknown type type")
	}
}
