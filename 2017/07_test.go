package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewTower(t *testing.T) {
	testCases := []struct {
		In     string
		Result string
	}{
		{`pbga (66)
xhth (57)
ebii (61)
havc (66)
ktlj (57)
fwft (72) -> ktlj, cntj, xhth
qoyq (66)
padx (45) -> pbga, havc, qoyq
tknk (41) -> ugml, padx, fwft
jptl (61)
ugml (68) -> gyxo, ebii, jptl
gyxo (61)
cntj (57)`, "tknk"},
	}

	for i, testCase := range testCases {
		assert.Equal(t, testCase.Result, NewTower(testCase.In).Name, "(case=%d)", i)
	}
}
