package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMineAdentCoin(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping slow MineAdventCoint tests")
	}

	testCases := []struct {
		Key string
		Out int
	}{
		{"abcdef", 609043},
		{"pqrstuv", 1048970},
	}

	for _, c := range testCases {
		assert.Equal(t, c.Out, MineAdventCoin(c.Key, 5), c.Key)
	}
}

func TestDay4Pt1(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping slow MineAdventCoint tests")
	}
	assert.Equal(t, 346386, MineAdventCoin(day4input, 5))
}

func TestDay4Pt2(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping slow MineAdventCoint tests")
	}
	assert.Equal(t, 9958218, MineAdventCoin(day4input, 6))
}
