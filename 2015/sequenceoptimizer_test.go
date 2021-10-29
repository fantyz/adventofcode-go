package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOptimize(t *testing.T) {
	testCases := map[string]struct {
		Elements  []string
		Distances [][]int
		Wrap      bool
		Min       int
		Max       int
	}{
		"London, Dublin, Belfast": {
			Elements: []string{"London", "Dublin", "Belfast"},
			Distances: [][]int{
				{0, 464, 518},
				{464, 0, 141},
				{518, 141, 0},
			},
			Wrap: false,
			Min:  605,
			Max:  982,
		},
		"Alice, Bob, Carol, David": {
			Elements: []string{"Alice", "Bob", "Carol", "David"},
			Distances: [][]int{
				{0, 137, -141, 44},
				{137, 0, 53, -70},
				{-141, 53, 0, 96},
				{44, -70, 96, 0},
			},
			Wrap: true,
			Min:  330,
			Max:  330,
		},
	}

	for name, c := range testCases {
		seq, err := NewSequence(c.Elements, c.Distances)
		if assert.NoError(t, err, name) {
			_, min := seq.Optimize(MinimizeOptimizationType, c.Wrap)
			assert.Equal(t, c.Min, min, "min (case=%s)", name)
			_, max := seq.Optimize(MaximizeOptimizationType, c.Wrap)
			assert.Equal(t, c.Max, max, "max (case=%s)", name)
		}
	}
}
