package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRequiredFuel(t *testing.T) {
	testCases := []struct {
		In    string
		Naive int
		Real  int
	}{
		{"12", 2, 2},
		{"14", 2, 2},
		{"1969", 654, 966},
		{"100756", 33583, 50346},
		{"12\n14", 4, 4},
	}

	for i, c := range testCases {
		n, r := RequiredFuel(c.In)
		assert.Equal(t, c.Naive, n, "Naive (case=%d)", i)
		assert.Equal(t, c.Real, r, "Real (case=%d)", i)
	}
}
