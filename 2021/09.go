package main

import (
	"fmt"
	"sort"

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
	cave, err := NewCave(day09Input)
	if err != nil {
		fmt.Println(errors.Wrap(err, "Failed to load cave"))
		return
	}
	fmt.Println("Sum of the risk levels of all low points:", LowPointRiskLevels(cave))
	fmt.Println("Size of the three largest basins multiplied:", LargestThreeBasinsMultiplied(cave))
}

// LowPointRiskLevels identifies and sums the risk of each low point within the cave.
func LowPointRiskLevels(cave Cave) int {
	sum := 0

	for y := 0; y < len(cave); y++ {
		for x := 0; x < len(cave[y]); x++ {
			isLowPoint := true
			for _, n := range cave.Neighbors(CaveLoc{x, y}, false) {
				if cave[y][x] >= cave[n.Y][n.X] {
					isLowPoint = false
					break
				}
			}

			if isLowPoint {
				sum += cave[y][x] + 1
			}
		}
	}

	return sum
}

// LargestThreeBasinsMultiplied identifies the various basins within the cave and returns the size
// of the larget three multiplied.
func LargestThreeBasinsMultiplied(cave Cave) int {
	var basinSize []int
	// locToBasinID tracks locations and what basin they belong to (the index of basinSize)
	locToBasinID := map[CaveLoc]int{}

	for y := 0; y < len(cave); y++ {
		for x := 0; x < len(cave[y]); x++ {
			if cave[y][x] == 9 {
				// no basin here
				continue
			}

			pos := CaveLoc{x, y}
			if _, found := locToBasinID[pos]; found {
				// in an already processed basin
				continue
			}

			// in a new basin, lets explore it
			basinSize = append(basinSize, exploreBasin(cave, len(basinSize), pos, locToBasinID))
		}
	}

	if len(basinSize) < 3 {
		return 0
	}

	sort.Ints(basinSize)
	n := len(basinSize)
	return basinSize[n-3] * basinSize[n-2] * basinSize[n-1]
}

// exploreBasin takes a cave, a basin id and start location of a basin and explores it, finding all
// locations belonging to it and updating locToBasinID appropriately before returning the number of
// poistions belonging to the basin.
func exploreBasin(cave Cave, id int, startLoc CaveLoc, locToBasinID map[CaveLoc]int) int {
	size := 0
	var pos CaveLoc
	unvisited := []CaveLoc{startLoc}
	for len(unvisited) > 0 {
		pos, unvisited = unvisited[len(unvisited)-1], unvisited[:len(unvisited)-1]

		if _, found := locToBasinID[pos]; found {
			// already processed
			continue
		}

		if cave[pos.Y][pos.X] == 9 {
			// not a basin
			locToBasinID[pos] = -1
			continue
		}

		locToBasinID[pos] = id
		size++

		for _, n := range cave.Neighbors(pos, false) {
			unvisited = append(unvisited, n)
		}
	}

	return size
}
