package main

import (
	"crypto/md5"
	"fmt"
)

func init() { days["4"] = Day4 }

/*
--- Day 4: The Ideal Stocking Stuffer ---
Santa needs help mining some AdventCoins (very similar to bitcoins) to use as gifts for all the economically forward-thinking little girls and boys.

To do this, he needs to find MD5 hashes which, in hexadecimal, start with at least five zeroes. The input to the MD5 hash is some secret key (your puzzle input, given below) followed by a number in decimal. To mine AdventCoins, you must find Santa the lowest positive number (no leading zeroes: 1, 2, 3, ...) that produces such a hash.

For example:

If your secret key is abcdef, the answer is 609043, because the MD5 hash of abcdef609043 starts with five zeroes (000001dbbfa...), and it is the lowest such number to do so.
If your secret key is pqrstuv, the lowest number it combines with to make an MD5 hash starting with five zeroes is 1048970; that is, the MD5 hash of pqrstuv1048970 looks like 000006136ef....
Your puzzle answer was 346386.

--- Part Two ---
Now find one that starts with six zeroes.

Your puzzle answer was 9958218.
*/

func Day4() {
	fmt.Println("Day 4: The Ideal Stocking Stuffer")
	fmt.Println("  First number for key to provide a 5-zero AdventCoin:", MineAdventCoin(day4input, 5))
	fmt.Println("  First number for key to provide a 6-zero AdventCoin:", MineAdventCoin(day4input, 6))
}

// MineAdventCoin takes a key to use and the number of leading zeroes needed for a hash to be a valid AdventCoin.
// It returns the lowest integer that needs to be concatenated to the key to generate a valid AdventCoin.
func MineAdventCoin(key string, zeroesNeeded int) int {
	// reuse the same hash function to avoid unessesary allocations
	hasher := md5.New()

	for i := 1; ; i++ {
		fmt.Fprintf(hasher, "%s%d", key, i)
		hash := fmt.Sprintf("%x", hasher.Sum(nil))
		zeroes := 0
		for j := 0; j < zeroesNeeded; j++ {
			if hash[j] == '0' {
				zeroes++
			}
		}
		if zeroes == zeroesNeeded {
			return i
		}
		hasher.Reset()
	}
}
