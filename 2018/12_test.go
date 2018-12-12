package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPots(t *testing.T) {
	testCases := []struct {
		Pots  string
		Rules string
		Ticks []string
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
			Ticks: []string{
				"#..#.#..##......###...###",
				"#...#....#.....#..#..#..#",
				"##..##...##....#..#..#..##",
			},
		},
	}

	for i, c := range testCases {
		p := NewPots(c.Pots, c.Rules)
		for _, exp := range c.Ticks {
			assert.Equal(t, exp, p.String(), "(case=%d)", i)
			p.Tick()
		}
	}
}

func TestSum(t *testing.T) {
	testCases := []struct {
		Pots   string
		Rules  string
		Ticks  int
		ExpSum int
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
			Ticks:  20,
			ExpSum: 325,
		},
	}

	for i, c := range testCases {
		p := NewPots(c.Pots, c.Rules)
		p.Ticks(c.Ticks)
		assert.Equal(t, c.ExpSum, p.Sum(), "(case=%d)", i)
	}
}
