package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRunMaze(t *testing.T) {
	testCases := []struct {
		In     string
		Result string
		Steps  int
	}{
		{`     |          
     |  +--+    
     A  |  C    
 F---|----E|--+ 
     |  |  |  D 
     +B-+  +--+`, "ABCDEF", 38},
	}

	for i, testCase := range testCases {
		letters, steps := NewMaze(testCase.In).Run()
		assert.Equal(t, testCase.Result, letters, "(case=%d)", i)
		assert.Equal(t, testCase.Steps, steps, "(case=%d)", i)
	}
}
