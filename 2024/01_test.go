package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDistanceHistoryLists(t *testing.T) {
	dist := DistanceHistoryLists(
		[]int{3,4,2,1,3,3},
		[]int{4,3,5,3,9,3},
	)

	assert.Equal(t, 11, dist)
}

func TestSimilarityScoreHistoryLists(t *testing.T) {
	score := SimilarityScoreHistoryLists(
		[]int{3,4,2,1,3,3},
		[]int{4,3,5,3,9,3},
	)

	assert.Equal(t, 31, score)
}

func TestDay01Pt1(t *testing.T) {
	listA, listB, err := ReadHistoryLists(day01Input)
	if assert.NoError(t, err) {
		assert.Equal(t, 1341714, DistanceHistoryLists(listA, listB))
	}
}

func TestDay01Pt2(t *testing.T) {
	listA, listB, err := ReadHistoryLists(day01Input)
	if assert.NoError(t, err) {
		assert.Equal(t, 27384707, SimilarityScoreHistoryLists(listA, listB))
	}
}

