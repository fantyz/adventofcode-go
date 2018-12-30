package main

import (
	"container/heap"
	"fmt"
	"strings"
)

/*


 */

func main() {
	fmt.Println("Day 15: Beverage Bandits")
	c := NewCave(puzzleInput)
	fmt.Println(c.String())
}

type Entity uint8

const (
	Nothing Entity = iota
	Wall
	Goblin
	Elf
)

func (e Entity) String() string {
	switch e {
	case Nothing:
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
	c [][]Entity
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
				row[x] = Elf
			case 'G':
				row[x] = Goblin
			default:
				panic("Unknown entity")
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

func AStar(from, to Coord, m [][]Entity) []Coord {
	abs := func(i int) int {
		if i >= 0 {
			return i
		}
		return -1 * i
	}
	heuristicCost := func(from, to Coord) int {
		return abs(from.X-to.X) + abs(from.Y-to.Y)
	}

	completed := map[Coord]struct{}{}

	openSet := []Coord{from}

	cameFrom := map[Coord]Coord{}
	gScore := map[Coord]int{from: 0}
	fScore := map[Coord]int{from: heuristicCost(from, to)}

	for len(openSet) > 0 {

	}

	return nil
}

// priority queue

type Item struct {
	value   Coord
	estDist int
	index   int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index, pq[j].index = i, j
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Item)
	item.index = len(*pq)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	item := (*pq)[len(*pq)-1]
	*pq = (*pq)[:len(*pq)-1]
	return item
}

func (pq *PriorityQueue) update(item *Item, value Coord, estDist int) {
	heap.Fix(pq, item.index)
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
