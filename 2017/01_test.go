package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInverseCapcha(t *testing.T) {
	testCases := []struct {
		Step   int
		In     string
		Result int
	}{
		{1, "1122", 3},
		{1, "1111", 4},
		{1, "91212129", 9},
		{2, "1212", 6},
		{2, "1221", 0},
		{3, "123425", 4},
		{3, "123123", 12},
		{4, "12131415", 4},
	}

	for i, testCase := range testCases {
		assert.Equal(t, testCase.Result, InverseCapcha(testCase.Step, testCase.In), "(case=%d)", i)
	}
}
