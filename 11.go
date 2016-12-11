package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Advent of Code 2016 - Day 11")
	initialBoard := &Board{
		elevator: 0,
		elements: [8]int{2, 2, 1, 1, 3, 3, 0, 0},
		//elements: [10]int{0, 1, 0, 0, 0, 1, 0, 0, 0, 0},
	}
	resultBoard := &Board{
		elevator: 0,
		elements: [8]int{topFloor, topFloor, topFloor, topFloor, topFloor, topFloor, topFloor, topFloor},
		//elements: [10]int{topFloor, topFloor, topFloor, topFloor, topFloor, topFloor, topFloor, topFloor, topFloor, topFloor},
	}
	root := &Node{}

	evalAllMoves(0, root, initialBoard)
	fmt.Println("Seen boards:", len(seenBoards))
	time.Sleep(time.Nanosecond)

	n := seenBoards[resultBoard.elements]
	fmt.Println("Did it in", n.depth, "moves")
	resultBoard.Print()
}

var s string

func evalAllMoves(move int, parent *Node, board *Board) {
	node, seen := SeenAlready(move, parent, board)
	if seen {
		return
	}

	fmt.Println("======== MOVE:", move)
	board.Print()
	//_, _ = fmt.Scanln(&s)
	time.Sleep(300 * time.Millisecond)

	// move all possible sets one up/down
	for i := 0; i < len(board.elements); i++ {
		for j := 0; j < len(board.elements); j++ {
			if i == j || board.elements[i] != board.elevator || board.elements[j] != board.elevator {
				continue
			}
			upBoard := board.Copy()
			upBoard.elements[i]++
			upBoard.elements[j]++
			upBoard.elevator++
			if isValid(upBoard) {
				evalAllMoves(move+1, node, upBoard)
			}

			downBoard := board.Copy()
			downBoard.elements[i]--
			downBoard.elements[j]--
			downBoard.elevator--
			if isValid(downBoard) {
				evalAllMoves(move+1, node, downBoard)
			}
		}
	}

	// move all single pieces one up/down
	for i := 0; i < len(board.elements); i++ {
		if board.elevator != board.elements[i] {
			// only move elements on the same level as the elevator
			continue
		}

		upBoard := board.Copy()
		upBoard.elements[i]++
		upBoard.elevator++
		if isValid(upBoard) {
			evalAllMoves(move+1, node, upBoard)
		}

		downBoard := board.Copy()
		downBoard.elements[i]--
		downBoard.elevator--
		if isValid(downBoard) {
			evalAllMoves(move+1, node, downBoard)
		}
	}

}

const topFloor = 3

type Board struct {
	elevator int
	elements [8]int
}

func (b *Board) Copy() *Board {
	new := &Board{
		elevator: b.elevator,
	}
	for i := 0; i < len(b.elements); i++ {
		new.elements[i] = b.elements[i]
	}
	return new
}

func (b *Board) Print() {
	for i := topFloor; i >= 0; i-- {
		fmt.Printf("Floor %d: ", i+1)
		if i == b.elevator {
			fmt.Print("E ")
		} else {
			fmt.Print("  ")
		}

		for j := 0; j < len(b.elements); j++ {
			if b.elements[j] == i {
				if j%2 == 0 {
					fmt.Print("G")
				} else {
					fmt.Print("M")
				}
			} else {
				fmt.Print(".")
			}
		}
		fmt.Print("\n")
	}
	fmt.Print("\n")
}

func (b *Board) Value() int {
	sum := 0
	for _, v := range b.elements {
		sum += v
	}
	return sum
}

func isValid(b *Board) bool {
	if b.elevator > topFloor || b.elevator < 0 {
		return false
	}
	gOnFloor := make([]bool, topFloor+1)
	mOnFloor := make([]bool, topFloor+1)

	for i := 0; i < len(b.elements)-1; i += 2 {
		gOnFloor[b.elements[i]] = true
		if b.elements[i] != b.elements[i+1] {
			mOnFloor[b.elements[i+1]] = true
		}
	}
	for i := 0; i < topFloor+1; i++ {
		if gOnFloor[i] && mOnFloor[i] {
			return false
		}
	}
	return true
}

var seenBoards = map[[8]int]*Node{}
var boardTree *Node

type Node struct {
	depth    int
	parent   *Node
	children []*Node
}

func (n *Node) SetDepth(d int) {
	if d >= n.depth {
		return
	}

	n.depth = d
	for i := range n.children {
		n.children[i].SetDepth(d + 1)
	}
}

func SeenAlready(m int, p *Node, b *Board) (*Node, bool) {
	if n, found := seenBoards[b.elements]; found {
		if m < n.depth {
			n.SetDepth(m)
		}
		return n, true
	}

	n := &Node{
		parent: p,
		depth:  m,
	}
	p.children = append(p.children, n)

	seenBoards[b.elements] = n

	return n, false
}
