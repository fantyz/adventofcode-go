package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumCarriedFromTopN(t *testing.T) {
	in := `1000
2000
3000

4000

5000
6000

7000
8000
9000

10000`

	testCases := map[string]struct {
		N      int
		ExpSum int
	}{
		"1": {1, 24000},
		"3": {3, 45000},
	}

	cals := NewCarriedCalories(in)

	for name, c := range testCases {
		assert.Equal(t, c.ExpSum, cals.SumCarriedFromTopN(c.N), name)
	}
}

func TestDay01Pt1(t *testing.T) {
	assert.Equal(t, 64929, NewCarriedCalories(day01Input).SumCarriedFromTopN(1))
}

func TestDay01Pt2(t *testing.T) {
	assert.Equal(t, 193697, NewCarriedCalories(day01Input).SumCarriedFromTopN(3))
}
