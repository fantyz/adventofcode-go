package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumCalibrationValues(t *testing.T) {

	testCases := map[string]struct {
		In         string
		DigitsOnly bool
		ExpSum     int
	}{
		"part1": {
			`1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`,
			true,
			142,
		},
		"part2": {
			`two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen`,
			false,
			281,
		},
	}

	for name, c := range testCases {
		sum := SumCalibrationValues(c.In, c.DigitsOnly)
		assert.Equal(t, c.ExpSum, sum, name)
	}
}

func TestDay01Pt1(t *testing.T) {
	assert.Equal(t, 54081, SumCalibrationValues(day01Input, true))
}

func TestDay01Pt2(t *testing.T) {
	assert.Equal(t, 54649, SumCalibrationValues(day01Input, false))
}
