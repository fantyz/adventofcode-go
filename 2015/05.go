package main

import (
	"fmt"
	"strings"
)

func init() { days["5"] = Day5 }

/*
--- Day 5: Doesn't He Have Intern-Elves For This? ---
Santa needs help figuring out which strings in his text file are naughty or nice.

A nice string is one with all of the following properties:

It contains at least three vowels (aeiou only), like aei, xazegov, or aeiouaeiouaeiou.
It contains at least one letter that appears twice in a row, like xx, abcdde (dd), or aabbccdd (aa, bb, cc, or dd).
It does not contain the strings ab, cd, pq, or xy, even if they are part of one of the other requirements.
For example:

ugknbfddgicrmopn is nice because it has at least three vowels (u...i...o...), a double letter (...dd...), and none of the disallowed substrings.
aaa is nice because it has at least three vowels and a double letter, even though the letters used by different rules overlap.
jchzalrnumimnmhp is naughty because it has no double letter.
haegwjzuvuyypxyu is naughty because it contains the string xy.
dvszwmarrgswjxmb is naughty because it contains only one vowel.
How many strings are nice?

Your puzzle answer was 238.

--- Part Two ---
Realizing the error of his ways, Santa has switched to a better model of determining whether a string is naughty or nice. None of the old rules apply, as they are all clearly ridiculous.

Now, a nice string is one with all of the following properties:

It contains a pair of any two letters that appears at least twice in the string without overlapping, like xyxy (xy) or aabcdefgaa (aa), but not like aaa (aa, but it overlaps).
It contains at least one letter which repeats with exactly one letter between them, like xyx, abcdefeghi (efe), or even aaa.
For example:

qjhvhtzxzqqjkmpb is nice because is has a pair that appears twice (qj) and a letter that repeats with exactly one letter between them (zxz).
xxyxx is nice because it has a pair that appears twice and a letter that repeats with one between, even though the letters used by each rule overlap.
uurcxstgmygtbstg is naughty because it has a pair (tg) but no repeat with a single letter between them.
ieodomkazucvgmuy is naughty because it has a repeating letter with one between (odo), but no pair that appears twice.
How many strings are nice under these new rules?

Your puzzle answer was 69.
*/

func Day5() {
	fmt.Println("--- Day 5: Doesn't He Have Intern-Elves For This? ---")
	fmt.Println("   Nice strings:", NiceStrings(day5Input))
	fmt.Println("  Nicer strings:", NicerStrings(day5Input))
}

// NiceStrings takes multiple lines separated by newlines and counts how many of them are
// nice according to the part 1 rules.
func NiceStrings(in string) int {
	nice := 0
	for _, s := range strings.Split(in, "\n") {
		if IsNice(s) {
			nice++
		}
	}
	return nice
}

// IsNice checks if a string and returns true if it is nice according to the part 1 rules.
func IsNice(s string) bool {
	var lastC rune
	vowels := 0
	twoInARow := false
	for _, c := range s {
		// rule 1: count vowels
		if c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' {
			vowels++
		}
		// rule 2: look for two characters in a row
		if lastC == c {
			twoInARow = true
		}
		// rule 3: disallowed two characters in a row
		if lastC == 'a' && c == 'b' || lastC == 'c' && c == 'd' || lastC == 'p' && c == 'q' || lastC == 'x' && c == 'y' {
			return false

		}

		lastC = c
	}

	// need at least 3 vowels and two in a row
	if vowels < 3 || !twoInARow {
		return false
	}

	return true
}

// NicerStrings takes multiple lines separated by newlines and counts how many of them are
// nice according to the part 2 rules.
func NicerStrings(in string) int {
	nicer := 0
	for _, s := range strings.Split(in, "\n") {
		if IsNicer(s) {
			nicer++
		}
	}
	return nicer
}

// IsNicer takes a string and returns true if it is nice according to the part 2 rules.
func IsNicer(s string) bool {
	rule1 := false
	rule2 := false

	// iterate through the string and compare the current character against the string leading up to it
	// the shortest sequence is that we need to check for is 3 chars long
	for i := 2; i < len(s); i++ {
		// rule 1: look for the last two letters of the string in the initial part of the string
		if strings.Contains(s[:i-1], s[i-1:i+1]) {
			rule1 = true
		}
		// rule 2: look for repetition with one char in between
		if s[i-2] == s[i] {
			rule2 = true
		}
	}

	return rule1 && rule2
}
