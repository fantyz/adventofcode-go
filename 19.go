package main

import (
	"fmt"
)

/*

--- Day 19: An Elephant Named Joseph ---

The Elves contact you over a highly secure emergency channel. Back at the North Pole, the Elves are busy misunderstanding White Elephant parties.

Each Elf brings a present. They all sit in a circle, numbered starting with position 1. Then, starting with the first Elf, they take turns stealing all the presents from the Elf to their left. An Elf with no presents is removed from the circle and does not take turns.

For example, with five Elves (numbered 1 to 5):

  1
5   2
 4 3
Elf 1 takes Elf 2's present.
Elf 2 has no presents and is skipped.
Elf 3 takes Elf 4's present.
Elf 4 has no presents and is also skipped.
Elf 5 takes Elf 1's two presents.
Neither Elf 1 nor Elf 2 have any presents, so both are skipped.
Elf 3 takes Elf 5's three presents.
So, with five Elves, the Elf that sits starting in position 3 gets all the presents.

With the number of Elves given in your puzzle input, which Elf gets all the presents?

Your puzzle answer was 1841611.

--- Part Two ---

Realizing the folly of their present-exchange rules, the Elves agree to instead steal presents from the Elf directly across the circle. If two Elves are across the circle, the one on the left (from the perspective of the stealer) is stolen from. The other rules remain unchanged: Elves with no presents are removed from the circle entirely, and the other elves move in slightly to keep the circle evenly spaced.

For example, with five Elves (again numbered 1 to 5):

The Elves sit in a circle; Elf 1 goes first:
  1
5   2
 4 3
Elves 3 and 4 are across the circle; Elf 3's present is stolen, being the one to the left. Elf 3 leaves the circle, and the rest of the Elves move in:
  1           1
5   2  -->  5   2
 4 -          4
Elf 2 steals from the Elf directly across the circle, Elf 5:
  1         1
-   2  -->     2
  4         4
Next is Elf 4 who, choosing between Elves 1 and 2, steals from Elf 1:
 -          2
    2  -->
 4          4
Finally, Elf 2 steals from Elf 4:
 2
    -->  2
 -
So, with five Elves, the Elf that sits starting in position 2 gets all the presents.

With the number of Elves given in your puzzle input, which Elf now gets all the presents?

*/

func main() {
	fmt.Println("Advent of Code 2016 - Day 19")
	stealFromLeft()
	stealFromAcross()
}

func stealFromAcross() {
	elves := make([]int, puzzleInput)
	for i := range elves {
		elves[i] = i + 1
	}

	i := 0
	for len(elves) > 1 {
		across := len(elves)/2 + i
		if across >= len(elves) {
			across = across - len(elves)
		}

		//fmt.Println("Elf", elves[i], "steals from", elves[across])

		if len(elves)%10000 == 0 {
			fmt.Println("elves left...", len(elves))
		}
		for j := across; j < len(elves)-1; j++ {
			elves[j] = elves[j+1]
		}
		elves = elves[:len(elves)-1]

		if across > i {
			i++
		}
		if i >= len(elves) {
			i = 0
		}
	}

	fmt.Println("Part 2: Elf number", elves[0], "sits back with all the presents")
}

func stealFromLeft() {
	var elves [puzzleInput]bool
	for i := range elves {
		elves[i] = true
	}

	steal := false
	lastIndex := -1
	i := 0
	for {
		if i >= len(elves) {
			i = 0
		}

		if !elves[i] {
			i++
			continue
		}
		if i == lastIndex {
			// done
			fmt.Println("Part 1: Elf number", i+1, "sits back with all the presents")
			break
		}
		lastIndex = i

		if steal {
			//fmt.Println("Elf", i+1)
			elves[i] = false
			steal = false
		} else {
			//fmt.Print("Elf ", i+1, " steals from... ")
			steal = true
		}

		i++
	}
}

const testpuzzleInput = 6
const puzzleInput = 3017957
