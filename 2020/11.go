package main

import (
	"fmt"
	"strings"
)

func init() { days["11"] = Day11 }

/*
--- Day 11: Seating System ---
Your plane lands with plenty of time to spare. The final leg of your journey is a ferry that goes directly to the tropical island where you can finally start your vacation. As you reach the waiting area to board the ferry, you realize you're so early, nobody else has even arrived yet!

By modeling the process people use to choose (or abandon) their seat in the waiting area, you're pretty sure you can predict the best place to sit. You make a quick map of the seat layout (your puzzle input).

The seat layout fits neatly on a grid. Each position is either floor (.), an empty seat (L), or an occupied seat (#). For example, the initial seat layout might look like this:

L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL
Now, you just need to model the people who will be arriving shortly. Fortunately, people are entirely predictable and always follow a simple set of rules. All decisions are based on the number of occupied seats adjacent to a given seat (one of the eight positions immediately up, down, left, right, or diagonal from the seat). The following rules are applied to every seat simultaneously:

If a seat is empty (L) and there are no occupied seats adjacent to it, the seat becomes occupied.
If a seat is occupied (#) and four or more seats adjacent to it are also occupied, the seat becomes empty.
Otherwise, the seat's state does not change.
Floor (.) never changes; seats don't move, and nobody sits on the floor.

After one round of these rules, every seat in the example layout becomes occupied:

#.##.##.##
#######.##
#.#.#..#..
####.##.##
#.##.##.##
#.#####.##
..#.#.....
##########
#.######.#
#.#####.##
After a second round, the seats with four or more occupied adjacent seats become empty again:

#.LL.L#.##
#LLLLLL.L#
L.L.L..L..
#LLL.LL.L#
#.LL.LL.LL
#.LLLL#.##
..L.L.....
#LLLLLLLL#
#.LLLLLL.L
#.#LLLL.##
This process continues for three more rounds:

#.##.L#.##
#L###LL.L#
L.#.#..#..
#L##.##.L#
#.##.LL.LL
#.###L#.##
..#.#.....
#L######L#
#.LL###L.L
#.#L###.##
#.#L.L#.##
#LLL#LL.L#
L.L.L..#..
#LLL.##.L#
#.LL.LL.LL
#.LL#L#.##
..L.L.....
#L#LLLL#L#
#.LLLLLL.L
#.#L#L#.##
#.#L.L#.##
#LLL#LL.L#
L.#.L..#..
#L##.##.L#
#.#L.LL.LL
#.#L#L#.##
..L.L.....
#L#L##L#L#
#.LLLLLL.L
#.#L#L#.##
At this point, something interesting happens: the chaos stabilizes and further applications of these rules cause no seats to change state! Once people stop moving around, you count 37 occupied seats.

Simulate your seating area by applying the seating rules repeatedly until no seats change state. How many seats end up occupied?

Your puzzle answer was 2406.

--- Part Two ---
As soon as people start to arrive, you realize your mistake. People don't just care about adjacent seats - they care about the first seat they can see in each of those eight directions!

Now, instead of considering just the eight immediately adjacent seats, consider the first seat in each of those eight directions. For example, the empty seat below would see eight occupied seats:

.......#.
...#.....
.#.......
.........
..#L....#
....#....
.........
#........
...#.....
The leftmost empty seat below would only see one empty seat, but cannot see any of the occupied ones:

.............
.L.L.#.#.#.#.
.............
The empty seat below would see no occupied seats:

.##.##.
#.#.#.#
##...##
...L...
##...##
#.#.#.#
.##.##.
Also, people seem to be more tolerant than you expected: it now takes five or more visible occupied seats for an occupied seat to become empty (rather than four or more from the previous rules). The other rules still apply: empty seats that see no occupied seats become occupied, seats matching no rule don't change, and floor never changes.

Given the same starting layout as above, these new rules cause the seating area to shift around as follows:

L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL
#.##.##.##
#######.##
#.#.#..#..
####.##.##
#.##.##.##
#.#####.##
..#.#.....
##########
#.######.#
#.#####.##
#.LL.LL.L#
#LLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLL#
#.LLLLLL.L
#.LLLLL.L#
#.L#.##.L#
#L#####.LL
L.#.#..#..
##L#.##.##
#.##.#L.##
#.#####.#L
..#.#.....
LLL####LL#
#.L#####.L
#.L####.L#
#.L#.L#.L#
#LLLLLL.LL
L.L.L..#..
##LL.LL.L#
L.LL.LL.L#
#.LLLLL.LL
..L.L.....
LLLLLLLLL#
#.LLLLL#.L
#.L#LL#.L#
#.L#.L#.L#
#LLLLLL.LL
L.L.L..#..
##L#.#L.L#
L.L#.#L.L#
#.L####.LL
..#.#.....
LLL###LLL#
#.LLLLL#.L
#.L#LL#.L#
#.L#.L#.L#
#LLLLLL.LL
L.L.L..#..
##L#.#L.L#
L.L#.LL.L#
#.LLLL#.LL
..#.L.....
LLL###LLL#
#.LLLLL#.L
#.L#LL#.L#
Again, at this point, people stop shifting around and the seating area reaches equilibrium. Once this occurs, you count 26 occupied seats.

Given the new visibility method and the rule change for occupied seats becoming empty, once equilibrium is reached, how many seats end up occupied?

Your puzzle answer was 2149.
*/

func Day11() {
	fmt.Println("--- Day 11: Seating System ---")
	fmt.Println("  Seats occupied after stabilizing (view distance=1, min adjacent=4):", NewSeatLayout(day11Input).StabilizeSeats(1, 4))
	fmt.Println("  Seats occupied after stabilizing (view distance=-1, min adjacent=5):", NewSeatLayout(day11Input).StabilizeSeats(-1, 5))
}

type SeatLayout [][]byte

func NewSeatLayout(in string) SeatLayout {
	var l SeatLayout
	for _, row := range strings.Split(in, "\n") {
		l = append(l, []byte(row))
	}
	return l
}

// ChangeSeats will apply the seat changing rules once using the viewing distance and minimum
// adjacent seated to cause a seating change provided and return the number of seats that changed
// state and the total number of occupied seats.
func (l SeatLayout) ChangeSeats(dist, minAdj int) (int, int) {
	// copy the initial layout to evaluate rules on why the existing layout is mutated
	oldL := make(SeatLayout, len(l))
	for y := range l {
		oldL[y] = make([]byte, len(l[y]))
		copy(oldL[y], l[y])
	}

	changes, total := 0, 0
	for y := range oldL {
		for x := range oldL[y] {
			switch oldL[y][x] {
			case 'L':
				if oldL.OccupiedAdjacentSeats(dist, x, y) == 0 {
					l[y][x] = '#'
					changes++
					total++
				}
			case '#':
				if oldL.OccupiedAdjacentSeats(dist, x, y) >= minAdj {
					l[y][x] = 'L'
					changes++
				} else {
					total++
				}
			}
		}
	}
	return changes, total
}

// StabilizeSeats will keep changing seats using the viewing distance and minimum adjacent seated
// to cause a seating chang eprovided until no seats changes state and returns the total number of
// occupied seats.
func (l SeatLayout) StabilizeSeats(dist, minAdj int) int {
	for {
		n, total := l.ChangeSeats(dist, minAdj)
		if n <= 0 {
			return total
		}
	}
}

// OccupiedAdjacentSeats takes a position and returns the number of occupied adjacent seats
// within the specified maximum distance next to it (among the 8 possible). Use -1 as maxDist
// to indicate no maximum distance.
func (l SeatLayout) OccupiedAdjacentSeats(maxDist, xPos, yPos int) int {
	count := 0

	// keep track of whether the view is blocked in the 8 possible directions
	viewBlocked := [8]bool{}

	for d := 1; ; d++ {
		if maxDist > 0 && d > maxDist {
			// maximum distance reached, all done
			break
		}
		direction := -1

		for x := xPos - d; x <= xPos+d; x += d {
			for y := yPos - d; y <= yPos+d; y += d {
				if x == xPos && y == yPos {
					// ignore the position being evaluated (center)
					continue
				}

				direction++
				if viewBlocked[direction] {
					// nothing further to see in this direction
					continue
				}

				if y < 0 || y >= len(l) || x < 0 || x >= len(l[y]) {
					// out of bounds, view is blocked by definition
					viewBlocked[direction] = true
					continue
				}

				switch l[y][x] {
				case '#':
					count++
					viewBlocked[direction] = true
				case 'L':
					viewBlocked[direction] = true
				}

			}
		}

		// check if all views are blocked
		allViewsBlocked := true
		for _, viewBlocked := range viewBlocked {
			if !viewBlocked {
				allViewsBlocked = false
				break
			}
		}
		if allViewsBlocked {
			// all views are blocked, all done
			break
		}
	}
	return count
}
