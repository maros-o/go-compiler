package main

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"github.com/fatih/color"
	"go-compiler/parser"
)

var tokenTypeToTypeStr = map[int]string{
	parser.StackLexerINT:    "I",
	parser.StackLexerFLOAT:  "F",
	parser.StackLexerSTRING: "S",
	parser.StackLexerBOOL:   "B",
}

var tokenTypeToDefaultValue = map[int]string{
	parser.StackLexerINT:    "0",
	parser.StackLexerFLOAT:  "0.0",
	parser.StackLexerSTRING: "\"\"",
	parser.StackLexerBOOL:   "true",
}

var tokenLiteralToType = map[int]int{
	parser.StackLexerINT_LITERAL:    parser.StackLexerINT,
	parser.StackLexerFLOAT_LITERAL:  parser.StackLexerFLOAT,
	parser.StackLexerSTRING_LITERAL: parser.StackLexerSTRING,
	parser.StackLexerBOOL_LITERAL:   parser.StackLexerBOOL,
}

var opToInstruction = map[int]string{
	parser.StackLexerADD:  "add",
	parser.StackLexerSUB:  "sub",
	parser.StackLexerMUL:  "mul",
	parser.StackLexerDIV:  "div",
	parser.StackLexerMOD:  "mod",
	parser.StackLexerAND:  "and",
	parser.StackLexerOR:   "or",
	parser.StackLexerGT:   "gt",
	parser.StackLexerLT:   "lt",
	parser.StackLexerEQ:   "eq",
	parser.StackParserNE:  "ne",
	parser.StackLexerNOT:  "not",
	parser.StackParserDOT: "concat",
}

type Visitor struct {
	*parser.BaseStackVisitor
	variables    map[string]Variable
	instructions []string
	depth        int
	labelCount   int
}

func (v *Visitor) NewVariable(typeTokenIndex int, id ...string) Variable {
	if len(id) == 0 {
		return Variable{typeTokenIndex: typeTokenIndex, depth: v.depth}
	}
	return Variable{typeTokenIndex: typeTokenIndex, depth: v.depth, id: id[0]}
}

func (v *Visitor) Visit(t antlr.ParseTree) interface{} {
	return t.Accept(v)
}

func (v *Visitor) VisitEmptySemStatement(_ *parser.EmptySemStatementContext) interface{} {
	return nil
}

func (v *Visitor) VisitProgram(c *parser.ProgramContext) interface{} {
	v.variables = make(map[string]Variable)
	v.instructions = make([]string, 0)
	v.depth = 0

	for _, ch := range c.GetChildren() {
		v.Visit(ch.(antlr.ParseTree))
	}
	return nil
}

func (v *Visitor) VisitVariableStatement(c *parser.VariableStatementContext) interface{} {
	typeTokenIndex := c.Type_().GetChild(0).(antlr.TerminalNode).GetSymbol().GetTokenType()
	ids := v.VisitVariableList(c.VariableList().(*parser.VariableListContext)).([]string)

	for _, id := range ids {
		if _, exists := v.variables[id]; exists {
			color.Red("VisitVariableStatement::Variable %s already exists, line: %d \n", id, c.GetStart().GetLine())
			break
		}
		v.variables[id] = v.NewVariable(typeTokenIndex, id)
		v.instructions = append(v.instructions, fmt.Sprintf("push %s %s", tokenTypeToTypeStr[typeTokenIndex], tokenTypeToDefaultValue[typeTokenIndex]))
		v.instructions = append(v.instructions, fmt.Sprintf("save %s", id))
	}
	return nil
}

func (v *Visitor) VisitExpressionStatement(c *parser.ExpressionStatementContext) interface{} {
	v.Visit(c.Expression())
	return nil
}

func (v *Visitor) VisitReadStatement(c *parser.ReadStatementContext) interface{} {
	ids := v.VisitVariableList(c.VariableList().(*parser.VariableListContext)).([]string)

	for _, id := range ids {
		variable, exists := v.variables[id]
		if !exists {
			color.Red("VisitReadStatement::Variable %s does not exist, line: %d\n", id, c.GetStart().GetLine())
			return nil
		}
		v.instructions = append(v.instructions, fmt.Sprintf("read %s", tokenTypeToTypeStr[variable.typeTokenIndex]))
		if variable.id != "" {
			v.instructions = append(v.instructions, fmt.Sprintf("save %s", id))
		}
	}
	return nil
}

func (v *Visitor) VisitWriteStatement(c *parser.WriteStatementContext) interface{} {
	n := 0
	for _, ch := range c.ExpressionList().GetChildren() {
		t, ok := ch.(antlr.TerminalNode)
		if ok && t.GetSymbol().GetTokenType() == parser.StackLexerCOMMA {
			continue
		}
		v.Visit(ch.(antlr.ParseTree))
		n++
	}
	v.instructions = append(v.instructions, fmt.Sprintf("print %d", n))
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

func (v *Visitor) VisitIfStatement(c *parser.IfStatementContext) interface{} {
	condition := v.Visit(c.Expression()).(Variable)
	if condition.typeTokenIndex != parser.StackLexerBOOL {
		color.Red("VisitIfStatement::If condition is not boolean, line: %d\n", c.GetStart().GetLine())
		return nil
	}
	v.instructions = append(v.instructions, fmt.Sprintf("fjmp %d", v.labelCount))
	v.Visit(c.Statement(0))
	v.instructions = append(v.instructions, fmt.Sprintf("jmp %d", v.labelCount+1))
	v.instructions = append(v.instructions, fmt.Sprintf("label %d", v.labelCount))
	v.labelCount++
	if c.Statement(1) != nil {
		v.Visit(c.Statement(1))
	}
	v.instructions = append(v.instructions, fmt.Sprintf("label %d", v.labelCount))
	v.labelCount++
	return nil
}

func (v *Visitor) VisitWhileStatement(c *parser.WhileStatementContext) interface{} {
	v.instructions = append(v.instructions, fmt.Sprintf("label %d", v.labelCount))
	v.labelCount++
	condition := v.Visit(c.Expression()).(Variable)
	if condition.typeTokenIndex != parser.StackLexerBOOL {
		color.Red("VisitWhileStatement::While condition is not boolean, line: %d\n", c.GetStart().GetLine())
		return nil
	}
	v.instructions = append(v.instructions, fmt.Sprintf("fjmp %d", v.labelCount))
	v.Visit(c.Statement())
	v.instructions = append(v.instructions, fmt.Sprintf("jmp %d", v.labelCount-1))
	v.instructions = append(v.instructions, fmt.Sprintf("label %d", v.labelCount))
	v.labelCount++
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

func (v *Visitor) VisitParenExpression(c *parser.ParenExpressionContext) interface{} {
	return v.Visit(c.Expression())
}

func (v *Visitor) VisitUminusExpression(c *parser.UminusExpressionContext) interface{} {
	right, ok := v.Visit(c.Expression()).(Variable)
	if !ok {
		return nil
	}
	if right.typeTokenIndex != parser.StackLexerINT && right.typeTokenIndex != parser.StackLexerFLOAT {
		color.Red("VisitUminusExpression::Uminus non-numeric value, line: %d\n", c.GetStart().GetLine())
		return nil
	}
	v.instructions = append(v.instructions, "uminus")
	return right
}

func (v *Visitor) VisitNotExpression(c *parser.NotExpressionContext) interface{} {
	right, ok := v.Visit(c.Expression()).(Variable)
	if !ok {
		return nil
	}
	if right.typeTokenIndex != parser.StackLexerBOOL {
		color.Red("VisitNotExpression::Negating non-boolean value, line: %d\n", c.GetStart().GetLine())
		return nil
	}
	v.instructions = append(v.instructions, fmt.Sprintf("%s", opToInstruction[c.GetOp().GetTokenType()]))
	return right
}

func (v *Visitor) VisitMulDivModExpression(c *parser.MulDivModExpressionContext) interface{} {
	left, ok := v.Visit(c.Expression(0)).(Variable)
	if !ok {
		return nil
	}
	right, ok := v.Visit(c.Expression(1)).(Variable)
	if !ok {
		return nil
	}
	op := c.GetOp().GetTokenType()

	if op == parser.StackLexerMOD && (left.typeTokenIndex != parser.StackLexerINT || right.typeTokenIndex != parser.StackLexerINT) {
		color.Red("VisitMulDivModExpression::Modulo non-integer value, line: %d\n", c.GetStart().GetLine())
		return nil
	}
	if left.typeTokenIndex == parser.StackLexerINT && right.typeTokenIndex == parser.StackLexerFLOAT {
		pop := v.instructions[len(v.instructions)-1]
		v.instructions = v.instructions[:len(v.instructions)-1]
		v.instructions = append(v.instructions, "itof")
		v.instructions = append(v.instructions, pop)
		left.typeTokenIndex = parser.StackLexerFLOAT
	}
	if left.typeTokenIndex == parser.StackLexerFLOAT && right.typeTokenIndex == parser.StackLexerINT {
		v.instructions = append(v.instructions, "itof")
	}
	v.instructions = append(v.instructions, fmt.Sprintf("%s", opToInstruction[c.GetOp().GetTokenType()]))
	return left
}

func (v *Visitor) VisitAddSubExpression(c *parser.AddSubExpressionContext) interface{} {
	left, ok := v.Visit(c.Expression(0)).(Variable)
	if !ok {
		return nil
	}
	right, ok := v.Visit(c.Expression(1)).(Variable)
	if !ok {
		return nil
	}

	if left.typeTokenIndex == parser.StackLexerINT && right.typeTokenIndex == parser.StackLexerFLOAT {
		pop := v.instructions[len(v.instructions)-1]
		v.instructions = v.instructions[:len(v.instructions)-1]
		v.instructions = append(v.instructions, "itof")
		v.instructions = append(v.instructions, pop)
		left.typeTokenIndex = parser.StackLexerFLOAT
	}
	if left.typeTokenIndex == parser.StackLexerFLOAT && right.typeTokenIndex == parser.StackLexerINT {
		v.instructions = append(v.instructions, "itof")
		right.typeTokenIndex = parser.StackLexerFLOAT
	}
	if left.typeTokenIndex != right.typeTokenIndex || left.typeTokenIndex == parser.StackLexerSTRING || left.typeTokenIndex == parser.StackLexerBOOL {
		color.Red("VisitAddSubExpression::Tried to %s %s %s, line: %d\n", tokenTypeToTypeStr[left.typeTokenIndex], c.GetOp().GetText(), tokenTypeToTypeStr[right.typeTokenIndex], c.GetStart().GetLine())
		return nil
	}
	v.instructions = append(v.instructions, fmt.Sprintf("%s", opToInstruction[c.GetOp().GetTokenType()]))
	return left
}

func (v *Visitor) VisitDotExpression(c *parser.DotExpressionContext) interface{} {
	left, ok := v.Visit(c.Expression(0)).(Variable)
	if !ok {
		return nil
	}
	right, ok := v.Visit(c.Expression(1)).(Variable)
	if !ok {
		return nil
	}

	if left.typeTokenIndex != right.typeTokenIndex || left.typeTokenIndex != parser.StackLexerSTRING {
		color.Red("VisitDotExpression::Tried to %s %s %s, line: %d\n", tokenTypeToTypeStr[left.typeTokenIndex], c.GetOp().GetText(), tokenTypeToTypeStr[right.typeTokenIndex], c.GetStart().GetLine())
		return nil
	}
	v.instructions = append(v.instructions, fmt.Sprintf("%s", opToInstruction[c.GetOp().GetTokenType()]))
	return left
}

func (v *Visitor) VisitComparisonExpression(c *parser.ComparisonExpressionContext) interface{} {
	left, ok := v.Visit(c.Expression(0)).(Variable)
	if !ok {
		return nil
	}
	right, ok := v.Visit(c.Expression(1)).(Variable)
	if !ok {
		return nil
	}

	if left.typeTokenIndex == parser.StackLexerINT && right.typeTokenIndex == parser.StackLexerFLOAT {
		pop := v.instructions[len(v.instructions)-1]
		v.instructions = v.instructions[:len(v.instructions)-1]
		v.instructions = append(v.instructions, "itof")
		v.instructions = append(v.instructions, pop)
		left.typeTokenIndex = parser.StackLexerFLOAT
	}
	if left.typeTokenIndex == parser.StackLexerFLOAT && right.typeTokenIndex == parser.StackLexerINT {
		v.instructions = append(v.instructions, "itof")
		right.typeTokenIndex = parser.StackLexerFLOAT
	}
	if left.typeTokenIndex != right.typeTokenIndex {
		color.Red("VisitComparisonExpression::Tried to %s %s %s, line: %d\n", tokenTypeToTypeStr[left.typeTokenIndex], c.GetOp().GetText(), tokenTypeToTypeStr[right.typeTokenIndex], c.GetStart().GetLine())
		return nil
	}
	v.instructions = append(v.instructions, fmt.Sprintf("%s", opToInstruction[c.GetOp().GetTokenType()]))
	left.typeTokenIndex = parser.StackLexerBOOL
	return left
}

func (v *Visitor) VisitEqualityExpression(c *parser.EqualityExpressionContext) interface{} {
	left, ok := v.Visit(c.Expression(0)).(Variable)
	if !ok {
		return nil
	}
	right, ok := v.Visit(c.Expression(1)).(Variable)
	if !ok {
		return nil
	}

	if left.typeTokenIndex != right.typeTokenIndex {
		color.Red("VisitEqualityExpression::Tried to %s %s %s, line: %d\n", tokenTypeToTypeStr[left.typeTokenIndex], c.GetOp().GetText(), tokenTypeToTypeStr[right.typeTokenIndex], c.GetStart().GetLine())
		return nil
	}
	op := c.GetOp().GetTokenType()
	if op == parser.StackLexerNE {
		v.instructions = append(v.instructions, "eq")
		v.instructions = append(v.instructions, "not")
	} else {
		v.instructions = append(v.instructions, fmt.Sprintf("%s", opToInstruction[op]))
	}
	left.typeTokenIndex = parser.StackLexerBOOL
	return left
}

func (v *Visitor) VisitAndExpression(c *parser.AndExpressionContext) interface{} {
	left, ok := v.Visit(c.Expression(0)).(Variable)
	if !ok {
		return nil
	}
	right, ok := v.Visit(c.Expression(1)).(Variable)
	if !ok {
		return nil
	}

	if left.typeTokenIndex != right.typeTokenIndex || left.typeTokenIndex != parser.StackLexerBOOL {
		color.Red("VisitAndExpression::Tried to %s %s %s, line: %d\n", tokenTypeToTypeStr[left.typeTokenIndex], c.GetOp().GetText(), tokenTypeToTypeStr[right.typeTokenIndex], c.GetStart().GetLine())
		return nil
	}
	v.instructions = append(v.instructions, fmt.Sprintf("%s", opToInstruction[c.GetOp().GetTokenType()]))
	return left
}

func (v *Visitor) VisitOrExpression(c *parser.OrExpressionContext) interface{} {
	left, ok := v.Visit(c.Expression(0)).(Variable)
	if !ok {
		return nil
	}
	right, ok := v.Visit(c.Expression(1)).(Variable)
	if !ok {
		return nil
	}

	if left.typeTokenIndex != right.typeTokenIndex || left.typeTokenIndex != parser.StackLexerBOOL {
		color.Red("VisitOrExpression::Tried to %s %s %s, line: %d\n", tokenTypeToTypeStr[left.typeTokenIndex], c.GetOp().GetText(), tokenTypeToTypeStr[right.typeTokenIndex], c.GetStart().GetLine())
		return nil
	}
	v.instructions = append(v.instructions, fmt.Sprintf("%s", opToInstruction[c.GetOp().GetTokenType()]))
	return left
}

func (v *Visitor) VisitAssignExpression(c *parser.AssignExpressionContext) interface{} {
	id := c.ID().GetText()
	variable, exists := v.variables[id]
	if !exists {
		color.Red("VisitAssignExpression::Variable %s does not exist, line: %d\n", id, c.ID().GetSymbol().GetLine())
		return nil
	}
	right, ok := v.Visit(c.Expression()).(Variable)
	if !ok {
		return nil
	}

	_, rightIsAssign := c.Expression().(*parser.AssignExpressionContext)
	if right.id != "" && rightIsAssign {
		v.instructions = v.instructions[:len(v.instructions)-2]
		v.instructions = append(v.instructions, fmt.Sprintf("load %s", right.id))
	}
	if variable.typeTokenIndex == parser.StackLexerFLOAT && right.typeTokenIndex == parser.StackLexerINT {
		v.instructions = append(v.instructions, "itof")
		right.typeTokenIndex = parser.StackLexerFLOAT
	}
	if variable.typeTokenIndex != right.typeTokenIndex {
		color.Red("VisitAssignExpression::Assigning type %s to variable %s of type %s, line: %d\n", tokenTypeToTypeStr[right.typeTokenIndex], id, tokenTypeToTypeStr[variable.typeTokenIndex], c.GetStart().GetLine())
		return nil
	}
	v.instructions = append(v.instructions, fmt.Sprintf("save %s", id))
	v.instructions = append(v.instructions, fmt.Sprintf("load %s", id))
	v.instructions = append(v.instructions, fmt.Sprintf("pop"))
	return variable
}

func (v *Visitor) VisitLiteralExpression(c *parser.LiteralExpressionContext) interface{} {
	literal := c.Literal().GetChild(0).(antlr.TerminalNode)
	value := literal.GetText()
	tokenType := tokenLiteralToType[literal.GetSymbol().GetTokenType()]
	v.instructions = append(v.instructions, fmt.Sprintf("push %s %s", tokenTypeToTypeStr[tokenType], value))
	return v.NewVariable(tokenType)
}

func (v *Visitor) VisitIdExpression(c *parser.IdExpressionContext) interface{} {
	id := c.ID().GetText()
	if _, exists := v.variables[id]; !exists {
		color.Red("VisitIdExpression::Variable %s does not exist, line: %d \n", id, c.ID().GetSymbol().GetLine())
	}
	v.instructions = append(v.instructions, fmt.Sprintf("load %s", id))
	return v.variables[id]
}
