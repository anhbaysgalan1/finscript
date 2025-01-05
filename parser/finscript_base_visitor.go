// Code generated from Finscript.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Finscript

import "github.com/antlr4-go/antlr/v4"

type BaseFinscriptVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseFinscriptVisitor) VisitProgram(ctx *ProgramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFinscriptVisitor) VisitSendStmt(ctx *SendStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFinscriptVisitor) VisitSourceClause(ctx *SourceClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFinscriptVisitor) VisitDestinationClause(ctx *DestinationClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFinscriptVisitor) VisitSourceBlock(ctx *SourceBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFinscriptVisitor) VisitInorderSrcList(ctx *InorderSrcListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFinscriptVisitor) VisitDestBlock(ctx *DestBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFinscriptVisitor) VisitAllotmentDestList(ctx *AllotmentDestListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFinscriptVisitor) VisitAllotmentClause(ctx *AllotmentClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFinscriptVisitor) VisitPortion(ctx *PortionContext) interface{} {
	return v.VisitChildren(ctx)
}
