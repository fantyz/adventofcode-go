package main

import (
	"fmt"
	"strconv"
	"strings"
)

func ExecuteOpcode(program []int, input <-chan int, output chan<- int) []int {
	p := make([]int, len(program))
	copy(p, program)

	i := 0
	for {
		op := opcode(p[i])
		regs := getRegs(p, i, op)

		switch op[0] {
		case 1:
			*regs[2] = *regs[0] + *regs[1]
			i += 4
		case 2:
			*regs[2] = *regs[0] * *regs[1]
			i += 4
		case 3:
			// input
			*regs[0] = <-input
			i += 2
		case 4:
			// output
			output <- *regs[0]
			i += 2
		case 5:
			// jump if true
			if *regs[0] != 0 {
				i = *regs[1]
				continue
			}
			i += 3
		case 6:
			// jump if false
			if *regs[0] == 0 {
				i = *regs[1]
				continue
			}
			i += 3
		case 7:
			// less than
			if *regs[0] < *regs[1] {
				*regs[2] = 1
			} else {
				*regs[2] = 0
			}
			i += 4
		case 8:
			// equals
			if *regs[0] == *regs[1] {
				*regs[2] = 1
			} else {
				*regs[2] = 0
			}
			i += 4
		case 99:
			close(output)
			return p
		default:
			panic(fmt.Sprintf("unknown opcode: %d (%v)", op[0], op))
		}
	}
}

func Inputter(in []int) <-chan int {
	c := make(chan int)
	go func() {
		for _, v := range in {
			c <- v
		}
		close(c)
	}()
	return c
}

func Outputter(c chan int) []int {
	var out []int
	for v := range c {
		out = append(out, v)
	}
	return out
}

func opcode(in int) [4]int {
	out := [4]int{}
	out[0], out[1], out[2], out[3] = in%100, (in/100)%10, (in/1000)%10, (in/10000)%10
	return out
}

func getRegs(p []int, pos int, op [4]int) [3]*int {
	out := [3]*int{}

	for i := 1; i <= 3; i++ {
		if pos+i >= len(p) {
			// outside of program, set to immidate mode with value 0
			v := 0
			out[i-1] = &v
			continue
		}

		switch op[i] {
		case 0:
			// position mode
			if p[pos+i] < 0 || p[pos+i] >= len(p) {
				// value outside of program, set to immidate mode with value 0
				v := 0
				out[i-1] = &v
				continue
			}
			out[i-1] = &p[p[pos+i]]
		case 1:
			// immidate mode
			v := p[pos+i]
			out[i-1] = &v
		default:
			panic(fmt.Sprintf("unknown parameter mode: %d (op=%v)", op[i], op))
		}
	}

	return out
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
