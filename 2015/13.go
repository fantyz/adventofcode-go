package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

func init() { days["13"] = Day13 }

/*
 */

func Day13() {
	fmt.Println("--- Day 13: Knights of the Dinner Table ---")
	seq, err := NewSequenceFromSeatingHappiness(day13Input)
	if err != nil {
		fmt.Println(errors.Wrap(err, "failed load seating happiness"))
		return
	}
	_, happiness := seq.Optimize(MaximizeOptimizationType, true)
	fmt.Println("The optimal seating to maximize happiness has a total sum of:", happiness)

	seq2, err := NewSequenceFromSeatingHappiness(day13Input)
	if err != nil {
		fmt.Println(errors.Wrap(err, "failed load seating happiness"))
		return
	}
	// having one additional person seated with happiness 0 to everyone has the same
	// as not wrapping the sequence.
	_, happiness2 := seq2.Optimize(MaximizeOptimizationType, false)
	fmt.Println("The optimal seating to maximize happiness with myself at the table has a total sum of:", happiness2)
}

// NewSequenceFromSeatingHappiness takes the puzzle input and return a Sequence.
// NewSequenceFromSeatingHappiness will return an error if the input is malformed.
// Any missing happiness numbers will default to 0.
func NewSequenceFromSeatingHappiness(input string) (*Sequence, error) {
	indexes := map[string]int{}
	var elements []string
	var distances [][]int

	getElement := func(e string) int {
		if idx, found := indexes[e]; found {
			return idx
		}

		// element does not exist yet, add element and expand the distances
		elements = append(elements, e)
		idx := len(elements) - 1
		for i := range distances {
			distances[i] = append(distances[i], 0)
		}
		distances = append(distances, make([]int, len(elements)))
		indexes[e] = idx
		return idx
	}

	hapExp := regexp.MustCompile(`^([A-Za-z]+) would (gain|lose) ([0-9]+) happiness units by sitting next to ([A-Za-z]+)\.$`)
	for _, line := range strings.Split(input, "\n") {
		m := hapExp.FindStringSubmatch(line)
		if len(m) != 5 {
			return nil, errors.Errorf("line did not match hapExp (line=%s)", line)
		}

		idx1 := getElement(m[1])
		dist, err := strconv.Atoi(m[3])
		if err != nil {
			panic(errors.Wrapf(err, "distance is not a number (line=%s)", line))
		}
		idx2 := getElement(m[4])

		if m[2] == "lose" {
			dist = -dist
		}

		distances[idx1][idx2] += dist
		distances[idx2][idx1] += dist
	}

	seq, err := NewSequence(elements, distances)
	if err != nil {
		panic(errors.Wrap(err, "unable to create sequence"))
	}

	return seq, nil
}
