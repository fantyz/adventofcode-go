package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShortestRouteVisitingAll(t *testing.T) {
	input := `London to Dublin = 464
London to Belfast = 518
Dublin to Belfast = 141`

	s, err := NewSequenceFromCityDistances(input)
	if assert.NoError(t, err) {
		_, dist := s.Optimize(MinimizeOptimizationType, false, false)
		assert.Equal(t, 605, dist)
	}
}

func TestLongestRouteVisitingAll(t *testing.T) {
	input := `London to Dublin = 464
London to Belfast = 518
Dublin to Belfast = 141`

	s, err := NewSequenceFromCityDistances(input)
	if assert.NoError(t, err) {
		_, dist := s.Optimize(MaximizeOptimizationType, false, false)
		assert.Equal(t, 982, dist)
	}
}

func TestDay09Pt1(t *testing.T) {
	s, err := NewSequenceFromCityDistances(day09Input)
	if assert.NoError(t, err) {
		_, dist := s.Optimize(MinimizeOptimizationType, false, false)
		assert.Equal(t, 117, dist)
	}
}

func TestDay09Pt2(t *testing.T) {
	s, err := NewSequenceFromCityDistances(day09Input)
	if assert.NoError(t, err) {
		_, dist := s.Optimize(MaximizeOptimizationType, false, false)
		assert.Equal(t, 909, dist)
	}
}
