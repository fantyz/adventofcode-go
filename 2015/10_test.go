package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLookAndSayRound(t *testing.T) {
	testCases := []struct {
		In  string
		Out string
	}{
		{"1", "11"},
		{"11", "21"},
		{"1211", "111221"},
		{"111221", "312211"},
	}

	for _, c := range testCases {
		assert.Equal(t, c.Out, LookAndSayRound(c.In), c.In)
	}
}

func TestLookAndSay(t *testing.T) {
	assert.Equal(t, "312211", LookAndSay("1", 5))
}

func TestDay10Pt1(t *testing.T) {
	assert.Equal(t, 329356, len(LookAndSay(day10Input, 40)))
}

func TestDay10Pt2(t *testing.T) {
	assert.Equal(t, 4666278, len(LookAndSay(day10Input, 50)))
}
