package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlayCombat(t *testing.T) {
	assert.Equal(t, 306, PlayCombat(NewDeck(9, 2, 6, 3, 1), NewDeck(5, 8, 4, 7, 10)))
}

func TestPlayRecursiveCombatInfiniteGame(t *testing.T) {
	_, score := PlayRecursiveCombat(NewDeck(43, 19), NewDeck(2, 29, 14))
	assert.Equal(t, 105, score)
}

func TestPlayRecursiveCombat(t *testing.T) {
	_, score := PlayRecursiveCombat(NewDeck(9, 2, 6, 3, 1), NewDeck(5, 8, 4, 7, 10))
	assert.Equal(t, 291, score)
}

func TestDay22Pt1(t *testing.T) {
	p1input, err1 := LoadInts(day22InputP1, "\n")
	p2input, err2 := LoadInts(day22InputP2, "\n")
	if assert.NoError(t, err1, "p1") && assert.NoError(t, err2, "p2") {
		assert.Equal(t, 31455, PlayCombat(NewDeck(p1input...), NewDeck(p2input...)))
	}
}

func TestDay22Pt2(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping slow running recursive combat")
	}

	p1input, err1 := LoadInts(day22InputP1, "\n")
	p2input, err2 := LoadInts(day22InputP2, "\n")
	if assert.NoError(t, err1, "p1") && assert.NoError(t, err2, "p2") {
		_, score := PlayRecursiveCombat(NewDeck(p1input...), NewDeck(p2input...))
		assert.Equal(t, 32528, score)
	}
}
