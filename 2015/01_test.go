package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeliverPresents(t *testing.T) {
	testCases := []struct {
		Instructions    string
		LastFloor       int
		FirstInBasement int
	}{
		{")", -1, 1},
		{"()())", -1, 5},
		{"(())", 0, -1},
		{"()()", 0, -1},
		{"(((", 3, -1},
		{"(()(()(", 3, -1},
		{"))(((((", 3, 1},
		{"())", -1, 3},
		{"))(", -1, 1},
		{")))", -3, 1},
		{")())())", -3, 1},
	}

	for _, c := range testCases {
		f, i := DeliverPresents(c.Instructions)
		assert.Equal(t, c.LastFloor, f, "LastFloor [%s]", c.Instructions)
		assert.Equal(t, c.FirstInBasement, i, "FirstInBasement [%s]", c.Instructions)
	}
}

func TestDay1(t *testing.T) {
	finalFloor, firstBasementIdx := DeliverPresents(day1Input)
	assert.Equal(t, 280, finalFloor, "pt1")
	assert.Equal(t, 1797, firstBasementIdx, "pt2")
}
