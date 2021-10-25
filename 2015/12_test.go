package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumJSONDocumentNumbers(t *testing.T) {
	testCases := []struct {
		Doc       string
		IgnoreRed bool
		Sum       int
	}{
		{`[1,2,3]`, true, 6},
		{`{"a":2,"b":4}`, false, 6},
		{`[[[3]]]`, false, 3},
		{`{"a":{"b":4},"c":-1}`, false, 3},
		{`{"a":[-1,1]}`, false, 0},
		{`[-1,{"a":1}]`, false, 0},
		{`[]`, false, 0},
		{`{}`, false, 0},
		{`[1,{"c":"red","b":2},3]`, true, 4},
		{`{"d":"red","e":[1,2,3,4],"f":5}`, true, 0},
		{`[1,"red",5]`, true, 6},
	}

	for _, c := range testCases {
		sum, err := SumJSONDocumentNumbers(c.Doc, c.IgnoreRed)
		if assert.NoError(t, err, c.Doc) {
			assert.Equal(t, c.Sum, sum, c.Doc)
		}
	}
}

func TestDay12Pt1(t *testing.T) {
	sum, err := SumJSONDocumentNumbers(day12Input, false)
	if assert.NoError(t, err) {
		assert.Equal(t, 111754, sum)
	}
}

func TestDay12Pt2(t *testing.T) {
	sum, err := SumJSONDocumentNumbers(day12Input, true)
	if assert.NoError(t, err) {
		assert.Equal(t, 65402, sum)
	}
}
