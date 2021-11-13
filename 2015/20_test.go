package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindLowestHouseNumber(t *testing.T) {
	assert.Equal(t, 8, FindLowestHouseNumber(150, -1, 10))
}

func TestDay20Pt1(t *testing.T) {
	assert.Equal(t, 776160, FindLowestHouseNumber(day20Input, -1, 10))
}

func TestDay20Pt2(t *testing.T) {
	assert.Equal(t, 786240, FindLowestHouseNumber(day20Input, 50, 11))
}
