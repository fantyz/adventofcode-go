package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStrongestBridge(t *testing.T) {
	testCases := []struct {
		In     string
		Result int
	}{
		{`Begin in state A.
Perform a diagnostic checksum after 6 steps.

In state A:
  If the current value is 0:
    - Write the value 1.
    - Move one slot to the right.
    - Continue with state B.
  If the current value is 1:
    - Write the value 0.
    - Move one slot to the left.
    - Continue with state B.

In state B:
  If the current value is 0:
    - Write the value 1.
    - Move one slot to the left.
    - Continue with state A.
  If the current value is 1:
    - Write the value 1.
    - Move one slot to the right.
    - Continue with state A.`, 3},
	}

	for i, testCase := range testCases {
		assert.Equal(t, testCase.Result, NewTuringMachine(testCase.In).Execute(), "(case=%d)", i)
	}
}
