package main

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"go-compiler/parser"
	"strconv"
)

type Visitor struct {
	*parser.BaseStackVisitor
	variables map[string]Variable
}

func (v *Visitor) PrintVariables() {
	fmt.Println("Variables:")
	fmt.Printf("%-10s %-10s %-10s \n", "ID", "Type", "Value")
	for name, variable := range v.variables {
		fmt.Printf("%-10s %-10s %-10s \n", name, variable.GetTypeString(), variable.GetValueString())
	}
}

func (v *Visitor) Visit(t antlr.ParseTree) interface{} {
	return t.Accept(v)
}

func (v *Visitor) VisitProgram(c *parser.ProgramContext) interface{} {
	println("VisitProgram")
	v.variables = make(map[string]Variable)
	for _, ch := range c.GetChildren() {
		v.Visit(ch.(antlr.ParseTree))
	}
	return nil
}

func (v *Visitor) VisitVariableStatement(c *parser.VariableStatementContext) interface{} {
	typeTokenIndex := c.Type_().GetChild(0).(antlr.TerminalNode).GetSymbol().GetTokenType()

	for _, ch := range c.VariableList().GetChildren() {
		id := ch.(antlr.TerminalNode)
		if id.GetSymbol().GetTokenType() != parser.StackLexerID {
			continue
		}

		name := id.GetText()
		if _, exists := v.variables[name]; exists {
			panic(fmt.Sprintf("Variable '%s' already exists, line: %d", id.GetSymbol().GetText(), id.GetSymbol().GetLine()))
		}

		if typeTokenIndex == parser.StackLexerINT {
			v.variables[name] = Variable{typeTokenIndex, 0}
		} else if typeTokenIndex == parser.StackLexerFLOAT {
			v.variables[name] = Variable{typeTokenIndex, 0.0}
		} else if typeTokenIndex == parser.StackLexerSTRING {
			v.variables[name] = Variable{typeTokenIndex, ""}
		} else if typeTokenIndex == parser.StackLexerBOOL {
			v.variables[name] = Variable{typeTokenIndex, false}
		} else {
			panic(fmt.Sprintf("Unknown type: %d", typeTokenIndex))
		}
	}
	return nil
}

func (v *Visitor) VisitExpressionStatement(c *parser.ExpressionStatementContext) interface{} {
	v.Visit(c.Expression())
	return nil
}

func (v *Visitor) VisitIdExpression(c *parser.IdExpressionContext) interface{} {
	id := c.ID().GetText()
	if _, exists := v.variables[id]; !exists {
		panic(fmt.Sprintf("Variable '%s' does not exist, line: %d", id, c.ID().GetSymbol().GetLine()))
	}
	return v.variables[id]
}

func (v *Visitor) VisitAddSubExpression(c *parser.AddSubExpressionContext) interface{} {
	left := v.Visit(c.Expression(0)).(Variable)
	right := v.Visit(c.Expression(1)).(Variable)

	if left.typeTokenIndex != right.typeTokenIndex {
		panic(fmt.Sprintf("AddSub different types of variables, line: %d", c.GetStart().GetLine()))
	}

	if left.typeTokenIndex == parser.StackLexerINT {
		if c.GetOp().GetTokenType() == parser.StackLexerADD {
			return Variable{parser.StackLexerINT, left.value.(int) + right.value.(int)}
		} else if c.GetOp().GetTokenType() == parser.StackLexerSUB {
			return Variable{parser.StackLexerINT, left.value.(int) - right.value.(int)}
		}
	} else if left.typeTokenIndex == parser.StackLexerFLOAT {
		if c.GetOp().GetTokenType() == parser.StackLexerADD {
			return Variable{parser.StackLexerFLOAT, left.value.(float64) + right.value.(float64)}
		} else if c.GetOp().GetTokenType() == parser.StackLexerSUB {
			return Variable{parser.StackLexerFLOAT, left.value.(float64) - right.value.(float64)}
		}
	}
	panic(fmt.Sprintf("Unknown type: %d", left.typeTokenIndex))
}

func (v *Visitor) VisitLiteralExpression(c *parser.LiteralExpressionContext) interface{} {
	literal := c.Literal().GetChild(0).(antlr.TerminalNode)
	value := literal.GetText()
	tokenType := literal.GetSymbol().GetTokenType()

	if tokenType == parser.StackLexerINT_LITERAL {
		intValue, _ := strconv.Atoi(value)
		return Variable{parser.StackLexerINT, intValue}
	} else if tokenType == parser.StackLexerFLOAT_LITERAL {
		floatValue, _ := strconv.ParseFloat(value, 64)
		return Variable{parser.StackLexerFLOAT, floatValue}
	} else if tokenType == parser.StackLexerSTRING_LITERAL {
		return Variable{parser.StackLexerSTRING, value}
	} else if tokenType == parser.StackLexerTRUE || tokenType == parser.StackLexerFALSE {
		return Variable{parser.StackLexerBOOL, tokenType == parser.StackLexerTRUE}
	}
	panic(fmt.Sprintf("Unknown token type: %d", tokenType))
}

func (v *Visitor) VisitAssignExpression(c *parser.AssignExpressionContext) interface{} {
	id := c.ID().GetText()
	if _, exists := v.variables[id]; !exists {
		panic(fmt.Sprintf("Variable '%s' does not exist, line: %d", id, c.ID().GetSymbol().GetLine()))
	}
	right := v.Visit(c.Expression()).(Variable)
	if v.variables[id].typeTokenIndex != right.typeTokenIndex {
		panic(fmt.Sprintf("Assigning different type, line: %d", c.GetStart().GetLine()))
	}
	v.variables[id] = right
	return v.variables[id]
}

func eval(input string) {
	is := antlr.NewInputStream(input)
	lexer := parser.NewStackLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewStackParser(stream)

	v := &Visitor{}
	v.Visit(p.Program())
	v.PrintVariables()
}
