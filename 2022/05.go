package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/pkg/errors"
)

func init() { days["5"] = Day5 }

/*
--- Day 5: Supply Stacks ---
The expedition can depart as soon as the final supplies have been unloaded from the ships. Supplies are stored in stacks of marked crates, but because the needed supplies are buried under many other crates, the crates need to be rearranged.

The ship has a giant cargo crane capable of moving crates between stacks. To ensure none of the crates get crushed or fall over, the crane operator will rearrange them in a series of carefully-planned steps. After the crates are rearranged, the desired crates will be at the top of each stack.

The Elves don't want to interrupt the crane operator during this delicate procedure, but they forgot to ask her which crate will end up where, and they want to be ready to unload them as soon as possible so they can embark.

They do, however, have a drawing of the starting stacks of crates and the rearrangement procedure (your puzzle input). For example:

    [D]
[N] [C]
[Z] [M] [P]
 1   2   3

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2
In this example, there are three stacks of crates. Stack 1 contains two crates: crate Z is on the bottom, and crate N is on top. Stack 2 contains three crates; from bottom to top, they are crates M, C, and D. Finally, stack 3 contains a single crate, P.

Then, the rearrangement procedure is given. In each step of the procedure, a quantity of crates is moved from one stack to a different stack. In the first step of the above rearrangement procedure, one crate is moved from stack 2 to stack 1, resulting in this configuration:

[D]
[N] [C]
[Z] [M] [P]
 1   2   3
In the second step, three crates are moved from stack 1 to stack 3. Crates are moved one at a time, so the first crate to be moved (D) ends up below the second and third crates:

        [Z]
        [N]
    [C] [D]
    [M] [P]
 1   2   3
Then, both crates are moved from stack 2 to stack 1. Again, because crates are moved one at a time, crate C ends up below crate M:

        [Z]
        [N]
[M]     [D]
[C]     [P]
 1   2   3
Finally, one crate is moved from stack 1 to stack 2:

        [Z]
        [N]
        [D]
[C] [M] [P]
 1   2   3
The Elves just need to know which crate will end up on top of each stack; in this example, the top crates are C in stack 1, M in stack 2, and Z in stack 3, so you should combine these together and give the Elves the message CMZ.

After the rearrangement procedure completes, what crate ends up on top of each stack?

Your puzzle answer was TLNGFGMFN.

--- Part Two ---
As you watch the crane operator expertly rearrange the crates, you notice the process isn't following your prediction.

Some mud was covering the writing on the side of the crane, and you quickly wipe it away. The crane isn't a CrateMover 9000 - it's a CrateMover 9001.

The CrateMover 9001 is notable for many new and exciting features: air conditioning, leather seats, an extra cup holder, and the ability to pick up and move multiple crates at once.

Again considering the example above, the crates begin in the same configuration:

    [D]
[N] [C]
[Z] [M] [P]
 1   2   3
Moving a single crate from stack 2 to stack 1 behaves the same as before:

[D]
[N] [C]
[Z] [M] [P]
 1   2   3
However, the action of moving three crates from stack 1 to stack 3 means that those three moved crates stay in the same order, resulting in this new configuration:

        [D]
        [N]
    [C] [Z]
    [M] [P]
 1   2   3
Next, as both crates are moved from stack 2 to stack 1, they retain their order as well:

        [D]
        [N]
[C]     [Z]
[M]     [P]
 1   2   3
Finally, a single crate is still moved from stack 1 to stack 2, but now it's crate C that gets moved:

        [D]
        [N]
        [Z]
[M] [C] [P]
 1   2   3
In this example, the CrateMover 9001 has put the crates in a totally different order: MCD.

Before the rearrangement process finishes, update your simulation so that the Elves know where they should stand to be ready to unload the final supplies. After the rearrangement procedure completes, what crate ends up on top of each stack?

Your puzzle answer was FGLQJCMBD.
*/

func Day5() {
	fmt.Println("--- Day 5: Supply Stacks ---")
	stacks := NewStacks(day05InputStacks)
	if err := stacks.MoveContainers(day05InputMoves, false); err != nil {
		fmt.Println(errors.Wrap(err, "Failed to move containers with CrateMover 9000"))
		return
	}
	top, err := stacks.TopContainers()
	if err != nil {
		fmt.Println(errors.Wrap(err, "Failed to get top containers after moving wiht CrateMover 9000"))
		return
	}
	fmt.Println("After moving containers with CrateMover 9000, the top containers are:", top)
	stacks = NewStacks(day05InputStacks)
	if err := stacks.MoveContainers(day05InputMoves, true); err != nil {
		fmt.Println(errors.Wrap(err, "Failed to move containers with CrateMover 9001"))
		return
	}
	top, err = stacks.TopContainers()
	if err != nil {
		fmt.Println(errors.Wrap(err, "Failed to get top containers after moving with CrateMover 9001"))
		return
	}
	fmt.Println("After moving containers with CrateMover 9001, the top containers are:", top)
}

// NewStacks takes stack input (only the initial container stacks - see `day05InputStacks`) and
// return the corresponding Stacks.
func NewStacks(in string) Stacks {
	lines := strings.Split(in, "\n")
	if len(lines) <= 0 {
		return nil
	}

	// use first line to identify the number of stacks knowing each stack takes up
	// 3 characers + 1 space (eg. "[X] "). Last stack is missing the space thus needs
	// to be added separately.
	stackCount := (len(lines[0]) / 4) + 1
	stacks := make(Stacks, stackCount)

	// go through the lines from the bottom, ignoring the first line containing numbers
	for n := len(lines) - 2; n >= 0; n-- {
		for i := 0; i < len(stacks); i++ {
			// the item in the stack is located at index 1, if any
			item := lines[n][i*4+1]
			if item != ' ' {
				stacks[i].Push(item)
			}
		}
	}

	return stacks
}

type Stacks []Stack

// MoveContainers takes a program (only the program - see `day05InputMoves`) and a boolean
// to indicate whether to lift containers one by one or whole stacks at a time. It executes
// the program and returns.
// An error is returned if any errors are encountered during the program.
func (s Stacks) MoveContainers(program string, liftStacks bool) error {
	moveExp := regexp.MustCompile(`^move (\d+) from (\d+) to (\d+)$`)
	for _, line := range strings.Split(program, "\n") {
		n := moveExp.FindStringSubmatch(line)
		if len(n) != 4 {
			return errors.Errorf("program line did not look as expected (line=%s)", line)
		}
		amount := ToIntOrPanic(n[1])
		from := ToIntOrPanic(n[2])
		to := ToIntOrPanic(n[3])
		if err := s.Move(liftStacks, amount, from, to); err != nil {
			return errors.Wrapf(err, "progam failed to move container (line=%s)", line)
		}
	}
	return nil
}

// Move takes a boolean to indicate whether to move whole stacks or containers one by one
// along with an amount, a from and a to and moves the containers accordingly.
// Note that from and to are 1-indexed with the first stack corresponding to 1.
// An error is returned if the move cannot be executed.
func (s Stacks) Move(liftStacks bool, amount, from, to int) error {
	// from and to are 1-indexed, correct to 0-index
	from--
	to--

	// from and to are 1-indexed
	if from < 0 || from >= len(s) {
		return errors.Errorf("from is not a valid stack (from=%d, stacks=%d)", from+1, len(s))
	}
	if to < 0 || to >= len(s) {
		return errors.Errorf("to is not a valid stack (to=%d, stacks=%d)", to+1, len(s))
	}

	containersPerMove := 1
	if liftStacks {
		containersPerMove = amount
	}
	moved := 0
	for moved < amount {
		c, err := s[from].Pop(containersPerMove)
		if err != nil {
			return errors.Wrapf(err, "not enough containers left to move from stack (amount=%d, from=%d, to=%d)", containersPerMove, from+1, to+1)
		}
		s[to].Push(c...)
		moved += containersPerMove
	}
	return nil
}

// TopContainers return a string representing the top container from each stack.
// An error is returned if any of the stacks are empty.
func (s Stacks) TopContainers() (string, error) {
	var res string
	for i := 0; i < len(s); i++ {
		c, err := s[i].Pop(1)
		if err != nil {
			return "", errors.Wrapf(err, "no top container left in stack (stack=%d)", i)
		}
		res += string(c[0])
	}
	return res, nil
}

// Print prints the stacks (horizontally with the bottom container leftmost) to stdout.
func (s Stacks) Print() {
	for n, stack := range s {
		fmt.Print(n+1, ": ")
		for i := 0; i < len(stack); i++ {
			fmt.Print(string(stack[i]))
		}
		fmt.Println()
	}
	fmt.Println()
}

type Stack []byte

// Push takes a list of one or more items push it on top of the stack in the order
// they were given.
func (s *Stack) Push(c ...byte) {
	(*s) = append(*s, c...)
}

// Pop removes the top number of items from the stack and returns them.
// Pop will return an error if more items are removed than exist in the stack.
func (s *Stack) Pop(count int) ([]byte, error) {
	if count > len(*s) {
		return nil, errors.Errorf("unable to pop items from stack (count=%d, stacksize=%d)", count, len(*s))
	}
	c := (*s)[len(*s)-count:]
	(*s) = (*s)[:len(*s)-count]
	return c, nil
}
