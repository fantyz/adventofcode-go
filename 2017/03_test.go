package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSpiralMemoryDistance(t *testing.T) {
	testCases := []struct {
		In     int
		Result int
	}{
		{1, 0},
		{12, 3},
		{23, 2},
		{1024, 31},
	}

	for i, testCase := range testCases {
		assert.Equal(t, testCase.Result, SpiralMemoryDistance(testCase.In), "in=%d (case=%d)", testCase.In, i)
	}
}

func TestSpiralAdjacentSum(t *testing.T) {
	testCases := []struct {
		In     int
		Result int
	}{
		{0, 1},
		{1, 2},
		{3, 4},
		{700, 747},
	}

	for i, testCase := range testCases {
		assert.Equal(t, testCase.Result, SpiralAdjacentSum(testCase.In), "in=%d (case=%d)", testCase.In, i)
	}
}
