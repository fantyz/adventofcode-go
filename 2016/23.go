package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

/*

--- Day 23: Safe Cracking ---

This is one of the top floors of the nicest tower in EBHQ. The Easter Bunny's private office is here, complete with a safe hidden behind a painting, and who wouldn't hide a star in a safe behind a painting?

The safe has a digital screen and keypad for code entry. A sticky note attached to the safe has a password hint on it: "eggs". The painting is of a large rabbit coloring some eggs. You see 7.

When you go to type the code, though, nothing appears on the display; instead, the keypad comes apart in your hands, apparently having been smashed. Behind it is some kind of socket - one that matches a connector in your prototype computer! You pull apart the smashed keypad and extract the logic circuit, plug it into your computer, and plug your computer into the safe.

Now, you just need to figure out what output the keypad would have sent to the safe. You extract the assembunny code from the logic chip (your puzzle input).
The code looks like it uses almost the same architecture and instruction set that the monorail computer used! You should be able to use the same assembunny interpreter for this as you did there, but with one new instruction:

tgl x toggles the instruction x away (pointing at instructions like jnz does: positive means forward; negative means backward):

For one-argument instructions, inc becomes dec, and all other one-argument instructions become inc.
For two-argument instructions, jnz becomes cpy, and all other two-instructions become jnz.
The arguments of a toggled instruction are not affected.
If an attempt is made to toggle an instruction outside the program, nothing happens.
If toggling produces an invalid instruction (like cpy 1 2) and an attempt is later made to execute that instruction, skip it instead.
If tgl toggles itself (for example, if a is 0, tgl a would target itself and become inc a), the resulting instruction is not executed until the next time it is reached.
For example, given this program:

cpy 2 a
tgl a
tgl a
tgl a
cpy 1 a
dec a
dec a
cpy 2 a initializes register a to 2.
The first tgl a toggles an instruction a (2) away from it, which changes the third tgl a into inc a.
The second tgl a also modifies an instruction 2 away from it, which changes the cpy 1 a into jnz 1 a.
The fourth line, which is now inc a, increments a to 3.
Finally, the fifth line, which is now jnz 1 a, jumps a (3) instructions ahead, skipping the dec a instructions.
In this example, the final value in register a is 3.

The rest of the electronics seem to place the keypad entry (the number of eggs, 7) in register a, run the code, and then send the value left in register a to the safe.

What value should be sent to the safe?

Your puzzle answer was 11610.

The first half of this puzzle is complete! It provides one gold star: *

--- Part Two ---

The safe doesn't open, but it does make several angry noises to express its frustration.

You're quite sure your logic is working correctly, so the only other thing is... you check the painting again. As it turns out, colored eggs are still eggs. Now you count 12.

As you run the program with this new input, the prototype computer begins to overheat. You wonder what's taking so long, and whether the lack of any instruction more powerful than "add one" has anything to do with it. Don't bunnies usually multiply?

Anyway, what value should actually be sent to the safe?

*/

func main() {
	fmt.Println("Advent of Code 2016 - Day 23")

	cmd := regexp.MustCompile(`^(cpy|inc|dec|jnz|tgl) ((-)?\d+|a|b|c|d)( ((-)?\d+|a|b|c|d))?$`)
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

		var instInfo InstInfo
		var inst Instruction
		switch c {
		case "cpy":
			instInfo.t = CpyInstType
			instInfo.regs = []string{x, y}
			inst = NewCpy(x, y)
		case "inc":
			instInfo.t = IncInstType
			instInfo.regs = []string{x}
			inst = NewInc(x)
		case "dec":
			instInfo.t = DecInstType
			instInfo.regs = []string{x}
			inst = NewDec(x)
		case "jnz":
			instInfo.t = JnzInstType
			instInfo.regs = []string{x, y}
			inst = NewJnz(x, y)
		case "tgl":
			instInfo.t = TglInstType
			instInfo.regs = []string{x}
			inst = NewTgl(x)
		default:
			panic("unknown command: " + c)
		}
		instInfos = append(instInfos, instInfo)
		instructions = append(instructions, inst)
	}

	// part 1
	//registers["a"] = 7

	// part 2
	registers["a"] = 12

	pc := 0
	for pc < len(instructions) {
		pc = instructions[pc](pc)
	}

	fmt.Println("Register a contains", registers["a"])
}

type InstInfo struct {
	t    InstType
	regs []string
}

type InstType uint8

const (
	CpyInstType InstType = iota
	IncInstType
	DecInstType
	JnzInstType
	TglInstType
)

var registers = map[string]int{}
var instructions = []Instruction{}
var instInfos = []InstInfo{}

func NewRegistry(x string) {
	v, err := strconv.Atoi(x)
	if err != nil {
		v = 0
	}

	registers[x] = v
}

// registers are defined as:
var isregisterexp = regexp.MustCompile(`^[a-z]$`)

func isRegister(x string) bool {
	return isregisterexp.MatchString(x)
}

type Instruction func(int) int

func NewCpy(x, y string) Instruction {
	return func(pc int) int {
		if !isRegister(y) {
			// invalid instruction becomes no-op
			fmt.Println(" > CPY (Invalid, NOOP)")
			return pc + 1
		}

		registers[y] = registers[x]
		//fmt.Println(" > CPY", x, y)

		return pc + 1
	}
}

func NewInc(x string) Instruction {
	return func(pc int) int {
		//fmt.Println(" > INC", x)
		registers[x]++
		return pc + 1
	}
}

func NewDec(x string) Instruction {
	return func(pc int) int {
		//fmt.Println(" > DEC", x)
		registers[x]--
		return pc + 1
	}
}

func NewJnz(x, y string) Instruction {
	return func(pc int) int {
		//fmt.Println(" > JNZ", x, y)
		if registers[x] != 0 {
			return pc + registers[y]
		}
		return pc + 1
	}
}

func NewTgl(x string) Instruction {
	return func(pc int) int {
		tgt := pc + registers[x]
		if tgt < 0 || tgt >= len(instructions) {
			//fmt.Println(" > TGL (Invalid, NOOP)")
			return pc + 1
		}

		//fmt.Println(" > TGL", x)
		info := instInfos[tgt]
		switch info.t {
		// one-argument instructions
		case IncInstType:
			//fmt.Println("Changing", tgt, "to Dec")
			instInfos[tgt].t = DecInstType
			instructions[tgt] = NewDec(info.regs[0])
		case DecInstType:
			fallthrough
		case TglInstType:
			//fmt.Println("Changing", tgt, "to Inc")
			instInfos[tgt].t = IncInstType
			instructions[tgt] = NewInc(info.regs[0])

			// two-argument instructions
		case JnzInstType:
			//fmt.Println("Changing", tgt, "to Cpy")
			instInfos[tgt].t = CpyInstType
			instructions[tgt] = NewCpy(info.regs[0], info.regs[1])
		case CpyInstType:
			//fmt.Println("Changing", tgt, "to Jnz")
			instInfos[tgt].t = JnzInstType
			instructions[tgt] = NewJnz(info.regs[0], info.regs[1])
		}
		return pc + 1
	}
}

const testpuzzleInput = `cpy 2 a
tgl a
tgl a
tgl a
cpy 1 a
dec a
dec a`

const puzzleInput = `cpy a b
dec b
cpy a d
cpy 0 a
cpy b c
inc a
dec c
jnz c -2
dec d
jnz d -5
dec b
cpy b c
cpy c d
dec d
inc c
jnz d -2
tgl c
cpy -16 c
jnz 1 c
cpy 90 c
jnz 73 d
inc a
inc d
jnz d -2
inc c
jnz c -5`
