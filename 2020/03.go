package main

import (
	"fmt"
	"strings"
)

func init() { days["3"] = Day3 }

/*
--- Day 3: Toboggan Trajectory ---
With the toboggan login problems resolved, you set off toward the airport. While travel by toboggan might be easy, it's certainly not safe: there's very minimal steering and the area is covered in trees. You'll need to see which angles will take you near the fewest trees.

Due to the local geology, trees in this area only grow on exact integer coordinates in a grid. You make a map (your puzzle input) of the open squares (.) and trees (#) you can see. For example:

..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#
These aren't the only trees, though; due to something you read about once involving arboreal genetics and biome stability, the same pattern repeats to the right many times:

..##.........##.........##.........##.........##.........##.......  --->
#...#...#..#...#...#..#...#...#..#...#...#..#...#...#..#...#...#..
.#....#..#..#....#..#..#....#..#..#....#..#..#....#..#..#....#..#.
..#.#...#.#..#.#...#.#..#.#...#.#..#.#...#.#..#.#...#.#..#.#...#.#
.#...##..#..#...##..#..#...##..#..#...##..#..#...##..#..#...##..#.
..#.##.......#.##.......#.##.......#.##.......#.##.......#.##.....  --->
.#.#.#....#.#.#.#....#.#.#.#....#.#.#.#....#.#.#.#....#.#.#.#....#
.#........#.#........#.#........#.#........#.#........#.#........#
#.##...#...#.##...#...#.##...#...#.##...#...#.##...#...#.##...#...
#...##....##...##....##...##....##...##....##...##....##...##....#
.#..#...#.#.#..#...#.#.#..#...#.#.#..#...#.#.#..#...#.#.#..#...#.#  --->
You start on the open square (.) in the top-left corner and need to reach the bottom (below the bottom-most row on your map).

The toboggan can only follow a few specific slopes (you opted for a cheaper model that prefers rational numbers); start by counting all the trees you would encounter for the slope right 3, down 1:

From your starting position at the top-left, check the position that is right 3 and down 1. Then, check the position that is right 3 and down 1 from there, and so on until you go past the bottom of the map.

The locations you'd check in the above example are marked here with O where there was an open square and X where there was a tree:

..##.........##.........##.........##.........##.........##.......  --->
#..O#...#..#...#...#..#...#...#..#...#...#..#...#...#..#...#...#..
.#....X..#..#....#..#..#....#..#..#....#..#..#....#..#..#....#..#.
..#.#...#O#..#.#...#.#..#.#...#.#..#.#...#.#..#.#...#.#..#.#...#.#
.#...##..#..X...##..#..#...##..#..#...##..#..#...##..#..#...##..#.
..#.##.......#.X#.......#.##.......#.##.......#.##.......#.##.....  --->
.#.#.#....#.#.#.#.O..#.#.#.#....#.#.#.#....#.#.#.#....#.#.#.#....#
.#........#.#........X.#........#.#........#.#........#.#........#
#.##...#...#.##...#...#.X#...#...#.##...#...#.##...#...#.##...#...
#...##....##...##....##...#X....##...##....##...##....##...##....#
.#..#...#.#.#..#...#.#.#..#...X.#.#..#...#.#.#..#...#.#.#..#...#.#  --->
In this example, traversing the map using this slope would cause you to encounter 7 trees.

Starting at the top-left corner of your map and following a slope of right 3 and down 1, how many trees would you encounter?

Your puzzle answer was 244.

--- Part Two ---
Time to check the rest of the slopes - you need to minimize the probability of a sudden arboreal stop, after all.

Determine the number of trees you would encounter if, for each of the following slopes, you start at the top-left corner and traverse the map all the way to the bottom:

Right 1, down 1.
Right 3, down 1. (This is the slope you already checked.)
Right 5, down 1.
Right 7, down 1.
Right 1, down 2.
In the above example, these slopes would find 2, 7, 3, 4, and 2 tree(s) respectively; multiplied together, these produce the answer 336.

What do you get if you multiply together the number of trees encountered on each of the listed slopes?

Your puzzle answer was 9406609920.
*/

func Day3() {
	fmt.Println("--- Day 3: Toboggan Trajectory ---")
	fmt.Println("Trees encountered:", TraverseForestWithRoutes(day3Input, []Route{{3, 1}}))
	fmt.Println("Trees encountered multipled across many routes:", TraverseForestWithRoutes(day3Input, []Route{{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2}}))
}

// NewForest returns a forest based on the provided map of trees.
func NewForest(trees string) Forest {
	var forest [][]bool
	for _, row := range strings.Split(trees, "\n") {
		forestRow := make([]bool, len(row))
		for x, tree := range row {
			if tree == '#' {
				forestRow[x] = true
			}
		}
		forest = append(forest, forestRow)
	}
	return forest
}

type Forest [][]bool // y,x -> hasTree

// HasTree takes a coordinate and returns whether the coordinate has a tree and whether the
// coordinate is out of bounds.
func (f Forest) HasTree(x, y int) (bool, bool) {
	if x < 0 || y < 0 || y >= len(f) {
		return false, true
	}
	return f[y][x%len(f[y])], false
}

// Route represent a specific route through the forest.
type Route struct {
	XStep, YStep int
}

// Traverse will step through the forest starting from (x,y) = (0,0) taking the specified route
// until it leaves the forest. It returns the number of trees encountered along the way.
func (f Forest) Traverse(route Route) int {
	treesEncoutered := 0
	x, y := 0, 0

	for {
		tree, done := f.HasTree(x, y)
		if done {
			// went out of bounds
			break
		}
		if tree {
			treesEncoutered++
		}
		x += route.XStep
		y += route.YStep
	}

	return treesEncoutered
}

// TraverseForestWithRoutes takes a map of trees and a list of routes and returns the number
// of trees encountered on each route multiplied together.
func TraverseForestWithRoutes(trees string, routes []Route) int {
	if len(routes) <= 0 {
		return 0
	}

	f := NewForest(trees)

	total := 1
	for _, route := range routes {
		total *= f.Traverse(route)
	}
	return total
}
