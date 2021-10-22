package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

func init() { days["7"] = Day7 }

/*
--- Day 7: Some Assembly Required ---
This year, Santa brought little Bobby Tables a set of wires and bitwise logic gates! Unfortunately, little Bobby is a little under the recommended age range, and he needs help assembling the circuit.

Each wire has an identifier (some lowercase letters) and can carry a 16-bit signal (a number from 0 to 65535). A signal is provided to each wire by a gate, another wire, or some specific value. Each wire can only get a signal from one source, but can provide its signal to multiple destinations. A gate provides no signal until all of its inputs have a signal.

The included instructions booklet describes how to connect the parts together: x AND y -> z means to connect wires x and y to an AND gate, and then connect its output to wire z.

For example:

123 -> x means that the signal 123 is provided to wire x.
x AND y -> z means that the bitwise AND of wire x and wire y is provided to wire z.
p LSHIFT 2 -> q means that the value from wire p is left-shifted by 2 and then provided to wire q.
NOT e -> f means that the bitwise complement of the value from wire e is provided to wire f.
Other possible gates include OR (bitwise OR) and RSHIFT (right-shift). If, for some reason, you'd like to emulate the circuit instead, almost all programming languages (for example, C, JavaScript, or Python) provide operators for these gates.

For example, here is a simple circuit:

123 -> x
456 -> y
x AND y -> d
x OR y -> e
x LSHIFT 2 -> f
y RSHIFT 2 -> g
NOT x -> h
NOT y -> i
After it is run, these are the signals on the wires:

d: 72
e: 507
f: 492
g: 114
h: 65412
i: 65079
x: 123
y: 456
In little Bobby's kit's instructions booklet (provided as your puzzle input), what signal is ultimately provided to wire a?

Your puzzle answer was 46065.

--- Part Two ---
Now, take the signal you got on wire a, override wire b to that signal, and reset the other wires (including wire a). What new signal is ultimately provided to wire a?

Your puzzle answer was 14134.
*/

func Day7() {
	fmt.Println("--- Day 7: Some Assembly Required ---")
	c, err := NewCircuit(day7Input)
	if err != nil {
		fmt.Println(errors.Wrap(err, "unable to load input"))
		return
	}
	a, err := c.ResolveWire("a")
	if err != nil {
		fmt.Println(errors.Wrap(err, "failed to resolve wire a"))
		return
	}
	fmt.Println("Signal on wire a:", int(a))

	c, err = NewCircuit(day7Input)
	if err != nil {
		fmt.Println(errors.Wrap(err, "unable to load input"))
		return
	}
	c.Override("b", a)
	a2, err := c.ResolveWire("a")
	if err != nil {
		fmt.Println(errors.Wrap(err, "failed to resolve wire a"))
		return
	}
	fmt.Printf("Signal on wire a after overriding b to %d: %d\n", int(a), int(a2))
}

// NewCircuit takes a list of circuit instructions and returns a Circuit
func NewCircuit(instructions string) (*Circuit, error) {
	c := Circuit{
		wireToGate: map[string]Gate{},
		wireValues: map[string]uint16{},
	}

	// instExp matches a general instruction into a gate expression and a wire label
	instExp := regexp.MustCompile(`^([0-9a-zA-Z ]+) -> ([a-z]+)$`)
	// gateExp matches a gate expression
	gateExp := regexp.MustCompile(`^([0-9a-z]+)$|^([0-9a-z]+) (AND|OR) ([0-9a-z]+)$|^([0-9a-z]+) (LSHIFT|RSHIFT) ([0-9]+)$|^NOT ([0-9a-z]+)$`)

	// constToWire is a helper function that takes a consant, inserts an extra InputGate and returns its input wire.
	// If a wire is passed (not-a-constant-number) the input wire is returned as is.
	constToWire := func(wire string) string {
		if v, err := strconv.ParseUint(wire, 10, 16); err == nil {
			c.wireToGate[wire] = InputGate(v)
		}
		return wire
	}

	for _, instruction := range strings.Split(instructions, "\n") {
		m := instExp.FindStringSubmatch(instruction)
		if len(m) != 3 {
			return nil, errors.Errorf("unknown instruction (instruction=%s)", instruction)
		}
		gateDef := m[1]
		wire := m[2]

		m = gateExp.FindStringSubmatch(gateDef)
		if len(m) != 9 {
			return nil, errors.Errorf("unknown gate definition (instruction=%s)", instruction)
		}
		switch {
		case m[1] != "":
			c.wireToGate[wire] = PassthroughGate(constToWire(m[1]))
		case m[3] == "AND":
			c.wireToGate[wire] = &AndGate{constToWire(m[2]), constToWire(m[4])}
		case m[3] == "OR":
			c.wireToGate[wire] = &OrGate{constToWire(m[2]), constToWire(m[4])}
		case m[6] == "LSHIFT":
			c.wireToGate[wire] = &LShiftGate{constToWire(m[5]), constToWire(m[7])}
		case m[6] == "RSHIFT":
			c.wireToGate[wire] = &RShiftGate{constToWire(m[5]), constToWire(m[7])}
		case m[8] != "":
			c.wireToGate[wire] = NotGate(constToWire(m[8]))
		default:
			panic(fmt.Sprintf("matched but unhandled gateExp (instruction=%s)", instruction))
		}
	}

	return &c, nil
}

// Circuit is a list of wires each getting their signal from exactly one gate
type Circuit struct {
	wireToGate map[string]Gate
	wireValues map[string]uint16
}

func (c *Circuit) ResolveWire(wire string) (uint16, error) {
	// check if wire already has been resolved
	if v, found := c.wireValues[wire]; found {
		return v, nil
	}

	// resolve the wire
	gate, found := c.wireToGate[wire]
	if !found {
		return 0, errors.Errorf("wire not found in circuit (wire=%s)", wire)
	}
	v, err := gate.Resolve(c)
	if err != nil {
		return 0, err
	}

	// save the resolved value
	c.wireValues[wire] = v

	return v, nil
}

// Override force the wire specified to the value specified
func (c *Circuit) Override(wire string, val uint16) {
	c.wireValues[wire] = val
}

// Gate is an interface that can resolve the output signal of itself
type Gate interface {
	Resolve(c *Circuit) (uint16, error)
}

// InputGate takes no input and provides a constant output signal
type InputGate uint16

func (g InputGate) Resolve(_ *Circuit) (uint16, error) {
	return uint16(g), nil
}

// PassthroughGate takes one input signal and outputs it directly
type PassthroughGate string

func (g PassthroughGate) Resolve(c *Circuit) (uint16, error) {
	return c.ResolveWire(string(g))
}

// AndGate takes two inputs and outputs a bitwise AND of the two
type AndGate struct {
	inA, inB string
}

func (g *AndGate) Resolve(c *Circuit) (uint16, error) {
	inA, err := c.ResolveWire(g.inA)
	if err != nil {
		return 0, errors.Wrapf(err, "unable to resolve inA for AND (inA=%s)", g.inA)
	}
	inB, err := c.ResolveWire(g.inB)
	if err != nil {
		return 0, errors.Wrapf(err, "unable to resolve inB for AND (inB=%s)", g.inB)
	}

	return inA & inB, nil
}

// OrGate takes two inputs and outputs a bitwise OR of the two
type OrGate struct {
	inA, inB string
}

func (g *OrGate) Resolve(c *Circuit) (uint16, error) {
	inA, err := c.ResolveWire(g.inA)
	if err != nil {
		return 0, errors.Wrapf(err, "unable to resolve inA for OR (inA=%s)", g.inA)
	}
	inB, err := c.ResolveWire(g.inB)
	if err != nil {
		return 0, errors.Wrapf(err, "unable to resolve inB for OR (inB=%s)", g.inB)
	}

	return inA | inB, nil
}

// LShiftGate takes two inputs and outputs a the first input left shifted by the second input
type LShiftGate struct {
	inA, inB string
}

func (g *LShiftGate) Resolve(c *Circuit) (uint16, error) {
	inA, err := c.ResolveWire(g.inA)
	if err != nil {
		return 0, errors.Wrapf(err, "unable to resolve inA for LSHIFT (inA=%s)", g.inA)
	}
	inB, err := c.ResolveWire(g.inB)
	if err != nil {
		return 0, errors.Wrapf(err, "unable to resolve inB for LSHIFT (inB=%s)", g.inB)
	}

	return inA << inB, nil
}

// RShiftGate takes two inputs and outputs a the first input left shifted by the second input
type RShiftGate struct {
	inA, inB string
}

func (g *RShiftGate) Resolve(c *Circuit) (uint16, error) {
	inA, err := c.ResolveWire(g.inA)
	if err != nil {
		return 0, errors.Wrapf(err, "unable to resolve inA for RSHIFT (inA=%s)", g.inA)
	}
	inB, err := c.ResolveWire(g.inB)
	if err != nil {
		return 0, errors.Wrapf(err, "unable to resolve inB for RSHIFT (inB=%s)", g.inB)
	}

	return inA >> inB, nil
}

// NotGate takes one input signal and outputs the bitwise not of it
type NotGate string

func (g NotGate) Resolve(c *Circuit) (uint16, error) {
	out, err := c.ResolveWire(string(g))
	if err != nil {
		return 0, errors.Wrapf(err, "unable to resolve input for NOT (input=%s)", string(g))
	}
	return ^out, nil
}
