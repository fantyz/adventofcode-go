package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindStartOfPacketMarker(t *testing.T) {
	testCases := []struct {
		In     string
		Marker int
	}{
		{"mjqjpqmgbljsphdztnvjfqwrcgsmlb", 7},
		{"bvwbjplbgvbhsrlpgdmjqwftvncz", 5},
		{"nppdvjthqldpwncqszvftbrmjlhg", 6},
		{"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 10},
		{"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 11},
	}

	for _, c := range testCases {
		assert.Equal(t, c.Marker, FindStartOfPacketMarker(c.In), "in=%s", c.In)
	}
}

func TestDay06Pt1(t *testing.T) {
}

func TestDay06Pt2(t *testing.T) {
}
