package main

/*

 */

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Advent of Code 2017 - Day 18")
	fmt.Println("Recovered frequency:", NewProgram(puzzle).RunUntilRcv())
}

func NewProgram(in string) *Program {
	p := &Program{
		regs: make(Registers),
	}

	cmdMatcher := regexp.MustCompile(`^(snd|set|add|mul|mod|rcv|jgz) ([a-z]|-?\d+)( ([a-z]|-?\d+))?$`)
	for _, line := range strings.Split(in, "\n") {
		m := cmdMatcher.FindStringSubmatch(line)
		if len(m) <= 0 {
			panic("line did not match regexp: " + line)
		}

		switch m[1] {
		case "snd":
			p.inst = append(p.inst, &Snd{x: p.regs.Get(m[2]), snd: p.regs.Get("snd")})
		case "set":
			p.inst = append(p.inst, &Set{x: p.regs.Get(m[2]), y: p.regs.GetRegOrConst(m[4])})
		case "add":
			p.inst = append(p.inst, &Add{x: p.regs.Get(m[2]), y: p.regs.GetRegOrConst(m[4])})
		case "mul":
			p.inst = append(p.inst, &Mul{x: p.regs.Get(m[2]), y: p.regs.GetRegOrConst(m[4])})
		case "mod":
			p.inst = append(p.inst, &Mod{x: p.regs.Get(m[2]), y: p.regs.GetRegOrConst(m[4])})
		case "rcv":
			p.inst = append(p.inst, &Rcv{x: p.regs.Get(m[2]), rcv: p.regs.Get("rcv"), snd: p.regs.Get("snd")})
		case "jgz":
			p.inst = append(p.inst, &Jgz{x: p.regs.GetRegOrConst(m[2]), y: p.regs.GetRegOrConst(m[4])})
		default:
			panic("unknown op: " + m[1])
		}
	}

	return p
}

type Program struct {
	regs Registers
	inst []Instruction
}

func (p *Program) RunUntilRcv() int {
	pc := 0
	for {
		/*
			fmt.Printf("[PC=%d] Registers:\n", pc)
			for k, v := range p.regs {
				fmt.Printf("  %s = %d\n", k, *v)
			}
		*/

		_, isRcv := p.inst[pc].(*Rcv)

		pc += p.inst[pc].Exec()
		if isRcv && p.Register("rcv") != 0 {
			return p.Register("rcv")
		}
	}
}

func (p *Program) Register(id string) int {
	r := p.regs.Get(id)
	return *r
}

type Registers map[string]*int

func (r Registers) Get(reg string) *int {
	v := r[reg]
	if v == nil {
		x := 0
		v = &x
		r[reg] = &x
	}
	return v
}

func (r Registers) GetRegOrConst(v string) *int {
	i, err := strconv.Atoi(v)
	if err != nil {
		// assume it is a register
		return r.Get(v)
	}
	// constant
	return &i
}

type Instruction interface {
	Exec() int
}

type Snd struct {
	x, snd *int
}

func (op *Snd) Exec() int {
	*op.snd = *op.x
	return 1
}

type Set struct {
	x, y *int
}

func (op *Set) Exec() int {
	*op.x = *op.y
	return 1
}

type Add struct {
	x, y *int
}

func (op *Add) Exec() int {
	*op.x = *op.x + *op.y
	return 1
}

type Mul struct {
	x, y *int
}

func (op *Mul) Exec() int {
	*op.x = *op.x * *op.y
	return 1
}

type Mod struct {
	x, y *int
}

func (op *Mod) Exec() int {
	*op.x = *op.x % *op.y
	return 1
}

type Rcv struct {
	x, snd, rcv *int
}

func (op *Rcv) Exec() int {
	if *op.x != 0 {
		*op.rcv = *op.snd
	}
	return 1
}

type Jgz struct {
	x, y *int
}

func (op *Jgz) Exec() int {
	if *op.x > 0 {
		return *op.y
	}
	return 1
}

const puzzle = `set i 31
set a 1
mul p 17
jgz p p
mul a 2
add i -1
jgz i -2
add a -1
set i 127
set p 618
mul p 8505
mod p a
mul p 129749
add p 12345
mod p a
set b p
mod b 10000
snd b
add i -1
jgz i -9
jgz a 3
rcv b
jgz b -1
set f 0
set i 126
rcv a
rcv b
set p a
mul p -1
add p b
jgz p 4
snd a
set a b
jgz 1 3
snd b
set f 1
add i -1
jgz i -11
snd a
jgz f -16
jgz a -19`
