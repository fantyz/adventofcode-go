package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDepthIncreases(t *testing.T) {
	testCases := map[string]struct {
		Depths    []int
		Window    int
		Increases int
	}{
		"window=1": {
			Depths:    []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263},
			Window:    1,
			Increases: 7,
		},
		"window=3": {
			Depths:    []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263},
			Window:    3,
			Increases: 5,
		},
	}

	for n, c := range testCases {
		assert.Equal(t, c.Increases, DepthIncreases(c.Window, c.Depths), n)
	}
}

func TestDay01Pt1(t *testing.T) {
	depths, err := LoadInts(day01Input, "\n")
	if assert.NoError(t, err) {
		assert.Equal(t, 1759, DepthIncreases(1, depths))
	}
}

func TestDay01Pt2(t *testing.T) {
	depths, err := LoadInts(day01Input, "\n")
	if assert.NoError(t, err) {
		assert.Equal(t, 1805, DepthIncreases(3, depths))
	}
}
