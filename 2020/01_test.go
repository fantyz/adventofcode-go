package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFind2ExpenseReportEntriesMultiplied(t *testing.T) {
	v := Find2ExpenseReportEntriesMultiplied([]int{1721, 979, 366, 299, 675, 1456})
	assert.Equal(t, 514579, v)
}

func TestFind3ExpenseReportEntriesMultiplied(t *testing.T) {
	v := Find3ExpenseReportEntriesMultiplied([]int{1721, 979, 366, 299, 675, 1456})
	assert.Equal(t, 241861950, v)
}

func TestDay1Pt1(t *testing.T) {
	ints, err := LoadInts(day1Input)
	if assert.NoError(t, err) {
		assert.Equal(t, 1020099, Find2ExpenseReportEntriesMultiplied(ints))
	}
}

func TestDay1Pt2(t *testing.T) {
	ints, err := LoadInts(day1Input)
	if assert.NoError(t, err) {
		assert.Equal(t, 49214880, Find3ExpenseReportEntriesMultiplied(ints))
	}
}
