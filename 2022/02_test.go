package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRockPaperScissorsStrategyGuideScore(t *testing.T) {
	in := `A Y
B X
C Z`
	guide, err := NewRockPaperScissorsStrategyGuide(in)
	if assert.NoError(t, err) {
		assert.Equal(t, 15, guide.Score(false), "moving")
		assert.Equal(t, 12, guide.Score(true), "responding")
	}
}

func TestDay02Pt1(t *testing.T) {
	guide, err := NewRockPaperScissorsStrategyGuide(day02Input)
	if assert.NoError(t, err) {
		assert.Equal(t, 14069, guide.Score(false))
	}
}

func TestDay02Pt2(t *testing.T) {
	guide, err := NewRockPaperScissorsStrategyGuide(day02Input)
	if assert.NoError(t, err) {
		assert.Equal(t, 12411, guide.Score(true))
	}
}
