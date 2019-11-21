// Code generated from Mixed.g by ANTLR 4.7.2. DO NOT EDIT.

package parser // Mixed

import "github.com/antlr/antlr4/runtime/Go/antlr"

// MixedListener is a complete listener for a parse tree produced by MixedParser.
type MixedListener interface {
	antlr.ParseTreeListener

	// EnterStart is called when entering the start production.
	EnterStart(c *StartContext)

	// EnterNumber is called when entering the Number production.
	EnterNumber(c *NumberContext)

	// EnterMixed is called when entering the Mixed production.
	EnterMixed(c *MixedContext)

	// EnterMulDiv is called when entering the MulDiv production.
	EnterMulDiv(c *MulDivContext)

	// EnterAddSub is called when entering the AddSub production.
	EnterAddSub(c *AddSubContext)

	// EnterFraction is called when entering the Fraction production.
	EnterFraction(c *FractionContext)

	// ExitStart is called when exiting the start production.
	ExitStart(c *StartContext)

	// ExitNumber is called when exiting the Number production.
	ExitNumber(c *NumberContext)

	// ExitMixed is called when exiting the Mixed production.
	ExitMixed(c *MixedContext)

	// ExitMulDiv is called when exiting the MulDiv production.
	ExitMulDiv(c *MulDivContext)

	// ExitAddSub is called when exiting the AddSub production.
	ExitAddSub(c *AddSubContext)

	// ExitFraction is called when exiting the Fraction production.
	ExitFraction(c *FractionContext)
}
