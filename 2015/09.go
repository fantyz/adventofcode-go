package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

func init() { days["9"] = Day9 }

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

Your puzzle answer was 117.

--- Part Two ---
The next year, just to show off, Santa decides to take the route with the longest distance instead.

He can still start and end at any two (different) locations he wants, and he still must visit each location exactly once.

For example, given the distances above, the longest route would be 982 via (for example) Dublin -> London -> Belfast.

What is the distance of the longest route?

Your puzzle answer was 909.

*/

// IMPROVEMENT IDEA: Represent locations with integers (and provide a lookup from integer to
// location name). This allow Distances be a [][]int instead of a map[string]map[string]int. That
// should speed up the lookups as well as reduce the extensive amount of []string fiddling to
// use cheaper []int.

func Day9() {
	fmt.Println("--- Day 9: All in a Single Night ---")
	dists, err := NewDistances(day09Input)
	if err != nil {
		fmt.Println(errors.Wrap(err, "failed to load distances"))
		return
	}
	_, shortestDist := dists.ShortestRouteVisitingAll()
	fmt.Println("Shortest route visiting all locations exactly once:", shortestDist)
	_, longestDist := dists.LongestRouteVisitingAll()
	fmt.Println("Longest route visiting all locations exactly once:", longestDist)
}

// Distances is a map that enables lookup by city to city giving the distance between the two.
type Distances map[string]map[string]int

// NewDistances takes the puzzle input and returns Distances.
// NewDistances will return an error if the input is malformed.
func NewDistances(input string) (Distances, error) {
	dists := Distances{}

	distExp := regexp.MustCompile(`^([a-zA-Z]+) to ([a-zA-Z]+) = ([0-9]+)$`)
	for _, line := range strings.Split(input, "\n") {
		m := distExp.FindStringSubmatch(line)
		if len(m) != 4 {
			return nil, errors.Errorf("unknown line (line=%s)", line)
		}
		if _, found := dists[m[1]]; !found {
			dists[m[1]] = map[string]int{}
		}
		if _, found := dists[m[2]]; !found {
			dists[m[2]] = map[string]int{}
		}
		dist, err := strconv.Atoi(m[3])
		if err != nil {
			return nil, errors.Wrapf(err, "unable to parse distance (line=%s)", line)
		}

		// allow looking up the distance from either direction
		dists[m[1]][m[2]] = dist
		dists[m[2]][m[1]] = dist
	}

	return dists, nil
}

type RouteType int

const (
	ShortestRouteType = iota
	LongestRouteType
)

// ShortestRouteVisitingAll returns the shortest route that visits all locations exactly once.
func (d Distances) ShortestRouteVisitingAll() ([]string, int) {
	return d.findRouteVisitingAllRecursive(ShortestRouteType, "", d.getLocations())
}

// LongestRouteVisitingAll returns the longest route that visits all locations exactly once.
func (d Distances) LongestRouteVisitingAll() ([]string, int) {
	return d.findRouteVisitingAllRecursive(LongestRouteType, "", d.getLocations())
}

// getLocations returns all possible locations
func (d Distances) getLocations() []string {
	locations := make([]string, 0, len(d))
	for loc := range d {
		locations = append(locations, loc)
	}
	return locations
}

// findRouteVisitingAllRecursive is a helper function to FindRouteVisitingAll. It will return the shortest
// or longest route found in reverse ending with the from location along with the distance of this route.
// It takes a location and starting with that location tries all possible routes to find the shortest one
// using itself recursively.
// If an empty string is provided as location, it will try all possible starting locations to find the
// shortest route.
//
// NOTE: The caller must reverse the route to get a route that starts with the from location. If the starting
// location doesn't matter the reverse route is just as good as the non-reversed as the distance of the route
// will remain the same.
func (d Distances) findRouteVisitingAllRecursive(routeType RouteType, from string, remainingLocations []string) ([]string, int) {
	if len(remainingLocations) <= 0 {
		if from == "" {
			return []string{}, 0
		}
		return []string{from}, 0
	}
	if len(remainingLocations) == 1 {
		// only one possible route
		if from == "" {
			return []string{remainingLocations[0]}, 0
		}
		return []string{remainingLocations[0], from}, d[from][remainingLocations[0]]
	}

	// using each possible remainingLocations as the next from location and call itself recurively
	var bestRoute []string
	bestRouteDist := -1
	for i := range remainingLocations {
		newRemainingLocations := make([]string, len(remainingLocations)-1)
		copy(newRemainingLocations, remainingLocations[:i])
		copy(newRemainingLocations[i:], remainingLocations[i+1:])
		route, dist := d.findRouteVisitingAllRecursive(routeType, remainingLocations[i], newRemainingLocations)

		if from != "" {
			// adjust distance to include the distance between from and the first location in the route
			dist += d[from][remainingLocations[i]]
		}

		var isBetter bool
		switch routeType {
		case ShortestRouteType:
			isBetter = bestRouteDist < 0 || dist < bestRouteDist
		case LongestRouteType:
			isBetter = bestRouteDist < 0 || dist > bestRouteDist
		default:
			panic("unsupported route type")
		}

		if isBetter {
			bestRoute = route
			bestRouteDist = dist
		}
	}

	if from != "" {
		// include from at the end of the best route found
		bestRoute = append(bestRoute, from)
	}

	return bestRoute, bestRouteDist
}
