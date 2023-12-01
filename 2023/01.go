package main

import (
	"fmt"
	"strings"
)

func init() { days["1"] = Day1 }

/*
--- Day 1: Trebuchet?! ---
Something is wrong with global snow production, and you've been selected to take a look. The Elves have even given you a map; on it, they've used stars to mark the top fifty locations that are likely to be having problems.

You've been doing this long enough to know that to restore snow operations, you need to check all fifty stars by December 25th.

Collect stars by solving puzzles. Two puzzles will be made available on each day in the Advent calendar; the second puzzle is unlocked when you complete the first. Each puzzle grants one star. Good luck!

You try to ask why they can't just use a weather machine ("not powerful enough") and where they're even sending you ("the sky") and why your map looks mostly blank ("you sure ask a lot of questions") and hang on did you just say the sky ("of course, where do you think snow comes from") when you realize that the Elves are already loading you into a trebuchet ("please hold still, we need to strap you in").

As they're making the final adjustments, they discover that their calibration document (your puzzle input) has been amended by a very young Elf who was apparently just excited to show off her art skills. Consequently, the Elves are having trouble reading the values on the document.

The newly-improved calibration document consists of lines of text; each line originally contained a specific calibration value that the Elves now need to recover. On each line, the calibration value can be found by combining the first digit and the last digit (in that order) to form a single two-digit number.

For example:

1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet
In this example, the calibration values of these four lines are 12, 38, 15, and 77. Adding these together produces 142.

Consider your entire calibration document. What is the sum of all of the calibration values?

Your puzzle answer was 54081.

--- Part Two ---
Your calculation isn't quite right. It looks like some of the digits are actually spelled out with letters: one, two, three, four, five, six, seven, eight, and nine also count as valid "digits".

Equipped with this new information, you now need to find the real first and last digit on each line. For example:

two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen
In this example, the calibration values are 29, 83, 13, 24, 42, 14, and 76. Adding these together produces 281.

What is the sum of all of the calibration values?

Your puzzle answer was 54649.
*/

func Day1() {
	fmt.Println("--- Day 1: Trebuchet?! ---")
	fmt.Println("Part 1: The sum of calibration values is", SumCalibrationValues(day01Input, true))
	fmt.Println("Part 2: The real sum of calibration values is", SumCalibrationValues(day01Input, false))
}

func SumCalibrationValues(in string, digitsOnly bool) int {
	sum := 0

	for _, line := range strings.Split(in, "\n") {
		sum += NewCalibrationValue(line, digitsOnly)
	}

	return sum
}

func NewCalibrationValue(s string, digitsOnly bool) int {
	val := 0
	for i := 0; i < len(s); i++ {
		first := -1

		switch {
		case s[i] == '1' || !digitsOnly && strings.HasPrefix(s[i:], "one"):
			first = 1
		case s[i] == '2' || !digitsOnly && strings.HasPrefix(s[i:], "two"):
			first = 2
		case s[i] == '3' || !digitsOnly && strings.HasPrefix(s[i:], "three"):
			first = 3
		case s[i] == '4' || !digitsOnly && strings.HasPrefix(s[i:], "four"):
			first = 4
		case s[i] == '5' || !digitsOnly && strings.HasPrefix(s[i:], "five"):
			first = 5
		case s[i] == '6' || !digitsOnly && strings.HasPrefix(s[i:], "six"):
			first = 6
		case s[i] == '7' || !digitsOnly && strings.HasPrefix(s[i:], "seven"):
			first = 7
		case s[i] == '8' || !digitsOnly && strings.HasPrefix(s[i:], "eight"):
			first = 8
		case s[i] == '9' || !digitsOnly && strings.HasPrefix(s[i:], "nine"):
			first = 9
		}

		if first >= 0 {
			// found it
			val = first * 10
			break
		}
	}
	for i := len(s) - 1; i >= 0; i-- {
		last := -1

		switch {
		case s[i] == '1' || !digitsOnly && strings.HasSuffix(s[:i+1], "one"):
			last = 1
		case s[i] == '2' || !digitsOnly && strings.HasSuffix(s[:i+1], "two"):
			last = 2
		case s[i] == '3' || !digitsOnly && strings.HasSuffix(s[:i+1], "three"):
			last = 3
		case s[i] == '4' || !digitsOnly && strings.HasSuffix(s[:i+1], "four"):
			last = 4
		case s[i] == '5' || !digitsOnly && strings.HasSuffix(s[:i+1], "five"):
			last = 5
		case s[i] == '6' || !digitsOnly && strings.HasSuffix(s[:i+1], "six"):
			last = 6
		case s[i] == '7' || !digitsOnly && strings.HasSuffix(s[:i+1], "seven"):
			last = 7
		case s[i] == '8' || !digitsOnly && strings.HasSuffix(s[:i+1], "eight"):
			last = 8
		case s[i] == '9' || !digitsOnly && strings.HasSuffix(s[:i+1], "nine"):
			last = 9
		}

		if last >= 0 {
			// found it
			val += last
			break
		}
	}
	return val
}
