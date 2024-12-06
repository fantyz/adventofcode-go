package main

import (
	"testing"


        "github.com/stretchr/testify/assert"
)

func TestGuardPatrolPositions(t *testing.T) {
	lab, err := NewLabMap(`....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`)
	if assert.NoError(t, err) {
		assert.Equal(t, 41, lab.GuardPatrolPositions())
	}
}

func TestNumberOfWaysToTrapGuard(t *testing.T) {
	lab, err := NewLabMap(`....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`)
	if assert.NoError(t, err) {
		assert.Equal(t, 6, lab.NumberOfWaysToTrapGuard())
	}
}

func TestDay06Pt1(t *testing.T) {
	lab, err := NewLabMap(day06Input)
	if assert.NoError(t, err) {
		assert.Equal(t, 5208, lab.GuardPatrolPositions())
	}
}

func TestDay06Pt2(t *testing.T) {
	lab, err := NewLabMap(day06Input)
	if assert.NoError(t, err) {
		assert.Equal(t, -1, lab.NumberOfWaysToTrapGuard())
	}
}

