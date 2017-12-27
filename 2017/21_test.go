package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewPainting(t *testing.T) {
	assert.Equal(t, 5, NewStartingPainting().PixelsOn())
}

func TestFlip(t *testing.T) {
	testCases := []struct {
		In     string
		Result string
	}{
		{`../.#`, `../#.`},
		{`.#./..#/##.`, `.#./#../.##`},
	}

	for i, testCase := range testCases {
		p := NewPainting(testCase.In)
		p.Flip()
		assert.Equal(t, testCase.Result, p.String(), "(case=%d)", i)
	}
}

func TestRotate(t *testing.T) {
	testCases := []struct {
		In     string
		Result string
	}{
		{`#./..`, `.#/..`},
		{`.#/..`, `../.#`},
		{`../.#`, `../#.`},
		{`../#.`, `#./..`},
		{`###/.../.#.`, `..#/#.#/..#`},
		{`..#/#.#/..#`, `.#./.../###`},
		{`.#./.../###`, `#../#.#/#..`},
		{`#../#.#/#..`, `###/.../.#.`},
	}

	for i, testCase := range testCases {
		p := NewPainting(testCase.In)
		p.Rotate()
		assert.Equal(t, testCase.Result, p.String(), "(case=%d)", i)
	}
}

func TestCopySubpainting(t *testing.T) {
	testCases := []struct {
		In         string
		Size, X, Y int
		Result     string
	}{
		{`...../.###./.#.#./.###./.....`, 3, 1, 1, `###/#.#/###`},
		{`...../.###./.#.#./.###./.....`, 3, 2, 2, `.#./##./...`},
		{`...../.###./.#.#./.###./.....`, 5, 0, 0, `...../.###./.#.#./.###./.....`},
		{`...../.###./.#.#./.###./.....`, 1, 0, 0, `.`},
	}

	for i, testCase := range testCases {
		assert.Equal(t, testCase.Result, NewPainting(testCase.In).CopySubpainting(testCase.X, testCase.Y, testCase.Size).String(), "(case=%d)", i)
	}
}

func TestPasteSubpainting(t *testing.T) {
	testCases := []struct {
		In     string
		X, Y   int
		Result string
	}{
		{`##/##`, 0, 0, `##../##../..../....`},
		{`##/##`, 1, 1, `..../.##./.##./....`},
		{`##/##`, 2, 2, `..../..../..##/..##`},
	}

	for i, testCase := range testCases {
		p := NewPainting(`..../..../..../....`)
		assert.Equal(t, testCase.Result, p.PasteSubpainting(testCase.X, testCase.Y, NewPainting(testCase.In)).String(), "(case=%d)", i)
	}
}

func TestMatch(t *testing.T) {
	ruleset := `../.# => ##./#../...
.#./..#/### => #..#/..../..../#..#`

	testCases := []struct {
		In     string
		Result string
	}{
		{`../.#`, `##./#../...`},
		{`.#./..#/###`, `#..#/..../..../#..#`},
	}

	rules := NewRules(ruleset)
	for i, testCase := range testCases {
		match, ok := rules.Match(NewPainting(testCase.In))
		if assert.True(t, ok, "(case=%d)", i) {
			assert.Equal(t, testCase.Result, match.String(), "(case=%d)", i)

		}
		match, ok = rules.Match(NewPainting(testCase.In).Flip())
		if assert.True(t, ok, "(case=%d)", i) {
			assert.Equal(t, testCase.Result, match.String(), "(case=%d)", i)

		}

	}
}
