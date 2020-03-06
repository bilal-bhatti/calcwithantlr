package calc

import (
	"fmt"

	"github.com/bilal-bhatti/calcwithantlr/parser"
)

// Listener ...
type Listener struct {
	*parser.BaseMixedListener

	stack []*Numby
}

// Push ...
func (l *Listener) push(i *Numby) {
	l.stack = append(l.stack, i)
}

// Pop ...
func (l *Listener) Pop() *Numby {
	if len(l.stack) < 1 {
		panic("No input")
	}

	// Get the last value from the stack.
	result := l.stack[len(l.stack)-1]

	// Remove the last element from the stack.
	l.stack = l.stack[:len(l.stack)-1]

	return result
}

// ExitMulDiv ...
func (l *Listener) ExitMulDiv(c *parser.MulDivContext) {
	right, left := l.Pop(), l.Pop()

	switch c.GetOp().GetTokenType() {
	case parser.MixedParserOP_MUL:
		l.push(left.Mul(right))
	case parser.MixedParserOP_DIV:
		l.push(left.Div(right))
	default:
		panic(fmt.Sprintf("unexpected op: %s", c.GetOp().GetText()))
	}
}

// ExitAddSub ...
func (l *Listener) ExitAddSub(c *parser.AddSubContext) {
	right, left := l.Pop(), l.Pop()

	switch c.GetOp().GetTokenType() {
	case parser.MixedLexerOP_ADD:
		l.push(left.Add(right))
	case parser.MixedLexerOP_SUB:
		l.push(left.Sub(right))
	default:
		panic(fmt.Sprintf("unexpected op: %s", c.GetOp().GetText()))
	}
}

// ExitNumber ...
func (l *Listener) ExitNumber(c *parser.NumberContext) {
	n := NewFromNumber(c.GetText())
	l.push(n)
}

// ExitFraction ...
func (l *Listener) ExitFraction(c *parser.FractionContext) {
	n := NewFromFraction(c.GetNumerator().GetText(), c.GetDenominator().GetText())
	l.push(n)
}

// ExitMixed ...
func (l *Listener) ExitMixed(c *parser.MixedContext) {
	n := NewFromMixed(c.GetWhole().GetText(), c.GetNumerator().GetText(), c.GetDenominator().GetText())
	l.push(n)
}
