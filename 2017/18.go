package main

/*

--- Day 18: Duet ---
You discover a tablet containing some strange assembly code labeled simply "Duet". Rather than bother the sound card with it, you decide to run the code yourself. Unfortunately, you don't see any documentation, so you're left to figure out what the instructions mean on your own.

It seems like the assembly is meant to operate on a set of registers that are each named with a single letter and that can each hold a single integer. You suppose each register should start with a value of 0.

There aren't that many instructions, so it shouldn't be hard to figure out what they do. Here's what you determine:

snd X plays a sound with a frequency equal to the value of X.
set X Y sets register X to the value of Y.
add X Y increases register X by the value of Y.
mul X Y sets register X to the result of multiplying the value contained in register X by the value of Y.
mod X Y sets register X to the remainder of dividing the value contained in register X by the value of Y (that is, it sets X to the result of X modulo Y).
rcv X recovers the frequency of the last sound played, but only when the value of X is not zero. (If it is zero, the command does nothing.)
jgz X Y jumps with an offset of the value of Y, but only if the value of X is greater than zero. (An offset of 2 skips the next instruction, an offset of -1 jumps to the previous instruction, and so on.)
Many of the instructions can take either a register (a single letter) or a number. The value of a register is the integer it contains; the value of a number is that number.

After each jump instruction, the program continues with the instruction to which the jump jumped. After any other instruction, the program continues with the next instruction. Continuing (or jumping) off either end of the program terminates it.

For example:

set a 1
add a 2
mul a a
mod a 5
snd a
set a 0
rcv a
jgz a -1
set a 1
jgz a -2
The first four instructions set a to 1, add 2 to it, square it, and then set it to itself modulo 5, resulting in a value of 4.
Then, a sound with frequency 4 (the value of a) is played.
After that, a is set to 0, causing the subsequent rcv and jgz instructions to both be skipped (rcv because a is 0, and jgz because a is not greater than 0).
Finally, a is set to 1, causing the next jgz instruction to activate, jumping back two instructions to another jump, which jumps again to the rcv, which ultimately triggers the recover operation.
At the time the recover operation is executed, the frequency of the last sound played is 4.

What is the value of the recovered frequency (the value of the most recently played sound) the first time a rcv instruction is executed with a non-zero value?

Your puzzle answer was 3423.

--- Part Two ---
As you congratulate yourself for a job well done, you notice that the documentation has been on the back of the tablet this entire time. While you actually got most of the instructions correct, there are a few key differences. This assembly code isn't about sound at all - it's meant to be run twice at the same time.

Each running copy of the program has its own set of registers and follows the code independently - in fact, the programs don't even necessarily run at the same speed. To coordinate, they use the send (snd) and receive (rcv) instructions:

snd X sends the value of X to the other program. These values wait in a queue until that program is ready to receive them. Each program has its own message queue, so a program can never receive a message it sent.
rcv X receives the next value and stores it in register X. If no values are in the queue, the program waits for a value to be sent to it. Programs do not continue to the next instruction until they have received a value. Values are received in the order they are sent.
Each program also has its own program ID (one 0 and the other 1); the register p should begin with this value.

For example:

snd 1
snd 2
snd p
rcv a
rcv b
rcv c
rcv d
Both programs begin by sending three values to the other. Program 0 sends 1, 2, 0; program 1 sends 1, 2, 1. Then, each program receives a value (both 1) and stores it in a, receives another value (both 2) and stores it in b, and then each receives the program ID of the other program (program 0 receives 1; program 1 receives 0) and stores it in c. Each program now sees a different value in its own copy of register c.

Finally, both programs try to rcv a fourth time, but no data is waiting for either of them, and they reach a deadlock. When this happens, both programs terminate.

It should be noted that it would be equally valid for the programs to run at different speeds; for example, program 0 might have sent all three values and then stopped at the first rcv before program 1 executed even its first instruction.

Once both of your programs have terminated (regardless of what caused them to do so), how many times did program 1 send a value?

Your puzzle answer was 7493.

*/

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
)

func main() {
	fmt.Println("Advent of Code 2017 - Day 18")
	fmt.Println("Recovered frequency:", NewProgram(puzzle).RunUntilRcv())
	fmt.Println("Program 1 sends:", RunInParallel(puzzle))
}

func RunInParallel(in string) int {
	p0p1 := make(chan int, 1000)
	p1p0 := make(chan int, 1000)

	p0 := NewProgram2(0, p1p0, p0p1, in)
	p1 := NewProgram2(1, p0p1, p1p0, in)

	wg := &sync.WaitGroup{}
	wg.Add(2)
	go func() {
		p0.Run()
		wg.Done()
	}()
	go func() {
		p1.Run()
		wg.Done()
	}()
	wg.Wait()

	return p1.Register("count")
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

func NewProgram2(id int, input <-chan int, output chan<- int, in string) *Program {
	p := &Program{
		id:   id,
		regs: make(Registers),
	}
	p.regs["p"] = &id

	cmdMatcher := regexp.MustCompile(`^(snd|set|add|mul|mod|rcv|jgz) ([a-z]|-?\d+)( ([a-z]|-?\d+))?$`)
	for _, line := range strings.Split(in, "\n") {
		m := cmdMatcher.FindStringSubmatch(line)
		if len(m) <= 0 {
			panic("line did not match regexp: " + line)
		}

		switch m[1] {
		case "snd":
			p.inst = append(p.inst, &Snd2{x: p.regs.GetRegOrConst(m[2]), count: p.regs.Get("count"), ch: output})
		case "set":
			p.inst = append(p.inst, &Set{x: p.regs.Get(m[2]), y: p.regs.GetRegOrConst(m[4])})
		case "add":
			p.inst = append(p.inst, &Add{x: p.regs.Get(m[2]), y: p.regs.GetRegOrConst(m[4])})
		case "mul":
			p.inst = append(p.inst, &Mul{x: p.regs.Get(m[2]), y: p.regs.GetRegOrConst(m[4])})
		case "mod":
			p.inst = append(p.inst, &Mod{x: p.regs.Get(m[2]), y: p.regs.GetRegOrConst(m[4])})
		case "rcv":
			p.inst = append(p.inst, &Rcv2{x: p.regs.Get(m[2]), ch: input})
		case "jgz":
			p.inst = append(p.inst, &Jgz{x: p.regs.GetRegOrConst(m[2]), y: p.regs.GetRegOrConst(m[4])})
		default:
			panic("unknown op: " + m[1])
		}
	}

	return p
}

type Program struct {
	id   int
	regs Registers
	inst []Instruction
}

func (p *Program) Run() {
	pc := 0
	for {
		/*
			fmt.Printf("[ID=%d | PC=%d] Registers:\n", p.id, pc)
			for k, v := range p.regs {
				fmt.Printf("  %s = %d\n", k, *v)
			}
		*/
		pc += p.inst[pc].Exec()
		if pc >= len(p.inst) {
			break
		}
	}
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

type Snd2 struct {
	x, count *int
	ch       chan<- int
}

func (op *Snd2) Exec() int {
	op.ch <- *op.x
	*op.count++
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

type Rcv2 struct {
	x  *int
	ch <-chan int
}

func (op *Rcv2) Exec() int {
	select {
	case *op.x = <-op.ch:
		return 1
	case _ = <-time.After(1 * time.Second):
		return 9999999
	}
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
