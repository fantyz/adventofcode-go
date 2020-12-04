package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsNice(t *testing.T) {
	testCases := []struct {
		Str    string
		IsNice bool
	}{
		{"ugknbfddgicrmopn", true},
		{"aaa", true},
		{"jchzalrnumimnmhp", false},
		{"haegwjzuvuyypxyu", false},
		{"dvszwmarrgswjxmb", false},
	}

	for _, c := range testCases {
		assert.Equal(t, c.IsNice, IsNice(c.Str), c.Str)
	}
}

func TestNiceStrings(t *testing.T) {
	testCase := `ugknbfddgicrmopn
aaa
jchzalrnumimnmhp
haegwjzuvuyypxyu
dvszwmarrgswjxmb`

	assert.Equal(t, 2, NiceStrings(testCase))
}

func TestIsNicer(t *testing.T) {
	testCases := []struct {
		Str     string
		IsNicer bool
	}{
		{"qjhvhtzxzqqjkmpb", true},
		{"xxyxx", true},
		{"uurcxstgmygtbstg", false},
		{"ieodomkazucvgmuy", false},
	}

	for _, c := range testCases {
		assert.Equal(t, c.IsNicer, IsNicer(c.Str), c.Str)
	}
}

func TestNicerStrings(t *testing.T) {
	testCase := `qjhvhtzxzqqjkmpb
xxyxx
uurcxstgmygtbstg
ieodomkazucvgmuy`

	assert.Equal(t, 2, NicerStrings(testCase))
}

func TestDay5Pt1(t *testing.T) {
	assert.Equal(t, 238, NiceStrings(day5Input))
}

func TestDay5Pt2(t *testing.T) {
	assert.Equal(t, 69, NicerStrings(day5Input))
}
