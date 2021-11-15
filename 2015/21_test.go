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

}

func TestDay21Pt2(t *testing.T) {
}
