package main

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

/*

--- Day 10: The Stars Align ---
It's no use; your navigation system simply isn't capable of providing walking directions in the arctic circle, and certainly not in 1018.

The Elves suggest an alternative. In times like these, North Pole rescue operations will arrange points of light in the sky to guide missing Elves back to base. Unfortunately, the message is easy to miss: the points move slowly enough that it takes hours to align them, but have so much momentum that they only stay aligned for a second. If you blink at the wrong time, it might be hours before another message appears.

You can see these points of light floating in the distance, and record their position in the sky and their velocity, the relative change in position per second (your puzzle input). The coordinates are all given from your perspective; given enough time, those positions and velocities will move the points into a cohesive message!

Rather than wait, you decide to fast-forward the process and calculate what the points will eventually spell.

For example, suppose you note the following points:

position=< 9,  1> velocity=< 0,  2>
position=< 7,  0> velocity=<-1,  0>
position=< 3, -2> velocity=<-1,  1>
position=< 6, 10> velocity=<-2, -1>
position=< 2, -4> velocity=< 2,  2>
position=<-6, 10> velocity=< 2, -2>
position=< 1,  8> velocity=< 1, -1>
position=< 1,  7> velocity=< 1,  0>
position=<-3, 11> velocity=< 1, -2>
position=< 7,  6> velocity=<-1, -1>
position=<-2,  3> velocity=< 1,  0>
position=<-4,  3> velocity=< 2,  0>
position=<10, -3> velocity=<-1,  1>
position=< 5, 11> velocity=< 1, -2>
position=< 4,  7> velocity=< 0, -1>
position=< 8, -2> velocity=< 0,  1>
position=<15,  0> velocity=<-2,  0>
position=< 1,  6> velocity=< 1,  0>
position=< 8,  9> velocity=< 0, -1>
position=< 3,  3> velocity=<-1,  1>
position=< 0,  5> velocity=< 0, -1>
position=<-2,  2> velocity=< 2,  0>
position=< 5, -2> velocity=< 1,  2>
position=< 1,  4> velocity=< 2,  1>
position=<-2,  7> velocity=< 2, -2>
position=< 3,  6> velocity=<-1, -1>
position=< 5,  0> velocity=< 1,  0>
position=<-6,  0> velocity=< 2,  0>
position=< 5,  9> velocity=< 1, -2>
position=<14,  7> velocity=<-2,  0>
position=<-3,  6> velocity=< 2, -1>
Each line represents one point. Positions are given as <X, Y> pairs: X represents how far left (negative) or right (positive) the point appears, while Y represents how far up (negative) or down (positive) the point appears.

At 0 seconds, each point has the position given. Each second, each point's velocity is added to its position. So, a point with velocity <1, -2> is moving to the right, but is moving upward twice as quickly. If this point's initial position were <3, 9>, after 3 seconds, its position would become <6, 3>.

Over time, the points listed above would move like this:

Initially:
........#.............
................#.....
.........#.#..#.......
......................
#..........#.#.......#
...............#......
....#.................
..#.#....#............
.......#..............
......#...............
...#...#.#...#........
....#..#..#.........#.
.......#..............
...........#..#.......
#...........#.........
...#.......#..........

After 1 second:
......................
......................
..........#....#......
........#.....#.......
..#.........#......#..
......................
......#...............
....##.........#......
......#.#.............
.....##.##..#.........
........#.#...........
........#...#.....#...
..#...........#.......
....#.....#.#.........
......................
......................

After 2 seconds:
......................
......................
......................
..............#.......
....#..#...####..#....
......................
........#....#........
......#.#.............
.......#...#..........
.......#..#..#.#......
....#....#.#..........
.....#...#...##.#.....
........#.............
......................
......................
......................

After 3 seconds:
......................
......................
......................
......................
......#...#..###......
......#...#...#.......
......#...#...#.......
......#####...#.......
......#...#...#.......
......#...#...#.......
......#...#...#.......
......#...#..###......
......................
......................
......................
......................

After 4 seconds:
......................
......................
......................
............#.........
........##...#.#......
......#.....#..#......
.....#..##.##.#.......
.......##.#....#......
...........#....#.....
..............#.......
....#......#...#......
.....#.....##.........
...............#......
...............#......
......................
......................
After 3 seconds, the message appeared briefly: HI. Of course, your message will be much longer and will take many more seconds to appear.

What message will eventually appear in the sky?

Your puzzle answer was GJNKBZEE.

--- Part Two ---
Good thing you didn't have to wait, because that would have taken a long time - much longer than the 3 seconds in the example above.

Impressed by your sub-hour communication capabilities, the Elves are curious: exactly how many seconds would they have needed to wait for that message to appear?

*/

func main() {
	fmt.Println("Day 10: The Stars Align")
	f := NewField(puzzleInput)
	fmt.Printf(" > Message (secs=%d)\n", f.FindMessage())
	fmt.Println(f.Plot())
}

type Field []*Point

type Point struct {
	X, Y       int
	velX, velY int
}

func NewField(in string) Field {
	var f []*Point

	re := regexp.MustCompile(`^position=<[ ]*(-?\d+),[ ]*(-?\d+)> velocity=<[ ]*(-?\d+),[ ]*(-?\d+)>$`)
	for _, l := range strings.Split(in, "\n") {
		m := re.FindStringSubmatch(l)
		if len(m) < 5 {
			panic("line did not match: " + l)
		}
		startX, err := strconv.Atoi(m[1])
		if err != nil {
			panic(err)
		}
		startY, err := strconv.Atoi(m[2])
		if err != nil {
			panic(err)
		}
		velX, err := strconv.Atoi(m[3])
		if err != nil {
			panic(err)
		}
		velY, err := strconv.Atoi(m[4])
		if err != nil {
			panic(err)
		}
		f = append(f, &Point{startX, startY, velX, velY})
	}

	return f
}

func (f Field) Plot() string {
	sort.Slice(f, func(i, j int) bool {
		return f[i].Y < f[j].Y || (f[i].Y == f[j].Y && f[i].X < f[j].X)
	})

	xmin, xmax := 999999, -999999
	ymin, ymax := 999999, -999999

	for _, p := range f {
		if xmin > p.X {
			xmin = p.X
		}
		if xmax < p.X {
			xmax = p.X
		}
		if ymin > p.Y {
			ymin = p.Y
		}
		if ymax < p.Y {
			ymax = p.Y
		}
	}

	idx := 0

	str := ""
	for y := ymin; y <= ymax; y++ {
		for x := xmin; x <= xmax; x++ {
			found := false
			for {
				if idx < len(f) && f[idx].X == x && f[idx].Y == y {
					idx++
					found = true
					continue
				}
				break
			}
			if found {
				str += "#"
			} else {
				str += "."
			}
		}

		for {
			if idx >= len(f) || f[idx].Y != y {
				break
			}
			idx++
		}

		str += "\n"
	}
	return str
}

func (f Field) Neighbours() float32 {
	neighbours := 0
	for _, p := range f {
		if f.HasNeighbours(p.X, p.Y) {
			neighbours++
		}
	}
	return float32(neighbours) / float32(len(f))
}

func (f Field) HasNeighbours(x, y int) bool {
	for _, p := range f {
		if (p.X == x+1 || p.X == x || p.X == x-1) && (p.Y == y+1 || p.Y == y || p.Y == y-1) && !(p.X == x && p.Y == y) {
			return true
		}
	}
	return false
}

func (f Field) Tick() {
	for _, p := range f {
		p.X += p.velX
		p.Y += p.velY
	}
}

func (f Field) FindMessage() int {
	sec := 0
	for {
		if f.Neighbours() > .95 {
			break
		}
		f.Tick()
		sec++
	}
	return sec
}

const puzzleInput = `position=< 53777,  21594> velocity=<-5, -2>
position=< 53761,  53776> velocity=<-5, -5>
position=<-32066,  53779> velocity=< 3, -5>
position=<-21287,  43043> velocity=< 2, -4>
position=< 10848, -42773> velocity=<-1,  4>
position=<-10596,  53770> velocity=< 1, -5>
position=<-42798,  53772> velocity=< 4, -5>
position=<-21308, -32037> velocity=< 2,  3>
position=<-21332,  10863> velocity=< 2, -1>
position=<-10596, -21313> velocity=< 1,  2>
position=<-42750, -53498> velocity=< 4,  5>
position=<-10569, -21315> velocity=< 1,  2>
position=<-21334,  53779> velocity=< 2, -5>
position=< 43055, -10586> velocity=<-4,  1>
position=< 21588, -21313> velocity=<-2,  2>
position=< 53795,  32316> velocity=<-5, -3>
position=< 43061, -21319> velocity=<-4,  2>
position=<-10590, -21317> velocity=< 1,  2>
position=<-21344, -32042> velocity=< 2,  3>
position=<-53491, -42769> velocity=< 5,  4>
position=< 53801,  21595> velocity=<-5, -2>
position=< 43055, -42771> velocity=<-4,  4>
position=<-32027, -21319> velocity=< 3,  2>
position=< 43068,  21598> velocity=<-4, -2>
position=< 32307, -53499> velocity=<-3,  5>
position=< 53787, -53491> velocity=<-5,  5>
position=<-21288,  53778> velocity=< 2, -5>
position=<-42787,  10862> velocity=< 4, -1>
position=<-53522,  32325> velocity=< 5, -3>
position=<-53513,  21595> velocity=< 5, -2>
position=<-42741, -32042> velocity=< 4,  3>
position=< 43022, -53495> velocity=<-4,  5>
position=< 43076,  21593> velocity=<-4, -2>
position=< 53757, -10589> velocity=<-5,  1>
position=<-32035,  43047> velocity=< 3, -4>
position=< 32328, -42768> velocity=<-3,  4>
position=< 21608, -53500> velocity=<-2,  5>
position=< 53790, -21318> velocity=<-5,  2>
position=<-10561, -32038> velocity=< 1,  3>
position=<-32068, -42768> velocity=< 3,  4>
position=< 10861, -10587> velocity=<-1,  1>
position=< 43068,  53779> velocity=<-4, -5>
position=< 21568, -10592> velocity=<-2,  1>
position=<-42766, -10584> velocity=< 4,  1>
position=<-53469, -10585> velocity=< 5,  1>
position=< 21564,  32322> velocity=<-2, -3>
position=<-42762, -32037> velocity=< 4,  3>
position=< 21576, -10587> velocity=<-2,  1>
position=<-10567,  53774> velocity=< 1, -5>
position=<-21283,  21589> velocity=< 2, -2>
position=<-32022, -32042> velocity=< 3,  3>
position=<-21309,  10871> velocity=< 2, -1>
position=< 53778,  53774> velocity=<-5, -5>
position=<-53493, -21317> velocity=< 5,  2>
position=< 43023, -42766> velocity=<-4,  4>
position=< 43030, -42768> velocity=<-4,  4>
position=< 10853,  21597> velocity=<-1, -2>
position=< 21584,  53778> velocity=<-2, -5>
position=<-53492,  53770> velocity=< 5, -5>
position=<-21309, -10592> velocity=< 2,  1>
position=<-10573, -53497> velocity=< 1,  5>
position=<-53513, -10592> velocity=< 5,  1>
position=<-42750, -42768> velocity=< 4,  4>
position=< 21589,  32321> velocity=<-2, -3>
position=< 43042,  21594> velocity=<-4, -2>
position=< 32308, -10590> velocity=<-3,  1>
position=<-32055, -21315> velocity=< 3,  2>
position=<-21318, -42767> velocity=< 2,  4>
position=<-53501,  10864> velocity=< 5, -1>
position=<-53474, -53496> velocity=< 5,  5>
position=<-53513, -21318> velocity=< 5,  2>
position=<-10617, -21318> velocity=< 1,  2>
position=< 10869, -42770> velocity=<-1,  4>
position=< 43034, -21318> velocity=<-4,  2>
position=< 53785,  21597> velocity=<-5, -2>
position=<-32070, -53491> velocity=< 3,  5>
position=<-42771, -32039> velocity=< 4,  3>
position=<-42786, -32045> velocity=< 4,  3>
position=<-10597,  53777> velocity=< 1, -5>
position=<-32023, -32042> velocity=< 3,  3>
position=< 53787, -10592> velocity=<-5,  1>
position=<-21318,  53776> velocity=< 2, -5>
position=<-10585, -53499> velocity=< 1,  5>
position=< 32315,  32323> velocity=<-3, -3>
position=< 53794, -53496> velocity=<-5,  5>
position=< 10864, -42766> velocity=<-1,  4>
position=<-32055, -42770> velocity=< 3,  4>
position=<-53467,  53770> velocity=< 5, -5>
position=< 32295,  53770> velocity=<-3, -5>
position=< 32291, -32038> velocity=<-3,  3>
position=<-32053,  21593> velocity=< 3, -2>
position=< 10888,  43047> velocity=<-1, -4>
position=< 53777,  53776> velocity=<-5, -5>
position=< 43070,  43052> velocity=<-4, -4>
position=<-53509,  53775> velocity=< 5, -5>
position=< 43047, -21310> velocity=<-4,  2>
position=< 10896,  53774> velocity=<-1, -5>
position=< 10861,  32320> velocity=<-1, -3>
position=< 21601,  43050> velocity=<-2, -4>
position=< 53749, -21314> velocity=<-5,  2>
position=<-10566,  43043> velocity=< 1, -4>
position=< 53747, -32037> velocity=<-5,  3>
position=< 21593, -42773> velocity=<-2,  4>
position=<-42782,  21598> velocity=< 4, -2>
position=<-21299, -10592> velocity=< 2,  1>
position=<-42766, -10592> velocity=< 4,  1>
position=< 43022, -21311> velocity=<-4,  2>
position=< 53801, -21314> velocity=<-5,  2>
position=<-42742,  53779> velocity=< 4, -5>
position=< 43042,  10862> velocity=<-4, -1>
position=<-32066,  32322> velocity=< 3, -3>
position=< 53746,  10871> velocity=<-5, -1>
position=< 43038,  43050> velocity=<-4, -4>
position=< 32349, -10588> velocity=<-3,  1>
position=< 10858,  53778> velocity=<-1, -5>
position=< 32304, -42773> velocity=<-3,  4>
position=< 43042, -53499> velocity=<-4,  5>
position=< 43035,  10864> velocity=<-4, -1>
position=<-32021,  21593> velocity=< 3, -2>
position=<-53501, -53494> velocity=< 5,  5>
position=< 43074,  32325> velocity=<-4, -3>
position=< 21620, -53499> velocity=<-2,  5>
position=<-21320,  32322> velocity=< 2, -3>
position=< 43050,  53779> velocity=<-4, -5>
position=<-10601, -32045> velocity=< 1,  3>
position=< 21608, -21310> velocity=<-2,  2>
position=<-42794, -10592> velocity=< 4,  1>
position=< 53750,  10863> velocity=<-5, -1>
position=<-21341, -53500> velocity=< 2,  5>
position=< 32348,  10871> velocity=<-3, -1>
position=< 32307, -21316> velocity=<-3,  2>
position=<-32019, -32042> velocity=< 3,  3>
position=< 53774, -32046> velocity=<-5,  3>
position=<-21286, -42764> velocity=< 2,  4>
position=<-10615,  21598> velocity=< 1, -2>
position=< 21604,  32316> velocity=<-2, -3>
position=< 32312, -53491> velocity=<-3,  5>
position=<-21339, -42764> velocity=< 2,  4>
position=<-42766,  43047> velocity=< 4, -4>
position=< 10888, -32037> velocity=<-1,  3>
position=< 53754,  21598> velocity=<-5, -2>
position=< 21576,  32320> velocity=<-2, -3>
position=< 21580, -21317> velocity=<-2,  2>
position=< 53750, -32041> velocity=<-5,  3>
position=<-21323, -53495> velocity=< 2,  5>
position=< 43066,  21592> velocity=<-4, -2>
position=< 10840, -32046> velocity=<-1,  3>
position=< 32343,  53774> velocity=<-3, -5>
position=<-10564,  53779> velocity=< 1, -5>
position=< 21590, -32043> velocity=<-2,  3>
position=< 21612,  43046> velocity=<-2, -4>
position=< 21576, -21312> velocity=<-2,  2>
position=< 53777, -32044> velocity=<-5,  3>
position=<-53493,  10863> velocity=< 5, -1>
position=< 32332,  10871> velocity=<-3, -1>
position=< 10880,  53774> velocity=<-1, -5>
position=<-53482,  32320> velocity=< 5, -3>
position=< 10889,  10862> velocity=<-1, -1>
position=<-10558, -32037> velocity=< 1,  3>
position=<-21304,  10870> velocity=< 2, -1>
position=<-32034,  21590> velocity=< 3, -2>
position=<-32039, -42770> velocity=< 3,  4>
position=< 53770,  53775> velocity=<-5, -5>
position=< 32332,  21595> velocity=<-3, -2>
position=< 32349,  53770> velocity=<-3, -5>
position=< 10886, -53496> velocity=<-1,  5>
position=<-32070,  21589> velocity=< 3, -2>
position=< 21606,  10871> velocity=<-2, -1>
position=<-42750, -53492> velocity=< 4,  5>
position=< 53772,  10869> velocity=<-5, -1>
position=<-42753,  43045> velocity=< 4, -4>
position=<-21316, -53499> velocity=< 2,  5>
position=<-32066, -53492> velocity=< 3,  5>
position=<-53509,  32319> velocity=< 5, -3>
position=<-42774, -42764> velocity=< 4,  4>
position=< 32318, -42771> velocity=<-3,  4>
position=<-21302,  53775> velocity=< 2, -5>
position=< 32349, -42764> velocity=<-3,  4>
position=< 10893, -32039> velocity=<-1,  3>
position=< 43022,  21597> velocity=<-4, -2>
position=< 43062, -10583> velocity=<-4,  1>
position=<-32066,  10870> velocity=< 3, -1>
position=< 10848, -42764> velocity=<-1,  4>
position=< 32325,  10866> velocity=<-3, -1>
position=<-53514,  53779> velocity=< 5, -5>
position=<-21317,  21591> velocity=< 2, -2>
position=<-53476, -32037> velocity=< 5,  3>
position=<-10617,  43046> velocity=< 1, -4>
position=<-42795,  53779> velocity=< 4, -5>
position=< 53758, -10592> velocity=<-5,  1>
position=<-10615, -53500> velocity=< 1,  5>
position=< 21620,  10867> velocity=<-2, -1>
position=< 10874, -42771> velocity=<-1,  4>
position=< 32315,  53778> velocity=<-3, -5>
position=<-53490, -32042> velocity=< 5,  3>
position=< 21596,  10869> velocity=<-2, -1>
position=<-32027, -10589> velocity=< 3,  1>
position=< 53782,  53778> velocity=<-5, -5>
position=<-53469,  10862> velocity=< 5, -1>
position=<-32068,  43052> velocity=< 3, -4>
position=< 32339, -53500> velocity=<-3,  5>
position=< 21624, -42764> velocity=<-2,  4>
position=< 43039,  53771> velocity=<-4, -5>
position=< 32307, -32046> velocity=<-3,  3>
position=<-53469,  10864> velocity=< 5, -1>
position=< 32323, -53491> velocity=<-3,  5>
position=< 32324, -10592> velocity=<-3,  1>
position=< 21584,  21596> velocity=<-2, -2>
position=<-42741,  32320> velocity=< 4, -3>
position=<-32066,  32317> velocity=< 3, -3>
position=< 10849,  32323> velocity=<-1, -3>
position=< 53778, -21310> velocity=<-5,  2>
position=<-10569,  10868> velocity=< 1, -1>
position=< 21607, -42764> velocity=<-2,  4>
position=< 32325, -53500> velocity=<-3,  5>
position=<-42742, -42770> velocity=< 4,  4>
position=< 21614,  21598> velocity=<-2, -2>
position=< 21596,  53771> velocity=<-2, -5>
position=< 10849, -10592> velocity=<-1,  1>
position=<-32039,  53775> velocity=< 3, -5>
position=< 43047, -21319> velocity=<-4,  2>
position=<-10583,  53779> velocity=< 1, -5>
position=<-21332,  43046> velocity=< 2, -4>
position=< 21566,  43052> velocity=<-2, -4>
position=< 53785, -21310> velocity=<-5,  2>
position=< 32303, -21319> velocity=<-3,  2>
position=< 43070, -21310> velocity=<-4,  2>
position=< 43062,  21589> velocity=<-4, -2>
position=<-53469, -42773> velocity=< 5,  4>
position=< 43070,  43052> velocity=<-4, -4>
position=<-42794,  43048> velocity=< 4, -4>
position=<-10582,  32316> velocity=< 1, -3>
position=< 21585,  32320> velocity=<-2, -3>
position=< 21585, -10591> velocity=<-2,  1>
position=< 32312,  32322> velocity=<-3, -3>
position=< 43036, -53497> velocity=<-4,  5>
position=<-42770,  43044> velocity=< 4, -4>
position=<-10569, -42772> velocity=< 1,  4>
position=< 21621, -53500> velocity=<-2,  5>
position=< 53777,  32322> velocity=<-5, -3>
position=<-10591,  43049> velocity=< 1, -4>
position=< 43037, -42768> velocity=<-4,  4>
position=<-53505, -53492> velocity=< 5,  5>
position=< 32296, -53495> velocity=<-3,  5>
position=< 21585, -53493> velocity=<-2,  5>
position=< 21600, -42769> velocity=<-2,  4>
position=< 53797, -53496> velocity=<-5,  5>
position=< 10885,  10865> velocity=<-1, -1>
position=<-42766, -21314> velocity=< 4,  2>
position=< 32304,  53770> velocity=<-3, -5>
position=< 21564, -53497> velocity=<-2,  5>
position=< 21621,  32325> velocity=<-2, -3>
position=< 10838,  53770> velocity=<-1, -5>
position=< 21598,  43052> velocity=<-2, -4>
position=< 53785, -32037> velocity=<-5,  3>
position=< 53806,  10871> velocity=<-5, -1>
position=<-21285, -21315> velocity=< 2,  2>
position=<-42761,  10870> velocity=< 4, -1>
position=<-42748, -42773> velocity=< 4,  4>
position=< 21600, -32046> velocity=<-2,  3>
position=<-32063,  21596> velocity=< 3, -2>
position=<-42782, -21317> velocity=< 4,  2>
position=< 32323,  53779> velocity=<-3, -5>
position=<-21328,  10866> velocity=< 2, -1>
position=< 43076, -53491> velocity=<-4,  5>
position=<-53477, -32040> velocity=< 5,  3>
position=<-32055, -21312> velocity=< 3,  2>
position=< 53805,  32316> velocity=<-5, -3>
position=<-32038,  21593> velocity=< 3, -2>
position=<-42761, -42770> velocity=< 4,  4>
position=<-21303,  53770> velocity=< 2, -5>
position=< 53766, -42771> velocity=<-5,  4>
position=<-32066, -21318> velocity=< 3,  2>
position=<-21320, -53497> velocity=< 2,  5>
position=< 43039, -53497> velocity=<-4,  5>
position=<-10616, -42764> velocity=< 1,  4>
position=<-53477, -10585> velocity=< 5,  1>
position=<-32012,  32316> velocity=< 3, -3>
position=<-32047, -53500> velocity=< 3,  5>
position=<-32023,  32321> velocity=< 3, -3>
position=<-32059, -21311> velocity=< 3,  2>
position=< 21620,  53773> velocity=<-2, -5>
position=<-10589,  53778> velocity=< 1, -5>
position=< 10837,  32323> velocity=<-1, -3>
position=<-53505, -53492> velocity=< 5,  5>
position=<-53525,  10864> velocity=< 5, -1>
position=<-32047,  32323> velocity=< 3, -3>
position=< 21575, -53491> velocity=<-2,  5>
position=<-42738,  53774> velocity=< 4, -5>
position=<-32015, -21315> velocity=< 3,  2>
position=< 32307, -42771> velocity=<-3,  4>
position=< 43042,  10862> velocity=<-4, -1>
position=<-10580, -42767> velocity=< 1,  4>
position=< 10877,  43050> velocity=<-1, -4>
position=<-32052, -10586> velocity=< 3,  1>
position=< 10856,  21594> velocity=<-1, -2>
position=<-10572,  10864> velocity=< 1, -1>
position=< 10849,  21591> velocity=<-1, -2>
position=<-32045, -21316> velocity=< 3,  2>
position=<-53477,  43052> velocity=< 5, -4>
position=< 43055, -53494> velocity=<-4,  5>
position=< 10869, -53496> velocity=<-1,  5>
position=< 10865,  43051> velocity=<-1, -4>
position=< 10886, -10583> velocity=<-1,  1>
position=< 21585, -10588> velocity=<-2,  1>
position=<-42758, -42766> velocity=< 4,  4>
position=< 21607, -53491> velocity=<-2,  5>
position=<-32012,  53779> velocity=< 3, -5>
position=< 53782, -42768> velocity=<-5,  4>
position=< 10837,  43044> velocity=<-1, -4>
position=<-32066, -32038> velocity=< 3,  3>
position=<-53483,  10871> velocity=< 5, -1>
position=<-21299,  53779> velocity=< 2, -5>
position=< 53794,  53770> velocity=<-5, -5>
position=< 32316, -21315> velocity=<-3,  2>
position=< 32303,  10867> velocity=<-3, -1>
position=< 43068,  53774> velocity=<-4, -5>
position=< 32315, -53493> velocity=<-3,  5>
position=< 53782, -53497> velocity=<-5,  5>
position=< 32307,  32323> velocity=<-3, -3>
position=<-10617, -10590> velocity=< 1,  1>
position=<-21335,  10871> velocity=< 2, -1>
position=<-21336, -21311> velocity=< 2,  2>
position=< 43052,  43043> velocity=<-4, -4>
position=< 32327, -32042> velocity=<-3,  3>
position=< 43062, -53497> velocity=<-4,  5>
position=< 32294, -10592> velocity=<-3,  1>
position=< 10853,  53779> velocity=<-1, -5>
position=< 32339,  10866> velocity=<-3, -1>
position=<-53477, -21319> velocity=< 5,  2>
position=<-42765,  53774> velocity=< 4, -5>
position=<-53485,  32324> velocity=< 5, -3>
position=<-53469,  21589> velocity=< 5, -2>
position=<-32071, -32043> velocity=< 3,  3>
position=< 10881, -42764> velocity=<-1,  4>
position=<-21344,  10866> velocity=< 2, -1>
position=<-10564,  32316> velocity=< 1, -3>
position=< 32307,  10866> velocity=<-3, -1>
position=< 43039, -21318> velocity=<-4,  2>
position=< 32303, -32038> velocity=<-3,  3>
position=<-21320,  21594> velocity=< 2, -2>
position=<-10574, -42764> velocity=< 1,  4>
position=<-21332,  43045> velocity=< 2, -4>
position=<-21344, -10587> velocity=< 2,  1>
position=< 43030, -42767> velocity=<-4,  4>
position=<-53484, -21319> velocity=< 5,  2>
position=< 43058,  53777> velocity=<-4, -5>
position=<-21341,  32321> velocity=< 2, -3>
position=<-10583, -21319> velocity=< 1,  2>
position=<-32054, -10591> velocity=< 3,  1>
position=<-53491, -53496> velocity=< 5,  5>
position=<-42742,  43050> velocity=< 4, -4>
position=<-32026, -10583> velocity=< 3,  1>
position=< 10853,  21594> velocity=<-1, -2>
position=<-53483,  32321> velocity=< 5, -3>
position=<-21284, -21310> velocity=< 2,  2>
position=< 32323,  43051> velocity=<-3, -4>
position=< 32352,  53779> velocity=<-3, -5>
position=< 21612,  21594> velocity=<-2, -2>
position=< 32316,  43048> velocity=<-3, -4>`
