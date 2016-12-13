package main

import (
	"fmt"
)

/*

--- Day 13: A Maze of Twisty Little Cubicles ---

You arrive at the first floor of this new building to discover a much less welcoming environment than the shiny atrium of the last one. Instead, you are in a maze of twisty little cubicles, all alike.

Every location in this area is addressed by a pair of non-negative integers (x,y). Each such coordinate is either a wall or an open space. You can't move diagonally. The cube maze starts at 0,0 and seems to extend infinitely toward positive x and y; negative values are invalid, as they represent a location outside the building. You are in a small waiting area at 1,1.

While it seems chaotic, a nearby morale-boosting poster explains, the layout is actually quite logical. You can determine whether a given x,y coordinate will be a wall or an open space using a simple system:

Find x*x + 3*x + 2*x*y + y + y*y.
Add the office designer's favorite number (your puzzle input).
Find the binary representation of that sum; count the number of bits that are 1.
If the number of bits that are 1 is even, it's an open space.
If the number of bits that are 1 is odd, it's a wall.
For example, if the office designer's favorite number were 10, drawing walls as # and open spaces as ., the corner of the building containing 0,0 would look like this:

  0123456789
0 .#.####.##
1 ..#..#...#
2 #....##...
3 ###.#.###.
4 .##..#..#.
5 ..##....#.
6 #...##.###
Now, suppose you wanted to reach 7,4. The shortest route you could take is marked as O:

  0123456789
0 .#.####.##
1 .O#..#...#
2 #OOO.##...
3 ###O#.###.
4 .##OO#OO#.
5 ..##OOO.#.
6 #...##.###
Thus, reaching 7,4 would take a minimum of 11 steps (starting from your current location, 1,1).

What is the fewest number of steps required for you to reach 31,39?

Your puzzle answer was 92.

--- Part Two ---

How many locations (distinct x,y coordinates, including your starting location) can you reach in at most 50 steps?

*/

const mazeWidth = 50
const mazeHeight = 50

func main() {
	fmt.Println("Advent of Code 2016 - Day 13")

	maze := NewMaze()
	fmt.Println("Distance to 31,39:", maze.Distance(31, 39))
	fmt.Println()
	maze.Print()
}

type Maze [mazeHeight][mazeWidth]bool

func NewMaze() Maze {
	m := Maze{}
	for y := 0; y < mazeHeight; y++ {
		for x := 0; x < mazeWidth; x++ {
			// count bits
			if Bits(x*x+3*x+2*x*y+y+y*y+puzzleInput)%2 != 0 {
				m[y][x] = true
			}
		}
	}
	return m
}

func (m *Maze) Print() {
	for y := 0; y < mazeWidth; y++ {
		for x := 0; x < mazeWidth; x++ {
			if m[y][x] {
				fmt.Print("#")
			} else {
				if _, found := beenTo[Coord{x: x, y: y}]; found {
					fmt.Print("O")
				} else {
					fmt.Print(".")
				}
			}
		}
		fmt.Print("\n")
	}
	fmt.Print("\n")
}

var beenTo = map[Coord]struct{}{}

func (m *Maze) Distance(x, y int) int {
	next := []Coord{{x: 1, y: 1}}
	beenTo[next[0]] = struct{}{}

	moves := -1
	for len(next) > 0 {
		moves++
		if moves == 50 {
			fmt.Println("Locations reached in 50 steps:", len(beenTo))
		}

		current := next
		next = []Coord{}

		for _, c := range current {
			if c.x == x && c.y == y {
				// found it!
				return moves
			}

			addCoord := func(c Coord) {
				if c.x < 0 || c.x >= mazeWidth || c.y < 0 || c.y >= mazeHeight || m[c.y][c.x] {
					// invalid pos
					return
				}
				if _, found := beenTo[c]; found {
					// already been there
					return
				}
				// good
				beenTo[c] = struct{}{}
				next = append(next, c)
			}

			up := c
			up.y += 1
			addCoord(up)

			down := c
			down.y -= 1
			addCoord(down)

			left := c
			left.x -= 1
			addCoord(left)

			right := c
			right.x += 1
			addCoord(right)
		}
	}

	return 0
}

type Coord struct {
	x, y int
}

func Bits(i int) int {
	count := 0
	for i > 0 {
		if i&1 == 1 {
			count++
		}
		i = i >> 1
	}
	return count
}

const puzzleInput = 1350
