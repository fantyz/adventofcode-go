package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInputToCoords(t *testing.T) {
	testCases := []struct {
		In  string
		Out []Coord
	}{
		{
			In: `1, 1
1, 6
8, 3`,
			Out: []Coord{{1, 1}, {1, 6}, {8, 3}},
		},
	}

	for i, c := range testCases {
		assert.Equal(t, c.Out, InputToCoordinates(c.In), "(case=%d)", i)
	}
}

func TestLargestArea(t *testing.T) {
	testCases := []struct {
		In  string
		Out int
	}{
		{
			In: `1, 1
1, 6
8, 3
3, 4
5, 5
8, 9`,
			Out: 17,
		},
	}

	for i, c := range testCases {
		assert.Equal(t, c.Out, LargestArea(InputToCoordinates(c.In)), "(case=%d)", i)
	}
}

func TestRegionWithinMaxDist(t *testing.T) {
	testCases := []struct {
		In      string
		MaxDist int
		Out     int
	}{
		{
			In: `1, 1
1, 6
8, 3
3, 4
5, 5
8, 9`,
			MaxDist: 32,
			Out:     16,
		},
	}

	for i, c := range testCases {
		assert.Equal(t, c.Out, RegionWithinMaxDist(c.MaxDist, InputToCoordinates(c.In)), "(case=%d)", i)
	}
}
