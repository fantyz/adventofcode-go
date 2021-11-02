package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

func init() { days["17"] = Day17 }

/*
--- Day 17: No Such Thing as Too Much ---
The elves bought too much eggnog again - 150 liters this time. To fit it all into your refrigerator, you'll need to move it into smaller containers. You take an inventory of the capacities of the available containers.

For example, suppose you have containers of size 20, 15, 10, 5, and 5 liters. If you need to store 25 liters, there are four ways to do it:

15 and 10
20 and 5 (the first 5)
20 and 5 (the second 5)
15, 5, and 5
Filling all containers entirely, how many different combinations of containers can exactly fit all 150 liters of eggnog?

Your puzzle answer was 4372.

--- Part Two ---
While playing with all the containers in the kitchen, another load of eggnog arrives! The shipping and receiving department is requesting as many containers as you can spare.

Find the minimum number of containers that can exactly fit all 150 liters of eggnog. How many different ways can you fill that number of containers and still hold exactly 150 litres?

In the example above, the minimum number of containers was two. There were three ways to use that many containers, and so the answer there would be 3.

Your puzzle answer was 4.
*/

func Day17() {
	fmt.Println("--- Day 17: No Such Thing as Too Much ---")
	containers, err := NewContainers(day17Input)
	if err != nil {
		fmt.Println(errors.Wrap(err, "failed to load containers"))
		return
	}
	combinations := ContainerCombinations(150, containers)

	fmt.Println("Number of container combinations that will hold the 150L of eggnog:", len(combinations))
	fmt.Println("Number of container combinations using the least amount of containers only:", MinimumContainerCountCombinations(combinations))
}

// NewContainers takes the puzzle input and returns the list of containers defined as
// how much each of them can hold.
func NewContainers(input string) ([]int, error) {
	var containers []int
	for _, l := range strings.Split(input, "\n") {
		i, err := strconv.Atoi(l)
		if err != nil {
			return nil, errors.Wrapf(err, "line did not contain a valid number (line=%s)", l)
		}
		containers = append(containers, i)
	}
	return containers, nil
}

// ContainerCombinations takes the number of liters and different containers that can
// be used and returns the number of different combinations of containers that, when
// filled completely, will contain the liters.
func ContainerCombinations(liters int, containers []int) [][]int {
	return containerCombinationsRecursive(liters, nil, containers)
}

// containerCombinationsRecursive is a helper function that help ContainerCombinations
// try all possible container combinations.
func containerCombinationsRecursive(litersLeft int, containersUsed []int, containersLeft []int) [][]int {
	if litersLeft == 0 {
		// all done
		return [][]int{containersUsed}
	}

	var combinations [][]int
	// try all the possible containersLeft and get all container combinations for any that could be used
	for i := range containersLeft {
		if containersLeft[i] <= litersLeft {
			// get all container combinations using this container
			newContainersUsed := make([]int, len(containersUsed)+1)
			copy(newContainersUsed, containersUsed)
			newContainersUsed[len(newContainersUsed)-1] = containersLeft[i]

			combinations = append(combinations, containerCombinationsRecursive(litersLeft-containersLeft[i], newContainersUsed, containersLeft[i+1:])...)
		}
	}
	return combinations
}

// MinimumContainerCountCombinations returns the number of different combinations using
// the least amount of containers.
func MinimumContainerCountCombinations(combinations [][]int) int {
	count, length := 0, -1
	for _, containers := range combinations {
		if length < 0 || length > len(containers) {
			// new shortest combination found, reset count
			count = 0
			length = len(containers)
		}
		if length == len(containers) {
			count++
		}
	}
	return count
}
