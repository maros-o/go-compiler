// Code generated from ./grammars/Stack.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // Stack

import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by StackParser.
type StackVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by StackParser#program.
	VisitProgram(ctx *ProgramContext) interface{}

	// Visit a parse tree produced by StackParser#type.
	VisitType(ctx *TypeContext) interface{}

	// Visit a parse tree produced by StackParser#intLiteral.
	VisitIntLiteral(ctx *IntLiteralContext) interface{}

	// Visit a parse tree produced by StackParser#floatLiteral.
	VisitFloatLiteral(ctx *FloatLiteralContext) interface{}

	// Visit a parse tree produced by StackParser#trueLiteral.
	VisitTrueLiteral(ctx *TrueLiteralContext) interface{}

	// Visit a parse tree produced by StackParser#falseLiteral.
	VisitFalseLiteral(ctx *FalseLiteralContext) interface{}

	// Visit a parse tree produced by StackParser#stringLiteral.
	VisitStringLiteral(ctx *StringLiteralContext) interface{}

	// Visit a parse tree produced by StackParser#variableList.
	VisitVariableList(ctx *VariableListContext) interface{}

	// Visit a parse tree produced by StackParser#expressionList.
	VisitExpressionList(ctx *ExpressionListContext) interface{}

	// Visit a parse tree produced by StackParser#emptySemStatement.
	VisitEmptySemStatement(ctx *EmptySemStatementContext) interface{}

	// Visit a parse tree produced by StackParser#variableStatement.
	VisitVariableStatement(ctx *VariableStatementContext) interface{}

	// Visit a parse tree produced by StackParser#expressionStatement.
	VisitExpressionStatement(ctx *ExpressionStatementContext) interface{}

	// Visit a parse tree produced by StackParser#readStatement.
	VisitReadStatement(ctx *ReadStatementContext) interface{}

	// Visit a parse tree produced by StackParser#writeStatement.
	VisitWriteStatement(ctx *WriteStatementContext) interface{}

	// Visit a parse tree produced by StackParser#blockStatement.
	VisitBlockStatement(ctx *BlockStatementContext) interface{}

	// Visit a parse tree produced by StackParser#ifStatement.
	VisitIfStatement(ctx *IfStatementContext) interface{}

	// Visit a parse tree produced by StackParser#whileStatement.
	VisitWhileStatement(ctx *WhileStatementContext) interface{}

	// Visit a parse tree produced by StackParser#mulDivModExpression.
	VisitMulDivModExpression(ctx *MulDivModExpressionContext) interface{}

	// Visit a parse tree produced by StackParser#idExpression.
	VisitIdExpression(ctx *IdExpressionContext) interface{}

	// Visit a parse tree produced by StackParser#addSubExpression.
	VisitAddSubExpression(ctx *AddSubExpressionContext) interface{}

	// Visit a parse tree produced by StackParser#assignExpression.
	VisitAssignExpression(ctx *AssignExpressionContext) interface{}

	// Visit a parse tree produced by StackParser#comparisonExpression.
	VisitComparisonExpression(ctx *ComparisonExpressionContext) interface{}

	// Visit a parse tree produced by StackParser#parenExpression.
	VisitParenExpression(ctx *ParenExpressionContext) interface{}

	// Visit a parse tree produced by StackParser#unaryExpression.
	VisitUnaryExpression(ctx *UnaryExpressionContext) interface{}

	// Visit a parse tree produced by StackParser#literalExpression.
	VisitLiteralExpression(ctx *LiteralExpressionContext) interface{}

	// Visit a parse tree produced by StackParser#logicalExpression.
	VisitLogicalExpression(ctx *LogicalExpressionContext) interface{}

	// Visit a parse tree produced by StackParser#stringConcatExpression.
	VisitStringConcatExpression(ctx *StringConcatExpressionContext) interface{}
}
