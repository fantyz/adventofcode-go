package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDistributePresentsToHouses(t *testing.T) {
	testCases := []struct {
		Moves       string
		Workforce   int
		TotalVisted int
	}{
		{">", 1, 2},
		{"^>v<", 1, 4},
		{"^v^v^v^v^v", 1, 2},
		{"^v", 2, 3},
		{"^>v<", 2, 3},
		{"^v^v^v^v^v", 2, 11},
	}

	for _, c := range testCases {
		moves, err := DistributePresentsToHouses(c.Workforce, c.Moves)
		if assert.NoError(t, err, "%s (workforce=%d)", c.Moves, c.Workforce) {
			assert.Equal(t, c.TotalVisted, moves, "%s (workforce=%d)", c.Moves, c.Workforce)
		}
	}
}

func TestDay3Pt1(t *testing.T) {
	totalVisited, err := DistributePresentsToHouses(1, day3Input)
	if assert.NoError(t, err) {
		assert.Equal(t, 2081, totalVisited)
	}
}

func TestDay3Pt2(t *testing.T) {
	totalVisited, err := DistributePresentsToHouses(2, day3Input)
	if assert.NoError(t, err) {
		assert.Equal(t, 2341, totalVisited)
	}
}
