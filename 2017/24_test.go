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
		{`0/2
2/2
2/3
3/4
3/5
0/1
10/1
9/10`, 31},
	}

	for i, testCase := range testCases {
		assert.Equal(t, testCase.Result, StrongestBridge(0, nil, NewComponents(testCase.In)), "(case=%d)", i)
	}
}
