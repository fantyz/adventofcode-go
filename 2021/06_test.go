package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestModelLanternfishGrowthRate18(t *testing.T) {
	assert.Equal(t, 26, ModelLanternfishGrowthRate([]int{3, 4, 3, 1, 2}, 18))
}

func TestModelLanternfishGrowthRate256(t *testing.T) {
	assert.Equal(t, 26984457539, ModelLanternfishGrowthRate([]int{3, 4, 3, 1, 2}, 256))
}

func TestDay06Pt1(t *testing.T) {
	assert.Equal(t, 353079, ModelLanternfishGrowthRate(LoadInts(day06Input), 80))
}

func TestDay06Pt2(t *testing.T) {
	assert.Equal(t, 1605400130036, ModelLanternfishGrowthRate(LoadInts(day06Input), 256))
}
