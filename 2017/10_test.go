package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestKnotHash(t *testing.T) {
	testCases := []struct {
		In     string
		Size   int
		Result int
	}{
		{`3,4,1,5`, 5, 12},
	}

	for i, testCase := range testCases {
		assert.Equal(t, testCase.Result, KnotHash(testCase.Size, ReadLengths(testCase.In)), "%v (case=%d)", testCase.In, i)
	}
}

func TestReadLengthsToBytes(t *testing.T) {
	testCases := []struct {
		In     string
		Result []byte
	}{
		{`1,2,3`, []byte{49, 44, 50, 44, 51, 17, 31, 73, 47, 23}},
	}

	for i, testCase := range testCases {
		assert.Equal(t, testCase.Result, ReadLengthsToBytes(testCase.In), "%v (case=%d)", testCase.In, i)
	}
}

func TestKnotHashBytes(t *testing.T) {
	testCases := []struct {
		In     string
		Result string
	}{
		{``, "a2582a3a0e66e6e86e3812dcb672a272"},
		{`AoC 2017`, "33efeb34ea91902bb2f59c9920caa6cd"},
		{`1,2,3`, "3efbe78a8d82f29979031a4aa0b16a9d"},
		{`1,2,4`, "63960835bcdc130f0b66d7ff4f6a5a8e"},
	}

	for i, testCase := range testCases {
		assert.Equal(t, testCase.Result, KnotHashBytes(ReadLengthsToBytes(testCase.In)), "%v (case=%d)", testCase.In, i)
	}
}
