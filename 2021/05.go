package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

func init() { days["5"] = Day5 }

/*
--- Day 5: Hydrothermal Venture ---
You come across a field of hydrothermal vents on the ocean floor! These vents constantly produce large, opaque clouds, so it would be best to avoid them if possible.

They tend to form in lines; the submarine helpfully produces a list of nearby lines of vents (your puzzle input) for you to review. For example:

0,9 -> 5,9
8,0 -> 0,8
9,4 -> 3,4
2,2 -> 2,1
7,0 -> 7,4
6,4 -> 2,0
0,9 -> 2,9
3,4 -> 1,4
0,0 -> 8,8
5,5 -> 8,2
Each line of vents is given as a line segment in the format x1,y1 -> x2,y2 where x1,y1 are the coordinates of one end the line segment and x2,y2 are the coordinates of the other end. These line segments include the points at both ends. In other words:

An entry like 1,1 -> 1,3 covers points 1,1, 1,2, and 1,3.
An entry like 9,7 -> 7,7 covers points 9,7, 8,7, and 7,7.
For now, only consider horizontal and vertical lines: lines where either x1 = x2 or y1 = y2.

So, the horizontal and vertical lines from the above list would produce the following diagram:

.......1..
..1....1..
..1....1..
.......1..
.112111211
..........
..........
..........
..........
222111....
In this diagram, the top left corner is 0,0 and the bottom right corner is 9,9. Each position is shown as the number of lines which cover that point or . if no line covers that point. The top-left pair of 1s, for example, comes from 2,2 -> 2,1; the very bottom row is formed by the overlapping lines 0,9 -> 5,9 and 0,9 -> 2,9.

To avoid the most dangerous areas, you need to determine the number of points where at least two lines overlap. In the above example, this is anywhere in the diagram with a 2 or larger - a total of 5 points.

Consider only horizontal and vertical lines. At how many points do at least two lines overlap?

Your puzzle answer was 5585.

--- Part Two ---
Unfortunately, considering only horizontal and vertical lines doesn't give you the full picture; you need to also consider diagonal lines.

Because of the limits of the hydrothermal vent mapping system, the lines in your list will only ever be horizontal, vertical, or a diagonal line at exactly 45 degrees. In other words:

An entry like 1,1 -> 3,3 covers points 1,1, 2,2, and 3,3.
An entry like 9,7 -> 7,9 covers points 9,7, 8,8, and 7,9.
Considering all lines from the above example would now produce the following diagram:

1.1....11.
.111...2..
..2.1.111.
...1.2.2..
.112313211
...1.2....
..1...1...
.1.....1..
1.......1.
222111....
You still need to determine the number of points where at least two lines overlap. In the above example, this is still anywhere in the diagram with a 2 or larger - now a total of 12 points.

Consider all of the lines. At how many points do at least two lines overlap?

Your puzzle answer was 17193.
*/

func Day5() {
	fmt.Println("--- Day 5: Hydrothermal Venture ---")
	vents, err := NewHydrothermalVents(day05Input)
	if err != nil {
		fmt.Println(errors.Wrap(err, "Failed to load hydrothermal vents"))
		return
	}
	fmt.Println("Points that has at least one overlap (without diagnonal):", FindHydrothermalVentOverlaps(vents, true))
	fmt.Println("Points that has at least one overlap (with diagnonals):", FindHydrothermalVentOverlaps(vents, false))
}

// NewHydrothermalVents takes the puzzle input and returns the hydrothermal vents it describes.
// NewHydrothermalVents will return an error if the input contains anything unexpected.
func NewHydrothermalVents(in string) ([]HydrothermalVent, error) {
	ventExp := regexp.MustCompile(`^([0-9]+),([0-9]+) -> ([0-9]+),([0-9]+)$`)
	var vents []HydrothermalVent

	for _, line := range strings.Split(in, "\n") {
		m := ventExp.FindStringSubmatch(line)
		if len(m) != 5 {
			return nil, errors.Errorf("line in input did not match (line=%s)", line)
		}

		x1, err := strconv.Atoi(m[1])
		if err != nil {
			// should never happen
			panic(fmt.Sprintf("x1 not a number! (x1=%s)", m[1]))
		}
		y1, err := strconv.Atoi(m[2])
		if err != nil {
			// should never happen
			panic(fmt.Sprintf("y1 not a number! (y1=%s)", m[2]))
		}
		x2, err := strconv.Atoi(m[3])
		if err != nil {
			// should never happen
			panic(fmt.Sprintf("x1 not a number! (x2=%s)", m[3]))
		}
		y2, err := strconv.Atoi(m[4])
		if err != nil {
			// should never happen
			panic(fmt.Sprintf("x1 not a number! (y2=%s)", m[4]))
		}

		// only allow horizontal, vertical and diagonal vents
		switch {
		case x1 == x2:
			// horizontal
		case y1 == y2:
			// vertical
		case x1-x2 == y1-y2 || x1-x2 == -(y1-y2):
			// diagnoal
		default:
			return nil, errors.Errorf("only horizontal, vertical and diagonal vents allowed (line=%s)", line)
		}

		vents = append(vents, HydrothermalVent{x1, y1, x2, y2})
	}

	return vents, nil
}

// HydrothermalVent is represented as a line from x1,y1 -> x2,y2.
type HydrothermalVent struct {
	X1, Y1 int
	X2, Y2 int
}

// FindHydrothermalVentOverlaps takes a slice of HydrothermalVent and finds any points in which the
// two or more of the vents overlap. It returns the number of such points.
func FindHydrothermalVentOverlaps(vents []HydrothermalVent, ignoreDiagonals bool) int {
	// find the higest value
	max := 0
	for _, vent := range vents {
		if vent.X1 > max {
			max = vent.X1
		}
		if vent.X2 > max {
			max = vent.X2
		}
		if vent.Y1 > max {
			max = vent.Y1
		}
		if vent.Y2 > max {
			max = vent.Y2
		}
	}

	// create a 2D slice to map the vents onto
	m := make([][]int, max+1)
	for i := range m {
		m[i] = make([]int, max+1)
	}

	// put the vents on the map
	for _, vent := range vents {
		xMod, yMod := 1, 1

		if vent.X1 == vent.X2 {
			xMod = 0
		}
		if vent.X1 > vent.X2 {
			xMod = -1
		}
		if vent.Y1 == vent.Y2 {
			yMod = 0
		}
		if vent.Y1 > vent.Y2 {
			yMod = -1
		}

		if ignoreDiagonals && xMod != 0 && yMod != 0 {
			continue
		}

		x, y := vent.X1, vent.Y1

		for {
			m[y][x]++

			if x == vent.X2 && y == vent.Y2 {
				// done
				break
			}

			x += xMod
			y += yMod

		}
	}

	// count overlaps
	overlaps := 0
	for y := 0; y < max+1; y++ {
		for x := 0; x < max+1; x++ {
			if m[y][x] > 1 {
				overlaps++
			}
		}
	}

	return overlaps
}
