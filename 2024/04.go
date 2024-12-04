package main

import (
	"fmt"
	"strings"
)

func init() { days["04"] = Day04 }

/*
--- Day 4: Ceres Search ---
"Looks like the Chief's not here. Next!" One of The Historians pulls out a device and pushes the only button on it. After a brief flash, you recognize the interior of the Ceres monitoring station!

As the search for the Chief continues, a small Elf who lives on the station tugs on your shirt; she'd like to know if you could help her with her word search (your puzzle input). She only has to find one word: XMAS.

This word search allows words to be horizontal, vertical, diagonal, written backwards, or even overlapping other words. It's a little unusual, though, as you don't merely need to find one instance of XMAS - you need to find all of them. Here are a few ways XMAS might appear, where irrelevant characters have been replaced with .:


..X...
.SAMX.
.A..A.
XMAS.S
.X....
The actual word search will be full of letters instead. For example:

MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX
In this word search, XMAS occurs a total of 18 times; here's the same word search again, but where letters not involved in any XMAS have been replaced with .:

....XXMAS.
.SAMXMS...
...S..A...
..A.A.MS.X
XMASAMX.MM
X.....XA.A
S.S.S.S.SS
.A.A.A.A.A
..M.M.M.MM
.X.X.XMASX
Take a look at the little Elf's word search. How many times does XMAS appear?

Your puzzle answer was 2507.

--- Part Two ---
The Elf looks quizzically at you. Did you misunderstand the assignment?

Looking for the instructions, you flip over the word search to find that this isn't actually an XMAS puzzle; it's an X-MAS puzzle in which you're supposed to find two MAS in the shape of an X. One way to achieve that is like this:

M.S
.A.
M.S
Irrelevant characters have again been replaced with . in the above diagram. Within the X, each MAS can be written forwards or backwards.

Here's the same example from before, but this time all of the X-MASes have been kept instead:

.M.S......
..A..MSMS.
.M.S.MAA..
..A.ASMSM.
.M.S.M....
..........
S.S.S.S.S.
.A.A.A.A..
M.M.M.M.M.
..........
In this example, an X-MAS appears 9 times.

Flip the word search from the instructions back over to the word search side and try again. How many times does an X-MAS appear?

Your puzzle answer was 1969.
*/

func Day04() {
	fmt.Println("--- Day 4: Ceres Search ---")
	fmt.Println("Part 1: Number of occurrences of XMAS:", WordSearch(day04Input, "XMAS"))
	fmt.Println("Part 2: Number of occurrences of X-MAS:", XMASSearch(day04Input))
}

// XMASSearch takes in input string in and returns the number of times the X-MAS appears in it.
// See puzzle description for details on X-MAS shape.
func XMASSearch(in string) int {
	if len(in) <= 0 {
		return 0
	}

	s := strings.Split(in, "\n")

	// X-MAS can be rotated into these four patterns
	//
	// M M  S M  S S  M S
	//  A    A    A    A
	// S S  S M  M M  M S
	//
	// Lets look for the top left character and see if the rest match up
	count := 0

	for y:=0; y<len(s)-2; y++ {
		for x:=0;x<len(s[y])-2; x++ {
			if s[y][x] == 'M' || s[y][x] == 'S' {
				pattern := fmt.Sprintf("%c%c%c%c%c", s[y][x], s[y][x+2], s[y+1][x+1], s[y+2][x], s[y+2][x+2])
				if pattern == "MMASS" {
					count++
				}
				if pattern == "SMASM" {
					count++
				}
				if pattern == "SSAMM" {
					count++
				}
				if pattern == "MSAMS" {
					count++
				}
			}
		}
	}

	return count
}

// WordSearch takes an input string in and returns the number times the word appears in it whether
// it is horizontally, vertically or diagonally - all in any direction.
func WordSearch(in string, word string) int {
	if len(word) <= 0 {
		return 0
	}

	s := strings.Split(in, "\n")

	// iterate through the input text to look for the first character of the word
	// if we find one check if the full word is represented in any of the directions
	count := 0
	for y := range s {
		for x := range s[y] {
			if s[y][x] == word[0] {
				// might be a word here in one of the directions
				if checkWord(s, word, x, y, 1, 0) { // horizontal
					count++
				}
				if checkWord(s, word, x, y, -1, 0) { // horizontal reversed
					count++
				}
				if checkWord(s, word, x, y, 0, 1) { // vertical
					count++
				}
				if checkWord(s, word, x, y, 0, -1) { // vertical reversed
					count++
				}
				if checkWord(s, word, x, y, 1, 1) { // bottom right diagonal
					count++
				}
				if checkWord(s, word, x, y, -1, 1) { // bottom left diagonal
					count++
				}
				if checkWord(s, word, x, y, -1, -1) { // top left diagonal
					count++
				}
				if checkWord(s, word, x, y, 1, -1) { // top right diagonal
					count++
				}
			}
		}
	}

	return count
}

// checkWord takes the input text as a slice of strings each representing a line, the word we
// are searching for a coordinate and x- and y coordinate modifiers specifying a direction to
// check in and checks if the word is present.
func checkWord(s []string, word string, x, y, xMod, yMod int) bool {
	if y < 0 || y >= len(s) || x < 0 || x >= len(s[y]) {
		// starting cordinate outside of s
		return false
	}

	xEnd := x + xMod * (len(word) - 1)
	yEnd := y + yMod * (len(word) - 1)
	if yEnd < 0 || yEnd >= len(s) || xEnd < 0 || xEnd >= len(s[y]) {
		// ending  cordinate outside of s
		return false
	}

	for i := range word {
		if word[i] != s[y+yMod*i][x+xMod*i] {
			// characther did not match the word
			return false
		}
	}

	return true
}
