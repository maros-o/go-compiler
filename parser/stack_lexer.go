// Code generated from ./grammars/Stack.g4 by ANTLR 4.13.1. DO NOT EDIT.

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

type StackLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var StackLexerLexerStaticData struct {
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

func stacklexerLexerInit() {
	staticData := &StackLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
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
		"INT", "FLOAT", "BOOL", "STRING", "TRUE", "FALSE", "WRITE", "READ",
		"IF", "ELSE", "WHILE", "ID", "INT_LITERAL", "FLOAT_LITERAL", "STRING_LITERAL",
		"ASSIGN", "OR", "AND", "EQ", "NE", "LT", "GT", "ADD", "SUB", "MUL",
		"DIV", "MOD", "NOT", "SEM", "LPAR", "RPAR", "LBRACE", "RBRACE", "COMMA",
		"DOT", "COMMENT", "WS",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 37, 227, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4,
		1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6,
		1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8,
		1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10,
		1, 11, 1, 11, 5, 11, 136, 8, 11, 10, 11, 12, 11, 139, 9, 11, 1, 12, 4,
		12, 142, 8, 12, 11, 12, 12, 12, 143, 1, 13, 4, 13, 147, 8, 13, 11, 13,
		12, 13, 148, 1, 13, 1, 13, 4, 13, 153, 8, 13, 11, 13, 12, 13, 154, 1, 14,
		1, 14, 5, 14, 159, 8, 14, 10, 14, 12, 14, 162, 9, 14, 1, 14, 1, 14, 1,
		15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18,
		1, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1, 21, 1, 21, 1, 22, 1, 22, 1, 23, 1,
		23, 1, 24, 1, 24, 1, 25, 1, 25, 1, 26, 1, 26, 1, 27, 1, 27, 1, 28, 1, 28,
		1, 29, 1, 29, 1, 30, 1, 30, 1, 31, 1, 31, 1, 32, 1, 32, 1, 33, 1, 33, 1,
		34, 1, 34, 1, 35, 1, 35, 1, 35, 1, 35, 5, 35, 214, 8, 35, 10, 35, 12, 35,
		217, 9, 35, 1, 35, 1, 35, 1, 36, 4, 36, 222, 8, 36, 11, 36, 12, 36, 223,
		1, 36, 1, 36, 0, 0, 37, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15,
		8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17,
		35, 18, 37, 19, 39, 20, 41, 21, 43, 22, 45, 23, 47, 24, 49, 25, 51, 26,
		53, 27, 55, 28, 57, 29, 59, 30, 61, 31, 63, 32, 65, 33, 67, 34, 69, 35,
		71, 36, 73, 37, 1, 0, 6, 2, 0, 65, 90, 97, 122, 3, 0, 48, 57, 65, 90, 97,
		122, 1, 0, 48, 57, 1, 0, 34, 34, 2, 0, 10, 10, 13, 13, 3, 0, 9, 10, 13,
		13, 32, 32, 233, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0,
		0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0,
		0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0,
		0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1,
		0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37,
		1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0,
		45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0,
		0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 57, 1, 0, 0, 0, 0, 59, 1, 0, 0,
		0, 0, 61, 1, 0, 0, 0, 0, 63, 1, 0, 0, 0, 0, 65, 1, 0, 0, 0, 0, 67, 1, 0,
		0, 0, 0, 69, 1, 0, 0, 0, 0, 71, 1, 0, 0, 0, 0, 73, 1, 0, 0, 0, 1, 75, 1,
		0, 0, 0, 3, 79, 1, 0, 0, 0, 5, 85, 1, 0, 0, 0, 7, 90, 1, 0, 0, 0, 9, 97,
		1, 0, 0, 0, 11, 102, 1, 0, 0, 0, 13, 108, 1, 0, 0, 0, 15, 114, 1, 0, 0,
		0, 17, 119, 1, 0, 0, 0, 19, 122, 1, 0, 0, 0, 21, 127, 1, 0, 0, 0, 23, 133,
		1, 0, 0, 0, 25, 141, 1, 0, 0, 0, 27, 146, 1, 0, 0, 0, 29, 156, 1, 0, 0,
		0, 31, 165, 1, 0, 0, 0, 33, 167, 1, 0, 0, 0, 35, 170, 1, 0, 0, 0, 37, 173,
		1, 0, 0, 0, 39, 176, 1, 0, 0, 0, 41, 179, 1, 0, 0, 0, 43, 181, 1, 0, 0,
		0, 45, 183, 1, 0, 0, 0, 47, 185, 1, 0, 0, 0, 49, 187, 1, 0, 0, 0, 51, 189,
		1, 0, 0, 0, 53, 191, 1, 0, 0, 0, 55, 193, 1, 0, 0, 0, 57, 195, 1, 0, 0,
		0, 59, 197, 1, 0, 0, 0, 61, 199, 1, 0, 0, 0, 63, 201, 1, 0, 0, 0, 65, 203,
		1, 0, 0, 0, 67, 205, 1, 0, 0, 0, 69, 207, 1, 0, 0, 0, 71, 209, 1, 0, 0,
		0, 73, 221, 1, 0, 0, 0, 75, 76, 5, 105, 0, 0, 76, 77, 5, 110, 0, 0, 77,
		78, 5, 116, 0, 0, 78, 2, 1, 0, 0, 0, 79, 80, 5, 102, 0, 0, 80, 81, 5, 108,
		0, 0, 81, 82, 5, 111, 0, 0, 82, 83, 5, 97, 0, 0, 83, 84, 5, 116, 0, 0,
		84, 4, 1, 0, 0, 0, 85, 86, 5, 98, 0, 0, 86, 87, 5, 111, 0, 0, 87, 88, 5,
		111, 0, 0, 88, 89, 5, 108, 0, 0, 89, 6, 1, 0, 0, 0, 90, 91, 5, 115, 0,
		0, 91, 92, 5, 116, 0, 0, 92, 93, 5, 114, 0, 0, 93, 94, 5, 105, 0, 0, 94,
		95, 5, 110, 0, 0, 95, 96, 5, 103, 0, 0, 96, 8, 1, 0, 0, 0, 97, 98, 5, 116,
		0, 0, 98, 99, 5, 114, 0, 0, 99, 100, 5, 117, 0, 0, 100, 101, 5, 101, 0,
		0, 101, 10, 1, 0, 0, 0, 102, 103, 5, 102, 0, 0, 103, 104, 5, 97, 0, 0,
		104, 105, 5, 108, 0, 0, 105, 106, 5, 115, 0, 0, 106, 107, 5, 101, 0, 0,
		107, 12, 1, 0, 0, 0, 108, 109, 5, 119, 0, 0, 109, 110, 5, 114, 0, 0, 110,
		111, 5, 105, 0, 0, 111, 112, 5, 116, 0, 0, 112, 113, 5, 101, 0, 0, 113,
		14, 1, 0, 0, 0, 114, 115, 5, 114, 0, 0, 115, 116, 5, 101, 0, 0, 116, 117,
		5, 97, 0, 0, 117, 118, 5, 100, 0, 0, 118, 16, 1, 0, 0, 0, 119, 120, 5,
		105, 0, 0, 120, 121, 5, 102, 0, 0, 121, 18, 1, 0, 0, 0, 122, 123, 5, 101,
		0, 0, 123, 124, 5, 108, 0, 0, 124, 125, 5, 115, 0, 0, 125, 126, 5, 101,
		0, 0, 126, 20, 1, 0, 0, 0, 127, 128, 5, 119, 0, 0, 128, 129, 5, 104, 0,
		0, 129, 130, 5, 105, 0, 0, 130, 131, 5, 108, 0, 0, 131, 132, 5, 101, 0,
		0, 132, 22, 1, 0, 0, 0, 133, 137, 7, 0, 0, 0, 134, 136, 7, 1, 0, 0, 135,
		134, 1, 0, 0, 0, 136, 139, 1, 0, 0, 0, 137, 135, 1, 0, 0, 0, 137, 138,
		1, 0, 0, 0, 138, 24, 1, 0, 0, 0, 139, 137, 1, 0, 0, 0, 140, 142, 7, 2,
		0, 0, 141, 140, 1, 0, 0, 0, 142, 143, 1, 0, 0, 0, 143, 141, 1, 0, 0, 0,
		143, 144, 1, 0, 0, 0, 144, 26, 1, 0, 0, 0, 145, 147, 7, 2, 0, 0, 146, 145,
		1, 0, 0, 0, 147, 148, 1, 0, 0, 0, 148, 146, 1, 0, 0, 0, 148, 149, 1, 0,
		0, 0, 149, 150, 1, 0, 0, 0, 150, 152, 5, 46, 0, 0, 151, 153, 7, 2, 0, 0,
		152, 151, 1, 0, 0, 0, 153, 154, 1, 0, 0, 0, 154, 152, 1, 0, 0, 0, 154,
		155, 1, 0, 0, 0, 155, 28, 1, 0, 0, 0, 156, 160, 5, 34, 0, 0, 157, 159,
		8, 3, 0, 0, 158, 157, 1, 0, 0, 0, 159, 162, 1, 0, 0, 0, 160, 158, 1, 0,
		0, 0, 160, 161, 1, 0, 0, 0, 161, 163, 1, 0, 0, 0, 162, 160, 1, 0, 0, 0,
		163, 164, 5, 34, 0, 0, 164, 30, 1, 0, 0, 0, 165, 166, 5, 61, 0, 0, 166,
		32, 1, 0, 0, 0, 167, 168, 5, 124, 0, 0, 168, 169, 5, 124, 0, 0, 169, 34,
		1, 0, 0, 0, 170, 171, 5, 38, 0, 0, 171, 172, 5, 38, 0, 0, 172, 36, 1, 0,
		0, 0, 173, 174, 5, 61, 0, 0, 174, 175, 5, 61, 0, 0, 175, 38, 1, 0, 0, 0,
		176, 177, 5, 33, 0, 0, 177, 178, 5, 61, 0, 0, 178, 40, 1, 0, 0, 0, 179,
		180, 5, 60, 0, 0, 180, 42, 1, 0, 0, 0, 181, 182, 5, 62, 0, 0, 182, 44,
		1, 0, 0, 0, 183, 184, 5, 43, 0, 0, 184, 46, 1, 0, 0, 0, 185, 186, 5, 45,
		0, 0, 186, 48, 1, 0, 0, 0, 187, 188, 5, 42, 0, 0, 188, 50, 1, 0, 0, 0,
		189, 190, 5, 47, 0, 0, 190, 52, 1, 0, 0, 0, 191, 192, 5, 37, 0, 0, 192,
		54, 1, 0, 0, 0, 193, 194, 5, 33, 0, 0, 194, 56, 1, 0, 0, 0, 195, 196, 5,
		59, 0, 0, 196, 58, 1, 0, 0, 0, 197, 198, 5, 40, 0, 0, 198, 60, 1, 0, 0,
		0, 199, 200, 5, 41, 0, 0, 200, 62, 1, 0, 0, 0, 201, 202, 5, 123, 0, 0,
		202, 64, 1, 0, 0, 0, 203, 204, 5, 125, 0, 0, 204, 66, 1, 0, 0, 0, 205,
		206, 5, 44, 0, 0, 206, 68, 1, 0, 0, 0, 207, 208, 5, 46, 0, 0, 208, 70,
		1, 0, 0, 0, 209, 210, 5, 47, 0, 0, 210, 211, 5, 47, 0, 0, 211, 215, 1,
		0, 0, 0, 212, 214, 8, 4, 0, 0, 213, 212, 1, 0, 0, 0, 214, 217, 1, 0, 0,
		0, 215, 213, 1, 0, 0, 0, 215, 216, 1, 0, 0, 0, 216, 218, 1, 0, 0, 0, 217,
		215, 1, 0, 0, 0, 218, 219, 6, 35, 0, 0, 219, 72, 1, 0, 0, 0, 220, 222,
		7, 5, 0, 0, 221, 220, 1, 0, 0, 0, 222, 223, 1, 0, 0, 0, 223, 221, 1, 0,
		0, 0, 223, 224, 1, 0, 0, 0, 224, 225, 1, 0, 0, 0, 225, 226, 6, 36, 0, 0,
		226, 74, 1, 0, 0, 0, 8, 0, 137, 143, 148, 154, 160, 215, 223, 1, 6, 0,
		0,
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

// StackLexerInit initializes any static state used to implement StackLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewStackLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func StackLexerInit() {
	staticData := &StackLexerLexerStaticData
	staticData.once.Do(stacklexerLexerInit)
}

// NewStackLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewStackLexer(input antlr.CharStream) *StackLexer {
	StackLexerInit()
	l := new(StackLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &StackLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "Stack.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// StackLexer tokens.
const (
	StackLexerINT            = 1
	StackLexerFLOAT          = 2
	StackLexerBOOL           = 3
	StackLexerSTRING         = 4
	StackLexerTRUE           = 5
	StackLexerFALSE          = 6
	StackLexerWRITE          = 7
	StackLexerREAD           = 8
	StackLexerIF             = 9
	StackLexerELSE           = 10
	StackLexerWHILE          = 11
	StackLexerID             = 12
	StackLexerINT_LITERAL    = 13
	StackLexerFLOAT_LITERAL  = 14
	StackLexerSTRING_LITERAL = 15
	StackLexerASSIGN         = 16
	StackLexerOR             = 17
	StackLexerAND            = 18
	StackLexerEQ             = 19
	StackLexerNE             = 20
	StackLexerLT             = 21
	StackLexerGT             = 22
	StackLexerADD            = 23
	StackLexerSUB            = 24
	StackLexerMUL            = 25
	StackLexerDIV            = 26
	StackLexerMOD            = 27
	StackLexerNOT            = 28
	StackLexerSEM            = 29
	StackLexerLPAR           = 30
	StackLexerRPAR           = 31
	StackLexerLBRACE         = 32
	StackLexerRBRACE         = 33
	StackLexerCOMMA          = 34
	StackLexerDOT            = 35
	StackLexerCOMMENT        = 36
	StackLexerWS             = 37
)