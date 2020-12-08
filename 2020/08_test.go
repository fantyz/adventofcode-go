package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBootCodeRun(t *testing.T) {
	testCases := map[string]struct {
		Program  string
		Debugger Debugger
		Acc      int
	}{
		"day8pt1test": {
			Program:  "nop +0\nacc +1\njmp +4\nacc +3\njmp -3\nacc -99\nacc +1\njmp -4\nacc +6",
			Debugger: &HaltOnFirstRepetitionDebugger{},
			Acc:      5,
		},
	}

	for n, c := range testCases {
		e, err := NewBootCodeExecuter(c.Program, c.Debugger)
		if assert.NoError(t, err, n) {
			assert.Equal(t, c.Acc, e.Run(), n)
		}
	}
}

func TestFindWorkingBootCodeProgram(t *testing.T) {
	testProgam := `nop +0
acc +1
jmp +4
acc +3
jmp -3
acc -99
acc +1
jmp -4
acc +6`
	acc, err := FindWorkingBootCodeProgram(testProgam)
	if assert.NoError(t, err) {
		assert.Equal(t, 8, acc)
	}
}

func TestDay8Pt1(t *testing.T) {
	e, err := NewBootCodeExecuter(day8Input, &HaltOnFirstRepetitionDebugger{})
	if assert.NoError(t, err) {
		assert.Equal(t, 1317, e.Run())
	}
}

func TestDay8Pt2(t *testing.T) {
	acc, err := FindWorkingBootCodeProgram(day8Input)
	if assert.NoError(t, err) {
		assert.Equal(t, 1033, acc)
	}
}
