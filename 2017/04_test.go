package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHasNoRepeatedWord(t *testing.T) {
	testCases := []struct {
		In     string
		Result bool
	}{
		{"aa bb cc dd ee", true},
		{"aa bb cc dd aa", false},
		{"aa bb cc dd aaa", true},
	}

	for i, testCase := range testCases {
		assert.Equal(t, testCase.Result, HasNoRepeatedWord(testCase.In), "in=%v (case=%d)", testCase.In, i)
	}
}

func TestHasNoAnagram(t *testing.T) {
	testCases := []struct {
		In     string
		Result bool
	}{
		{"abcde fghij", true},
		{"abcde xyx ecdab", false},
		{"a ab abc abd abf abj", true},
		{"iiii oiii ooii oooi oooo", true},
		{"oiii ioii iioi iiio", false},
	}

	for i, testCase := range testCases {
		assert.Equal(t, testCase.Result, HasNoAnagram(testCase.In), "in=%v (case=%d)", testCase.In, i)
	}
}
