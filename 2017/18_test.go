package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProgramRun(t *testing.T) {
	testCases := []struct {
		In     string
		Result int
	}{
		{`set a 1
add a 2
mul a a
mod a 5
snd a
set a 0
rcv a
jgz a -1
set a 1
jgz a -2`, 4},
	}

	for i, testCase := range testCases {
		assert.Equal(t, testCase.Result, NewProgram(testCase.In).RunUntilRcv(), "(case=%d)", i)
	}
}

func TestProgram2Run(t *testing.T) {
	testCases := []struct {
		In     string
		Result int
	}{
		{`snd 1
snd 2
snd p
rcv a
rcv b
rcv c
rcv d`, 3},
	}

	for i, testCase := range testCases {
		assert.Equal(t, testCase.Result, RunInParallel(testCase.In), "(case=%d)", i)
	}
}
