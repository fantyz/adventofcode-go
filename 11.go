package main

import (
	"fmt"
)

/*

--- Day 11: Radioisotope Thermoelectric Generators ---

You come upon a column of four floors that have been entirely sealed off from the rest of the building except for a small dedicated lobby. There are some radiation warnings and a big sign which reads "Radioisotope Testing Facility".

According to the project status board, this facility is currently being used to experiment with Radioisotope Thermoelectric Generators (RTGs, or simply "generators") that are designed to be paired with specially-constructed microchips. Basically, an RTG is a highly radioactive rock that generates electricity through heat.

The experimental RTGs have poor radiation containment, so they're dangerously radioactive. The chips are prototypes and don't have normal radiation shielding, but they do have the ability to generate an elecromagnetic radiation shield when powered. Unfortunately, they can only be powered by their corresponding RTG. An RTG powering a microchip is still dangerous to other microchips.

In other words, if a chip is ever left in the same area as another RTG, and it's not connected to its own RTG, the chip will be fried. Therefore, it is assumed that you will follow procedure and keep chips connected to their corresponding RTG when they're in the same room, and away from other RTGs otherwise.

These microchips sound very interesting and useful to your current activities, and you'd like to try to retrieve them. The fourth floor of the facility has an assembling machine which can make a self-contained, shielded computer for you to take with you - that is, if you can bring it all of the RTGs and microchips.

Within the radiation-shielded part of the facility (in which it's safe to have these pre-assembly RTGs), there is an elevator that can move between the four floors. Its capacity rating means it can carry at most yourself and two RTGs or microchips in any combination. (They're rigged to some heavy diagnostic equipment - the assembling machine will detach it for you.) As a security measure, the elevator will only function if it contains at least one RTG or microchip. The elevator always stops on each floor to recharge, and this takes long enough that the items within it and the items on that floor can irradiate each other. (You can prevent this if a Microchip and its Generator end up on the same floor in this way, as they can be connected while the elevator is recharging.)

You make some notes of the locations of each component of interest (your puzzle input). Before you don a hazmat suit and start moving things around, you'd like to have an idea of what you need to do.

When you enter the containment area, you and the elevator will start on the first floor.

For example, suppose the isolated area has the following arrangement:

The first floor contains a hydrogen-compatible microchip and a lithium-compatible microchip.
The second floor contains a hydrogen generator.
The third floor contains a lithium generator.
The fourth floor contains nothing relevant.
As a diagram (F# for a Floor number, E for Elevator, H for Hydrogen, L for Lithium, M for Microchip, and G for Generator), the initial state looks like this:

F4 .  .  .  .  .
F3 .  .  .  LG .
F2 .  HG .  .  .
F1 E  .  HM .  LM
Then, to get everything up to the assembling machine on the fourth floor, the following steps could be taken:

Bring the Hydrogen-compatible Microchip to the second floor, which is safe because it can get power from the Hydrogen Generator:

F4 .  .  .  .  .
F3 .  .  .  LG .
F2 E  HG HM .  .
F1 .  .  .  .  LM
Bring both Hydrogen-related items to the third floor, which is safe because the Hydrogen-compatible microchip is getting power from its generator:

F4 .  .  .  .  .
F3 E  HG HM LG .
F2 .  .  .  .  .
F1 .  .  .  .  LM
Leave the Hydrogen Generator on floor three, but bring the Hydrogen-compatible Microchip back down with you so you can still use the elevator:

F4 .  .  .  .  .
F3 .  HG .  LG .
F2 E  .  HM .  .
F1 .  .  .  .  LM
At the first floor, grab the Lithium-compatible Microchip, which is safe because Microchips don't affect each other:

F4 .  .  .  .  .
F3 .  HG .  LG .
F2 .  .  .  .  .
F1 E  .  HM .  LM
Bring both Microchips up one floor, where there is nothing to fry them:

F4 .  .  .  .  .
F3 .  HG .  LG .
F2 E  .  HM .  LM
F1 .  .  .  .  .
Bring both Microchips up again to floor three, where they can be temporarily connected to their corresponding generators while the elevator recharges, preventing either of them from being fried:

F4 .  .  .  .  .
F3 E  HG HM LG LM
F2 .  .  .  .  .
F1 .  .  .  .  .
Bring both Microchips to the fourth floor:

F4 E  .  HM .  LM
F3 .  HG .  LG .
F2 .  .  .  .  .
F1 .  .  .  .  .
Leave the Lithium-compatible microchip on the fourth floor, but bring the Hydrogen-compatible one so you can still use the elevator; this is safe because although the Lithium Generator is on the destination floor, you can connect Hydrogen-compatible microchip to the Hydrogen Generator there:

F4 .  .  .  .  LM
F3 E  HG HM LG .
F2 .  .  .  .  .
F1 .  .  .  .  .
Bring both Generators up to the fourth floor, which is safe because you can connect the Lithium-compatible Microchip to the Lithium Generator upon arrival:

F4 E  HG .  LG LM
F3 .  .  HM .  .
F2 .  .  .  .  .
F1 .  .  .  .  .
Bring the Lithium Microchip with you to the third floor so you can use the elevator:

F4 .  HG .  LG .
F3 E  .  HM .  LM
F2 .  .  .  .  .
F1 .  .  .  .  .
Bring both Microchips to the fourth floor:

F4 E  HG HM LG LM
F3 .  .  .  .  .
F2 .  .  .  .  .
F1 .  .  .  .  .
In this arrangement, it takes 11 steps to collect all of the objects at the fourth floor for assembly. (Each elevator stop counts as one step, even if nothing is added to or removed from it.)

In your situation, what is the minimum number of steps required to bring all of the objects to the fourth floor?

Your puzzle answer was 47.

The first half of this puzzle is complete! It provides one gold star: *

--- Part Two ---

You step into the cleanroom separating the lobby from the isolated area and put on the hazmat suit.

Upon entering the isolated containment area, however, you notice some extra parts on the first floor that weren't listed on the record outside:

An elerium generator.
An elerium-compatible microchip.
A dilithium generator.
A dilithium-compatible microchip.
These work just like the other generators and microchips. You'll have to get them up to assembly as well.

What is the minimum number of steps required to bring all of the objects, including these four new ones, to the fourth floor?

Your puzzle answer was 71.

*/

/*

The solution does not run with 14 elements.

Instead I arrived at 71 by looking at 8, 10 and 12.

8 is done in 35 steps
10 is done in 47 steps
12 is done in 59 steps

35 -> 47 is 12 steps
47 -> 59 is 12 steps

Ergo it takes 12 steps for every 2 extra elements.

59 + 12 = 71.

Win.

*/

//const elementCount = 4
//const elementCount = 10
const elementCount = 12
const topFloor = 3

var endBoard = &Board{
	//elevator: 1,
	//elements: [elementCount]int{2, 1, 2, 0},

	elevator: 3,
	//elements: [elementCount]int{topFloor, topFloor, topFloor, topFloor}, // 4
	//elements: [elementCount]int{topFloor, topFloor, topFloor, topFloor, topFloor, topFloor, topFloor, topFloor}, // 8
	//elements: [elementCount]int{topFloor, topFloor, topFloor, topFloor, topFloor, topFloor, topFloor, topFloor, topFloor, topFloor}, // 10
	elements: [elementCount]int{topFloor, topFloor, topFloor, topFloor, topFloor, topFloor, topFloor, topFloor, topFloor, topFloor, topFloor, topFloor}, // 12
	//elements: [elementCount]int{topFloor, topFloor, topFloor, topFloor, topFloor, topFloor, topFloor, topFloor, topFloor, topFloor, topFloor, topFloor, topFloor, topFloor},  // 14
}
var endBoardChecksum = endBoard.Checksum()

var seenBoards = map[[elementCount + 1]int]*Board{}
var incomplete = []*Board{
	&Board{
		elevator: 0,
		//elements: [elementCount]int{1, 0, 2, 0},  // 4
		//elements: [elementCount]int{0, 1, 0, 0, 0, 1}, // 8
		//elements: [elementCount]int{0, 1, 0, 0, 0, 1, 0, 0}, // 10
		elements: [elementCount]int{0, 1, 0, 0, 0, 1, 0, 0, 0, 0}, // 12
		//elements: [elementCount]int{0, 1, 0, 0, 0, 1, 0, 0, 0, 0, 0,  0}, // 14
	},
}
var maxDepth = 0

func main() {
	fmt.Println("Advent of Code 2016 - Day 11")

	for {
		boards := make([]*Board, len(incomplete))
		copy(boards, incomplete)
		incomplete = incomplete[0:0]
		for i := range boards {
			calcAllMoves(boards[i])
		}
		n := seenBoards[endBoardChecksum]
		if n != nil {
			n.Print()
			fmt.Println("Did it in", n.Depth(), "moves")
			fmt.Println("")
			fmt.Println("Seen boards:", len(seenBoards))
			break
		}

		fmt.Println("Increasing maxDepth to", maxDepth)
		maxDepth++
	}
}

type Board struct {
	parent   *Board
	children []*Board
	depth    int
	elevator int
	elements [elementCount]int
}

func (b *Board) Child() *Board {
	c := &Board{
		parent:   b,
		depth:    b.depth + 1,
		elevator: b.elevator,
	}
	for i := 0; i < len(b.elements); i++ {
		c.elements[i] = b.elements[i]
	}
	return c
}

func (b *Board) SwitchParent(p *Board) {
	b.parent = p
	b.RecalculateDepth()
	for _, c := range b.children {
		c.ResetDepth()
	}
}

func (b *Board) ResetDepth() {
	b.depth = b.parent.depth + 1
}

func (b *Board) RecalculateDepth() {
	p := b.parent
	b.depth = 0
	for p != nil {
		p = p.parent
		b.depth++
	}
}

func (b *Board) Checksum() [elementCount + 1]int {
	v := [elementCount + 1]int{}
	for i := 0; i < elementCount; i++ {
		v[i] = b.elements[i]
	}
	v[elementCount] = b.elevator
	return v
}

func (b *Board) Depth() int {
	if b == nil || b.parent == nil {
		return 0
	}
	return b.parent.Depth() + 1
}

func (b *Board) Print() {
	if b == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println("DEPTH:", b.Depth())
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

func calcAllMoves(b *Board) {
	if b.depth >= maxDepth {
		incomplete = append(incomplete, b)
		return
	}

	if b.Checksum() == endBoardChecksum {
		maxDepth = b.depth
		fmt.Println("END DEPTH:", b.Depth())
	}

	if existingBoard, found := seenBoards[b.Checksum()]; found {
		if existingBoard.depth > b.depth {
			// replace existingBoard's parent with with b's parent
			existingBoard.SwitchParent(b.parent)
		} else {
			return
		}
	}

	seenBoards[b.Checksum()] = b

	/*
		b.Print()
		var s string
		_, _ = fmt.Scanln(&s)
	*/

	// move all possible sets one up/down
	for i := 0; i < len(b.elements); i++ {
		for j := 0; j < len(b.elements); j++ {
			if i == j || b.elements[i] != b.elevator || b.elements[j] != b.elevator {
				continue
			}
			upBoard := b.Child()
			upBoard.elements[i]++
			upBoard.elements[j]++
			upBoard.elevator++
			if isValid(upBoard) {
				calcAllMoves(upBoard)
			}

			downBoard := b.Child()
			downBoard.elements[i]--
			downBoard.elements[j]--
			downBoard.elevator--
			if isValid(downBoard) {
				calcAllMoves(downBoard)
			}
		}
	}

	// move all single pieces one up/down
	for i := 0; i < len(b.elements); i++ {
		if b.elevator != b.elements[i] {
			// only move elements on the same level as the elevator
			continue
		}

		upBoard := b.Child()
		upBoard.elements[i]++
		upBoard.elevator++
		if isValid(upBoard) {
			calcAllMoves(upBoard)
		}

		downBoard := b.Child()
		downBoard.elements[i]--
		downBoard.elevator--
		if isValid(downBoard) {
			calcAllMoves(downBoard)
		}
	}

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

const puzzleInput = `The first floor contains a polonium generator, a thulium generator, a thulium-compatible microchip, a promethium generator, a ruthenium generator, a ruthenium-compatible microchip, a cobalt generator, and a cobalt-compatible microchip.
The second floor contains a polonium-compatible microchip and a promethium-compatible microchip.
The third floor contains nothing relevant.
The fourth floor contains nothing relevant.`
