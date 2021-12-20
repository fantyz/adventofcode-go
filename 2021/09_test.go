package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLowPointRiskLevels(t *testing.T) {
	in := `2199943210
3987894921
9856789892
8767896789
9899965678`

	cave, err := NewCave(in)
	if assert.NoError(t, err) {
		assert.Equal(t, 15, LowPointRiskLevels(cave))
	}
}

func TestLargestThreeBasinsMultiplied(t *testing.T) {
	in := `2199943210
3987894921
9856789892
8767896789
9899965678`

	cave, err := NewCave(in)
	if assert.NoError(t, err) {
		assert.Equal(t, 1134, LargestThreeBasinsMultiplied(cave))
	}
}

func TestDay09Pt1(t *testing.T) {
	cave, err := NewCave(day09Input)
	if assert.NoError(t, err) {
		assert.Equal(t, 478, LowPointRiskLevels(cave))
	}
}

func TestDay09Pt2(t *testing.T) {
	cave, err := NewCave(day09Input)
	if assert.NoError(t, err) {
		assert.Equal(t, 1327014, LargestThreeBasinsMultiplied(cave))
	}
}
