package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCountNodesThatCanReach(t *testing.T) {
	testCases := []struct {
		In     string
		Size   int
		Result int
		Groups int
	}{
		{`0 <-> 2
1 <-> 1
2 <-> 0, 3, 4
3 <-> 2, 4
4 <-> 2, 3, 6
5 <-> 6
6 <-> 4, 5`, 7, 6, 2},
	}

	for i, testCase := range testCases {
		g := NewGraph(testCase.Size, testCase.In)
		assert.Equal(t, testCase.Result, g.CountNodesThatCanReach(0), "Nodes can reach 0 (case=%d)", i)
		assert.Equal(t, testCase.Groups, g.CountGroups(), "Groups (case=%d)", i)
	}
}
