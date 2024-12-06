package main

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

func init() { days["06"] = Day06 }

/*
 */

func Day06() {
	fmt.Println("--- Day 6: Guard Gallivant ---")
	lab, err := NewLabMap(day06Input)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Part 1: Locations visited by the guard:", lab.GuardPatrolPositions())
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

// GuardPatrolPositions returns the number of distinct coordinates that the guard visites on his patrol.
// GuardPatrolPositions returns -1 if the guard ends up going in a loop and never would leave the map.
func (l *LabMap) GuardPatrolPositions() int {
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
			return len(visitedCoords)
		case 0:
			// the location is empty, move foward
			pos.X, pos.Y = pos.X + dir.X, pos.Y + dir.Y

			// check if the guard has entered a loop by having been at this location moving
			// the same direction before
			for _, oldDir := range visitedCoords[pos] {
				if dir == oldDir {
					// caught in loop!
					return -1
				}
			}

			visitedCoords[pos] = append(visitedCoords[pos], dir)

		case 1:
			// the location is occupied, rotate right
			rotateRight()
		}
	}

	return len(visitedCoords)
}

// NumberOfWaysToTrapGuard returns the number of ways the guard could be caught in loop by inserting a single obstacle
// into the map.
func (l *LabMap) NumberOfWaysToTrapGuard() int {
	// brute forcing this- I cannot think of a better solution to just trying to insert an obstacle in each
	// possible place in the map and test it.
	count := 0

	for y := range l.m {
		for x := range l.m[y] {
			if l.Location(x, y) != 0 || (l.guardPos.X == x && l.guardPos.Y == y) {
				// non-empty location, skip it
				continue
			}

			// insert obstacle
			l.m[y][x] = true

			// test if guard gets caught in loop
			if steps := l.GuardPatrolPositions(); steps < 0 {
				count++
			}

			l.m[y][x] = false
		}
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
