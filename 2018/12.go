package main

import (
	"fmt"
	"strings"
)

/*


 */

func main() {
	fmt.Println("Day 12: Subterranean Sustainability")
	p := NewPots(initialPuzzleState, rules)
	p.Ticks(20)
	fmt.Println(" > Sum after 20 generations:", p.Sum())
	p.Ticks(50000000000 - 20)
	fmt.Println(" > Sum after 50000000000 generations:", p.Sum())
}

type Pot struct {
	next, prev *Pot
	id         int
	val        bool
	prevVal bool
}

type Pots struct {
	pots  *Pot
	rules []*Rule
}

func NewPots(initialPots string, rules string) *Pots {
	pots := &Pots{
		pots: &Pot{
			id:  0,
			val: initialPots[0] == '#',
		},
	}
	tail := pots.pots

	for i := 1; i < len(initialPots); i++ {
		tail.next = &Pot{
			prev: tail,
			id:   i,
			val:  initialPots[i] == '#',
		}
		tail = tail.next
	}
	for _, l := range strings.Split(rules, "\n") {
		if len(l) <= 0 {
			continue
		}
		r := Rule{
			cond: [5]bool{},
		}
		for i := 0; i < 5; i++ {
			r.cond[i] = l[i] == '#'
		}
		r.res = l[9] == '#'
		p.rules = append(p.rules, &r)
	}

	return &p
}

func (p *Pots) Ticks(n int) *Pots {
	for i := 0; i < n; i++ {
		if i%1000000 == 0 && i > 0 {
			fmt.Print(".")
		}
		p.Tick()
	}
	fmt.Println()
	return p
}

func (p *Pots) Tick() *Pots {
	for 

	
	new := make([]bool, len(p.pots)+8)
	for i := 0; i < len(new); i++ {
		p1, p2, p3, p4, p5 := false, false, false, false, false
		if i-2-4 >= 0 && i-2-4 < len(p.pots) {
			p1 = p.pots[i-2-4]
		}
		if i-1-4 >= 0 && i-1-4 < len(p.pots) {
			p2 = p.pots[i-1-4]
		}
		if i-4 >= 0 && i-4 < len(p.pots) {
			p3 = p.pots[i-4]
		}
		if i+1-4 >= 0 && i+1-4 < len(p.pots) {
			p4 = p.pots[i+1-4]
		}
		if i+2-4 >= 0 && i+2-4 < len(p.pots) {
			p5 = p.pots[i+2-4]
		}

		for _, r := range p.rules {
			if p1 == r.cond[0] && p2 == r.cond[1] && p3 == r.cond[2] && p4 == r.cond[3] && p5 == r.cond[4] {
				new[i] = r.res
				break
			}
		}
	}

	first := len(new)
	last := 0
	for i := range new {
		if new[i] {
			if first > i {
				first = i
			}
			if last < i {
				last = i
			}
		}
	}

	p.pots = new[first : last+1]

	p.zero += 4 - first

	return p
}

func (p *Pots) Sum() int {
	sum := 0
	cur = p.pots
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
	cur = p.pots
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

type Rule struct {
	cond [5]bool
	res  bool
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
