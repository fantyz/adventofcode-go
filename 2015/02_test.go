package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWrapBox(t *testing.T) {
	testCases := map[string]struct {
		Length, Width, Height int
		PaperNeeded           int
		RibbonNeeded          int
	}{
		"2x3x4":  {2, 3, 4, 58, 34},
		"1x1x10": {1, 1, 10, 43, 14},
		"1x10x1": {1, 10, 1, 43, 14},
		"10x1x1": {10, 1, 1, 43, 14},
	}

	for n, c := range testCases {
		p, r := WrapBox(c.Length, c.Width, c.Height)
		assert.Equal(t, c.PaperNeeded, p, "Paper needed (case=%s)", n)
		assert.Equal(t, c.RibbonNeeded, r, "Ribbon needed (case=%s)", n)
	}
}

func TestWrapBoxes(t *testing.T) {
	testCase := `2x3x4
1x1x10`
	p, r, err := WrapBoxes(testCase)
	if assert.NoError(t, err) {
		assert.Equal(t, 101, p, "Paper needed")
		assert.Equal(t, 48, r, "Ribbon needed")
	}
}

func TestDay2(t *testing.T) {
	paperNeeded, ribbonNeeded, err := WrapBoxes(day2Input)
	if assert.NoError(t, err) {
		assert.Equal(t, 1598415, paperNeeded, "pt1")
		assert.Equal(t, 3812909, ribbonNeeded, "pt2")
	}
}
