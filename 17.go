package main

import (
	"crypto/md5"
	"fmt"
)

/*



 */

const (
	exitX = 3
	exitY = 3

	startX = 0
	startY = 0
)

func main() {
	fmt.Println("Advent of Code 2016 - Day 17")
	moves := []*Pos{&Pos{}}
	longest := 0
	i := 0
	for len(moves) > 0 {
		fmt.Println("ITERATION:", i)

		var newMoves []*Pos
		for _, p := range moves {
			for _, newP := range p.Moves() {
				if newP.x == exitX && newP.y == exitY {
					longest = i + 1
				} else {
					newMoves = append(newMoves, newP)
				}
			}
		}
		moves = newMoves
		i++
	}

	fmt.Println("Longest path:", longest)
}

type Pos struct {
	x, y  int
	moves string
}

func (p *Pos) Moves() (res []*Pos) {
	chk := fmt.Sprintf("%x", md5.Sum([]byte(puzzleInput+p.moves)))
	if p.y > 0 && isOpen(chk[0]) {
		// up
		res = append(res, &Pos{p.x, p.y - 1, p.moves + "U"})
	}
	if p.y < exitY && isOpen(chk[1]) {
		// down
		res = append(res, &Pos{p.x, p.y + 1, p.moves + "D"})
	}
	if p.x > 0 && isOpen(chk[2]) {
		// left
		res = append(res, &Pos{p.x - 1, p.y, p.moves + "L"})
	}
	if p.x < exitX && isOpen(chk[3]) {
		// right
		res = append(res, &Pos{p.x + 1, p.y, p.moves + "R"})
	}
	return
}

func isOpen(b byte) bool {
	switch b {
	case 'b', 'c', 'd', 'e', 'f':
		return true
	default:
		return false
	}
}

const testpuzzleInput = `kglvqrro` // `ulqzkmiv`
const puzzleInput = `ioramepc`
