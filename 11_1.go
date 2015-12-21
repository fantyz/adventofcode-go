package main

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

Your puzzle input is hxbxwxba.

*/

import (
	"fmt"
)

var input = []byte(`hxbxwxba`)

func main() {
	for {
		// increment password
		pos := len(input) - 1
		for input[pos]+1 > 'z' {
			input[pos] = 'a'
			pos--
			if pos >= 0 {
				continue
			}

			input = append([]byte{'a'-1}, input...)
			pos = 0
		}
		input[pos] = input[pos] + 1

		// validate
		badPass := false
		seqOk := false
		lastPair := -1
		pairs := 0
		
		prev1 := byte('z')+10
		prev2 := byte('z')+20
		for i:=0; i < len(input); i++ {
			if input[i] == 'i' || input[i] == 'l' || input[i] == 'o' {
				badPass = true
				break
			}

			if prev1 == input[i] && lastPair != i-1 {
				pairs++
				lastPair = i
			}
			
			if prev1 == prev2+1 && input[i] == prev1 + 1 {
				seqOk = true
			}
			prev2, prev1 = prev1, input[i]
		}
		if seqOk && !badPass && pairs >= 2 {
			break
		}
		
	}
	fmt.Println(string(input))
}
