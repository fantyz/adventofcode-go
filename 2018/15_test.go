package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewCave(t *testing.T) {
	testCases := []struct {
		In string
	}{
		{`#########
#G..G..G#
#.......#
#.......#
#G..E..G#
#.......#
#.......#
#G..G..G#
#########
`},
	}

	for i, c := range testCases {
		cave := NewCave(c.In)
		assert.Equal(t, c.In, cave.String(), "(case=%d)", i)
	}
}
