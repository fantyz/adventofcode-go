package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExecuteOpcodeDay2(t *testing.T) {
	testCases := []struct {
		In      []int
		Input   []int
		Outprog []int
		Output  []int
	}{
		{Load("1,0,0,0,99"), nil, Load("2,0,0,0,99"), nil},
		{Load("2,3,0,3,99"), nil, Load("2,3,0,6,99"), nil},
		{Load("2,4,4,5,99,0"), nil, Load("2,4,4,5,99,9801"), nil},
		{Load("1,1,1,4,99,5,6,0,99"), nil, Load("30,1,1,4,2,5,6,0,99"), nil},
		{Load("1,9,10,3,2,3,11,0,99,30,40,50"), nil, Load("3500,9,10,70,2,3,11,0,99,30,40,50"), nil},
		{Load("1101,100,-1,4,0"), nil, Load("1101,100,-1,4,99"), nil},
		{Load("3,9,8,9,10,9,4,9,99,-1,8"), []int{8}, Load("3,9,8,9,10,9,4,9,99,1,8"), []int{1}},
		{Load("3,9,8,9,10,9,4,9,99,-1,8"), []int{7}, Load("3,9,8,9,10,9,4,9,99,0,8"), []int{0}},
		{Load("3,9,7,9,10,9,4,9,99,-1,8"), []int{7}, Load("3,9,7,9,10,9,4,9,99,1,8"), []int{1}},
		{Load("3,9,7,9,10,9,4,9,99,-1,8"), []int{8}, Load("3,9,7,9,10,9,4,9,99,0,8"), []int{0}},
		{Load("3,3,1108,-1,8,3,4,3,99"), []int{8}, Load("3,3,1108,1,8,3,4,3,99"), []int{1}},
		{Load("3,3,1108,-1,8,3,4,3,99"), []int{7}, Load("3,3,1108,0,8,3,4,3,99"), []int{0}},
		{Load("3,3,1107,-1,8,3,4,3,99"), []int{7}, Load("3,3,1107,1,8,3,4,3,99"), []int{1}},
		{Load("3,3,1107,-1,8,3,4,3,99"), []int{8}, Load("3,3,1107,0,8,3,4,3,99"), []int{0}},
		{Load("109,1,204,-1,1001,100,1,100,1008,100,16,101,1006,101,0,99"), []int{}, nil, Load("109,1,204,-1,1001,100,1,100,1008,100,16,101,1006,101,0,99")},
		{Load("1102,34915192,34915192,7,4,7,99,0"), []int{}, Load("1102,34915192,34915192,7,4,7,99,1219070632396864"), []int{1219070632396864}},
		{Load("104,1125899906842624,99"), []int{}, Load("104,1125899906842624,99"), []int{1125899906842624}},
	}

	for i, c := range testCases {
		outprog, out := ExecuteOpcode(c.In, Inputter(c.Input))
		if c.Outprog != nil {
			assert.Equal(t, c.Outprog, outprog, "outprog (case=%d)", i)
		}
		assert.Equal(t, c.Output, out, "output (case=%d)", i)
	}
}

func TestOpcode(t *testing.T) {
	testCases := []struct {
		In int
		Op [4]int
	}{
		{1, [4]int{1, 0, 0, 0}},
		{1002, [4]int{2, 0, 1, 0}},
		{20131, [4]int{31, 1, 0, 2}},
	}

	for i, c := range testCases {
		assert.Equal(t, c.Op, opcode(c.In), "(case=%d)", i)
	}
}
