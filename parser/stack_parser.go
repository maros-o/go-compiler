// Code generated from ./grammars/Stack.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // Stack

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

type StackParser struct {
	*antlr.BaseParser
}

var StackParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func stackParserInit() {
	staticData := &StackParserStaticData
	staticData.LiteralNames = []string{
		"", "'int'", "'float'", "'bool'", "'string'", "'true'", "'false'", "'write'",
		"'read'", "'if'", "'else'", "'while'", "", "", "", "", "'='", "'||'",
		"'&&'", "'=='", "'!='", "'<'", "'>'", "'+'", "'-'", "'*'", "'/'", "'%'",
		"'!'", "';'", "'('", "')'", "'{'", "'}'", "','", "'.'",
	}
	staticData.SymbolicNames = []string{
		"", "INT", "FLOAT", "BOOL", "STRING", "TRUE", "FALSE", "WRITE", "READ",
		"IF", "ELSE", "WHILE", "ID", "INT_LITERAL", "FLOAT_LITERAL", "STRING_LITERAL",
		"ASSIGN", "OR", "AND", "EQ", "NE", "LT", "GT", "ADD", "SUB", "MUL",
		"DIV", "MOD", "NOT", "SEM", "LPAR", "RPAR", "LBRACE", "RBRACE", "COMMA",
		"DOT", "COMMENT", "WS",
	}
	staticData.RuleNames = []string{
		"program", "type", "variableList", "literal", "statement", "expression",
		"expressionList",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 37, 121, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 1, 0, 5, 0, 16, 8, 0, 10, 0, 12, 0, 19, 9, 0,
		1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 5, 2, 26, 8, 2, 10, 2, 12, 2, 29, 9, 2, 1,
		3, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 36, 8, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4,
		1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4,
		1, 4, 5, 4, 56, 8, 4, 10, 4, 12, 4, 59, 9, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1,
		4, 1, 4, 1, 4, 1, 4, 3, 4, 69, 8, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4,
		3, 4, 77, 8, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1,
		5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 94, 8, 5, 1, 5, 1, 5, 1, 5, 1, 5,
		1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 5, 5, 108, 8, 5, 10, 5,
		12, 5, 111, 9, 5, 1, 6, 1, 6, 1, 6, 5, 6, 116, 8, 6, 10, 6, 12, 6, 119,
		9, 6, 1, 6, 0, 1, 10, 7, 0, 2, 4, 6, 8, 10, 12, 0, 6, 1, 0, 1, 4, 2, 0,
		24, 24, 28, 28, 1, 0, 25, 27, 1, 0, 23, 24, 1, 0, 19, 22, 1, 0, 17, 18,
		138, 0, 17, 1, 0, 0, 0, 2, 20, 1, 0, 0, 0, 4, 22, 1, 0, 0, 0, 6, 35, 1,
		0, 0, 0, 8, 76, 1, 0, 0, 0, 10, 93, 1, 0, 0, 0, 12, 112, 1, 0, 0, 0, 14,
		16, 3, 8, 4, 0, 15, 14, 1, 0, 0, 0, 16, 19, 1, 0, 0, 0, 17, 15, 1, 0, 0,
		0, 17, 18, 1, 0, 0, 0, 18, 1, 1, 0, 0, 0, 19, 17, 1, 0, 0, 0, 20, 21, 7,
		0, 0, 0, 21, 3, 1, 0, 0, 0, 22, 27, 5, 12, 0, 0, 23, 24, 5, 34, 0, 0, 24,
		26, 5, 12, 0, 0, 25, 23, 1, 0, 0, 0, 26, 29, 1, 0, 0, 0, 27, 25, 1, 0,
		0, 0, 27, 28, 1, 0, 0, 0, 28, 5, 1, 0, 0, 0, 29, 27, 1, 0, 0, 0, 30, 36,
		5, 13, 0, 0, 31, 36, 5, 14, 0, 0, 32, 36, 5, 5, 0, 0, 33, 36, 5, 6, 0,
		0, 34, 36, 5, 15, 0, 0, 35, 30, 1, 0, 0, 0, 35, 31, 1, 0, 0, 0, 35, 32,
		1, 0, 0, 0, 35, 33, 1, 0, 0, 0, 35, 34, 1, 0, 0, 0, 36, 7, 1, 0, 0, 0,
		37, 77, 5, 29, 0, 0, 38, 39, 3, 2, 1, 0, 39, 40, 3, 4, 2, 0, 40, 41, 5,
		29, 0, 0, 41, 77, 1, 0, 0, 0, 42, 43, 3, 10, 5, 0, 43, 44, 5, 29, 0, 0,
		44, 77, 1, 0, 0, 0, 45, 46, 5, 8, 0, 0, 46, 47, 3, 4, 2, 0, 47, 48, 5,
		29, 0, 0, 48, 77, 1, 0, 0, 0, 49, 50, 5, 7, 0, 0, 50, 51, 3, 12, 6, 0,
		51, 52, 5, 29, 0, 0, 52, 77, 1, 0, 0, 0, 53, 57, 5, 32, 0, 0, 54, 56, 3,
		8, 4, 0, 55, 54, 1, 0, 0, 0, 56, 59, 1, 0, 0, 0, 57, 55, 1, 0, 0, 0, 57,
		58, 1, 0, 0, 0, 58, 60, 1, 0, 0, 0, 59, 57, 1, 0, 0, 0, 60, 77, 5, 33,
		0, 0, 61, 62, 5, 9, 0, 0, 62, 63, 5, 30, 0, 0, 63, 64, 3, 10, 5, 0, 64,
		65, 5, 31, 0, 0, 65, 68, 3, 8, 4, 0, 66, 67, 5, 10, 0, 0, 67, 69, 3, 8,
		4, 0, 68, 66, 1, 0, 0, 0, 68, 69, 1, 0, 0, 0, 69, 77, 1, 0, 0, 0, 70, 71,
		5, 11, 0, 0, 71, 72, 5, 30, 0, 0, 72, 73, 3, 10, 5, 0, 73, 74, 5, 31, 0,
		0, 74, 75, 3, 8, 4, 0, 75, 77, 1, 0, 0, 0, 76, 37, 1, 0, 0, 0, 76, 38,
		1, 0, 0, 0, 76, 42, 1, 0, 0, 0, 76, 45, 1, 0, 0, 0, 76, 49, 1, 0, 0, 0,
		76, 53, 1, 0, 0, 0, 76, 61, 1, 0, 0, 0, 76, 70, 1, 0, 0, 0, 77, 9, 1, 0,
		0, 0, 78, 79, 6, 5, -1, 0, 79, 94, 3, 6, 3, 0, 80, 94, 5, 12, 0, 0, 81,
		82, 5, 12, 0, 0, 82, 83, 5, 16, 0, 0, 83, 94, 3, 10, 5, 8, 84, 85, 7, 1,
		0, 0, 85, 94, 3, 10, 5, 3, 86, 87, 5, 30, 0, 0, 87, 88, 3, 10, 5, 0, 88,
		89, 5, 31, 0, 0, 89, 94, 1, 0, 0, 0, 90, 91, 5, 15, 0, 0, 91, 92, 5, 35,
		0, 0, 92, 94, 5, 15, 0, 0, 93, 78, 1, 0, 0, 0, 93, 80, 1, 0, 0, 0, 93,
		81, 1, 0, 0, 0, 93, 84, 1, 0, 0, 0, 93, 86, 1, 0, 0, 0, 93, 90, 1, 0, 0,
		0, 94, 109, 1, 0, 0, 0, 95, 96, 10, 7, 0, 0, 96, 97, 7, 2, 0, 0, 97, 108,
		3, 10, 5, 8, 98, 99, 10, 6, 0, 0, 99, 100, 7, 3, 0, 0, 100, 108, 3, 10,
		5, 7, 101, 102, 10, 5, 0, 0, 102, 103, 7, 4, 0, 0, 103, 108, 3, 10, 5,
		6, 104, 105, 10, 4, 0, 0, 105, 106, 7, 5, 0, 0, 106, 108, 3, 10, 5, 5,
		107, 95, 1, 0, 0, 0, 107, 98, 1, 0, 0, 0, 107, 101, 1, 0, 0, 0, 107, 104,
		1, 0, 0, 0, 108, 111, 1, 0, 0, 0, 109, 107, 1, 0, 0, 0, 109, 110, 1, 0,
		0, 0, 110, 11, 1, 0, 0, 0, 111, 109, 1, 0, 0, 0, 112, 117, 3, 10, 5, 0,
		113, 114, 5, 34, 0, 0, 114, 116, 3, 10, 5, 0, 115, 113, 1, 0, 0, 0, 116,
		119, 1, 0, 0, 0, 117, 115, 1, 0, 0, 0, 117, 118, 1, 0, 0, 0, 118, 13, 1,
		0, 0, 0, 119, 117, 1, 0, 0, 0, 10, 17, 27, 35, 57, 68, 76, 93, 107, 109,
		117,
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

// StackParserInit initializes any static state used to implement StackParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewStackParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func StackParserInit() {
	staticData := &StackParserStaticData
	staticData.once.Do(stackParserInit)
}

// NewStackParser produces a new parser instance for the optional input antlr.TokenStream.
func NewStackParser(input antlr.TokenStream) *StackParser {
	StackParserInit()
	this := new(StackParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &StackParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "Stack.g4"

	return this
}

// StackParser tokens.
const (
	StackParserEOF            = antlr.TokenEOF
	StackParserINT            = 1
	StackParserFLOAT          = 2
	StackParserBOOL           = 3
	StackParserSTRING         = 4
	StackParserTRUE           = 5
	StackParserFALSE          = 6
	StackParserWRITE          = 7
	StackParserREAD           = 8
	StackParserIF             = 9
	StackParserELSE           = 10
	StackParserWHILE          = 11
	StackParserID             = 12
	StackParserINT_LITERAL    = 13
	StackParserFLOAT_LITERAL  = 14
	StackParserSTRING_LITERAL = 15
	StackParserASSIGN         = 16
	StackParserOR             = 17
	StackParserAND            = 18
	StackParserEQ             = 19
	StackParserNE             = 20
	StackParserLT             = 21
	StackParserGT             = 22
	StackParserADD            = 23
	StackParserSUB            = 24
	StackParserMUL            = 25
	StackParserDIV            = 26
	StackParserMOD            = 27
	StackParserNOT            = 28
	StackParserSEM            = 29
	StackParserLPAR           = 30
	StackParserRPAR           = 31
	StackParserLBRACE         = 32
	StackParserRBRACE         = 33
	StackParserCOMMA          = 34
	StackParserDOT            = 35
	StackParserCOMMENT        = 36
	StackParserWS             = 37
)

// StackParser rules.
const (
	StackParserRULE_program        = 0
	StackParserRULE_type           = 1
	StackParserRULE_variableList   = 2
	StackParserRULE_literal        = 3
	StackParserRULE_statement      = 4
	StackParserRULE_expression     = 5
	StackParserRULE_expressionList = 6
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext

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
	p.RuleIndex = StackParserRULE_program
	return p
}

func InitEmptyProgramContext(p *ProgramContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = StackParserRULE_program
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = StackParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
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

	return t.(IStatementContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StackListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StackListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (p *StackParser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, StackParserRULE_program)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(17)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&6190857214) != 0 {
		{
			p.SetState(14)
			p.Statement()
		}

		p.SetState(19)
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

// ITypeContext is an interface to support dynamic dispatch.
type ITypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INT() antlr.TerminalNode
	FLOAT() antlr.TerminalNode
	BOOL() antlr.TerminalNode
	STRING() antlr.TerminalNode

	// IsTypeContext differentiates from other interfaces.
	IsTypeContext()
}

type TypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeContext() *TypeContext {
	var p = new(TypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = StackParserRULE_type
	return p
}

func InitEmptyTypeContext(p *TypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = StackParserRULE_type
}

func (*TypeContext) IsTypeContext() {}

func NewTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeContext {
	var p = new(TypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = StackParserRULE_type

	return p
}

func (s *TypeContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeContext) INT() antlr.TerminalNode {
	return s.GetToken(StackParserINT, 0)
}

func (s *TypeContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(StackParserFLOAT, 0)
}

func (s *TypeContext) BOOL() antlr.TerminalNode {
	return s.GetToken(StackParserBOOL, 0)
}

func (s *TypeContext) STRING() antlr.TerminalNode {
	return s.GetToken(StackParserSTRING, 0)
}

func (s *TypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StackListener); ok {
		listenerT.EnterType(s)
	}
}

func (s *TypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StackListener); ok {
		listenerT.ExitType(s)
	}
}

func (p *StackParser) Type_() (localctx ITypeContext) {
	localctx = NewTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, StackParserRULE_type)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(20)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&30) != 0) {
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

// IVariableListContext is an interface to support dynamic dispatch.
type IVariableListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllID() []antlr.TerminalNode
	ID(i int) antlr.TerminalNode
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsVariableListContext differentiates from other interfaces.
	IsVariableListContext()
}

type VariableListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableListContext() *VariableListContext {
	var p = new(VariableListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = StackParserRULE_variableList
	return p
}

func InitEmptyVariableListContext(p *VariableListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = StackParserRULE_variableList
}

func (*VariableListContext) IsVariableListContext() {}

func NewVariableListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableListContext {
	var p = new(VariableListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = StackParserRULE_variableList

	return p
}

func (s *VariableListContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableListContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(StackParserID)
}

func (s *VariableListContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(StackParserID, i)
}

func (s *VariableListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(StackParserCOMMA)
}

func (s *VariableListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(StackParserCOMMA, i)
}

func (s *VariableListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StackListener); ok {
		listenerT.EnterVariableList(s)
	}
}

func (s *VariableListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StackListener); ok {
		listenerT.ExitVariableList(s)
	}
}

func (p *StackParser) VariableList() (localctx IVariableListContext) {
	localctx = NewVariableListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, StackParserRULE_variableList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(22)
		p.Match(StackParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(27)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == StackParserCOMMA {
		{
			p.SetState(23)
			p.Match(StackParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(24)
			p.Match(StackParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(29)
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

// ILiteralContext is an interface to support dynamic dispatch.
type ILiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsLiteralContext differentiates from other interfaces.
	IsLiteralContext()
}

type LiteralContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralContext() *LiteralContext {
	var p = new(LiteralContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = StackParserRULE_literal
	return p
}

func InitEmptyLiteralContext(p *LiteralContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = StackParserRULE_literal
}

func (*LiteralContext) IsLiteralContext() {}

func NewLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralContext {
	var p = new(LiteralContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = StackParserRULE_literal

	return p
}

func (s *LiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralContext) CopyAll(ctx *LiteralContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *LiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type FalseLiteralContext struct {
	LiteralContext
}

func NewFalseLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FalseLiteralContext {
	var p = new(FalseLiteralContext)

	InitEmptyLiteralContext(&p.LiteralContext)
	p.parser = parser
	p.CopyAll(ctx.(*LiteralContext))

	return p
}

func (s *FalseLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FalseLiteralContext) FALSE() antlr.TerminalNode {
	return s.GetToken(StackParserFALSE, 0)
}

func (s *FalseLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StackListener); ok {
		listenerT.EnterFalseLiteral(s)
	}
}

func (s *FalseLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StackListener); ok {
		listenerT.ExitFalseLiteral(s)
	}
}

type StringLiteralContext struct {
	LiteralContext
}

func NewStringLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StringLiteralContext {
	var p = new(StringLiteralContext)

	InitEmptyLiteralContext(&p.LiteralContext)
	p.parser = parser
	p.CopyAll(ctx.(*LiteralContext))

	return p
}

func (s *StringLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringLiteralContext) STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(StackParserSTRING_LITERAL, 0)
}

func (s *StringLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StackListener); ok {
		listenerT.EnterStringLiteral(s)
	}
}

func (s *StringLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StackListener); ok {
		listenerT.ExitStringLiteral(s)
	}
}

type IntLiteralContext struct {
	LiteralContext
}

func NewIntLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IntLiteralContext {
	var p = new(IntLiteralContext)

	InitEmptyLiteralContext(&p.LiteralContext)
	p.parser = parser
	p.CopyAll(ctx.(*LiteralContext))

	return p
}

func (s *IntLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntLiteralContext) INT_LITERAL() antlr.TerminalNode {
	return s.GetToken(StackParserINT_LITERAL, 0)
}

func (s *IntLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StackListener); ok {
		listenerT.EnterIntLiteral(s)
	}
}

func (s *IntLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StackListener); ok {
		listenerT.ExitIntLiteral(s)
	}
}

type FloatLiteralContext struct {
	LiteralContext
}

func NewFloatLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FloatLiteralContext {
	var p = new(FloatLiteralContext)

	InitEmptyLiteralContext(&p.LiteralContext)
	p.parser = parser
	p.CopyAll(ctx.(*LiteralContext))

	return p
}

func (s *FloatLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FloatLiteralContext) FLOAT_LITERAL() antlr.TerminalNode {
	return s.GetToken(StackParserFLOAT_LITERAL, 0)
}

func (s *FloatLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StackListener); ok {
		listenerT.EnterFloatLiteral(s)
	}
}

func (s *FloatLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StackListener); ok {
		listenerT.ExitFloatLiteral(s)
	}
}

type TrueLiteralContext struct {
	LiteralContext
}

func NewTrueLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TrueLiteralContext {
	var p = new(TrueLiteralContext)

	InitEmptyLiteralContext(&p.LiteralContext)
	p.parser = parser
	p.CopyAll(ctx.(*LiteralContext))

	return p
}

func (s *TrueLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TrueLiteralContext) TRUE() antlr.TerminalNode {
	return s.GetToken(StackParserTRUE, 0)
}

func (s *TrueLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StackListener); ok {
		listenerT.EnterTrueLiteral(s)
	}
}

func (s *TrueLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StackListener); ok {
		listenerT.ExitTrueLiteral(s)
	}
}

func (p *StackParser) Literal() (localctx ILiteralContext) {
	localctx = NewLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, StackParserRULE_literal)
	p.SetState(35)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case StackParserINT_LITERAL:
		localctx = NewIntLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(30)
			p.Match(StackParserINT_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case StackParserFLOAT_LITERAL:
		localctx = NewFloatLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(31)
			p.Match(StackParserFLOAT_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case StackParserTRUE:
		localctx = NewTrueLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(32)
			p.Match(StackParserTRUE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case StackParserFALSE:
		localctx = NewFalseLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(33)
			p.Match(StackParserFALSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case StackParserSTRING_LITERAL:
		localctx = NewStringLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(34)
			p.Match(StackParserSTRING_LITERAL)
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

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = StackParserRULE_statement
	return p
}

func InitEmptyStatementContext(p *StatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = StackParserRULE_statement
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = StackParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) CopyAll(ctx *StatementContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type VariableStatementContext struct {
	StatementContext
}

func NewVariableStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VariableStatementContext {
	var p = new(VariableStatementContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *VariableStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableStatementContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *VariableStatementContext) VariableList() IVariableListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariableListContext)
}

func (s *VariableStatementContext) SEM() antlr.TerminalNode {
	return s.GetToken(StackParserSEM, 0)
}

func (s *VariableStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StackListener); ok {
		listenerT.EnterVariableStatement(s)
	}
}

func (s *VariableStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StackListener); ok {
		listenerT.ExitVariableStatement(s)
	}
}

type WhileStatementContext struct {
	StatementContext
}

func NewWhileStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *WhileStatementContext {
	var p = new(WhileStatementContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *WhileStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhileStatementContext) WHILE() antlr.TerminalNode {
	return s.GetToken(StackParserWHILE, 0)
}

func (s *WhileStatementContext) LPAR() antlr.TerminalNode {
	return s.GetToken(StackParserLPAR, 0)
}

func (s *WhileStatementContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *WhileStatementContext) RPAR() antlr.TerminalNode {
	return s.GetToken(StackParserRPAR, 0)
}

func (s *WhileStatementContext) Statement() IStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *WhileStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StackListener); ok {
		listenerT.EnterWhileStatement(s)
	}
}

func (s *WhileStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StackListener); ok {
		listenerT.ExitWhileStatement(s)
	}
}

type WriteStatementContext struct {
	StatementContext
}

func NewWriteStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *WriteStatementContext {
	var p = new(WriteStatementContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *WriteStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WriteStatementContext) WRITE() antlr.TerminalNode {
	return s.GetToken(StackParserWRITE, 0)
}

func (s *WriteStatementContext) ExpressionList() IExpressionListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionListContext)
}

func (s *WriteStatementContext) SEM() antlr.TerminalNode {
	return s.GetToken(StackParserSEM, 0)
}

func (s *WriteStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StackListener); ok {
		listenerT.EnterWriteStatement(s)
	}
}

func (s *WriteStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StackListener); ok {
		listenerT.ExitWriteStatement(s)
	}
}

type BlockStatementContext struct {
	StatementContext
}

func NewBlockStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BlockStatementContext {
	var p = new(BlockStatementContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *BlockStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockStatementContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(StackParserLBRACE, 0)
}

func (s *BlockStatementContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(StackParserRBRACE, 0)
}

func (s *BlockStatementContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *BlockStatementContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
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

	return t.(IStatementContext)
}

func (s *BlockStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StackListener); ok {
		listenerT.EnterBlockStatement(s)
	}
}

func (s *BlockStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StackListener); ok {
		listenerT.ExitBlockStatement(s)
	}
}

type ReadStatementContext struct {
	StatementContext
}

func NewReadStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ReadStatementContext {
	var p = new(ReadStatementContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *ReadStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReadStatementContext) READ() antlr.TerminalNode {
	return s.GetToken(StackParserREAD, 0)
}

func (s *ReadStatementContext) VariableList() IVariableListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariableListContext)
}

func (s *ReadStatementContext) SEM() antlr.TerminalNode {
	return s.GetToken(StackParserSEM, 0)
}

func (s *ReadStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StackListener); ok {
		listenerT.EnterReadStatement(s)
	}
}

func (s *ReadStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StackListener); ok {
		listenerT.ExitReadStatement(s)
	}
}

type ExpressionStatementContext struct {
	StatementContext
}

func NewExpressionStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpressionStatementContext {
	var p = new(ExpressionStatementContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *ExpressionStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionStatementContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionStatementContext) SEM() antlr.TerminalNode {
	return s.GetToken(StackParserSEM, 0)
}

func (s *ExpressionStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StackListener); ok {
		listenerT.EnterExpressionStatement(s)
	}
}

func (s *ExpressionStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StackListener); ok {
		listenerT.ExitExpressionStatement(s)
	}
}

type IfStatementContext struct {
	StatementContext
}

func NewIfStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IfStatementContext {
	var p = new(IfStatementContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *IfStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfStatementContext) IF() antlr.TerminalNode {
	return s.GetToken(StackParserIF, 0)
}

func (s *IfStatementContext) LPAR() antlr.TerminalNode {
	return s.GetToken(StackParserLPAR, 0)
}

func (s *IfStatementContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *IfStatementContext) RPAR() antlr.TerminalNode {
	return s.GetToken(StackParserRPAR, 0)
}

func (s *IfStatementContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *IfStatementContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
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

	return t.(IStatementContext)
}

func (s *IfStatementContext) ELSE() antlr.TerminalNode {
	return s.GetToken(StackParserELSE, 0)
}

func (s *IfStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StackListener); ok {
		listenerT.EnterIfStatement(s)
	}
}

func (s *IfStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StackListener); ok {
		listenerT.ExitIfStatement(s)
	}
}

type EmptySemStatementContext struct {
	StatementContext
}

func NewEmptySemStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EmptySemStatementContext {
	var p = new(EmptySemStatementContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *EmptySemStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EmptySemStatementContext) SEM() antlr.TerminalNode {
	return s.GetToken(StackParserSEM, 0)
}

func (s *EmptySemStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StackListener); ok {
		listenerT.EnterEmptySemStatement(s)
	}
}

func (s *EmptySemStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StackListener); ok {
		listenerT.ExitEmptySemStatement(s)
	}
}

func (p *StackParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, StackParserRULE_statement)
	var _la int

	p.SetState(76)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case StackParserSEM:
		localctx = NewEmptySemStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(37)
			p.Match(StackParserSEM)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case StackParserINT, StackParserFLOAT, StackParserBOOL, StackParserSTRING:
		localctx = NewVariableStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(38)
			p.Type_()
		}
		{
			p.SetState(39)
			p.VariableList()
		}
		{
			p.SetState(40)
			p.Match(StackParserSEM)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case StackParserTRUE, StackParserFALSE, StackParserID, StackParserINT_LITERAL, StackParserFLOAT_LITERAL, StackParserSTRING_LITERAL, StackParserSUB, StackParserNOT, StackParserLPAR:
		localctx = NewExpressionStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(42)
			p.expression(0)
		}
		{
			p.SetState(43)
			p.Match(StackParserSEM)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case StackParserREAD:
		localctx = NewReadStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(45)
			p.Match(StackParserREAD)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(46)
			p.VariableList()
		}
		{
			p.SetState(47)
			p.Match(StackParserSEM)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case StackParserWRITE:
		localctx = NewWriteStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(49)
			p.Match(StackParserWRITE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(50)
			p.ExpressionList()
		}
		{
			p.SetState(51)
			p.Match(StackParserSEM)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case StackParserLBRACE:
		localctx = NewBlockStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(53)
			p.Match(StackParserLBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(57)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&6190857214) != 0 {
			{
				p.SetState(54)
				p.Statement()
			}

			p.SetState(59)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(60)
			p.Match(StackParserRBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case StackParserIF:
		localctx = NewIfStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(61)
			p.Match(StackParserIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(62)
			p.Match(StackParserLPAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(63)
			p.expression(0)
		}
		{
			p.SetState(64)
			p.Match(StackParserRPAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(65)
			p.Statement()
		}
		p.SetState(68)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(66)
				p.Match(StackParserELSE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(67)
				p.Statement()
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}

	case StackParserWHILE:
		localctx = NewWhileStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(70)
			p.Match(StackParserWHILE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(71)
			p.Match(StackParserLPAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(72)
			p.expression(0)
		}
		{
			p.SetState(73)
			p.Match(StackParserRPAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(74)
			p.Statement()
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

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = StackParserRULE_expression
	return p
}

func InitEmptyExpressionContext(p *ExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = StackParserRULE_expression
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = StackParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) CopyAll(ctx *ExpressionContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type MulDivModExpressionContext struct {
	ExpressionContext
	op antlr.Token
}

func NewMulDivModExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MulDivModExpressionContext {
	var p = new(MulDivModExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *MulDivModExpressionContext) GetOp() antlr.Token { return s.op }

func (s *MulDivModExpressionContext) SetOp(v antlr.Token) { s.op = v }

func (s *MulDivModExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MulDivModExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *MulDivModExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
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

	return t.(IExpressionContext)
}

func (s *MulDivModExpressionContext) MUL() antlr.TerminalNode {
	return s.GetToken(StackParserMUL, 0)
}

func (s *MulDivModExpressionContext) DIV() antlr.TerminalNode {
	return s.GetToken(StackParserDIV, 0)
}

func (s *MulDivModExpressionContext) MOD() antlr.TerminalNode {
	return s.GetToken(StackParserMOD, 0)
}

func (s *MulDivModExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StackListener); ok {
		listenerT.EnterMulDivModExpression(s)
	}
}

func (s *MulDivModExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StackListener); ok {
		listenerT.ExitMulDivModExpression(s)
	}
}

type IdExpressionContext struct {
	ExpressionContext
}

func NewIdExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IdExpressionContext {
	var p = new(IdExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *IdExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdExpressionContext) ID() antlr.TerminalNode {
	return s.GetToken(StackParserID, 0)
}

func (s *IdExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StackListener); ok {
		listenerT.EnterIdExpression(s)
	}
}

func (s *IdExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StackListener); ok {
		listenerT.ExitIdExpression(s)
	}
}

type AddSubExpressionContext struct {
	ExpressionContext
	op antlr.Token
}

func NewAddSubExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AddSubExpressionContext {
	var p = new(AddSubExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *AddSubExpressionContext) GetOp() antlr.Token { return s.op }

func (s *AddSubExpressionContext) SetOp(v antlr.Token) { s.op = v }

func (s *AddSubExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AddSubExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *AddSubExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
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

	return t.(IExpressionContext)
}

func (s *AddSubExpressionContext) ADD() antlr.TerminalNode {
	return s.GetToken(StackParserADD, 0)
}

func (s *AddSubExpressionContext) SUB() antlr.TerminalNode {
	return s.GetToken(StackParserSUB, 0)
}

func (s *AddSubExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StackListener); ok {
		listenerT.EnterAddSubExpression(s)
	}
}

func (s *AddSubExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StackListener); ok {
		listenerT.ExitAddSubExpression(s)
	}
}

type AssignExpressionContext struct {
	ExpressionContext
}

func NewAssignExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignExpressionContext {
	var p = new(AssignExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *AssignExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignExpressionContext) ID() antlr.TerminalNode {
	return s.GetToken(StackParserID, 0)
}

func (s *AssignExpressionContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(StackParserASSIGN, 0)
}

func (s *AssignExpressionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AssignExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StackListener); ok {
		listenerT.EnterAssignExpression(s)
	}
}

func (s *AssignExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StackListener); ok {
		listenerT.ExitAssignExpression(s)
	}
}

type ComparisonExpressionContext struct {
	ExpressionContext
	op antlr.Token
}

func NewComparisonExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ComparisonExpressionContext {
	var p = new(ComparisonExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *ComparisonExpressionContext) GetOp() antlr.Token { return s.op }

func (s *ComparisonExpressionContext) SetOp(v antlr.Token) { s.op = v }

func (s *ComparisonExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComparisonExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ComparisonExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
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

	return t.(IExpressionContext)
}

func (s *ComparisonExpressionContext) LT() antlr.TerminalNode {
	return s.GetToken(StackParserLT, 0)
}

func (s *ComparisonExpressionContext) GT() antlr.TerminalNode {
	return s.GetToken(StackParserGT, 0)
}

func (s *ComparisonExpressionContext) EQ() antlr.TerminalNode {
	return s.GetToken(StackParserEQ, 0)
}

func (s *ComparisonExpressionContext) NE() antlr.TerminalNode {
	return s.GetToken(StackParserNE, 0)
}

func (s *ComparisonExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StackListener); ok {
		listenerT.EnterComparisonExpression(s)
	}
}

func (s *ComparisonExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StackListener); ok {
		listenerT.ExitComparisonExpression(s)
	}
}

type ParenExpressionContext struct {
	ExpressionContext
}

func NewParenExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParenExpressionContext {
	var p = new(ParenExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *ParenExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParenExpressionContext) LPAR() antlr.TerminalNode {
	return s.GetToken(StackParserLPAR, 0)
}

func (s *ParenExpressionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ParenExpressionContext) RPAR() antlr.TerminalNode {
	return s.GetToken(StackParserRPAR, 0)
}

func (s *ParenExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StackListener); ok {
		listenerT.EnterParenExpression(s)
	}
}

func (s *ParenExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StackListener); ok {
		listenerT.ExitParenExpression(s)
	}
}

type StringConcatenationContext struct {
	ExpressionContext
}

func NewStringConcatenationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StringConcatenationContext {
	var p = new(StringConcatenationContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *StringConcatenationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringConcatenationContext) AllSTRING_LITERAL() []antlr.TerminalNode {
	return s.GetTokens(StackParserSTRING_LITERAL)
}

func (s *StringConcatenationContext) STRING_LITERAL(i int) antlr.TerminalNode {
	return s.GetToken(StackParserSTRING_LITERAL, i)
}

func (s *StringConcatenationContext) DOT() antlr.TerminalNode {
	return s.GetToken(StackParserDOT, 0)
}

func (s *StringConcatenationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StackListener); ok {
		listenerT.EnterStringConcatenation(s)
	}
}

func (s *StringConcatenationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StackListener); ok {
		listenerT.ExitStringConcatenation(s)
	}
}

type LiteralExpressionContext struct {
	ExpressionContext
}

func NewLiteralExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LiteralExpressionContext {
	var p = new(LiteralExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *LiteralExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralExpressionContext) Literal() ILiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *LiteralExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StackListener); ok {
		listenerT.EnterLiteralExpression(s)
	}
}

func (s *LiteralExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StackListener); ok {
		listenerT.ExitLiteralExpression(s)
	}
}

type UnaryExpressionContext struct {
	ExpressionContext
	op antlr.Token
}

func NewUnaryExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UnaryExpressionContext {
	var p = new(UnaryExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *UnaryExpressionContext) GetOp() antlr.Token { return s.op }

func (s *UnaryExpressionContext) SetOp(v antlr.Token) { s.op = v }

func (s *UnaryExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryExpressionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *UnaryExpressionContext) NOT() antlr.TerminalNode {
	return s.GetToken(StackParserNOT, 0)
}

func (s *UnaryExpressionContext) SUB() antlr.TerminalNode {
	return s.GetToken(StackParserSUB, 0)
}

func (s *UnaryExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StackListener); ok {
		listenerT.EnterUnaryExpression(s)
	}
}

func (s *UnaryExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StackListener); ok {
		listenerT.ExitUnaryExpression(s)
	}
}

type LogicalExpressionContext struct {
	ExpressionContext
	op antlr.Token
}

func NewLogicalExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LogicalExpressionContext {
	var p = new(LogicalExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *LogicalExpressionContext) GetOp() antlr.Token { return s.op }

func (s *LogicalExpressionContext) SetOp(v antlr.Token) { s.op = v }

func (s *LogicalExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogicalExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *LogicalExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
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

	return t.(IExpressionContext)
}

func (s *LogicalExpressionContext) AND() antlr.TerminalNode {
	return s.GetToken(StackParserAND, 0)
}

func (s *LogicalExpressionContext) OR() antlr.TerminalNode {
	return s.GetToken(StackParserOR, 0)
}

func (s *LogicalExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StackListener); ok {
		listenerT.EnterLogicalExpression(s)
	}
}

func (s *LogicalExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StackListener); ok {
		listenerT.ExitLogicalExpression(s)
	}
}

func (p *StackParser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *StackParser) expression(_p int) (localctx IExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 10
	p.EnterRecursionRule(localctx, 10, StackParserRULE_expression, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(93)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		localctx = NewLiteralExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(79)
			p.Literal()
		}

	case 2:
		localctx = NewIdExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(80)
			p.Match(StackParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		localctx = NewAssignExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(81)
			p.Match(StackParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(82)
			p.Match(StackParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(83)
			p.expression(8)
		}

	case 4:
		localctx = NewUnaryExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(84)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*UnaryExpressionContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == StackParserSUB || _la == StackParserNOT) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*UnaryExpressionContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(85)
			p.expression(3)
		}

	case 5:
		localctx = NewParenExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(86)
			p.Match(StackParserLPAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(87)
			p.expression(0)
		}
		{
			p.SetState(88)
			p.Match(StackParserRPAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		localctx = NewStringConcatenationContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(90)
			p.Match(StackParserSTRING_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(91)
			p.Match(StackParserDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(92)
			p.Match(StackParserSTRING_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(109)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(107)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext()) {
			case 1:
				localctx = NewMulDivModExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, StackParserRULE_expression)
				p.SetState(95)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
					goto errorExit
				}
				{
					p.SetState(96)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*MulDivModExpressionContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&234881024) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*MulDivModExpressionContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(97)
					p.expression(8)
				}

			case 2:
				localctx = NewAddSubExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, StackParserRULE_expression)
				p.SetState(98)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}
				{
					p.SetState(99)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*AddSubExpressionContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == StackParserADD || _la == StackParserSUB) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*AddSubExpressionContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(100)
					p.expression(7)
				}

			case 3:
				localctx = NewComparisonExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, StackParserRULE_expression)
				p.SetState(101)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(102)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ComparisonExpressionContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&7864320) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ComparisonExpressionContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(103)
					p.expression(6)
				}

			case 4:
				localctx = NewLogicalExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, StackParserRULE_expression)
				p.SetState(104)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(105)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*LogicalExpressionContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == StackParserOR || _la == StackParserAND) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*LogicalExpressionContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(106)
					p.expression(5)
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(111)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext())
		if p.HasError() {
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
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpressionListContext is an interface to support dynamic dispatch.
type IExpressionListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsExpressionListContext differentiates from other interfaces.
	IsExpressionListContext()
}

type ExpressionListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionListContext() *ExpressionListContext {
	var p = new(ExpressionListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = StackParserRULE_expressionList
	return p
}

func InitEmptyExpressionListContext(p *ExpressionListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = StackParserRULE_expressionList
}

func (*ExpressionListContext) IsExpressionListContext() {}

func NewExpressionListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionListContext {
	var p = new(ExpressionListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = StackParserRULE_expressionList

	return p
}

func (s *ExpressionListContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionListContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ExpressionListContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
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

	return t.(IExpressionContext)
}

func (s *ExpressionListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(StackParserCOMMA)
}

func (s *ExpressionListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(StackParserCOMMA, i)
}

func (s *ExpressionListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StackListener); ok {
		listenerT.EnterExpressionList(s)
	}
}

func (s *ExpressionListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StackListener); ok {
		listenerT.ExitExpressionList(s)
	}
}

func (p *StackParser) ExpressionList() (localctx IExpressionListContext) {
	localctx = NewExpressionListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, StackParserRULE_expressionList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(112)
		p.expression(0)
	}
	p.SetState(117)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == StackParserCOMMA {
		{
			p.SetState(113)
			p.Match(StackParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(114)
			p.expression(0)
		}

		p.SetState(119)
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

func (p *StackParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 5:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *StackParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 4)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}