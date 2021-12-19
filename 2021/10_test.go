package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSyntaxErrorScore(t *testing.T) {
	in := `[({(<(())[]>[[{[]{<()<>>
[(()[<>])]({[<{<<[]>>(
{([(<{}[<>[]}>{[]{[(<()>
(((({<>}<{<{<>}{[]{[]{}
[[<[([]))<([[{}[[()]]]
[{[{({}]{}}([{[{{{}}([]
{<[[]]>}<{[{[{[]{()[[[]
[<(<(<(<{}))><([]([]()
<{([([[(<>()){}]>(<<{{
<{([{{}}[<[[[<>{}]]]>[]]`

	nav, err := NewNavigationSubsystem(in)
	if assert.NoError(t, err) {
		assert.Equal(t, 26397, nav.SyntaxErrorScore())
	}
}

func TestMiddleClosignSequenceScore(t *testing.T) {
	in := `[({(<(())[]>[[{[]{<()<>>
[(()[<>])]({[<{<<[]>>(
{([(<{}[<>[]}>{[]{[(<()>
(((({<>}<{<{<>}{[]{[]{}
[[<[([]))<([[{}[[()]]]
[{[{({}]{}}([{[{{{}}([]
{<[[]]>}<{[{[{[]{()[[[]
[<(<(<(<{}))><([]([]()
<{([([[(<>()){}]>(<<{{
<{([{{}}[<[[[<>{}]]]>[]]`

	nav, err := NewNavigationSubsystem(in)
	if assert.NoError(t, err) {
		assert.Equal(t, 288957, nav.MiddleClosignSequenceScore())
	}
}

func TestDay10Pt1(t *testing.T) {
	nav, err := NewNavigationSubsystem(day10Input)
	if assert.NoError(t, err) {
		assert.Equal(t, 392097, nav.SyntaxErrorScore())
	}
}

func TestDay10Pt2(t *testing.T) {
	nav, err := NewNavigationSubsystem(day10Input)
	if assert.NoError(t, err) {
		assert.Equal(t, 4263222782, nav.MiddleClosignSequenceScore())
	}
}
