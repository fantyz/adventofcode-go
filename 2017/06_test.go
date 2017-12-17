package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRedistribute(t *testing.T) {
	testCases := []struct {
		In     Blocks
		Result Blocks
	}{
		{Blocks{0, 2, 7, 0}, Blocks{2, 4, 1, 2}},
		{Blocks{2, 4, 1, 2}, Blocks{3, 1, 2, 3}},
		{Blocks{3, 1, 2, 3}, Blocks{0, 2, 3, 4}},
		{Blocks{0, 2, 3, 4}, Blocks{1, 3, 4, 1}},
		{Blocks{1, 3, 4, 1}, Blocks{2, 4, 1, 2}},
	}

	for i, testCase := range testCases {
		testCase.In.Redistribute()
		assert.Equal(t, testCase.Result, testCase.In, "(case=%d)", i)
	}
}

func TestRedistributeUntilSeenStateSteps(t *testing.T) {
	testCases := []struct {
		In     Blocks
		Result int
	}{
		{Blocks{0, 2, 7, 0}, 5},
	}

	for i, testCase := range testCases {
		assert.Equal(t, testCase.Result, testCase.In.RedistributeUntilSeenStateSteps(), "(case=%d)", i)
	}
}
