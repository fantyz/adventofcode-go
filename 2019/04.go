package main

import (
	"fmt"
)

/*

--- Day 4: Secure Container ---
You arrive at the Venus fuel depot only to discover it's protected by a password. The Elves had written the password on a sticky note, but someone threw it out.

However, they do remember a few key facts about the password:

It is a six-digit number.
The value is within the range given in your puzzle input.
Two adjacent digits are the same (like 22 in 122345).
Going from left to right, the digits never decrease; they only ever increase or stay the same (like 111123 or 135679).
Other than the range rule, the following are true:

111111 meets these criteria (double 11, never decreases).
223450 does not meet these criteria (decreasing pair of digits 50).
123789 does not meet these criteria (no double).
How many different passwords within the range given in your puzzle input meet these criteria?

Your puzzle answer was 2779.

--- Part Two ---
An Elf just remembered one more important detail: the two adjacent matching digits are not part of a larger group of matching digits.

Given this additional criterion, but still ignoring the range rule, the following are now true:

112233 meets these criteria because the digits never decrease and all repeated digits are exactly two digits long.
123444 no longer meets the criteria (the repeated 44 is part of a larger group of 444).
111122 meets the criteria (even though 1 is repeated more than twice, it still contains a double 22).
How many different passwords within the range given in your puzzle input meet all of the criteria?

Your puzzle answer was 1972.

---

Note to self: Clearification on part two rule is "must contain at least one group of exactly 2 digits"

*/

func main() {
	fmt.Println("Day 4 - Secure Container")

	count := 0
	for i := 108457; i <= 562041; i++ {
		if IsValid(i) {
			count++
		}
	}

	fmt.Println("Valid passwords in range:", count)
}

func IsValid(in int) bool {
	digit := 0
	divisor := 100000

	group := 0
	double := false

	for divisor > 0 {
		in = in - digit*10*divisor
		last := digit
		digit = in / divisor
		divisor = divisor / 10

		if last == digit {
			group++
		} else {
			if group == 1 {
				double = true
			}
			group = 0
		}

		if last > digit {
			return false
		}
	}

	return double || group == 1
}

const puzzleInput = `108457-562041`
