// Code generated from Finscript.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Finscript

import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by FinscriptParser.
type FinscriptVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by FinscriptParser#program.
	VisitProgram(ctx *ProgramContext) interface{}

	// Visit a parse tree produced by FinscriptParser#sendStmt.
	VisitSendStmt(ctx *SendStmtContext) interface{}

	// Visit a parse tree produced by FinscriptParser#sourceClause.
	VisitSourceClause(ctx *SourceClauseContext) interface{}

	// Visit a parse tree produced by FinscriptParser#destinationClause.
	VisitDestinationClause(ctx *DestinationClauseContext) interface{}

	// Visit a parse tree produced by FinscriptParser#sourceBlock.
	VisitSourceBlock(ctx *SourceBlockContext) interface{}

	// Visit a parse tree produced by FinscriptParser#inorderSrcList.
	VisitInorderSrcList(ctx *InorderSrcListContext) interface{}

	// Visit a parse tree produced by FinscriptParser#destBlock.
	VisitDestBlock(ctx *DestBlockContext) interface{}

	// Visit a parse tree produced by FinscriptParser#allotmentDestList.
	VisitAllotmentDestList(ctx *AllotmentDestListContext) interface{}

	// Visit a parse tree produced by FinscriptParser#allotmentClause.
	VisitAllotmentClause(ctx *AllotmentClauseContext) interface{}

	// Visit a parse tree produced by FinscriptParser#portion.
	VisitPortion(ctx *PortionContext) interface{}
}
