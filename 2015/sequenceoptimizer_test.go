package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOptimize(t *testing.T) {
	testCases := map[string]struct {
		Elements         []string
		Distances        [][]int
		Wrap             bool
		SingleDistMaxSum int
		SingleDistMinSum int
		BothDistMaxSum   int
		BothDistMinSum   int
	}{
		"London, Dublin, Belfast": {
			Elements: []string{"London", "Dublin", "Belfast"},
			Distances: [][]int{
				{0, 464, 518},
				{464, 0, 141},
				{518, 141, 0},
			},
			Wrap:             false,
			SingleDistMaxSum: 982,
			SingleDistMinSum: 605,
			BothDistMaxSum:   982 * 2,
			BothDistMinSum:   605 * 2,
		},
		"Alice, Bob, Carol, David": {
			Elements: []string{"Alice", "Bob", "Carol", "David"},
			Distances: [][]int{
				{0, 54, -79, -2},
				{83, 0, -7, -63},
				{-62, 60, 0, 55},
				{46, -7, 41, 0},
			},
			Wrap:             true,
			SingleDistMaxSum: 86,
			SingleDistMinSum: 86,
			BothDistMaxSum:   330,
			BothDistMinSum:   330,
		},
	}

	for name, c := range testCases {
		seq, err := NewSequence(c.Elements, c.Distances)
		if assert.NoError(t, err, name) {
			_, singleMax := seq.Optimize(MaximizeOptimizationType, false, c.Wrap)
			assert.Equal(t, c.SingleDistMaxSum, singleMax, "single dist, max sum (case=%s)", name)
			_, singleMin := seq.Optimize(MinimizeOptimizationType, false, c.Wrap)
			assert.Equal(t, c.SingleDistMinSum, singleMin, "single dist, min sum (case=%s)", name)
			_, bothMax := seq.Optimize(MaximizeOptimizationType, true, c.Wrap)
			assert.Equal(t, c.BothDistMaxSum, bothMax, "both dist, max sum (case=%s)", name)
			_, bothMin := seq.Optimize(MinimizeOptimizationType, true, c.Wrap)
			assert.Equal(t, c.BothDistMinSum, bothMin, "both dist, min sum (case=%s)", name)
		}
	}
}
