package main

import (
	"fmt"
	"strings"
)

func init() { days["11"] = Day11 }

/*
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
