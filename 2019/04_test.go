package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValid(t *testing.T) {
	testCases := []struct {
		In    int
		Valid bool
	}{
		{111111, false},
		{223450, false},
		{123789, false},
		{112233, true},
		{123444, false},
		{111122, true},
	}

	for i, c := range testCases {
		assert.Equal(t, c.Valid, IsValid(c.In), "%d (case=%d)", c.In, i)
	}
}
