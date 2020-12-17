package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMemoryGame(t *testing.T) {
	testCases := map[string]struct {
		Init   []int
		Turn   int
		Number int
	}{
		"0,3,6": {[]int{0, 3, 6}, 10, 0},
		"1,3,2": {[]int{1, 3, 2}, 2020, 1},
		"2,1,3": {[]int{2, 1, 3}, 2020, 10},
		"1,2,3": {[]int{1, 2, 3}, 2020, 27},
		"2,3,1": {[]int{2, 3, 1}, 2020, 78},
		"3,2,1": {[]int{3, 2, 1}, 2020, 438},
		"3,1,2": {[]int{3, 1, 2}, 2020, 1836},
	}

	for n, c := range testCases {
		assert.Equal(t, c.Number, MemoryGame(c.Init, c.Turn), n)
	}
}

func TestLongMemoryGame(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping long running memory games")
	}

	testCases := map[string]struct {
		Init   []int
		Turn   int
		Number int
	}{
		"0,3,6": {[]int{0, 3, 6}, 30000000, 175594},
		"1,3,2": {[]int{1, 3, 2}, 30000000, 2578},
		"2,1,3": {[]int{2, 1, 3}, 30000000, 3544142},
		"1,2,3": {[]int{1, 2, 3}, 30000000, 261214},
		"2,3,1": {[]int{2, 3, 1}, 30000000, 6895259},
		"3,2,1": {[]int{3, 2, 1}, 30000000, 18},
		"3,1,2": {[]int{3, 1, 2}, 30000000, 362},
	}

	for n, c := range testCases {
		assert.Equal(t, c.Number, MemoryGame(c.Init, c.Turn), n)
	}
}

func TestDay15Pt1(t *testing.T) {
	init, err := LoadInts(day15Input, ",")
	if assert.NoError(t, err) {
		assert.Equal(t, 706, MemoryGame(init, 2020))
	}
}

func TestDay15Pt2(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping slow running memory game")
	}
	init, err := LoadInts(day15Input, ",")
	if assert.NoError(t, err) {
		assert.Equal(t, 19331, MemoryGame(init, 30000000))
	}
}
