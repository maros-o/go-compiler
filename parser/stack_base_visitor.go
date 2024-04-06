// Code generated from ./grammars/Stack.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // Stack

import "github.com/antlr4-go/antlr/v4"

type BaseStackVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseStackVisitor) VisitProgram(ctx *ProgramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStackVisitor) VisitType(ctx *TypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStackVisitor) VisitIntLiteral(ctx *IntLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStackVisitor) VisitFloatLiteral(ctx *FloatLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStackVisitor) VisitTrueLiteral(ctx *TrueLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStackVisitor) VisitFalseLiteral(ctx *FalseLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStackVisitor) VisitStringLiteral(ctx *StringLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStackVisitor) VisitVariableList(ctx *VariableListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStackVisitor) VisitExpressionList(ctx *ExpressionListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStackVisitor) VisitEmptySemStatement(ctx *EmptySemStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStackVisitor) VisitVariableStatement(ctx *VariableStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStackVisitor) VisitExpressionStatement(ctx *ExpressionStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStackVisitor) VisitReadStatement(ctx *ReadStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStackVisitor) VisitWriteStatementBBN(ctx *WriteStatementBBNContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStackVisitor) VisitBlockStatement(ctx *BlockStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStackVisitor) VisitIfStatement(ctx *IfStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStackVisitor) VisitWhileStatement(ctx *WhileStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStackVisitor) VisitMulDivModExpression(ctx *MulDivModExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStackVisitor) VisitIdExpression(ctx *IdExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStackVisitor) VisitAddSubExpression(ctx *AddSubExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStackVisitor) VisitAssignExpression(ctx *AssignExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStackVisitor) VisitComparisonExpression(ctx *ComparisonExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStackVisitor) VisitParenExpression(ctx *ParenExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStackVisitor) VisitLiteralExpression(ctx *LiteralExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStackVisitor) VisitUnaryExpression(ctx *UnaryExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStackVisitor) VisitLogicalExpression(ctx *LogicalExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStackVisitor) VisitStringConcatExpression(ctx *StringConcatExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}
