package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReact(t *testing.T) {
	testCases := []struct {
		In    string
		Out   string
		Units int
	}{
		{"aA", "", 0},
		{"abBA", "", 0},
		{"abAB", "abAB", 4},
		{"aabAAB", "aabAAB", 6},
		{"dabAcCaCBAcCcaDA", "dabCBAcaDA", 10},
	}

	for i, c := range testCases {
		p := NewPolymers(c.In)
		p.React()

		assert.Equal(t, c.Out, p.String(), "(case=%d)", i)
		assert.Equal(t, c.Units, p.Units(), "(case=%d)", i)
	}
}

func TestRemoveUnit(t *testing.T) {
	testCases := []struct {
		In     string
		Remove string
		Out    string
	}{
		{"aA", "a", ""},
		{"aA", "A", ""},
		{"abBA", "b", "aA"},
		{"abAB", "B", "aA"},
		{"aabAAB", "a", "bB"},
		{"dabAcCaCBAcCcaDA", "c", "dabAaBAaDA"},
	}

	for i, c := range testCases {
		p := NewPolymers(c.In)
		p = p.RemoveUnit(c.Remove)
		assert.Equal(t, c.Out, p.String(), "(case=%d)", i)
	}
}

func TestShortestUnitsWithRemove(t *testing.T) {
	testCases := []struct {
		In  string
		Out int
	}{
		{"dabAcCaCBAcCcaDA", 4},
	}

	for i, c := range testCases {
		assert.Equal(t, c.Out, ShortestUnitsWithRemove(NewPolymers(c.In)), "(case=%d)", i)
	}
}
