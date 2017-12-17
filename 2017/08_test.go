package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLargest(t *testing.T) {
	testCases := []struct {
		In     string
		Result int
		Max    int
	}{
		{`b inc 5 if a > 1
a inc 1 if b < 5
c dec -10 if a >= 1
c inc -20 if c == 10`, 1, 10},
	}

	for i, testCase := range testCases {
		regs, max := NewProgram(testCase.In).Exec()
		assert.Equal(t, testCase.Result, regs.Largest(), "(case=%d)", i)
		assert.Equal(t, testCase.Max, max, "(case=%d)", i)
	}
}
