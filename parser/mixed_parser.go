// Code generated from Mixed.g by ANTLR 4.7.2. DO NOT EDIT.

package parser // Mixed

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 9, 33, 4,
	2, 9, 2, 4, 3, 9, 3, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 20, 10, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 7, 3, 28, 10, 3, 12, 3, 14, 3, 31, 11, 3, 3, 3, 2, 3, 4, 4, 2,
	4, 2, 4, 3, 2, 6, 7, 3, 2, 4, 5, 2, 34, 2, 6, 3, 2, 2, 2, 4, 19, 3, 2,
	2, 2, 6, 7, 5, 4, 3, 2, 7, 8, 7, 2, 2, 3, 8, 3, 3, 2, 2, 2, 9, 10, 8, 3,
	1, 2, 10, 11, 7, 8, 2, 2, 11, 12, 7, 7, 2, 2, 12, 20, 7, 8, 2, 2, 13, 14,
	7, 8, 2, 2, 14, 15, 7, 3, 2, 2, 15, 16, 7, 8, 2, 2, 16, 17, 7, 7, 2, 2,
	17, 20, 7, 8, 2, 2, 18, 20, 7, 8, 2, 2, 19, 9, 3, 2, 2, 2, 19, 13, 3, 2,
	2, 2, 19, 18, 3, 2, 2, 2, 20, 29, 3, 2, 2, 2, 21, 22, 12, 7, 2, 2, 22,
	23, 9, 2, 2, 2, 23, 28, 5, 4, 3, 8, 24, 25, 12, 6, 2, 2, 25, 26, 9, 3,
	2, 2, 26, 28, 5, 4, 3, 7, 27, 21, 3, 2, 2, 2, 27, 24, 3, 2, 2, 2, 28, 31,
	3, 2, 2, 2, 29, 27, 3, 2, 2, 2, 29, 30, 3, 2, 2, 2, 30, 5, 3, 2, 2, 2,
	31, 29, 3, 2, 2, 2, 5, 19, 27, 29,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'_'", "'+'", "'-'", "'*'", "'/'",
}
var symbolicNames = []string{
	"", "", "OP_ADD", "OP_SUB", "OP_MUL", "OP_DIV", "NUM", "WS",
}

var ruleNames = []string{
	"start", "expr",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type MixedParser struct {
	*antlr.BaseParser
}

func NewMixedParser(input antlr.TokenStream) *MixedParser {
	this := new(MixedParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Mixed.g"

	return this
}

// MixedParser tokens.
const (
	MixedParserEOF    = antlr.TokenEOF
	MixedParserT__0   = 1
	MixedParserOP_ADD = 2
	MixedParserOP_SUB = 3
	MixedParserOP_MUL = 4
	MixedParserOP_DIV = 5
	MixedParserNUM    = 6
	MixedParserWS     = 7
)

// MixedParser rules.
const (
	MixedParserRULE_start = 0
	MixedParserRULE_expr  = 1
)

// IStartContext is an interface to support dynamic dispatch.
type IStartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStartContext differentiates from other interfaces.
	IsStartContext()
}

type StartContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStartContext() *StartContext {
	var p = new(StartContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MixedParserRULE_start
	return p
}

func (*StartContext) IsStartContext() {}

func NewStartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StartContext {
	var p = new(StartContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MixedParserRULE_start

	return p
}

func (s *StartContext) GetParser() antlr.Parser { return s.parser }

func (s *StartContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *StartContext) EOF() antlr.TerminalNode {
	return s.GetToken(MixedParserEOF, 0)
}

func (s *StartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StartContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MixedListener); ok {
		listenerT.EnterStart(s)
	}
}

func (s *StartContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MixedListener); ok {
		listenerT.ExitStart(s)
	}
}

func (p *MixedParser) Start() (localctx IStartContext) {
	localctx = NewStartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, MixedParserRULE_start)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(4)
		p.expr(0)
	}
	{
		p.SetState(5)
		p.Match(MixedParserEOF)
	}

	return localctx
}

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MixedParserRULE_expr
	return p
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MixedParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) CopyFrom(ctx *ExprContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type NumberContext struct {
	*ExprContext
	value antlr.Token
}

func NewNumberContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NumberContext {
	var p = new(NumberContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *NumberContext) GetValue() antlr.Token { return s.value }

func (s *NumberContext) SetValue(v antlr.Token) { s.value = v }

func (s *NumberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberContext) NUM() antlr.TerminalNode {
	return s.GetToken(MixedParserNUM, 0)
}

func (s *NumberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MixedListener); ok {
		listenerT.EnterNumber(s)
	}
}

func (s *NumberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MixedListener); ok {
		listenerT.ExitNumber(s)
	}
}

type MixedContext struct {
	*ExprContext
	whole       antlr.Token
	numerator   antlr.Token
	denominator antlr.Token
}

func NewMixedContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MixedContext {
	var p = new(MixedContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *MixedContext) GetWhole() antlr.Token { return s.whole }

func (s *MixedContext) GetNumerator() antlr.Token { return s.numerator }

func (s *MixedContext) GetDenominator() antlr.Token { return s.denominator }

func (s *MixedContext) SetWhole(v antlr.Token) { s.whole = v }

func (s *MixedContext) SetNumerator(v antlr.Token) { s.numerator = v }

func (s *MixedContext) SetDenominator(v antlr.Token) { s.denominator = v }

func (s *MixedContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MixedContext) OP_DIV() antlr.TerminalNode {
	return s.GetToken(MixedParserOP_DIV, 0)
}

func (s *MixedContext) AllNUM() []antlr.TerminalNode {
	return s.GetTokens(MixedParserNUM)
}

func (s *MixedContext) NUM(i int) antlr.TerminalNode {
	return s.GetToken(MixedParserNUM, i)
}

func (s *MixedContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MixedListener); ok {
		listenerT.EnterMixed(s)
	}
}

func (s *MixedContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MixedListener); ok {
		listenerT.ExitMixed(s)
	}
}

type MulDivContext struct {
	*ExprContext
	left  IExprContext
	op    antlr.Token
	right IExprContext
}

func NewMulDivContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MulDivContext {
	var p = new(MulDivContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *MulDivContext) GetOp() antlr.Token { return s.op }

func (s *MulDivContext) SetOp(v antlr.Token) { s.op = v }

func (s *MulDivContext) GetLeft() IExprContext { return s.left }

func (s *MulDivContext) GetRight() IExprContext { return s.right }

func (s *MulDivContext) SetLeft(v IExprContext) { s.left = v }

func (s *MulDivContext) SetRight(v IExprContext) { s.right = v }

func (s *MulDivContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MulDivContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *MulDivContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *MulDivContext) OP_MUL() antlr.TerminalNode {
	return s.GetToken(MixedParserOP_MUL, 0)
}

func (s *MulDivContext) OP_DIV() antlr.TerminalNode {
	return s.GetToken(MixedParserOP_DIV, 0)
}

func (s *MulDivContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MixedListener); ok {
		listenerT.EnterMulDiv(s)
	}
}

func (s *MulDivContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MixedListener); ok {
		listenerT.ExitMulDiv(s)
	}
}

type AddSubContext struct {
	*ExprContext
	left  IExprContext
	op    antlr.Token
	right IExprContext
}

func NewAddSubContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AddSubContext {
	var p = new(AddSubContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *AddSubContext) GetOp() antlr.Token { return s.op }

func (s *AddSubContext) SetOp(v antlr.Token) { s.op = v }

func (s *AddSubContext) GetLeft() IExprContext { return s.left }

func (s *AddSubContext) GetRight() IExprContext { return s.right }

func (s *AddSubContext) SetLeft(v IExprContext) { s.left = v }

func (s *AddSubContext) SetRight(v IExprContext) { s.right = v }

func (s *AddSubContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AddSubContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *AddSubContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AddSubContext) OP_ADD() antlr.TerminalNode {
	return s.GetToken(MixedParserOP_ADD, 0)
}

func (s *AddSubContext) OP_SUB() antlr.TerminalNode {
	return s.GetToken(MixedParserOP_SUB, 0)
}

func (s *AddSubContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MixedListener); ok {
		listenerT.EnterAddSub(s)
	}
}

func (s *AddSubContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MixedListener); ok {
		listenerT.ExitAddSub(s)
	}
}

type FractionContext struct {
	*ExprContext
	numerator   antlr.Token
	denominator antlr.Token
}

func NewFractionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FractionContext {
	var p = new(FractionContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *FractionContext) GetNumerator() antlr.Token { return s.numerator }

func (s *FractionContext) GetDenominator() antlr.Token { return s.denominator }

func (s *FractionContext) SetNumerator(v antlr.Token) { s.numerator = v }

func (s *FractionContext) SetDenominator(v antlr.Token) { s.denominator = v }

func (s *FractionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FractionContext) OP_DIV() antlr.TerminalNode {
	return s.GetToken(MixedParserOP_DIV, 0)
}

func (s *FractionContext) AllNUM() []antlr.TerminalNode {
	return s.GetTokens(MixedParserNUM)
}

func (s *FractionContext) NUM(i int) antlr.TerminalNode {
	return s.GetToken(MixedParserNUM, i)
}

func (s *FractionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MixedListener); ok {
		listenerT.EnterFraction(s)
	}
}

func (s *FractionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MixedListener); ok {
		listenerT.ExitFraction(s)
	}
}

func (p *MixedParser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *MixedParser) expr(_p int) (localctx IExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 2
	p.EnterRecursionRule(localctx, 2, MixedParserRULE_expr, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(17)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext()) {
	case 1:
		localctx = NewFractionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(8)

			var _m = p.Match(MixedParserNUM)

			localctx.(*FractionContext).numerator = _m
		}
		{
			p.SetState(9)
			p.Match(MixedParserOP_DIV)
		}
		{
			p.SetState(10)

			var _m = p.Match(MixedParserNUM)

			localctx.(*FractionContext).denominator = _m
		}

	case 2:
		localctx = NewMixedContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(11)

			var _m = p.Match(MixedParserNUM)

			localctx.(*MixedContext).whole = _m
		}
		{
			p.SetState(12)
			p.Match(MixedParserT__0)
		}
		{
			p.SetState(13)

			var _m = p.Match(MixedParserNUM)

			localctx.(*MixedContext).numerator = _m
		}
		{
			p.SetState(14)
			p.Match(MixedParserOP_DIV)
		}
		{
			p.SetState(15)

			var _m = p.Match(MixedParserNUM)

			localctx.(*MixedContext).denominator = _m
		}

	case 3:
		localctx = NewNumberContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(16)

			var _m = p.Match(MixedParserNUM)

			localctx.(*NumberContext).value = _m
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(27)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(25)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
			case 1:
				localctx = NewMulDivContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*MulDivContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, MixedParserRULE_expr)
				p.SetState(19)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(20)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*MulDivContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == MixedParserOP_MUL || _la == MixedParserOP_DIV) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*MulDivContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(21)

					var _x = p.expr(6)

					localctx.(*MulDivContext).right = _x
				}

			case 2:
				localctx = NewAddSubContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*AddSubContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, MixedParserRULE_expr)
				p.SetState(22)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(23)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*AddSubContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == MixedParserOP_ADD || _la == MixedParserOP_SUB) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*AddSubContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(24)

					var _x = p.expr(5)

					localctx.(*AddSubContext).right = _x
				}

			}

		}
		p.SetState(29)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())
	}

	return localctx
}

func (p *MixedParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 1:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *MixedParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 4)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
