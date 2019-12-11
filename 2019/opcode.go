package main

import (
	"fmt"
	"strconv"
	"strings"
)

func ExecuteOpcode(program []int, input <-chan int) <-chan int {
	out := make(chan int)

	p := make([]int, len(program))
	copy(p, program)

	go func() {
		relativeBase := 0
		i := 0
		for {
			op := opcode(p[i])
			switch op[0] {
			case 1:
				regs := getRegs(&p, i, op, relativeBase, 3)
				*regs[2] = *regs[0] + *regs[1]
				i += 4
			case 2:
				regs := getRegs(&p, i, op, relativeBase, 3)
				*regs[2] = *regs[0] * *regs[1]
				i += 4
			case 3:
				// input
				regs := getRegs(&p, i, op, relativeBase, 1)
				*regs[0] = <-input
				i += 2
			case 4:
				// output
				regs := getRegs(&p, i, op, relativeBase, 1)
				out <- *regs[0]
				i += 2
			case 5:
				// jump if true
				regs := getRegs(&p, i, op, relativeBase, 2)
				if *regs[0] != 0 {
					i = *regs[1]
					continue
				}
				i += 3
			case 6:
				// jump if false
				regs := getRegs(&p, i, op, relativeBase, 2)
				if *regs[0] == 0 {
					i = *regs[1]
					continue
				}
				i += 3
			case 7:
				// less than
				regs := getRegs(&p, i, op, relativeBase, 3)
				if *regs[0] < *regs[1] {
					*regs[2] = 1
				} else {
					*regs[2] = 0
				}
				i += 4
			case 8:
				// equals
				regs := getRegs(&p, i, op, relativeBase, 3)
				if *regs[0] == *regs[1] {
					*regs[2] = 1
				} else {
					*regs[2] = 0
				}
				i += 4
			case 9:
				// adjust the relative base
				regs := getRegs(&p, i, op, relativeBase, 1)
				relativeBase += *regs[0]
				i += 2
			case 99:
				close(out)
				return
			default:
				panic(fmt.Sprintf("unknown opcode: %d (%v)", op[0], op))
			}
		}
	}()

	return out
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

func Outputter(c <-chan int) []int {
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

func getRegs(p *[]int, pos int, op [4]int, relativeBase int, regcount int) [3]*int {
	out := [3]*int{}

	for i := 1; i <= regcount; i++ {
		switch op[i] {
		case 0:
			// position mode (same as relative with base 0)
			relativeBase = 0
			fallthrough
		case 2:
			// relative mode
			reg := getReg(p, pos+i)
			out[i-1] = getReg(p, *reg+relativeBase)
		case 1:
			// immidate mode
			out[i-1] = getReg(p, pos+i)
		default:
			panic(fmt.Sprintf("unknown parameter mode: %d (op=%v)", op[i], op))
		}
	}

	return out
}

func getReg(p *[]int, pos int) *int {
	if len(*p) <= pos {
		newP := make([]int, pos+1)
		copy(newP, *p)
		*p = newP
	}
	return &(*p)[pos]
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
