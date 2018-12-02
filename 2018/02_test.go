package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestChecksum(t *testing.T) {
	testCases := []struct {
		In  []string
		Out int
	}{
		{[]string{"abcdef", "bababc", "abbcde", "abcccd", "aabcdd", "abcdee", "ababab"}, 12},
	}

	for i, c := range testCases {
		chksum := Checksum(c.In)
		assert.Equal(t, c.Out, chksum, "(case=%d)", i)
	}
}

func TestCommonLetters(t *testing.T) {
	testCases := []struct {
		In  []string
		Out string
	}{
		{[]string{"abcde", "fghij", "klmno", "pqrst", "fguij", "axcye", "wvxyz"}, "fgij"},
	}

	for i, c := range testCases {
		common := CommonLetters(c.In)
		assert.Equal(t, c.Out, common, "(case=%d)", i)
	}
}
