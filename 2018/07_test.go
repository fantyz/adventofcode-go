package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewGraph(t *testing.T) {
	testCases := []struct {
		In      string
		Workers int
		Offset  int
		Out     string
		OutTime int
	}{
		{
			In: `Step C must be finished before step A can begin.
Step C must be finished before step F can begin.
Step A must be finished before step B can begin.
Step A must be finished before step D can begin.
Step B must be finished before step E can begin.
Step D must be finished before step E can begin.
Step F must be finished before step E can begin.`,
			Workers: 1,
			Offset:  0,
			Out:     "CABDFE",
			OutTime: 21,
		},
		{
			In: `Step C must be finished before step A can begin.
Step C must be finished before step F can begin.
Step A must be finished before step B can begin.
Step A must be finished before step D can begin.
Step B must be finished before step E can begin.
Step D must be finished before step E can begin.
Step F must be finished before step E can begin.`,
			Workers: 2,
			Offset:  0,
			Out:     "CABFDE",
			OutTime: 15,
		},
	}

	for i, c := range testCases {
		order, time := NewGraph(c.In).Order(c.Workers, c.Offset)
		assert.Equal(t, c.Out, order, "(case=%d)", i)
		assert.Equal(t, c.OutTime, time, "(case=%d)", i)
	}
}

func TestTime(t *testing.T) {
	testCases := []struct {
		Step *Step
		Out  int
	}{
		{&Step{Name: "A"}, 1},
		{&Step{Name: "Z"}, 26},
	}

	for i, c := range testCases {
		assert.Equal(t, c.Out, c.Step.Time(), "(case=%d)", i)
	}
}
