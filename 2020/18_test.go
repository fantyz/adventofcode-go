package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEvaluate(t *testing.T) {
	testCases := []struct {
		Exp             string
		EqualPrecedence bool
		Res             int
	}{
		{"2 * 3 + (4 * 5)", true, 26},
		{"5 + (8 * 3 + 9 + 3 * 4 * 3)", true, 437},
		{"5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))", true, 12240},
		{"((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2", true, 13632},
		{"1 + (2 * 3) + (4 * (5 + 6))", false, 51},
		{"2 * 3 + (4 * 5)", false, 46},
		{"5 + (8 * 3 + 9 + 3 * 4 * 3)", false, 1445},
		{"5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))", false, 669060},
		{"((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2", false, 23340},
		{"1 + 2 * 3 + 4 * 5 + 6", false, 231},
	}

	for _, c := range testCases {
		res, err := Evaluate(c.Exp, c.EqualPrecedence)
		if assert.NoError(t, err, c.Exp) {
			assert.Equal(t, c.Res, res, c.Exp)
		}
	}
}

func TestDay18Pt1(t *testing.T) {
	sum, err := SumExpressions(day18Input, true)
	if assert.NoError(t, err) {
		assert.Equal(t, 1890866893020, sum)
	}
}

func TestDay18Pt2(t *testing.T) {
	sum, err := SumExpressions(day18Input, false)
	if assert.NoError(t, err) {
		assert.Equal(t, 34646237037193, sum)
	}
}
