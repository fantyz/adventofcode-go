package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewInfectedNodes(t *testing.T) {
	testCases := []struct {
		In     string
		Result map[string]struct{}
	}{
		{`..#
#..
...`, map[string]struct{}{
			"-1,0": struct{}{},
			"1,1":  struct{}{},
		},
		},
	}

	for i, testCase := range testCases {
		assert.Equal(t, testCase.Result, NewInfectedNodes(testCase.In).nodes, "(case=%d)", i)
	}
}

func TestBurst(t *testing.T) {
	const grid = `..#
#..
...`

	testCases := []struct {
		Bursts     int
		Infections int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{70, 41},
		{10000, 5587},
	}

	for i, testCase := range testCases {
		assert.Equal(t, testCase.Infections, NewInfectedNodes(grid).Run(testCase.Bursts).Infections(), "(case=%d)", i)
	}
}

func TestEvolvedBurst(t *testing.T) {
	const grid = `..#
#..
...`

	testCases := []struct {
		Bursts     int
		Infections int
	}{
		{100, 26},
		{10000000, 2511944},
	}

	for i, testCase := range testCases {
		assert.Equal(t, testCase.Infections, NewEvolvedInfectedNodes(grid).Run(testCase.Bursts).Infections(), "(case=%d)", i)
	}
}
