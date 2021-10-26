package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

func init() { days["14"] = Day14 }

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

Your puzzle answer was 2655.

--- Part Two ---
Seeing how reindeer move in bursts, Santa decides he's not pleased with the old scoring system.

Instead, at the end of each second, he awards one point to the reindeer currently in the lead. (If there are multiple reindeer tied for the lead, they each get one point.) He keeps the traditional 2503 second time limit, of course, as doing otherwise would be entirely ridiculous.

Given the example reindeer from above, after the first second, Dancer is in the lead and gets one point. He stays in the lead until several seconds into Comet's second burst: after the 140th second, Comet pulls into the lead and gets his first point. Of course, since Dancer had been in the lead for the 139 seconds before that, he has accumulated 139 points by the 140th second.

After the 1000th second, Dancer has accumulated 689 points, while poor Comet, our old champion, only has 312. So, with the new scoring system, Dancer would win (if the race ended at 1000 seconds).

Again given the descriptions of each reindeer (in your puzzle input), after exactly 2503 seconds, how many points does the winning reindeer have?

Your puzzle answer was 1059.
*/

func Day14() {
	fmt.Println("--- Day 14: Reindeer Olympics ---")
	raindeers, err := NewRaindeers(day14Input)
	if err != nil {
		fmt.Println(errors.Wrap(err, "failed to load raindeers"))
		return
	}
	_, dist := raindeers.Race(2503)
	fmt.Println("After a race of 2503 seconds the winning raindeer has traveled a distance of:", dist)
	_, points := raindeers.PointsRace(2503)
	fmt.Println("After a points race of 2503 seconds the winning raindeer has a score of:", points)
}

// NewRaindeers takes the puzzle input and returns []Raindeers.
func NewRaindeers(input string) (Raindeers, error) {
	var raindeers Raindeers
	raindeerExp := regexp.MustCompile(`^([A-Za-z]+) can fly ([0-9]+) km/s for ([0-9]+) seconds, but then must rest for ([0-9]+) seconds\.$`)
	for _, line := range strings.Split(input, "\n") {
		m := raindeerExp.FindStringSubmatch(line)
		if len(m) != 5 {
			fmt.Println(m)
			return nil, errors.Errorf("line did not match raindeerExp (line=%s)", line)
		}
		name := m[1]
		speed, err := strconv.Atoi(m[2])
		if err != nil {
			// should never happen
			panic(fmt.Sprintf("invalid speed (line=%s, speed=%s)", line, m[2]))
		}
		travelTime, err := strconv.Atoi(m[3])
		if err != nil {
			// should never happen
			panic(fmt.Sprintf("invalid travelTime (line=%s, travelTime=%s)", line, m[3]))
		}
		restTime, err := strconv.Atoi(m[4])
		if err != nil {
			// should never happen
			panic(fmt.Sprintf("invalid restTime (line=%s, restTime=%s)", line, m[4]))
		}

		raindeers = append(raindeers, Raindeer{
			Name:       name,
			Speed:      speed,
			TravelTime: travelTime,
			RestTime:   restTime,
		})
	}
	return raindeers, nil
}

// Raindeer represents the stats of a single raindeer.
type Raindeer struct {
	Name       string
	Speed      int
	TravelTime int
	RestTime   int
}

// Raindeers represents a group of raindeers.
type Raindeers []Raindeer

// Race races the raindeers for the duration in seconds specified. The winning raindeers name
// along with the distance traveled is returned. In case of a tie only one of the winning
// raindeers names is returned.
func (r Raindeers) Race(seconds int) (string, int) {
	bestDist := 0
	bestRaindeer := ""
	for i := range r {
		// calculate the distance travled
		// figure out how many cycles of travel...
		cycles := seconds / (r[i].TravelTime + r[i].RestTime)
		// ...plus how much of the last incomplete cycle (if any) is travel time
		remainder := seconds % (r[i].TravelTime + r[i].RestTime)
		if remainder > r[i].TravelTime {
			remainder = r[i].TravelTime
		}
		dist := cycles*r[i].TravelTime*r[i].Speed + remainder*r[i].Speed
		if dist > bestDist {
			bestRaindeer = r[i].Name
			bestDist = dist
		}
	}
	return bestRaindeer, bestDist
}

// PointsRace races the raindeers in a points-race for the duration in seconds specified. The
// winning raindeer along with the points scored is returned. In case of a tie, only one of the
// winning raindeers names is returned.
func (r Raindeers) PointsRace(seconds int) (string, int) {
	scores := make([]int, len(r))
	dist := make([]int, len(r))

	for sec := 0; sec < seconds; sec++ {
		highDist := 0
		for i := range r {
			// determine whether the raindeer is moving or resting
			if sec%(r[i].TravelTime+r[i].RestTime) < r[i].TravelTime {
				dist[i] += r[i].Speed
			}
			if dist[i] > highDist {
				highDist = dist[i]
			}
		}

		// increment the score of leading raindeer(s)
		for i := range r {
			if dist[i] == highDist {
				scores[i]++
			}
		}
	}

	raindeer := ""
	highscore := 0
	for i := range scores {
		if scores[i] > highscore {
			raindeer = r[i].Name
			highscore = scores[i]
		}
	}

	return raindeer, highscore
}
