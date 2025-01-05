// Code generated from Finscript.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Finscript

import "github.com/antlr4-go/antlr/v4"

// BaseFinscriptListener is a complete listener for a parse tree produced by FinscriptParser.
type BaseFinscriptListener struct{}

var _ FinscriptListener = &BaseFinscriptListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseFinscriptListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseFinscriptListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseFinscriptListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseFinscriptListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BaseFinscriptListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BaseFinscriptListener) ExitProgram(ctx *ProgramContext) {}

// EnterSendStmt is called when production sendStmt is entered.
func (s *BaseFinscriptListener) EnterSendStmt(ctx *SendStmtContext) {}

// ExitSendStmt is called when production sendStmt is exited.
func (s *BaseFinscriptListener) ExitSendStmt(ctx *SendStmtContext) {}

// EnterSourceClause is called when production sourceClause is entered.
func (s *BaseFinscriptListener) EnterSourceClause(ctx *SourceClauseContext) {}

// ExitSourceClause is called when production sourceClause is exited.
func (s *BaseFinscriptListener) ExitSourceClause(ctx *SourceClauseContext) {}

// EnterDestinationClause is called when production destinationClause is entered.
func (s *BaseFinscriptListener) EnterDestinationClause(ctx *DestinationClauseContext) {}

// ExitDestinationClause is called when production destinationClause is exited.
func (s *BaseFinscriptListener) ExitDestinationClause(ctx *DestinationClauseContext) {}

// EnterSourceBlock is called when production sourceBlock is entered.
func (s *BaseFinscriptListener) EnterSourceBlock(ctx *SourceBlockContext) {}

// ExitSourceBlock is called when production sourceBlock is exited.
func (s *BaseFinscriptListener) ExitSourceBlock(ctx *SourceBlockContext) {}

// EnterInorderSrcList is called when production inorderSrcList is entered.
func (s *BaseFinscriptListener) EnterInorderSrcList(ctx *InorderSrcListContext) {}

// ExitInorderSrcList is called when production inorderSrcList is exited.
func (s *BaseFinscriptListener) ExitInorderSrcList(ctx *InorderSrcListContext) {}

// EnterDestBlock is called when production destBlock is entered.
func (s *BaseFinscriptListener) EnterDestBlock(ctx *DestBlockContext) {}

// ExitDestBlock is called when production destBlock is exited.
func (s *BaseFinscriptListener) ExitDestBlock(ctx *DestBlockContext) {}

// EnterAllotmentDestList is called when production allotmentDestList is entered.
func (s *BaseFinscriptListener) EnterAllotmentDestList(ctx *AllotmentDestListContext) {}

// ExitAllotmentDestList is called when production allotmentDestList is exited.
func (s *BaseFinscriptListener) ExitAllotmentDestList(ctx *AllotmentDestListContext) {}

// EnterAllotmentClause is called when production allotmentClause is entered.
func (s *BaseFinscriptListener) EnterAllotmentClause(ctx *AllotmentClauseContext) {}

// ExitAllotmentClause is called when production allotmentClause is exited.
func (s *BaseFinscriptListener) ExitAllotmentClause(ctx *AllotmentClauseContext) {}

// EnterPortion is called when production portion is entered.
func (s *BaseFinscriptListener) EnterPortion(ctx *PortionContext) {}

// ExitPortion is called when production portion is exited.
func (s *BaseFinscriptListener) ExitPortion(ctx *PortionContext) {}
