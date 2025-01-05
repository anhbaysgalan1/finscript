// Code generated from Finscript.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Finscript

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type FinscriptParser struct {
	*antlr.BaseParser
}

var FinscriptParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func finscriptParserInit() {
	staticData := &FinscriptParserStaticData
	staticData.LiteralNames = []string{
		"", "'to'", "'send'", "'source'", "'destination'", "'remaining'", "'='",
		"'('", "')'", "'['", "']'", "'{'", "'}'", "','",
	}
	staticData.SymbolicNames = []string{
		"", "", "SEND", "SOURCE", "DESTINATION", "REMAINING", "EQ", "LPAREN",
		"RPAREN", "LBRACK", "RBRACK", "LBRACE", "RBRACE", "COMMA", "ASSET",
		"AMOUNT", "ACCOUNT", "PERCENTAGE_PORTION_LITERAL", "RATIO_PORTION_LITERAL",
		"WS", "LINE_COMMENT",
	}
	staticData.RuleNames = []string{
		"program", "sendStmt", "sourceClause", "destinationClause", "sourceBlock",
		"inorderSrcList", "destBlock", "allotmentDestList", "allotmentClause",
		"portion",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 20, 89, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 1, 0, 4,
		0, 22, 8, 0, 11, 0, 12, 0, 23, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3,
		1, 3, 1, 4, 1, 4, 1, 4, 3, 4, 49, 8, 4, 1, 4, 3, 4, 52, 8, 4, 1, 5, 1,
		5, 1, 5, 5, 5, 57, 8, 5, 10, 5, 12, 5, 60, 9, 5, 1, 6, 1, 6, 1, 6, 3, 6,
		65, 8, 6, 1, 6, 3, 6, 68, 8, 6, 1, 7, 1, 7, 1, 7, 5, 7, 73, 8, 7, 10, 7,
		12, 7, 76, 9, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 3, 8, 85, 8,
		8, 1, 9, 1, 9, 1, 9, 0, 0, 10, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 0, 1,
		1, 0, 17, 18, 86, 0, 21, 1, 0, 0, 0, 2, 27, 1, 0, 0, 0, 4, 37, 1, 0, 0,
		0, 6, 41, 1, 0, 0, 0, 8, 51, 1, 0, 0, 0, 10, 53, 1, 0, 0, 0, 12, 67, 1,
		0, 0, 0, 14, 69, 1, 0, 0, 0, 16, 84, 1, 0, 0, 0, 18, 86, 1, 0, 0, 0, 20,
		22, 3, 2, 1, 0, 21, 20, 1, 0, 0, 0, 22, 23, 1, 0, 0, 0, 23, 21, 1, 0, 0,
		0, 23, 24, 1, 0, 0, 0, 24, 25, 1, 0, 0, 0, 25, 26, 5, 0, 0, 1, 26, 1, 1,
		0, 0, 0, 27, 28, 5, 2, 0, 0, 28, 29, 5, 9, 0, 0, 29, 30, 5, 14, 0, 0, 30,
		31, 5, 15, 0, 0, 31, 32, 5, 10, 0, 0, 32, 33, 5, 7, 0, 0, 33, 34, 3, 4,
		2, 0, 34, 35, 3, 6, 3, 0, 35, 36, 5, 8, 0, 0, 36, 3, 1, 0, 0, 0, 37, 38,
		5, 3, 0, 0, 38, 39, 5, 6, 0, 0, 39, 40, 3, 8, 4, 0, 40, 5, 1, 0, 0, 0,
		41, 42, 5, 4, 0, 0, 42, 43, 5, 6, 0, 0, 43, 44, 3, 12, 6, 0, 44, 7, 1,
		0, 0, 0, 45, 52, 5, 16, 0, 0, 46, 48, 5, 11, 0, 0, 47, 49, 3, 10, 5, 0,
		48, 47, 1, 0, 0, 0, 48, 49, 1, 0, 0, 0, 49, 50, 1, 0, 0, 0, 50, 52, 5,
		12, 0, 0, 51, 45, 1, 0, 0, 0, 51, 46, 1, 0, 0, 0, 52, 9, 1, 0, 0, 0, 53,
		58, 3, 8, 4, 0, 54, 55, 5, 13, 0, 0, 55, 57, 3, 8, 4, 0, 56, 54, 1, 0,
		0, 0, 57, 60, 1, 0, 0, 0, 58, 56, 1, 0, 0, 0, 58, 59, 1, 0, 0, 0, 59, 11,
		1, 0, 0, 0, 60, 58, 1, 0, 0, 0, 61, 68, 5, 16, 0, 0, 62, 64, 5, 11, 0,
		0, 63, 65, 3, 14, 7, 0, 64, 63, 1, 0, 0, 0, 64, 65, 1, 0, 0, 0, 65, 66,
		1, 0, 0, 0, 66, 68, 5, 12, 0, 0, 67, 61, 1, 0, 0, 0, 67, 62, 1, 0, 0, 0,
		68, 13, 1, 0, 0, 0, 69, 74, 3, 16, 8, 0, 70, 71, 5, 13, 0, 0, 71, 73, 3,
		16, 8, 0, 72, 70, 1, 0, 0, 0, 73, 76, 1, 0, 0, 0, 74, 72, 1, 0, 0, 0, 74,
		75, 1, 0, 0, 0, 75, 15, 1, 0, 0, 0, 76, 74, 1, 0, 0, 0, 77, 78, 3, 18,
		9, 0, 78, 79, 5, 1, 0, 0, 79, 80, 5, 16, 0, 0, 80, 85, 1, 0, 0, 0, 81,
		82, 5, 5, 0, 0, 82, 83, 5, 1, 0, 0, 83, 85, 5, 16, 0, 0, 84, 77, 1, 0,
		0, 0, 84, 81, 1, 0, 0, 0, 85, 17, 1, 0, 0, 0, 86, 87, 7, 0, 0, 0, 87, 19,
		1, 0, 0, 0, 8, 23, 48, 51, 58, 64, 67, 74, 84,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// FinscriptParserInit initializes any static state used to implement FinscriptParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewFinscriptParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func FinscriptParserInit() {
	staticData := &FinscriptParserStaticData
	staticData.once.Do(finscriptParserInit)
}

// NewFinscriptParser produces a new parser instance for the optional input antlr.TokenStream.
func NewFinscriptParser(input antlr.TokenStream) *FinscriptParser {
	FinscriptParserInit()
	this := new(FinscriptParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &FinscriptParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "Finscript.g4"

	return this
}

// FinscriptParser tokens.
const (
	FinscriptParserEOF                        = antlr.TokenEOF
	FinscriptParserT__0                       = 1
	FinscriptParserSEND                       = 2
	FinscriptParserSOURCE                     = 3
	FinscriptParserDESTINATION                = 4
	FinscriptParserREMAINING                  = 5
	FinscriptParserEQ                         = 6
	FinscriptParserLPAREN                     = 7
	FinscriptParserRPAREN                     = 8
	FinscriptParserLBRACK                     = 9
	FinscriptParserRBRACK                     = 10
	FinscriptParserLBRACE                     = 11
	FinscriptParserRBRACE                     = 12
	FinscriptParserCOMMA                      = 13
	FinscriptParserASSET                      = 14
	FinscriptParserAMOUNT                     = 15
	FinscriptParserACCOUNT                    = 16
	FinscriptParserPERCENTAGE_PORTION_LITERAL = 17
	FinscriptParserRATIO_PORTION_LITERAL      = 18
	FinscriptParserWS                         = 19
	FinscriptParserLINE_COMMENT               = 20
)

// FinscriptParser rules.
const (
	FinscriptParserRULE_program           = 0
	FinscriptParserRULE_sendStmt          = 1
	FinscriptParserRULE_sourceClause      = 2
	FinscriptParserRULE_destinationClause = 3
	FinscriptParserRULE_sourceBlock       = 4
	FinscriptParserRULE_inorderSrcList    = 5
	FinscriptParserRULE_destBlock         = 6
	FinscriptParserRULE_allotmentDestList = 7
	FinscriptParserRULE_allotmentClause   = 8
	FinscriptParserRULE_portion           = 9
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
	AllSendStmt() []ISendStmtContext
	SendStmt(i int) ISendStmtContext

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FinscriptParserRULE_program
	return p
}

func InitEmptyProgramContext(p *ProgramContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FinscriptParserRULE_program
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FinscriptParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) EOF() antlr.TerminalNode {
	return s.GetToken(FinscriptParserEOF, 0)
}

func (s *ProgramContext) AllSendStmt() []ISendStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISendStmtContext); ok {
			len++
		}
	}

	tst := make([]ISendStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISendStmtContext); ok {
			tst[i] = t.(ISendStmtContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) SendStmt(i int) ISendStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISendStmtContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISendStmtContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FinscriptListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FinscriptListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (s *ProgramContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FinscriptVisitor:
		return t.VisitProgram(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FinscriptParser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, FinscriptParserRULE_program)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(21)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == FinscriptParserSEND {
		{
			p.SetState(20)
			p.SendStmt()
		}

		p.SetState(23)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(25)
		p.Match(FinscriptParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISendStmtContext is an interface to support dynamic dispatch.
type ISendStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SEND() antlr.TerminalNode
	LBRACK() antlr.TerminalNode
	ASSET() antlr.TerminalNode
	AMOUNT() antlr.TerminalNode
	RBRACK() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	SourceClause() ISourceClauseContext
	DestinationClause() IDestinationClauseContext
	RPAREN() antlr.TerminalNode

	// IsSendStmtContext differentiates from other interfaces.
	IsSendStmtContext()
}

type SendStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySendStmtContext() *SendStmtContext {
	var p = new(SendStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FinscriptParserRULE_sendStmt
	return p
}

func InitEmptySendStmtContext(p *SendStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FinscriptParserRULE_sendStmt
}

func (*SendStmtContext) IsSendStmtContext() {}

func NewSendStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SendStmtContext {
	var p = new(SendStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FinscriptParserRULE_sendStmt

	return p
}

func (s *SendStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *SendStmtContext) SEND() antlr.TerminalNode {
	return s.GetToken(FinscriptParserSEND, 0)
}

func (s *SendStmtContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(FinscriptParserLBRACK, 0)
}

func (s *SendStmtContext) ASSET() antlr.TerminalNode {
	return s.GetToken(FinscriptParserASSET, 0)
}

func (s *SendStmtContext) AMOUNT() antlr.TerminalNode {
	return s.GetToken(FinscriptParserAMOUNT, 0)
}

func (s *SendStmtContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(FinscriptParserRBRACK, 0)
}

func (s *SendStmtContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(FinscriptParserLPAREN, 0)
}

func (s *SendStmtContext) SourceClause() ISourceClauseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISourceClauseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISourceClauseContext)
}

func (s *SendStmtContext) DestinationClause() IDestinationClauseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDestinationClauseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDestinationClauseContext)
}

func (s *SendStmtContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(FinscriptParserRPAREN, 0)
}

func (s *SendStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SendStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SendStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FinscriptListener); ok {
		listenerT.EnterSendStmt(s)
	}
}

func (s *SendStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FinscriptListener); ok {
		listenerT.ExitSendStmt(s)
	}
}

func (s *SendStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FinscriptVisitor:
		return t.VisitSendStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FinscriptParser) SendStmt() (localctx ISendStmtContext) {
	localctx = NewSendStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, FinscriptParserRULE_sendStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(27)
		p.Match(FinscriptParserSEND)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(28)
		p.Match(FinscriptParserLBRACK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(29)
		p.Match(FinscriptParserASSET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(30)
		p.Match(FinscriptParserAMOUNT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(31)
		p.Match(FinscriptParserRBRACK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(32)
		p.Match(FinscriptParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(33)
		p.SourceClause()
	}
	{
		p.SetState(34)
		p.DestinationClause()
	}
	{
		p.SetState(35)
		p.Match(FinscriptParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISourceClauseContext is an interface to support dynamic dispatch.
type ISourceClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SOURCE() antlr.TerminalNode
	EQ() antlr.TerminalNode
	SourceBlock() ISourceBlockContext

	// IsSourceClauseContext differentiates from other interfaces.
	IsSourceClauseContext()
}

type SourceClauseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySourceClauseContext() *SourceClauseContext {
	var p = new(SourceClauseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FinscriptParserRULE_sourceClause
	return p
}

func InitEmptySourceClauseContext(p *SourceClauseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FinscriptParserRULE_sourceClause
}

func (*SourceClauseContext) IsSourceClauseContext() {}

func NewSourceClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SourceClauseContext {
	var p = new(SourceClauseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FinscriptParserRULE_sourceClause

	return p
}

func (s *SourceClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *SourceClauseContext) SOURCE() antlr.TerminalNode {
	return s.GetToken(FinscriptParserSOURCE, 0)
}

func (s *SourceClauseContext) EQ() antlr.TerminalNode {
	return s.GetToken(FinscriptParserEQ, 0)
}

func (s *SourceClauseContext) SourceBlock() ISourceBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISourceBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISourceBlockContext)
}

func (s *SourceClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SourceClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SourceClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FinscriptListener); ok {
		listenerT.EnterSourceClause(s)
	}
}

func (s *SourceClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FinscriptListener); ok {
		listenerT.ExitSourceClause(s)
	}
}

func (s *SourceClauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FinscriptVisitor:
		return t.VisitSourceClause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FinscriptParser) SourceClause() (localctx ISourceClauseContext) {
	localctx = NewSourceClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, FinscriptParserRULE_sourceClause)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(37)
		p.Match(FinscriptParserSOURCE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(38)
		p.Match(FinscriptParserEQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(39)
		p.SourceBlock()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDestinationClauseContext is an interface to support dynamic dispatch.
type IDestinationClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DESTINATION() antlr.TerminalNode
	EQ() antlr.TerminalNode
	DestBlock() IDestBlockContext

	// IsDestinationClauseContext differentiates from other interfaces.
	IsDestinationClauseContext()
}

type DestinationClauseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDestinationClauseContext() *DestinationClauseContext {
	var p = new(DestinationClauseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FinscriptParserRULE_destinationClause
	return p
}

func InitEmptyDestinationClauseContext(p *DestinationClauseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FinscriptParserRULE_destinationClause
}

func (*DestinationClauseContext) IsDestinationClauseContext() {}

func NewDestinationClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DestinationClauseContext {
	var p = new(DestinationClauseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FinscriptParserRULE_destinationClause

	return p
}

func (s *DestinationClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *DestinationClauseContext) DESTINATION() antlr.TerminalNode {
	return s.GetToken(FinscriptParserDESTINATION, 0)
}

func (s *DestinationClauseContext) EQ() antlr.TerminalNode {
	return s.GetToken(FinscriptParserEQ, 0)
}

func (s *DestinationClauseContext) DestBlock() IDestBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDestBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDestBlockContext)
}

func (s *DestinationClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DestinationClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DestinationClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FinscriptListener); ok {
		listenerT.EnterDestinationClause(s)
	}
}

func (s *DestinationClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FinscriptListener); ok {
		listenerT.ExitDestinationClause(s)
	}
}

func (s *DestinationClauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FinscriptVisitor:
		return t.VisitDestinationClause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FinscriptParser) DestinationClause() (localctx IDestinationClauseContext) {
	localctx = NewDestinationClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, FinscriptParserRULE_destinationClause)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(41)
		p.Match(FinscriptParserDESTINATION)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(42)
		p.Match(FinscriptParserEQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(43)
		p.DestBlock()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISourceBlockContext is an interface to support dynamic dispatch.
type ISourceBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ACCOUNT() antlr.TerminalNode
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	InorderSrcList() IInorderSrcListContext

	// IsSourceBlockContext differentiates from other interfaces.
	IsSourceBlockContext()
}

type SourceBlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySourceBlockContext() *SourceBlockContext {
	var p = new(SourceBlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FinscriptParserRULE_sourceBlock
	return p
}

func InitEmptySourceBlockContext(p *SourceBlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FinscriptParserRULE_sourceBlock
}

func (*SourceBlockContext) IsSourceBlockContext() {}

func NewSourceBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SourceBlockContext {
	var p = new(SourceBlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FinscriptParserRULE_sourceBlock

	return p
}

func (s *SourceBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *SourceBlockContext) ACCOUNT() antlr.TerminalNode {
	return s.GetToken(FinscriptParserACCOUNT, 0)
}

func (s *SourceBlockContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(FinscriptParserLBRACE, 0)
}

func (s *SourceBlockContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(FinscriptParserRBRACE, 0)
}

func (s *SourceBlockContext) InorderSrcList() IInorderSrcListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInorderSrcListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInorderSrcListContext)
}

func (s *SourceBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SourceBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SourceBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FinscriptListener); ok {
		listenerT.EnterSourceBlock(s)
	}
}

func (s *SourceBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FinscriptListener); ok {
		listenerT.ExitSourceBlock(s)
	}
}

func (s *SourceBlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FinscriptVisitor:
		return t.VisitSourceBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FinscriptParser) SourceBlock() (localctx ISourceBlockContext) {
	localctx = NewSourceBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, FinscriptParserRULE_sourceBlock)
	var _la int

	p.SetState(51)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case FinscriptParserACCOUNT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(45)
			p.Match(FinscriptParserACCOUNT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case FinscriptParserLBRACE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(46)
			p.Match(FinscriptParserLBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(48)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == FinscriptParserLBRACE || _la == FinscriptParserACCOUNT {
			{
				p.SetState(47)
				p.InorderSrcList()
			}

		}
		{
			p.SetState(50)
			p.Match(FinscriptParserRBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IInorderSrcListContext is an interface to support dynamic dispatch.
type IInorderSrcListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllSourceBlock() []ISourceBlockContext
	SourceBlock(i int) ISourceBlockContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsInorderSrcListContext differentiates from other interfaces.
	IsInorderSrcListContext()
}

type InorderSrcListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInorderSrcListContext() *InorderSrcListContext {
	var p = new(InorderSrcListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FinscriptParserRULE_inorderSrcList
	return p
}

func InitEmptyInorderSrcListContext(p *InorderSrcListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FinscriptParserRULE_inorderSrcList
}

func (*InorderSrcListContext) IsInorderSrcListContext() {}

func NewInorderSrcListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InorderSrcListContext {
	var p = new(InorderSrcListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FinscriptParserRULE_inorderSrcList

	return p
}

func (s *InorderSrcListContext) GetParser() antlr.Parser { return s.parser }

func (s *InorderSrcListContext) AllSourceBlock() []ISourceBlockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISourceBlockContext); ok {
			len++
		}
	}

	tst := make([]ISourceBlockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISourceBlockContext); ok {
			tst[i] = t.(ISourceBlockContext)
			i++
		}
	}

	return tst
}

func (s *InorderSrcListContext) SourceBlock(i int) ISourceBlockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISourceBlockContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISourceBlockContext)
}

func (s *InorderSrcListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(FinscriptParserCOMMA)
}

func (s *InorderSrcListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(FinscriptParserCOMMA, i)
}

func (s *InorderSrcListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InorderSrcListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InorderSrcListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FinscriptListener); ok {
		listenerT.EnterInorderSrcList(s)
	}
}

func (s *InorderSrcListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FinscriptListener); ok {
		listenerT.ExitInorderSrcList(s)
	}
}

func (s *InorderSrcListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FinscriptVisitor:
		return t.VisitInorderSrcList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FinscriptParser) InorderSrcList() (localctx IInorderSrcListContext) {
	localctx = NewInorderSrcListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, FinscriptParserRULE_inorderSrcList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(53)
		p.SourceBlock()
	}
	p.SetState(58)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == FinscriptParserCOMMA {
		{
			p.SetState(54)
			p.Match(FinscriptParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(55)
			p.SourceBlock()
		}

		p.SetState(60)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDestBlockContext is an interface to support dynamic dispatch.
type IDestBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ACCOUNT() antlr.TerminalNode
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	AllotmentDestList() IAllotmentDestListContext

	// IsDestBlockContext differentiates from other interfaces.
	IsDestBlockContext()
}

type DestBlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDestBlockContext() *DestBlockContext {
	var p = new(DestBlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FinscriptParserRULE_destBlock
	return p
}

func InitEmptyDestBlockContext(p *DestBlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FinscriptParserRULE_destBlock
}

func (*DestBlockContext) IsDestBlockContext() {}

func NewDestBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DestBlockContext {
	var p = new(DestBlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FinscriptParserRULE_destBlock

	return p
}

func (s *DestBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *DestBlockContext) ACCOUNT() antlr.TerminalNode {
	return s.GetToken(FinscriptParserACCOUNT, 0)
}

func (s *DestBlockContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(FinscriptParserLBRACE, 0)
}

func (s *DestBlockContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(FinscriptParserRBRACE, 0)
}

func (s *DestBlockContext) AllotmentDestList() IAllotmentDestListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAllotmentDestListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAllotmentDestListContext)
}

func (s *DestBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DestBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DestBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FinscriptListener); ok {
		listenerT.EnterDestBlock(s)
	}
}

func (s *DestBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FinscriptListener); ok {
		listenerT.ExitDestBlock(s)
	}
}

func (s *DestBlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FinscriptVisitor:
		return t.VisitDestBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FinscriptParser) DestBlock() (localctx IDestBlockContext) {
	localctx = NewDestBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, FinscriptParserRULE_destBlock)
	var _la int

	p.SetState(67)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case FinscriptParserACCOUNT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(61)
			p.Match(FinscriptParserACCOUNT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case FinscriptParserLBRACE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(62)
			p.Match(FinscriptParserLBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(64)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&393248) != 0 {
			{
				p.SetState(63)
				p.AllotmentDestList()
			}

		}
		{
			p.SetState(66)
			p.Match(FinscriptParserRBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAllotmentDestListContext is an interface to support dynamic dispatch.
type IAllotmentDestListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllAllotmentClause() []IAllotmentClauseContext
	AllotmentClause(i int) IAllotmentClauseContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsAllotmentDestListContext differentiates from other interfaces.
	IsAllotmentDestListContext()
}

type AllotmentDestListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAllotmentDestListContext() *AllotmentDestListContext {
	var p = new(AllotmentDestListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FinscriptParserRULE_allotmentDestList
	return p
}

func InitEmptyAllotmentDestListContext(p *AllotmentDestListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FinscriptParserRULE_allotmentDestList
}

func (*AllotmentDestListContext) IsAllotmentDestListContext() {}

func NewAllotmentDestListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AllotmentDestListContext {
	var p = new(AllotmentDestListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FinscriptParserRULE_allotmentDestList

	return p
}

func (s *AllotmentDestListContext) GetParser() antlr.Parser { return s.parser }

func (s *AllotmentDestListContext) AllAllotmentClause() []IAllotmentClauseContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAllotmentClauseContext); ok {
			len++
		}
	}

	tst := make([]IAllotmentClauseContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAllotmentClauseContext); ok {
			tst[i] = t.(IAllotmentClauseContext)
			i++
		}
	}

	return tst
}

func (s *AllotmentDestListContext) AllotmentClause(i int) IAllotmentClauseContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAllotmentClauseContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAllotmentClauseContext)
}

func (s *AllotmentDestListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(FinscriptParserCOMMA)
}

func (s *AllotmentDestListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(FinscriptParserCOMMA, i)
}

func (s *AllotmentDestListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AllotmentDestListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AllotmentDestListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FinscriptListener); ok {
		listenerT.EnterAllotmentDestList(s)
	}
}

func (s *AllotmentDestListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FinscriptListener); ok {
		listenerT.ExitAllotmentDestList(s)
	}
}

func (s *AllotmentDestListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FinscriptVisitor:
		return t.VisitAllotmentDestList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FinscriptParser) AllotmentDestList() (localctx IAllotmentDestListContext) {
	localctx = NewAllotmentDestListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, FinscriptParserRULE_allotmentDestList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(69)
		p.AllotmentClause()
	}
	p.SetState(74)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == FinscriptParserCOMMA {
		{
			p.SetState(70)
			p.Match(FinscriptParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(71)
			p.AllotmentClause()
		}

		p.SetState(76)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAllotmentClauseContext is an interface to support dynamic dispatch.
type IAllotmentClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Portion() IPortionContext
	ACCOUNT() antlr.TerminalNode
	REMAINING() antlr.TerminalNode

	// IsAllotmentClauseContext differentiates from other interfaces.
	IsAllotmentClauseContext()
}

type AllotmentClauseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAllotmentClauseContext() *AllotmentClauseContext {
	var p = new(AllotmentClauseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FinscriptParserRULE_allotmentClause
	return p
}

func InitEmptyAllotmentClauseContext(p *AllotmentClauseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FinscriptParserRULE_allotmentClause
}

func (*AllotmentClauseContext) IsAllotmentClauseContext() {}

func NewAllotmentClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AllotmentClauseContext {
	var p = new(AllotmentClauseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FinscriptParserRULE_allotmentClause

	return p
}

func (s *AllotmentClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *AllotmentClauseContext) Portion() IPortionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPortionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPortionContext)
}

func (s *AllotmentClauseContext) ACCOUNT() antlr.TerminalNode {
	return s.GetToken(FinscriptParserACCOUNT, 0)
}

func (s *AllotmentClauseContext) REMAINING() antlr.TerminalNode {
	return s.GetToken(FinscriptParserREMAINING, 0)
}

func (s *AllotmentClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AllotmentClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AllotmentClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FinscriptListener); ok {
		listenerT.EnterAllotmentClause(s)
	}
}

func (s *AllotmentClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FinscriptListener); ok {
		listenerT.ExitAllotmentClause(s)
	}
}

func (s *AllotmentClauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FinscriptVisitor:
		return t.VisitAllotmentClause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FinscriptParser) AllotmentClause() (localctx IAllotmentClauseContext) {
	localctx = NewAllotmentClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, FinscriptParserRULE_allotmentClause)
	p.SetState(84)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case FinscriptParserPERCENTAGE_PORTION_LITERAL, FinscriptParserRATIO_PORTION_LITERAL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(77)
			p.Portion()
		}
		{
			p.SetState(78)
			p.Match(FinscriptParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(79)
			p.Match(FinscriptParserACCOUNT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case FinscriptParserREMAINING:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(81)
			p.Match(FinscriptParserREMAINING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(82)
			p.Match(FinscriptParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(83)
			p.Match(FinscriptParserACCOUNT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPortionContext is an interface to support dynamic dispatch.
type IPortionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PERCENTAGE_PORTION_LITERAL() antlr.TerminalNode
	RATIO_PORTION_LITERAL() antlr.TerminalNode

	// IsPortionContext differentiates from other interfaces.
	IsPortionContext()
}

type PortionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPortionContext() *PortionContext {
	var p = new(PortionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FinscriptParserRULE_portion
	return p
}

func InitEmptyPortionContext(p *PortionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FinscriptParserRULE_portion
}

func (*PortionContext) IsPortionContext() {}

func NewPortionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PortionContext {
	var p = new(PortionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FinscriptParserRULE_portion

	return p
}

func (s *PortionContext) GetParser() antlr.Parser { return s.parser }

func (s *PortionContext) PERCENTAGE_PORTION_LITERAL() antlr.TerminalNode {
	return s.GetToken(FinscriptParserPERCENTAGE_PORTION_LITERAL, 0)
}

func (s *PortionContext) RATIO_PORTION_LITERAL() antlr.TerminalNode {
	return s.GetToken(FinscriptParserRATIO_PORTION_LITERAL, 0)
}

func (s *PortionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PortionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PortionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FinscriptListener); ok {
		listenerT.EnterPortion(s)
	}
}

func (s *PortionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FinscriptListener); ok {
		listenerT.ExitPortion(s)
	}
}

func (s *PortionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FinscriptVisitor:
		return t.VisitPortion(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FinscriptParser) Portion() (localctx IPortionContext) {
	localctx = NewPortionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, FinscriptParserRULE_portion)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(86)
		_la = p.GetTokenStream().LA(1)

		if !(_la == FinscriptParserPERCENTAGE_PORTION_LITERAL || _la == FinscriptParserRATIO_PORTION_LITERAL) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}
