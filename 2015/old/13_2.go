package main

/*
--- Day 13: Knights of the Dinner Table ---

In years past, the holiday feast with your family hasn't gone so well. Not everyone gets along! This year, you resolve, will be different. You're going to find the optimal seating arrangement and avoid all those awkward conversations.

You start by writing up a list of everyone invited and the amount their happiness would increase or decrease if they were to find themselves sitting next to each other person. You have a circular table that will be just big enough to fit everyone comfortably, and so each person will have exactly two neighbors.

For example, suppose you have only four attendees planned, and you calculate their potential happiness as follows:

Alice would gain 54 happiness units by sitting next to Bob.
Alice would lose 79 happiness units by sitting next to Carol.
Alice would lose 2 happiness units by sitting next to David.
Bob would gain 83 happiness units by sitting next to Alice.
Bob would lose 7 happiness units by sitting next to Carol.
Bob would lose 63 happiness units by sitting next to David.
Carol would lose 62 happiness units by sitting next to Alice.
Carol would gain 60 happiness units by sitting next to Bob.
Carol would gain 55 happiness units by sitting next to David.
David would gain 46 happiness units by sitting next to Alice.
David would lose 7 happiness units by sitting next to Bob.
David would gain 41 happiness units by sitting next to Carol.
Then, if you seat Alice next to David, Alice would lose 2 happiness units (because David talks so much), but David would gain 46 happiness units (because Alice is such a good listener), for a total change of 44.

If you continue around the table, you could then seat Bob next to Alice (Bob gains 83, Alice gains 54). Finally, seat Carol, who sits next to Bob (Carol gains 60, Bob loses 7) and David (Carol gains 55, David gains 41). The arrangement looks like this:

     +41 +46
+55   David    -2
Carol       Alice
+60    Bob    +54
     -7  +83
After trying every other seating arrangement in this hypothetical scenario, you find that this one is the most optimal, with a total change in happiness of 330.

What is the total change in happiness for the optimal seating arrangement of the actual guest list?

Your puzzle answer was 709.

The first half of this puzzle is complete! It provides one gold star: *

--- Part Two ---

In all the commotion, you realize that you forgot to seat yourself. At this point, you're pretty apathetic toward the whole thing, and your happiness wouldn't really go up or down regardless of who you sit next to. You assume everyone else would be just as ambivalent about sitting next to you, too.

So, add yourself to the list, and give all happiness relationships that involve you a score of 0.

What is the total change in happiness for the optimal seating arrangement that actually includes yourself?
*/

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
)

type Happiness map[string]map[string]int

func (h Happiness) AddSeating(a, b string, happiness int) {
	if _, found := h[a]; !found {
		h[a] = make(map[string]int)
	}
	h[a][b] = happiness
}

func (h Happiness) CalculateBasedOnSeating(seating []string) int {
	happiness := 0
	for i:=0; i < len(seating)-1; i++ {
		happiness += h[seating[i]][seating[i+1]] + h[seating[i+1]][seating[i]]
	}
	happiness += h[seating[len(seating)-1]][seating[0]] + h[seating[0]][seating[len(seating)-1]]
	return happiness
}

func main() {
	re := regexp.MustCompile("^([a-zA-Z]+) would (gain|lose) ([0-9]+) happiness units by sitting next to ([a-zA-Z]+)\\.$")

	happiness := make(Happiness)

	for {
		line, found := readNextInputLine()
		if !found {
			break
		}
		matches := re.FindStringSubmatch(line)
		if len(matches) <= 0 {
			panic("input did not match expected format")
		}
		
		happy, err := strconv.Atoi(matches[3])
		if err != nil {
			panic("unable to convert happiness to integer: " + err.Error())
		}
		if matches[2] == "lose" {
			happy = -happy
		}
		happiness.AddSeating(matches[1], matches[4], happy)
	}

	people := make([]string, 0, len(happiness)+1)
	for name := range happiness {
		people = append(people, name)
	}
	people = append(people, "Me")
	sort.Strings(people)

	seatings := generateSeatings(people)
	maxHappiness := -9999999
	for i := range seatings {
		if seatingHappiness := happiness.CalculateBasedOnSeating(seatings[i]); seatingHappiness > maxHappiness {
			maxHappiness = seatingHappiness
		}
	}
	
	fmt.Println("Max happiness:", maxHappiness)
}

func factorial(n int) int {
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result
}

func generateSeatings(people []string) [][]string {
	if len(people) < 2 {
		panic("not enough people left to seat")
	}
	if len(people) == 2 {
		return [][]string{
			[]string{people[0], people[1]},
			[]string{people[1], people[0]},
		}
	}
	
	result := make([][]string, factorial(len(people)))
	s := 0
	for i := range people {
		peopleLeft := make([]string, 0, len(people)-1)
		for j := 0; j <= i-1; j++ {
			peopleLeft = append(peopleLeft, people[j])
		}
		for j := i + 1; j < len(people); j++ {
			peopleLeft = append(peopleLeft, people[j])
		}

		seatings := generateSeatings(peopleLeft)
		for seating := range seatings {
			result[s] = append(seatings[seating], people[i])
			s++
		}
	}
	return result
}

var pos = 0

func readNextInputLine() (string, bool) {
	start := pos
	for i := pos + 1; i <= len(input); i++ {
		if i == len(input) || input[i] == '\n' {
			pos = i + 1
			return input[start : pos-1], true
		}
	}
	return "", false
}

var input = `Alice would gain 54 happiness units by sitting next to Bob.
Alice would lose 81 happiness units by sitting next to Carol.
Alice would lose 42 happiness units by sitting next to David.
Alice would gain 89 happiness units by sitting next to Eric.
Alice would lose 89 happiness units by sitting next to Frank.
Alice would gain 97 happiness units by sitting next to George.
Alice would lose 94 happiness units by sitting next to Mallory.
Bob would gain 3 happiness units by sitting next to Alice.
Bob would lose 70 happiness units by sitting next to Carol.
Bob would lose 31 happiness units by sitting next to David.
Bob would gain 72 happiness units by sitting next to Eric.
Bob would lose 25 happiness units by sitting next to Frank.
Bob would lose 95 happiness units by sitting next to George.
Bob would gain 11 happiness units by sitting next to Mallory.
Carol would lose 83 happiness units by sitting next to Alice.
Carol would gain 8 happiness units by sitting next to Bob.
Carol would gain 35 happiness units by sitting next to David.
Carol would gain 10 happiness units by sitting next to Eric.
Carol would gain 61 happiness units by sitting next to Frank.
Carol would gain 10 happiness units by sitting next to George.
Carol would gain 29 happiness units by sitting next to Mallory.
David would gain 67 happiness units by sitting next to Alice.
David would gain 25 happiness units by sitting next to Bob.
David would gain 48 happiness units by sitting next to Carol.
David would lose 65 happiness units by sitting next to Eric.
David would gain 8 happiness units by sitting next to Frank.
David would gain 84 happiness units by sitting next to George.
David would gain 9 happiness units by sitting next to Mallory.
Eric would lose 51 happiness units by sitting next to Alice.
Eric would lose 39 happiness units by sitting next to Bob.
Eric would gain 84 happiness units by sitting next to Carol.
Eric would lose 98 happiness units by sitting next to David.
Eric would lose 20 happiness units by sitting next to Frank.
Eric would lose 6 happiness units by sitting next to George.
Eric would gain 60 happiness units by sitting next to Mallory.
Frank would gain 51 happiness units by sitting next to Alice.
Frank would gain 79 happiness units by sitting next to Bob.
Frank would gain 88 happiness units by sitting next to Carol.
Frank would gain 33 happiness units by sitting next to David.
Frank would gain 43 happiness units by sitting next to Eric.
Frank would gain 77 happiness units by sitting next to George.
Frank would lose 3 happiness units by sitting next to Mallory.
George would lose 14 happiness units by sitting next to Alice.
George would lose 12 happiness units by sitting next to Bob.
George would lose 52 happiness units by sitting next to Carol.
George would gain 14 happiness units by sitting next to David.
George would lose 62 happiness units by sitting next to Eric.
George would lose 18 happiness units by sitting next to Frank.
George would lose 17 happiness units by sitting next to Mallory.
Mallory would lose 36 happiness units by sitting next to Alice.
Mallory would gain 76 happiness units by sitting next to Bob.
Mallory would lose 34 happiness units by sitting next to Carol.
Mallory would gain 37 happiness units by sitting next to David.
Mallory would gain 40 happiness units by sitting next to Eric.
Mallory would gain 18 happiness units by sitting next to Frank.
Mallory would gain 7 happiness units by sitting next to George.`
