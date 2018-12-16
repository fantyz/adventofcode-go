package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerate(t *testing.T) {
	testCases := []struct {
		In  int
		Out string
	}{
		{9, "5158916779"},
		{5, "0124515891"},
		{18, "9251071085"},
		{2018, "5941429882"},
	}

	for i, c := range testCases {
		sb := NewScoreboard()
		assert.Equal(t, c.Out, sb.Generate(c.In), "(case=%d)", i)
	}
}

func TestPattern(t *testing.T) {
	testCases := []struct {
		In  []int
		Out int
	}{
		{[]int{5, 1, 5, 8, 9}, 9},
		{[]int{0, 1, 2, 4, 5}, 5},
		{[]int{9, 2, 5, 1, 0}, 18},
		{[]int{5, 9, 4, 1, 4}, 2018},
	}

	for i, c := range testCases {
		sb := NewScoreboard()
		assert.Equal(t, c.Out, sb.Pattern(c.In), "(case=%d)", i)
	}
}
