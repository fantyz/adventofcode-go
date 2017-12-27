package main

/*

--- Day 22: Sporifica Virus ---
Diagnostics indicate that the local grid computing cluster has been contaminated with the Sporifica Virus. The grid computing cluster is a seemingly-infinite two-dimensional grid of compute nodes. Each node is either clean or infected by the virus.

To prevent overloading the nodes (which would render them useless to the virus) or detection by system administrators, exactly one virus carrier moves through the network, infecting or cleaning nodes as it moves. The virus carrier is always located on a single node in the network (the current node) and keeps track of the direction it is facing.

To avoid detection, the virus carrier works in bursts; in each burst, it wakes up, does some work, and goes back to sleep. The following steps are all executed in order one time each burst:

If the current node is infected, it turns to its right. Otherwise, it turns to its left. (Turning is done in-place; the current node does not change.)
If the current node is clean, it becomes infected. Otherwise, it becomes cleaned. (This is done after the node is considered for the purposes of changing direction.)
The virus carrier moves forward one node in the direction it is facing.
Diagnostics have also provided a map of the node infection status (your puzzle input). Clean nodes are shown as .; infected nodes are shown as #. This map only shows the center of the grid; there are many more nodes beyond those shown, but none of them are currently infected.

The virus carrier begins in the middle of the map facing up.

For example, suppose you are given a map like this:

..#
#..
...
Then, the middle of the infinite grid looks like this, with the virus carrier's position marked with [ ]:

. . . . . . . . .
. . . . . . . . .
. . . . . . . . .
. . . . . # . . .
. . . #[.]. . . .
. . . . . . . . .
. . . . . . . . .
. . . . . . . . .
The virus carrier is on a clean node, so it turns left, infects the node, and moves left:

. . . . . . . . .
. . . . . . . . .
. . . . . . . . .
. . . . . # . . .
. . .[#]# . . . .
. . . . . . . . .
. . . . . . . . .
. . . . . . . . .
The virus carrier is on an infected node, so it turns right, cleans the node, and moves up:

. . . . . . . . .
. . . . . . . . .
. . . . . . . . .
. . .[.]. # . . .
. . . . # . . . .
. . . . . . . . .
. . . . . . . . .
. . . . . . . . .
Four times in a row, the virus carrier finds a clean, infects it, turns left, and moves forward, ending in the same place and still facing up:

. . . . . . . . .
. . . . . . . . .
. . . . . . . . .
. . #[#]. # . . .
. . # # # . . . .
. . . . . . . . .
. . . . . . . . .
. . . . . . . . .
Now on the same node as before, it sees an infection, which causes it to turn right, clean the node, and move forward:

. . . . . . . . .
. . . . . . . . .
. . . . . . . . .
. . # .[.]# . . .
. . # # # . . . .
. . . . . . . . .
. . . . . . . . .
. . . . . . . . .
After the above actions, a total of 7 bursts of activity had taken place. Of them, 5 bursts of activity caused an infection.

After a total of 70, the grid looks like this, with the virus carrier facing up:

. . . . . # # . .
. . . . # . . # .
. . . # . . . . #
. . # . #[.]. . #
. . # . # . . # .
. . . . . # # . .
. . . . . . . . .
. . . . . . . . .
By this time, 41 bursts of activity caused an infection (though most of those nodes have since been cleaned).

After a total of 10000 bursts of activity, 5587 bursts will have caused an infection.

Given your actual map, after 10000 bursts of activity, how many bursts cause a node to become infected? (Do not count nodes that begin infected.)

Your puzzle answer was 5196.

--- Part Two ---
As you go to remove the virus from the infected nodes, it evolves to resist your attempt.

Now, before it infects a clean node, it will weaken it to disable your defenses. If it encounters an infected node, it will instead flag the node to be cleaned in the future. So:

Clean nodes become weakened.
Weakened nodes become infected.
Infected nodes become flagged.
Flagged nodes become clean.
Every node is always in exactly one of the above states.

The virus carrier still functions in a similar way, but now uses the following logic during its bursts of action:

Decide which way to turn based on the current node:
If it is clean, it turns left.
If it is weakened, it does not turn, and will continue moving in the same direction.
If it is infected, it turns right.
If it is flagged, it reverses direction, and will go back the way it came.
Modify the state of the current node, as described above.
The virus carrier moves forward one node in the direction it is facing.
Start with the same map (still using . for clean and # for infected) and still with the virus carrier starting in the middle and facing up.

Using the same initial state as the previous example, and drawing weakened as W and flagged as F, the middle of the infinite grid looks like this, with the virus carrier's position again marked with [ ]:

. . . . . . . . .
. . . . . . . . .
. . . . . . . . .
. . . . . # . . .
. . . #[.]. . . .
. . . . . . . . .
. . . . . . . . .
. . . . . . . . .
This is the same as before, since no initial nodes are weakened or flagged. The virus carrier is on a clean node, so it still turns left, instead weakens the node, and moves left:

. . . . . . . . .
. . . . . . . . .
. . . . . . . . .
. . . . . # . . .
. . .[#]W . . . .
. . . . . . . . .
. . . . . . . . .
. . . . . . . . .
The virus carrier is on an infected node, so it still turns right, instead flags the node, and moves up:

. . . . . . . . .
. . . . . . . . .
. . . . . . . . .
. . .[.]. # . . .
. . . F W . . . .
. . . . . . . . .
. . . . . . . . .
. . . . . . . . .
This process repeats three more times, ending on the previously-flagged node and facing right:

. . . . . . . . .
. . . . . . . . .
. . . . . . . . .
. . W W . # . . .
. . W[F]W . . . .
. . . . . . . . .
. . . . . . . . .
. . . . . . . . .
Finding a flagged node, it reverses direction and cleans the node:

. . . . . . . . .
. . . . . . . . .
. . . . . . . . .
. . W W . # . . .
. .[W]. W . . . .
. . . . . . . . .
. . . . . . . . .
. . . . . . . . .
The weakened node becomes infected, and it continues in the same direction:

. . . . . . . . .
. . . . . . . . .
. . . . . . . . .
. . W W . # . . .
.[.]# . W . . . .
. . . . . . . . .
. . . . . . . . .
. . . . . . . . .
Of the first 100 bursts, 26 will result in infection. Unfortunately, another feature of this evolved virus is speed; of the first 10000000 bursts, 2511944 will result in infection.

Given your actual map, after 10000000 bursts of activity, how many bursts cause a node to become infected? (Do not count nodes that begin infected.)

Your puzzle answer was 2511633.

*/

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Advent of Code 2017 - Day 22")
	fmt.Println("Infections caused by 10000 bursts:           ", NewInfectedNodes(puzzle).Run(10000).Infections())
	fmt.Println("Evolved infections caused by 10000000 bursts:", NewEvolvedInfectedNodes(puzzle).Run(10000000).Infections())
}

func NewInfectedNodes(in string) *InfectedNodes {
	nodes := &InfectedNodes{
		nodes: map[string]struct{}{},
		dy:    1,
	}
	for i, line := range strings.Split(in, "\n") {
		size := len(line)
		y := (size / 2) - i
		for j := 0; j < len(line); j++ {
			x := j - (size / 2)
			if line[j] == '#' {
				nodes.nodes[Coord(x, y)] = struct{}{}
			}
		}
	}
	return nodes
}

type InfectedNodes struct {
	nodes      map[string]struct{}
	x, y       int
	dx, dy     int
	infections int
}

func (n *InfectedNodes) Run(i int) *InfectedNodes {
	for j := 0; j < i; j++ {
		n.Burst()
	}
	return n
}

func (n *InfectedNodes) Burst() {
	coord := Coord(n.x, n.y)
	if _, found := n.nodes[coord]; found {
		n.dx, n.dy = TurnRight(n.dx, n.dy)
		delete(n.nodes, coord)
	} else {
		n.dx, n.dy = TurnLeft(n.dx, n.dy)
		n.infections++
		n.nodes[coord] = struct{}{}
	}
	n.x, n.y = n.x+n.dx, n.y+n.dy
}

func (n *InfectedNodes) Infections() int {
	return n.infections
}

func NewEvolvedInfectedNodes(in string) *EvolvedInfectedNodes {
	nodes := &EvolvedInfectedNodes{
		nodes: map[string]byte{},
		dy:    1,
	}
	for i, line := range strings.Split(in, "\n") {
		size := len(line)
		y := (size / 2) - i
		for j := 0; j < len(line); j++ {
			x := j - (size / 2)
			if line[j] == '#' {
				nodes.nodes[Coord(x, y)] = 'I'
			}
		}
	}
	return nodes
}

type EvolvedInfectedNodes struct {
	nodes      map[string]byte
	x, y       int
	dx, dy     int
	infections int
}

func (n *EvolvedInfectedNodes) Run(i int) *EvolvedInfectedNodes {
	for j := 0; j < i; j++ {
		n.Burst()
	}
	return n
}

func (n *EvolvedInfectedNodes) Burst() {
	coord := Coord(n.x, n.y)

	switch n.nodes[coord] {
	case 0:
		n.dx, n.dy = TurnLeft(n.dx, n.dy)
		n.nodes[coord] = 'W'
	case 'W':
		n.infections++
		n.nodes[coord] = 'I'
	case 'I':
		n.dx, n.dy = TurnRight(n.dx, n.dy)
		n.nodes[coord] = 'F'
	case 'F':
		n.dx, n.dy = Reverse(n.dx, n.dy)
		delete(n.nodes, coord)
	default:
		panic("Unknown state: " + string(n.nodes[coord]))
	}

	n.x, n.y = n.x+n.dx, n.y+n.dy
}

func (n *EvolvedInfectedNodes) Infections() int {
	return n.infections
}

func Coord(x, y int) string {
	return fmt.Sprintf("%d,%d", x, y)
}

func TurnRight(dx, dy int) (int, int) {
	switch {
	case dy > 0:
		return 1, 0
	case dx > 0:
		return 0, -1
	case dy < 0:
		return -1, 0
	case dx < 0:
		return 0, 1
	}
	panic("No facing")

}

func TurnLeft(dx, dy int) (int, int) {
	switch {
	case dy > 0:
		return -1, 0
	case dx > 0:
		return 0, 1
	case dy < 0:
		return 1, 0
	case dx < 0:
		return 0, -1
	}
	panic("No facing")
}

func Reverse(dx, dy int) (int, int) {
	return -1 * dx, -1 * dy
}

const puzzle = `.#.....##..##..###.###..#
..##..######.#.###.##.#.#
###..#..#####.##.##.#...#
###......##..###.#...#.#.
.#.###.##..#.####.#..#...
..#.#.#####...##.####.###
..#..#.#..###.#..###.###.
#########...#....##..#.#.
.###..#######..####...###
#####...#..##...###..##..
..#......##.#....#...####
.##.#..#####.#####.##.##.
####.##.###.#..#.#.#.....
#....##.####.#.#..#.#.##.
###...##...#.###.#.#.####
.#.#...#.#.##.##....##.#.
#..##.#.#..#....###..####
#####...#..#.###...##.###
##.#..####.###...#....###
###.#####.....#....#.##..
####.##.....######.#..#.#
.#.....####.##...###..##.
....########.#..###.#..##
##.##..#...#...##.#....##
.#.######.##....####.#.##`
