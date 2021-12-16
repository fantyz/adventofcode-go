package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

func init() { days["9"] = Day9 }

/*
--- Day 9: Smoke Basin ---
These caves seem to be lava tubes. Parts are even still volcanically active; small hydrothermal vents release smoke into the caves that slowly settles like rain.

If you can model how the smoke flows through the caves, you might be able to avoid it and be that much safer. The submarine generates a heightmap of the floor of the nearby caves for you (your puzzle input).

Smoke flows to the lowest point of the area it's in. For example, consider the following heightmap:

2199943210
3987894921
9856789892
8767896789
9899965678
Each number corresponds to the height of a particular location, where 9 is the highest and 0 is the lowest a location can be.

Your first goal is to find the low points - the locations that are lower than any of its adjacent locations. Most locations have four adjacent locations (up, down, left, and right); locations on the edge or corner of the map have three or two adjacent locations, respectively. (Diagonal locations do not count as adjacent.)

In the above example, there are four low points, all highlighted: two are in the first row (a 1 and a 0), one is in the third row (a 5), and one is in the bottom row (also a 5). All other locations on the heightmap have some lower adjacent location, and so are not low points.

The risk level of a low point is 1 plus its height. In the above example, the risk levels of the low points are 2, 1, 6, and 6. The sum of the risk levels of all low points in the heightmap is therefore 15.

Find all of the low points on your heightmap. What is the sum of the risk levels of all low points on your heightmap?

Your puzzle answer was 478.

--- Part Two ---
Next, you need to find the largest basins so you know what areas are most important to avoid.

A basin is all locations that eventually flow downward to a single low point. Therefore, every low point has a basin, although some basins are very small. Locations of height 9 do not count as being in any basin, and all other locations will always be part of exactly one basin.

The size of a basin is the number of locations within the basin, including the low point. The example above has four basins.

The top-left basin, size 3:

2199943210
3987894921
9856789892
8767896789
9899965678
The top-right basin, size 9:

2199943210
3987894921
9856789892
8767896789
9899965678
The middle basin, size 14:

2199943210
3987894921
9856789892
8767896789
9899965678
The bottom-right basin, size 9:

2199943210
3987894921
9856789892
8767896789
9899965678
Find the three largest basins and multiply their sizes together. In the above example, this is 9 * 14 * 9 = 1134.

What do you get if you multiply together the sizes of the three largest basins?

Your puzzle answer was 1327014.
*/

func Day9() {
	fmt.Println("--- Day 9: Smoke Basin ---")
	cave, err := NewLavaTubesCave(day09Input)
	if err != nil {
		fmt.Println(errors.Wrap(err, "Failed to load cave"))
		return
	}
	fmt.Println("Sum of the risk levels of all low points:", cave.LowPointRiskLevels())
	fmt.Println("Size of the three largest basins multiplied:", cave.LargestThreeBasinsMultiplied())
}

// LavaTubesCave is a hightmap representing the cave.
type LavaTubesCave [][]int

// NewLavaTubesCave takes the puzzle input and returns a NewLavaTubesCave.
// NewLavaTubesCave returns an error on any unexpected in the input.
func NewLavaTubesCave(in string) (LavaTubesCave, error) {
	var c LavaTubesCave
	width := -1

	for _, line := range strings.Split(in, "\n") {
		if width < 0 {
			width = len(line)
		}
		if len(line) != width {
			return nil, errors.Errorf("unexpected length of row in lava tubes cave (count=%d, expected=%d)", len(line), width)
		}

		row := make([]int, 0, width)
		for _, ch := range line {
			val, err := strconv.Atoi(string(ch))
			if err != nil {
				return nil, errors.Wrapf(err, "bad lava tube cave value (val=%s)", string(ch))
			}
			row = append(row, val)
		}
		c = append(c, row)
	}

	return c, nil
}

// LavaTubesCaveLoc represents a position within the cave.
type LavaTubesCaveLoc struct {
	X, Y int
}

// LowPointRiskLevels identifies and sums the risk of each low point within the cave.
func (c LavaTubesCave) LowPointRiskLevels() int {
	sum := 0

	for y := 0; y < len(c); y++ {
		for x := 0; x < len(c[y]); x++ {
			neighbors := []LavaTubesCaveLoc{
				{x - 1, y},
				{x, y + 1},
				{x + 1, y},
				{x, y - 1},
			}

			isLowPoint := true
			for _, n := range neighbors {
				if n.X < 0 || n.X >= len(c[0]) || n.Y < 0 || n.Y >= len(c) {
					continue
				}
				if c[y][x] >= c[n.Y][n.X] {
					isLowPoint = false
					break
				}
			}

			if isLowPoint {
				sum += c[y][x] + 1
			}
		}
	}

	return sum

}

// LargestThreeBasinsMultiplied identifies the various basins within the cave and returns the size
// of the larget three multiplied.
func (c LavaTubesCave) LargestThreeBasinsMultiplied() int {
	var basinSize []int
	// locToBasinID tracks locations and what basin they belong to (the index of basinSize)
	locToBasinID := map[LavaTubesCaveLoc]int{}

	for y := 0; y < len(c); y++ {
		for x := 0; x < len(c[y]); x++ {
			if c[y][x] == 9 {
				// no basin here
				continue
			}

			pos := LavaTubesCaveLoc{x, y}
			if _, found := locToBasinID[pos]; found {
				// in an already processed basin
				continue
			}

			// in a new basin, lets explore it
			basinSize = append(basinSize, c.exploreBasin(len(basinSize), pos, locToBasinID))
		}
	}

	if len(basinSize) < 3 {
		return 0
	}

	sort.Ints(basinSize)
	n := len(basinSize)
	return basinSize[n-3] * basinSize[n-2] * basinSize[n-1]
}

// exploreBasin takes a basin id and start location of a basin and explores it, finding all locations
// belonging to it and updating locToBasinID appropriately before returning the number of poistions
// belonging to the basin.
func (c LavaTubesCave) exploreBasin(id int, startLoc LavaTubesCaveLoc, locToBasinID map[LavaTubesCaveLoc]int) int {
	size := 0
	var pos LavaTubesCaveLoc
	unvisited := []LavaTubesCaveLoc{startLoc}
	for len(unvisited) > 0 {
		pos, unvisited = unvisited[len(unvisited)-1], unvisited[:len(unvisited)-1]

		if _, found := locToBasinID[pos]; found {
			// already processed
			continue
		}

		if c[pos.Y][pos.X] == 9 {
			// not a basin
			locToBasinID[pos] = -1
			continue
		}

		locToBasinID[pos] = id
		size++

		neighbors := []LavaTubesCaveLoc{
			{pos.X - 1, pos.Y},
			{pos.X, pos.Y + 1},
			{pos.X + 1, pos.Y},
			{pos.X, pos.Y - 1},
		}
		for _, n := range neighbors {
			if n.X < 0 || n.X >= len(c[0]) || n.Y < 0 || n.Y >= len(c) {
				continue
			}
			unvisited = append(unvisited, n)
		}
	}

	return size
}
