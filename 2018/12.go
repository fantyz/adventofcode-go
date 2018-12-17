package main

import (
	"fmt"
	"strings"
)

/*

--- Day 12: Subterranean Sustainability ---
The year 518 is significantly more underground than your history books implied. Either that, or you've arrived in a vast cavern network under the North Pole.

After exploring a little, you discover a long tunnel that contains a row of small pots as far as you can see to your left and right. A few of them contain plants - someone is trying to grow things in these geothermally-heated caves.

The pots are numbered, with 0 in front of you. To the left, the pots are numbered -1, -2, -3, and so on; to the right, 1, 2, 3.... Your puzzle input contains a list of pots from 0 to the right and whether they do (#) or do not (.) currently contain a plant, the initial state. (No other pots currently contain plants.) For example, an initial state of #..##.... indicates that pots 0, 3, and 4 currently contain plants.

Your puzzle input also contains some notes you find on a nearby table: someone has been trying to figure out how these plants spread to nearby pots. Based on the notes, for each generation of plants, a given pot has or does not have a plant based on whether that pot (and the two pots on either side of it) had a plant in the last generation. These are written as LLCRR => N, where L are pots to the left, C is the current pot being considered, R are the pots to the right, and N is whether the current pot will have a plant in the next generation. For example:

A note like ..#.. => . means that a pot that contains a plant but with no plants within two pots of it will not have a plant in it during the next generation.
A note like ##.## => . means that an empty pot with two plants on each side of it will remain empty in the next generation.
A note like .##.# => # means that a pot has a plant in a given generation if, in the previous generation, there were plants in that pot, the one immediately to the left, and the one two pots to the right, but not in the ones immediately to the right and two to the left.
It's not clear what these plants are for, but you're sure it's important, so you'd like to make sure the current configuration of plants is sustainable by determining what will happen after 20 generations.

For example, given the following input:

initial state: #..#.#..##......###...###

...## => #
..#.. => #
.#... => #
.#.#. => #
.#.## => #
.##.. => #
.#### => #
#.#.# => #
#.### => #
##.#. => #
##.## => #
###.. => #
###.# => #
####. => #
For brevity, in this example, only the combinations which do produce a plant are listed. (Your input includes all possible combinations.) Then, the next 20 generations will look like this:

                 1         2         3
       0         0         0         0
 0: ...#..#.#..##......###...###...........
 1: ...#...#....#.....#..#..#..#...........
 2: ...##..##...##....#..#..#..##..........
 3: ..#.#...#..#.#....#..#..#...#..........
 4: ...#.#..#...#.#...#..#..##..##.........
 5: ....#...##...#.#..#..#...#...#.........
 6: ....##.#.#....#...#..##..##..##........
 7: ...#..###.#...##..#...#...#...#........
 8: ...#....##.#.#.#..##..##..##..##.......
 9: ...##..#..#####....#...#...#...#.......
10: ..#.#..#...#.##....##..##..##..##......
11: ...#...##...#.#...#.#...#...#...#......
12: ...##.#.#....#.#...#.#..##..##..##.....
13: ..#..###.#....#.#...#....#...#...#.....
14: ..#....##.#....#.#..##...##..##..##....
15: ..##..#..#.#....#....#..#.#...#...#....
16: .#.#..#...#.#...##...#...#.#..##..##...
17: ..#...##...#.#.#.#...##...#....#...#...
18: ..##.#.#....#####.#.#.#...##...##..##..
19: .#..###.#..#.#.#######.#.#.#..#.#...#..
20: .#....##....#####...#######....#.#..##.
The generation is shown along the left, where 0 is the initial state. The pot numbers are shown along the top, where 0 labels the center pot, negative-numbered pots extend to the left, and positive pots extend toward the right. Remember, the initial state begins at pot 0, which is not the leftmost pot used in this example.

After one generation, only seven plants remain. The one in pot 0 matched the rule looking for ..#.., the one in pot 4 matched the rule looking for .#.#., pot 9 matched .##.., and so on.

In this example, after 20 generations, the pots shown as # contain plants, the furthest left of which is pot -2, and the furthest right of which is pot 34. Adding up all the numbers of plant-containing pots after the 20th generation produces 325.

After 20 generations, what is the sum of the numbers of all pots which contain a plant?

Your puzzle answer was 1787.

--- Part Two ---
You realize that 20 generations aren't enough. After all, these plants will need to last another 1500 years to even reach your timeline, not to mention your future.

After fifty billion (50000000000) generations, what is the sum of the numbers of all pots which contain a plant?

*/

func main() {
	fmt.Println("Day 12: Subterranean Sustainability")
	p1 := NewPots(initialPuzzleState, rules).Ticks(20)
	fmt.Println(" > Sum after 20 generations:", p1.Sum())
	//p2 := NewPots(initialPuzzleState, rules).Ticks(50000000000)
	//p2 := NewPots(initialPuzzleState, rules).Ticks(3000000)
	//fmt.Println(" > Sum after 50000000000 generations:", p2.Sum())

	// Part 2 was solved using a spreadsheet. There is a clear pattern emerging after a very short amount of time. The result was extrapolated from this.
	// One could make a program as well that could detect this and do the math to make it run very fast.
}

type Pot struct {
	next, prev *Pot
	id         int
	val        bool
	prevVal    bool
}

type Pots struct {
	head, tail *Pot
	rules      map[uint8]bool
}

func NewPots(initialPots string, rules string) *Pots {
	pots := &Pots{
		head: &Pot{
			id:  0,
			val: initialPots[0] == '#',
		},
		rules: make(map[uint8]bool),
	}
	pots.tail = pots.head

	for i := 1; i < len(initialPots); i++ {
		pots.tail.next = &Pot{
			prev: pots.tail,
			id:   i,
			val:  initialPots[i] == '#',
		}
		pots.tail = pots.tail.next
	}

	for _, l := range strings.Split(rules, "\n") {
		if len(l) <= 0 || l[9] != '#' {
			continue
		}
		r := uint8(0)
		for i := 0; i < 5; i++ {
			r = r << 1
			if l[i] == '#' {
				r++
			}
		}
		pots.rules[r] = true
	}

	return pots
}

func (p *Pots) Ticks(n int) *Pots {
	for i := 0; i < n; i++ {
		if i%10 == 0 && i > 0 {
			fmt.Println(p.FullString())
		}
		if i%1000000 == 0 && i > 0 {
			fmt.Print(".")
		}
		p.Tick()
	}
	fmt.Println()
	return p
}

func (p *Pots) Tick() *Pots {
	cur := p.head
	for cur != nil {
		cur.prevVal = cur.val
		cur.val = false
		cur = cur.next
	}

	spareHeads := -4
	cur = p.head
	for cur != nil && !cur.prevVal {
		spareHeads++
		cur = cur.next
	}
	cur = p.tail
	spareTails := -4
	for cur != nil && !cur.prevVal {
		spareTails++
		cur = cur.prev
	}
	if spareHeads < 0 {
		for spareHeads < 0 {
			p.head = &Pot{
				next: p.head,
				id:   p.head.id - 1,
			}
			p.head.next.prev = p.head
			spareHeads++
		}
	} else {
		for spareHeads >= 0 {
			p.head = p.head.next
			spareHeads--
		}
	}
	if spareTails < 0 {
		for spareTails < 0 {
			p.tail = &Pot{
				prev: p.tail,
				id:   p.tail.id + 1,
			}
			p.tail.prev.next = p.tail
			spareTails++
		}
	} else {
		for spareTails >= 0 {
			p.tail = p.tail.prev
			spareTails--
		}
	}

	cur = p.head.next.next
	for cur.next.next != nil {
		r := uint8(0)
		tmp := cur.prev.prev
		for i := 0; i < 5; i++ {
			r = r << 1
			if tmp.prevVal {
				r++
			}
			tmp = tmp.next
		}

		if _, found := p.rules[r]; found {
			cur.val = true
		}
		cur = cur.next
	}

	return p
}

func (p *Pots) Sum() int {
	sum := 0
	cur := p.head
	for cur != nil {
		if cur.val {
			sum += cur.id
		}
		cur = cur.next
	}
	return sum
}

func (p *Pots) String() string {
	str := ""

	cur := p.head
	tail := p.tail

	for cur != nil && !cur.val {
		cur = cur.next
	}
	for tail != nil && !tail.val {
		tail = tail.prev
	}

	for {
		if cur.val {
			str += "#"
		} else {
			str += "."
		}
		cur = cur.next
		if cur == tail {
			str += "#"
			break
		}
	}
	return str
}

func (p *Pots) FullString() string {
	str := ""
	cur := p.head
	for cur != nil {
		if cur.val {
			str += "#"
		} else {
			str += "."
		}
		cur = cur.next
	}
	return str
}

const initialPuzzleState = `#...#..###.#.###.####.####.#..#.##..#..##..#.....#.#.#.##.#...###.#..##..#.##..###..#..##.#..##`

const rules = `...#. => #
#..## => #
..... => .
##.## => .
.##.. => #
.##.# => .
####. => #
.#.#. => .
..#.# => .
.#.## => .
.#..# => .
##... => #
#...# => #
##### => .
#.### => #
..### => #
###.. => .
#.#.# => #
##..# => #
..#.. => #
.#### => .
#.##. => .
....# => .
...## => .
#.... => .
#..#. => .
..##. => .
###.# => #
#.#.. => #
##.#. => #
.###. => .
.#... => .`

const puzzleInput = `initial state: #...#..###.#.###.####.####.#..#.##..#..##..#.....#.#.#.##.#...###.#..##..#.##..###..#..##.#..##...

...#. => #
#..## => #
..... => .
##.## => .
.##.. => #
.##.# => .
####. => #
.#.#. => .
..#.# => .
.#.## => .
.#..# => .
##... => #
#...# => #
##### => .
#.### => #
..### => #
###.. => .
#.#.# => #
##..# => #
..#.. => #
.#### => .
#.##. => .
....# => .
...## => .
#.... => .
#..#. => .
..##. => .
###.# => #
#.#.. => #
##.#. => #
.###. => .
.#... => .`
