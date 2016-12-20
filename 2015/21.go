package main

import (
	"fmt"
)

/*

--- Day 21: RPG Simulator 20XX ---

Little Henry Case got a new video game for Christmas. It's an RPG, and he's stuck on a boss. He needs to know what equipment to buy at the shop. He hands you the controller.

In this game, the player (you) and the enemy (the boss) take turns attacking. The player always goes first. Each attack reduces the opponent's hit points by at least 1. The first character at or below 0 hit points loses.

Damage dealt by an attacker each turn is equal to the attacker's damage score minus the defender's armor score. An attacker always does at least 1 damage. So, if the attacker has a damage score of 8, and the defender has an armor score of 3, the defender loses 5 hit points. If the defender had an armor score of 300, the defender would still lose 1 hit point.

Your damage score and armor score both start at zero. They can be increased by buying items in exchange for gold. You start with no items and have as much gold as you need. Your total damage or armor is equal to the sum of those stats from all of your items. You have 100 hit points.

Here is what the item shop is selling:

Weapons:    Cost  Damage  Armor
Dagger        8     4       0
Shortsword   10     5       0
Warhammer    25     6       0
Longsword    40     7       0
Greataxe     74     8       0

Armor:      Cost  Damage  Armor
Leather      13     0       1
Chainmail    31     0       2
Splintmail   53     0       3
Bandedmail   75     0       4
Platemail   102     0       5

Rings:      Cost  Damage  Armor
Damage +1    25     1       0
Damage +2    50     2       0
Damage +3   100     3       0
Defense +1   20     0       1
Defense +2   40     0       2
Defense +3   80     0       3
You must buy exactly one weapon; no dual-wielding. Armor is optional, but you can't use more than one. You can buy 0-2 rings (at most one for each hand). You must use any items you buy. The shop only has one of each item, so you can't buy, for example, two rings of Damage +3.

For example, suppose you have 8 hit points, 5 damage, and 5 armor, and that the boss has 12 hit points, 7 damage, and 2 armor:

The player deals 5-2 = 3 damage; the boss goes down to 9 hit points.
The boss deals 7-5 = 2 damage; the player goes down to 6 hit points.
The player deals 5-2 = 3 damage; the boss goes down to 6 hit points.
The boss deals 7-5 = 2 damage; the player goes down to 4 hit points.
The player deals 5-2 = 3 damage; the boss goes down to 3 hit points.
The boss deals 7-5 = 2 damage; the player goes down to 2 hit points.
The player deals 5-2 = 3 damage; the boss goes down to 0 hit points.
In this scenario, the player wins! (Barely.)

You have 100 hit points. The boss's actual stats are in your puzzle input. What is the least amount of gold you can spend and still win the fight?

Your puzzle answer was 121.

--- Part Two ---

Turns out the shopkeeper is working with the boss, and can persuade you to buy whatever items he wants. The other rules still apply, and he still only has one of each item.

What is the most amount of gold you can spend and still lose the fight?

Your puzzle answer was 201.

*/

func main() {
	fmt.Println("Advent of Code 2015 - Day 21")

	weapons := []Item{
		{8, Stats{Damage: 4}},
		{10, Stats{Damage: 5}},
		{25, Stats{Damage: 6}},
		{40, Stats{Damage: 7}},
		{74, Stats{Damage: 8}},
	}

	armors := []Item{
		{0, Stats{}},
		{13, Stats{Armor: 1}},
		{31, Stats{Armor: 2}},
		{53, Stats{Armor: 3}},
		{75, Stats{Armor: 4}},
		{102, Stats{Armor: 5}},
	}

	rings := []Item{
		{0, Stats{}},
		{0, Stats{}},
		{20, Stats{Armor: 1}},
		{25, Stats{Damage: 1}},
		{40, Stats{Armor: 2}},
		{50, Stats{Damage: 2}},
		{80, Stats{Armor: 3}},
		{100, Stats{Damage: 3}},
	}

	bestCost := 999999
	worstCost := 0
	for weapon := 0; weapon < len(weapons); weapon++ {
		for armor := 0; armor < len(armors); armor++ {
			for ring1 := 0; ring1 < len(rings); ring1++ {
				for ring2 := ring1 + 1; ring2 < len(rings); ring2++ {
					cost := weapons[weapon].Cost + armors[armor].Cost + rings[ring1].Cost + rings[ring2].Cost
					if Fight(Stats{
						Hp:     100,
						Damage: weapons[weapon].Stats.Damage + armors[armor].Stats.Damage + rings[ring1].Stats.Damage + rings[ring2].Stats.Damage,
						Armor:  weapons[weapon].Stats.Armor + armors[armor].Stats.Armor + rings[ring1].Stats.Armor + rings[ring2].Stats.Armor,
					}) {
						if bestCost > cost {
							bestCost = cost
						}
					} else {
						if worstCost < cost {
							worstCost = cost
						}
					}
				}
			}
		}
	}

	fmt.Println("Part 1: Best cost is", bestCost)
	fmt.Println("Part 2: Worst cost is", worstCost)
}

type Item struct {
	Cost  int
	Stats Stats
}

type Stats struct {
	Hp     int
	Damage int
	Armor  int
}

var Boss = Stats{
	Hp:     103,
	Damage: 9,
	Armor:  2,
}

func Fight(p Stats) bool {
	b := Boss
	pDmg := p.Damage - b.Armor
	if pDmg <= 0 {
		pDmg = 1
	}
	bDmg := b.Damage - p.Armor
	if bDmg <= 0 {
		bDmg = 1
	}

	for {
		b.Hp -= pDmg
		if b.Hp <= 0 {
			return true
		}
		p.Hp -= bDmg
		if p.Hp <= 0 {
			return false
		}
	}
}

const puzzleInput = `Hit Points: 103
Damage: 9
Armor: 2`
