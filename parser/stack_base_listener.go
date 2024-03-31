// Code generated from ./grammars/Stack.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // Stack

import "github.com/antlr4-go/antlr/v4"

// BaseStackListener is a complete listener for a parse tree produced by StackParser.
type BaseStackListener struct{}

var _ StackListener = &BaseStackListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseStackListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseStackListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseStackListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseStackListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BaseStackListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BaseStackListener) ExitProgram(ctx *ProgramContext) {}

// EnterType is called when production type is entered.
func (s *BaseStackListener) EnterType(ctx *TypeContext) {}

// ExitType is called when production type is exited.
func (s *BaseStackListener) ExitType(ctx *TypeContext) {}

// EnterVariableList is called when production variableList is entered.
func (s *BaseStackListener) EnterVariableList(ctx *VariableListContext) {}

// ExitVariableList is called when production variableList is exited.
func (s *BaseStackListener) ExitVariableList(ctx *VariableListContext) {}

// EnterIntLiteral is called when production intLiteral is entered.
func (s *BaseStackListener) EnterIntLiteral(ctx *IntLiteralContext) {}

// ExitIntLiteral is called when production intLiteral is exited.
func (s *BaseStackListener) ExitIntLiteral(ctx *IntLiteralContext) {}

// EnterFloatLiteral is called when production floatLiteral is entered.
func (s *BaseStackListener) EnterFloatLiteral(ctx *FloatLiteralContext) {}

// ExitFloatLiteral is called when production floatLiteral is exited.
func (s *BaseStackListener) ExitFloatLiteral(ctx *FloatLiteralContext) {}

// EnterTrueLiteral is called when production trueLiteral is entered.
func (s *BaseStackListener) EnterTrueLiteral(ctx *TrueLiteralContext) {}

// ExitTrueLiteral is called when production trueLiteral is exited.
func (s *BaseStackListener) ExitTrueLiteral(ctx *TrueLiteralContext) {}

// EnterFalseLiteral is called when production falseLiteral is entered.
func (s *BaseStackListener) EnterFalseLiteral(ctx *FalseLiteralContext) {}

// ExitFalseLiteral is called when production falseLiteral is exited.
func (s *BaseStackListener) ExitFalseLiteral(ctx *FalseLiteralContext) {}

// EnterStringLiteral is called when production stringLiteral is entered.
func (s *BaseStackListener) EnterStringLiteral(ctx *StringLiteralContext) {}

// ExitStringLiteral is called when production stringLiteral is exited.
func (s *BaseStackListener) ExitStringLiteral(ctx *StringLiteralContext) {}

// EnterEmptySemStatement is called when production emptySemStatement is entered.
func (s *BaseStackListener) EnterEmptySemStatement(ctx *EmptySemStatementContext) {}

// ExitEmptySemStatement is called when production emptySemStatement is exited.
func (s *BaseStackListener) ExitEmptySemStatement(ctx *EmptySemStatementContext) {}

// EnterVariableStatement is called when production variableStatement is entered.
func (s *BaseStackListener) EnterVariableStatement(ctx *VariableStatementContext) {}

// ExitVariableStatement is called when production variableStatement is exited.
func (s *BaseStackListener) ExitVariableStatement(ctx *VariableStatementContext) {}

// EnterExpressionStatement is called when production expressionStatement is entered.
func (s *BaseStackListener) EnterExpressionStatement(ctx *ExpressionStatementContext) {}

// ExitExpressionStatement is called when production expressionStatement is exited.
func (s *BaseStackListener) ExitExpressionStatement(ctx *ExpressionStatementContext) {}

// EnterReadStatement is called when production readStatement is entered.
func (s *BaseStackListener) EnterReadStatement(ctx *ReadStatementContext) {}

// ExitReadStatement is called when production readStatement is exited.
func (s *BaseStackListener) ExitReadStatement(ctx *ReadStatementContext) {}

// EnterWriteStatement is called when production writeStatement is entered.
func (s *BaseStackListener) EnterWriteStatement(ctx *WriteStatementContext) {}

// ExitWriteStatement is called when production writeStatement is exited.
func (s *BaseStackListener) ExitWriteStatement(ctx *WriteStatementContext) {}

// EnterBlockStatement is called when production blockStatement is entered.
func (s *BaseStackListener) EnterBlockStatement(ctx *BlockStatementContext) {}

// ExitBlockStatement is called when production blockStatement is exited.
func (s *BaseStackListener) ExitBlockStatement(ctx *BlockStatementContext) {}

// EnterIfStatement is called when production ifStatement is entered.
func (s *BaseStackListener) EnterIfStatement(ctx *IfStatementContext) {}

// ExitIfStatement is called when production ifStatement is exited.
func (s *BaseStackListener) ExitIfStatement(ctx *IfStatementContext) {}

// EnterWhileStatement is called when production whileStatement is entered.
func (s *BaseStackListener) EnterWhileStatement(ctx *WhileStatementContext) {}

// ExitWhileStatement is called when production whileStatement is exited.
func (s *BaseStackListener) ExitWhileStatement(ctx *WhileStatementContext) {}

// EnterMulDivModExpression is called when production mulDivModExpression is entered.
func (s *BaseStackListener) EnterMulDivModExpression(ctx *MulDivModExpressionContext) {}

// ExitMulDivModExpression is called when production mulDivModExpression is exited.
func (s *BaseStackListener) ExitMulDivModExpression(ctx *MulDivModExpressionContext) {}

// EnterIdExpression is called when production idExpression is entered.
func (s *BaseStackListener) EnterIdExpression(ctx *IdExpressionContext) {}

// ExitIdExpression is called when production idExpression is exited.
func (s *BaseStackListener) ExitIdExpression(ctx *IdExpressionContext) {}

// EnterAddSubExpression is called when production addSubExpression is entered.
func (s *BaseStackListener) EnterAddSubExpression(ctx *AddSubExpressionContext) {}

// ExitAddSubExpression is called when production addSubExpression is exited.
func (s *BaseStackListener) ExitAddSubExpression(ctx *AddSubExpressionContext) {}

// EnterAssignExpression is called when production assignExpression is entered.
func (s *BaseStackListener) EnterAssignExpression(ctx *AssignExpressionContext) {}

// ExitAssignExpression is called when production assignExpression is exited.
func (s *BaseStackListener) ExitAssignExpression(ctx *AssignExpressionContext) {}

// EnterComparisonExpression is called when production comparisonExpression is entered.
func (s *BaseStackListener) EnterComparisonExpression(ctx *ComparisonExpressionContext) {}

// ExitComparisonExpression is called when production comparisonExpression is exited.
func (s *BaseStackListener) ExitComparisonExpression(ctx *ComparisonExpressionContext) {}

// EnterParenExpression is called when production parenExpression is entered.
func (s *BaseStackListener) EnterParenExpression(ctx *ParenExpressionContext) {}

// ExitParenExpression is called when production parenExpression is exited.
func (s *BaseStackListener) ExitParenExpression(ctx *ParenExpressionContext) {}

// EnterStringConcatenation is called when production stringConcatenation is entered.
func (s *BaseStackListener) EnterStringConcatenation(ctx *StringConcatenationContext) {}

// ExitStringConcatenation is called when production stringConcatenation is exited.
func (s *BaseStackListener) ExitStringConcatenation(ctx *StringConcatenationContext) {}

// EnterLiteralExpression is called when production literalExpression is entered.
func (s *BaseStackListener) EnterLiteralExpression(ctx *LiteralExpressionContext) {}

// ExitLiteralExpression is called when production literalExpression is exited.
func (s *BaseStackListener) ExitLiteralExpression(ctx *LiteralExpressionContext) {}

// EnterUnaryExpression is called when production unaryExpression is entered.
func (s *BaseStackListener) EnterUnaryExpression(ctx *UnaryExpressionContext) {}

// ExitUnaryExpression is called when production unaryExpression is exited.
func (s *BaseStackListener) ExitUnaryExpression(ctx *UnaryExpressionContext) {}

// EnterLogicalExpression is called when production logicalExpression is entered.
func (s *BaseStackListener) EnterLogicalExpression(ctx *LogicalExpressionContext) {}

// ExitLogicalExpression is called when production logicalExpression is exited.
func (s *BaseStackListener) ExitLogicalExpression(ctx *LogicalExpressionContext) {}

// EnterExpressionList is called when production expressionList is entered.
func (s *BaseStackListener) EnterExpressionList(ctx *ExpressionListContext) {}

// ExitExpressionList is called when production expressionList is exited.
func (s *BaseStackListener) ExitExpressionList(ctx *ExpressionListContext) {}
