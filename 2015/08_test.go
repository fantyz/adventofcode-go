package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnescape(t *testing.T) {
	testCases := []struct {
		In, Out string
	}{
		{`""`, ``},
		{`"abc"`, `abc`},
		{`"aaa\"aaa"`, `aaa"aaa`},
		{`"\x27"`, `'`},
	}

	for _, c := range testCases {
		s, err := Unescape(c.In)
		if assert.NoError(t, err, c.In) {
			assert.Equal(t, c.Out, s, c.In)
		}
	}
}

func TestEscape(t *testing.T) {
	testCases := []struct {
		In, Out string
	}{
		{`""`, `"\"\""`},
		{`"abc"`, `"\"abc\""`},
		{`"aaa\"aaa"`, `"\"aaa\\\"aaa\""`},
		{`"\x27"`, `"\"\\x27\""`},
	}

	for _, c := range testCases {
		assert.Equal(t, c.Out, Escape(c.In), c.In)
	}
}

func TestSumEscapedCharactersMinusUnescapedCharacters(t *testing.T) {
	in := `""
"abc"
"aaa\"aaa"
"\x27"`
	sumEsc, sumUnesc, err := SumEscapedCharactersMinusUnescapedCharacters(in)
	if assert.NoError(t, err) {
		assert.Equal(t, 19, sumEsc, "escape")
		assert.Equal(t, 12, sumUnesc, "unescape")
	}
}

func TestDay8Pt1(t *testing.T) {
	_, sumUnesc, err := SumEscapedCharactersMinusUnescapedCharacters(day8Input)
	if assert.NoError(t, err) {
		assert.Equal(t, 1342, sumUnesc)
	}
}

func TestDay8Pt2(t *testing.T) {
	sumEsc, _, err := SumEscapedCharactersMinusUnescapedCharacters(day8Input)
	if assert.NoError(t, err) {
		assert.Equal(t, 2074, sumEsc)
	}
}
