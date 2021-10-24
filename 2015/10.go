package main

import (
	"fmt"
	"strconv"
)

func init() { days["10"] = Day10 }

/*
--- Day 10: Elves Look, Elves Say ---
Today, the Elves are playing a game called look-and-say. They take turns making sequences by reading aloud the previous sequence and using that reading as the next sequence. For example, 211 is read as "one two, two ones", which becomes 1221 (1 2, 2 1s).

Look-and-say sequences are generated iteratively, using the previous value as input for the next step. For each step, take the previous value, and replace each run of digits (like 111) with the number of digits (3) followed by the digit itself (1).

For example:

1 becomes 11 (1 copy of digit 1).
11 becomes 21 (2 copies of digit 1).
21 becomes 1211 (one 2 followed by one 1).
1211 becomes 111221 (one 1, one 2, and two 1s).
111221 becomes 312211 (three 1s, two 2s, and one 1).
Starting with the digits in your puzzle input, apply this process 40 times. What is the length of the result?

Your puzzle answer was 329356.

--- Part Two ---
Neat, right? You might also enjoy hearing John Conway talking about this sequence (that's Conway of Conway's Game of Life fame).

Now, starting again with the digits in your puzzle input, apply this process 50 times. What is the length of the new result?

Your puzzle answer was 4666278.
*/

func Day10() {
	fmt.Println("--- Day 10: Elves Look, Elves Say ---")
	fmt.Println("40 rounds of look-and-say produce a result with the length:", len(LookAndSay(day10Input, 40)))
	fmt.Println("50 rounds of look-and-say produce a result with the length:", len(LookAndSay(day10Input, 50)))
}

// LookAndSay takes a string and use it for input to play n rounds of look-and-say before
// outputting the result.
func LookAndSay(in string, n int) string {
	res := in
	for i := 0; i < n; i++ {
		res = LookAndSayRound(res)
	}
	return res
}

// LookAndSayRound takes a string and plays a round of look-and-say to outputs the result.
func LookAndSayRound(in string) string {
	if len(in) <= 0 {
		return ""
	}

	out := make([]byte, 0, len(in))

	// initialize v to the first character of the input to bootstrap the process
	i, v, count := 0, in[0], 0
	for {
		if i == len(in) || v != in[i] {
			// current sequence finsihed, create output
			out = append(out, []byte(strconv.Itoa(count))...)
			out = append(out, v)

			if i == len(in) {
				// game finished
				return string(out)
			}

			// prepare next sequence
			v = in[i]
			count = 0
		}
		i++
		count++
	}
}
