package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSpinlock(t *testing.T) {
	testCases := []struct {
		In     int
		Steps  int
		Result int
	}{
		{9, 3, 5},
		{4, 3, 3},
		{2017, 3, 638},
	}

	for i, testCase := range testCases {
		sl := Spinlock(testCase.Steps, testCase.In)
		t.Log(sl.String())

		assert.Equal(t, testCase.Result, ValueAfter(testCase.In, sl), "(case=%d)", i)
	}
}
