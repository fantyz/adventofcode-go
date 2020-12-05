package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLitToggleLights(t *testing.T) {
	testCases := []struct {
		Command   string
		StartOn   bool
		LitLights int
	}{
		{"turn on 0,0 through 999,999", false, 1000 * 1000},
		{"toggle 0,0 through 999,0", false, 1000},
		{"turn off 499,499 through 500,500", false, 0},
		{"turn on 0,0 through 999,999", true, 1000 * 1000},
		{"toggle 0,0 through 999,0", true, 1000*1000 - 1000},
		{"turn off 499,499 through 500,500", true, 1000*1000 - 4},
	}

	for _, c := range testCases {
		l := NewToggleLights(1000, 1000, c.StartOn)
		_ = ExecuteInstructions(l, c.Command)
		assert.Equal(t, c.LitLights, l.LitLights(), c.Command)
	}
}

func TestLitBrightnessLights(t *testing.T) {
	testCases := []struct {
		Command   string
		LitLights int
	}{
		{"turn on 0,0 through 0,0", 1},
		{"toggle 0,0 through 999,999", 2000000},
	}

	for _, c := range testCases {
		l := NewBrightnessLights(1000, 1000, 0)
		_ = ExecuteInstructions(l, c.Command)
		assert.Equal(t, c.LitLights, l.LitLights(), c.Command)
	}
}

func TestDay6Pt1(t *testing.T) {
	l := NewToggleLights(1000, 1000, false)
	err := ExecuteInstructions(l, day6Input)
	if assert.NoError(t, err) {
		assert.Equal(t, 543903, l.LitLights())
	}
}

func TestDay6Pt2(t *testing.T) {
	l := NewBrightnessLights(1000, 1000, 0)
	err := ExecuteInstructions(l, day6Input)
	if assert.NoError(t, err) {
		assert.Equal(t, 14687245, l.LitLights())
	}
}
