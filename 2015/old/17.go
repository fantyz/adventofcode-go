package main

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

*/

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	containers := []int{}
	for _, line := range strings.Split(puzzleInput, "\n") {
		capacity, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		containers = append(containers, capacity)
	}

	sum := 0
	for i := 0; i < len(containers); i++ {
		sum += fillIt(150, 0, containers[i:])
	}

	fmt.Println("Possible ways to store 150L eggnog:", sum)
	fmt.Println("Least containers used is:", containersUsed)
	fmt.Println("Number of solutions using the last possible containers:", best)
}

var containersUsed int = 99999
var best int

func fillIt(vol int, used int, containers []int) int {
	if len(containers) <= 0 {
		return 0
	}

	vol -= containers[0]
	if vol < 0 {
		return 0
	}
	if vol == 0 {
		if containersUsed > used {
			containersUsed = used
			best = 1
		} else if containersUsed == used {
			best++
		}
		return 1
	}

	possiblities := 0
	for i := 1; i < len(containers); i++ {
		possiblities += fillIt(vol, used+1, containers[i:])
	}
	return possiblities
}

var puzzleInput = `11
30
47
31
32
36
3
1
5
3
32
36
15
11
46
26
28
1
19
3`
