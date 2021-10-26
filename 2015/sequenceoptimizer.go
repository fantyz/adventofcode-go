package main

import (
	"fmt"

	"github.com/pkg/errors"
)

// NewSequence takes a list of elements along with a lookup table that specifies the distance between any two
// elements.
// The distances table use the index of the Elements slice to lookup.
// The distance between two places can be asymetrical depending on whether going from A to B or B to A.
func NewSequence(elements []string, distances [][]int) (*Sequence, error) {
	// validate the distances lookup table match the number of elements
	lengthMismatch := len(elements) != len(distances)
	for i := 0; i < len(distances); i++ {
		if len(distances[i]) != len(elements) {
			lengthMismatch = true
			break
		}
	}
	if lengthMismatch {
		return nil, errors.New("number of elements does mot match the dimensions of the distances specified")
	}
	return &Sequence{
		elements:  elements,
		distances: distances,
	}, nil
}

// Sequence contains a list of elements and a distance table containing the distance between any two elements.
// See NewSequence for more details.
type Sequence struct {
	elements  []string
	distances [][]int
}

// Optimize will return the sequnce of elements that is optimized in accordance with the paramters provided
// in regards to the distance between elements.
// If sumBothDirections is set to false only the distance going in one direction is included in the sum. If
// set to true both directions will be included in the sum (eg. A -> B has a distance of 1 and B -> A has a
// distance of 2 and we maximize, Optimize would return [A,B] with a sum of 3).
// If wrap is set to true the sequence will take into account that the distance between the last and first
// element must be included in the distance as well.
func (s *Sequence) Optimize(t OptimizationType, sumBothDirections bool, wrap bool) ([]string, int) {
	// default from and to to any element
	from, to := -1, -1
	elements := s.getElementIndexes()
	if wrap && len(elements) > 0 {
		// wrapping is handled by setting from and to to the same element and subsequently
		// removing the duplicate element from the result.
		from, to = elements[0], elements[0]
		elements = elements[1:]
	}

	revSeq, sum := s.optimizeRecursive(t, sumBothDirections, from, to, elements)
	seq := s.reverseAndConvertToElements(revSeq)
	if wrap && len(elements) > 0 {
		// remove the duplicate element due to wrapping from the sequence
		seq = seq[:len(seq)-1]
	}

	return seq, sum
}

type OptimizationType uint8

const (
	MaximizeOptimizationType = iota
	MinimizeOptimizationType
)

// optimizeRecursive is a recursive helper function that will create a *reverse* sequence of elements that
// is optimized according to the optimization type specified. Additionally it will return the total distance
// of the sequence.
// If from is set to -1 all possible starting elements will be considered to find the optimal solution.
func (s *Sequence) optimizeRecursive(t OptimizationType, sumBoth bool, from int, to int, elmsLeft []int) ([]int, int) {
	if len(elmsLeft) <= 1 {
		// only one possible sequence to make
		var seq []int
		if to != -1 {
			seq = append(seq, to)
		}
		if len(elmsLeft) > 0 {
			seq = append(seq, elmsLeft[0])
		}
		if from != -1 {
			seq = append(seq, from)
		}

		// sum up the distance
		dist := 0
		for i := 0; i < len(seq)-1; i++ {
			dist += s.distances[seq[i]][seq[i+1]]
			if sumBoth {
				dist += s.distances[seq[i+1]][seq[i]]
			}
		}
		return seq, dist
	}

	// multiple elements left, use each possible elmsLeft as the next from element and call itself recursively
	var bestSeq []int
	bestSeqDist := -1
	for i := range elmsLeft {
		newElmsLeft := make([]int, len(elmsLeft)-1)
		copy(newElmsLeft, elmsLeft[:i])
		copy(newElmsLeft[i:], elmsLeft[i+1:])
		seq, dist := s.optimizeRecursive(t, sumBoth, elmsLeft[i], to, newElmsLeft)

		if from >= 0 {
			dist += s.distances[from][elmsLeft[i]]
			if sumBoth {
				dist += s.distances[elmsLeft[i]][from]
			}
		}

		var isBetter bool
		switch t {
		case MinimizeOptimizationType:
			isBetter = bestSeqDist < 0 || dist < bestSeqDist
		case MaximizeOptimizationType:
			isBetter = bestSeqDist < 0 || dist > bestSeqDist
		default:
			panic(fmt.Sprintf("unsupported optimization type (type=%d)", t))
		}

		if isBetter {
			bestSeq = seq
			bestSeqDist = dist
		}
	}

	if from >= 0 {
		// include the from at the end of the best route found before returning it
		bestSeq = append(bestSeq, from)
	}

	return bestSeq, bestSeqDist
}

// getElementIndexs is a helper function that returns a list of integers with the indexes of all
// elements in the sequence.
func (s *Sequence) getElementIndexes() []int {
	res := make([]int, len(s.elements))
	for i := 0; i < len(res); i++ {
		res[i] = i
	}
	return res
}

// reverseAndConvertToElements is a helper function that reverse and convert the output of optimize
// to a list of named elements.
func (s *Sequence) reverseAndConvertToElements(seq []int) []string {
	res := make([]string, 0, len(seq))
	for i := len(seq) - 1; i >= 0; i-- {
		res = append(res, s.elements[seq[i]])
	}
	return res
}
