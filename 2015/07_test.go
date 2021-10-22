package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCircuit(t *testing.T) {
	instructions := `123 -> x
456 -> y
x AND y -> d
x OR y -> e
x LSHIFT 2 -> f
y RSHIFT 2 -> g
NOT x -> h
NOT y -> i`

	testCases := []struct {
		Wire   string
		ExpVal uint16
	}{
		{"d", 72},
		{"e", 507},
		{"f", 492},
		{"g", 114},
		{"h", 65412},
		{"i", 65079},
		{"x", 123},
		{"y", 456},
	}

	circuit, err := NewCircuit(instructions)
	if assert.NoError(t, err) {
		for _, c := range testCases {
			v, err := circuit.ResolveWire(c.Wire)
			if assert.NoError(t, err, "wire=%s", c.Wire) {
				assert.Equal(t, c.ExpVal, v, "wire=%s", c.Wire)
			}
		}
	}
}

func TestDay7Pt1(t *testing.T) {
	c, err := NewCircuit(day7Input)
	if assert.NoError(t, err, "failed to load circuit") {
		v, err := c.ResolveWire("a")
		if assert.NoError(t, err, "failed to resolve wire") {
			assert.Equal(t, uint16(46065), v)
		}
	}
}

func TestDay7Pt2(t *testing.T) {
	c, err := NewCircuit(day7Input)
	if assert.NoError(t, err, "failed to load circuit") {
		c.Override("b", 46065)
		v, err := c.ResolveWire("a")
		if assert.NoError(t, err, "failed to resolve wire") {
			assert.Equal(t, uint16(14134), v)
		}
	}
}
