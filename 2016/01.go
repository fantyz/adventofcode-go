package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*

You're airdropped near Easter Bunny Headquarters in a city somewhere. "Near", unfortunately, is as close as you can get - the instructions on the Easter Bunny Recruiting Document the Elves intercepted start here, and nobody had time to work them out further.

The Document indicates that you should start at the given coordinates (where you just landed) and face North. Then, follow the provided sequence: either turn left (L) or right (R) 90 degrees, then walk forward the given number of blocks, ending at a new intersection.

There's no time to follow such ridiculous instructions on foot, though, so you take a moment and work out the destination. Given that you can only walk on the street grid of the city, how far is the shortest path to the destination?

For example:

Following R2, L3 leaves you 2 blocks East and 3 blocks North, or 5 blocks away.
R2, R2, R2 leaves you 2 blocks due South of your starting position, which is 2 blocks away.
R5, L5, R5, R3 leaves you 12 blocks away.
How many blocks away is Easter Bunny HQ?

--- Part Two ---

Then, you notice the instructions continue on the back of the Recruiting Document. Easter Bunny HQ is actually at the first location you visit twice.

For example, if your instructions are R8, R4, R4, R8, the first location you visit twice is 4 blocks away, due East.

How many blocks away is the first location you visit twice?

*/

const puzzleInput = `R5, R4, R2, L3, R1, R1, L4, L5, R3, L1, L1, R4, L2, R1, R4, R4, L2, L2, R4, L4, R1, R3, L3, L1, L2, R1, R5, L5, L1, L1, R3, R5, L1, R4, L5, R5, R1, L185, R4, L1, R51, R3, L2, R78, R1, L4, R188, R1, L5, R5, R2, R3, L5, R3, R4, L1, R2, R2, L4, L4, L5, R5, R4, L4, R2, L5, R2, L1, L4, R4, L4, R2, L3, L4, R2, L3, R3, R2, L2, L3, R4, R3, R1, L4, L2, L5, R4, R4, L1, R1, L5, L1, R3, R1, L2, R1, R1, R3, L4, L1, L3, R2, R4, R2, L2, R1, L5, R3, L3, R3, L1, R4, L3, L3, R4, L2, L1, L3, R2, R3, L2, L1, R4, L3, L5, L2, L4, R1, L4, L4, R3, R5, L4, L1, L1, R4, L2, R5, R1, R1, R2, R1, R5, L1, L3, L5, R2`

type Direction uint8

const (
	North Direction = iota
	East
	South
	West
)

func main() {
	fmt.Println("Advent of Code 2016 - Day 01")

	moves := strings.Split(puzzleInput, ", ")

	pos := NewPos()
	for _, move := range moves {
		dist, err := strconv.Atoi(move[1:])
		if err != nil {
			panic(err)
		}

		var direction Direction
		switch pos.Direction() {
		case North:
			if move[:1] == "R" {
				direction = East
			} else {
				direction = West
			}
		case East:
			if move[:1] == "R" {
				direction = South
			} else {
				direction = North
			}
		case South:
			if move[:1] == "R" {
				direction = West
			} else {
				direction = East
			}
		case West:
			if move[:1] == "R" {
				direction = North
			} else {
				direction = South
			}
		}

		pos.Move(direction, dist)
	}

	fmt.Printf("Bunny HQ Coordinates: %d, %d (distance=%d)\n", pos.x, pos.y, abs(pos.x)+abs(pos.y))
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func NewPos() *Pos {
	p := &Pos{
		visitedCoords: make(map[string]struct{}),
	}
	p.visit()
	return p
}

type Pos struct {
	x, y          int
	d             Direction
	visitedCoords map[string]struct{}
}

func (p *Pos) Direction() Direction {
	return p.d
}

func (p *Pos) Move(direction Direction, distance int) {
	p.d = direction
	switch direction {
	case North:
		p.moveNorth(distance)
	case South:
		p.moveSouth(distance)
	case East:
		p.moveEast(distance)
	case West:
		p.moveWest(distance)
	}
}

func (p *Pos) moveEast(n int) {
	for i := 0; i < n; i++ {
		p.x--
		p.visit()
	}
}

func (p *Pos) moveWest(n int) {
	for i := 0; i < n; i++ {
		p.x++
		p.visit()
	}
}

func (p *Pos) moveNorth(n int) {
	for i := 0; i < n; i++ {
		p.y++
		p.visit()
	}
}
func (p *Pos) moveSouth(n int) {
	for i := 0; i < n; i++ {
		p.y--
		p.visit()
	}
}

func (p *Pos) visit() {
	coord := fmt.Sprintf("%d,%d", p.x, p.y)
	if _, found := p.visitedCoords[coord]; found {
		fmt.Printf("We've already been here: %d, %d (distance=%d)\n", p.x, p.y, abs(p.x)+abs(p.y))
	}
	p.visitedCoords[coord] = struct{}{}
}
