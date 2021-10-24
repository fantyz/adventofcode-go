package main

import (
	"fmt"
)

func init() { days["11"] = Day11 }

/*
--- Day 11: Corporate Policy ---
Santa's previous password expired, and he needs help choosing a new one.

To help him remember his new password after the old one expires, Santa has devised a method of coming up with a password based on the previous one. Corporate policy dictates that passwords must be exactly eight lowercase letters (for security reasons), so he finds his new password by incrementing his old password string repeatedly until it is valid.

Incrementing is just like counting with numbers: xx, xy, xz, ya, yb, and so on. Increase the rightmost letter one step; if it was z, it wraps around to a, and repeat with the next letter to the left until one doesn't wrap around.

Unfortunately for Santa, a new Security-Elf recently started, and he has imposed some additional password requirements:

Passwords must include one increasing straight of at least three letters, like abc, bcd, cde, and so on, up to xyz. They cannot skip letters; abd doesn't count.
Passwords may not contain the letters i, o, or l, as these letters can be mistaken for other characters and are therefore confusing.
Passwords must contain at least two different, non-overlapping pairs of letters, like aa, bb, or zz.
For example:

hijklmmn meets the first requirement (because it contains the straight hij) but fails the second requirement requirement (because it contains i and l).
abbceffg meets the third requirement (because it repeats bb and ff) but fails the first requirement.
abbcegjk fails the third requirement, because it only has one double letter (bb).
The next password after abcdefgh is abcdffaa.
The next password after ghijklmn is ghjaabcc, because you eventually skip all the passwords that start with ghi..., since i is not allowed.
Given Santa's current password (your puzzle input), what should his next password be?

Your puzzle answer was cqjxxyzz.

--- Part Two ---
Santa's password expired again. What's the next one?

Your puzzle answer was cqkaabcc.
*/

func Day11() {
	fmt.Println("--- Day 11: Corporate Policy ---")
	newPass := GeneratePassword(day11Input, WithStraightOfThreePolicy(), WithNoIOLPolicy(), WithTwoNonOverlappingPairs())
	fmt.Println("New password should be:", newPass)
	newPass2 := GeneratePassword(newPass, WithStraightOfThreePolicy(), WithNoIOLPolicy(), WithTwoNonOverlappingPairs())
	fmt.Println("New password after that should be:", newPass2)
}

// GeneratePassword takes an old password and "increments" it until reaching a password
// that all provided policies can validate successfully.
// If the old password provided contains any characters not a-z this character will either
// be incremented to "aa".
func GeneratePassword(oldPwd string, policies ...Policy) string {
	isValid := false

	p := []byte(oldPwd)
	for !isValid {
		// increment existing password starting with the last character
		i := len(p) - 1
		for {
			p[i]++
			if p[i] >= 'a' && p[i] <= 'z' {
				// password successfully incremented
				break
			}
			// wrap character around to 'a' and increment the next character
			p[i] = 'a'
			i--
			if i < 0 {
				// increase the length of the password by adding an 'a' to the beginning of it
				p = append([]byte{'a'}, p...)
				break
			}
		}

		// validate incremented password against policies
		isValid = true
		for _, validator := range policies {
			isValid = validator(p)
			if !isValid {
				break
			}
		}
	}

	return string(p)
}

// Policy is a function that takes a password and returns whether it pass the policy or not.
type Policy func([]byte) bool

// WithStraightOfThreePolicy returns a policy that require three subsequent letters to be
// a straight (eg. abc, bcd, etc).
func WithStraightOfThreePolicy() Policy {
	return func(p []byte) bool {
		for i := 0; i <= len(p)-3; i++ {
			if p[i] == p[i+1]-1 && p[i] == p[i+2]-2 {
				return true
			}
		}
		return false
	}
}

// WithNoIOLPolicy returns a Policy that rejects any password containing the letters i, o or l.
func WithNoIOLPolicy() Policy {
	return func(p []byte) bool {
		for _, c := range p {
			switch c {
			case 'i', 'o', 'l':
				return false
			}
		}
		return true
	}
}

// WithTwoNonOverlappingPairs returns a Policy that require any password to have at least two
// different non-overlapping pairs within them (eg. "aa", "bb", etc)
func WithTwoNonOverlappingPairs() Policy {
	return func(p []byte) bool {
		pairs := 0
		for i := 0; i <= len(p)-2; i++ {
			if p[i] == p[i+1] {
				pairs++
				if pairs >= 2 {
					return true
				}
				// skip a character to avoid overlap
				i++
			}
		}
		return false
	}
}
