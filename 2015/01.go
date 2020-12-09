package main

import "fmt"

func init() { days["1"] = Day1 }

/*
--- Day 1: Not Quite Lisp ---
Santa was hoping for a white Christmas, but his weather machine's "snow" function is powered by stars, and he's fresh out! To save Christmas, he needs you to collect fifty stars by December 25th.

Collect stars by helping Santa solve puzzles. Two puzzles will be made available on each day in the Advent calendar; the second puzzle is unlocked when you complete the first. Each puzzle grants one star. Good luck!

Here's an easy puzzle to warm you up.

Santa is trying to deliver presents in a large apartment building, but he can't find the right floor - the directions he got are a little confusing. He starts on the ground floor (floor 0) and then follows the instructions one character at a time.

An opening parenthesis, (, means he should go up one floor, and a closing parenthesis, ), means he should go down one floor.

The apartment building is very tall, and the basement is very deep; he will never find the top or bottom floors.

For example:

(()) and ()() both result in floor 0.
((( and (()(()( both result in floor 3.
))((((( also results in floor 3.
()) and ))( both result in floor -1 (the first basement level).
))) and )())()) both result in floor -3.
To what floor do the instructions take Santa?

Your puzzle answer was 280.

--- Part Two ---
Now, given the same instructions, find the position of the first character that causes him to enter the basement (floor -1). The first character in the instructions has position 1, the second character has position 2, and so on.

For example:

) causes him to enter the basement at character position 1.
()()) causes him to enter the basement at character position 5.
What is the position of the character that causes Santa to first enter the basement?

Your puzzle answer was 1797.
*/

func Day1() {
	fmt.Println("--- Day 1: Not Quite Lisp ---")
	finalFloor, firstBasementIdx := DeliverPresents(day1Input)
	fmt.Println("          Final floor after delivering presents:", finalFloor)
	fmt.Println("  First time entering basement happens at index:", firstBasementIdx)
}

// DeliverPresents will navigate the infinite floors of the appartment building and return
// the final floor number along with the index (1-indexed) of the first instruction that
// takes santa into the basement. The index return will be -1 if santa never goes into the
// basement.
func DeliverPresents(instructions string) (int, int) {
	floor := 0
	firstInBasement := -1
	for n, inst := range instructions {
		switch inst {
		case ')':
			floor--
		case '(':
			floor++
		default:
			panic("Unknown instruction: " + string(inst))
		}
		if firstInBasement == -1 && floor == -1 {
			firstInBasement = n + 1
		}
	}
	return floor, firstInBasement
}
