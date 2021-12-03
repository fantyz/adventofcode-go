package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBitIsSet(t *testing.T) {
	testCases := []struct {
		Val   uint64
		IsSet []bool
	}{
		{5, []bool{true, false, true, false, false, false}},
	}

	for _, c := range testCases {
		for i := range c.IsSet {
			assert.Equal(t, c.IsSet[i], bitIsSet(i, c.Val), "val=%d, bit=%d", c.Val, i)
		}
	}
}

func TestCalculateGammaAndEpsilon(t *testing.T) {
	input := `00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010`
	r, err := NewDiagnosticsReport(input)
	if assert.NoError(t, err) {
		g, e := r.CalculateGammaAndEpsilon()
		assert.Equal(t, 22, g, "gamma")
		assert.Equal(t, 9, e, "epsilon")
	}
}

func TestGetOxygenGeneratorRating(t *testing.T) {
	input := `00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010`
	r, err := NewDiagnosticsReport(input)
	if assert.NoError(t, err) {
		assert.Equal(t, 23, r.GetOxygenGeneratorRating())
	}
}

func TestGetCO2ScrubberRating(t *testing.T) {
	input := `00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010`
	r, err := NewDiagnosticsReport(input)
	if assert.NoError(t, err) {
		assert.Equal(t, 10, r.GetCO2ScrubberRating())
	}
}

func TestDay03Pt1(t *testing.T) {
	r, err := NewDiagnosticsReport(day03Input)
	if assert.NoError(t, err) {
		g, e := r.CalculateGammaAndEpsilon()
		assert.Equal(t, 4147524, g*e)
	}
}

func TestDay03Pt2(t *testing.T) {
	r, err := NewDiagnosticsReport(day03Input)
	if assert.NoError(t, err) {
		assert.Equal(t, 3570354, r.GetOxygenGeneratorRating()*r.GetCO2ScrubberRating())
	}
}
