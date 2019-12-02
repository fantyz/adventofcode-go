package main

import (
	"fmt"
	"strconv"
	"strings"
)

func ExecuteOpcode(input []int) []int {
	program := make([]int, len(input))
	copy(program, input)

	i := 0
	for {
		switch program[i] {
		case 1:
			program[program[i+3]] = program[program[i+1]] + program[program[i+2]]
			i += 4
		case 2:
			program[program[i+3]] = program[program[i+1]] * program[program[i+2]]
			i += 4
		case 99:
			return program
		default:
			panic(fmt.Sprintf("unknown opcode: %d", program[i]))
		}
	}
}

func Load(in string) []int {
	var out []int
	for _, l := range strings.Split(in, ",") {
		i, err := strconv.Atoi(l)
		if err != nil {
			panic(err)
		}
		out = append(out, i)
	}
	return out
}
