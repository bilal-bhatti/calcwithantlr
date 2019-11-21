// Code generated from Mixed.g by ANTLR 4.7.2. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 9, 36, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9,
	7, 4, 8, 9, 8, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3,
	6, 3, 7, 6, 7, 29, 10, 7, 13, 7, 14, 7, 30, 3, 8, 3, 8, 3, 8, 3, 8, 2,
	2, 9, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 3, 2, 4, 3, 2, 50, 59,
	5, 2, 11, 12, 15, 15, 34, 34, 2, 36, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2,
	2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2,
	2, 2, 15, 3, 2, 2, 2, 3, 17, 3, 2, 2, 2, 5, 19, 3, 2, 2, 2, 7, 21, 3, 2,
	2, 2, 9, 23, 3, 2, 2, 2, 11, 25, 3, 2, 2, 2, 13, 28, 3, 2, 2, 2, 15, 32,
	3, 2, 2, 2, 17, 18, 7, 97, 2, 2, 18, 4, 3, 2, 2, 2, 19, 20, 7, 45, 2, 2,
	20, 6, 3, 2, 2, 2, 21, 22, 7, 47, 2, 2, 22, 8, 3, 2, 2, 2, 23, 24, 7, 44,
	2, 2, 24, 10, 3, 2, 2, 2, 25, 26, 7, 49, 2, 2, 26, 12, 3, 2, 2, 2, 27,
	29, 9, 2, 2, 2, 28, 27, 3, 2, 2, 2, 29, 30, 3, 2, 2, 2, 30, 28, 3, 2, 2,
	2, 30, 31, 3, 2, 2, 2, 31, 14, 3, 2, 2, 2, 32, 33, 9, 3, 2, 2, 33, 34,
	3, 2, 2, 2, 34, 35, 8, 8, 2, 2, 35, 16, 3, 2, 2, 2, 4, 2, 30, 3, 8, 2,
	2,
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
	"", "'_'", "'+'", "'-'", "'*'", "'/'",
}

var lexerSymbolicNames = []string{
	"", "", "OP_ADD", "OP_SUB", "OP_MUL", "OP_DIV", "NUM", "WS",
}

var lexerRuleNames = []string{
	"T__0", "OP_ADD", "OP_SUB", "OP_MUL", "OP_DIV", "NUM", "WS",
}

type MixedLexer struct {
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

func NewMixedLexer(input antlr.CharStream) *MixedLexer {

	l := new(MixedLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Mixed.g"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// MixedLexer tokens.
const (
	MixedLexerT__0   = 1
	MixedLexerOP_ADD = 2
	MixedLexerOP_SUB = 3
	MixedLexerOP_MUL = 4
	MixedLexerOP_DIV = 5
	MixedLexerNUM    = 6
	MixedLexerWS     = 7
)
