package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadInts(t *testing.T) {
	testCases := []struct {
		In   string
		Ints []int
	}{
		{In: "1,2,3", Ints: []int{1, 2, 3}},
		{In: "1  2 3", Ints: []int{1, 2, 3}},
		{In: "1\n2\n3\n", Ints: []int{1, 2, 3}},
	}

	for _, c := range testCases {
		assert.Equal(t, c.Ints, LoadInts(c.In), "in=%s", c.In)
	}
}
