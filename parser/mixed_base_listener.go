// Code generated from Mixed.g by ANTLR 4.7.2. DO NOT EDIT.

package parser // Mixed

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseMixedListener is a complete listener for a parse tree produced by MixedParser.
type BaseMixedListener struct{}

var _ MixedListener = &BaseMixedListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseMixedListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseMixedListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseMixedListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseMixedListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterStart is called when production start is entered.
func (s *BaseMixedListener) EnterStart(ctx *StartContext) {}

// ExitStart is called when production start is exited.
func (s *BaseMixedListener) ExitStart(ctx *StartContext) {}

// EnterNumber is called when production Number is entered.
func (s *BaseMixedListener) EnterNumber(ctx *NumberContext) {}

// ExitNumber is called when production Number is exited.
func (s *BaseMixedListener) ExitNumber(ctx *NumberContext) {}

// EnterMixed is called when production Mixed is entered.
func (s *BaseMixedListener) EnterMixed(ctx *MixedContext) {}

// ExitMixed is called when production Mixed is exited.
func (s *BaseMixedListener) ExitMixed(ctx *MixedContext) {}

// EnterMulDiv is called when production MulDiv is entered.
func (s *BaseMixedListener) EnterMulDiv(ctx *MulDivContext) {}

// ExitMulDiv is called when production MulDiv is exited.
func (s *BaseMixedListener) ExitMulDiv(ctx *MulDivContext) {}

// EnterAddSub is called when production AddSub is entered.
func (s *BaseMixedListener) EnterAddSub(ctx *AddSubContext) {}

// ExitAddSub is called when production AddSub is exited.
func (s *BaseMixedListener) ExitAddSub(ctx *AddSubContext) {}

// EnterFraction is called when production Fraction is entered.
func (s *BaseMixedListener) EnterFraction(ctx *FractionContext) {}

// ExitFraction is called when production Fraction is exited.
func (s *BaseMixedListener) ExitFraction(ctx *FractionContext) {}
