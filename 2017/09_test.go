package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestScoreStream(t *testing.T) {
	testCases := []struct {
		In     string
		Result int
		Garb   int
	}{
		{`{}`, 1, 0},
		{`{{{}}}`, 6, 0},
		{`{{},{}}`, 5, 0},
		{`{{{},{},{{}}}}`, 16, 0},
		{`{<a>,<a>,<a>,<a>}`, 1, 4},
		{`{{<ab>},{<ab>},{<ab>},{<ab>}}`, 9, 8},
		{`{{<!!>},{<!!>},{<!!>},{<!!>}}`, 9, 0},
		{`{{<a!>},{<a!>},{<a!>},{<ab>}}`, 3, 17},
		{`<>`, 0, 0},
		{`<random characters>`, 0, 17},
		{`<<<<>`, 0, 3},
		{`<{!>}>`, 0, 2},
		{`<!!>`, 0, 0},
		{`<!!!>>`, 0, 0},
		{`<{o"i!a,<{i<a>`, 0, 10},
	}

	for i, testCase := range testCases {
		score, garb := ScoreStream(testCase.In)
		assert.Equal(t, testCase.Result, score, "%v (case=%d)", testCase.In, i)
		assert.Equal(t, testCase.Garb, garb, "%v (case=%d)", testCase.In, i)
	}
}
