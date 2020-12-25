package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlayCrabCups(t *testing.T) {
	testCups := []int{3, 8, 9, 1, 2, 5, 4, 6, 7}
	testCases := []struct {
		Rounds  int
		LastCup int
		Result  int
	}{
		{0, len(testCups), 25467389},
		{10, len(testCups), 92658374},
		{100, len(testCups), 67384529},
		{10000000, 1000000, 149245887792},
	}
	for _, c := range testCases {
		assert.Equal(t, c.Result, PlayCrabCups(c.Rounds, testCups, c.LastCup), c.Rounds)
	}
}

func TestDay23Pt1(t *testing.T) {
	cups, err := LoadInts(day23Input, "")
	if assert.NoError(t, err) {
		assert.Equal(t, 45983627, PlayCrabCups(100, cups, 0))
	}
}

func TestDay23Pt2(t *testing.T) {
	cups, err := LoadInts(day23Input, "")
	if assert.NoError(t, err) {
		assert.Equal(t, 111080192688, PlayCrabCups(10000000, cups, 1000000))
	}
}
