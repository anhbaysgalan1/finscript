// Code generated from Finscript.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type FinscriptLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var FinscriptLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func finscriptlexerLexerInit() {
	staticData := &FinscriptLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
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
		"T__0", "SEND", "SOURCE", "DESTINATION", "REMAINING", "EQ", "LPAREN",
		"RPAREN", "LBRACK", "RBRACK", "LBRACE", "RBRACE", "COMMA", "ASSET",
		"AMOUNT", "ACCOUNT", "PERCENTAGE_PORTION_LITERAL", "RATIO_PORTION_LITERAL",
		"WS", "LINE_COMMENT",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 20, 157, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 1, 0, 1,
		0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1,
		3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1,
		5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 1, 10,
		1, 11, 1, 11, 1, 12, 1, 12, 1, 13, 4, 13, 96, 8, 13, 11, 13, 12, 13, 97,
		1, 14, 4, 14, 101, 8, 14, 11, 14, 12, 14, 102, 1, 14, 3, 14, 106, 8, 14,
		1, 15, 1, 15, 4, 15, 110, 8, 15, 11, 15, 12, 15, 111, 1, 16, 4, 16, 115,
		8, 16, 11, 16, 12, 16, 116, 1, 16, 1, 16, 4, 16, 121, 8, 16, 11, 16, 12,
		16, 122, 3, 16, 125, 8, 16, 1, 16, 1, 16, 1, 17, 4, 17, 130, 8, 17, 11,
		17, 12, 17, 131, 1, 17, 1, 17, 4, 17, 136, 8, 17, 11, 17, 12, 17, 137,
		1, 18, 4, 18, 141, 8, 18, 11, 18, 12, 18, 142, 1, 18, 1, 18, 1, 19, 1,
		19, 1, 19, 1, 19, 5, 19, 151, 8, 19, 10, 19, 12, 19, 154, 9, 19, 1, 19,
		1, 19, 0, 0, 20, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17,
		9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35,
		18, 37, 19, 39, 20, 1, 0, 5, 2, 0, 47, 57, 65, 90, 1, 0, 48, 57, 5, 0,
		45, 45, 48, 58, 65, 90, 95, 95, 97, 122, 3, 0, 9, 10, 13, 13, 32, 32, 2,
		0, 10, 10, 13, 13, 167, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0,
		0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1,
		0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21,
		1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0,
		29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0,
		0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 1, 41, 1, 0, 0, 0, 3, 44, 1, 0, 0,
		0, 5, 49, 1, 0, 0, 0, 7, 56, 1, 0, 0, 0, 9, 68, 1, 0, 0, 0, 11, 78, 1,
		0, 0, 0, 13, 80, 1, 0, 0, 0, 15, 82, 1, 0, 0, 0, 17, 84, 1, 0, 0, 0, 19,
		86, 1, 0, 0, 0, 21, 88, 1, 0, 0, 0, 23, 90, 1, 0, 0, 0, 25, 92, 1, 0, 0,
		0, 27, 95, 1, 0, 0, 0, 29, 105, 1, 0, 0, 0, 31, 107, 1, 0, 0, 0, 33, 114,
		1, 0, 0, 0, 35, 129, 1, 0, 0, 0, 37, 140, 1, 0, 0, 0, 39, 146, 1, 0, 0,
		0, 41, 42, 5, 116, 0, 0, 42, 43, 5, 111, 0, 0, 43, 2, 1, 0, 0, 0, 44, 45,
		5, 115, 0, 0, 45, 46, 5, 101, 0, 0, 46, 47, 5, 110, 0, 0, 47, 48, 5, 100,
		0, 0, 48, 4, 1, 0, 0, 0, 49, 50, 5, 115, 0, 0, 50, 51, 5, 111, 0, 0, 51,
		52, 5, 117, 0, 0, 52, 53, 5, 114, 0, 0, 53, 54, 5, 99, 0, 0, 54, 55, 5,
		101, 0, 0, 55, 6, 1, 0, 0, 0, 56, 57, 5, 100, 0, 0, 57, 58, 5, 101, 0,
		0, 58, 59, 5, 115, 0, 0, 59, 60, 5, 116, 0, 0, 60, 61, 5, 105, 0, 0, 61,
		62, 5, 110, 0, 0, 62, 63, 5, 97, 0, 0, 63, 64, 5, 116, 0, 0, 64, 65, 5,
		105, 0, 0, 65, 66, 5, 111, 0, 0, 66, 67, 5, 110, 0, 0, 67, 8, 1, 0, 0,
		0, 68, 69, 5, 114, 0, 0, 69, 70, 5, 101, 0, 0, 70, 71, 5, 109, 0, 0, 71,
		72, 5, 97, 0, 0, 72, 73, 5, 105, 0, 0, 73, 74, 5, 110, 0, 0, 74, 75, 5,
		105, 0, 0, 75, 76, 5, 110, 0, 0, 76, 77, 5, 103, 0, 0, 77, 10, 1, 0, 0,
		0, 78, 79, 5, 61, 0, 0, 79, 12, 1, 0, 0, 0, 80, 81, 5, 40, 0, 0, 81, 14,
		1, 0, 0, 0, 82, 83, 5, 41, 0, 0, 83, 16, 1, 0, 0, 0, 84, 85, 5, 91, 0,
		0, 85, 18, 1, 0, 0, 0, 86, 87, 5, 93, 0, 0, 87, 20, 1, 0, 0, 0, 88, 89,
		5, 123, 0, 0, 89, 22, 1, 0, 0, 0, 90, 91, 5, 125, 0, 0, 91, 24, 1, 0, 0,
		0, 92, 93, 5, 44, 0, 0, 93, 26, 1, 0, 0, 0, 94, 96, 7, 0, 0, 0, 95, 94,
		1, 0, 0, 0, 96, 97, 1, 0, 0, 0, 97, 95, 1, 0, 0, 0, 97, 98, 1, 0, 0, 0,
		98, 28, 1, 0, 0, 0, 99, 101, 7, 1, 0, 0, 100, 99, 1, 0, 0, 0, 101, 102,
		1, 0, 0, 0, 102, 100, 1, 0, 0, 0, 102, 103, 1, 0, 0, 0, 103, 106, 1, 0,
		0, 0, 104, 106, 5, 42, 0, 0, 105, 100, 1, 0, 0, 0, 105, 104, 1, 0, 0, 0,
		106, 30, 1, 0, 0, 0, 107, 109, 5, 64, 0, 0, 108, 110, 7, 2, 0, 0, 109,
		108, 1, 0, 0, 0, 110, 111, 1, 0, 0, 0, 111, 109, 1, 0, 0, 0, 111, 112,
		1, 0, 0, 0, 112, 32, 1, 0, 0, 0, 113, 115, 7, 1, 0, 0, 114, 113, 1, 0,
		0, 0, 115, 116, 1, 0, 0, 0, 116, 114, 1, 0, 0, 0, 116, 117, 1, 0, 0, 0,
		117, 124, 1, 0, 0, 0, 118, 120, 5, 46, 0, 0, 119, 121, 7, 1, 0, 0, 120,
		119, 1, 0, 0, 0, 121, 122, 1, 0, 0, 0, 122, 120, 1, 0, 0, 0, 122, 123,
		1, 0, 0, 0, 123, 125, 1, 0, 0, 0, 124, 118, 1, 0, 0, 0, 124, 125, 1, 0,
		0, 0, 125, 126, 1, 0, 0, 0, 126, 127, 5, 37, 0, 0, 127, 34, 1, 0, 0, 0,
		128, 130, 7, 1, 0, 0, 129, 128, 1, 0, 0, 0, 130, 131, 1, 0, 0, 0, 131,
		129, 1, 0, 0, 0, 131, 132, 1, 0, 0, 0, 132, 133, 1, 0, 0, 0, 133, 135,
		5, 47, 0, 0, 134, 136, 7, 1, 0, 0, 135, 134, 1, 0, 0, 0, 136, 137, 1, 0,
		0, 0, 137, 135, 1, 0, 0, 0, 137, 138, 1, 0, 0, 0, 138, 36, 1, 0, 0, 0,
		139, 141, 7, 3, 0, 0, 140, 139, 1, 0, 0, 0, 141, 142, 1, 0, 0, 0, 142,
		140, 1, 0, 0, 0, 142, 143, 1, 0, 0, 0, 143, 144, 1, 0, 0, 0, 144, 145,
		6, 18, 0, 0, 145, 38, 1, 0, 0, 0, 146, 147, 5, 47, 0, 0, 147, 148, 5, 47,
		0, 0, 148, 152, 1, 0, 0, 0, 149, 151, 8, 4, 0, 0, 150, 149, 1, 0, 0, 0,
		151, 154, 1, 0, 0, 0, 152, 150, 1, 0, 0, 0, 152, 153, 1, 0, 0, 0, 153,
		155, 1, 0, 0, 0, 154, 152, 1, 0, 0, 0, 155, 156, 6, 19, 0, 0, 156, 40,
		1, 0, 0, 0, 12, 0, 97, 102, 105, 111, 116, 122, 124, 131, 137, 142, 152,
		1, 6, 0, 0,
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

// FinscriptLexerInit initializes any static state used to implement FinscriptLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewFinscriptLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func FinscriptLexerInit() {
	staticData := &FinscriptLexerLexerStaticData
	staticData.once.Do(finscriptlexerLexerInit)
}

// NewFinscriptLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewFinscriptLexer(input antlr.CharStream) *FinscriptLexer {
	FinscriptLexerInit()
	l := new(FinscriptLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &FinscriptLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "Finscript.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// FinscriptLexer tokens.
const (
	FinscriptLexerT__0                       = 1
	FinscriptLexerSEND                       = 2
	FinscriptLexerSOURCE                     = 3
	FinscriptLexerDESTINATION                = 4
	FinscriptLexerREMAINING                  = 5
	FinscriptLexerEQ                         = 6
	FinscriptLexerLPAREN                     = 7
	FinscriptLexerRPAREN                     = 8
	FinscriptLexerLBRACK                     = 9
	FinscriptLexerRBRACK                     = 10
	FinscriptLexerLBRACE                     = 11
	FinscriptLexerRBRACE                     = 12
	FinscriptLexerCOMMA                      = 13
	FinscriptLexerASSET                      = 14
	FinscriptLexerAMOUNT                     = 15
	FinscriptLexerACCOUNT                    = 16
	FinscriptLexerPERCENTAGE_PORTION_LITERAL = 17
	FinscriptLexerRATIO_PORTION_LITERAL      = 18
	FinscriptLexerWS                         = 19
	FinscriptLexerLINE_COMMENT               = 20
)
