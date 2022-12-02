package main

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

func init() { days["2"] = Day2 }

/*
--- Day 2: Rock Paper Scissors ---
The Elves begin to set up camp on the beach. To decide whose tent gets to be closest to the snack storage, a giant Rock Paper Scissors tournament is already in progress.

Rock Paper Scissors is a game between two players. Each game contains many rounds; in each round, the players each simultaneously choose one of Rock, Paper, or Scissors using a hand shape. Then, a winner for that round is selected: Rock defeats Scissors, Scissors defeats Paper, and Paper defeats Rock. If both players choose the same shape, the round instead ends in a draw.

Appreciative of your help yesterday, one Elf gives you an encrypted strategy guide (your puzzle input) that they say will be sure to help you win. "The first column is what your opponent is going to play: A for Rock, B for Paper, and C for Scissors. The second column--" Suddenly, the Elf is called away to help with someone's tent.

The second column, you reason, must be what you should play in response: X for Rock, Y for Paper, and Z for Scissors. Winning every time would be suspicious, so the responses must have been carefully chosen.

The winner of the whole tournament is the player with the highest score. Your total score is the sum of your scores for each round. The score for a single round is the score for the shape you selected (1 for Rock, 2 for Paper, and 3 for Scissors) plus the score for the outcome of the round (0 if you lost, 3 if the round was a draw, and 6 if you won).

Since you can't be sure if the Elf is trying to help you or trick you, you should calculate the score you would get if you were to follow the strategy guide.

For example, suppose you were given the following strategy guide:

A Y
B X
C Z
This strategy guide predicts and recommends the following:

In the first round, your opponent will choose Rock (A), and you should choose Paper (Y). This ends in a win for you with a score of 8 (2 because you chose Paper + 6 because you won).
In the second round, your opponent will choose Paper (B), and you should choose Rock (X). This ends in a loss for you with a score of 1 (1 + 0).
The third round is a draw with both players choosing Scissors, giving you a score of 3 + 3 = 6.
In this example, if you were to follow the strategy guide, you would get a total score of 15 (8 + 1 + 6).

What would your total score be if everything goes exactly according to your strategy guide?

Your puzzle answer was 14069.

--- Part Two ---
The Elf finishes helping with the tent and sneaks back over to you. "Anyway, the second column says how the round needs to end: X means you need to lose, Y means you need to end the round in a draw, and Z means you need to win. Good luck!"

The total score is still calculated in the same way, but now you need to figure out what shape to choose so the round ends as indicated. The example above now goes like this:

In the first round, your opponent will choose Rock (A), and you need the round to end in a draw (Y), so you also choose Rock. This gives you a score of 1 + 3 = 4.
In the second round, your opponent will choose Paper (B), and you choose Rock so you lose (X) with a score of 1 + 0 = 1.
In the third round, you will defeat your opponent's Scissors with Rock for a score of 1 + 6 = 7.
Now that you're correctly decrypting the ultra top secret strategy guide, you would get a total score of 12.

Following the Elf's instructions for the second column, what would your total score be if everything goes exactly according to your strategy guide?

Your puzzle answer was 12411.
*/

func Day2() {
	fmt.Println("--- Day 2: Rock Paper Scissors ---")

	guide, err := NewRockPaperScissorsStrategyGuide(day02Input)
	if err != nil {
		fmt.Println(errors.Wrap(err, "Unable to create guide"))
		return
	}
	fmt.Println("Total score if everything goes according to the guide:", guide.Score(false))
	fmt.Println("Total score if everything goes according to the guide responding correctly:", guide.Score(true))
}

func NewRockPaperScissorsStrategyGuide(in string) (RockPaperScissorsStrategyGuide, error) {
	var guide RockPaperScissorsStrategyGuide
	for _, line := range strings.Split(in, "\n") {
		if len(line) != 3 {
			return nil, errors.Errorf("Line not 3 long (line=%s)", line)
		}
		var r Round
		switch line[0] {
		case 'A':
			r.Opponent = RockMoveOrLooseResponse
		case 'B':
			r.Opponent = PaperMoveOrDrawResponse
		case 'C':
			r.Opponent = ScissorsMoveOrWinResponse
		default:
			return nil, errors.Errorf("Unexpected opponent move (line=%s)", line)
		}

		switch line[2] {
		case 'X':
			r.You = RockMoveOrLooseResponse
		case 'Y':
			r.You = PaperMoveOrDrawResponse
		case 'Z':
			r.You = ScissorsMoveOrWinResponse
		default:
			return nil, errors.Errorf("Unexpected response (line=%s)", line)
		}
		guide = append(guide, r)
	}
	return guide, nil
}

type RockPaperScissorsStrategyGuide []Round

func (g RockPaperScissorsStrategyGuide) Score(respond bool) int {
	score := 0
	for _, r := range g {
		score += r.Score(respond)
	}
	return score
}

type Round struct {
	Opponent RockPaperScissorsMoveOrResponse
	You      RockPaperScissorsMoveOrResponse
}

type RockPaperScissorsMoveOrResponse int

const (
	RockMoveOrLooseResponse   RockPaperScissorsMoveOrResponse = 0
	PaperMoveOrDrawResponse   RockPaperScissorsMoveOrResponse = 1
	ScissorsMoveOrWinResponse RockPaperScissorsMoveOrResponse = 2
)

func (r *Round) Score(respond bool) int {
	switch r.Opponent {
	case RockMoveOrLooseResponse:
		switch r.You {
		case RockMoveOrLooseResponse:
			if respond {
				// scissors loose
				return 3 + 0
			}

			// rock draw
			return 1 + 3
		case PaperMoveOrDrawResponse:
			if respond {
				// rocks draw
				return 1 + 3
			}

			// paper wins
			return 2 + 6
		case ScissorsMoveOrWinResponse:
			if respond {
				// paper wins
				return 2 + 6
			}

			// scissor loose
			return 3 + 0
		}
	case PaperMoveOrDrawResponse:
		switch r.You {
		case RockMoveOrLooseResponse:
			if respond {
				// rock loose
				return 1 + 0
			}

			// rock loose
			return 1 + 0
		case PaperMoveOrDrawResponse:
			if respond {
				// paper draw
				return 2 + 3
			}

			// paper draw
			return 2 + 3
		case ScissorsMoveOrWinResponse:
			if respond {
				// scissors wins
				return 3 + 6
			}

			// scissor win
			return 3 + 6
		}
	case ScissorsMoveOrWinResponse:
		switch r.You {
		case RockMoveOrLooseResponse:
			if respond {
				// paper loose
				return 2 + 0
			}

			// rock win
			return 1 + 6
		case PaperMoveOrDrawResponse:
			if respond {
				// scissor draw
				return 3 + 3
			}

			// paper loose
			return 2 + 0
		case ScissorsMoveOrWinResponse:
			if respond {
				// rock wins
				return 1 + 6
			}

			// scissor draw
			return 3 + 3
		}
	}

	panic(fmt.Sprintf("Unexpected move (opponent=%d, you=%d)", r.Opponent, r.You))
}
