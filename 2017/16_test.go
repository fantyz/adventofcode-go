package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDanceSpin(t *testing.T) {
	testCases := []struct {
		In      int
		Dancers string
		Result  string
	}{
		{0, "abcde", "abcde"},
		{1, "abcde", "eabcd"},
		{2, "abcde", "deabc"},
		{3, "abcde", "cdeab"},
		{4, "abcde", "bcdea"},
		{5, "abcde", "abcde"},
	}

	for i, testCase := range testCases {
		d := NewDance(testCase.Dancers)
		d.Spin(testCase.In)
		assert.Equal(t, testCase.Result, d.String(), "(case=%d)", i)
	}
}

func TestDanceExchange(t *testing.T) {
	testCases := []struct {
		InA     int
		InB     int
		Dancers string
		Result  string
	}{
		{0, 1, "abcde", "bacde"},
		{3, 4, "eabcd", "eabdc"},
	}

	for i, testCase := range testCases {
		d := NewDance(testCase.Dancers)
		d.Exchange(testCase.InA, testCase.InB)
		assert.Equal(t, testCase.Result, d.String(), "(case=%d)", i)
	}
}

func TestDanceExchangeWithSpin(t *testing.T) {
	testCases := []struct {
		InA     int
		InB     int
		Spin    int
		Dancers string
		Result  string
	}{
		{0, 1, 1, "abcde", "aebcd"},
		{3, 4, 1, "eabcd", "deacb"},
	}

	for i, testCase := range testCases {
		d := NewDance(testCase.Dancers)
		d.Spin(testCase.Spin)
		d.Exchange(testCase.InA, testCase.InB)
		assert.Equal(t, testCase.Result, d.String(), "(case=%d)", i)
	}
}

func TestDancePartner(t *testing.T) {
	testCases := []struct {
		InA     byte
		InB     byte
		Dancers string
		Result  string
	}{
		{'a', 'b', "abcde", "bacde"},
		{'e', 'b', "eabdc", "baedc"},
	}

	for i, testCase := range testCases {
		d := NewDance(testCase.Dancers)
		d.Swap(testCase.InA, testCase.InB)
		assert.Equal(t, testCase.Result, d.String(), "(case=%d)", i)
	}
}

func TestDoTheDance(t *testing.T) {
	testCases := []struct {
		In      string
		Dancers string
		Times   int
		Result  string
	}{
		{`s1,x3/4,pe/b`, "abcde", 1, "baedc"},
		{`s1,x3/4,pe/b`, "abcde", 2, "ceadb"},
	}

	for i, testCase := range testCases {
		assert.Equal(t, testCase.Result, DoTheDance(testCase.Dancers, testCase.In, testCase.Times), "(case=%d)", i)
	}
}
