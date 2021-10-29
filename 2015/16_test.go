package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay16Pt1(t *testing.T) {
	giftDetails := map[string]int{
		"children":    3,
		"cats":        7,
		"samoyeds":    2,
		"pomeranians": 3,
		"akitas":      0,
		"vizslas":     0,
		"goldfish":    5,
		"trees":       3,
		"cars":        2,
		"perfumes":    1,
	}
	details, err := NewSueDetails(day16Input)
	if assert.NoError(t, err) {
		assert.Equal(t, 103, details.FindFirstMatch(giftDetails, false))
	}
}

func TestDay16Pt2(t *testing.T) {
	giftDetails := map[string]int{
		"children":    3,
		"cats":        7,
		"samoyeds":    2,
		"pomeranians": 3,
		"akitas":      0,
		"vizslas":     0,
		"goldfish":    5,
		"trees":       3,
		"cars":        2,
		"perfumes":    1,
	}
	details, err := NewSueDetails(day16Input)
	if assert.NoError(t, err) {
		assert.Equal(t, 405, details.FindFirstMatch(giftDetails, true))
	}
}
