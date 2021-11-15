package main

import (
	"fmt"
)

func init() { days["21"] = Day21 }

/*
 */

func Day21() {
	fmt.Println("--- Day 21: RPG Simulator 20XX ---")
}

func RPGSimulator20XX(player, boss Character) bool {
	for {
		boss.Hitpoints -= calcDamage(player.Damage, boss.Armor)
		if boss.Hitpoints <= 0 {
			return true
		}
		player.Hitpoints -= calcDamage(boss.Damage, player.Armor)
		if player.Hitpoints <= 0 {
			return false
		}
	}
}

func calcDamage(damage int, armor int) int {
	damage -= armor
	if damage <= 0 {
		damage = 1
	}
	return damage
}

type Character struct {
	Hitpoints int
	Armor     int
	Damage    int
}
