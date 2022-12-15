package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewStacks(t *testing.T) {
	in := `    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 `
	stacks := NewStacks(in)
	if assert.Equal(t, 3, len(stacks), "# of stack in stacks") {
		assert.Equal(t, Stack{'Z', 'N'}, stacks[0], "stack 1")
		assert.Equal(t, Stack{'M', 'C', 'D'}, stacks[1], "stack 2")
		assert.Equal(t, Stack{'P'}, stacks[2], "stack 3")
	}
}

func TestMoveContainers(t *testing.T) {
	in := `    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 `

	prog := `move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2`

	stacks := NewStacks(in)
	if assert.NoError(t, stacks.MoveContainers(prog, false), "move containers") {
		top, err := stacks.TopContainers()
		if assert.NoError(t, err, "top containers") {
			assert.Equal(t, "CMZ", top)
		}
	}

}

func TestDay05Pt1(t *testing.T) {
	stacks := NewStacks(day05InputStacks)
	if assert.NoError(t, stacks.MoveContainers(day05InputMoves, false)) {
		top, err := stacks.TopContainers()
		if assert.NoError(t, err) {
			assert.Equal(t, "TLNGFGMFN", top)
		}
	}
}

func TestDay05Pt2(t *testing.T) {
	stacks := NewStacks(day05InputStacks)
	if assert.NoError(t, stacks.MoveContainers(day05InputMoves, true)) {
		top, err := stacks.TopContainers()
		if assert.NoError(t, err) {
			assert.Equal(t, "FGLQJCMBD", top)
		}
	}
}
