package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindFirstXMASWeakness(t *testing.T) {
	testData := []int{35, 20, 15, 25, 47, 40, 62, 55, 65, 95, 102, 117, 150, 182, 127, 219, 299, 277, 309, 576}
	assert.Equal(t, 14, FindFirstXMASWeakness(5, testData)) // entry 14 = 127
}

func TestFindSecondXMASWeakness(t *testing.T) {
	testData := []int{35, 20, 15, 25, 47, 40, 62, 55, 65, 95, 102, 117, 150, 182, 127, 219, 299, 277, 309, 576}
	assert.Equal(t, 62, FindSecondXMASWeakness(14, testData))
}

func TestDay9(t *testing.T) {
	input, err := LoadInts(day9Input, "\n")
	if assert.NoError(t, err) {
		weak := FindFirstXMASWeakness(25, input)
		if assert.GreaterOrEqual(t, weak, 0) {
			assert.Equal(t, 85848519, input[weak])
			assert.Equal(t, 13414198, FindSecondXMASWeakness(weak, input))
		}
	}
}
