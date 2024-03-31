// Code generated from ./grammars/Stack.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // Stack

import "github.com/antlr4-go/antlr/v4"

// StackListener is a complete listener for a parse tree produced by StackParser.
type StackListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterType is called when entering the type production.
	EnterType(c *TypeContext)

	// EnterVariableList is called when entering the variableList production.
	EnterVariableList(c *VariableListContext)

	// EnterIntLiteral is called when entering the intLiteral production.
	EnterIntLiteral(c *IntLiteralContext)

	// EnterFloatLiteral is called when entering the floatLiteral production.
	EnterFloatLiteral(c *FloatLiteralContext)

	// EnterTrueLiteral is called when entering the trueLiteral production.
	EnterTrueLiteral(c *TrueLiteralContext)

	// EnterFalseLiteral is called when entering the falseLiteral production.
	EnterFalseLiteral(c *FalseLiteralContext)

	// EnterStringLiteral is called when entering the stringLiteral production.
	EnterStringLiteral(c *StringLiteralContext)

	// EnterEmptySemStatement is called when entering the emptySemStatement production.
	EnterEmptySemStatement(c *EmptySemStatementContext)

	// EnterVariableStatement is called when entering the variableStatement production.
	EnterVariableStatement(c *VariableStatementContext)

	// EnterExpressionStatement is called when entering the expressionStatement production.
	EnterExpressionStatement(c *ExpressionStatementContext)

	// EnterReadStatement is called when entering the readStatement production.
	EnterReadStatement(c *ReadStatementContext)

	// EnterWriteStatement is called when entering the writeStatement production.
	EnterWriteStatement(c *WriteStatementContext)

	// EnterBlockStatement is called when entering the blockStatement production.
	EnterBlockStatement(c *BlockStatementContext)

	// EnterIfStatement is called when entering the ifStatement production.
	EnterIfStatement(c *IfStatementContext)

	// EnterWhileStatement is called when entering the whileStatement production.
	EnterWhileStatement(c *WhileStatementContext)

	// EnterMulDivModExpression is called when entering the mulDivModExpression production.
	EnterMulDivModExpression(c *MulDivModExpressionContext)

	// EnterIdExpression is called when entering the idExpression production.
	EnterIdExpression(c *IdExpressionContext)

	// EnterAddSubExpression is called when entering the addSubExpression production.
	EnterAddSubExpression(c *AddSubExpressionContext)

	// EnterAssignExpression is called when entering the assignExpression production.
	EnterAssignExpression(c *AssignExpressionContext)

	// EnterComparisonExpression is called when entering the comparisonExpression production.
	EnterComparisonExpression(c *ComparisonExpressionContext)

	// EnterParenExpression is called when entering the parenExpression production.
	EnterParenExpression(c *ParenExpressionContext)

	// EnterStringConcatenation is called when entering the stringConcatenation production.
	EnterStringConcatenation(c *StringConcatenationContext)

	// EnterLiteralExpression is called when entering the literalExpression production.
	EnterLiteralExpression(c *LiteralExpressionContext)

	// EnterUnaryExpression is called when entering the unaryExpression production.
	EnterUnaryExpression(c *UnaryExpressionContext)

	// EnterLogicalExpression is called when entering the logicalExpression production.
	EnterLogicalExpression(c *LogicalExpressionContext)

	// EnterExpressionList is called when entering the expressionList production.
	EnterExpressionList(c *ExpressionListContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitType is called when exiting the type production.
	ExitType(c *TypeContext)

	// ExitVariableList is called when exiting the variableList production.
	ExitVariableList(c *VariableListContext)

	// ExitIntLiteral is called when exiting the intLiteral production.
	ExitIntLiteral(c *IntLiteralContext)

	// ExitFloatLiteral is called when exiting the floatLiteral production.
	ExitFloatLiteral(c *FloatLiteralContext)

	// ExitTrueLiteral is called when exiting the trueLiteral production.
	ExitTrueLiteral(c *TrueLiteralContext)

	// ExitFalseLiteral is called when exiting the falseLiteral production.
	ExitFalseLiteral(c *FalseLiteralContext)

	// ExitStringLiteral is called when exiting the stringLiteral production.
	ExitStringLiteral(c *StringLiteralContext)

	// ExitEmptySemStatement is called when exiting the emptySemStatement production.
	ExitEmptySemStatement(c *EmptySemStatementContext)

	// ExitVariableStatement is called when exiting the variableStatement production.
	ExitVariableStatement(c *VariableStatementContext)

	// ExitExpressionStatement is called when exiting the expressionStatement production.
	ExitExpressionStatement(c *ExpressionStatementContext)

	// ExitReadStatement is called when exiting the readStatement production.
	ExitReadStatement(c *ReadStatementContext)

	// ExitWriteStatement is called when exiting the writeStatement production.
	ExitWriteStatement(c *WriteStatementContext)

	// ExitBlockStatement is called when exiting the blockStatement production.
	ExitBlockStatement(c *BlockStatementContext)

	// ExitIfStatement is called when exiting the ifStatement production.
	ExitIfStatement(c *IfStatementContext)

	// ExitWhileStatement is called when exiting the whileStatement production.
	ExitWhileStatement(c *WhileStatementContext)

	// ExitMulDivModExpression is called when exiting the mulDivModExpression production.
	ExitMulDivModExpression(c *MulDivModExpressionContext)

	// ExitIdExpression is called when exiting the idExpression production.
	ExitIdExpression(c *IdExpressionContext)

	// ExitAddSubExpression is called when exiting the addSubExpression production.
	ExitAddSubExpression(c *AddSubExpressionContext)

	// ExitAssignExpression is called when exiting the assignExpression production.
	ExitAssignExpression(c *AssignExpressionContext)

	// ExitComparisonExpression is called when exiting the comparisonExpression production.
	ExitComparisonExpression(c *ComparisonExpressionContext)

	// ExitParenExpression is called when exiting the parenExpression production.
	ExitParenExpression(c *ParenExpressionContext)

	// ExitStringConcatenation is called when exiting the stringConcatenation production.
	ExitStringConcatenation(c *StringConcatenationContext)

	// ExitLiteralExpression is called when exiting the literalExpression production.
	ExitLiteralExpression(c *LiteralExpressionContext)

	// ExitUnaryExpression is called when exiting the unaryExpression production.
	ExitUnaryExpression(c *UnaryExpressionContext)

	// ExitLogicalExpression is called when exiting the logicalExpression production.
	ExitLogicalExpression(c *LogicalExpressionContext)

	// ExitExpressionList is called when exiting the expressionList production.
	ExitExpressionList(c *ExpressionListContext)
}
