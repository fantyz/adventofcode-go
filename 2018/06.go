package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

/*

--- Day 6: Chronal Coordinates ---
The device on your wrist beeps several times, and once again you feel like you're falling.

"Situation critical," the device announces. "Destination indeterminate. Chronal interference detected. Please specify new target coordinates."

The device then produces a list of coordinates (your puzzle input). Are they places it thinks are safe or dangerous? It recommends you check manual page 729. The Elves did not give you a manual.

If they're dangerous, maybe you can minimize the danger by finding the coordinate that gives the largest distance from the other points.

Using only the Manhattan distance, determine the area around each coordinate by counting the number of integer X,Y locations that are closest to that coordinate (and aren't tied in distance to any other coordinate).

Your goal is to find the size of the largest area that isn't infinite. For example, consider the following list of coordinates:

1, 1
1, 6
8, 3
3, 4
5, 5
8, 9
If we name these coordinates A through F, we can draw them on a grid, putting 0,0 at the top left:

..........
.A........
..........
........C.
...D......
.....E....
.B........
..........
..........
........F.
This view is partial - the actual grid extends infinitely in all directions. Using the Manhattan distance, each location's closest coordinate can be determined, shown here in lowercase:

aaaaa.cccc
aAaaa.cccc
aaaddecccc
aadddeccCc
..dDdeeccc
bb.deEeecc
bBb.eeee..
bbb.eeefff
bbb.eeffff
bbb.ffffFf
Locations shown as . are equally far from two or more coordinates, and so they don't count as being closest to any.

In this example, the areas of coordinates A, B, C, and F are infinite - while not shown here, their areas extend forever outside the visible grid. However, the areas of coordinates D and E are finite: D is closest to 9 locations, and E is closest to 17 (both including the coordinate's location itself). Therefore, in this example, the size of the largest area is 17.

What is the size of the largest area that isn't infinite?

Your puzzle answer was 3604.

--- Part Two ---
On the other hand, if the coordinates are safe, maybe the best you can do is try to find a region near as many coordinates as possible.

For example, suppose you want the sum of the Manhattan distance to all of the coordinates to be less than 32. For each location, add up the distances to all of the given coordinates; if the total of those distances is less than 32, that location is within the desired region. Using the same coordinates as above, the resulting region looks like this:

..........
.A........
..........
...###..C.
..#D###...
..###E#...
.B.###....
..........
..........
........F.
In particular, consider the highlighted location 4,3 located at the top middle of the region. Its calculation is as follows, where abs() is the absolute value function:

Distance to coordinate A: abs(4-1) + abs(3-1) =  5
Distance to coordinate B: abs(4-1) + abs(3-6) =  6
Distance to coordinate C: abs(4-8) + abs(3-3) =  4
Distance to coordinate D: abs(4-3) + abs(3-4) =  2
Distance to coordinate E: abs(4-5) + abs(3-5) =  3
Distance to coordinate F: abs(4-8) + abs(3-9) = 10
Total distance: 5 + 6 + 4 + 2 + 3 + 10 = 30
Because the total distance to all coordinates (30) is less than 32, the location is within the region.

This region, which also includes coordinates D and E, has a total size of 16.

Your actual region will need to be much larger than this example, though, instead including all locations with a total distance of less than 10000.

What is the size of the region containing all locations which have a total distance to all given coordinates of less than 10000?

*/

func main() {
	fmt.Println("Day 6: Chronal Coordinates")
	coords := InputToCoordinates(puzzleInput)
	fmt.Println(" > Largest area:", LargestArea(coords))
	fmt.Println(" > Largest Region:", RegionWithinMaxDist(10000, coords))
}

type Coord struct {
	X, Y int
}

func InputToCoordinates(in string) []Coord {
	coords := []Coord{}
	re := regexp.MustCompile(`^(\d+), (\d+)$`)
	for _, l := range strings.Split(in, "\n") {
		m := re.FindStringSubmatch(l)
		x, err := strconv.Atoi(m[1])
		if err != nil {
			panic(err)
		}
		y, err := strconv.Atoi(m[2])
		if err != nil {
			panic(err)
		}
		coords = append(coords, Coord{x, y})
	}
	return coords
}

func RegionWithinMaxDist(maxDist int, coords []Coord) int {
	xmax, ymax := 0, 0
	for _, c := range coords {
		if c.X > xmax {
			xmax = c.X + 1
		}
		if c.Y > ymax {
			ymax = c.Y + 1
		}
	}

	count := 0
	for y := 0; y <= ymax; y++ {
	next:
		for x := 0; x <= xmax; x++ {
			sum := 0
			for _, c := range coords {
				sum += dist(c, x, y)
				if sum >= maxDist {
					continue next
				}
			}
			count++
		}
	}

	return count
}

func LargestArea(coords []Coord) int {
	xmax, ymax := 0, 0
	for _, c := range coords {
		if c.X > xmax {
			xmax = c.X + 1
		}
		if c.Y > ymax {
			ymax = c.Y + 1
		}
	}

	f := make([][]int, ymax)
	for i := range f {
		f[i] = make([]int, xmax)
		for j := range f[i] {
			f[i][j] = -1
		}
	}

	infIds := map[int]struct{}{}
	areas := map[int]int{}
	for y := range f {
		for x := range f[y] {
			min := 99999
			minID := -1
			for id, c := range coords {
				if dist(c, x, y) < min {
					minID = id
					min = dist(c, x, y)
				}
			}
			areas[minID]++

			if y == 0 || y == len(f)-1 || x == 0 || x == len(f[y])-1 {
				infIds[minID] = struct{}{}
			}
		}
	}

	max := 0
	for id, size := range areas {
		if size > max {
			if _, found := infIds[id]; found {
				continue
			}
			max = size
		}
	}

	return max
}

func dist(c Coord, x, y int) int {
	xp, yp := 0, 0
	if c.X > x {
		xp = c.X - x
	} else {
		xp = x - c.X
	}
	if c.Y > y {
		yp = c.Y - y
	} else {
		yp = y - c.Y
	}
	return xp + yp
}

const puzzleInput = `252, 125
128, 333
89, 324
141, 171
266, 338
117, 175
160, 236
234, 202
165, 192
204, 232
83, 192
229, 178
333, 57
70, 243
108, 350
161, 63
213, 277
87, 299
163, 68
135, 312
290, 87
73, 246
283, 146
80, 357
66, 312
159, 214
221, 158
175, 54
298, 342
348, 162
249, 90
189, 322
311, 181
194, 244
53, 295
80, 301
262, 332
268, 180
139, 287
115, 53
163, 146
220, 268
79, 85
95, 112
349, 296
179, 274
113, 132
158, 264
316, 175
268, 215`
