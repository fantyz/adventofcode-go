package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPlayGame(t *testing.T) {
	testCases := []struct {
		Players      int
		LastMarble   int
		ExpHighscore int
	}{
		{9, 25, 32},
		{10, 1618, 8317},
		{13, 7999, 146373},
		{17, 1104, 2764},
		{21, 6111, 54718},
		{30, 5807, 37305},
	}

	for i, c := range testCases {
		assert.Equal(t, c.ExpHighscore, PlayGame(c.Players, c.LastMarble), "(case=%d)", i)
	}
}
