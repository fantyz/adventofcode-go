package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProcessGroupForms(t *testing.T) {
	testGroup := []string{"abcx", "abcy", "abcz"}
	yesAnswers := map[byte]int{'a': 3, 'b': 3, 'c': 3, 'x': 1, 'y': 1, 'z': 1}
	assert.Equal(t, yesAnswers, ProcessGroupForms(testGroup))
}

func TestProcessAllFormsAny(t *testing.T) {
	testData := `abc

a
b
c

ab
ac

a
a
a
a

b`

	assert.Equal(t, 11, ProcessAllForms(testData, true))
}

func TestProcessAllFormsAll(t *testing.T) {
	testData := `abc

a
b
c

ab
ac

a
a
a
a

b`

	assert.Equal(t, 6, ProcessAllForms(testData, false))
}

func TestDay6Pt1(t *testing.T) {
	assert.Equal(t, 6351, ProcessAllForms(day6Input, true))
}

func TestDay6Pt2(t *testing.T) {
	assert.Equal(t, 3143, ProcessAllForms(day6Input, false))
}
