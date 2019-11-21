package calc

import (
	"fmt"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testcase struct {
	left, right            *Numby
	numerator, denominator int64
	simple                 string
}

func TestAdd(t *testing.T) {
	tests := []testcase{
		{
			left:        NewFromNumber("2"),
			right:       NewFromNumber("2"),
			numerator:   4,
			denominator: 1,
			simple:      "4",
		},
		{
			left:        NewFromFraction("1", "2"),
			right:       NewFromFraction("1", "2"),
			numerator:   4,
			denominator: 4,
			simple:      "1",
		},
		{
			left:        NewFromFraction("1", "6"),
			right:       NewFromFraction("2", "6"),
			numerator:   18,
			denominator: 36,
			simple:      "1/2",
		},
		{
			left:        NewFromFraction("1", "3"),
			right:       NewFromFraction("1", "2"),
			numerator:   5,
			denominator: 6,
			simple:      "5/6",
		},
		{
			left:        NewFromFraction("3", "4"),
			right:       NewFromFraction("4", "5"),
			numerator:   31,
			denominator: 20,
			simple:      "1_11/20",
		},
		{
			left:        NewFromMixed("2", "3", "8"),
			right:       NewFromFraction("9", "8"),
			numerator:   224,
			denominator: 64,
			simple:      "3_1/2",
		},
	}

	for _, test := range tests {
		left := test.left
		right := test.right

		result := left.Add(right)
		log.Printf("Simplified %d/%d to %s ", result.numerator, result.denominator, result.Simplify())

		assert.Equal(t, test.numerator, result.numerator, fmt.Sprintf("Numerator should be equal to %d", test.numerator))
		assert.Equal(t, test.denominator, result.denominator, fmt.Sprintf("Denominator should be equal to %d", test.denominator))
		assert.Equal(t, test.simple, result.Simplify(), fmt.Sprintf("Simplified fraction should be equal to %s", test.simple))
	}
}

func TestSub(t *testing.T) {
	tests := []testcase{
		{
			left:        NewFromFraction("1", "3"),
			right:       NewFromFraction("1", "2"),
			numerator:   -1,
			denominator: 6,
			simple:      "-1/6",
		},
		{
			left:        NewFromFraction("1", "2"),
			right:       NewFromFraction("1", "3"),
			numerator:   1,
			denominator: 6,
			simple:      "1/6",
		},
		{
			left:        NewFromFraction("3", "4"),
			right:       NewFromFraction("4", "5"),
			numerator:   -1,
			denominator: 20,
			simple:      "-1/20",
		},
		{
			left:        NewFromMixed("2", "3", "8"),
			right:       NewFromFraction("9", "8"),
			numerator:   80,
			denominator: 64,
			simple:      "1_1/4",
		},
		{
			left:        NewFromMixed("2", "3", "8"),
			right:       NewFromNumber("5"),
			numerator:   -21,
			denominator: 8,
			simple:      "-21/8",
		},
	}

	for _, test := range tests {
		left := test.left
		right := test.right

		result := left.Sub(right)
		log.Printf("Simplified %d/%d to %s ", result.numerator, result.denominator, result.Simplify())

		assert.Equal(t, test.numerator, result.numerator, fmt.Sprintf("Numerator should be equal to %d", test.numerator))
		assert.Equal(t, test.denominator, result.denominator, fmt.Sprintf("Denominator should be equal to %d", test.denominator))
		assert.Equal(t, test.simple, result.Simplify(), fmt.Sprintf("Simplified fraction should be equal to %s", test.simple))
	}
}

func TestMul(t *testing.T) {
	tests := []testcase{
		{
			left:        NewFromFraction("1", "2"),
			right:       NewFromNumber("2"),
			numerator:   2,
			denominator: 2,
			simple:      "1",
		},
		{
			left:        NewFromFraction("1", "3"),
			right:       NewFromFraction("1", "2"),
			numerator:   1,
			denominator: 6,
			simple:      "1/6",
		},
		{
			left:        NewFromFraction("3", "4"),
			right:       NewFromFraction("4", "5"),
			numerator:   12,
			denominator: 20,
			simple:      "3/5",
		},
		{
			left:        NewFromFraction("3", "4"),
			right:       NewFromFraction("12", "5"),
			numerator:   36,
			denominator: 20,
			simple:      "1_4/5",
		},
		{
			left:        NewFromMixed("2", "3", "8"),
			right:       NewFromFraction("9", "8"),
			numerator:   171,
			denominator: 64,
			simple:      "2_43/64",
		},
		{
			left:        NewFromMixed("2", "3", "8"),
			right:       NewFromNumber("5"),
			numerator:   95,
			denominator: 8,
			simple:      "11_7/8",
		},
		{
			left:        NewFromFraction("1", "2"),
			right:       NewFromMixed("3", "3", "4"),
			numerator:   15,
			denominator: 8,
			simple:      "1_7/8",
		},
	}

	for _, test := range tests {
		left := test.left
		right := test.right

		result := left.Mul(right)
		log.Printf("Simplified %d/%d to %s ", result.numerator, result.denominator, result.Simplify())

		assert.Equal(t, test.numerator, result.numerator, fmt.Sprintf("Numerator should be equal to %d", test.numerator))
		assert.Equal(t, test.denominator, result.denominator, fmt.Sprintf("Denominator should be equal to %d", test.denominator))
		assert.Equal(t, test.simple, result.Simplify(), fmt.Sprintf("Simplified fraction should be equal to %s", test.simple))
	}
}

func TestDiv(t *testing.T) {
	tests := []testcase{
		{
			left:        NewFromNumber("2"),
			right:       NewFromFraction("1", "2"),
			numerator:   4,
			denominator: 1,
			simple:      "4",
		},
		{
			left:        NewFromFraction("1", "3"),
			right:       NewFromFraction("1", "2"),
			numerator:   2,
			denominator: 3,
			simple:      "2/3",
		},
		{
			left:        NewFromFraction("3", "4"),
			right:       NewFromFraction("4", "5"),
			numerator:   15,
			denominator: 16,
			simple:      "15/16",
		},
		{
			left:        NewFromMixed("2", "3", "8"),
			right:       NewFromFraction("9", "8"),
			numerator:   152,
			denominator: 72,
			simple:      "2_1/9",
		},
		{
			left:        NewFromMixed("2", "3", "8"),
			right:       NewFromNumber("5"),
			numerator:   19,
			denominator: 40,
			simple:      "19/40",
		},
	}

	for _, test := range tests {
		left := test.left
		right := test.right

		result := left.Div(right)
		log.Printf("Simplified %d/%d to %s ", result.numerator, result.denominator, result.Simplify())

		assert.Equal(t, test.numerator, result.numerator, fmt.Sprintf("Numerator should be equal to %d", test.numerator))
		assert.Equal(t, test.denominator, result.denominator, fmt.Sprintf("Denominator should be equal to %d", test.denominator))
		assert.Equal(t, test.simple, result.Simplify(), fmt.Sprintf("Simplified fraction should be equal to %s", test.simple))
	}
}
