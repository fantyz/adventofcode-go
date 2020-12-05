package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSeatID(t *testing.T) {
	testCases := []struct {
		Seat   string
		SeatID int
	}{
		{"FBFBBFFRLR", 357},
		{"BFFFBBFRRR", 567},
		{"FFFBBBFRRR", 119},
		{"BBFFBBFRLL", 820},
	}

	for _, c := range testCases {
		x, y := SeatPos(c.Seat)
		assert.Equal(t, c.SeatID, SeatID(x, y), c.Seat)
	}
}

func TestDay5Pt1(t *testing.T) {
	assert.Equal(t, 989, HighestSeatID(day5Input))
}

func TestDay5Pt2(t *testing.T) {
	assert.Equal(t, 548, FindEmptySeat(day5Input))
}
