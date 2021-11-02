package main

import (
	"fmt"
	"strings"
)

func init() { days["18"] = Day18 }

/*
--- Day 18: Like a GIF For Your Yard ---
After the million lights incident, the fire code has gotten stricter: now, at most ten thousand lights are allowed. You arrange them in a 100x100 grid.

Never one to let you down, Santa again mails you instructions on the ideal lighting configuration. With so few lights, he says, you'll have to resort to animation.

Start by setting your lights to the included initial configuration (your puzzle input). A # means "on", and a . means "off".

Then, animate your grid in steps, where each step decides the next configuration based on the current one. Each light's next state (either on or off) depends on its current state and the current states of the eight lights adjacent to it (including diagonals). Lights on the edge of the grid might have fewer than eight neighbors; the missing ones always count as "off".

For example, in a simplified 6x6 grid, the light marked A has the neighbors numbered 1 through 8, and the light marked B, which is on an edge, only has the neighbors marked 1 through 5:

1B5...
234...
......
..123.
..8A4.
..765.
The state a light should have next is based on its current state (on or off) plus the number of neighbors that are on:

A light which is on stays on when 2 or 3 neighbors are on, and turns off otherwise.
A light which is off turns on if exactly 3 neighbors are on, and stays off otherwise.
All of the lights update simultaneously; they all consider the same current state before moving to the next.

Here's a few steps from an example configuration of another 6x6 grid:

Initial state:
.#.#.#
...##.
#....#
..#...
#.#..#
####..

After 1 step:
..##..
..##.#
...##.
......
#.....
#.##..

After 2 steps:
..###.
......
..###.
......
.#....
.#....

After 3 steps:
...#..
......
...#..
..##..
......
......

After 4 steps:
......
......
..##..
..##..
......
......
After 4 steps, this example has four lights on.

In your grid of 100x100 lights, given your initial configuration, how many lights are on after 100 steps?

Your puzzle answer was 821.

--- Part Two ---
You flip the instructions over; Santa goes on to point out that this is all just an implementation of Conway's Game of Life. At least, it was, until you notice that something's wrong with the grid of lights you bought: four lights, one in each corner, are stuck on and can't be turned off. The example above will actually run like this:

Initial state:
##.#.#
...##.
#....#
..#...
#.#..#
####.#

After 1 step:
#.##.#
####.#
...##.
......
#...#.
#.####

After 2 steps:
#..#.#
#....#
.#.##.
...##.
.#..##
##.###

After 3 steps:
#...##
####.#
..##.#
......
##....
####.#

After 4 steps:
#.####
#....#
...#..
.##...
#.....
#.#..#

After 5 steps:
##.###
.##..#
.##...
.##...
#.#...
##...#
After 5 steps, this example now has 17 lights on.

In your grid of 100x100 lights, given your initial configuration, but with the four corners always in the on state, how many lights are on after 100 steps?

Your puzzle answer was 886.
*/

func Day18() {
	fmt.Println("--- Day 18: Like a GIF For Your Yard ---")
	lights := NewAnimatedLights(day18Input)
	lights.Animate(100, false)
	fmt.Println("Number of lights turned on after 100 steps:", lights.LightsOn())
	lights = NewAnimatedLights(day18Input)
	lights.Animate(100, true)
	fmt.Println("Number of lights turned on after 100 steps with the four corners stuck on:", lights.LightsOn())
}

// AnimatedLights is a 2D slice of bool each representing a light indexed y,x and with an inverted y axis.
type AnimatedLights [][]bool

// NewAnimatedLights takes the puzzle input and returns AnimatedLights.
func NewAnimatedLights(input string) AnimatedLights {
	var lights AnimatedLights
	for _, line := range strings.Split(input, "\n") {
		row := make([]bool, len(line))
		for i, c := range line {
			if c == '#' {
				row[i] = true
			}
		}
		lights = append(lights, row)
	}
	return lights
}

// Animate animates the display in accordance with the adjancy rules for the number of steps specified.
// If cornersStuck is set to true, the four corners will be turned on initially as well as after each
// step.
func (lights AnimatedLights) Animate(steps int, cornersStuck bool) {
	if cornersStuck {
		lights.TurnOnCorners()
	}

	for i := 0; i < steps; i++ {
		// copy the current state to allow both evaluating and changing the lights at the same time
		currentState := make(AnimatedLights, len(lights))
		for n := range lights {
			currentState[n] = make([]bool, len(lights[n]))
			copy(currentState[n], lights[n])
		}

		// evalulate each light
		for y := range lights {
			for x := range lights[y] {
				neighbors := currentState.Neighbors(x, y)

				if currentState[y][x] {
					// light is on
					if neighbors != 2 && neighbors != 3 {
						lights[y][x] = false
					}
				} else {
					// light is off
					if neighbors == 3 {
						lights[y][x] = true
					}
				}
			}
		}

		if cornersStuck {
			lights.TurnOnCorners()
		}
	}
}

// Neighbors returns the number of the 8 neighboring lights that are on.
func (lights AnimatedLights) Neighbors(x, y int) int {
	count := 0

	for j := y - 1; j <= y+1; j++ {
		if j < 0 || j >= len(lights) {
			// out of bounds, ignore
			continue
		}
		for i := x - 1; i <= x+1; i++ {
			if i < 0 || i >= len(lights[j]) || (i == x && j == y) {
				// out of bounds or the input coordinate itself, ignore
				continue
			}
			if lights[j][i] {
				count++
			}
		}
	}

	return count
}

// LightsOn returns the total number of lights that are turned on.
func (lights AnimatedLights) LightsOn() int {
	count := 0
	for y := range lights {
		for x := range lights[y] {
			if lights[y][x] {
				count++
			}
		}
	}
	return count
}

// TurnOnCorners turns on the four lights in the corners.
func (lights AnimatedLights) TurnOnCorners() {
	lights[0][0] = true
	lights[len(lights)-1][0] = true
	lights[len(lights)-1][len(lights[len(lights)-1])-1] = true
	lights[0][len(lights[0])-1] = true
}

// Print outputs the lights to stdout.
func (lights AnimatedLights) Print() {
	for y := range lights {
		for x := range lights[y] {
			if lights[y][x] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}
