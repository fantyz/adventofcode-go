package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

// Cave represents a generic cave.
// Note that Cave is indexed `val = cave[y][x]`.
type Cave [][]int

// CaveLoc represents a 2D position within a cave.
type CaveLoc struct {
	X, Y int
}

// NewCave takes a string representing a 2D rectangular cave having each location represented by a
// single digit integer.
// NewCave will return an error if any non-integers are found in the input or the lines do form a
// rectangular shape.
func NewCave(in string) (Cave, error) {
	var c Cave
	width := -1

	for _, line := range strings.Split(in, "\n") {
		if width < 0 {
			width = len(line)
		}
		if len(line) != width {
			return nil, errors.Errorf("unexpected length of row in cave (count=%d, expected=%d)", len(line), width)
		}

		row := make([]int, 0, width)
		for _, c := range line {
			val, err := strconv.Atoi(string(c))
			if err != nil {
				return nil, errors.Wrapf(err, "bad cave value (val=%s)", string(c))
			}
			row = append(row, val)
		}
		c = append(c, row)
	}

	return c, nil
}

// Neighbors return all neighbors adjacant to the specified location. Neighbors only return
// valid locations.
func (c Cave) Neighbors(pos CaveLoc, inclDiagonal bool) []CaveLoc {
	if len(c) <= 0 {
		return nil
	}

	n := []CaveLoc{
		{pos.X - 1, pos.Y},
		{pos.X, pos.Y + 1},
		{pos.X + 1, pos.Y},
		{pos.X, pos.Y - 1},
	}
	if inclDiagonal {
		n = append(n,
			CaveLoc{pos.X - 1, pos.Y - 1},
			CaveLoc{pos.X + 1, pos.Y - 1},
			CaveLoc{pos.X - 1, pos.Y + 1},
			CaveLoc{pos.X + 1, pos.Y + 1},
		)
	}

	// remove any neighbors that are ouside the cave
	i := 0
	for i < len(n) {
		if n[i].X >= 0 && n[i].X < len(c[0]) && n[i].Y >= 0 && n[i].Y < len(c) {
			// good neighbor
			i++
		} else {
			// neighbor is outside cave, remove it from list
			n[i], n = n[len(n)-1], n[:len(n)-1]
		}
	}

	return n
}

// Enlarge enlarges the cave with the factor specified. It is enlarged by repeating the existing cave
// and adding +1 to the repeated values for each repetion.
// See puzzle input for day 15 for more information.
// If the factor specified is 1 or less Enlarge does nothing.
func (c *Cave) Enlarge(factor int) {
	if factor <= 1 {
		return
	}

	incWithWrap := func(n int) int {
		n++
		if n >= 10 {
			n = 1
		}
		return n
	}

	startWidth, startHeight := c.Width(), c.Height()

	for y := 0; y < startHeight*factor; y++ {
		row := make([]int, startWidth*factor)
		if y < startHeight {
			// initialize first values of row using the original cave values
			for x := 0; x < startWidth; x++ {
				row[x] = (*c)[y][x]
			}
		} else {
			// initialize first values of row using the previous repetion + 1
			for x := 0; x < startWidth; x++ {
				row[x] = incWithWrap((*c)[y-startWidth][x])
			}
		}

		// finish the row
		for x := startWidth; x < startWidth*factor; x++ {
			row[x] = incWithWrap(row[x-startWidth])
		}

		// update cave with the new row
		if y < len(*c) {
			(*c)[y] = row
		} else {
			*c = append(*c, row)
		}
	}
}

// Width returns the width of the cave.
func (c Cave) Width() int {
	if len(c) <= 0 {
		return 0
	}
	return len(c[0])
}

// Height returns the height of the cave.
func (c Cave) Height() int {
	return len(c)
}

// TopLeft returns the cave location of the top left position in the cave.
func (c Cave) TopLeft() CaveLoc {
	return CaveLoc{0, 0}
}

// BottomRight returns the cave location of the bottom right position of the cave.
func (c Cave) BottomRight() CaveLoc {
	return CaveLoc{c.Width() - 1, c.Height() - 1}
}

// String returns a string representation of the cave.
func (c Cave) String() string {
	if len(c) <= 0 {
		return ""
	}

	// very inefficient implemntation - lots of garbage strings here
	var str string
	for y := 0; y < len(c); y++ {
		for x := 0; x < len(c[y]); x++ {
			str += fmt.Sprintf("%d", c[y][x])
		}
		str += "\n"
	}

	return str[:len(str)-1]
}
