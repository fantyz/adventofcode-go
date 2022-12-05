package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRucksack(t *testing.T) {
	rucksack := Rucksack("vJrwpWtwJgWrhcsFMMfFFhFp")
	assert.Equal(t, "vJrwpWtwJgWr", rucksack.FirstCompartment(), "first compartment")
	assert.Equal(t, "hcsFMMfFFhFp", rucksack.SecondCompartment(), "second compartment")

	itemType, err := rucksack.FindFirstError()
	if assert.NoError(t, err) {
		assert.Equal(t, byte('p'), itemType, "find first error")
	}
}

func TestRucksacksSumFirstErrorPriority(t *testing.T) {
	in := `vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw`
	rucksacks, err := NewRucksacks(in)
	if assert.NoError(t, err) {
		assert.Equal(t, 157, rucksacks.SumFirstErrorPriority())
	}
}

func TestRucksacksSumGroupBadgePriorities(t *testing.T) {
	in := `vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw`
	rucksacks, err := NewRucksacks(in)
	if assert.NoError(t, err) {
		sum, err := rucksacks.SumGroupBadgePriorities()
		if assert.NoError(t, err) {
			assert.Equal(t, 70, sum)
		}
	}
}

func TestDay03Pt1(t *testing.T) {
	rucksacks, err := NewRucksacks(day03Input)
	if assert.NoError(t, err) {
		assert.Equal(t, 7826, rucksacks.SumFirstErrorPriority())
	}
}

func TestDay03Pt2(t *testing.T) {
	rucksacks, err := NewRucksacks(day03Input)
	if assert.NoError(t, err) {
		sum, err := rucksacks.SumGroupBadgePriorities()
		if assert.NoError(t, err) {
			assert.Equal(t, 2577, sum)
		}
	}
}
