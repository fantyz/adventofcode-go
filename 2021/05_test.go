package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindHydrothermalVentOverlaps(t *testing.T) {
	input := `0,9 -> 5,9
8,0 -> 0,8
9,4 -> 3,4
2,2 -> 2,1
7,0 -> 7,4
6,4 -> 2,0
0,9 -> 2,9
3,4 -> 1,4
0,0 -> 8,8
5,5 -> 8,2`
	vents, err := NewHydrothermalVents(input)
	if assert.NoError(t, err) {
		assert.Equal(t, 5, FindHydrothermalVentOverlaps(vents, true), "without diagonals")
		assert.Equal(t, 12, FindHydrothermalVentOverlaps(vents, false), "with diagonals")
	}
}

func TestDay05Pt1(t *testing.T) {
	vents, err := NewHydrothermalVents(day05Input)
	if assert.NoError(t, err) {
		assert.Equal(t, 5585, FindHydrothermalVentOverlaps(vents, true))
	}
}

func TestDay05Pt2(t *testing.T) {
	vents, err := NewHydrothermalVents(day05Input)
	if assert.NoError(t, err) {
		assert.Equal(t, 17193, FindHydrothermalVentOverlaps(vents, false))
	}
}
