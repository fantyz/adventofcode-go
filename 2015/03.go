package main

import (
	"fmt"
)

func init() { days["3"] = Day3 }

/*
--- Day 3: Perfectly Spherical Houses in a Vacuum ---
Santa is delivering presents to an infinite two-dimensional grid of houses.

He begins by delivering a present to the house at his starting location, and then an elf at the North Pole calls him via radio and tells him where to move next. Moves are always exactly one house to the north (^), south (v), east (>), or west (<). After each move, he delivers another present to the house at his new location.

However, the elf back at the north pole has had a little too much eggnog, and so his directions are a little off, and Santa ends up visiting some houses more than once. How many houses receive at least one present?

For example:

> delivers presents to 2 houses: one at the starting location, and one to the east.
^>v< delivers presents to 4 houses in a square, including twice to the house at his starting/ending location.
^v^v^v^v^v delivers a bunch of presents to some very lucky children at only 2 houses.
Your puzzle answer was 2081.

--- Part Two ---
The next year, to speed up the process, Santa creates a robot version of himself, Robo-Santa, to deliver presents with him.

Santa and Robo-Santa start at the same location (delivering two presents to the same starting house), then take turns moving based on instructions from the elf, who is eggnoggedly reading from the same script as the previous year.

This year, how many houses receive at least one present?

For example:

^v delivers presents to 3 houses, because Santa goes north, and then Robo-Santa goes south.
^>v< now delivers presents to 3 houses, and Santa and Robo-Santa end up back where they started.
^v^v^v^v^v now delivers presents to 11 houses, with Santa going one direction and Robo-Santa going the other.
Your puzzle answer was 2341.
*/

func Day3() {
	fmt.Println("Day 3: Perfectly Spherical Houses in a Vacuum")
	totalVisited, err := DistributePresentsToHouses(1, day3Input)
	if err != nil {
		panic(err)
	}
	fmt.Println("   Total visted houses first year:", totalVisited)
	totalVisited, err = DistributePresentsToHouses(2, day3Input)
	if err != nil {
		panic(err)
	}
	fmt.Println("  Total visted houses second year:", totalVisited)

}

// Coord is a 2D simple representation of a coordinate.
type Coord struct {
	X, Y int
}

// Houses will keep track of what houses has been visisted.
//
// A map seems like a great structure to keep track of visited coordinates. We need it to be able to
// quickly lookup and add whether a specific coordinate has been visited already. We do not need to
// keep track of anything other than the fact that a given coordinate has been visted thus no need for
// the actual values in the map.
//
// Alternatively we could use a 2d array. However, given the unknown size needed, we would need extra
// logic that would be able to grow it easily in any direction once we go out of bounds. And it would
// likely consume a significant amount of additional memory as many coordinates wouldn't be used.
type Houses map[Coord]struct{}

// Visit takes a coordinate and mark it as visited.
func (h Houses) Visit(c Coord) {
	h[c] = struct{}{}
}

// TotalVisisted returns the total number of coordinates visited.
func (h Houses) TotalVisisted() int {
	return len(h)
}

// DistributePresentsToHouses take a workforce number specifying how many people are delivering presents
// as well as a list of moves they take turns taking. It returns the number of houses visited total.
func DistributePresentsToHouses(workforce int, moves string) (int, error) {
	if workforce <= 0 {
		return 0, nil
	}

	// each person in the workforce has his own set of coordinates
	coords := make([]Coord, workforce)
	activeIdx := 0

	h := make(Houses)
	h.Visit(Coord{0, 0})

	for _, m := range moves {
		switch m {
		case '>':
			coords[activeIdx].X++
		case '<':
			coords[activeIdx].X--
		case '^':
			coords[activeIdx].Y++
		case 'v':
			coords[activeIdx].Y--
		default:
			return 0, fmt.Errorf("unknown move (move=%v)", m)
		}
		h.Visit(coords[activeIdx])

		// next person in the workforce turn to move
		activeIdx++
		if activeIdx >= len(coords) {
			activeIdx = 0
		}
	}

	return h.TotalVisisted(), nil
}
