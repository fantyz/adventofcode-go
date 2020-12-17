package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApplyAndOr(t *testing.T) {
	testCases := []struct {
		Mask   string
		Input  int
		Output int
	}{
		{"XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X", 11, 73},
		{"XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X", 101, 101},
		{"XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X", 0, 64},
	}

	for _, c := range testCases {
		bitmask, err := NewBitmask(c.Mask)
		if assert.NoError(t, err, "%s - %d", c.Mask, c.Input) {
			assert.Equal(t, c.Output, bitmask.ApplyAndOr(c.Input), "%s - %d", c.Mask, c.Input)
		}
	}
}

func TestRunAndSumInitProgramAndOr(t *testing.T) {
	testProg := `mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X
mem[8] = 11
mem[7] = 101
mem[8] = 0`
	sum, err := RunAndSumInitProgram(testProg, "andor")
	if assert.NoError(t, err) {
		assert.Equal(t, 165, sum)
	}
}

func TestRunAndSumInitProgramOrX(t *testing.T) {
	testProg := `mask = 000000000000000000000000000000X1001X
mem[42] = 100
mask = 00000000000000000000000000000000X0XX
mem[26] = 1`
	sum, err := RunAndSumInitProgram(testProg, "orx")
	if assert.NoError(t, err) {
		assert.Equal(t, 208, sum)
	}
}

func TestDay14Pt1(t *testing.T) {
	sum, err := RunAndSumInitProgram(day14Input, "andor")
	if assert.NoError(t, err) {
		assert.Equal(t, 5875750429995, sum)
	}
}

func TestDay14Pt2(t *testing.T) {
	sum, err := RunAndSumInitProgram(day14Input, "orx")
	if assert.NoError(t, err) {
		assert.Equal(t, 5272149590143, sum)
	}
}
