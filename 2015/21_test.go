package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRPGSimulator20XX(t *testing.T) {
	assert.True(t, RPGSimulator20XX(
		Character{Hitpoints: 8, Damage: 5, Armor: 5},
		Character{Hitpoints: 12, Damage: 7, Armor: 2},
	))
}

func TestDay21Pt1(t *testing.T) {
	lowestToWin, _ := OptimizeRPGSimulator20XXGear(Character{
		Hitpoints: day21InputHP,
		Damage:    day21InputDamage,
		Armor:     day21InputArmor,
	})
	assert.Equal(t, 91, lowestToWin)
}

func TestDay21Pt2(t *testing.T) {
	_, highestToLoose := OptimizeRPGSimulator20XXGear(Character{
		Hitpoints: day21InputHP,
		Damage:    day21InputDamage,
		Armor:     day21InputArmor,
	})
	assert.Equal(t, 158, highestToLoose)
}
