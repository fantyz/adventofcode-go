package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExecuteOpcodeDay2(t *testing.T) {
	testCases := []struct {
		In     []int
		Input  []int
		Output []int
	}{
		{Load("3,9,8,9,10,9,4,9,99,-1,8"), []int{8}, []int{1}},
		{Load("3,9,8,9,10,9,4,9,99,-1,8"), []int{7}, []int{0}},
		{Load("3,9,7,9,10,9,4,9,99,-1,8"), []int{7}, []int{1}},
		{Load("3,9,7,9,10,9,4,9,99,-1,8"), []int{8}, []int{0}},
		{Load("3,3,1108,-1,8,3,4,3,99"), []int{8}, []int{1}},
		{Load("3,3,1108,-1,8,3,4,3,99"), []int{7}, []int{0}},
		{Load("3,3,1107,-1,8,3,4,3,99"), []int{7}, []int{1}},
		{Load("3,3,1107,-1,8,3,4,3,99"), []int{8}, []int{0}},
		{Load("109,1,204,-1,1001,100,1,100,1008,100,16,101,1006,101,0,99"), []int{}, Load("109,1,204,-1,1001,100,1,100,1008,100,16,101,1006,101,0,99")},
		{Load("1102,34915192,34915192,7,4,7,99,0"), []int{}, []int{1219070632396864}},
		{Load("104,1125899906842624,99"), []int{}, []int{1125899906842624}},
	}

	for i, c := range testCases {
		assert.Equal(t, c.Output, Outputter(ExecuteOpcode(c.In, Inputter(c.Input))), "output (case=%d)", i)
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
