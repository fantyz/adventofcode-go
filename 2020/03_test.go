package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTraverseForestWithRoutes(t *testing.T) {
	testForest := `..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#`

	testCases := map[string]struct {
		Routes []Route
		Trees  int
	}{
		"(3,1)":                         {[]Route{{3, 1}}, 7},
		"(1,1),(3,1),(5,1),(7,1),(1,2)": {[]Route{{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2}}, 336},
	}

	for n, c := range testCases {
		assert.Equal(t, c.Trees, TraverseForestWithRoutes(testForest, c.Routes), n)
	}
}
