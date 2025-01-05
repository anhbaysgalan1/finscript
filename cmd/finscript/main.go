package main

import (
	"fmt"
	"os"

	"github.com/antlr4-go/antlr/v4"

	"github.com/anhbaysgalan1/finscript/pkg/ast"
	"github.com/anhbaysgalan1/finscript/pkg/parser"
)

// A custom listener that collects the program's "send" statements into AST
type FinscriptListener struct {
	*parser.BaseFinscriptListener
	ASTNodes []ast.SendStmtAST

	currentSend ast.SendStmtAST
}

func (l *FinscriptListener) EnterSendStmt(ctx *parser.SendStmtContext) {
	l.currentSend = ast.SendStmtAST{}
}
func (l *FinscriptListener) ExitSendStmt(ctx *parser.SendStmtContext) {
	// Example: ctx.ASSET(0) is the token
	l.currentSend.Asset = ctx.ASSET().GetText()
	l.currentSend.Amount = ctx.AMOUNT().GetText()
	l.ASTNodes = append(l.ASTNodes, l.currentSend)
}

// Source
func (l *FinscriptListener) ExitSourceClause(ctx *parser.SourceClauseContext) {
	srcAST := buildSourceAST(ctx.SourceBlock())
	l.currentSend.Source = srcAST
}

// Destination
func (l *FinscriptListener) ExitDestinationClause(ctx *parser.DestinationClauseContext) {
	dAST := buildDestAST(ctx.DestBlock())
	l.currentSend.Destination = dAST
}

func buildSourceAST(sc parser.ISourceBlockContext) SourceAST {
	switch real := sc.(type) {
	case *parser.AccountContext:
		return SingleSrc{Account: real.ACCOUNT().GetText()}
	case *parser.BlockInorderContext:
		// parse sub
		var items []SourceAST
		lst := real.InorderSrcList()
		if lst != nil {
			if realList, ok := lst.(*parser.InorderSrcListContext); ok {
				for _, sb := range realList.AllSourceBlock() {
					items = append(items, buildSourceAST(sb))
				}
			}
		}
		return InorderSrc{Sources: items}
	default:
		return SingleSrc{Account: "@unknownSource"}
	}
}

func buildDestAST(db parser.IDestBlockContext) DestAST {
	switch real := db.(type) {
	case *parser.DestAccountContext:
		return SingleDest{Account: real.ACCOUNT().GetText()}
	case *parser.BlockAllotmentContext:
		var items []AllotmentItem
		lst := real.AllotmentDestList()
		if lst != nil {
			if realList, ok := lst.(*parser.AllotmentDestListContext); ok {
				for _, c := range realList.AllAllotmentClause() {
					items = append(items, buildAllotmentClause(c))
				}
			}
		}
		return AllotmentDest{Items: items}
	default:
		return SingleDest{Account: "@unknownDest"}
	}
}

func buildAllotmentClause(c parser.IAllotmentClauseContext) AllotmentItem {
	switch real := c.(type) {
	case *parser.PortionToAccountContext:
		return AllotmentItem{
			IsRemaining: false,
			Portion:     real.Portion().GetText(),
			Account:     real.ACCOUNT().GetText(),
		}
	case *parser.RemainingToAccountContext:
		return AllotmentItem{
			IsRemaining: true,
			Account:     real.ACCOUNT().GetText(),
		}
	default:
		return AllotmentItem{IsRemaining: true, Account: "@unknownRem"}
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <.fins file>\n", os.Args[0])
		os.Exit(1)
	}
	filePath := os.Args[1]
	data, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	input := antlr.NewInputStream(string(data))
	lexer := parser.NewFinscriptLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewFinscriptParser(stream)

	listener := &FinscriptListener{}
	antlr.ParseTreeWalkerDefault.Walk(listener, p.Program())

	// interpret
	var allPosts []Posting
	for _, sendAst := range listener.ASTNodes {
		posts, err := EvaluateSendStmt(&sendAst)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Interpret error: %v\n", err)
			os.Exit(1)
		}
		allPosts = append(allPosts, posts...)
	}

	fmt.Println(`{"postings":[`)
	for i, ps := range allPosts {
		fmt.Printf(`  {"source":%q, "destination":%q, "asset":%q, "amount":%q}`,
			ps.Source, ps.Destination, ps.Asset, ps.Amount.String())
		if i < len(allPosts)-1 {
			fmt.Print(",")
		}
		fmt.Println()
	}
	fmt.Println("]}")
}
