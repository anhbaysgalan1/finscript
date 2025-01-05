// Code generated from Finscript.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Finscript

import "github.com/antlr4-go/antlr/v4"

// FinscriptListener is a complete listener for a parse tree produced by FinscriptParser.
type FinscriptListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterSendStmt is called when entering the sendStmt production.
	EnterSendStmt(c *SendStmtContext)

	// EnterSourceClause is called when entering the sourceClause production.
	EnterSourceClause(c *SourceClauseContext)

	// EnterDestinationClause is called when entering the destinationClause production.
	EnterDestinationClause(c *DestinationClauseContext)

	// EnterSourceBlock is called when entering the sourceBlock production.
	EnterSourceBlock(c *SourceBlockContext)

	// EnterInorderSrcList is called when entering the inorderSrcList production.
	EnterInorderSrcList(c *InorderSrcListContext)

	// EnterDestBlock is called when entering the destBlock production.
	EnterDestBlock(c *DestBlockContext)

	// EnterAllotmentDestList is called when entering the allotmentDestList production.
	EnterAllotmentDestList(c *AllotmentDestListContext)

	// EnterAllotmentClause is called when entering the allotmentClause production.
	EnterAllotmentClause(c *AllotmentClauseContext)

	// EnterPortion is called when entering the portion production.
	EnterPortion(c *PortionContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitSendStmt is called when exiting the sendStmt production.
	ExitSendStmt(c *SendStmtContext)

	// ExitSourceClause is called when exiting the sourceClause production.
	ExitSourceClause(c *SourceClauseContext)

	// ExitDestinationClause is called when exiting the destinationClause production.
	ExitDestinationClause(c *DestinationClauseContext)

	// ExitSourceBlock is called when exiting the sourceBlock production.
	ExitSourceBlock(c *SourceBlockContext)

	// ExitInorderSrcList is called when exiting the inorderSrcList production.
	ExitInorderSrcList(c *InorderSrcListContext)

	// ExitDestBlock is called when exiting the destBlock production.
	ExitDestBlock(c *DestBlockContext)

	// ExitAllotmentDestList is called when exiting the allotmentDestList production.
	ExitAllotmentDestList(c *AllotmentDestListContext)

	// ExitAllotmentClause is called when exiting the allotmentClause production.
	ExitAllotmentClause(c *AllotmentClauseContext)

	// ExitPortion is called when exiting the portion production.
	ExitPortion(c *PortionContext)
}
