package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExecuteOpcode(t *testing.T) {
	testCases := []struct {
		In  []int
		Out []int
	}{
		{Load("1,0,0,0,99"), Load("2,0,0,0,99")},
		{Load("2,3,0,3,99"), Load("2,3,0,6,99")},
		{Load("2,4,4,5,99,0"), Load("2,4,4,5,99,9801")},
		{Load("1,1,1,4,99,5,6,0,99"), Load("30,1,1,4,2,5,6,0,99")},
		{Load("1,9,10,3,2,3,11,0,99,30,40,50"), Load("3500,9,10,70,2,3,11,0,99,30,40,50")},
	}

	for i, c := range testCases {
		out := ExecuteOpcode(c.In)
		assert.Equal(t, c.Out, out, "(case=%d)", i)
	}
}
