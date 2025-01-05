package ast

// AST for "send [ASSET AMOUNT] ( source=... destination=... )"

type SendStmtAST struct {
	Asset       string
	Amount      string // e.g. "599" or "*"
	Source      SourceAST
	Destination DestAST
}

// SourceAST can be a single account or an inorder block
type SourceAST interface {
	srcNode()
}

type SingleSrc struct {
	Account string
}

func (SingleSrc) srcNode() {}

type InorderSrc struct {
	Sources []SourceAST
}

func (InorderSrc) srcNode() {}

// DestAST can be a single account or an allotment block
type DestAST interface {
	destNode()
}

type SingleDest struct {
	Account string
}

func (SingleDest) destNode() {}

type AllotmentDest struct {
	Items []AllotmentItem
}

func (AllotmentDest) destNode() {}

// Each allotment item can be "portion -> account" or "remaining -> account"
type AllotmentItem struct {
	IsRemaining bool
	Portion     string // if not remaining
	Account     string
}
