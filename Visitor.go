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
	depth     int
}

func (v *Visitor) PrintVariables() {
	fmt.Println("Variables:")
	fmt.Printf("%-10s %-10s %-10s \n", "ID", "Type", "Value")
	for name, variable := range v.variables {
		fmt.Printf("%-10s %-10s %-10s \n", name, variable.GetTypeString(), variable.GetValueString())
	}
}

func (v *Visitor) NewVariable(typeTokenIndex int, value interface{}) Variable {
	return Variable{typeTokenIndex: typeTokenIndex, value: value, depth: v.depth}
}

func (v *Visitor) Visit(t antlr.ParseTree) interface{} {
	return t.Accept(v)
}

func (v *Visitor) VisitProgram(c *parser.ProgramContext) interface{} {
	v.variables = make(map[string]Variable)
	for _, ch := range c.GetChildren() {
		v.Visit(ch.(antlr.ParseTree))
	}
	return nil
}

func (v *Visitor) VisitVariableStatement(c *parser.VariableStatementContext) interface{} {
	typeTokenIndex := c.Type_().GetChild(0).(antlr.TerminalNode).GetSymbol().GetTokenType()
	ids := v.VisitVariableList(c.VariableList().(*parser.VariableListContext)).([]string)

	for _, name := range ids {
		if _, exists := v.variables[name]; exists {
			panic(fmt.Sprintf("Variable '%s' already exists, line: %d", name, c.GetStart().GetLine()))
		}

		if typeTokenIndex == parser.StackLexerINT {
			v.variables[name] = v.NewVariable(typeTokenIndex, 0)
		} else if typeTokenIndex == parser.StackLexerFLOAT {
			v.variables[name] = v.NewVariable(typeTokenIndex, 0.0)
		} else if typeTokenIndex == parser.StackLexerSTRING {
			v.variables[name] = v.NewVariable(typeTokenIndex, "")
		} else if typeTokenIndex == parser.StackLexerBOOL {
			v.variables[name] = v.NewVariable(typeTokenIndex, false)
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
			return v.NewVariable(parser.StackLexerINT, left.value.(int)+right.value.(int))
		} else if c.GetOp().GetTokenType() == parser.StackLexerSUB {
			return v.NewVariable(parser.StackLexerINT, left.value.(int)-right.value.(int))
		}
	} else if left.typeTokenIndex == parser.StackLexerFLOAT {
		if c.GetOp().GetTokenType() == parser.StackLexerADD {
			return v.NewVariable(parser.StackLexerFLOAT, left.value.(float64)+right.value.(float64))
		} else if c.GetOp().GetTokenType() == parser.StackLexerSUB {
			return v.NewVariable(parser.StackLexerFLOAT, left.value.(float64)-right.value.(float64))
		}
	}
	panic(fmt.Sprintf("Unknown type: %d", left.typeTokenIndex))
}

func (v *Visitor) VisitMulDivModExpression(c *parser.MulDivModExpressionContext) interface{} {
	left := v.Visit(c.Expression(0)).(Variable)
	right := v.Visit(c.Expression(1)).(Variable)

	if left.typeTokenIndex != right.typeTokenIndex {
		panic(fmt.Sprintf("MulDivMod different types of variables, line: %d", c.GetStart().GetLine()))
	}

	if left.typeTokenIndex == parser.StackLexerINT {
		if c.GetOp().GetTokenType() == parser.StackLexerMUL {
			return v.NewVariable(parser.StackLexerINT, left.value.(int)*right.value.(int))
		} else if c.GetOp().GetTokenType() == parser.StackLexerDIV {
			return v.NewVariable(parser.StackLexerINT, left.value.(int)/right.value.(int))
		} else if c.GetOp().GetTokenType() == parser.StackLexerMOD {
			return v.NewVariable(parser.StackLexerINT, left.value.(int)%right.value.(int))
		}
	} else if left.typeTokenIndex == parser.StackLexerFLOAT {
		if c.GetOp().GetTokenType() == parser.StackLexerMUL {
			return v.NewVariable(parser.StackLexerFLOAT, left.value.(float64)*right.value.(float64))
		} else if c.GetOp().GetTokenType() == parser.StackLexerDIV {
			return v.NewVariable(parser.StackLexerFLOAT, left.value.(float64)/right.value.(float64))
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
		return v.NewVariable(parser.StackLexerINT, intValue)
	} else if tokenType == parser.StackLexerFLOAT_LITERAL {
		floatValue, _ := strconv.ParseFloat(value, 64)
		return v.NewVariable(parser.StackLexerFLOAT, floatValue)
	} else if tokenType == parser.StackLexerSTRING_LITERAL {
		return v.VisitStringLiteral(c.Literal().(*parser.StringLiteralContext))
	} else if tokenType == parser.StackLexerTRUE || tokenType == parser.StackLexerFALSE {
		return v.NewVariable(parser.StackLexerBOOL, tokenType == parser.StackLexerTRUE)
	}
	panic(fmt.Sprintf("Unknown token type: %d", tokenType))
}

func (v *Visitor) VisitAssignExpression(c *parser.AssignExpressionContext) interface{} {
	id := c.ID().GetText()
	variable, exists := v.variables[id]
	if !exists {
		panic(fmt.Sprintf("Variable '%s' does not exist, line: %d", id, c.ID().GetSymbol().GetLine()))
	}
	right := v.Visit(c.Expression()).(Variable)
	if v.variables[id].typeTokenIndex != right.typeTokenIndex {
		panic(fmt.Sprintf("Assigning different type, line: %d", c.GetStart().GetLine()))
	}
	v.variables[id] = Variable{typeTokenIndex: variable.typeTokenIndex, value: right.value, depth: variable.depth}
	return v.variables[id]
}

func (v *Visitor) VisitWriteStatement(c *parser.WriteStatementContext) interface{} {
	for _, ch := range c.ExpressionList().GetChildren() {
		t, ok := ch.(antlr.TerminalNode)
		if ok && t.GetSymbol().GetTokenType() == parser.StackLexerCOMMA {
			continue
		}
		res := v.Visit(ch.(antlr.ParseTree)).(Variable)
		fmt.Print(res.GetValueString())
	}
	fmt.Println("")
	return nil
}

func (v *Visitor) VisitEmptySemStatement(_ *parser.EmptySemStatementContext) interface{} {
	return nil
}

func (v *Visitor) VisitUnaryExpression(c *parser.UnaryExpressionContext) interface{} {
	right := v.Visit(c.Expression()).(Variable)
	op := c.GetOp().GetTokenType()

	if op == parser.StackLexerNOT && right.typeTokenIndex == parser.StackLexerBOOL {
		return v.NewVariable(parser.StackLexerBOOL, !right.value.(bool))
	}
	if right.typeTokenIndex == parser.StackLexerINT {
		return v.NewVariable(parser.StackLexerINT, -right.value.(int))
	} else if right.typeTokenIndex == parser.StackLexerFLOAT {
		return v.NewVariable(parser.StackLexerFLOAT, -right.value.(float64))
	}
	panic(fmt.Sprintf("Unknown type: %d", right.typeTokenIndex))
}

func (v *Visitor) VisitStringLiteral(c *parser.StringLiteralContext) interface{} {
	text := c.STRING_LITERAL().GetText()
	return v.NewVariable(parser.StackLexerSTRING, text[1:len(text)-1])
}

func (v *Visitor) VisitStringConcatExpression(c *parser.StringConcatExpressionContext) interface{} {
	left := c.STRING_LITERAL(0).GetSymbol().GetText()
	right := c.STRING_LITERAL(1).GetSymbol().GetText()
	return v.NewVariable(parser.StackLexerSTRING, left[1:len(left)-1]+right[1:len(right)-1])
}

func (v *Visitor) VisitReadStatement(c *parser.ReadStatementContext) interface{} {
	ids := v.VisitVariableList(c.VariableList().(*parser.VariableListContext)).([]string)
	var input string

	for _, name := range ids {
		fmt.Printf("Enter value for %s\n", name)
		_, err := fmt.Scan(&input)
		if err != nil {
			panic(fmt.Sprintf("Error reading input: %s", err))
		}
		tokenType := v.variables[name].typeTokenIndex
		if tokenType == parser.StackLexerINT {
			intValue, err := strconv.Atoi(input)
			if err != nil {
				panic(fmt.Sprintf("Wrong user input, expected int, got %s", input))
			}
			v.variables[name] = v.NewVariable(tokenType, intValue)
		} else if tokenType == parser.StackLexerFLOAT {
			floatValue, err := strconv.ParseFloat(input, 64)
			if err != nil {
				panic(fmt.Sprintf("Wrong user input, expected float, got %s", input))
			}
			v.variables[name] = v.NewVariable(tokenType, floatValue)
		} else if tokenType == parser.StackLexerSTRING {
			v.variables[name] = v.NewVariable(tokenType, input)
		} else if tokenType == parser.StackLexerBOOL {
			boolValue, err := strconv.ParseBool(input)
			if err != nil {
				panic(fmt.Sprintf("Wrong user input, expected bool, got %s", input))
			}
			v.variables[name] = v.NewVariable(tokenType, boolValue)
		} else {
			panic(fmt.Sprintf("Unknown type: %d", tokenType))
		}
	}
	return nil
}

func (v *Visitor) VisitVariableList(c *parser.VariableListContext) interface{} {
	res := make([]string, 0)

	for _, ch := range c.GetChildren() {
		id := ch.(antlr.TerminalNode)
		if id.GetSymbol().GetTokenType() != parser.StackLexerID {
			continue
		}
		name := id.GetText()
		res = append(res, name)
	}
	return res
}

func (v *Visitor) VisitComparisonExpression(c *parser.ComparisonExpressionContext) interface{} {
	left := v.Visit(c.Expression(0)).(Variable)
	right := v.Visit(c.Expression(1)).(Variable)

	if left.typeTokenIndex != right.typeTokenIndex {
		panic(fmt.Sprintf("Comparison different types of variables, line: %d", c.GetStart().GetLine()))
	}

	op := c.GetOp().GetTokenType()
	if left.typeTokenIndex == parser.StackLexerINT {
		if op == parser.StackLexerEQ {
			return v.NewVariable(parser.StackLexerBOOL, left.value.(int) == right.value.(int))
		} else if op == parser.StackLexerNE {
			return v.NewVariable(parser.StackLexerBOOL, left.value.(int) != right.value.(int))
		} else if op == parser.StackLexerLT {
			return v.NewVariable(parser.StackLexerBOOL, left.value.(int) < right.value.(int))
		} else if op == parser.StackParserGT {
			return v.NewVariable(parser.StackLexerBOOL, left.value.(int) > right.value.(int))
		}
	} else if left.typeTokenIndex == parser.StackLexerFLOAT {
		if op == parser.StackLexerEQ {
			return v.NewVariable(parser.StackLexerBOOL, left.value.(float64) == right.value.(float64))
		} else if op == parser.StackLexerNE {
			return v.NewVariable(parser.StackLexerBOOL, left.value.(float64) != right.value.(float64))
		} else if op == parser.StackLexerLT {
			return v.NewVariable(parser.StackLexerBOOL, left.value.(float64) < right.value.(float64))
		} else if op == parser.StackParserGT {
			return v.NewVariable(parser.StackLexerBOOL, left.value.(float64) > right.value.(float64))
		}
	} else if left.typeTokenIndex == parser.StackLexerSTRING {
		if op == parser.StackLexerEQ {
			return v.NewVariable(parser.StackLexerBOOL, left.value.(string) == right.value.(string))
		} else if op == parser.StackLexerNE {
			return v.NewVariable(parser.StackLexerBOOL, left.value.(string) != right.value.(string))
		} else if op == parser.StackLexerLT {
			return v.NewVariable(parser.StackLexerBOOL, left.value.(string) < right.value.(string))
		} else if op == parser.StackParserGT {
			return v.NewVariable(parser.StackLexerBOOL, left.value.(string) > right.value.(string))
		}
	} else if left.typeTokenIndex == parser.StackLexerBOOL {
		if op == parser.StackLexerEQ {
			return v.NewVariable(parser.StackLexerBOOL, left.value.(bool) == right.value.(bool))
		} else if op == parser.StackLexerNE {
			return v.NewVariable(parser.StackLexerBOOL, left.value.(bool) != right.value.(bool))
		}
	}
	panic(fmt.Sprintf("Unknown type: %d", left.typeTokenIndex))
}

func (v *Visitor) VisitLogicalExpression(c *parser.LogicalExpressionContext) interface{} {
	left := v.Visit(c.Expression(0)).(Variable)
	right := v.Visit(c.Expression(1)).(Variable)

	if left.typeTokenIndex != parser.StackLexerBOOL || right.typeTokenIndex != parser.StackLexerBOOL {
		panic(fmt.Sprintf("Logical expression with non-boolean types, line: %d", c.GetStart().GetLine()))
	}

	op := c.GetOp().GetTokenType()
	if op == parser.StackLexerAND {
		return v.NewVariable(parser.StackLexerBOOL, left.value.(bool) && right.value.(bool))
	} else if op == parser.StackLexerOR {
		return v.NewVariable(parser.StackLexerBOOL, left.value.(bool) || right.value.(bool))
	}
	panic(fmt.Sprintf("Unknown operator: %d", op))
}

func (v *Visitor) VisitParenExpression(c *parser.ParenExpressionContext) interface{} {
	return v.Visit(c.Expression())
}

func (v *Visitor) VisitIfStatement(c *parser.IfStatementContext) interface{} {
	condition := v.Visit(c.Expression()).(Variable)
	if condition.typeTokenIndex != parser.StackLexerBOOL {
		panic(fmt.Sprintf("If condition is not boolean, line: %d", c.GetStart().GetLine()))
	}
	if condition.value.(bool) {
		v.Visit(c.Statement(0))
	} else if c.Statement(1) != nil {
		v.Visit(c.Statement(1))
	}
	return nil
}

func (v *Visitor) VisitBlockStatement(c *parser.BlockStatementContext) interface{} {
	v.depth++
	for _, ch := range c.GetChildren()[1 : len(c.GetChildren())-1] {
		v.Visit(ch.(antlr.ParseTree))
	}
	for name, variable := range v.variables {
		if variable.depth == v.depth {
			delete(v.variables, name)
		}
	}
	v.depth--
	return nil
}

func (v *Visitor) VisitWhileStatement(c *parser.WhileStatementContext) interface{} {
	for {
		condition := v.Visit(c.Expression()).(Variable)
		if condition.typeTokenIndex != parser.StackLexerBOOL {
			panic(fmt.Sprintf("While condition is not boolean, line: %d", c.GetStart().GetLine()))
		}
		if !condition.value.(bool) {
			break
		}
		v.Visit(c.Statement())
	}
	return nil
}
