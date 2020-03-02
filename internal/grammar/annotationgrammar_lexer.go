// Code generated from internal/grammar/AnnotationGrammar.g4 by ANTLR 4.8. DO NOT EDIT.

package grammar

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 11, 68, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9,
	7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4,
	3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 7, 6, 7, 33, 10, 7, 13, 7, 14, 7, 34,
	3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 5, 8, 45, 10, 8, 3, 8,
	7, 8, 48, 10, 8, 12, 8, 14, 8, 51, 11, 8, 3, 9, 7, 9, 54, 10, 9, 12, 9,
	14, 9, 57, 11, 9, 3, 9, 6, 9, 60, 10, 9, 13, 9, 14, 9, 61, 3, 10, 6, 10,
	65, 10, 10, 13, 10, 14, 10, 66, 2, 2, 11, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7,
	13, 8, 15, 9, 17, 10, 19, 11, 3, 2, 4, 7, 2, 11, 12, 34, 34, 42, 44, 93,
	93, 95, 95, 4, 2, 11, 11, 34, 34, 2, 74, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2,
	2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3,
	2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 3, 21,
	3, 2, 2, 2, 5, 23, 3, 2, 2, 2, 7, 25, 3, 2, 2, 2, 9, 27, 3, 2, 2, 2, 11,
	29, 3, 2, 2, 2, 13, 32, 3, 2, 2, 2, 15, 36, 3, 2, 2, 2, 17, 55, 3, 2, 2,
	2, 19, 64, 3, 2, 2, 2, 21, 22, 7, 44, 2, 2, 22, 4, 3, 2, 2, 2, 23, 24,
	7, 42, 2, 2, 24, 6, 3, 2, 2, 2, 25, 26, 7, 43, 2, 2, 26, 8, 3, 2, 2, 2,
	27, 28, 7, 93, 2, 2, 28, 10, 3, 2, 2, 2, 29, 30, 7, 95, 2, 2, 30, 12, 3,
	2, 2, 2, 31, 33, 10, 2, 2, 2, 32, 31, 3, 2, 2, 2, 33, 34, 3, 2, 2, 2, 34,
	32, 3, 2, 2, 2, 34, 35, 3, 2, 2, 2, 35, 14, 3, 2, 2, 2, 36, 44, 5, 17,
	9, 2, 37, 38, 7, 34, 2, 2, 38, 45, 7, 34, 2, 2, 39, 45, 7, 11, 2, 2, 40,
	41, 7, 34, 2, 2, 41, 42, 7, 34, 2, 2, 42, 43, 7, 34, 2, 2, 43, 45, 7, 34,
	2, 2, 44, 37, 3, 2, 2, 2, 44, 39, 3, 2, 2, 2, 44, 40, 3, 2, 2, 2, 45, 49,
	3, 2, 2, 2, 46, 48, 5, 19, 10, 2, 47, 46, 3, 2, 2, 2, 48, 51, 3, 2, 2,
	2, 49, 47, 3, 2, 2, 2, 49, 50, 3, 2, 2, 2, 50, 16, 3, 2, 2, 2, 51, 49,
	3, 2, 2, 2, 52, 54, 5, 19, 10, 2, 53, 52, 3, 2, 2, 2, 54, 57, 3, 2, 2,
	2, 55, 53, 3, 2, 2, 2, 55, 56, 3, 2, 2, 2, 56, 59, 3, 2, 2, 2, 57, 55,
	3, 2, 2, 2, 58, 60, 7, 12, 2, 2, 59, 58, 3, 2, 2, 2, 60, 61, 3, 2, 2, 2,
	61, 59, 3, 2, 2, 2, 61, 62, 3, 2, 2, 2, 62, 18, 3, 2, 2, 2, 63, 65, 9,
	3, 2, 2, 64, 63, 3, 2, 2, 2, 65, 66, 3, 2, 2, 2, 66, 64, 3, 2, 2, 2, 66,
	67, 3, 2, 2, 2, 67, 20, 3, 2, 2, 2, 9, 2, 34, 44, 49, 55, 61, 66, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'*'", "'('", "')'", "'['", "']'",
}

var lexerSymbolicNames = []string{
	"", "INTENT_NAME_START", "OPEN_PAREN", "CLOSE_PAREN", "OPEN_SB", "CLOSE_SB",
	"WORD", "INDENT", "END", "WHITESPACE",
}

var lexerRuleNames = []string{
	"INTENT_NAME_START", "OPEN_PAREN", "CLOSE_PAREN", "OPEN_SB", "CLOSE_SB",
	"WORD", "INDENT", "END", "WHITESPACE",
}

type AnnotationGrammarLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewAnnotationGrammarLexer(input antlr.CharStream) *AnnotationGrammarLexer {

	l := new(AnnotationGrammarLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "AnnotationGrammar.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// AnnotationGrammarLexer tokens.
const (
	AnnotationGrammarLexerINTENT_NAME_START = 1
	AnnotationGrammarLexerOPEN_PAREN        = 2
	AnnotationGrammarLexerCLOSE_PAREN       = 3
	AnnotationGrammarLexerOPEN_SB           = 4
	AnnotationGrammarLexerCLOSE_SB          = 5
	AnnotationGrammarLexerWORD              = 6
	AnnotationGrammarLexerINDENT            = 7
	AnnotationGrammarLexerEND               = 8
	AnnotationGrammarLexerWHITESPACE        = 9
)