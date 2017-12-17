package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShortestPath(t *testing.T) {
	testCases := []struct {
		In     string
		X      int
		Y      int
		Z      int
		Result int
	}{
		{`ne,ne,ne`, 3, 0, -3, 3},
		{`ne,ne,sw,sw`, 0, 0, 0, 0},
		{`ne,ne,s,s`, 2, -2, 0, 2},
		{`se,sw,se,sw,sw`, -1, -2, 3, 3},
	}

	for i, testCase := range testCases {
		x, y, z, _ := TracePath(testCase.In)
		assert.Equal(t, testCase.X, x, "x=%d (case=%d)", x, i)
		assert.Equal(t, testCase.Y, y, "y=%d (case=%d)", y, i)
		assert.Equal(t, testCase.Z, z, "z=%d (case=%d)", z, i)
		assert.Equal(t, testCase.Result, ShortestPath(x, y, z), "%v (case=%d)", testCase.In, i)
	}
}
