package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVerifyImage(t *testing.T) {
	testCases := []struct {
		In    string
		Check int
	}{
		{"123456789012", 1},
		{"111222789012", 9},
	}

	for i, c := range testCases {
		assert.Equal(t, c.Check, VerifyImage(NewSpaceImage(Load(c.In), 3, 2)), "(case=%d)", i)
	}
}
