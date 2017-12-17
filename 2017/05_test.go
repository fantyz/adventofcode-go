package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestJumpsNeededToExit(t *testing.T) {
	testCases := []struct {
		In     string
		Result int
	}{
		{"0\n3\n0\n1\n-3", 5},
	}

	for i, testCase := range testCases {
		assert.Equal(t, testCase.Result, JumpsNeededToExit(testCase.In), "in=%v (case=%d)", testCase.In, i)
	}
}

func TestWeirderJumpsNeededToExit(t *testing.T) {
	testCases := []struct {
		In     string
		Result int
	}{
		{"0\n3\n0\n1\n-3", 10},
	}

	for i, testCase := range testCases {
		assert.Equal(t, testCase.Result, WeirderJumpsNeededToExit(testCase.In), "in=%v (case=%d)", testCase.In, i)
	}
}
