package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPots(t *testing.T) {
	testCases := []struct {
		Pots  string
		Rules string
	}{
		{
			Pots: `#..#.#..##......###...###`,
			Rules: `...## => #
..#.. => #
.#... => #
.#.#. => #
.#.## => #
.##.. => #
.#### => #
#.#.# => #
#.### => #
##.#. => #
##.## => #
###.. => #
###.# => #
####. => #`,
		},
	}

	for i, c := range testCases {
		assert.Equal(t, c.Pots, NewPots(c.Pots, c.Rules).String(), "(case=%d)", i)
	}
}
