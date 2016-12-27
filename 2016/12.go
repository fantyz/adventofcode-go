package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

/*

--- Day 12: Leonardo's Monorail ---

You finally reach the top floor of this building: a garden with a slanted glass ceiling. Looks like there are no more stars to be had.

While sitting on a nearby bench amidst some tiger lilies, you manage to decrypt some of the files you extracted from the servers downstairs.

According to these documents, Easter Bunny HQ isn't just this building - it's a collection of buildings in the nearby area. They're all connected by a local monorail, and there's another building not far from here! Unfortunately, being night, the monorail is currently not operating.

You remotely connect to the monorail control systems and discover that the boot sequence expects a password. The password-checking logic (your puzzle input) is easy to extract, but the code it uses is strange: it's assembunny code designed for the new computer you just assembled. You'll have to execute the code and get the password.

The assembunny code you've extracted operates on four registers (a, b, c, and d) that start at 0 and can hold any integer. However, it seems to make use of only a few instructions:

cpy x y copies x (either an integer or the value of a register) into register y.
inc x increases the value of register x by one.
dec x decreases the value of register x by one.
jnz x y jumps to an instruction y away (positive means forward; negative means backward), but only if x is not zero.
The jnz instruction moves relative to itself: an offset of -1 would continue at the previous instruction, while an offset of 2 would skip over the next instruction.

For example:

cpy 41 a
inc a
inc a
dec a
jnz a 2
dec a
The above code would set register a to 41, increase its value by 2, decrease its value by 1, and then skip the last dec a (because a is not zero, so the jnz a 2 skips it), leaving register a at 42. When you move past the last instruction, the program halts.

After executing the assembunny code in your puzzle input, what value is left in register a?

Your puzzle answer was 318083.

The first half of this puzzle is complete! It provides one gold star: *

--- Part Two ---

As you head down the fire escape to the monorail, you notice it didn't start; register c needs to be initialized to the position of the ignition key.

If you instead initialize register c to be 1, what value is now left in register a

*/

func main() {
	fmt.Println("Advent of Code 2016 - Day 12")

	cmd := regexp.MustCompile(`^(cpy|inc|dec|jnz) ((-)?\d+|a|b|c|d)( ((-)?\d+|a|b|c|d))?$`)
	for _, line := range strings.Split(puzzleInput, "\n") {
		m := cmd.FindStringSubmatch(line)

		if len(m) <= 0 {
			panic("Unknown line: " + line)
		}
		c := m[1]
		x := m[2]
		y := m[5]

		if _, found := registers[x]; !found {
			NewRegistry(x)
		}
		if _, found := registers[y]; !found {
			NewRegistry(y)
		}

		var inst Instruction
		switch c {
		case "cpy":
			inst = NewCpy(x, y)
		case "inc":
			inst = NewInc(x)
		case "dec":
			inst = NewDec(x)
		case "jnz":
			inst = NewJnz(x, y)
		default:
			panic("unknown command: " + c)
		}
		instructions = append(instructions, inst)
	}

	pc := 0
	for pc < len(instructions) {
		pc = instructions[pc](pc)
	}

	fmt.Println("Register a contains", registers["a"])
}

var registers = map[string]int{}
var instructions = []Instruction{}

func NewRegistry(x string) {
	if x == "" {
		return
	}

	v, err := strconv.Atoi(x)
	if err != nil {
		v = 0
	}

	// part 2
	if x == "c" {
		v = 1
	}

	registers[x] = v
}

type Instruction func(int) int

func NewCpy(x, y string) Instruction {
	return func(pc int) int {
		registers[y] = registers[x]
		return pc + 1
	}
}

func NewInc(x string) Instruction {
	return func(pc int) int {
		registers[x]++
		return pc + 1
	}
}

func NewDec(x string) Instruction {
	return func(pc int) int {
		registers[x]--
		return pc + 1
	}
}

func NewJnz(x, y string) Instruction {
	return func(pc int) int {
		if registers[x] != 0 {
			return pc + registers[y]
		}
		return pc + 1
	}
}

const testpuzzleInput = `cpy 41 a
inc a
inc a
dec a
jnz a 2
dec a`
const puzzleInput = `cpy 1 a
cpy 1 b
cpy 26 d
jnz c 2
jnz 1 5
cpy 7 c
inc d
dec c
jnz c -2
cpy a c
inc a
dec b
jnz b -2
cpy c b
dec d
jnz d -6
cpy 16 c
cpy 17 d
inc a
dec d
jnz d -2
dec c
jnz c -5`
