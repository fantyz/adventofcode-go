package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestChronalCalibration(t *testing.T) {
	testCases := []struct {
		In  string
		Out int
	}{}

	for i, c := range testCases {
		i := a(c.In)
		assert.Equal(t, c.Out, i, "(case=%d)", i)
	}
}
