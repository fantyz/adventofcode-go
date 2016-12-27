package main

import (
	"fmt"
)

func main() {
	fmt.Println("Advent of Code 2015 - Day 22")

	initialFight := Fight{
		Boss: Stats{
			Hp: 55,
		},
		Player: Stats{
			Hp:   50,
			Mana: 500,
		},
	}

	best := 99999999
	outcomes := []Fight{initialFight}
	for len(outcomes) > 0 {
		newOutcomes := []Fight{}
		for i := range outcomes {
			newOutcomes = append(newOutcomes, tryAll(outcomes[i])...)
		}
		/*
			var s string
			fmt.Scanf("%s", s)
			_ = s
		*/

		// find wins
		outcomes = newOutcomes
		for i := range outcomes {
			if outcomes[i].Boss.Hp <= 0 && outcomes[i].Player.ManaSpent < best {
				best = outcomes[i].Player.ManaSpent
				fmt.Println("Player won using", outcomes[i].Player.ManaSpent, "mana")
			}
		}
	}
}

func tryAll(f Fight) (newFights []Fight) {
	if f.Player.Hp <= 0 || f.Boss.Hp <= 0 {
		return nil
	}

	if f.Player.Recharge > 0 {
		f.Player.Recharge--
		f.Player.Mana += 101
	}

	if f.Player.Mana >= 53 {
		mmFight := f
		mmFight.Player.ManaSpent += 53
		mmFight.Boss.Hp -= 4
		bossStep(&mmFight)
		newFights = append(newFights, mmFight)
	}
	if f.Player.Mana >= 73 {
		dFight := f
		dFight.Player.ManaSpent += 73
		dFight.Player.Hp += 2
		dFight.Boss.Hp -= 2
		bossStep(&dFight)
		newFights = append(newFights, dFight)
	}
	if f.Player.Mana >= 113 && f.Player.Shield == 0 {
		sFight := f
		sFight.Player.ManaSpent += 113
		sFight.Player.Shield = 3
		bossStep(&sFight)
		newFights = append(newFights, sFight)
	}
	if f.Player.Mana >= 173 && f.Boss.Poison == 0 {
		pFight := f
		pFight.Player.ManaSpent += 173
		pFight.Boss.Poison = 3
		bossStep(&pFight)
		newFights = append(newFights, pFight)
	}
	if f.Player.Mana >= 229 && f.Player.Recharge == 0 {
		rFight := f
		rFight.Player.ManaSpent += 229
		rFight.Player.Recharge = 5
		bossStep(&rFight)
		newFights = append(newFights, rFight)
	}

	return newFights
}

func bossStep(f *Fight) {
	if f.Player.Recharge > 0 {
		f.Player.Recharge--
		f.Player.Mana += 101
	}
	if f.Boss.Poison > 0 {
		f.Boss.Poison--
		f.Boss.Hp -= 6
	}
	if f.Player.Shield > 0 {
		f.Player.Shield--
		f.Player.Hp -= 1
	} else {
		f.Player.Hp -= 8
	}
}

type Fight struct {
	Boss, Player Stats
}

func (f *Fight) Print() {
	fmt.Printf("== PLAYER  hp=%d\tmana=%d (spent=%d)\n", f.Player.Hp, f.Player.Mana, f.Player.ManaSpent)
	if f.Player.Shield > 0 {
		fmt.Println("  > Has shield for", f.Player.Shield, "more rounds")
	}
	if f.Player.Recharge > 0 {
		fmt.Println("  > Has recharge for", f.Player.Recharge, "more rounds")
	}
	fmt.Printf("== BOSS    hp=%d\n", f.Boss.Hp)
	if f.Boss.Poison > 0 {
		fmt.Println("  > Has poison for", f.Boss.Poison, "more rounds")
	}
	if f.Boss.Hp <= 0 {
		fmt.Println("PLAYER WON!")
	} else if f.Player.Hp <= 0 {
		fmt.Println("BOSS WON!")
	}
	fmt.Println()
}

type Stats struct {
	Hp        int
	Mana      int
	ManaSpent int
	Shield    int
	Poison    int
	Recharge  int
}

const puzzleInput = `Hit Points: 55
Damage: 8`
