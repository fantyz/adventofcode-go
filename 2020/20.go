package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

func init() { days["20"] = Day20 }

/*
--- Day 20: Jurassic Jigsaw ---
The high-speed train leaves the forest and quickly carries you south. You can even see a desert in the distance! Since you have some spare time, you might as well see if there was anything interesting in the image the Mythical Information Bureau satellite captured.

After decoding the satellite messages, you discover that the data actually contains many small images created by the satellite's camera array. The camera array consists of many cameras; rather than produce a single square image, they produce many smaller square image tiles that need to be reassembled back into a single image.

Each camera in the camera array returns a single monochrome image tile with a random unique ID number. The tiles (your puzzle input) arrived in a random order.

Worse yet, the camera array appears to be malfunctioning: each image tile has been rotated and flipped to a random orientation. Your first task is to reassemble the original image by orienting the tiles so they fit together.

To show how the tiles should be reassembled, each tile's image data includes a border that should line up exactly with its adjacent tiles. All tiles have this border, and the border lines up exactly when the tiles are both oriented correctly. Tiles at the edge of the image also have this border, but the outermost edges won't line up with any other tiles.

For example, suppose you have the following nine tiles:

Tile 2311:
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
..#.###...
By rotating, flipping, and rearranging them, you can find a square arrangement that causes all adjacent borders to line up:

#...##.#.. ..###..### #.#.#####.
..#.#..#.# ###...#.#. .#..######
.###....#. ..#....#.. ..#.......
###.##.##. .#.#.#..## ######....
.###.##### ##...#.### ####.#..#.
.##.#....# ##.##.###. .#...#.##.
#...###### ####.#...# #.#####.##
.....#..## #...##..#. ..#.###...
#.####...# ##..#..... ..#.......
#.##...##. ..##.#..#. ..#.###...

#.##...##. ..##.#..#. ..#.###...
##..#.##.. ..#..###.# ##.##....#
##.####... .#.####.#. ..#.###..#
####.#.#.. ...#.##### ###.#..###
.#.####... ...##..##. .######.##
.##..##.#. ....#...## #.#.#.#...
....#..#.# #.#.#.##.# #.###.###.
..#.#..... .#.##.#..# #.###.##..
####.#.... .#..#.##.. .######...
...#.#.#.# ###.##.#.. .##...####

...#.#.#.# ###.##.#.. .##...####
..#.#.###. ..##.##.## #..#.##..#
..####.### ##.#...##. .#.#..#.##
#..#.#..#. ...#.#.#.. .####.###.
.#..####.# #..#.#.#.# ####.###..
.#####..## #####...#. .##....##.
##.##..#.. ..#...#... .####...#.
#.#.###... .##..##... .####.##.#
#...###... ..##...#.. ...#..####
..#.#....# ##.#.#.... ...##.....
For reference, the IDs of the above tiles are:

1951    2311    3079
2729    1427    2473
2971    1489    1171
To check that you've assembled the image correctly, multiply the IDs of the four corner tiles together. If you do this with the assembled tiles from the example above, you get 1951 * 3079 * 2971 * 1171 = 20899048083289.

Assemble the tiles into an image. What do you get if you multiply together the IDs of the four corner tiles?

Your puzzle answer was 30425930368573.

--- Part Two ---
Now, you're ready to check the image for sea monsters.

The borders of each tile are not part of the actual image; start by removing them.

In the example above, the tiles become:

.#.#..#. ##...#.# #..#####
###....# .#....#. .#......
##.##.## #.#.#..# #####...
###.#### #...#.## ###.#..#
##.#.... #.##.### #...#.##
...##### ###.#... .#####.#
....#..# ...##..# .#.###..
.####... #..#.... .#......

#..#.##. .#..###. #.##....
#.####.. #.####.# .#.###..
###.#.#. ..#.#### ##.#..##
#.####.. ..##..## ######.#
##..##.# ...#...# .#.#.#..
...#..#. .#.#.##. .###.###
.#.#.... #.##.#.. .###.##.
###.#... #..#.##. ######..

.#.#.### .##.##.# ..#.##..
.####.## #.#...## #.#..#.#
..#.#..# ..#.#.#. ####.###
#..####. ..#.#.#. ###.###.
#####..# ####...# ##....##
#.##..#. .#...#.. ####...#
.#.###.. ##..##.. ####.##.
...###.. .##...#. ..#..###
Remove the gaps to form the actual image:

.#.#..#.##...#.##..#####
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
...###...##...#...#..###
Now, you're ready to search for sea monsters! Because your image is monochrome, a sea monster will look like this:

                  #
#    ##    ##    ###
 #  #  #  #  #  #
When looking for this pattern in the image, the spaces can be anything; only the # need to match. Also, you might need to rotate or flip your image before it's oriented correctly to find sea monsters. In the above image, after flipping and rotating it to the appropriate orientation, there are two sea monsters (marked with O):

.####...#####..#...###..
#####..#..#.#.####..#.#.
.#.#...#.###...#.##.O#..
#.O.##.OO#.#.OO.##.OOO##
..#O.#O#.O##O..O.#O##.##
...#.#..##.##...#..#..##
#.##.#..#.#..#..##.#.#..
.###.##.....#...###.#...
#.####.#.#....##.#..#.#.
##...#..#....#..#...####
..#.##...###..#.#####..#
....#.##.#.#####....#...
..##.##.###.....#.##..#.
#...#...###..####....##.
.#.##...#.##.#.#.###...#
#.###.#..####...##..#...
#.###...#.##...#.##O###.
.O##.#OO.###OO##..OOO##.
..O#.O..O..O.#O##O##.###
#.#..##.########..#..##.
#.#####..#.#...##..#....
#....##..#.#########..##
#...#.....#..##...###.##
#..###....##.#...##.##.#
Determine how rough the waters are in the sea monsters' habitat by counting the number of # that are not part of a sea monster. In the above example, the habitat's water roughness is 273.

How many # are not part of a sea monster?

Your puzzle answer was 2453.
*/

func Day20() {
	fmt.Println("--- Day 20: Jurassic Jigsaw ---")
	tiles, err := LoadTiles(day20Input)
	if err != nil {
		fmt.Println(err)
		return
	}
	if !SolveTilePuzzle(tiles) {
		fmt.Println("Failed to solve the puzzle")
	}
	fmt.Println("Corners multiplied after solving:", MultiplyCornerIDs(tiles))
	fmt.Println("Water roughness after seamonster removal:", WaterRoughnessAfterFindingSeaMonsters(NewMonsterPicture(tiles)))
}

// NewMonsterPicture converts a square of tiles into a monster picture. NewMonsterPicture returns nil if
// tiles does for a square.
func NewMonsterPicture(tiles []*Tile) [][]bool {
	w := int(math.Sqrt(float64(len(tiles))))
	if w == 0 || w*w != len(tiles) {
		// not possible to form a square
		return nil
	}

	pic := make([][]bool, 0, w*(len(tiles[0].Grid)-2))
	// for each vertical tile
	for n := 0; n < w && n*w < len(tiles); n++ {
		// for each row minus the borders
		for y := 1; y < len(tiles[n*w].Grid)-1; y++ {
			row := make([]bool, 0, (len(tiles[n*w].Grid[y])-2)*w)
			// for each horizontal tile
			for m := 0; m < w && n*w+m < len(tiles); m++ {
				// add the row minus the borders
				row = append(row, tiles[n*w+m].Grid[y][1:len(tiles[n*w+m].Grid[y])-1]...)
			}
			pic = append(pic, row)
		}
	}
	return pic
}

// WaterRoughnessAfterFindingSeaMonsters will return the water roughness after finding at least one sea
// monster and removing it. WaterRoughnessAfterFindingSeaMonsters return -1 if no sea monsters can be found
// after rotating and flipping the image.
func WaterRoughnessAfterFindingSeaMonsters(pic [][]bool) int {
	// monster looks like (20x3):
	//                   #
	// #    ##    ##    ###
	//  #  #  #  #  #  #

	type Coord struct {
		X, Y int
	}

	const monsterWidth = 20
	const monsterHeight = 3
	monsterCoords := []Coord{{0, 1}, {1, 2}, {4, 2}, {5, 1}, {6, 1}, {7, 2}, {10, 2},
		{11, 1}, {12, 1}, {13, 2}, {16, 2}, {17, 1}, {18, 0}, {18, 1}, {19, 1}}

	// use Tile to represent the pic to allow easily rotating/flipping it
	p := &Tile{Grid: pic}

	// for each rotation/flip possible, check to find (and remove) sea monsters
	for i := 0; i < 8; i++ {
		monstersFound := 0
		for y := 0; y < len(p.Grid)-monsterHeight; y++ {
			for x := 0; x < len(p.Grid[y])-monsterWidth; x++ {
				isMonster := true
				for _, c := range monsterCoords {
					if !p.Grid[y+c.Y][x+c.X] {
						isMonster = false
						break
					}
				}
				if isMonster {
					monstersFound++
					// remove the monster
					for _, c := range monsterCoords {
						p.Grid[y+c.Y][x+c.X] = false
					}
				}
			}
		}

		if monstersFound > 0 {
			// monsters found - no need to look further
			roughWater := 0
			for y := 0; y < len(p.Grid); y++ {
				for x := 0; x < len(p.Grid[y]); x++ {
					if p.Grid[y][x] {
						roughWater++
					}
				}
			}
			return roughWater
		}

		if i == 3 {
			p.Rotate()
			p.FlipX()
		} else {
			p.Rotate()
		}
	}

	return -1
}

// MultiplyCornerIDs treats the tiles as if they are arranged in a square, one row after the other, and
// returns the corner IDs multipied together. MultiplyCornerIDs will return -1 if it is unable to form a
// square from the tiles.
func MultiplyCornerIDs(tiles []*Tile) int {
	w := int(math.Sqrt(float64(len(tiles))))
	if w*w != len(tiles) {
		// not possible to form a square
		return -1
	}
	return tiles[0].ID * tiles[w-1].ID * tiles[w*w-w].ID * tiles[w*w-1].ID
}

// SolveTilePuzzle takes a list of tiles and solve the tile puzzle by rotating, flipping and reordering
// them so they fit into a square where the edges align with each other. SolveTilePuzzle will return true
// if successful and false if no solution could be found.
func SolveTilePuzzle(tiles []*Tile) bool {
	// determine the width (and indirectly height) of the square
	w := int(math.Sqrt(float64(len(tiles))))
	if w*w != len(tiles) {
		// not possible to form a square
		return false
	}

	// we need all possible tile layouts to iterate over for solving the puzzle
	// we could save some space at the cost of runtime here by computing them as we need them
	// rather than pre-create all of them.
	allLayouts := allTileLayouts(tiles)

	// generate a map of left and top edge values of all possible tile layouts to allow easily
	// finding a matching tile when solving it
	leftEdges := map[int][]*Tile{}
	topEdges := map[int][]*Tile{}
	for _, t := range allLayouts {
		leftEdges[t.Edges[LeftEdge]] = append(leftEdges[t.Edges[LeftEdge]], t)
		topEdges[t.Edges[TopEdge]] = append(topEdges[t.Edges[TopEdge]], t)
	}

	tiles = tiles[:0]

	return solveRemainingPuzzle(w, tiles, allLayouts, leftEdges, topEdges)
}

// solveRemainingPuzzle is a recursive function that attempts to place a single tile.
func solveRemainingPuzzle(width int, tiles, allLayouts []*Tile, leftEdges, topEdges map[int][]*Tile) bool {
	if width*width == len(tiles) {
		// all done!
		return true
	}

	if len(tiles) <= 0 {
		// no tiles laid yet, try all possible layouts as the first piece
		for _, t := range allLayouts {
			tiles = tiles[:0]
			tiles = append(tiles, t)
			if solveRemainingPuzzle(width, tiles, allLayouts, leftEdges, topEdges) {
				// found solution! tiles now contain the solved puzzle
				return true
			}
		}
		// no solution could be found
		return false
	}

	isInUse := func(id int) bool {
		for i := range tiles {
			if tiles[i].ID == id {
				return true
			}
		}
		return false
	}

	// a tile matches if it has not been previously used
	// and if its left edge has the same value as the right edge of tiles[len(tiles)-1]
	//   (except when layingt he first tile in a new row (eg len(tiles)%width == 0) in)
	// and its top edge has the same value as the bottom edge of tiles[len(tiles)-width]

	var possibleTiles []*Tile
	if len(tiles) >= width && len(tiles)%width == 0 {
		//fmt.Println("Finding matches from above")
		bottomEdgeAbove := tiles[len(tiles)-width].Edges[BottomEdge]
		possibleTiles = topEdges[bottomEdgeAbove]
	} else {
		//fmt.Println("Finding matches from left")
		rightEdgeToTheLeft := tiles[len(tiles)-1].Edges[RightEdge] // tiles is guaranteed to be at least 1 long
		possibleTiles = leftEdges[rightEdgeToTheLeft]
	}

	for _, possibleTile := range possibleTiles {
		if len(tiles) > width && len(tiles)%width != 0 {
			// tile must fit the tile above it
			if possibleTile.Edges[TopEdge] != tiles[len(tiles)-width].Edges[BottomEdge] {
				continue
			}
		}

		if isInUse(possibleTile.ID) {
			continue
		}

		// place the tile
		tiles = append(tiles, possibleTile)
		if solveRemainingPuzzle(width, tiles, allLayouts, leftEdges, topEdges) {
			return true
		}
		// it didnt work- remove it again
		tiles = tiles[:len(tiles)-1]
	}
	return false
}

// allTileLayouts returns a new list of tiles that contains all possible layouts for the provided
// list of tiles.
func allTileLayouts(tiles []*Tile) []*Tile {
	// each tile can be in 8 different positions each with a unique combination of edges
	// each layout has 4 rotations + 4 additional layout when the tile is flipped.
	// note that it doesn't matter if we flip around the x or y axis, as rotating the
	// flipped tile will produce the same layout.
	allLayouts := make([]*Tile, 0, len(tiles)*16)
	for i := range tiles {
		// as is layout
		allLayouts = append(allLayouts, tiles[i].Copy())
		for j := 0; j < 3; j++ {
			allLayouts = append(allLayouts, allLayouts[len(allLayouts)-1].Copy())
			allLayouts[len(allLayouts)-1].Rotate()
		}
		// flip x
		allLayouts = append(allLayouts, tiles[i].Copy())
		allLayouts[len(allLayouts)-1].FlipX()
		for j := 0; j < 3; j++ {
			allLayouts = append(allLayouts, allLayouts[len(allLayouts)-1].Copy())
			allLayouts[len(allLayouts)-1].Rotate()
		}
	}
	return allLayouts
}

// LoadTiles loads the raw tile data and returns it corresponding tiles. An error is returned
// if the tile data is invalid.
func LoadTiles(tileData string) ([]*Tile, error) {
	var tiles []*Tile

	// make sure the last tile is followed by a blank line to let the loop handle adding it to tiles
	tileData += "\n"

	t := &Tile{}
	for _, line := range strings.Split(tileData, "\n") {
		switch {
		case strings.HasPrefix(line, "Tile "):
			// new tile starting
			id, err := strconv.Atoi(line[5 : len(line)-1])
			if err != nil {
				return nil, errors.Wrapf(err, "bad tile id (line=%s)", line)
			}
			t.ID = id
		case line == "":
			// tile ending- generate integer representation of edges and add the finished tile to tiles
			t.calcEdges()
			tiles = append(tiles, t)
			t = &Tile{}
		default:
			// tile grid line
			row := make([]bool, len(line))
			for i := 0; i < len(line); i++ {
				if line[i] == '#' {
					row[i] = true
				}
			}
			t.Grid = append(t.Grid, row)
		}
	}

	return tiles, nil
}

// Tile represents a single tile.
type Tile struct {
	ID    int
	Grid  [][]bool
	Edges [4]int
}

const (
	TopEdge = iota
	RightEdge
	BottomEdge
	LeftEdge
)

// Copy creates a copy of a tile and returns it.
func (t *Tile) Copy() *Tile {
	new := Tile{
		ID:    t.ID,
		Edges: t.Edges,
	}

	new.Grid = make([][]bool, len(t.Grid))
	for y := range t.Grid {
		new.Grid[y] = make([]bool, len(t.Grid[y]))
		copy(new.Grid[y], t.Grid[y])
	}
	return &new
}

// Rotate rotates the tile clockwise 90 degrees.
func (t *Tile) Rotate() {
	// The tile contains width/2 layers that needs to be rotated independently
	// eg.                                       1  2  3  4
	//      1  2  3                              5  6  7  8
	//      4  5  6                              9 10 11 12
	//      7  8  9  contains 1 layer           13 14 15 16  contains 2 layers
	//               1, 2, 3, 4, 6, 7, 8, 9                  1, 2, 3, 4, 5, 8, 9, 12, 13, 14, 15, 16
	//               (5 doesn't need to be moved)            6, 7, 10, 11
	//
	// We rotate in place by moving four numbers at the same time starting at the corners then moving around the layer.
	//
	// The coordinate of the four can be defined as:
	//      (n+l,l), (width-1-l, n+l), (width-1-n, width-1+l), (0, width-1-n)
	// Where
	//      w = width
	//      l = layer
	//      n = index running from 0 to width-1-2*l
	width := len(t.Grid)
	for l := 0; l < width/2; l++ {
		for n := 0; n < width-1-2*l; n++ {
			x1, y1 := n+l, l
			x2, y2 := width-1-l, n+l
			x3, y3 := width-1-n-l, width-1-l
			x4, y4 := l, width-1-n-l

			t.Grid[y2][x2], t.Grid[y3][x3], t.Grid[y4][x4], t.Grid[y1][x1] =
				t.Grid[y1][x1], t.Grid[y2][x2], t.Grid[y3][x3], t.Grid[y4][x4]
		}
	}
	t.calcEdges()
}

// FlipX flips the tile around the X axis.
func (t *Tile) FlipX() {
	for y := 0; y < len(t.Grid)/2; y++ {
		for x := 0; x < len(t.Grid[y]); x++ {
			t.Grid[y][x], t.Grid[len(t.Grid)-1-y][x] = t.Grid[len(t.Grid)-1-y][x], t.Grid[y][x]
		}
	}
	t.calcEdges()
}

// calcEdges updates the bitfield representation of the edges.
func (t *Tile) calcEdges() {
	width := len(t.Grid)
	t.Edges = [4]int{0, 0, 0, 0}

	currentBit := 1
	for n := width - 1; n >= 0; n-- {
		if t.Grid[0][n] {
			t.Edges[TopEdge] += currentBit
		}
		if t.Grid[n][width-1] {
			t.Edges[RightEdge] += currentBit
		}
		if t.Grid[width-1][n] {
			t.Edges[BottomEdge] += currentBit
		}
		if t.Grid[n][0] {
			t.Edges[LeftEdge] += currentBit
		}
		currentBit *= 2
	}
}

// Print will print the tile to stdout.
func (t Tile) Print(withID bool) {
	if withID {
		fmt.Printf("Tile %d:\n", t.ID)
	}
	for y := 0; y < len(t.Grid); y++ {
		for x := 0; x < len(t.Grid[y]); x++ {
			switch t.Grid[y][x] {
			case true:
				fmt.Print("#")
			case false:
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

// PrintTilesSquare will print the entire square of tiles to stdout.
func PrintTilesSquare(width int, tiles []*Tile) {
	for n := 0; n < width && n*width < len(tiles); n++ {
		for y := 0; y < len(tiles[n*width].Grid); y++ {
			for m := 0; m < width && n*width+m < len(tiles); m++ {
				for x := 0; x < len(tiles[n*width+m].Grid[y]); x++ {
					if tiles[n*width+m].Grid[y][x] {
						fmt.Print("#")
					} else {
						fmt.Print(".")
					}
				}
				fmt.Print(" ")
			}
			fmt.Println()
		}
		fmt.Println()
	}
}
