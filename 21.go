package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

/*

--- Day 21: Scrambled Letters and Hash ---

The computer system you're breaking into uses a weird scrambling function to store its passwords. It shouldn't be much trouble to create your own scrambled password so you can add it to the system; you just have to implement the scrambler.

The scrambling function is a series of operations (the exact list is provided in your puzzle input). Starting with the password to be scrambled, apply each operation in succession to the string. The individual operations behave as follows:

swap position X with position Y means that the letters at indexes X and Y (counting from 0) should be swapped.
swap letter X with letter Y means that the letters X and Y should be swapped (regardless of where they appear in the string).
rotate left/right X steps means that the whole string should be rotated; for example, one right rotation would turn abcd into dabc.
rotate based on position of letter X means that the whole string should be rotated to the right based on the index of letter X (counting from 0) as determined before this instruction does any rotations. Once the index is determined, rotate the string to the right one time, plus a number of times equal to that index, plus one additional time if the index was at least 4.
reverse positions X through Y means that the span of letters at indexes X through Y (including the letters at X and Y) should be reversed in order.
move position X to position Y means that the letter which is at index X should be removed from the string, then inserted such that it ends up at index Y.
For example, suppose you start with abcde and perform the following operations:

swap position 4 with position 0 swaps the first and last letters, producing the input for the next step, ebcda.
swap letter d with letter b swaps the positions of d and b: edcba.
reverse positions 0 through 4 causes the entire string to be reversed, producing abcde.
rotate left 1 step shifts all letters left one position, causing the first letter to wrap to the end of the string: bcdea.
move position 1 to position 4 removes the letter at position 1 (c), then inserts it at position 4 (the end of the string): bdeac.
move position 3 to position 0 removes the letter at position 3 (a), then inserts it at position 0 (the front of the string): abdec.
rotate based on position of letter b finds the index of letter b (1), then rotates the string right once plus a number of times equal to that index (2): ecabd.
rotate based on position of letter d finds the index of letter d (4), then rotates the string right once, plus a number of times equal to that index, plus an additional time because the index was at least 4, for a total of 6 right rotations: decab.
After these steps, the resulting scrambled password is decab.

Now, you just need to generate a new scrambled password and you can access the system. Given the list of scrambling operations in your puzzle input, what is the result of scrambling abcdefgh?

Your puzzle answer was ghfacdbe.

The first half of this puzzle is complete! It provides one gold star: *

--- Part Two ---

You scrambled the password correctly, but you discover that you can't actually modify the password file on the system. You'll need to un-scramble one of the existing passwords by reversing the scrambling process.

What is the un-scrambled version of the scrambled password fbgdceah?

*/

func main() {
	fmt.Println("Advent of code 2016: Day 21")

	code := "abcdefgh"
	code = "abcde"
	fmt.Println("Part 1: Scrambling " + code + "...")
	for _, line := range strings.Split(testpuzzleInput, "\n") {
		code = NewOp(line).Execute(code)
		fmt.Println(code)
	}

	//code = "fbgdceah"
	fmt.Println("Part 2: Descrambling " + code + "...")

	lines := strings.Split(testpuzzleInput, "\n")
	for i := len(lines) - 1; i >= 0; i-- {
		code = NewOp(lines[i]).Reverse().Execute(code)
		fmt.Println(code)
	}
}

type Op interface {
	Execute(code string) string
	Reverse() Op
}

var (
	swapPosExp    = regexp.MustCompile(`^swap position (\d+) with position (\d+)$`)
	swapLetterExp = regexp.MustCompile(`^swap letter ([a-z]) with letter ([a-z])$`)
	rotateExp     = regexp.MustCompile(`^rotate (left|right) (\d+) steps?$`)
	rotateRelExp  = regexp.MustCompile(`^rotate based on position of letter ([a-z])$`)
	reverseExp    = regexp.MustCompile(`^reverse positions (\d+) through (\d+)$`)
	moveExp       = regexp.MustCompile(`^move position (\d+) to position (\d+)$`)
)

func NewOp(raw string) (o Op) {
	switch {
	case swapPosExp.MatchString(raw):
		m := swapPosExp.FindStringSubmatch(raw)
		p1, err := strconv.Atoi(m[1])
		if err != nil {
			panic(err)
		}
		p2, err := strconv.Atoi(m[2])
		if err != nil {
			panic(err)
		}
		o = NewSwapPosOp(p1, p2)
	case swapLetterExp.MatchString(raw):
		m := swapLetterExp.FindStringSubmatch(raw)
		o = NewSwapLetterOp(m[1], m[2])
	case rotateExp.MatchString(raw):
		m := rotateExp.FindStringSubmatch(raw)
		steps, err := strconv.Atoi(m[2])
		if err != nil {
			panic(err)
		}
		o = NewRotateOp(m[1], steps)
	case rotateRelExp.MatchString(raw):
		m := rotateRelExp.FindStringSubmatch(raw)
		o = NewRotateRelOp(m[1])
	case reverseExp.MatchString(raw):
		m := reverseExp.FindStringSubmatch(raw)
		p1, err := strconv.Atoi(m[1])
		if err != nil {
			panic(err)
		}
		p2, err := strconv.Atoi(m[2])
		if err != nil {
			panic(err)
		}
		o = NewReverseOp(p1, p2)
	case moveExp.MatchString(raw):
		m := moveExp.FindStringSubmatch(raw)
		p1, err := strconv.Atoi(m[1])
		if err != nil {
			panic(err)
		}
		p2, err := strconv.Atoi(m[2])
		if err != nil {
			panic(err)
		}
		o = NewMoveOp(p1, p2)
	default:
		panic("Unknown operation: " + raw)
	}
	return
}

func NewSwapPosOp(p1, p2 int) *SwapPosOp {
	return &SwapPosOp{
		p1: p1,
		p2: p2,
	}
}

type SwapPosOp struct {
	p1, p2 int
}

func (o *SwapPosOp) Reverse() Op {
	return o
}

func (o *SwapPosOp) Execute(code string) string {
	l1 := code[o.p1]
	l2 := code[o.p2]

	code = code[:o.p1] + string(l2) + code[o.p1+1:]
	code = code[:o.p2] + string(l1) + code[o.p2+1:]

	return code
}

func NewSwapLetterOp(l1, l2 string) *SwapLetterOp {
	return &SwapLetterOp{
		l1: l1,
		l2: l2,
	}
}

type SwapLetterOp struct {
	l1, l2 string
}

func (o *SwapLetterOp) Execute(code string) string {
	p1 := strings.Index(code, o.l1)
	p2 := strings.Index(code, o.l2)

	code = code[:p1] + o.l2 + code[p1+1:]
	code = code[:p2] + o.l1 + code[p2+1:]

	return code
}

func (o *SwapLetterOp) Reverse() Op {
	return o
}

func NewRotateOp(direction string, steps int) *RotateOp {
	return &RotateOp{
		direction: direction,
		steps:     steps,
	}
}

type RotateOp struct {
	direction string
	steps     int
}

func (o *RotateOp) Execute(code string) string {
	switch o.direction {
	case "left":
		return code[o.steps:] + code[:o.steps]
	case "right":
		return code[len(code)-o.steps:] + code[:len(code)-o.steps]
	default:
		panic("unknown direction: " + o.direction)
	}
}

func (o *RotateOp) Reverse() Op {
	if o.direction == "left" {
		o.direction = "right"
	} else {
		o.direction = "left"
	}
	return o
}

func NewRotateRelOp(letter string) *RotateRelOp {
	return &RotateRelOp{
		letter:    letter,
		direction: "right",
	}
}

type RotateRelOp struct {
	letter    string
	direction string
}

func (o *RotateRelOp) Execute(code string) string {
	i := strings.Index(code, o.letter)
	if i < 0 {
		panic("Letter not found!")
	}
	if i >= 4 {
		i++
	}
	i += 1
	for i >= len(code) {
		i -= len(code)
	}

	rotateOp := NewRotateOp(o.direction, i)

	return rotateOp.Execute(code)
}

func (o *RotateRelOp) Reverse() Op {
	panic("non trivial")
	return o
}

func NewReverseOp(p1, p2 int) *ReverseOp {
	return &ReverseOp{
		p1: p1,
		p2: p2,
	}
}

type ReverseOp struct {
	p1, p2 int
}

func (o *ReverseOp) Execute(code string) string {
	rev := ""
	for i := o.p2; i >= o.p1; i-- {
		rev += string(code[i])
	}

	return code[:o.p1] + rev + code[o.p2+1:]
}

func (o *ReverseOp) Reverse() Op {
	return o
}

func NewMoveOp(p1, p2 int) *MoveOp {
	return &MoveOp{
		p1: p1,
		p2: p2,
	}
}

type MoveOp struct {
	p1, p2 int
}

func (o *MoveOp) Execute(code string) string {
	l := string(code[o.p1])
	code = code[:o.p1] + code[o.p1+1:]
	code = code[:o.p2] + l + code[o.p2:]
	return code
}

func (o *MoveOp) Reverse() Op {
	panic("non trivial")
	return o
}

const testpuzzleInput = `swap position 4 with position 0
swap letter d with letter b
reverse positions 0 through 4
rotate left 1 step
move position 1 to position 4
move position 3 to position 0
rotate based on position of letter b
rotate based on position of letter d`

const puzzleInput = `rotate based on position of letter d
move position 1 to position 6
swap position 3 with position 6
rotate based on position of letter c
swap position 0 with position 1
rotate right 5 steps
rotate left 3 steps
rotate based on position of letter b
swap position 0 with position 2
rotate based on position of letter g
rotate left 0 steps
reverse positions 0 through 3
rotate based on position of letter a
rotate based on position of letter h
rotate based on position of letter a
rotate based on position of letter g
rotate left 5 steps
move position 3 to position 7
rotate right 5 steps
rotate based on position of letter f
rotate right 7 steps
rotate based on position of letter a
rotate right 6 steps
rotate based on position of letter a
swap letter c with letter f
reverse positions 2 through 6
rotate left 1 step
reverse positions 3 through 5
rotate based on position of letter f
swap position 6 with position 5
swap letter h with letter e
move position 1 to position 3
swap letter c with letter h
reverse positions 4 through 7
swap letter f with letter h
rotate based on position of letter f
rotate based on position of letter g
reverse positions 3 through 4
rotate left 7 steps
swap letter h with letter a
rotate based on position of letter e
rotate based on position of letter f
rotate based on position of letter g
move position 5 to position 0
rotate based on position of letter c
reverse positions 3 through 6
rotate right 4 steps
move position 1 to position 2
reverse positions 3 through 6
swap letter g with letter a
rotate based on position of letter d
rotate based on position of letter a
swap position 0 with position 7
rotate left 7 steps
rotate right 2 steps
rotate right 6 steps
rotate based on position of letter b
rotate right 2 steps
swap position 7 with position 4
rotate left 4 steps
rotate left 3 steps
swap position 2 with position 7
move position 5 to position 4
rotate right 3 steps
rotate based on position of letter g
move position 1 to position 2
swap position 7 with position 0
move position 4 to position 6
move position 3 to position 0
rotate based on position of letter f
swap letter g with letter d
swap position 1 with position 5
reverse positions 0 through 2
swap position 7 with position 3
rotate based on position of letter g
swap letter c with letter a
rotate based on position of letter g
reverse positions 3 through 5
move position 6 to position 3
swap letter b with letter e
reverse positions 5 through 6
move position 6 to position 7
swap letter a with letter e
swap position 6 with position 2
move position 4 to position 5
rotate left 5 steps
swap letter a with letter d
swap letter e with letter g
swap position 3 with position 7
reverse positions 0 through 5
swap position 5 with position 7
swap position 1 with position 7
swap position 1 with position 7
rotate right 7 steps
swap letter f with letter a
reverse positions 0 through 7
rotate based on position of letter d
reverse positions 2 through 4
swap position 7 with position 1
swap letter a with letter h`
