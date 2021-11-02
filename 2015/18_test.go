package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAnimateLight(t *testing.T) {
	input := `.#.#.#
...##.
#....#
..#...
#.#..#
####..`
	lights := NewAnimatedLights(input)
	lights.Animate(4, false)
	assert.Equal(t, 4, lights.LightsOn())
}

func TestAnimateLightWithStuckCorners(t *testing.T) {
	input := `.#.#.#
...##.
#....#
..#...
#.#..#
####..`
	lights := NewAnimatedLights(input)
	lights.Animate(5, true)
	assert.Equal(t, 17, lights.LightsOn())
}

func TestDay18Pt1(t *testing.T) {
	lights := NewAnimatedLights(day18Input)
	lights.Animate(100, false)
	assert.Equal(t, 821, lights.LightsOn())
}

func TestDay18Pt2(t *testing.T) {
	lights := NewAnimatedLights(day18Input)
	lights.Animate(100, true)
	assert.Equal(t, 886, lights.LightsOn())
}
