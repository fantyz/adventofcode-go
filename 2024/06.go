package main

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

func init() { days["06"] = Day06 }

/*
--- Day 6: Guard Gallivant ---
The Historians use their fancy device again, this time to whisk you all away to the North Pole prototype suit manufacturing lab... in the year 1518! It turns out that having direct access to history is very convenient for a group of historians.

You still have to be careful of time paradoxes, and so it will be important to avoid anyone from 1518 while The Historians search for the Chief. Unfortunately, a single guard is patrolling this part of the lab.

Maybe you can work out where the guard will go ahead of time so that The Historians can search safely?

You start by making a map (your puzzle input) of the situation. For example:

....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...
The map shows the current position of the guard with ^ (to indicate the guard is currently facing up from the perspective of the map). Any obstructions - crates, desks, alchemical reactors, etc. - are shown as #.

Lab guards in 1518 follow a very strict patrol protocol which involves repeatedly following these steps:

If there is something directly in front of you, turn right 90 degrees.
Otherwise, take a step forward.
Following the above protocol, the guard moves up several times until she reaches an obstacle (in this case, a pile of failed suit prototypes):

....#.....
....^....#
..........
..#.......
.......#..
..........
.#........
........#.
#.........
......#...
Because there is now an obstacle in front of the guard, she turns right before continuing straight in her new facing direction:

....#.....
........>#
..........
..#.......
.......#..
..........
.#........
........#.
#.........
......#...
Reaching another obstacle (a spool of several very long polymers), she turns right again and continues downward:

....#.....
.........#
..........
..#.......
.......#..
..........
.#......v.
........#.
#.........
......#...
This process continues for a while, but the guard eventually leaves the mapped area (after walking past a tank of universal solvent):

....#.....
.........#
..........
..#.......
.......#..
..........
.#........
........#.
#.........
......#v..
By predicting the guard's route, you can determine which specific positions in the lab will be in the patrol path. Including the guard's starting position, the positions visited by the guard before leaving the area are marked with an X:

....#.....
....XXXXX#
....X...X.
..#.X...X.
..XXXXX#X.
..X.X.X.X.
.#XXXXXXX.
.XXXXXXX#.
#XXXXXXX..
......#X..
In this example, the guard will visit 41 distinct positions on your map.

Predict the path of the guard. How many distinct positions will the guard visit before leaving the mapped area?

Your puzzle answer was 5208.

--- Part Two ---
While The Historians begin working around the guard's patrol route, you borrow their fancy device and step outside the lab. From the safety of a supply closet, you time travel through the last few months and record the nightly status of the lab's guard post on the walls of the closet.

Returning after what seems like only a few seconds to The Historians, they explain that the guard's patrol area is simply too large for them to safely search the lab without getting caught.

Fortunately, they are pretty sure that adding a single new obstruction won't cause a time paradox. They'd like to place the new obstruction in such a way that the guard will get stuck in a loop, making the rest of the lab safe to search.

To have the lowest chance of creating a time paradox, The Historians would like to know all of the possible positions for such an obstruction. The new obstruction can't be placed at the guard's starting position - the guard is there right now and would notice.

In the above example, there are only 6 different positions where a new obstruction would cause the guard to get stuck in a loop. The diagrams of these six situations use O to mark the new obstruction, | to show a position where the guard moves up/down, - to show a position where the guard moves left/right, and + to show a position where the guard moves both up/down and left/right.

Option one, put a printing press next to the guard's starting position:

....#.....
....+---+#
....|...|.
..#.|...|.
....|..#|.
....|...|.
.#.O^---+.
........#.
#.........
......#...
Option two, put a stack of failed suit prototypes in the bottom right quadrant of the mapped area:


....#.....
....+---+#
....|...|.
..#.|...|.
..+-+-+#|.
..|.|.|.|.
.#+-^-+-+.
......O.#.
#.........
......#...
Option three, put a crate of chimney-squeeze prototype fabric next to the standing desk in the bottom right quadrant:

....#.....
....+---+#
....|...|.
..#.|...|.
..+-+-+#|.
..|.|.|.|.
.#+-^-+-+.
.+----+O#.
#+----+...
......#...
Option four, put an alchemical retroencabulator near the bottom left corner:

....#.....
....+---+#
....|...|.
..#.|...|.
..+-+-+#|.
..|.|.|.|.
.#+-^-+-+.
..|...|.#.
#O+---+...
......#...
Option five, put the alchemical retroencabulator a bit to the right instead:

....#.....
....+---+#
....|...|.
..#.|...|.
..+-+-+#|.
..|.|.|.|.
.#+-^-+-+.
....|.|.#.
#..O+-+...
......#...
Option six, put a tank of sovereign glue right next to the tank of universal solvent:

....#.....
....+---+#
....|...|.
..#.|...|.
..+-+-+#|.
..|.|.|.|.
.#+-^-+-+.
.+----++#.
#+----++..
......#O..
It doesn't really matter what you choose to use as an obstacle so long as you and The Historians can put it into position without the guard noticing. The important thing is having enough options that you can find one that minimizes time paradoxes, and in this example, there are 6 different positions you could choose.

You need to get the guard stuck in a loop by adding a single new obstruction. How many different positions could you choose for this obstruction?

Your puzzle answer was 1972.
*/

func Day06() {
	fmt.Println("--- Day 6: Guard Gallivant ---")
	lab, err := NewLabMap(day06Input)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Part 1: Locations visited by the guard:", len(lab.GuardPatrolPositions()))
	fmt.Println("Part 2: Number of ways a guard could get trapped in a loop:", lab.NumberOfWaysToTrapGuard())
}

type LabMap struct {
	m [][]bool
	guardPos Coord
	guardDir Vector
}

func NewLabMap(in string) (*LabMap, error) {
	// note the y axis is flipped, so the y-axis moves from down to up, hence guards nort starting direction is (0,-1)
	lab := LabMap{nil, Coord{-1,-1}, Vector{0, -1}}

	for y, line := range strings.Split(in, "\n") {
		row := make([]bool, len(line))
		for x := range line {
			switch line[x] {
			case '.':
				// empty, do nothing
			case '#':
				// obstacle
				row[x] = true
			case '^':
				// guard starting position, otherwise empty
				if lab.guardPos.X != -1 {
					return nil, errors.Errorf("multiple starting positions (pos1=%s, pos2=%s)", lab.guardPos, Coord{x,y})
				}
				lab.guardPos = Coord{x,y}
			default:
				// unexpected character
				return nil, errors.Errorf("unexpected character (pos=%s, char=%s)", Coord{x,y}, string(line[x]))
			}
		}
		lab.m = append(lab.m, row)
	}

	return &lab, nil
}

// GuardPatrolPositions returns a list of the distinct coordinates visited by the guard as he does his patrol.
// GuardPatrolPositions returns nil if the guard is caught in a loop along the way.
func (l *LabMap) GuardPatrolPositions() []Coord {
	pos := l.guardPos
	dir := l.guardDir

	rotateRight := func() {
		switch {
		case dir.X == 0 && dir.Y == -1:
			// north, rotate to east
			dir.X, dir.Y = 1, 0
		case dir.X == 1 && dir.Y == 0:
			// east, rotate south
			dir.X, dir.Y = 0, 1
		case dir.X == 0 && dir.Y == 1:
			// south, rotate west
			dir.X, dir.Y = -1, 0
		case dir.X == -1 && dir.Y == 0:
			// west, rotate north
			dir.X, dir.Y = 0, -1
		default:
			panic("unexpected direction")
		}
	}

	// iterate through the different coordinates visted by the guard
	visitedCoords := map[Coord][]Vector{}
	visitedCoords[pos] = append(visitedCoords[pos], dir)
	for {
		loc := l.Location(pos.X + dir.X, pos.Y + dir.Y)

		switch loc {
		case -1:
			// the guard has left the map, we're done

			// build list of visited coordinates
			coords := make([]Coord, 0, len(visitedCoords))
			for coord := range visitedCoords {
				coords = append(coords, coord)
			}

			return coords
		case 0:
			// the location is empty, move foward
			pos.X, pos.Y = pos.X + dir.X, pos.Y + dir.Y

			// check if the guard has entered a loop by having been at this location moving
			// the same direction before
			for _, oldDir := range visitedCoords[pos] {
				if dir == oldDir {
					// caught in loop!
					return nil
				}
			}

			visitedCoords[pos] = append(visitedCoords[pos], dir)
		case 1:
			// the location is occupied, rotate right
			rotateRight()
		}
	}
}

// NumberOfWaysToTrapGuard returns the number of ways the guard could be caught in loop by inserting a single obstacle
// into the map.
func (l *LabMap) NumberOfWaysToTrapGuard() int {
	count := 0

	potentialBlockingPoints := l.GuardPatrolPositions()

	for _, coord := range potentialBlockingPoints {
		if l.guardPos.X == coord.X && l.guardPos.Y == coord.Y {
			// skip the guard starting point
			continue
		}

		// insert obstacle
		l.m[coord.Y][coord.X] = true

		// test if guard gets caught in loop
		if steps := l.GuardPatrolPositions(); steps == nil {
			count++
		}

		l.m[coord.Y][coord.X] = false
	}

	return count
}

// Location returns what is at the location specified by x,y. It will return 0 if the location is empty,
// 1 if it contains an obstacle and -1 if the location is outside of the map.
func (l *LabMap) Location(x, y int) int {
	if y < 0 || y >= len(l.m) || x < 0 || x >= len(l.m[y]) {
		return -1
	}
	if l.m[y][x] {
		return 1
	}
	return 0
}

func (l *LabMap) Print() {
	for y := range l.m {
		for x := range l.m[y] {
			switch l.m[y][x] {
			case true:
				fmt.Print("#")
			case false:
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

type Vector struct {
	X, Y int
}

func (v Vector) String() string {
	return fmt.Sprintf("[%d,%d]", v.X, v.Y)
}

type Coord struct {
	X, Y int
}

func (c Coord) String() string {
	return fmt.Sprintf("(%d,%d)", c.X, c.Y)
}
