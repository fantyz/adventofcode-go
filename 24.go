package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*

--- Day 24: Air Duct Spelunking ---

You've finally met your match; the doors that provide access to the roof are locked tight, and all of the controls and related electronics are inaccessible. You simply can't reach them.

The robot that cleans the air ducts, however, can.

It's not a very fast little robot, but you reconfigure it to be able to interface with some of the exposed wires that have been routed through the HVAC system. If you can direct it to each of those locations, you should be able to bypass the security controls.

You extract the duct layout for this area from some blueprints you acquired and create a map with the relevant locations marked (your puzzle input). 0 is your current location, from which the cleaning robot embarks; the other numbers are (in no particular order) the locations the robot needs to visit at least once each. Walls are marked as #, and open passages are marked as .. Numbers behave like open passages.

For example, suppose you have a map like the following:

###########
#0.1.....2#
#.#######.#
#4.......3#
###########
To reach all of the points of interest as quickly as possible, you would have the robot take the following path:

0 to 4 (2 steps)
4 to 1 (4 steps; it can't move diagonally)
1 to 2 (6 steps)
2 to 3 (2 steps)
Since the robot isn't very fast, you need to find it the shortest route. This path is the fewest steps (in the above example, a total of 14) required to start at 0 and then visit every other location at least once.

Given your actual map, and starting from location 0, what is the fewest number of steps required to visit every non-0 number marked on the map at least once?

Your puzzle answer was 474.

--- Part Two ---

Of course, if you leave the cleaning robot somewhere weird, someone is bound to notice.

What is the fewest number of steps required to start at 0, visit every non-0 number marked on the map at least once, and then return to 0?

*/

func main() {
	fmt.Println("Advent of code 2016: Day 24")

	l := NewLabyrinth(puzzleInput)
	l.Print()

	dist, route := l.ShortestPath(false)
	fmt.Println("Part 1: Shorest path is", dist, "steps")
	for i := 1; i < len(route); i++ {
		fmt.Println("  >", route[i-1], "=>", route[i], "is", l.dist[route[i-1]][route[i]])
	}

	dist, route = l.ShortestPath(true)
	fmt.Println("Part 2: Shorest path is", dist, "steps")
	for i := 1; i < len(route); i++ {
		fmt.Println("  >", route[i-1], "=>", route[i], "is", l.dist[route[i-1]][route[i]])
	}
}

func NewLabyrinth(layout string) *Labyrinth {
	l := &Labyrinth{
		waypoints: map[int]Pos{},
	}
	// read maze
	for y, line := range strings.Split(layout, "\n") {
		l.walls = append(l.walls, []bool{})
		for x := range line {
			switch line[x] {
			case '#':
				l.walls[y] = append(l.walls[y], true)
			default:
				i, err := strconv.Atoi(string(line[x]))
				if err != nil {
					panic("Unknown character: " + string(line[x]))
				}
				l.waypoints[i] = Pos{x: x, y: y}
				fallthrough
			case '.':
				l.walls[y] = append(l.walls[y], false)
			}
		}
	}

	// determine distance between waypoints
	l.dist = make([][]int, len(l.waypoints))
	for i := 0; i < len(l.dist); i++ {
		l.dist[i] = make([]int, len(l.waypoints))
	}

	for from := range l.waypoints {
		for to := range l.waypoints {
			l.dist[from][to] = l.ShortestAB(from, to)
			l.dist[to][from] = l.dist[from][to]
		}
	}

	return l
}

type Labyrinth struct {
	dist      [][]int // from+to = dist
	walls     [][]bool
	waypoints map[int]Pos
}

func (l *Labyrinth) ShortestPath(endAtStart bool) (int, []int) {
	// skip wp 0 as we know 0 must come first
	wps := []int{}
	for i := 1; i < len(l.waypoints); i++ {
		wps = append(wps, i)
	}

	p := NewPermutator(wps)
	var bestRoute []int
	best := 9999999
	for {
		route := p.Next()
		if route == nil {
			break
		}
		d := l.dist[0][route[0]]
		for i := 1; i < len(route); i++ {
			d += l.dist[route[i-1]][route[i]]
		}
		if endAtStart {
			d += l.dist[route[len(route)-1]][0]
		}

		if d < best {
			best = d
			bestRoute = route
		}
	}

	bestRoute = append([]int{0}, bestRoute...)
	if endAtStart {
		bestRoute = append(bestRoute, 0)
	}
	return best, bestRoute
}

func NewPermutator(list []int) *Permutator {
	org := make([]int, len(list))
	copy(org, list)

	return &Permutator{
		org: org,
		s:   make([]int, len(list)),
	}
}

func (p *Permutator) Next() []int {
	// step 1: find the last digit in p.s that can be incremented
	done := true
	for i := len(p.s) - 1; i >= 0; i-- {
		// either p.s[i] can be incremented or it should be reset to 0
		if p.s[i] < len(p.s)-i-1 {
			p.s[i]++
			done = false
			break
		}
		p.s[i] = 0
	}
	if done {
		return nil
	}

	// step 2: copy p.org and swap according to p.s
	res := make([]int, len(p.org))
	copy(res, p.org)

	for i, v := range p.s {
		res[i], res[i+v] = res[i+v], res[i]
	}

	return res
}

type Permutator struct {
	org []int
	s   []int
}

func (l *Labyrinth) ShortestAB(from, to int) int {
	if from == to {
		return 0
	}

	visited := map[Pos]struct{}{}
	next := []Pos{l.waypoints[from]}
	visitNext := func() {
		newNext := []Pos{}
		for _, p := range next {
			if p == l.waypoints[to] {
				// done
				next = nil
				return
			}

			up := p
			up.y += 1
			if _, been := visited[up]; !been && up.y < len(l.walls) && !l.walls[up.y][up.x] {
				newNext = append(newNext, up)
				visited[up] = struct{}{}
			}

			down := p
			down.y -= 1
			if _, been := visited[down]; !been && down.y >= 0 && !l.walls[down.y][down.x] {
				newNext = append(newNext, down)
				visited[down] = struct{}{}
			}

			right := p
			right.x += 1
			if _, been := visited[right]; !been && right.x < len(l.walls[0]) && !l.walls[right.y][right.x] {
				newNext = append(newNext, right)
				visited[right] = struct{}{}
			}

			left := p
			left.x -= 1
			if _, been := visited[left]; !been && left.x >= 0 && !l.walls[left.y][left.x] {
				newNext = append(newNext, left)
				visited[left] = struct{}{}
			}
		}
		if len(newNext) <= 0 {
			panic("No solution found")
		}
		next = newNext
	}

	steps := 0
	for len(next) > 0 {
		steps++
		visitNext()
	}

	return steps - 1
}

func (l *Labyrinth) Print() {
	for y := 0; y < len(l.walls); y++ {
		for x := 0; x < len(l.walls[y]); x++ {
			found := false
			for i, p := range l.waypoints {
				if p.x == x && p.y == y {
					found = true
					fmt.Printf("%d", i)
				}
			}
			if !found {
				if l.walls[y][x] {
					fmt.Print("#")
				} else {
					fmt.Print(".")
				}
			}
		}
		fmt.Println()
	}
}

type Pos struct {
	x, y int
}

const testpuzzleInput = `###########
#0.1.....2#
#.#######.#
#4.......3#
###########`

const puzzleInput = `#######################################################################################################################################################################################
#.....#.....#.....#...#...#.#...#...............................#.#.......#.#...#...........#...............#.#...#.....#.....#.....#...................#.......#.#.......#.....#...#.#
###.#.#.###.#.#.#.#.###.#.#.###.#.#.#.#.###.###.#.#.#.#.#.#.###.#.#.###.#.#.#####.###.#.###.#.###.###.#.#.###.#.###.#.#.#.###.###.###.#.#.#.#.#.#.#.#.#.#.#.###.#.#.#.#.#.#.#.#.#.###.#
#.#......4#.#.....#...#...#.#...........#...#.#.#...#.#...#...#.#.........#.#.#.........#...........#.............#.....#...#...#.#...#...#...#....3#.....#.....#.............#.......#
#.#####.###.#.#####.#.#.#.#.#.#####.###.###.#.#.###.#.#####.#.#.#.###.#.#.#.#.#.#.#######.#.#.#.#.###.#####.###.#.#.###.###.#.#.#.#.#####.###.###.#.#.###.#.#####.#####.###.#####.###.#
#.#...#.....#.#.#...#.#.#.#.#...........#...#...#.....#...#...#.........#.............#.....#...#...#.........#.#...#.......#.....#.....#...#...#...........#.....#.....#...#...#.#...#
###.#.#.###.#.#.#.#.#.#.#.#.#####.###.###.#.#.###.#######.#.#.###.#.#.###.#######.###.#.#######.#.#####.#.###.#.#####.#.#.#.#.###.#.###.#.#.#.###.###.#.#.###.#.#.###.#.#.#.#.#.###.#.#
#...#.......#.#.#.........#...#.....#.#.....#...#...........#.#...#.....#.#.......................#.......#...#.............#...#...#...#.........#.....#.#...#...#.#.....#.....#.#...#
#.#.###.#.#.#.#.#.#####.###.#.#.###.#.#.#.#.###.###.#####.#.#####.#.###.#.#.#.#.###.#.###.#.#####.#.###.#.#.#.#.#.###.#.#.#.#.#.#####.#.#########.#####.#.#.#.#.#.#.#.###.#####.#.#.###
#...#...#.#.#...........#...#...#...#.........#.#...#.....#...#...#.#...#.....#.....#.....#...#...........#.#.#...#...#.....#.....#.#.......#...........#.#.......#...#.....#...#.#.#.#
#.#.###.#.#.#.#.###.#.###.#.#.#.###.###.#.###.#.#.#.#.#.#.#.#.#.###.#.###.#.#####.#.#.#.#.#####.#.###.###.#.#.#.#####.#####.#.#.#.#.###.#.#.#.#########.#.###.#.###.#.#.#.#.#.#.#.###.#
#...#...#...#.#...#...#.......#...............#.#...#.........#...#.#...#...#.............#.#...#.#.......#.....#...#...#.....#.#...#.........#6..#...#.....#.#.....#...#...#...#.....#
#.#.#.#######.#.#.#.#.#.#####.#.#####.#.#.#.###.#.#.#.#.#.#.#.#.#######.#.###.#.###.#####.#.#.#.#.#.#.#####.#######.#.###.#.#.#.#.#.#.#.###.#####.#.#.###.#####.#####.#.#####.#.#.#.###
#.....#.....#...#.#...#.#...........#...#...#.......#.....#.#.#...#.#.....#.....#.....#.......#.....#.#...#.#...#.......#.........#...#.......#...#.#.....#...#.....#.#.#.#.#.....#...#
###.#.###.#.#.###.#.#.#.#.###.###.#.###.#.#.#.#.###.#.###.#######.#.#.###.#.#.#.#.#.#.#.#######.###.###.#.#.#.#.#.#.#.###.#.#.#.###.#.###.###.#.#.###.#####.#.#.#.#.#.#.#.#.#.###.#####
#.#.#.#.....#...#.#...#...#...#.........#.....#.....#.#...#.......#...#.....#...#...#...#...#.....#...#.#...#.........#.....#.....#.......#.....#.#.....#...#...#...#.#.....#.........#
#.###.#.#.###.#.#.#.#.#.#.###.#####.###.#.###.#.#.#.#.#.#.#.#####.#.#.#.###.#.#.#####.#####.#.#####.#.#.#######.#.#.#.#.#.#.###.#.#.#.#.#.#######.#.###.#.#####.###.###.#.#.#.#.#.###.#
#.....#.#5#...#...#...#.#.....#.........#.....#.#.......#...#.......#.#.......#.#.#...........#.......#...#.#.....#.#.........#...#.#.....#.......#..0#.........#...#.....#.#.....#.#.#
#.###.#.###.#.#.#.#.#.#.###.###.#####.#########.#.#.#.#.###.#.#.###.#.#.#.#.#.###.#.###.#####.#.#.#####.#.#.#.#.###.#######.#####.#########.#.#.#.#.#.#.###.#.#.#.#.#.###.#.#####.#.#.#
#.....#.........#...#...#.........#.....#.......#...#.#.......#.....#.#...#.#.#...#.#.#.....#...............#.#.....#.......#.#.......#.......#.#.....#.#.......#.........#.......#...#
#.###.#.#.#.###.#.###.#########.#.###.#.#.###.#.#.#####.#.#.#.#####.#.#.#.#.###.#.#.#.#.#.###.#.#.#.#.#.###.#.#.#####.#.#.###.#.###.#.#.#####.#.#.#####.#.#.###.#########.#.#.#######.#
#.....#.......................#.#.#...#...#.............#.#.....#...........#...#.#...#.........#.#...#.....#...........#.#.#.......#.#...#.......#.........#.................#.#.....#
#.#.#.#.#.#.#.###.#######.#.#.#.###.#.#.###.#.#.#.#.#.#.#.#.###.#.#.#.#####.#.#.#.#.#.#######.###.#.#.###.#.#######.###.#.#.#####.#.#.#.#.#.#.#.#.#.#.#.#######.#.#####.#.#####.###.#.#
#...#...#...#.......#.#...#.#.#...#.#...........#.#.......#.....#.......#.......#.............#...#.#.#...#.......#.#.....#...#...#.....#.......#.....#...#.....#.........#.....#.#...#
#.#.###.#.#.#.###.#.#.#.#.#.#.###.#.###.###.###.#.#.###.#.#.#.#.#.#.###.#####.#.#.#.###.#.#.#.#.#.#.#.#.#.###.###.#.#.#.#.#.#.#.#.#.###.#.#.#.#.###.###.#.#.###.#####.#.#.#.###.#.#.#.#
#...............#.....#...#...#.....#.........#.....#...#...#.#.........#...........#.#.....#.#.....#...#...#...#.#.........#.....#...#.............#1#...#.#.........#.#...#.......#.#
#####.###.#.#.###.#.#####.#.#.#.###.#.#.#.#.#.#.#.#.#.#.###.###.###.#.###.#.#.###.###.#.#.#.###.#.#.#.###.###.#.#.#.#.#.#.#.###.#.#.#.#.#####.###.#.#.#.#.#.#######.#.#.#.###.#.#.#.###
#.#...........#.....#...#...#.#.........#...#...#.......#.....#.#.....#.........#.#...#...#...#.......#...#.......#.#.#...#.......#...#...............#.........#.......#.#...#.....#.#
#.#.#.#.#.###.#.#.#.#.#.#.#.#.###.###.###.#.#.#.#.#######.#.#.###.#####.###.###.#.#.#.#######.#######.#.###.###.#.#.#.#.#.#.#.#.#.###.###.#.#.#####.#.###.#.#####.#########.###.#.#.#.#
#...........#.....#.......#...#...#...#7..#.....#.......#.....#...#.....#.......#.#...#...#.....#.......#.#.#.....#...............#...#.#...#.#...#.......#.#.#.#...#...#.#.#.....#.#.#
###.###.###.#####.#.#########.#####.#.###.#######.###.#.###.#.###.###.#.###.#.#.#.###.#.#.#.#.#.#.#.#.#.#.#.#####.#.###.#.#####.#.#.#.#.###.#.#.#.#.###.#.#.#.#.###.#.#.#.#####.#.#.#.#
#...#...#.#.......#...........#.......#.....#.#.......#.....#.#...#...#...#...#.......#...#.#.........#.............#...........#.#...#.#...#.#...#...#...#...#...#.#.....#.......#...#
#.###.#.#.#####.#.###.#.#.#####.#.#.#.#.###.#.#.###.#.#.#.###.#.#.#.#.#######.#.###.###.#.###.#####.#.#.#.###.#.#.#.#####.#.#.#.#####.#.#.###.#.#.#.#.#.###.#.###.###.#.#.#.#.#.#.#.###
#.#.....#.#.#.........#...........#...............#...#.#.#.#.#...#...#.........#...#...#.#.#.#...#...#.......#.#.....#.....#...#.#...#.#...........#.....#.#.....#...#.........#.....#
#.#.###.#.#.#.#.#.#.#.#####.#.#.#.#.#.#####.###.###.###.###.###.###.#############.###.###.#.#.#.#.#.#.#.###.#.#.#####.#.#.#######.###.#.#.#.###.#.#.#.#.###.#######.#.#.#.#.###.#####.#
#...#.#.#.#...#.#...#.....#...........#.#...#.......#...........#...#.....#.#...........#.#.....#...#.#...#...#...............#...#...#.#...#...#...#...#.....#2......#...#.#.#.#.....#
#######################################################################################################################################################################################`
