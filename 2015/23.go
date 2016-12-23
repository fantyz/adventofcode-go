package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

/*

 */

func main() {
	fmt.Println("Advent of Code 2015 - Day 23")

	p := NewProgram(0)
	for _, line := range strings.Split(puzzleInput, "\n") {
		p.AddInstruction(line)
	}
	p.Execute()
	fmt.Println("Part 1: Register a contains", *p.registers["b"])

	p = NewProgram(1)
	for _, line := range strings.Split(puzzleInput, "\n") {
		p.AddInstruction(line)
	}
	p.Execute()
	fmt.Println("Part 2: Register a contains", *p.registers["b"])

}

func NewProgram(a int) *Program {
	registers := make(map[string]*int)
	registers["a"] = &a
	b := 0
	registers["b"] = &b

	return &Program{
		registers: registers,
	}
}

type Program struct {
	instructions []Instruction
	registers    map[string]*int
}

var instExp = regexp.MustCompile(`^(hlf|tpl|inc|jmp|jie|jio) (a|b|[+-]\d+)(, ([+-]\d+))?$`)

func (p *Program) AddInstruction(line string) {
	m := instExp.FindStringSubmatch(line)
	if len(m) <= 0 {
		panic("Line didnt match: " + line)
	}
	var inst Instruction
	switch m[1] {
	case "hlf":
		inst = &Hlf{p.registers[m[2]]}
	case "tpl":
		inst = &Tpl{p.registers[m[2]]}
	case "inc":
		inst = &Inc{p.registers[m[2]]}
	case "jmp":
		offset, err := strconv.Atoi(m[2])
		if err != nil {
			panic("Invalid offset: " + err.Error())
		}
		inst = &Jmp{offset}
	case "jie":
		offset, err := strconv.Atoi(m[4])
		if err != nil {
			panic("Invalid offset: " + err.Error())
		}
		inst = &Jie{p.registers[m[2]], offset}
	case "jio":
		offset, err := strconv.Atoi(m[4])
		if err != nil {
			panic("Invalid offset: " + err.Error())
		}
		inst = &Jio{p.registers[m[2]], offset}
	default:
		panic("Unknown instruction: " + m[1])
	}
	p.instructions = append(p.instructions, inst)
}

func (p Program) Execute() {
	pc := 0
	for pc < len(p.instructions) {
		pc = p.instructions[pc].Execute(pc)
	}
}

type Instruction interface {
	Execute(int) int
}

type Hlf struct {
	r *int
}

func (i *Hlf) Execute(pc int) int {
	//fmt.Println(" > HLF", i.r, *i.r)
	*i.r = *i.r / 2
	return pc + 1
}

type Tpl struct {
	r *int
}

func (i *Tpl) Execute(pc int) int {
	//fmt.Println(" > TPL", i.r, *i.r)
	*i.r = *i.r * 3
	return pc + 1
}

type Inc struct {
	r *int
}

func (i *Inc) Execute(pc int) int {
	//fmt.Println(" > INC", i.r, *i.r)
	*i.r++
	return pc + 1
}

type Jmp struct {
	offset int
}

func (i *Jmp) Execute(pc int) int {
	//fmt.Println(" > JMP", i.offset)
	return pc + i.offset
}

type Jie struct {
	r      *int
	offset int
}

func (i *Jie) Execute(pc int) int {
	//fmt.Println(" > JIE", i.r, *i.r, i.offset)
	if *i.r%2 != 0 {
		return pc + 1
	}
	return pc + i.offset
}

type Jio struct {
	r      *int
	offset int
}

func (i *Jio) Execute(pc int) int {
	//fmt.Println(" > JIO", i.r, *i.r, i.offset)
	if *i.r != 1 {
		return pc + 1
	}
	return pc + i.offset
}

const puzzleInput = `jio a, +16
inc a
inc a
tpl a
tpl a
tpl a
inc a
inc a
tpl a
inc a
inc a
tpl a
tpl a
tpl a
inc a
jmp +23
tpl a
inc a
inc a
tpl a
inc a
inc a
tpl a
tpl a
inc a
inc a
tpl a
inc a
tpl a
inc a
tpl a
inc a
inc a
tpl a
inc a
tpl a
tpl a
inc a
jio a, +8
inc b
jie a, +4
tpl a
inc a
jmp +2
hlf a
jmp -7`
