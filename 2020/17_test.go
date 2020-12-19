package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestActiveCubes3D(t *testing.T) {
	testInitialState := `.#.
..#
###`

	testCases := []struct {
		Cycles      int
		ActiveCubes int
	}{
		{0, 5},
		{1, 11},
		{2, 21},
		{3, 38},
		{6, 112},
	}

	for _, c := range testCases {
		p := NewPocketDimension3D(testInitialState)
		p.RunCycles(c.Cycles)
		assert.Equal(t, c.ActiveCubes, p.ActiveCubes(), "cycles=%d", c.Cycles)
	}
}

func TestActiveCubes4D(t *testing.T) {
	testInitialState := `.#.
..#
###`

	testCases := []struct {
		Cycles      int
		ActiveCubes int
	}{
		{0, 5},
		{1, 29},
		{6, 848},
	}

	for _, c := range testCases {
		p := NewPocketDimension4D(testInitialState)
		p.RunCycles(c.Cycles)
		assert.Equal(t, c.ActiveCubes, p.ActiveCubes(), "cycles=%d", c.Cycles)
	}
}

func TestDay17Pt1(t *testing.T) {
	p := NewPocketDimension3D(day17Input)
	p.RunCycles(6)
	assert.Equal(t, 372, p.ActiveCubes())
}

func TestDay17Pt2(t *testing.T) {
	p := NewPocketDimension4D(day17Input)
	p.RunCycles(6)
	assert.Equal(t, 1896, p.ActiveCubes())
}
