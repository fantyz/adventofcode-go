package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestManhattanDistanceToShipAfterInstructions(t *testing.T) {
	testInst := []ShipInstruction{{'F', 10}, {'N', 3}, {'F', 7}, {'R', 90}, {'F', 11}}
	assert.Equal(t, 25, ManhattanDistanceToShipAfterInstructions(testInst))
}

func TestManhattanDistanceToShipWithWaypointAfterInstructions(t *testing.T) {
	testInst := []ShipInstruction{{'F', 10}, {'N', 3}, {'F', 7}, {'R', 90}, {'F', 11}}
	assert.Equal(t, 286, ManhattanDistanceToShipWithWaypointAfterInstructions(10, 1, testInst))
}

func TestDay12Pt1(t *testing.T) {
	inst, err := LoadShipInstructions(day12Input)
	if assert.NoError(t, err) {
		assert.Equal(t, 962, ManhattanDistanceToShipAfterInstructions(inst))
	}
}

func TestDay12Pt2(t *testing.T) {
	inst, err := LoadShipInstructions(day12Input)
	if assert.NoError(t, err) {
		assert.Equal(t, 56135, ManhattanDistanceToShipWithWaypointAfterInstructions(10, 1, inst))
	}
}
