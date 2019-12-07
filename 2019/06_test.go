package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrbitCount(t *testing.T) {
	testCases := []struct {
		In    string
		Count int
	}{
		{"COM)B\nB)C\nC)D\nD)E\nE)F\nB)G\nG)H\nD)I\nE)J\nJ)K\nK)L", 42},
	}

	for i, c := range testCases {
		byMoons, _ := Load(c.In)
		assert.Equal(t, c.Count, OrbitCount("COM", 0, byMoons), "(case=%d)", i)
	}
}

func TestTransfers(t *testing.T) {
	testCases := []struct {
		In    string
		Count int
	}{
		{"COM)B\nB)C\nC)D\nD)E\nE)F\nB)G\nG)H\nD)I\nE)J\nJ)K\nK)L\nK)YOU\nI)SAN", 4},
	}

	for i, c := range testCases {
		byMoons, byOrbits := Load(c.In)
		assert.Equal(t, c.Count, Transfers("YOU", "SAN", "", 0, byMoons, byOrbits), "(case=%d)", i)
	}
}
