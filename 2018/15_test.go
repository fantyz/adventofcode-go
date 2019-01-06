package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
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

func TestMove(t *testing.T) {
	testCases := []struct {
		In       string
		ExpSteps []string
	}{
		{`#########
#G......#
#.......#
#.......#
#...E...#
#.......#
#.......#
#.......#
#########
`, []string{`#########
#.G.....#
#.......#
#...E...#
#.......#
#.......#
#.......#
#.......#
#########
`},
		},
		{`#########
#G..G..G#
#.......#
#.......#
#G..E..G#
#.......#
#.......#
#G..G..G#
#########
`, []string{`#########
#.G...G.#
#...G...#
#...E..G#
#.G.....#
#.......#
#G..G..G#
#.......#
#########
`, `#########
#..G.G..#
#...G...#
#.G.E.G.#
#.......#
#G..G..G#
#.......#
#.......#
#########
`, `#########
#.......#
#..GGG..#
#..GEG..#
#G..G...#
#......G#
#.......#
#.......#
#########
`},
		},
		{`#####
#E..#
#...#
#..G#
#####`, []string{`#####
#.E.#
#..G#
#...#
#####
`},
		},
	}

	for i, c := range testCases {
		cave := NewCave(c.In)
		for j, expCave := range c.ExpSteps {
			cave.Tick()
			if !assert.Equal(t, expCave, cave.String(), "(step=%d, case=%d)", j, i) {
				break
			}
		}
	}
}
