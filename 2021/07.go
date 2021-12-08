package main

import (
	"fmt"
	"math"
	"sort"
)

func init() { days["7"] = Day7 }

/*
--- Day 7: The Treachery of Whales ---
A giant whale has decided your submarine is its next meal, and it's much faster than you are. There's nowhere to run!

Suddenly, a swarm of crabs (each in its own tiny submarine - it's too deep for them otherwise) zooms in to rescue you! They seem to be preparing to blast a hole in the ocean floor; sensors indicate a massive underground cave system just beyond where they're aiming!

The crab submarines all need to be aligned before they'll have enough power to blast a large enough hole for your submarine to get through. However, it doesn't look like they'll be aligned before the whale catches you! Maybe you can help?

There's one major catch - crab submarines can only move horizontally.

You quickly make a list of the horizontal position of each crab (your puzzle input). Crab submarines have limited fuel, so you need to find a way to make all of their horizontal positions match while requiring them to spend as little fuel as possible.

For example, consider the following horizontal positions:

16,1,2,0,4,2,7,1,2,14
This means there's a crab with horizontal position 16, a crab with horizontal position 1, and so on.

Each change of 1 step in horizontal position of a single crab costs 1 fuel. You could choose any horizontal position to align them all on, but the one that costs the least fuel is horizontal position 2:

Move from 16 to 2: 14 fuel
Move from 1 to 2: 1 fuel
Move from 2 to 2: 0 fuel
Move from 0 to 2: 2 fuel
Move from 4 to 2: 2 fuel
Move from 2 to 2: 0 fuel
Move from 7 to 2: 5 fuel
Move from 1 to 2: 1 fuel
Move from 2 to 2: 0 fuel
Move from 14 to 2: 12 fuel
This costs a total of 37 fuel. This is the cheapest possible outcome; more expensive outcomes include aligning at position 1 (41 fuel), position 3 (39 fuel), or position 10 (71 fuel).

Determine the horizontal position that the crabs can align to using the least fuel possible. How much fuel must they spend to align to that position?

Your puzzle answer was 328187.

--- Part Two ---
The crabs don't seem interested in your proposed solution. Perhaps you misunderstand crab engineering?

As it turns out, crab submarine engines don't burn fuel at a constant rate. Instead, each change of 1 step in horizontal position costs 1 more unit of fuel than the last: the first step costs 1, the second step costs 2, the third step costs 3, and so on.

As each crab moves, moving further becomes more expensive. This changes the best horizontal position to align them all on; in the example above, this becomes 5:

Move from 16 to 5: 66 fuel
Move from 1 to 5: 10 fuel
Move from 2 to 5: 6 fuel
Move from 0 to 5: 15 fuel
Move from 4 to 5: 1 fuel
Move from 2 to 5: 6 fuel
Move from 7 to 5: 3 fuel
Move from 1 to 5: 10 fuel
Move from 2 to 5: 6 fuel
Move from 14 to 5: 45 fuel
This costs a total of 168 fuel. This is the new cheapest possible outcome; the old alignment position (2) now costs 206 fuel instead.

Determine the horizontal position that the crabs can align to using the least fuel possible so they can make you an escape route! How much fuel must they spend to align to that position?

Your puzzle answer was 91257582.
*/

func Day7() {
	fmt.Println("--- Day 7: The Treachery of Whales ---")
	fmt.Println("Fuel cost to align crabs using constant fuel consumption:", FuelNeededToAlignSimpleCrabs(LoadInts(day07Input)))
	fmt.Println("Fuel cost to align crabs using non-constant fuel consumption:", FuelNeededToAlignAdvancedCrabs(LoadInts(day07Input)))
}

// FuelNeededToAlignSimpleCrabs takes a slice of crab positions and finds the least amount of fuel
// needed using a simple 1 fuel per distance cost to move all crabs to the same position. The total
// fuel consumption needed is returned.
func FuelNeededToAlignSimpleCrabs(crabs []int) int {
	// The constant fuel consumption means that the alignment where we move the least distance
	// is the optimal solution. By finding the median we have this solution.
	if len(crabs) <= 0 {
		return 0
	}

	// find the median
	sort.Ints(crabs)
	median := crabs[len(crabs)/2]
	if len(crabs)%2 == 0 {
		// even number of crabs, the median is the average of the two middle crabs
		median = int(math.Round(float64(crabs[len(crabs)/2]+crabs[(len(crabs)/2)-1]) / 2))
	}

	// calculate fuel cost for all crabs to reach the median
	fuel := 0
	for _, c := range crabs {
		cost := c - median
		if cost < 0 {
			cost = -cost
		}
		fuel += cost
	}

	return fuel
}

// FuelNeededToAlignAdvancedCrabs takes a slice of crab positions and finds the least amount of fuel
// needed using an advanced non-constant rate fuel consumption to move all crabs to the same position.
// The total fuel consumption needed is returned.
func FuelNeededToAlignAdvancedCrabs(crabs []int) int {
	// The non-constant fuel consumption rate means we need to search for the optimal solution.
	// A good starting point for the position to align around is the average of the starting
	// positions. From there we try stepping in either direction to see if a better solution
	// exists. We repeart stepping in that direction until we have a optimal solution.
	if len(crabs) <= 0 {
		return 0
	}

	// find the average
	sum := 0
	for _, c := range crabs {
		sum += c
	}
	avg := int(math.Round(float64(sum) / float64(len(crabs))))

	// we will need to calculate the fuel cost of moving the crabs to a certain position a few times
	fuelCost := func(pos int) int {
		fuel := 0
		for _, c := range crabs {
			diff := c - pos
			if diff < 0 {
				diff = -diff
			}
			for i := 1; i <= diff; i++ {
				fuel += i
			}
		}
		return fuel
	}

	// figure out whether we need to step with +1 or -1 from the average for a better solution
	bestFuelCost := fuelCost(avg)
	pos := avg
	step := 1
	if bestFuelCost < fuelCost(pos+step) {
		step = -1
	}

	// keep stepping in that direction until we have an optimal solution
	for {
		pos += step
		fuel := fuelCost(pos)
		if fuel > bestFuelCost {
			// done
			return bestFuelCost
		}
		bestFuelCost = fuel
	}
}
