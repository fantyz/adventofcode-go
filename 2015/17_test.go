package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContainerCombinations(t *testing.T) {
	assert.Equal(t, 4, len(ContainerCombinations(25, []int{20, 15, 10, 5, 5})))
}

func TestMinimumContainerCountCombinations(t *testing.T) {
	assert.Equal(t, 3, MinimumContainerCountCombinations([][]int{{15, 10}, {20, 5}, {20, 5}, {15, 5, 5}}))
}

func TestDay17Pt1(t *testing.T) {
	containers, err := NewContainers(day17Input)
	if assert.NoError(t, err) {
		assert.Equal(t, 4372, len(ContainerCombinations(150, containers)))
	}
}

func TestDay17Pt2(t *testing.T) {
	containers, err := NewContainers(day17Input)
	if assert.NoError(t, err) {
		assert.Equal(t, 4, MinimumContainerCountCombinations(ContainerCombinations(150, containers)))
	}
}
