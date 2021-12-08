package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFuelNeededToAlignSimpleCrabs(t *testing.T) {
	assert.Equal(t, 37, FuelNeededToAlignSimpleCrabs([]int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}))
}

func TestFuelNeededToAlignAdvancedCrabs(t *testing.T) {
	assert.Equal(t, 168, FuelNeededToAlignAdvancedCrabs([]int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}))
}

func TestDay07Pt1(t *testing.T) {
	assert.Equal(t, 328187, FuelNeededToAlignSimpleCrabs(LoadInts(day07Input)))
}

func TestDay07Pt2(t *testing.T) {
	assert.Equal(t, 91257582, FuelNeededToAlignAdvancedCrabs(LoadInts(day07Input)))
}
