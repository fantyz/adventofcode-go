package main

import (
	"fmt"
)

func init() { days["15"] = Day15 }

func Day15() {
	fmt.Println("--- Day 15: Rambunctious Recitation ---")
	init, err := LoadInts(day15Input, ",")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("The 2020th number spoken:", MemoryGame(init, 2020))
	fmt.Println("The 30.000.000th number spoken:", MemoryGame(init, 30000000))
}

// MemoryGame plays the memory game using the initial numbers provided and keeps going until
// the turn specified and returns the number spoken.
func MemoryGame(init []int, turn int) int {
	numbers := map[int]int{}

	for i := 0; i < len(init)-1; i++ {
		numbers[init[i]] = i + 1
	}
	n := 0
	if len(init) > 0 {
		n = init[len(init)-1]
	}

	for i := len(init) + 1; i <= turn; i++ {
		// this turns number is n

		// last turn the number n was heard
		last := numbers[n]

		// for future rounds; last time the number n was heard was in the previous round
		numbers[n] = i - 1

		if last == 0 {
			// n was not heard before the previous round
			n = 0
			continue
		}

		// the number spoken is the previous round number - the last round it was heard prior to that
		n = i - 1 - last
	}

	// all done
	return n
}
