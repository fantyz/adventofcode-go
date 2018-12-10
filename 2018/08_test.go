package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewInput(t *testing.T) {
	testCases := []struct {
		In  string
		Out []int
	}{
		{"2 3 0 3 10 11 12 1 1 0 1 99 2 1 1 2", []int{2, 3, 0, 3, 10, 11, 12, 1, 1, 0, 1, 99, 2, 1, 1, 2}},
	}

	for i, c := range testCases {
		assert.Equal(t, c.Out, NewInput(c.In), "(case=%d)", i)
	}
}

func TestSumMetadata(t *testing.T) {
	testCases := []struct {
		In  []int
		Out int
	}{
		{[]int{2, 3, 0, 3, 10, 11, 12, 1, 1, 0, 1, 99, 2, 1, 1, 2}, 138},
	}

	for i, c := range testCases {
		assert.Equal(t, c.Out, NewNodes(c.In).SumMetadata(), "(case=%d)", i)
	}
}

func TestValue(t *testing.T) {
	testCases := []struct {
		In  []int
		Out int
	}{
		{[]int{2, 3, 0, 3, 10, 11, 12, 1, 1, 0, 1, 99, 2, 1, 1, 2}, 66},
	}

	for i, c := range testCases {
		assert.Equal(t, c.Out, NewNodes(c.In).Value(), "(case=%d)", i)
	}
}
