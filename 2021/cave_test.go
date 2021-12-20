package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCaveNeighbors(t *testing.T) {
	in := `123
456
789`

	testCases := map[string]struct {
		Pos       CaveLoc
		InclDiag  bool
		Neighbors []CaveLoc
	}{
		"0,0 - diag":   {CaveLoc{0, 0}, false, []CaveLoc{{1, 0}, {0, 1}}},
		"0,0 + diag":   {CaveLoc{0, 0}, true, []CaveLoc{{1, 0}, {0, 1}, {1, 1}}},
		"1,1 - diag":   {CaveLoc{1, 1}, false, []CaveLoc{{1, 0}, {0, 1}, {2, 1}, {1, 2}}},
		"1,1 + diag":   {CaveLoc{1, 1}, true, []CaveLoc{{0, 0}, {1, 0}, {2, 0}, {0, 1}, {2, 1}, {0, 2}, {1, 2}, {2, 2}}},
		"-1,0 + daig":  {CaveLoc{-1, 0}, true, []CaveLoc{{0, 0}, {0, 1}}},
		"-2,-2 + daig": {CaveLoc{-2, 2}, true, []CaveLoc{}},
	}

	for name, c := range testCases {
		cave, err := NewCave(in)
		if assert.NoError(t, err, name) {
			neighbors := cave.Neighbors(c.Pos, c.InclDiag)
			assert.Equal(t, len(c.Neighbors), len(neighbors), name)
			for _, n := range neighbors {
				assert.Contains(t, c.Neighbors, n, name)
			}
		}
	}
}

func TestCaveEnlarge(t *testing.T) {
	cave := `1`
	expEnlargedCave := `12345
23456
34567
45678
56789`

	c, err := NewCave(cave)
	if assert.NoError(t, err) {
		c.Enlarge(5)
		assert.Equal(t, expEnlargedCave, c.String())
	}
}
