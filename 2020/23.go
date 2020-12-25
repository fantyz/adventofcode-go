package main

import (
	"fmt"
)

func init() { days["23"] = Day23 }

/*
--- Day 23: Crab Cups ---
The small crab challenges you to a game! The crab is going to mix up some cups, and you have to predict where they'll end up.

The cups will be arranged in a circle and labeled clockwise (your puzzle input). For example, if your labeling were 32415, there would be five cups in the circle; going clockwise around the circle from the first cup, the cups would be labeled 3, 2, 4, 1, 5, and then back to 3 again.

Before the crab starts, it will designate the first cup in your list as the current cup. The crab is then going to do 100 moves.

Each move, the crab does the following actions:

The crab picks up the three cups that are immediately clockwise of the current cup. They are removed from the circle; cup spacing is adjusted as necessary to maintain the circle.
The crab selects a destination cup: the cup with a label equal to the current cup's label minus one. If this would select one of the cups that was just picked up, the crab will keep subtracting one until it finds a cup that wasn't just picked up. If at any point in this process the value goes below the lowest value on any cup's label, it wraps around to the highest value on any cup's label instead.
The crab places the cups it just picked up so that they are immediately clockwise of the destination cup. They keep the same order as when they were picked up.
The crab selects a new current cup: the cup which is immediately clockwise of the current cup.
For example, suppose your cup labeling were 389125467. If the crab were to do merely 10 moves, the following changes would occur:

-- move 1 --
cups: (3) 8  9  1  2  5  4  6  7
pick up: 8, 9, 1
destination: 2

-- move 2 --
cups:  3 (2) 8  9  1  5  4  6  7
pick up: 8, 9, 1
destination: 7

-- move 3 --
cups:  3  2 (5) 4  6  7  8  9  1
pick up: 4, 6, 7
destination: 3

-- move 4 --
cups:  7  2  5 (8) 9  1  3  4  6
pick up: 9, 1, 3
destination: 7

-- move 5 --
cups:  3  2  5  8 (4) 6  7  9  1
pick up: 6, 7, 9
destination: 3

-- move 6 --
cups:  9  2  5  8  4 (1) 3  6  7
pick up: 3, 6, 7
destination: 9

-- move 7 --
cups:  7  2  5  8  4  1 (9) 3  6
pick up: 3, 6, 7
destination: 8

-- move 8 --
cups:  8  3  6  7  4  1  9 (2) 5
pick up: 5, 8, 3
destination: 1

-- move 9 --
cups:  7  4  1  5  8  3  9  2 (6)
pick up: 7, 4, 1
destination: 5

-- move 10 --
cups: (5) 7  4  1  8  3  9  2  6
pick up: 7, 4, 1
destination: 3

-- final --
cups:  5 (8) 3  7  4  1  9  2  6
In the above example, the cups' values are the labels as they appear moving clockwise around the circle; the current cup is marked with ( ).

After the crab is done, what order will the cups be in? Starting after the cup labeled 1, collect the other cups' labels clockwise into a single string with no extra characters; each number except 1 should appear exactly once. In the above example, after 10 moves, the cups clockwise from 1 are labeled 9, 2, 6, 5, and so on, producing 92658374. If the crab were to complete all 100 moves, the order after cup 1 would be 67384529.

Using your labeling, simulate 100 moves. What are the labels on the cups after cup 1?

Your puzzle answer was 45983627.

--- Part Two ---
Due to what you can only assume is a mistranslation (you're not exactly fluent in Crab), you are quite surprised when the crab starts arranging many cups in a circle on your raft - one million (1000000) in total.

Your labeling is still correct for the first few cups; after that, the remaining cups are just numbered in an increasing fashion starting from the number after the highest number in your list and proceeding one by one until one million is reached. (For example, if your labeling were 54321, the cups would be numbered 5, 4, 3, 2, 1, and then start counting up from 6 until one million is reached.) In this way, every number from one through one million is used exactly once.

After discovering where you made the mistake in translating Crab Numbers, you realize the small crab isn't going to do merely 100 moves; the crab is going to do ten million (10000000) moves!

The crab is going to hide your stars - one each - under the two cups that will end up immediately clockwise of cup 1. You can have them if you predict what the labels on those cups will be when the crab is finished.

In the above example (389125467), this would be 934001 and then 159792; multiplying these together produces 149245887792.

Determine which two cups will end up immediately clockwise of cup 1. What do you get if you multiply their labels together?

Your puzzle answer was 111080192688.
*/

func Day23() {
	fmt.Println("--- Day 23: Crab Cups ---")
	cups, err := LoadInts(day23Input, "")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Cup order after 100 rounds:", PlayCrabCups(100, cups, len(cups)))
	fmt.Println("Two cups following label 1 multiplied after 10M rounds:", PlayCrabCups(10000000, cups, 1000000))
}

// PlayCrabCups plays a game of crab cups with the cups provided. Extra specifies how many extra cups
// to add in addition to the ones specified (with labels starting at len(cups)+1 and incrementing.
// It will return the sequence of cups clockwise following the cup with the label 1 if no extra cups
// are added. If extra cups are added then the returned value will be that of the labels of the two
// cups following cup 1 multiplied.
// PlayCrabCups will return -1 if the cups provided does not allow completing the game.
func PlayCrabCups(rounds int, cups []int, lastCup int) int {
	if len(cups) < 5 {
		// need at least 5 cups to play crab cups
		return -1
	}
	if lastCup < len(cups) {
		lastCup = len(cups)
	}

	// we represent the cup circle as a linked list.
	c := NewCupCircle(cups, lastCup)

	// we need an easy way to find a cup with a given label - we know labels start from 1
	// and go to the number of cups. This allow us to use a slice and let its index+1 contain
	// a pointer to the cup with that label.
	lbls := make([]*Cup, lastCup)
	for i := 0; i < lastCup; i++ {
		lbls[c.Label-1] = c
		c = c.Next
	}

	for round := 0; round < rounds; round++ {
		// remove the three cups from the circle
		c1, c2, c3 := c.Next, c.Next.Next, c.Next.Next.Next
		c.Next = c3.Next

		// identify the cup destination
		var dst *Cup
		dstLbl := c.Label - 1
		for {
			if dstLbl <= 0 {
				// higest label match the number of cups
				dstLbl = lastCup
			}
			dst = lbls[dstLbl-1]
			if dst != c1 && dst != c2 && dst != c3 {
				// found one that was not among the three cups
				break
			}
			dstLbl--
		}

		// insert the three cups after the dst cup
		c3.Next = dst.Next
		dst.Next = c1

		// move the current
		c = c.Next
	}

	// result is relative to cup labeled 1
	c = lbls[0]

	// if no additional cups beyond cups then the result should be the labels of the cups combined into a single value
	if lastCup == len(cups) {
		res := 0
		for i := 1; i < len(cups); i++ {
			c = c.Next
			res = (res + c.Label) * 10
		}
		return res / 10

	}

	// ...otherwise it should be the label of the following two cups multiplied together
	return c.Next.Label * c.Next.Next.Label
}

// NewCupCircle takes a list of cups along with the last cup label to use beyond the initial cups and returns the cup
// corresponding to the first label provided in the cups.
func NewCupCircle(cups []int, last int) *Cup {
	if len(cups) <= 0 {
		return nil
	}

	c := &Cup{Label: cups[0]}
	c.Next = c

	for i := 1; i < len(cups); i++ {
		c = c.InsertAfter(cups[i])
	}
	for i := len(cups) + 1; i <= last; i++ {
		c = c.InsertAfter(i)
	}

	// return the first cup
	return c.Next
}

type Cup struct {
	Label int
	Next  *Cup
}

// InsertAfter inserts a new cup with the specified label after itself and returns the new cup.
func (c *Cup) InsertAfter(label int) *Cup {
	cup := &Cup{
		Label: label,
		Next:  c.Next,
	}
	c.Next = cup
	return cup
}
