package main

/*
--- Day 14: Reindeer Olympics ---

This year is the Reindeer Olympics! Reindeer can fly at high speeds, but must rest occasionally to recover their energy. Santa would like to know which of his reindeer is fastest, and so he has them race.

Reindeer can only either be flying (always at their top speed) or resting (not moving at all), and always spend whole seconds in either state.

For example, suppose you have the following Reindeer:

Comet can fly 14 km/s for 10 seconds, but then must rest for 127 seconds.
Dancer can fly 16 km/s for 11 seconds, but then must rest for 162 seconds.
After one second, Comet has gone 14 km, while Dancer has gone 16 km. After ten seconds, Comet has gone 140 km, while Dancer has gone 160 km. On the eleventh second, Comet begins resting (staying at 140 km), and Dancer continues on for a total distance of 176 km. On the 12th second, both reindeer are resting. They continue to rest until the 138th second, when Comet flies for another ten seconds. On the 174th second, Dancer flies for another 11 seconds.

In this example, after the 1000th second, both reindeer are resting, and Comet is in the lead at 1120 km (poor Dancer has only gotten 1056 km by that point). So, in this situation, Comet would win (if the race ended at 1000 seconds).

Given the descriptions of each reindeer (in your puzzle input), after exactly 2503 seconds, what distance has the winning reindeer traveled?

Your puzzle answer was 2696.

The first half of this puzzle is complete! It provides one gold star: *

--- Part Two ---

Seeing how reindeer move in bursts, Santa decides he's not pleased with the old scoring system.

Instead, at the end of each second, he awards one point to the reindeer currently in the lead. (If there are multiple reindeer tied for the lead, they each get one point.) He keeps the traditional 2503 second time limit, of course, as doing otherwise would be entirely ridiculous.

Given the example reindeer from above, after the first second, Dancer is in the lead and gets one point. He stays in the lead until several seconds into Comet's second burst: after the 140th second, Comet pulls into the lead and gets his first point. Of course, since Dancer had been in the lead for the 139 seconds before that, he has accumulated 139 points by the 140th second.

After the 1000th second, Dancer has accumulated 689 points, while poor Comet, our old champion, only has 312. So, with the new scoring system, Dancer would win (if the race ended at 1000 seconds).

Again given the descriptions of each reindeer (in your puzzle input), after exactly 2503 seconds, how many points does the winning reindeer have?

*/

import (
	"fmt"
	"regexp"
	"strconv"
)

type Raindeer struct {
	Name string
	Speed, Duration, Rest int
	distance int
	resting bool
	timeInState int
	score int
}

func (r *Raindeer) Tick() {
	r.timeInState++

	if !r.resting {
		r.distance += r.Speed
		if r.timeInState == r.Duration {
			r.resting = true
			r.timeInState = 0
		}
	} else {
		if r.timeInState == r.Rest {
			r.resting = false
			r.timeInState = 0
		}
	}
}

func main() {
	re := regexp.MustCompile("^([a-zA-Z]+) can fly ([0-9]+) km/s for ([0-9]+) seconds, but then must rest for ([0-9]+) seconds\\.$")

	raceTime := 2503
	raindeers := []Raindeer{}
	
	for {
		line, found := readNextInputLine()
		if !found {
			break
		}
		matches := re.FindStringSubmatch(line)
		if len(matches) <= 0 {
			panic("input did not match expected format")
		}

		name := matches[1]
		speed, err := strconv.Atoi(matches[2])
		if err != nil {
			panic("could not convert speed to int")
		}
		duration, err := strconv.Atoi(matches[3])
		if err != nil {
			panic("could not convert duration to int")
		}
		rest, err := strconv.Atoi(matches[4])
		if err != nil {
			panic("could not convert rest to int")
		}

		raindeers = append(raindeers, Raindeer{
			Name: name,
			Speed: speed,
			Duration: duration,
			Rest: rest,
		})
	}

	for t:=0; t < raceTime; t++ {
		raindeerInLead := -1
		raindeerInLeadDist := 0
		
		for i:= range raindeers {
			raindeers[i].Tick()
			if raindeers[i].distance > raindeerInLeadDist {
				raindeerInLead = i
				raindeerInLeadDist = raindeers[i].distance
			}
		}
		raindeers[raindeerInLead].score++
	}

	winnerName := ""
	winnerScore := 0
	for i:= range raindeers {
		if raindeers[i].score > winnerScore {
			winnerScore = raindeers[i].score
			winnerName = raindeers[i].Name
		}
	}
	fmt.Println("Winner score:", winnerScore, "by", winnerName)
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

var input = `Rudolph can fly 22 km/s for 8 seconds, but then must rest for 165 seconds.
Cupid can fly 8 km/s for 17 seconds, but then must rest for 114 seconds.
Prancer can fly 18 km/s for 6 seconds, but then must rest for 103 seconds.
Donner can fly 25 km/s for 6 seconds, but then must rest for 145 seconds.
Dasher can fly 11 km/s for 12 seconds, but then must rest for 125 seconds.
Comet can fly 21 km/s for 6 seconds, but then must rest for 121 seconds.
Blitzen can fly 18 km/s for 3 seconds, but then must rest for 50 seconds.
Vixen can fly 20 km/s for 4 seconds, but then must rest for 75 seconds.
Dancer can fly 7 km/s for 20 seconds, but then must rest for 119 seconds.`
