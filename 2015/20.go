package main

import (
	"fmt"
)

func init() { days["20"] = Day20 }

/*
--- Day 20: Infinite Elves and Infinite Houses ---
To keep the Elves busy, Santa has them deliver some presents by hand, door-to-door. He sends them down a street with infinite houses numbered sequentially: 1, 2, 3, 4, 5, and so on.

Each Elf is assigned a number, too, and delivers presents to houses based on that number:

The first Elf (number 1) delivers presents to every house: 1, 2, 3, 4, 5, ....
The second Elf (number 2) delivers presents to every second house: 2, 4, 6, 8, 10, ....
Elf number 3 delivers presents to every third house: 3, 6, 9, 12, 15, ....
There are infinitely many Elves, numbered starting with 1. Each Elf delivers presents equal to ten times his or her number at each house.

So, the first nine houses on the street end up like this:

House 1 got 10 presents.
House 2 got 30 presents.
House 3 got 40 presents.
House 4 got 70 presents.
House 5 got 60 presents.
House 6 got 120 presents.
House 7 got 80 presents.
House 8 got 150 presents.
House 9 got 130 presents.
The first house gets 10 presents: it is visited only by Elf 1, which delivers 1 * 10 = 10 presents. The fourth house gets 70 presents, because it is visited by Elves 1, 2, and 4, for a total of 10 + 20 + 40 = 70 presents.

What is the lowest house number of the house to get at least as many presents as the number in your puzzle input?

Your puzzle answer was 776160.

--- Part Two ---
The Elves decide they don't want to visit an infinite number of houses. Instead, each Elf will stop after delivering presents to 50 houses. To make up for it, they decide to deliver presents equal to eleven times their number at each house.

With these changes, what is the new lowest house number of the house to get at least as many presents as the number in your puzzle input?

Your puzzle answer was 786240.
*/

func Day20() {
	fmt.Println("--- Day 20: Infinite Elves and Infinite Houses ---")
	fmt.Println("Lowest house number to receive", day20Input, "presents:", FindLowestHouseNumber(day20Input, -1, 10))
	fmt.Println("Lowest house number to receive", day20Input, "presents with elves only visting 50 houses:", FindLowestHouseNumber(day20Input, 50, 11))
}

// FindLowestHouseNumber finds the lowest house number that receives at least as many
// presents as specified in the input.
func FindLowestHouseNumber(presentsLimit int, maxVisits int, presentsPerElf int) int {
	// In reality we're looking for integers (houses) that has a lot of divisors
	// (elves visiting them).
	//
	// As far as I'm aware, just as with primes, it is non-trivial and very
	// unpredictable just how many divisors any given number would have.
	//
	// This leave no other options than just trying them out one by one the hard
	// way.
	//
	// Luckily we can avoid trying a lot of numbers by finding the house with the
	// highest highly composite number that still produce a number of presents that
	// is equal to or less than the limit. The property of the highly composite
	// numbers is that they have more divisors than any integer that is smaller
	// than it. While it would need to be proven properly, I believe that it is
	// safe to assume that no house with a non highly composite number smaller
	// than a given highly composite number would be able to produce more presents.
	//
	// Along the same lines, there would be no need to test any odd house numbers
	// as odd numbers have much fewer divisors than even ones.
	//
	// Highly composite numbers are not easily generated, so I've taken a list of
	// the first 100 which should be significantly more than what we would ever
	// need.
	// Source: https://gist.github.com/dario2994/fb4713f252ca86c1254d
	hcns := []int{1, 2, 4, 6, 12, 24, 36, 48, 60, 120, 180, 240, 360, 720, 840, 1260, 1680, 2520, 5040, 7560,
		10080, 15120, 20160, 25200, 27720, 45360, 50400, 55440, 83160, 110880, 166320, 221760, 277200, 332640,
		498960, 554400, 665280, 720720, 1081080, 1441440, 2162160, 2882880, 3603600, 4324320, 6486480, 7207200,
		8648640, 10810800, 14414400, 17297280, 21621600, 32432400, 36756720, 43243200, 61261200, 73513440,
		110270160, 122522400, 147026880, 183783600, 245044800, 294053760, 367567200, 551350800, 698377680,
		735134400, 1102701600, 1396755360, 2095133040, 2205403200, 2327925600, 2793510720, 3491888400,
		4655851200, 5587021440, 6983776800, 10475665200, 13967553600, 20951330400, 27935107200, 41902660800,
		48886437600, 64250746560, 73329656400, 80313433200, 97772875200, 128501493120, 146659312800, 160626866400,
		240940299600, 293318625600, 321253732800, 481880599200, 642507465600, 963761198400, 1124388064800,
		1606268664000}

	// find the highest highly composite number house that receives a number of presents
	// lower than or equal to the present limit
	var start int
	for _, hcn := range hcns {
		if PresentsAtHouse(hcn, maxVisits, presentsPerElf) >= presentsLimit {
			break
		}
		start = hcn
	}

	// start searching for the first number which house receives equal to or more presents
	for i := start; ; i += 2 {
		if PresentsAtHouse(i, maxVisits, presentsPerElf) >= presentsLimit {
			return i
		}
	}
}

// PresentsAtHouse takes a house number and returns the number of presents the elves will leave at this house.
func PresentsAtHouse(n int, maxVisits int, presentsPerElf int) int {
	// The naive approach to calculating the number of presents would be to run from 1..n and check each
	// if it is n divisble with it. Given we will be calling this function a lot and with large numbers
	// we need to do a bit better than that.
	//
	// An alterantive approach would be to start with n and for m=1..n check whether n is divisble with n/m.
	// This is obviously just as slow as the first approach.
	//
	// However, we can combine the two to dramatically cut down the number of calculations needed. The
	// second approch can quickly evaluate the big fractions and starts to become slow as the fractions
	// gets smaller. We can stop using the second approach when these fractions becomes too small and
	// switch to the first approach for the small amount of numbers remaining.

	// determine the first elf that potentially could deliver presents
	firstElf := n / maxVisits
	if firstElf <= 0 {
		firstElf = 1
	}

	presents := 0
	var elf int
	lastElf := 0
	for i := 2; ; i++ {
		elf = n / i

		if elf == 0 || elf == lastElf || elf < firstElf {
			// stop trying to find more elves using this approach when the fraction is too small
			// to produce steps that change the elf number
			lastElf = elf
			break
		}
		if n%elf == 0 {
			presents += elf * presentsPerElf
		}
		lastElf = elf
	}

	// iterate from 1 to lastElf to account for the remaning elves
	for elf = firstElf; elf <= lastElf; elf++ {
		if n%elf == 0 {
			presents += elf * presentsPerElf
		}
	}

	// all done, only missing presents from elf n that is always guaranteed to contribute his presents
	return presents + n*presentsPerElf
}
