package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

/*

--- Day 25: Clock Signal ---

You open the door and find yourself on the roof. The city sprawls away from you for miles and miles.

There's not much time now - it's already Christmas, but you're nowhere near the North Pole, much too far to deliver these stars to the sleigh in time.

However, maybe the huge antenna up here can offer a solution. After all, the sleigh doesn't need the stars, exactly; it needs the timing data they provide, and you happen to have a massive signal generator right here.

You connect the stars you have to your prototype computer, connect that to the antenna, and begin the transmission.

Nothing happens.

You call the service number printed on the side of the antenna and quickly explain the situation. "I'm not sure what kind of equipment you have connected over there," he says, "but you need a clock signal." You try to explain that this is a signal for a clock.

"No, no, a clock signal - timing information so the antenna computer knows how to read the data you're sending it. An endless, alternating pattern of 0, 1, 0, 1, 0, 1, 0, 1, 0, 1...." He trails off.

You ask if the antenna can handle a clock signal at the frequency you would need to use for the data from the stars. "There's no way it can! The only antenna we've installed capable of that is on top of a top-secret Easter Bunny installation, and you're definitely not-" You hang up the phone.

You've extracted the antenna's clock signal generation assembunny code (your puzzle input); it looks mostly compatible with code you worked on just recently.

This antenna code, being a signal generator, uses one extra instruction:

out x transmits x (either an integer or the value of a register) as the next value for the clock signal.
The code takes a value (via register a) that describes the signal to generate, but you're not sure how it's used. You'll have to find the input to produce the right signal through experimentation.

What is the lowest positive integer that can be used to initialize register a and cause the code to output a clock signal of 0, 1, 0, 1... repeating forever?

Your puzzle answer was 189.

--- Part Two ---

The antenna is ready. Now, all you need is the fifty stars required to generate the signal for the sleigh, but you don't have enough.

You look toward the sky in desperation... suddenly noticing that a lone star has been installed at the top of the antenna! Only 49 more to go.

If you like, you can [Retransmit the Signal].

*/

func main() {
	fmt.Println("Advent of Code 2016 - Day 25")

	i := 0
outer:
	for {
		i++

		registers = map[string]int{}
		instructions = []Instruction{}
		instInfos = []InstInfo{}

		cmd := regexp.MustCompile(`^(cpy|inc|dec|jnz|out) ((-)?\d+|a|b|c|d)( ((-)?\d+|a|b|c|d))?$`)
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
			case "out":
				instInfo.t = OutInstType
				instInfo.regs = []string{x}
				inst = NewOut(x)
			default:
				panic("unknown command: " + c)
			}
			instInfos = append(instInfos, instInfo)
			instructions = append(instructions, inst)
		}

		fmt.Printf("Initializing a with %d: ", i)
		registers["a"] = i

		pc := 0
		expects := 0
		count := 0
		for pc < len(instructions) {
			pc = instructions[pc](pc)
			if registers["out"] == 1 {
				fmt.Print(registers["b"])
				if registers["b"] != expects {
					fmt.Println()
					continue outer
				}
				count++
				if count > 50 {
					// done
					fmt.Println()
					break outer
				}

				expects = (expects + 1) % 2
				registers["out"] = 0
			}
		}
	}

	fmt.Println("Part 1: Initializing a to", i, "gave the desired output")
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
	OutInstType
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

func NewOut(x string) Instruction {
	return func(pc int) int {
		registers["out"] = 1
		return pc + 1
	}
}

const puzzleInput = `cpy a d
cpy 11 c
cpy 231 b
inc d
dec b
jnz b -2
dec c
jnz c -5
cpy d a
jnz 0 0
cpy a b
cpy 0 a
cpy 2 c
jnz b 2
jnz 1 6
dec b
dec c
jnz c -4
inc a
jnz 1 -7
cpy 2 b
jnz c 2
jnz 1 4
dec b
dec c
jnz 1 -4
jnz 0 0
out b
jnz a -19
jnz 1 -21`
