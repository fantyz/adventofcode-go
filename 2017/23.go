package main

/*

--- Day 23: Coprocessor Conflagration ---
You decide to head directly to the CPU and fix the printer from there. As you get close, you find an experimental coprocessor doing so much work that the local programs are afraid it will halt and catch fire. This would cause serious issues for the rest of the computer, so you head in and see what you can do.

The code it's running seems to be a variant of the kind you saw recently on that tablet. The general functionality seems very similar, but some of the instructions are different:

set X Y sets register X to the value of Y.
sub X Y decreases register X by the value of Y.
mul X Y sets register X to the result of multiplying the value contained in register X by the value of Y.
jnz X Y jumps with an offset of the value of Y, but only if the value of X is not zero. (An offset of 2 skips the next instruction, an offset of -1 jumps to the previous instruction, and so on.)
Only the instructions listed above are used. The eight registers here, named a through h, all start at 0.

The coprocessor is currently set to some kind of debug mode, which allows for testing, but prevents it from doing any meaningful work.

If you run the program (your puzzle input), how many times is the mul instruction invoked?

Your puzzle answer was 3025.

--- Part Two ---
Now, it's time to fix the problem.

The debug mode switch is wired directly to register a. You flip the switch, which makes register a now start at 1 when the program is executed.

Immediately, the coprocessor begins to overheat. Whoever wrote this program obviously didn't choose a very efficient implementation. You'll need to optimize the program if it has any hope of completing before Santa needs that printer working.

The coprocessor's ultimate goal is to determine the final value left in register h once the program completes. Technically, if it had that... it wouldn't even need to run the program.

After setting register a to 1, if the program were to run to completion, what value would be left in register h?

Your puzzle answer was 915.

*/

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const debug = false

func main() {
	fmt.Println("Advent of Code 2017 - Day 23")
	p := NewProgram(puzzle)
	p.Run()
	fmt.Println("MUL instruction invokations:", *p.regs.Get("mulcount"))
	fmt.Println("Register f after setting a=1:", PuzzleProg())

	/*
			// Deciphered the puzzle to be the sum of non-primes between b and c in steps of 17
	        // this code does not work.
			b := 105700
			c := 122700
			primes := PrimesBetween(b, c)
			idx := 0
			sum := 0
			for i := b; i < c; i += 17 {
				if primes[idx] != i {
					sum++
				}
				for primes[idx] <= i {
					idx++
				}
			}
			fmt.Println("Sum of non-primes:", sum)
	*/
}

func PrimesBetween(n, m int) []int {
	primes := make([]int, 0)
	l := make([]bool, m-n)
	for i := 2; i < m; i++ {
		if i >= n && !l[i-n] {
			primes = append(primes, i)
		}
		for j := i; j < m; j += i {
			if j < n {
				continue
			}
			l[j-n] = true
		}
	}
	return primes
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

func NewProgram(in string) *Program {
	p := &Program{
		regs: make(Registers),
	}

	cmdMatcher := regexp.MustCompile(`^(set|sub|mul|jnz) ([a-z]|-?\d+)( ([a-z]|-?\d+))?$`)
	for _, line := range strings.Split(in, "\n") {
		m := cmdMatcher.FindStringSubmatch(line)
		if len(m) <= 0 {
			panic("line did not match regexp: " + line)
		}

		switch m[1] {
		case "set":
			p.inst = append(p.inst, &Set{x: p.regs.GetRegOrConst(m[2]), y: p.regs.GetRegOrConst(m[4])})
		case "sub":
			p.inst = append(p.inst, &Sub{x: p.regs.GetRegOrConst(m[2]), y: p.regs.GetRegOrConst(m[4])})
		case "mul":
			p.inst = append(p.inst, &Mul{x: p.regs.GetRegOrConst(m[2]), y: p.regs.GetRegOrConst(m[4]), count: p.regs.Get("mulcount")})
		case "jnz":
			p.inst = append(p.inst, &Jnz{x: p.regs.GetRegOrConst(m[2]), y: p.regs.GetRegOrConst(m[4])})
		default:
			panic("unknown op: " + m[1])
		}
	}

	return p
}

type Instruction interface {
	Exec() int
}

type Program struct {
	regs Registers
	inst []Instruction
}

func (p *Program) Run() {
	pc := 0
	for {
		pc += p.inst[pc].Exec()
		if pc >= len(p.inst) {
			break
		}
	}
}

type Set struct {
	x, y *int
}

func (op *Set) Exec() int {
	if debug {
		fmt.Println("SET", op.x, *op.y)
	}
	*op.x = *op.y
	return 1
}

type Sub struct {
	x, y *int
}

func (op *Sub) Exec() int {
	if debug {
		fmt.Println("SUB", op.x, *op.y)
	}
	*op.x -= *op.y
	return 1
}

type Mul struct {
	x, y  *int
	count *int
}

func (op *Mul) Exec() int {
	if debug {
		fmt.Println("MUL", op.x, *op.y, *op.count)
	}
	*op.count++
	*op.x *= *op.y
	return 1
}

type Jnz struct {
	x, y *int
}

func (op *Jnz) Exec() int {
	if debug {
		fmt.Println("JNZ", op.x, *op.x, *op.y)
	}
	if *op.x != 0 {
		return *op.y
	}
	return 1
}

func PuzzleProg() int {
	h := 0

	//b := 57 // set b 57
	//c := b  // set c b
	// jnz a 2, a == 1
	// jnz 1 5
	//b *= 100    // mul b 100
	//b += 100000 // sub b -100000
	//c = b       // set c b
	//c += 17000  // sub c -17000

	//fmt.Println("b", b)
	//fmt.Println("c", c)
	//return 0

	b := 105700
	c := 122700

	for {
		f := 1 // set f 1
		d := 2 // set d 2
	outer:
		for {
			e := 2 // set e 2
			for {
				//fmt.Printf("d=%d e=%d b=%d\n", d, e, b)
				if (d*e)-b == 0 { // set g d / mul g e / sub g b / jnz g 2
					f = 0 // set f 0
					break outer
				}
				e++           // sub e -1
				if e-b == 0 { // set g e + sub g b + jnz g -8
					break
				}
			}
			d++           // sub d -1
			if d-b == 0 { // set g d + sub g b + jnz g -13
				break
			}
		}

		fmt.Printf("b=%d c=%d d=%d f=%d h=%d\n", b, c, d, f, h)
		//return -1

		if f == 0 { // jnz f 2
			h++ // sub h -1
		}
		if b-c == 0 { // set g b + sub g c + jnz g 2
			return h // jnz 1 3
		}
		b += 17 // sub b -17
	} // jnz 1 -23

	return h
}

const puzzle = `set b 57
set c b
jnz a 2
jnz 1 5
mul b 100
sub b -100000
set c b
sub c -17000
set f 1
set d 2
set e 2
set g d
mul g e
sub g b
jnz g 2
set f 0
sub e -1
set g e
sub g b
jnz g -8
sub d -1
set g d
sub g b
jnz g -13
jnz f 2
sub h -1
set g b
sub g c
jnz g 2
jnz 1 3
sub b -17
jnz 1 -23`
