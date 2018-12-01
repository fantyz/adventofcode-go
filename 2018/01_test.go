package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestChronalCalibration(t *testing.T) {
	testCases := []struct {
		In   string
		Sum  int
		Freq int
	}{
		{"+1\n-1", 0, 0},
		{"+3\n+3\n+4\n-2\n-4", 4, 10},
		{"-6\n+3\n+8\n+5\n-6", 4, 5},
		{"+7\n+7\n-2\n-7\n-4", 1, 14},
		{"+1\n-2\n+3\n+1", 3, 2},
	}

	for i, c := range testCases {
		sum, freq := ChronalCalibration(c.In)
		assert.Equal(t, c.Sum, sum, "(case=%d)", i)
		assert.Equal(t, c.Freq, freq, "(case=%d)", i)
	}
}
