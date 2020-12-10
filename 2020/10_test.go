package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJoltDifference(t *testing.T) {
	testCases := map[string]struct {
		Adapters []int
		Diff     int
	}{
		"short": {[]int{16, 10, 15, 5, 1, 11, 7, 19, 6, 12, 4}, 7 * 5},
		"long":  {[]int{28, 33, 18, 42, 31, 14, 46, 20, 48, 47, 24, 23, 49, 45, 19, 38, 39, 11, 1, 32, 25, 35, 8, 17, 7, 9, 4, 2, 34, 10, 3}, 22 * 10},
	}

	for n, c := range testCases {
		diff, err := JoltDifference(c.Adapters)
		if assert.NoError(t, err, n) {
			assert.Equal(t, c.Diff, diff, n)
		}
	}
}

func TestAdapterCombinations(t *testing.T) {
	testCases := map[string]struct {
		Adapters     []int
		Combinations int
	}{
		"short": {[]int{16, 10, 15, 5, 1, 11, 7, 19, 6, 12, 4}, 8},
		"long":  {[]int{28, 33, 18, 42, 31, 14, 46, 20, 48, 47, 24, 23, 49, 45, 19, 38, 39, 11, 1, 32, 25, 35, 8, 17, 7, 9, 4, 2, 34, 10, 3}, 19208},
	}

	for n, c := range testCases {
		assert.Equal(t, c.Combinations, AdapterCombinations(c.Adapters), n)
	}
}

func TestDay10Pt1(t *testing.T) {
	in, err := LoadInts(day10Input)
	if assert.NoError(t, err) {
		diff, err := JoltDifference(in)
		if assert.NoError(t, err) {
			assert.Equal(t, 1885, diff)
		}
	}
}

func TestDay10Pt2(t *testing.T) {
	in, err := LoadInts(day10Input)
	if assert.NoError(t, err) {
		assert.Equal(t, 2024782584832, AdapterCombinations(in))
	}
}
