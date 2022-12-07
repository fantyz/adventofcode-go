package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHasFullOverlaps(t *testing.T) {
	in := `2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8`

	assignments, err := NewSectionAssignments(in)
	if assert.NoError(t, err) {
		assert.Equal(t, 2, assignments.FullOverlaps())
	}
}

func TestHasPartialOverlaps(t *testing.T) {
	in := `2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8`

	assignments, err := NewSectionAssignments(in)
	if assert.NoError(t, err) {
		assert.Equal(t, 4, assignments.PartialOverlaps())
	}
}

func TestDay04Pt1(t *testing.T) {
	assignments, err := NewSectionAssignments(day04Input)
	if assert.NoError(t, err) {
		assert.Equal(t, 509, assignments.FullOverlaps())
	}
}

func TestDay04Pt2(t *testing.T) {
	assignments, err := NewSectionAssignments(day04Input)
	if assert.NoError(t, err) {
		assert.Equal(t, 870, assignments.PartialOverlaps())
	}
}
