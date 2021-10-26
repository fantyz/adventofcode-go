package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

func init() { days["9"] = Day9 }

/*
--- Day 9: All in a Single Night ---
Every year, Santa manages to deliver all of his presents in a single night.

This year, however, he has some new locations to visit; his elves have provided him the distances between every pair of locations. He can start and end at any two (different) locations he wants, but he must visit each location exactly once. What is the shortest distance he can travel to achieve this?

For example, given the following distances:

London to Dublin = 464
London to Belfast = 518
Dublin to Belfast = 141
The possible routes are therefore:

Dublin -> London -> Belfast = 982
London -> Dublin -> Belfast = 605
London -> Belfast -> Dublin = 659
Dublin -> Belfast -> London = 659
Belfast -> Dublin -> London = 605
Belfast -> London -> Dublin = 982
The shortest of these is London -> Dublin -> Belfast = 605, and so the answer is 605 in this example.

What is the distance of the shortest route?

Your puzzle answer was 117.

--- Part Two ---
The next year, just to show off, Santa decides to take the route with the longest distance instead.

He can still start and end at any two (different) locations he wants, and he still must visit each location exactly once.

For example, given the distances above, the longest route would be 982 via (for example) Dublin -> London -> Belfast.

What is the distance of the longest route?

Your puzzle answer was 909.

*/

func Day9() {
	fmt.Println("--- Day 9: All in a Single Night ---")
	seq, err := NewSequenceFromCityDistances(day09Input)
	if err != nil {
		fmt.Println(errors.Wrap(err, "failed to load distances"))
		return
	}
	_, shortestDist := seq.Optimize(MinimizeOptimizationType, false, false)
	fmt.Println("Shortest route visiting all locations exactly once:", shortestDist)
	_, longestDist := seq.Optimize(MaximizeOptimizationType, false, false)
	fmt.Println("Longest route visiting all locations exactly once:", longestDist)
}

// NewSequenceFromCityDistances takes the puzzle input and returns a Sequence.
// NewSequenceFromCityDistances will return an error if the input is malformed.
// If distances are missing between any two cities it will default to 0.
func NewSequenceFromCityDistances(input string) (*Sequence, error) {
	indexes := map[string]int{}
	var elements []string
	var distances [][]int

	getElementIndex := func(element string) int {
		if idx, found := indexes[element]; found {
			return idx
		}

		// not found, add to elements and expand the distances table
		elements = append(elements, element)
		idx := len(elements) - 1
		for i := range distances {
			distances[i] = append(distances[i], 0)
		}
		distances = append(distances, make([]int, len(elements)))
		indexes[element] = idx
		return idx
	}

	distExp := regexp.MustCompile(`^([a-zA-Z]+) to ([a-zA-Z]+) = ([0-9]+)$`)
	for _, line := range strings.Split(input, "\n") {
		m := distExp.FindStringSubmatch(line)
		if len(m) != 4 {
			return nil, errors.Errorf("unknown line (line=%s)", line)
		}
		idx1 := getElementIndex(m[1])
		idx2 := getElementIndex(m[2])
		dist, err := strconv.Atoi(m[3])
		if err != nil {
			return nil, errors.Wrapf(err, "unable to parse distance (line=%s)", line)
		}
		distances[idx1][idx2] = dist
		distances[idx2][idx1] = dist
	}

	seq, err := NewSequence(elements, distances)
	if err != nil {
		// should not be possible
		panic(errors.Wrap(err, "unable to create new sequence"))
	}

	return seq, nil
}
