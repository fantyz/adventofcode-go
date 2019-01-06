package main

import (
	"fmt"
	"sort"
	"strings"
)

/*


 */

func main() {
	fmt.Println("Day 15: Beverage Bandits")
	c := NewCave(puzzleInput)
	fmt.Println(c.String())
	AStar(Coord{0, 0}, Coord{10, 10}, nil)
}

type Entity uint8

const (
	Empty Entity = iota
	Wall
	Goblin
	Elf
)

func (e Entity) String() string {
	switch e {
	case Empty:
		return "."
	case Wall:
		return "#"
	case Goblin:
		return "G"
	case Elf:
		return "E"
	default:
		panic("unknwon")
	}
}

type Cave struct {
	c      [][]Entity
	actors []*Actor
}

const (
	elfAttack = 3
	elfHP     = 200
	orcAttack = 3
	orcHP     = 200
)

type Actor struct {
	c      Coord
	hp     int
	attack int
}

type Coord struct {
	X, Y int
}

func NewCave(in string) *Cave {
	c := Cave{}
	for _, l := range strings.Split(in, "\n") {
		if l == "" {
			continue
		}

		row := make([]Entity, len(l))
		for x := range l {
			switch l[x] {
			case '.':
			case '#':
				row[x] = Wall
			case 'E':
				c.actors = append(c.actors, &Actor{Coord{x, len(c.c)}, elfHP, elfAttack})
				row[x] = Elf
			case 'G':
				c.actors = append(c.actors, &Actor{Coord{x, len(c.c)}, orcHP, orcAttack})
				row[x] = Goblin
			default:
				panic("Unknown entity: " + string(l[x]))
			}
		}
		c.c = append(c.c, row)
	}
	return &c
}

func (c *Cave) String() string {
	var str string
	for y := 0; y < len(c.c); y++ {
		for x := 0; x < len(c.c[y]); x++ {
			str += c.c[y][x].String()
		}
		str += "\n"
	}
	return str
}

func (c *Cave) Tick() {
	//fmt.Println()
	sort.Slice(c.actors, func(i, j int) bool { return c.actors[i].c.Less(c.actors[j].c) })

	for i, actor := range c.actors {
		e := c.c[actor.c.Y][actor.c.X]

		//fmt.Println(c.String())
		//fmt.Printf("Moving %v %v\n", e, actor)

		enemyClose := false
		for _, n := range []Coord{{actor.c.X + 1, actor.c.Y}, {actor.c.X - 1, actor.c.Y}, {actor.c.X, actor.c.Y + 1}, {actor.c.X, actor.c.Y - 1}} {
			if (e == Elf && c.c[n.Y][n.X] == Goblin) ||
				(e == Goblin && c.c[n.Y][n.X] == Elf) {
				// enemy adjacant
				enemyClose = true
			}
		}
		if enemyClose {
			// attack!

			//fmt.Println(" > Close to enemy, not moving")
			continue
		}

		var targetLoc, nextStep Coord
		dist := -1

		for _, enemy := range c.actors {
			if c.c[enemy.c.Y][enemy.c.X] == e {
				// not an enemy!
				continue
			}

			for _, n := range []Coord{{enemy.c.X + 1, enemy.c.Y}, {enemy.c.X - 1, enemy.c.Y}, {enemy.c.X, enemy.c.Y + 1}, {enemy.c.X, enemy.c.Y - 1}} {
				if c.c[n.Y][n.X] != Empty {
					// can't move into something
					continue
				}

				path := AStar(actor.c, n, c.c)
				if len(path) > 0 && (dist < 0 || len(path) <= dist) {
					//fmt.Println("Possible to move to:", path[0], " (dist=", len(path), ") ; path:", path)
					if len(path) == dist && targetLoc.Less(n) {
						// tagetLoc takes precedence due to being earlier in "read order"
						//fmt.Println("Skipping path due to being later in reading order!")
						continue
					}

					targetLoc = n
					nextStep = path[0]
					dist = len(path)
				}
			}
		}

		if dist < 0 {
			//fmt.Println("No paths found...")
			continue
		}

		//fmt.Printf("Moving... %v to %v (dest=%v)\n", actor, nextStep, targetLoc)
		// move to nextStep
		c.c[nextStep.Y][nextStep.X] = e
		c.c[actor.c.Y][actor.c.X] = Empty
		c.actors[i].c = nextStep
	}

}

func AStar(from, to Coord, m [][]Entity) []Coord {
	abs := func(i int) int {
		if i < 0 {
			return -1 * i
		}
		return i
	}
	simpleDist := func(a, b Coord) int {
		return abs(a.X-b.X) + abs(a.Y-b.Y)
	}

	type node struct {
		c Coord
		d int
	}

	//fmt.Println()
	//fmt.Println("  >>", from, " => ", to)

	closedSet := map[Coord]struct{}{}
	openSet := []node{node{from, simpleDist(from, to)}}
	cameFrom := map[Coord]Coord{}
	gScore := map[Coord]int{from: 0}

	for len(openSet) > 0 {
		// grab the next one that looks to be cloest to where we're going
		sort.Slice(openSet, func(i, j int) bool {
			if openSet[i].d != openSet[j].d {
				return openSet[i].d <= openSet[j].d
			}
			// sort equal distance by "reading order" to give preference to this when multiple equally lenth paths exist
			return openSet[i].c.Y < openSet[j].c.Y || (openSet[i].c.Y == openSet[j].c.Y && openSet[i].c.X < openSet[j].c.X)
		})
		//fmt.Println()
		//fmt.Println("  >>  openSet:", openSet)
		//fmt.Println("  >> cameFrom:", cameFrom)
		cur := openSet[0].c
		if cur == to {
			// done, retrace our steps
			var path []Coord
			for {
				path = append([]Coord{cur}, path...)
				var found bool
				cur, found = cameFrom[cur]
				if !found {
					panic("missing coord from cameFrom when tracing our steps")
				}
				if cur == from {
					//fmt.Println("  >> path:", path)
					return path
				}
			}
		}

		openSet = openSet[1:]
		closedSet[cur] = struct{}{}

		for _, neighbor := range []Coord{{cur.X, cur.Y - 1}, {cur.X - 1, cur.Y}, {cur.X + 1, cur.Y}, {cur.X, cur.Y + 1}} {
			if _, found := closedSet[neighbor]; found {
				// already evaluated
				continue
			}

			if m[neighbor.Y][neighbor.X] != Empty {
				// something is blocking this coordinate, can't go there
				continue
			}

			// distance from start to neighbor (distance from current to neighbor is 1 by design)
			d := gScore[cur] + 1

			idx := sort.Search(len(openSet), func(i int) bool { return openSet[i].c == neighbor })
			if idx == len(openSet) {
				// not found, add it
				openSet = append(openSet, node{neighbor, d + simpleDist(neighbor, to)})
			}

			if gs, found := gScore[neighbor]; found && d >= gs {
				// already found a better path
				continue
			}

			// this path is better than anything previously seen
			cameFrom[neighbor] = cur
			openSet[idx].d = d + simpleDist(neighbor, to)
			gScore[neighbor] = d
		}
	}

	return nil
}

func (c Coord) Less(c2 Coord) bool {
	return c.Y < c2.Y || (c.Y == c2.Y && c.X < c2.X)
}

const puzzleInput = `################################
########.#######################
#######..#######################
######..########################
###....####...##################
###.#..####G..##################
###G#.G#####..####G#############
##....G..###.......#############
#G#####...#..G.....#############
#G.###..#..G........############
#..G.G..........G.....#.G.######
###......GG..G............######
#######....G..#####.G...#.######
#######......#######....########
#######.....#########..........#
#######.....#########.........##
#######...#.#########.........##
#######.....#########........###
#######.....#########.........##
#######....E.#######........#..#
#######.......#####E........####
###.#.E..#.....G.........#..####
###......#E......E..G...E...####
##...........#.............#####
#####.###..............E...#####
#############..............#####
#############..E.....###...#####
###############..E...###...#####
#################.E#.####..#####
#################..#.###########
#################..#.###########
################################`
