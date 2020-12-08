package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

func init() { days["8"] = Day8 }

/*

--- Day 8: Handheld Halting ---
Your flight to the major airline hub reaches cruising altitude without incident. While you consider checking the in-flight menu for one of those drinks that come with a little umbrella, you are interrupted by the kid sitting next to you.

Their handheld game console won't turn on! They ask if you can take a look.

You narrow the problem down to a strange infinite loop in the boot code (your puzzle input) of the device. You should be able to fix it, but first you need to be able to run the code in isolation.

The boot code is represented as a text file with one instruction per line of text. Each instruction consists of an operation (acc, jmp, or nop) and an argument (a signed number like +4 or -20).

acc increases or decreases a single global value called the accumulator by the value given in the argument. For example, acc +7 would increase the accumulator by 7. The accumulator starts at 0. After an acc instruction, the instruction immediately below it is executed next.
jmp jumps to a new instruction relative to itself. The next instruction to execute is found using the argument as an offset from the jmp instruction; for example, jmp +2 would skip the next instruction, jmp +1 would continue to the instruction immediately below it, and jmp -20 would cause the instruction 20 lines above to be executed next.
nop stands for No OPeration - it does nothing. The instruction immediately below it is executed next.
For example, consider the following program:

nop +0
acc +1
jmp +4
acc +3
jmp -3
acc -99
acc +1
jmp -4
acc +6
These instructions are visited in this order:

nop +0  | 1
acc +1  | 2, 8(!)
jmp +4  | 3
acc +3  | 6
jmp -3  | 7
acc -99 |
acc +1  | 4
jmp -4  | 5
acc +6  |
First, the nop +0 does nothing. Then, the accumulator is increased from 0 to 1 (acc +1) and jmp +4 sets the next instruction to the other acc +1 near the bottom. After it increases the accumulator from 1 to 2, jmp -4 executes, setting the next instruction to the only acc +3. It sets the accumulator to 5, and jmp -3 causes the program to continue back at the first acc +1.

This is an infinite loop: with this sequence of jumps, the program will run forever. The moment the program tries to run any instruction a second time, you know it will never terminate.

Immediately before the program would run an instruction a second time, the value in the accumulator is 5.

Run your copy of the boot code. Immediately before any instruction is executed a second time, what value is in the accumulator?

Your puzzle answer was 1317.

--- Part Two ---
After some careful analysis, you believe that exactly one instruction is corrupted.

Somewhere in the program, either a jmp is supposed to be a nop, or a nop is supposed to be a jmp. (No acc instructions were harmed in the corruption of this boot code.)

The program is supposed to terminate by attempting to execute an instruction immediately after the last instruction in the file. By changing exactly one jmp or nop, you can repair the boot code and make it terminate correctly.

For example, consider the same program from above:

nop +0
acc +1
jmp +4
acc +3
jmp -3
acc -99
acc +1
jmp -4
acc +6
If you change the first instruction from nop +0 to jmp +0, it would create a single-instruction infinite loop, never leaving that instruction. If you change almost any of the jmp instructions, the program will still eventually find another jmp instruction and loop forever.

However, if you change the second-to-last instruction (from jmp -4 to nop -4), the program terminates! The instructions are visited in this order:

nop +0  | 1
acc +1  | 2
jmp +4  | 3
acc +3  |
jmp -3  |
acc -99 |
acc +1  | 4
nop -4  | 5
acc +6  | 6
After the last instruction (acc +6), the program terminates by attempting to run the instruction below the last instruction in the file. With this change, after the program terminates, the accumulator contains the value 8 (acc +1, acc +1, acc +6).

Fix the program so that it terminates normally by changing exactly one jmp (to nop) or nop (to jmp). What is the value of the accumulator after the program terminates?

Your puzzle answer was 1033.

*/

func Day8() {
	fmt.Println("--- Day 8: Handheld Halting ---")
	e, err := NewBootCodeExecuter(day8Input, &HaltOnFirstRepetitionDebugger{})
	if err != nil {
		fmt.Println(errors.Wrap(err, "Failed to create executor"))
		return
	}
	fmt.Println("  Accumulator when an instruction would be run the second time:", e.Run())
	acc, err := FindWorkingBootCodeProgram(day8Input)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("  Accumulator after finding a working program:", acc)
}

// FindWorkingBootCodeProgram takes a program and modifies it by swapping Nop with Jmp and
// vice versa to find a working bootcode program that exists normally. It will return the
// value of the accumulator for the first working program it encounters. An error is returned
// if no working program is found.
func FindWorkingBootCodeProgram(program string) (int, error) {
	e, err := NewBootCodeExecuter(program, &HaltOnFirstRepetitionDebugger{})
	if err != nil {
		return 0, errors.Wrap(err, "Failed to create executor")
	}

	for i := range e.program {
		orgInst := e.program[i]

		switch inst := e.program[i].(type) {
		case JmpBootCode:
			e.program[i] = NopBootCode{N: inst.Offset}
		case NopBootCode:
			e.program[i] = JmpBootCode{Offset: inst.N}
		default:
			// nothing to replace, skip
			continue
		}

		acc := e.Run()
		if e.pc >= len(e.program) {
			// program exited normally, working program found!
			return acc, nil
		}

		// program looped, reset instruction and executor
		e.program[i] = orgInst
		e.Reset()
	}

	return 0, errors.New("no working program found")
}

// NewBootCodeExecuter takes a program and debugger and returns a BootCodeExecutor that
// is capable of running the bootcode. An error will be returned if the input program
// is invalid.
func NewBootCodeExecuter(program string, debugger Debugger) (*BootCodeExecuter, error) {
	instExp := regexp.MustCompile(`^([a-z]+) ([+\-0-9]+)$`)

	var instructions []BootCodeInst
	for _, line := range strings.Split(program, "\n") {
		m := instExp.FindStringSubmatch(line)
		if len(m) != 3 {
			return nil, errors.Errorf("line did not parse properly (line=%s)", line)
		}

		arg, err := strconv.Atoi(m[2])
		if err != nil {
			return nil, errors.Wrapf(err, "failed to read argument (line=%s)", line)
		}

		var inst BootCodeInst
		switch m[1] {
		case "acc":
			inst = AccBootCode{N: arg}
		case "jmp":
			inst = JmpBootCode{Offset: arg}
		case "nop":
			inst = NopBootCode{N: arg}
		default:
			return nil, errors.Errorf("unknown op (line=%s, op=%s)", line, m[1])
		}

		instructions = append(instructions, inst)
	}

	e := &BootCodeExecuter{program: instructions}
	if debugger != nil {
		debugger.Init(e)
		e.debugger = debugger
	}

	return e, nil
}

type BootCodeExecuter struct {
	pc          int
	accumulator int
	debugger    Debugger
	program     []BootCodeInst
}

// Run executes the the program.
func (e *BootCodeExecuter) Run() int {
	for e.pc >= 0 && e.pc < len(e.program) {
		if e.debugger != nil {
			if e.debugger.RunPreExec(e) {
				return e.accumulator
			}
		}
		e.pc = e.program[e.pc].Execute(e)
	}
	return e.accumulator
}

// Reset resets BootCodeExecutor to allow running it again.
func (e *BootCodeExecuter) Reset() {
	e.pc = 0
	e.accumulator = 0
	if e.debugger != nil {
		e.debugger.Init(e)
	}
}

// Instructions
//
// The program is represented in BootCodeExecutor as a slice of BootCodeInst instructions.
// Once the BootCodeExecutor needs to run a given instruction it calls its Execute function
// with the exeecutor as an argument and expects this function to return the index to the
// following instruction (usually the BootCodeExecutor pc value + 1).
type BootCodeInst interface {
	Execute(*BootCodeExecuter) int
}

type AccBootCode struct {
	N int
}

func (inst AccBootCode) Execute(e *BootCodeExecuter) int {
	e.accumulator += inst.N
	return e.pc + 1
}

type JmpBootCode struct {
	Offset int
}

func (inst JmpBootCode) Execute(e *BootCodeExecuter) int {
	return e.pc + inst.Offset
}

type NopBootCode struct {
	N int
}

func (_ NopBootCode) Execute(e *BootCodeExecuter) int {
	return e.pc + 1
}

// Debuggers
//
// Debuggers allow an easy way to interact with a running bootcode program and possibly modify
// its behavior.
//
type Debugger interface {
	// Init resets the debugger to a state where it is ready for teh program to run.
	Init(*BootCodeExecuter)
	// RunPreExec is called right before the BootCodeExecutor executes its next instruction.
	// The return value is used to determine whether or not to halt the execution of the
	// program. It halts if it returns true and continues normally otherwise.
	RunPreExec(*BootCodeExecuter) bool
}

// HaltOnFirstRepetitionDebugger keeps track of how many times the individual instructions have
// been executed. It terminates the execution of the program if any instruction is about to be
// executed a second time.
type HaltOnFirstRepetitionDebugger []int

func (debug *HaltOnFirstRepetitionDebugger) Init(e *BootCodeExecuter) {
	*debug = make([]int, len(e.program))
}

func (debug HaltOnFirstRepetitionDebugger) RunPreExec(e *BootCodeExecuter) bool {
	if e.pc >= 0 && e.pc < len(debug) {
		if debug[e.pc] >= 1 {
			// next instruction has already been run before, forece termination
			return true
		}

		// mark next instruction as executed
		debug[e.pc]++
	}
	return false
}
