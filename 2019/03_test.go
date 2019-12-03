package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddWire(t *testing.T) {
	testCases := []struct {
		In    string
		Dist  int
		Delay int
	}{
		{"R8,U5,L5,D3\nU7,R6,D4,L4", 6, 30},
		{"R75,D30,R83,U83,L12,D49,R71,U7,L72\nU62,R66,U55,R34,D71,R55,D58,R83", 159, 610},
		{"R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51\nU98,R91,D20,R16,D67,R40,U7,R15,U6,R7", 135, 410},
	}

	for i, c := range testCases {
		wiremoves := Load(c.In)
		board := NewCircuitboard()
		_, _ = board.AddWire(1, wiremoves[0])
		dist, delay := board.AddWire(2, wiremoves[1])
		assert.Equal(t, c.Dist, dist, "Dist (case=%d)", i)
		assert.Equal(t, c.Delay, delay, "Delay (case=%d)", i)
	}
}
