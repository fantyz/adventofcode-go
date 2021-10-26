package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRace(t *testing.T) {
	input := `Comet can fly 14 km/s for 10 seconds, but then must rest for 127 seconds.
Dancer can fly 16 km/s for 11 seconds, but then must rest for 162 seconds.`

	testCases := []struct {
		Duration int
		Dist     int
	}{
		{1, 16},
		{10, 160},
		{11, 176},
		{12, 176},
		{1000, 1120},
	}

	raindeers, err := NewRaindeers(input)
	if assert.NoError(t, err) {
		for _, c := range testCases {
			_, dist := raindeers.Race(c.Duration)
			assert.Equal(t, c.Dist, dist, c.Duration)
		}
	}
}

func TestPointsRace(t *testing.T) {
	input := `Comet can fly 14 km/s for 10 seconds, but then must rest for 127 seconds.
Dancer can fly 16 km/s for 11 seconds, but then must rest for 162 seconds.`

	raindeers, err := NewRaindeers(input)
	if assert.NoError(t, err) {
		name, points := raindeers.PointsRace(1000)
		assert.Equal(t, "Dancer", name)
		assert.Equal(t, 689, points)
	}
}

func TestDay14Pt1(t *testing.T) {
	raindeers, err := NewRaindeers(day14Input)
	if assert.NoError(t, err) {
		_, dist := raindeers.Race(2503)
		assert.Equal(t, 2655, dist)
	}
}

func TestDay14Pt2(t *testing.T) {
	raindeers, err := NewRaindeers(day14Input)
	if assert.NoError(t, err) {
		_, dist := raindeers.PointsRace(2503)
		assert.Equal(t, 1059, dist)
	}
}
