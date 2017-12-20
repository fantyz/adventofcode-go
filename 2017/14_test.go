package main

import (
	"encoding/hex"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHashToBitRow(t *testing.T) {
	testCases := []struct {
		In     string
		Result string
	}{
		{`a0c20170000000000000000000000000`, `10100000110000100000000101110000`},
	}

	for i, testCase := range testCases {
		hash, err := hex.DecodeString(testCase.In)
		if err != nil {
			panic(err)
		}

		row := HashToBitRow(hash)
		str := ""
		for i := 0; i < len(row); i++ {
			if row[i] {
				str += "1"
			} else {
				str += "0"
			}
		}

		assert.Equal(t, testCase.Result, str[:len(testCase.Result)], "(case=%d)", i)
	}
}

func TestGridRow(t *testing.T) {
	testCases := []struct {
		In     string
		Result string
	}{
		{`flqrgnkx-0`, `11010100`},
		{`flqrgnkx-1`, `01010101`},
		{`flqrgnkx-2`, `00001010`},
	}

	for i, testCase := range testCases {
		hash := KnotHash(testCase.In)

		row := HashToBitRow(hash)
		str := ""
		for i := 0; i < len(row); i++ {
			if row[i] {
				str += "1"
			} else {
				str += "0"
			}
		}

		assert.Equal(t, testCase.Result, str[:len(testCase.Result)], "(case=%d)", i)
	}
}

func TestUsedSquares(t *testing.T) {
	testCases := []struct {
		In     string
		Result int
	}{
		{`flqrgnkx`, 8108},
	}

	for i, testCase := range testCases {
		assert.Equal(t, testCase.Result, NewGrid(128, testCase.In).UsedSquares(), "(case=%d)", i)
	}
}

func TestRegions(t *testing.T) {
	testCases := []struct {
		In     string
		Result int
	}{
		{`flqrgnkx`, 1242},
	}

	for i, testCase := range testCases {
		assert.Equal(t, testCase.Result, NewGrid(128, testCase.In).Regions(), "(case=%d)", i)
	}
}
