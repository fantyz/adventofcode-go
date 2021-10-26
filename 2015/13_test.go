package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewSequenceFromSeatingHappiness(t *testing.T) {
	input := `Alice would gain 54 happiness units by sitting next to Bob.
Alice would lose 79 happiness units by sitting next to Carol.
Alice would lose 2 happiness units by sitting next to David.
Bob would gain 83 happiness units by sitting next to Alice.
Bob would lose 7 happiness units by sitting next to Carol.
Bob would lose 63 happiness units by sitting next to David.
Carol would lose 62 happiness units by sitting next to Alice.
Carol would gain 60 happiness units by sitting next to Bob.
Carol would gain 55 happiness units by sitting next to David.
David would gain 46 happiness units by sitting next to Alice.
David would lose 7 happiness units by sitting next to Bob.
David would gain 41 happiness units by sitting next to Carol.`

	seq, err := NewSequenceFromSeatingHappiness(input)
	if assert.NoError(t, err) {
		_, sum := seq.Optimize(MaximizeOptimizationType, true, true)
		assert.Equal(t, 330, sum)
	}
}

func TestDay13Pt1(t *testing.T) {
	seq, err := NewSequenceFromSeatingHappiness(day13Input)
	if assert.NoError(t, err) {
		_, happiness := seq.Optimize(MaximizeOptimizationType, true, true)
		assert.Equal(t, 733, happiness)
	}
}

func TestDay13Pt2(t *testing.T) {
	seq, err := NewSequenceFromSeatingHappiness(day13Input)
	if assert.NoError(t, err) {
		_, happiness := seq.Optimize(MaximizeOptimizationType, true, false)
		assert.Equal(t, 725, happiness)
	}
}
