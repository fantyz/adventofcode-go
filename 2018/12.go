package main

import (
	"fmt"
	"strings"
)

/*


 */

func main() {
	fmt.Println("Day 12: Subterranean Sustainability")
}

type Pots struct {
	pots  []bool
	zero  int
	rules []*Rule
}

func NewPots(initialPots string, rules string) *Pots {
	p := Pots{
		pots: make([]bool, len(initialPots)),
	}
	for i := 0; i < len(initialPots); i++ {
		p.pots[i] = initialPots[i] == '#'
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

func (p *Pots) Tick() *Pots {
	new := make([]bool, len(p.pots)+4)
	for i := 0; i < len(p.pots); i++ {
		for _, r := range p.rules {
			if new[i-2] == r.cond[0] && new[i-1] == r.cond[1] && new[i] == r.cond[2] && new[i+1] == r.cond[3] && new[i+2] == r.cond[4] {
				new[i] = 
			}
		}
	}

	return p
}

func (p *Pots) String() string {
	str := ""
	for i := range p.pots {
		if p.pots[i] {
			str += "#"
		} else {
			str += "."
		}
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
