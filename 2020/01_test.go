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
