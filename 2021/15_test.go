package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLowestRiskRoute(t *testing.T) {
	in := `1163751742
1381373672
2136511328
3694931569
7463417111
1319128137
1359912421
3125421639
1293138521
2311944581`

	cave, err := NewCave(in)
	if assert.NoError(t, err) {
		assert.Equal(t, 40, LowestRiskRoute(cave, CaveLoc{0, 0}, CaveLoc{cave.Width() - 1, cave.Height() - 1}))
	}
}

func TestLowestRiskRouteInEnlargedCave(t *testing.T) {
	in := `1163751742
1381373672
2136511328
3694931569
7463417111
1319128137
1359912421
3125421639
1293138521
2311944581`

	cave, err := NewCave(in)
	if assert.NoError(t, err) {
		cave.Enlarge(5)
		assert.Equal(t, 315, LowestRiskRoute(cave, CaveLoc{0, 0}, CaveLoc{cave.Width() - 1, cave.Height() - 1}))
	}
}

func TestDay15Pt1(t *testing.T) {
	cave, err := NewCave(day15Input)
	if assert.NoError(t, err) {
		assert.Equal(t, 685, LowestRiskRoute(cave, cave.TopLeft(), cave.BottomRight()))
	}
}

func TestDay15Pt2(t *testing.T) {
	cave, err := NewCave(day15Input)
	if assert.NoError(t, err) {
		cave.Enlarge(5)
		assert.Equal(t, 2995, LowestRiskRoute(cave, cave.TopLeft(), cave.BottomRight()))
	}
}
