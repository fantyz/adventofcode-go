package main

/*

--- Day 14: Disk Defragmentation ---
Suddenly, a scheduled job activates the system's disk defragmenter. Were the situation different, you might sit and watch it for a while, but today, you just don't have that kind of time. It's soaking up valuable system resources that are needed elsewhere, and so the only option is to help it finish its task as soon as possible.

The disk in question consists of a 128x128 grid; each square of the grid is either free or used. On this disk, the state of the grid is tracked by the bits in a sequence of knot hashes.

A total of 128 knot hashes are calculated, each corresponding to a single row in the grid; each hash contains 128 bits which correspond to individual grid squares. Each bit of a hash indicates whether that square is free (0) or used (1).

The hash inputs are a key string (your puzzle input), a dash, and a number from 0 to 127 corresponding to the row. For example, if your key string were flqrgnkx, then the first row would be given by the bits of the knot hash of flqrgnkx-0, the second row from the bits of the knot hash of flqrgnkx-1, and so on until the last row, flqrgnkx-127.

The output of a knot hash is traditionally represented by 32 hexadecimal digits; each of these digits correspond to 4 bits, for a total of 4 * 32 = 128 bits. To convert to bits, turn each hexadecimal digit to its equivalent binary value, high-bit first: 0 becomes 0000, 1 becomes 0001, e becomes 1110, f becomes 1111, and so on; a hash that begins with a0c2017... in hexadecimal would begin with 10100000110000100000000101110000... in binary.

Continuing this process, the first 8 rows and columns for key flqrgnkx appear as follows, using # to denote used squares, and . to denote free ones:

##.#.#..-->
.#.#.#.#
....#.#.
#.#.##.#
.##.#...
##..#..#
.#...#..
##.#.##.-->
|      |
V      V
In this example, 8108 squares are used across the entire 128x128 grid.

Given your actual key string, how many squares are used?

Your puzzle answer was 8194.

--- Part Two ---
Now, all the defragmenter needs to know is the number of regions. A region is a group of used squares that are all adjacent, not including diagonals. Every used square is in exactly one region: lone used squares form their own isolated regions, while several adjacent squares all count as a single region.

In the example above, the following nine regions are visible, each marked with a distinct digit:

11.2.3..-->
.1.2.3.4
....5.6.
7.8.55.9
.88.5...
88..5..8
.8...8..
88.8.88.-->
|      |
V      V
Of particular interest is the region marked 8; while it does not appear contiguous in this small view, all of the squares marked 8 are connected when considering the whole 128x128 grid. In total, in this example, 1242 regions are present.

How many regions are present given your key string?

Your puzzle answer was 1141.

*/

import (
	"fmt"
)

func main() {
	fmt.Println("Advent of Code 2017 - Day 14")
	fmt.Println("Used squares:", NewGrid(128, puzzle).UsedSquares())
	fmt.Println("Used squares:", NewGrid(128, puzzle).Regions())
}

type Grid [][]bool

func (g Grid) Print() {
	for y := 0; y < len(g); y++ {
		for x := 0; x < len(g[y]); x++ {
			if g[y][x] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Print("\n")
	}
}

func (g Grid) UsedSquares() int {
	sum := 0
	for y := 0; y < len(g); y++ {
		for x := 0; x < len(g[y]); x++ {
			if g[y][x] {
				sum++
			}
		}
	}
	return sum
}

func (g Grid) Regions() int {
	sum := 0
	for y := 0; y < len(g); y++ {
		for x := 0; x < len(g[y]); x++ {
			if g[y][x] {
				sum++
				g.zeroRegion(x, y)
			}
		}
	}
	return sum
}

func (g Grid) zeroRegion(x, y int) {
	if !g[y][x] {
		return
	}

	g[y][x] = false
	if y+1 < len(g) {
		g.zeroRegion(x, y+1)
	}
	if y-1 >= 0 {
		g.zeroRegion(x, y-1)
	}
	if x+1 < len(g[y]) {
		g.zeroRegion(x+1, y)
	}
	if x-1 >= 0 {
		g.zeroRegion(x-1, y)
	}
}

func NewGrid(size int, in string) Grid {
	grid := make(Grid, 0, size)
	for i := 0; i < size; i++ {
		grid = append(grid, HashToBitRow(KnotHash(fmt.Sprintf("%s-%d", in, i))))
	}
	return grid
}

func HashToBitRow(hash []byte) []bool {
	res := make([]bool, 0, len(hash)*4)
	for _, b := range hash {
		for i := 0; i < 8; i++ {
			res = append(res, b&128 == 128)
			b = b << 1
		}
	}

	return res
}

// From: Advent of Code Day 10, slightly modified
func KnotHash(in string) []byte {
	const rounds = 64
	const size = 256
	const blocks = 16

	lengths := append([]byte(in), 17, 31, 73, 47, 23)

	elms := make([]byte, size)
	for i := 0; i < size; i++ {
		elms[i] = byte(i)
	}
	pos := 0
	skip := 0

	for round := 0; round < rounds; round++ {
		for _, length := range lengths {
			for step := 0; pos+step < pos+int(length)-step-1; step++ {
				elms[(pos+step)%size], elms[(pos+int(length)-step-1)%size] = elms[(pos+int(length)-step-1)%size], elms[(pos+step)%size]
			}
			pos += (int(length) + skip) % size
			skip++
		}
	}

	var dense [blocks]byte
	for b := 0; b < blocks; b++ {
		sparse := elms[b*blocks : (b+1)*blocks]
		dense[b] = sparse[0] ^ sparse[1] ^ sparse[2] ^ sparse[3] ^ sparse[4] ^ sparse[5] ^ sparse[6] ^ sparse[7] ^
			sparse[8] ^ sparse[9] ^ sparse[10] ^ sparse[11] ^ sparse[12] ^ sparse[13] ^ sparse[14] ^ sparse[15]
	}

	return dense[:]
}

const puzzle = `uugsqrei`
