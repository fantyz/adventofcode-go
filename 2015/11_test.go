package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGeneratePassword(t *testing.T) {
	testCases := map[string]struct {
		Old      string
		Policies []Policy
		New      string
	}{
		"xx, no policies": {"xx", nil, "xy"},
		"xz, no policies": {"xz", nil, "ya"},
		"z, no policies":  {"z", nil, "aa"},
		"h, no-iol":       {"h", []Policy{WithNoIOLPolicy()}, "j"},
		"n, no-iol":       {"n", []Policy{WithNoIOLPolicy()}, "p"},
		"k, no-iol":       {"k", []Policy{WithNoIOLPolicy()}, "m"},
		"a, straight":     {"a", []Policy{WithStraightOfThreePolicy()}, "abc"},
		"aba, straight":   {"aba", []Policy{WithStraightOfThreePolicy()}, "abc"},
		"aaba, pairs":     {"aaba", []Policy{WithTwoNonOverlappingPairs()}, "aabb"},
		"zzz, pairs":      {"zzz", []Policy{WithTwoNonOverlappingPairs()}, "aaaa"},
	}

	for name, c := range testCases {
		assert.Equal(t, c.New, GeneratePassword(c.Old, c.Policies...), name)
	}
}

func TestDay11Pt1(t *testing.T) {
	assert.Equal(t, "cqjxxyzz", GeneratePassword(day11Input, WithStraightOfThreePolicy(), WithNoIOLPolicy(), WithTwoNonOverlappingPairs()))
}

func TestDay11Pt2(t *testing.T) {
	assert.Equal(t, "cqkaabcc", GeneratePassword("cqjxxyzz", WithStraightOfThreePolicy(), WithNoIOLPolicy(), WithTwoNonOverlappingPairs()))
}
