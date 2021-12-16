package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func LowPointRiskLevels(t *testing.T) {
	in := `2199943210
3987894921
9856789892
8767896789
9899965678`

	cave, err := NewLavaTubesCave(in)
	if assert.NoError(t, err) {
		assert.Equal(t, 15, cave.LowPointRiskLevels())
	}
}

func TestLargestThreeBasinsMultiplied(t *testing.T) {
	in := `2199943210
3987894921
9856789892
8767896789
9899965678`

	cave, err := NewLavaTubesCave(in)
	if assert.NoError(t, err) {
		assert.Equal(t, 1134, cave.LargestThreeBasinsMultiplied())
	}
}

func TestDay09Pt1(t *testing.T) {
	cave, err := NewLavaTubesCave(day09Input)
	if assert.NoError(t, err) {
		assert.Equal(t, 478, cave.LowPointRiskLevels())
	}
}

func TestDay09Pt2(t *testing.T) {
	cave, err := NewLavaTubesCave(day09Input)
	if assert.NoError(t, err) {
		assert.Equal(t, 1327014, cave.LargestThreeBasinsMultiplied())
	}
}
