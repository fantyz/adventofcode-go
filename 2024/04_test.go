package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWordSearch(t *testing.T) {
	assert.Equal(t, 18, WordSearch(`MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`, "XMAS"))
}

func TestXMASSearch(t *testing.T) {
	assert.Equal(t, 9, XMASSearch(`MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`))
}

func TestDay04Pt1(t *testing.T) {
	assert.Equal(t, 2507, WordSearch(day04Input, "XMAS"))
}

func TestDay04Pt2(t *testing.T) {
	assert.Equal(t, 1969, XMASSearch(day04Input))
}
