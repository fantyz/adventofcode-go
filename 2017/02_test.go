package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSpreadsheetChecksum(t *testing.T) {
	testCases := []struct {
		In     string
		Result int
	}{
		{"5 1 9 5", 8},
		{"7 5 3", 4},
		{"2 4 6 8", 6},
		{"5 1 9 5\n7 5 3\n2 4 6 8", 18},
	}

	for i, testCase := range testCases {
		assert.Equal(t, testCase.Result, SpreadsheetChecksum(testCase.In), "(case=%d)", i)
	}
}

func TestEvenlyDivisibleValueSum(t *testing.T) {
	testCases := []struct {
		In     string
		Result int
	}{
		{"5 9 2 8", 4},
		{"9 4 7 3", 3},
		{"3 8 6 5", 2},
		{"5 9 2 8\n9 4 7 3\n3 8 6 5", 9},
	}

	for i, testCase := range testCases {
		assert.Equal(t, testCase.Result, EvenlyDivisibleValueSum(testCase.In), "(case=%d)", i)
	}
}
