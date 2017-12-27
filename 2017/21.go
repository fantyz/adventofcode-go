package main

/*

--- Day 21: Fractal Art ---
You find a program trying to generate some art. It uses a strange process that involves repeatedly enhancing the detail of an image through a set of rules.

The image consists of a two-dimensional square grid of pixels that are either on (#) or off (.). The program always begins with this pattern:

.#.
..#
###
Because the pattern is both 3 pixels wide and 3 pixels tall, it is said to have a size of 3.

Then, the program repeats the following process:

If the size is evenly divisible by 2, break the pixels up into 2x2 squares, and convert each 2x2 square into a 3x3 square by following the corresponding enhancement rule.
Otherwise, the size is evenly divisible by 3; break the pixels up into 3x3 squares, and convert each 3x3 square into a 4x4 square by following the corresponding enhancement rule.
Because each square of pixels is replaced by a larger one, the image gains pixels and so its size increases.

The artist's book of enhancement rules is nearby (your puzzle input); however, it seems to be missing rules. The artist explains that sometimes, one must rotate or flip the input pattern to find a match. (Never rotate or flip the output pattern, though.) Each pattern is written concisely: rows are listed as single units, ordered top-down, and separated by slashes. For example, the following rules correspond to the adjacent patterns:

../.#  =  ..
          .#

                .#.
.#./..#/###  =  ..#
                ###

                        #..#
#..#/..../#..#/.##.  =  ....
                        #..#
                        .##.
When searching for a rule to use, rotate and flip the pattern as necessary. For example, all of the following patterns match the same rule:

.#.   .#.   #..   ###
..#   #..   #.#   ..#
###   ###   ##.   .#.
Suppose the book contained the following two rules:

../.# => ##./#../...
.#./..#/### => #..#/..../..../#..#
As before, the program begins with this pattern:

.#.
..#
###
The size of the grid (3) is not divisible by 2, but it is divisible by 3. It divides evenly into a single square; the square matches the second rule, which produces:

#..#
....
....
#..#
The size of this enhanced grid (4) is evenly divisible by 2, so that rule is used. It divides evenly into four squares:

#.|.#
..|..
--+--
..|..
#.|.#
Each of these squares matches the same rule (../.# => ##./#../...), three of which require some flipping and rotation to line up with the rule. The output for the rule is the same in all four cases:

##.|##.
#..|#..
...|...
---+---
##.|##.
#..|#..
...|...
Finally, the squares are joined into a new grid:

##.##.
#..#..
......
##.##.
#..#..
......
Thus, after 2 iterations, the grid contains 12 pixels that are on.

How many pixels stay on after 5 iterations?

Your puzzle answer was 152.

--- Part Two ---
How many pixels stay on after 18 iterations?

Your puzzle answer was 1956174.

*/

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	fmt.Println("Advent of Code 2017 - Day 21")

	rules := NewRules(puzzle)
	p := NewStartingPainting()
	for i := 0; i < 5; i++ {
		p.Print()
		p = p.Enhance(rules)
	}

	fmt.Println("Pixels on after 5 iterations:", p.PixelsOn())

	p = NewStartingPainting()
	for i := 0; i < 18; i++ {
		p = p.Enhance(rules)
	}
	fmt.Println("Pixels on after 5 iterations:", p.PixelsOn())
}

func NewStartingPainting() *Painting {
	return NewPainting(".#./..#/###")
}

func NewPainting(in string) *Painting {
	rows := strings.Split(in, "/")
	p := &Painting{
		canvas: make([][]bool, len(rows)),
	}

	for y := range rows {
		p.canvas[y] = make([]bool, 0, len(rows[y]))
		for x := 0; x < len(rows[y]); x++ {
			switch rows[y][x] {
			case '#':
				p.canvas[y] = append(p.canvas[y], true)
			case '.':
				p.canvas[y] = append(p.canvas[y], false)
			default:
				panic("Unsupported painting input: " + rows[y])
			}
		}
	}
	return p
}

type Painting struct {
	canvas [][]bool
}

func (p *Painting) Enhance(r Rules) *Painting {
	var step int
	switch {
	case len(p.canvas)%2 == 0:
		step = 2
	case len(p.canvas)%3 == 0:
		step = 3
	default:
		panic("Not divisable by neither 2 or 3")
	}

	// initiate new canvas
	size := (len(p.canvas) / step) * (step + 1)
	np := &Painting{
		canvas: make([][]bool, size),
	}
	for i := 0; i < size; i++ {
		np.canvas[i] = make([]bool, size)
	}

	for y := 0; y < len(p.canvas)/step; y++ {
		for x := 0; x < len(p.canvas[y])/step; x++ {
			enhanced, found := r.Match(p.CopySubpainting(x*step, y*step, step))
			if !found {
				panic("Enhanced painting not found: " + p.CopySubpainting(x, y, step).String())
			}
			np.PasteSubpainting(x*(step+1), y*(step+1), enhanced)
		}
	}
	return np
}

func (p *Painting) Flip() *Painting {
	for y := 0; y < len(p.canvas); y++ {
		for x := 0; x < len(p.canvas[y])/2; x++ {
			p.canvas[y][x], p.canvas[y][len(p.canvas[y])-1-x] = p.canvas[y][len(p.canvas[y])-1-x], p.canvas[y][x]
		}
	}
	return p
}

func (p *Painting) Rotate() *Painting {
	// transpose
	for y := 0; y < len(p.canvas); y++ {
		for x := 0; x < y; x++ {
			p.canvas[y][x], p.canvas[x][y] = p.canvas[x][y], p.canvas[y][x]
		}
	}
	// reverse rows
	p.Flip()
	return p
}

func (p *Painting) PixelsOn() int {
	on := 0
	for y := 0; y < len(p.canvas); y++ {
		for x := 0; x < len(p.canvas[y]); x++ {
			if p.canvas[y][x] {
				on++
			}
		}
	}
	return on
}

func (p *Painting) Print() {
	for y := 0; y < len(p.canvas); y++ {
		for x := 0; x < len(p.canvas[y]); x++ {
			if p.canvas[y][x] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Print("\n")
	}
	fmt.Println()
}

func (p *Painting) String() string {
	str := ""
	for y := 0; y < len(p.canvas); y++ {
		for x := 0; x < len(p.canvas[y]); x++ {
			if p.canvas[y][x] {
				str += "#"
			} else {
				str += "."
			}
		}
		str += "/"
	}
	return str[:len(str)-1]
}

func (p *Painting) CopySubpainting(xoff, yoff, size int) *Painting {
	if xoff+size > len(p.canvas) {
		panic(fmt.Sprintf("Invalid x offset (xoff=%d, size=%d)", xoff, size))
	}
	if xoff+size > len(p.canvas) {
		panic(fmt.Sprintf("Invalid y offset (yoff=%d, size=%d)", yoff, size))
	}

	subp := &Painting{
		canvas: make([][]bool, size),
	}
	for y := 0; y < size; y++ {
		subp.canvas[y] = make([]bool, 0, size)
		for x := xoff; x < xoff+size; x++ {
			subp.canvas[y] = append(subp.canvas[y], p.canvas[y+yoff][x])
		}
	}
	return subp
}

func (p *Painting) PasteSubpainting(xoff, yoff int, subp *Painting) *Painting {
	for y := 0; y < len(subp.canvas); y++ {
		for x := 0; x < len(subp.canvas[y]); x++ {
			p.canvas[yoff+y][xoff+x] = subp.canvas[y][x]
		}
	}
	return p
}

func NewRules(in string) Rules {
	r := make(Rules)
	exp := regexp.MustCompile(`^([.#/]+) => ([.#/]+)$`)
	for _, line := range strings.Split(in, "\n") {
		m := exp.FindStringSubmatch(line)
		if len(m) != 3 {
			panic("Line did not match: " + line)
		}

		key := NewPainting(m[1])
		val := NewPainting(m[2])

		// all rotations
		for i := 0; i < 4; i++ {
			key = key.Rotate()
			if v, found := r[key.String()]; found {
				if v.String() != val.String() {
					panic("Ambigous rule: " + line)
				}
			}
			r[key.String()] = val
		}

		// mirrored
		key = key.Flip()
		if v, found := r[key.String()]; found {
			if v.String() != val.String() {
				panic("Ambigous rule: " + line)
			}
		}
		for i := 0; i < 4; i++ {
			key = key.Rotate()
			if v, found := r[key.String()]; found {
				if v.String() != val.String() {
					panic("Ambigous rule: " + line)
				}
			}
			r[key.String()] = val
		}
	}
	return r
}

type Rules map[string]*Painting

func (r Rules) Match(p *Painting) (*Painting, bool) {
	newp, found := r[p.String()]
	return newp, found
}

const puzzle = `../.. => ###/.../#..
#./.. => #../.#./###
##/.. => #../###/#..
.#/#. => .#./##./#.#
##/#. => .#./.../...
##/## => .../#../#.#
.../.../... => #.##/#.../..##/.##.
#../.../... => .#../#..#/###./#.##
.#./.../... => ..../###./.###/##..
##./.../... => .###/##../.#.#/..#.
#.#/.../... => ..#./...#/.##./#.##
###/.../... => ##.#/..../#.../....
.#./#../... => .##./..../.##./.###
##./#../... => .##./##../#.#./#.#.
..#/#../... => .#../####/##.#/##.#
#.#/#../... => .##./##.#/#..#/#...
.##/#../... => .#../.#.#/####/....
###/#../... => .###/#.#./.#../...#
.../.#./... => ##.#/####/##../#..#
#../.#./... => #..#/..##/.###/.##.
.#./.#./... => #.../#..#/...#/#.#.
##./.#./... => .#.#/#..#/..#./##.#
#.#/.#./... => .##./#.#./##../#.##
###/.#./... => #.../...#/.#../.#.#
.#./##./... => ####/#.../#.../###.
##./##./... => ##../##../.#../##..
..#/##./... => ##.#/####/..##/.##.
#.#/##./... => .#.#/.##./..../#..#
.##/##./... => ..../.###/..#./###.
###/##./... => .###/.##./..../..#.
.../#.#/... => .###/####/..##/#.#.
#../#.#/... => .#../.##./..#./....
.#./#.#/... => ..##/#.##/###./#...
##./#.#/... => ...#/#..#/#.../##..
#.#/#.#/... => ####/#.#./####/.#..
###/#.#/... => #.##/####/..../#.#.
.../###/... => #..#/.##./.##./.#..
#../###/... => #..#/...#/#..#/...#
.#./###/... => ..##/###./..##/#.##
##./###/... => ###./####/#.../####
#.#/###/... => ..../#..#/.##./.##.
###/###/... => .###/#.#./#.#./#...
..#/.../#.. => #..#/...#/.###/.#.#
#.#/.../#.. => ...#/.###/..../.#..
.##/.../#.. => ..##/..##/#.../..#.
###/.../#.. => ..../..##/#.../.#.#
.##/#../#.. => ...#/#..#/##../.##.
###/#../#.. => .##./###./.##./#.#.
..#/.#./#.. => .##./..#./###./.#..
#.#/.#./#.. => .#.#/...#/#..#/.#..
.##/.#./#.. => #.../#.##/.###/.#..
###/.#./#.. => ##../#.#./#.##/#...
.##/##./#.. => ..#./##.#/..#./..##
###/##./#.. => ..##/..#./.##./###.
#../..#/#.. => ...#/.#.#/..../.#..
.#./..#/#.. => ..../.#.#/#..#/...#
##./..#/#.. => #.#./#.##/..#./.###
#.#/..#/#.. => #.#./...#/.#.#/....
.##/..#/#.. => .###/.##./..#./....
###/..#/#.. => .##./###./#.../.###
#../#.#/#.. => #.../..##/...#/..##
.#./#.#/#.. => ..##/###./###./..#.
##./#.#/#.. => ...#/.##./#..#/#.#.
..#/#.#/#.. => #..#/...#/###./###.
#.#/#.#/#.. => ####/...#/..#./##..
.##/#.#/#.. => ..../..#./..../#.##
###/#.#/#.. => .#../#.#./.###/#...
#../.##/#.. => .###/####/.#.#/.#..
.#./.##/#.. => ###./#.#./..../.#..
##./.##/#.. => #.##/..#./#.#./.##.
#.#/.##/#.. => ..#./..#./..../..#.
.##/.##/#.. => #..#/#.##/.#.#/###.
###/.##/#.. => .###/..../.#.#/....
#../###/#.. => ###./..#./..../#.##
.#./###/#.. => ..../##.#/####/####
##./###/#.. => ..#./##.#/.###/.###
..#/###/#.. => .###/.#../####/#.##
#.#/###/#.. => ...#/##../..../.#.#
.##/###/#.. => #.##/..#./#.../##..
###/###/#.. => ###./#.##/.###/....
.#./#.#/.#. => #.#./#.../..#./#.##
##./#.#/.#. => ####/.##./...#/.##.
#.#/#.#/.#. => ###./##.#/#.../#..#
###/#.#/.#. => .###/.#.#/..../...#
.#./###/.#. => #..#/.###/..#./.##.
##./###/.#. => .#.#/#.##/.#.#/###.
#.#/###/.#. => ...#/...#/##.#/....
###/###/.#. => ####/#.##/##../.#..
#.#/..#/##. => #.../...#/####/#...
###/..#/##. => ####/###./.##./.##.
.##/#.#/##. => .#.#/#.../####/####
###/#.#/##. => .#../#.#./.#../#.##
#.#/.##/##. => ##.#/#..#/#.../..##
###/.##/##. => ####/.#../.##./###.
.##/###/##. => .#../..#./#..#/.##.
###/###/##. => ...#/.#../..../..#.
#.#/.../#.# => #..#/##../#.##/###.
###/.../#.# => ..##/.#.#/#.#./#.#.
###/#../#.# => .#.#/###./..../.###
#.#/.#./#.# => ..##/#.##/#.../####
###/.#./#.# => ####/###./...#/.###
###/##./#.# => ..../##../####/.###
#.#/#.#/#.# => ..../..../.##./####
###/#.#/#.# => ##../#..#/.#.#/#..#
#.#/###/#.# => #..#/..../####/#.##
###/###/#.# => ####/..../.##./.#.#
###/#.#/### => .#.#/#..#/###./.##.
###/###/### => ##../#.##/.##./#..#`
