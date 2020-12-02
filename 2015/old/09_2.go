package main

/*
--- Day 9: All in a Single Night ---

Every year, Santa manages to deliver all of his presents in a single night.

This year, however, he has some new locations to visit; his elves have provided him the distances between every pair of locations. He can start and end at any two (different) locations he wants, but he must visit each location exactly once. What is the shortest distance he can travel to achieve this?

For example, given the following distances:

London to Dublin = 464
London to Belfast = 518
Dublin to Belfast = 141
The possible routes are therefore:

Dublin -> London -> Belfast = 982
London -> Dublin -> Belfast = 605
London -> Belfast -> Dublin = 659
Dublin -> Belfast -> London = 659
Belfast -> Dublin -> London = 605
Belfast -> London -> Dublin = 982
The shortest of these is London -> Dublin -> Belfast = 605, and so the answer is 605 in this example.

What is the distance of the shortest route?

--- Part Two ---

The next year, just to show off, Santa decides to take the route with the longest distance instead.

He can still start and end at any two (different) locations he wants, and he still must visit each location exactly once.

For example, given the distances above, the longest route would be 982 via (for example) Dublin -> London -> Belfast.

What is the distance of the longest route?

*/

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
)

type Routes map[string]map[string]int

func (r Routes) AddRoute(from, to string, distance string) {
	dist, err := strconv.Atoi(distance)
	if err != nil {
		panic("unable to convert distance to integer: " + err.Error())
	}
	if _, found := r[from]; !found {
		r[from] = make(map[string]int)
	}
	r[from][to] = dist
}

func main() {
	re := regexp.MustCompile("^([a-zA-Z]+) to ([a-zA-Z]+) = ([0-9]+)$")

	routes := make(Routes)

	for {
		line, found := readNextInputLine()
		if !found {
			break
		}
		matches := re.FindStringSubmatch(line)
		if len(matches) <= 0 {
			panic("input did not match expected format")
		}

		routes.AddRoute(matches[1], matches[2], matches[3])
		routes.AddRoute(matches[2], matches[1], matches[3])
	}

	destinations := make([]string, 0, len(routes))
	for city := range routes {
		destinations = append(destinations, city)
	}
	sort.Strings(destinations)

	longestRoute := 0
	for i := 0; i < len(destinations); i++ {
		newDestinations := make([]string, 0, len(destinations)-1)
		for j := 0; j <= i-1; j++ {
			newDestinations = append(newDestinations, destinations[j])
		}
		for j := i + 1; j < len(destinations); j++ {
			newDestinations = append(newDestinations, destinations[j])
		}
		r := routes.getShortestDistance(destinations[i], newDestinations)
		if longestRoute < r.Length() {
			longestRoute = r.Length()
		}
	}

	fmt.Println("Longest route:", longestRoute)
}

type dest struct {
	to     string
	length int
}

type path []dest

func (p path) Length() int {
	l := 0
	for i := 0; i < len(p); i++ {
		l += p[i].length
	}
	return l
}

func (r Routes) getShortestDistance(from string, destinations []string) path {
	if len(destinations) == 0 {
		panic("no destinations")
	}

	if len(destinations) == 1 {
		return path{
			dest{
				to:     destinations[0],
				length: r[from][destinations[0]],
			},
		}
	}

	var bestPathLength int
	var bestPath path

	for i := range destinations {
		newDestinations := make([]string, 0, len(destinations)-1)
		for j := 0; j <= i-1; j++ {
			newDestinations = append(newDestinations, destinations[j])
		}
		for j := i + 1; j < len(destinations); j++ {
			newDestinations = append(newDestinations, destinations[j])
		}
		p := r.getShortestDistance(destinations[i], newDestinations)
		p = append(path{dest{
			to:     destinations[i],
			length: r[from][destinations[i]],
		}}, p...)

		if bestPath == nil {
			bestPathLength = p.Length()
			bestPath = p
		} else {
			if bestPathLength < p.Length() {
				bestPathLength = p.Length()
				bestPath = p
			}
		}
	}

	return bestPath
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

var input = `Tristram to AlphaCentauri = 34
Tristram to Snowdin = 100
Tristram to Tambi = 63
Tristram to Faerun = 108
Tristram to Norrath = 111
Tristram to Straylight = 89
Tristram to Arbre = 132
AlphaCentauri to Snowdin = 4
AlphaCentauri to Tambi = 79
AlphaCentauri to Faerun = 44
AlphaCentauri to Norrath = 147
AlphaCentauri to Straylight = 133
AlphaCentauri to Arbre = 74
Snowdin to Tambi = 105
Snowdin to Faerun = 95
Snowdin to Norrath = 48
Snowdin to Straylight = 88
Snowdin to Arbre = 7
Tambi to Faerun = 68
Tambi to Norrath = 134
Tambi to Straylight = 107
Tambi to Arbre = 40
Faerun to Norrath = 11
Faerun to Straylight = 66
Faerun to Arbre = 144
Norrath to Straylight = 115
Norrath to Arbre = 135
Straylight to Arbre = 127`
