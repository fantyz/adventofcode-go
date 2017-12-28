package main

/*

--- Day 24: Electromagnetic Moat ---
The CPU itself is a large, black building surrounded by a bottomless pit. Enormous metal tubes extend outward from the side of the building at regular intervals and descend down into the void. There's no way to cross, but you need to get inside.

No way, of course, other than building a bridge out of the magnetic components strewn about nearby.

Each component has two ports, one on each end. The ports come in all different types, and only matching types can be connected. You take an inventory of the components by their port types (your puzzle input). Each port is identified by the number of pins it uses; more pins mean a stronger connection for your bridge. A 3/7 component, for example, has a type-3 port on one side, and a type-7 port on the other.

Your side of the pit is metallic; a perfect surface to connect a magnetic, zero-pin port. Because of this, the first port you use must be of type 0. It doesn't matter what type of port you end with; your goal is just to make the bridge as strong as possible.

The strength of a bridge is the sum of the port types in each component. For example, if your bridge is made of components 0/3, 3/7, and 7/4, your bridge has a strength of 0+3 + 3+7 + 7+4 = 24.

For example, suppose you had the following components:

0/2
2/2
2/3
3/4
3/5
0/1
10/1
9/10
With them, you could make the following valid bridges:

0/1
0/1--10/1
0/1--10/1--9/10
0/2
0/2--2/3
0/2--2/3--3/4
0/2--2/3--3/5
0/2--2/2
0/2--2/2--2/3
0/2--2/2--2/3--3/4
0/2--2/2--2/3--3/5
(Note how, as shown by 10/1, order of ports within a component doesn't matter. However, you may only use each port on a component once.)

Of these bridges, the strongest one is 0/1--10/1--9/10; it has a strength of 0+1 + 1+10 + 10+9 = 31.

What is the strength of the strongest bridge you can make with the components you have available?

Your puzzle answer was 1906.

--- Part Two ---
The bridge you've built isn't long enough; you can't jump the rest of the way.

In the example above, there are two longest bridges:

0/2--2/2--2/3--3/4
0/2--2/2--2/3--3/5
Of them, the one which uses the 3/5 component is stronger; its strength is 0+2 + 2+2 + 2+3 + 3+5 = 19.

What is the strength of the longest bridge you can make? If you can make multiple bridges of the longest length, pick the strongest one.

Your puzzle answer was 1824.

*/

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Advent of Code 2017 - Day 24")
	fmt.Println("Strongest bridge:          ", StrongestBridge(0, nil, NewComponents(puzzle)))
	_, s := LongestBridge(0, nil, NewComponents(puzzle))
	fmt.Println("Strength of longest bridge:", s)
}

func StrongestBridge(next int, chain []*Component, comps Components) int {
	if chain == nil {
		chain = make([]*Component, 0)
	}

	// for all possible next components
	strongest := 0
	for _, c := range comps[next] {
		used := false
		for i := range chain {
			if c.ID == chain[i].ID {
				used = true
				break
			}
		}
		if used {
			continue
		}

		newChain := append(chain, c)
		newNext := c.P1
		if c.P1 == next {
			newNext = c.P2
		}
		s := StrongestBridge(newNext, newChain, comps)
		s += c.P1 + c.P2
		if s > strongest {
			strongest = s
		}
	}

	return strongest
}

func LongestBridge(next int, chain []*Component, comps Components) (int, int) {
	if chain == nil {
		chain = make([]*Component, 0)
	}

	// for all possible next components
	longest := 0
	strength := 0
	for _, c := range comps[next] {
		used := false
		for i := range chain {
			if c.ID == chain[i].ID {
				used = true
				break
			}
		}
		if used {
			continue
		}

		newChain := append(chain, c)
		newNext := c.P1
		if c.P1 == next {
			newNext = c.P2
		}
		l, s := LongestBridge(newNext, newChain, comps)
		l++
		s += c.P1 + c.P2
		if l > longest || (l == longest && s > strength) {
			longest = l
			strength = s
		}
	}

	return longest, strength
}

func NewComponents(in string) Components {
	components := make(Components)
	for id, line := range strings.Split(in, "\n") {
		ends := strings.Split(line, "/")
		if len(ends) != 2 {
			panic("Invalid input line: " + line)
		}
		port1, err := strconv.Atoi(ends[0])
		if err != nil {
			panic("Invalid port: " + line + " " + err.Error())
		}
		port2, err := strconv.Atoi(ends[1])
		if err != nil {
			panic("Invalid port: " + line + " " + err.Error())
		}
		b := &Component{
			ID: id,
			P1: port1,
			P2: port2,
		}
		components[port1] = append(components[port1], b)
		components[port2] = append(components[port2], b)
	}
	return components
}

type Component struct {
	ID     int
	P1, P2 int
}

type Components map[int][]*Component

const puzzle = `31/13
34/4
49/49
23/37
47/45
32/4
12/35
37/30
41/48
0/47
32/30
12/5
37/31
7/41
10/28
35/4
28/35
20/29
32/20
31/43
48/14
10/11
27/6
9/24
8/28
45/48
8/1
16/19
45/45
0/4
29/33
2/5
33/9
11/7
32/10
44/1
40/32
2/45
16/16
1/18
38/36
34/24
39/44
32/37
26/46
25/33
9/10
0/29
38/8
33/33
49/19
18/20
49/39
18/39
26/13
19/32`
