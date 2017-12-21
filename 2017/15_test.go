package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCompareLow16Bits(t *testing.T) {
	testCases := []struct {
		InA    int
		InB    int
		Result bool
	}{
		{1092455, 430625591, false},
		{1181022009, 1233683848, false},
		{245556042, 1431495498, true},
		{1744312007, 137874439, false},
		{1352636452, 285222916, false},
	}

	for i, testCase := range testCases {
		assert.Equal(t, testCase.Result, CompareLow16Bits(testCase.InA, testCase.InB), "(case=%d)", i)
	}
}

func TestGenerator(t *testing.T) {
	testCases := []struct {
		State  int
		Mult   int
		Picky  int
		Result []int
	}{
		{65, 16807, 1, []int{1092455, 1181022009, 245556042, 1744312007, 1352636452}},
		{8921, 48271, 1, []int{430625591, 1233683848, 1431495498, 137874439, 285222916}},
		{65, 16807, 4, []int{1352636452, 1992081072, 530830436, 1980017072, 740335192}},
	}

	for i, testCase := range testCases {
		g := NewGenerator(testCase.Mult, testCase.State, testCase.Picky)
		var res []int
		for j := 0; j < len(testCase.Result); j++ {
			res = append(res, g.Next())
		}
		assert.Equal(t, testCase.Result, res, "(case=%d)", i)
	}
}

func TestMatchPair(t *testing.T) {
	testCases := []struct {
		Pairs  int
		StateA int
		PickyA int
		StateB int
		PickyB int
		Result int
	}{
		{40000000, 65, 1, 8921, 1, 588},
		{5000000, 65, 4, 8921, 8, 309},
	}

	for i, testCase := range testCases {
		assert.Equal(t, testCase.Result, MatchPairs(testCase.Pairs, testCase.StateA, testCase.PickyA, testCase.StateB, testCase.PickyB), "(case=%d)", i)
	}
}
