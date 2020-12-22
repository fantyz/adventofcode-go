package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolveTilePuzzle(t *testing.T) {
	testTiles := `Tile 2311:
..##.#..#.
##..#.....
#...##..#.
####.#...#
##.##.###.
##...#.###
.#.#.#..##
..#....#..
###...#.#.
..###..###

Tile 1951:
#.##...##.
#.####...#
.....#..##
#...######
.##.#....#
.###.#####
###.##.##.
.###....#.
..#.#..#.#
#...##.#..

Tile 1171:
####...##.
#..##.#..#
##.#..#.#.
.###.####.
..###.####
.##....##.
.#...####.
#.##.####.
####..#...
.....##...

Tile 1427:
###.##.#..
.#..#.##..
.#.##.#..#
#.#.#.##.#
....#...##
...##..##.
...#.#####
.#.####.#.
..#..###.#
..##.#..#.

Tile 1489:
##.#.#....
..##...#..
.##..##...
..#...#...
#####...#.
#..#.#.#.#
...#.#.#..
##.#...##.
..##.##.##
###.##.#..

Tile 2473:
#....####.
#..#.##...
#.##..#...
######.#.#
.#...#.#.#
.#########
.###.#..#.
########.#
##...##.#.
..###.#.#.

Tile 2971:
..#.#....#
#...###...
#.#.###...
##.##..#..
.#####..##
.#..####.#
#..#.#..#.
..####.###
..#.#.###.
...#.#.#.#

Tile 2729:
...#.#.#.#
####.#....
..#.#.....
....#..#.#
.##..##.#.
.#.####...
####.#.#..
##.####...
##..#.##..
#.##...##.

Tile 3079:
#.#.#####.
.#..######
..#.......
######....
####.#..#.
.#...#.##.
#.#####.##
..#.###...
..#.......
..#.###...`

	tiles, err := LoadTiles(testTiles)
	if assert.NoError(t, err) {
		if assert.True(t, SolveTilePuzzle(tiles)) {
			assert.Equal(t, 20899048083289, MultiplyCornerIDs(tiles))
		}
	}
}

func TestNewMonsterPic(t *testing.T) {
	testTiles := `Tile 1:
###
###
###

Tile 2:
...
...
...

Tile 3:
###
###
###

Tile 4:
...
...
...

Tile 5:
###
###
###

Tile 6:
...
...
...

Tile 7:
###
###
###

Tile 8:
...
...
...

Tile 9:
###
###
###`
	expPic := [][]bool{
		{true, false, true},
		{false, true, false},
		{true, false, true},
	}

	tiles, err := LoadTiles(testTiles)
	if assert.NoError(t, err) {
		assert.Equal(t, expPic, NewMonsterPicture(tiles))
	}
}

func TestWaterRoughnessAfterFindingSeaMonsters(t *testing.T) {
	testPicInput := `.#.#..#.##...#.##..#####
###....#.#....#..#......
##.##.###.#.#..######...
###.#####...#.#####.#..#
##.#....#.##.####...#.##
...########.#....#####.#
....#..#...##..#.#.###..
.####...#..#.....#......
#..#.##..#..###.#.##....
#.####..#.####.#.#.###..
###.#.#...#.######.#..##
#.####....##..########.#
##..##.#...#...#.#.#.#..
...#..#..#.#.##..###.###
.#.#....#.##.#...###.##.
###.#...#..#.##.######..
.#.#.###.##.##.#..#.##..
.####.###.#...###.#..#.#
..#.#..#..#.#.#.####.###
#..####...#.#.#.###.###.
#####..#####...###....##
#.##..#..#...#..####...#
.#.###..##..##..####.##.
...###...##...#...#..###`

	var testPic [][]bool
	for _, l := range strings.Split(testPicInput, "\n") {
		row := make([]bool, len(l))
		for i := 0; i < len(l); i++ {
			if l[i] == '#' {
				row[i] = true
			}
		}
		testPic = append(testPic, row)
	}

	assert.Equal(t, 273, WaterRoughnessAfterFindingSeaMonsters(testPic))
}

func TestDay20(t *testing.T) {
	tiles, err := LoadTiles(day20Input)
	if assert.NoError(t, err) && assert.True(t, SolveTilePuzzle(tiles)) {
		assert.Equal(t, 30425930368573, MultiplyCornerIDs(tiles))
		pic := NewMonsterPicture(tiles)
		assert.Equal(t, 2453, WaterRoughnessAfterFindingSeaMonsters(pic))
	}
}

func BenchmarkDay20Pt1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		tiles, _ := LoadTiles(day20Input)
		b.StartTimer()
		SolveTilePuzzle(tiles)
	}
}
