package main

import (
	"fmt"
	"sort"
	"strings"
)

func init() { days["5"] = Day5 }

/*
--- Day 5: Binary Boarding ---
You board your plane only to discover a new problem: you dropped your boarding pass! You aren't sure which seat is yours, and all of the flight attendants are busy with the flood of people that suddenly made it through passport control.

You write a quick program to use your phone's camera to scan all of the nearby boarding passes (your puzzle input); perhaps you can find your seat through process of elimination.

Instead of zones or groups, this airline uses binary space partitioning to seat people. A seat might be specified like FBFBBFFRLR, where F means "front", B means "back", L means "left", and R means "right".

The first 7 characters will either be F or B; these specify exactly one of the 128 rows on the plane (numbered 0 through 127). Each letter tells you which half of a region the given seat is in. Start with the whole list of rows; the first letter indicates whether the seat is in the front (0 through 63) or the back (64 through 127). The next letter indicates which half of that region the seat is in, and so on until you're left with exactly one row.

For example, consider just the first seven characters of FBFBBFFRLR:

Start by considering the whole range, rows 0 through 127.
F means to take the lower half, keeping rows 0 through 63.
B means to take the upper half, keeping rows 32 through 63.
F means to take the lower half, keeping rows 32 through 47.
B means to take the upper half, keeping rows 40 through 47.
B keeps rows 44 through 47.
F keeps rows 44 through 45.
The final F keeps the lower of the two, row 44.
The last three characters will be either L or R; these specify exactly one of the 8 columns of seats on the plane (numbered 0 through 7). The same process as above proceeds again, this time with only three steps. L means to keep the lower half, while R means to keep the upper half.

For example, consider just the last 3 characters of FBFBBFFRLR:

Start by considering the whole range, columns 0 through 7.
R means to take the upper half, keeping columns 4 through 7.
L means to take the lower half, keeping columns 4 through 5.
The final R keeps the upper of the two, column 5.
So, decoding FBFBBFFRLR reveals that it is the seat at row 44, column 5.

Every seat also has a unique seat ID: multiply the row by 8, then add the column. In this example, the seat has ID 44 * 8 + 5 = 357.

Here are some other boarding passes:

BFFFBBFRRR: row 70, column 7, seat ID 567.
FFFBBBFRRR: row 14, column 7, seat ID 119.
BBFFBBFRLL: row 102, column 4, seat ID 820.
As a sanity check, look through your list of boarding passes. What is the highest seat ID on a boarding pass?

Your puzzle answer was 989.

--- Part Two ---
Ding! The "fasten seat belt" signs have turned on. Time to find your seat.

It's a completely full flight, so your seat should be the only missing boarding pass in your list. However, there's a catch: some of the seats at the very front and back of the plane don't exist on this aircraft, so they'll be missing from your list as well.

Your seat wasn't at the very front or back, though; the seats with IDs +1 and -1 from yours will be in your list.

What is the ID of your seat?

Your puzzle answer was 548.
*/

func Day5() {
	fmt.Println("--- Day 5: Binary Boarding ---")
	fmt.Println("Highest seat ID:", HighestSeatID(day5Input))
	fmt.Println("Empty seat ID:", FindEmptySeat(day5Input))
}

// FindEmptySeat takes a list of seats and returns the first empty seatID found where both seatID-1
// and seatID+1 is occupied. FindEmptySeat will return -1 if no empty seats are found or the input
// is invalid.
func FindEmptySeat(seats string) int {
	// given the way the seatIDs are specified (row*8+seat) a completely full plane would
	// consist of one long sequence of seat IDs without any gaps. Knowing this, we simply need
	// to find a gap in the full list of seat IDs.
	var seatIDs []int

	for _, seat := range strings.Split(seats, "\n") {
		x, y := SeatPos(seat)
		if x == -1 {
			// bad seat
			return -1
		}
		seatIDs = append(seatIDs, SeatID(x, y))
	}

	sort.Ints(seatIDs)
	// there must be an occupied seat in front and behind, so start at 1
	for i := 1; i < len(seatIDs); i++ {
		if seatIDs[i-1]+2 == seatIDs[i] {
			// found empty seat between seatIDs[i-1] and seatID[i]
			return seatIDs[i] - 1
		}
	}

	// no empty seats found
	return -1
}

// HighestSeatID takes a list of seats and returns the highest encoutered seat ID.
func HighestSeatID(seats string) int {
	highest := 0
	for _, seat := range strings.Split(seats, "\n") {
		col, row := SeatPos(seat)
		seatID := SeatID(col, row)
		if seatID > highest {
			highest = seatID
		}
	}
	return highest
}

// SeatID takes a column, row seating position and returns the corresponding seat ID.
func SeatID(col, row int) int {
	return row*8 + col
}

// SeatPos takes a boarding pass seat and returns the column and row it corresponds to.
// If the input is invalid -1, -1 is returned.
func SeatPos(seat string) (int, int) {
	if len(seat) != 10 {
		return -1, -1
	}

	// plane has seats in 128 rows and 8 columns
	rMin, rMax := 0, 127
	cMin, cMax := 0, 7

	for i := 0; i < 7; i++ {
		switch seat[i] {
		case 'F':
			rMax -= (rMax - rMin + 1) / 2
		case 'B':
			rMin += (rMax - rMin + 1) / 2
		default:
			// bad seat definition
			return -1, -1
		}
	}
	if rMin != rMax {
		panic("rMin should equal rMax!")
	}

	for i := 7; i < 10; i++ {
		switch seat[i] {
		case 'L':
			cMax -= (cMax - cMin + 1) / 2
		case 'R':
			cMin += (cMax - cMin + 1) / 2
		default:
			// bad seat definition
			return -1, -1
		}
	}
	if cMin != cMax {
		panic("cMin should equal cMax!")
	}

	return cMin, rMin
}
