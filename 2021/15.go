package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

func init() { days["15"] = Day15 }

/*
--- Day 15: Chiton ---
You've almost reached the exit of the cave, but the walls are getting closer together. Your submarine can barely still fit, though; the main problem is that the walls of the cave are covered in chitons, and it would be best not to bump any of them.

The cavern is large, but has a very low ceiling, restricting your motion to two dimensions. The shape of the cavern resembles a square; a quick scan of chiton density produces a map of risk level throughout the cave (your puzzle input). For example:

1163751742
1381373672
2136511328
3694931569
7463417111
1319128137
1359912421
3125421639
1293138521
2311944581
You start in the top left position, your destination is the bottom right position, and you cannot move diagonally. The number at each position is its risk level; to determine the total risk of an entire path, add up the risk levels of each position you enter (that is, don't count the risk level of your starting position unless you enter it; leaving it adds no risk to your total).

Your goal is to find a path with the lowest total risk. In this example, a path with the lowest total risk is highlighted here:

1163751742
1381373672
2136511328
3694931569
7463417111
1319128137
1359912421
3125421639
1293138521
2311944581
The total risk of this path is 40 (the starting position is never entered, so its risk is not counted).

What is the lowest total risk of any path from the top left to the bottom right?

Your puzzle answer was 685.

--- Part Two ---
Now that you know how to find low-risk paths in the cave, you can try to find your way out.

The entire cave is actually five times larger in both dimensions than you thought; the area you originally scanned is just one tile in a 5x5 tile area that forms the full map. Your original map tile repeats to the right and downward; each time the tile repeats to the right or downward, all of its risk levels are 1 higher than the tile immediately up or left of it. However, risk levels above 9 wrap back around to 1. So, if your original map had some position with a risk level of 8, then that same position on each of the 25 total tiles would be as follows:

8 9 1 2 3
9 1 2 3 4
1 2 3 4 5
2 3 4 5 6
3 4 5 6 7
Each single digit above corresponds to the example position with a value of 8 on the top-left tile. Because the full map is actually five times larger in both dimensions, that position appears a total of 25 times, once in each duplicated tile, with the values shown above.

Here is the full five-times-as-large version of the first example above, with the original map in the top left corner highlighted:

11637517422274862853338597396444961841755517295286
13813736722492484783351359589446246169155735727126
21365113283247622439435873354154698446526571955763
36949315694715142671582625378269373648937148475914
74634171118574528222968563933317967414442817852555
13191281372421239248353234135946434524615754563572
13599124212461123532357223464346833457545794456865
31254216394236532741534764385264587549637569865174
12931385212314249632342535174345364628545647573965
23119445813422155692453326671356443778246755488935
22748628533385973964449618417555172952866628316397
24924847833513595894462461691557357271266846838237
32476224394358733541546984465265719557637682166874
47151426715826253782693736489371484759148259586125
85745282229685639333179674144428178525553928963666
24212392483532341359464345246157545635726865674683
24611235323572234643468334575457944568656815567976
42365327415347643852645875496375698651748671976285
23142496323425351743453646285456475739656758684176
34221556924533266713564437782467554889357866599146
33859739644496184175551729528666283163977739427418
35135958944624616915573572712668468382377957949348
43587335415469844652657195576376821668748793277985
58262537826937364893714847591482595861259361697236
96856393331796741444281785255539289636664139174777
35323413594643452461575456357268656746837976785794
35722346434683345754579445686568155679767926678187
53476438526458754963756986517486719762859782187396
34253517434536462854564757396567586841767869795287
45332667135644377824675548893578665991468977611257
44961841755517295286662831639777394274188841538529
46246169155735727126684683823779579493488168151459
54698446526571955763768216687487932779859814388196
69373648937148475914825958612593616972361472718347
17967414442817852555392896366641391747775241285888
46434524615754563572686567468379767857948187896815
46833457545794456865681556797679266781878137789298
64587549637569865174867197628597821873961893298417
45364628545647573965675868417678697952878971816398
56443778246755488935786659914689776112579188722368
55172952866628316397773942741888415385299952649631
57357271266846838237795794934881681514599279262561
65719557637682166874879327798598143881961925499217
71484759148259586125936169723614727183472583829458
28178525553928963666413917477752412858886352396999
57545635726865674683797678579481878968159298917926
57944568656815567976792667818781377892989248891319
75698651748671976285978218739618932984172914319528
56475739656758684176786979528789718163989182927419
67554889357866599146897761125791887223681299833479
Equipped with the full map, you can now find a path from the top left corner to the bottom right corner with the lowest total risk:

11637517422274862853338597396444961841755517295286
13813736722492484783351359589446246169155735727126
21365113283247622439435873354154698446526571955763
36949315694715142671582625378269373648937148475914
74634171118574528222968563933317967414442817852555
13191281372421239248353234135946434524615754563572
13599124212461123532357223464346833457545794456865
31254216394236532741534764385264587549637569865174
12931385212314249632342535174345364628545647573965
23119445813422155692453326671356443778246755488935
22748628533385973964449618417555172952866628316397
24924847833513595894462461691557357271266846838237
32476224394358733541546984465265719557637682166874
47151426715826253782693736489371484759148259586125
85745282229685639333179674144428178525553928963666
24212392483532341359464345246157545635726865674683
24611235323572234643468334575457944568656815567976
42365327415347643852645875496375698651748671976285
23142496323425351743453646285456475739656758684176
34221556924533266713564437782467554889357866599146
33859739644496184175551729528666283163977739427418
35135958944624616915573572712668468382377957949348
43587335415469844652657195576376821668748793277985
58262537826937364893714847591482595861259361697236
96856393331796741444281785255539289636664139174777
35323413594643452461575456357268656746837976785794
35722346434683345754579445686568155679767926678187
53476438526458754963756986517486719762859782187396
34253517434536462854564757396567586841767869795287
45332667135644377824675548893578665991468977611257
44961841755517295286662831639777394274188841538529
46246169155735727126684683823779579493488168151459
54698446526571955763768216687487932779859814388196
69373648937148475914825958612593616972361472718347
17967414442817852555392896366641391747775241285888
46434524615754563572686567468379767857948187896815
46833457545794456865681556797679266781878137789298
64587549637569865174867197628597821873961893298417
45364628545647573965675868417678697952878971816398
56443778246755488935786659914689776112579188722368
55172952866628316397773942741888415385299952649631
57357271266846838237795794934881681514599279262561
65719557637682166874879327798598143881961925499217
71484759148259586125936169723614727183472583829458
28178525553928963666413917477752412858886352396999
57545635726865674683797678579481878968159298917926
57944568656815567976792667818781377892989248891319
75698651748671976285978218739618932984172914319528
56475739656758684176786979528789718163989182927419
67554889357866599146897761125791887223681299833479
The total risk of this path is 315 (the starting position is still never entered, so its risk is not counted).

Using the full map, what is the lowest total risk of any path from the top left to the bottom right?

Your puzzle answer was 2995.
*/

func Day15() {
	fmt.Println("--- Day 15: Chiton ---")
	cave, err := NewCave(day15Input)
	if err != nil {
		fmt.Println(errors.Wrap(err, "Failed to load cave"))
		return
	}
	fmt.Println("Lowest risk route through the cave has a total risk of:", LowestRiskRoute(cave, cave.TopLeft(), cave.BottomRight()))
	cave.Enlarge(5)
	fmt.Println("Lowest risk route through the enlarged cave has a total risk of:", LowestRiskRoute(cave, cave.TopLeft(), cave.BottomRight()))
}

// Cave represents the cave, each coordinate representing the risk of the given position.
// The risk of the position X, Y is found in cave[y][x].
type Cave [][]int

// NewCave takes the puzzle input and returns the cave it represents.
// NewCave returns an error if the input contains anything unexpected.
func NewCave(in string) (Cave, error) {
	c := make(Cave, 0)
	width := -1
	for _, line := range strings.Split(in, "\n") {
		if width < 0 {
			width = len(line)
		}
		if len(line) != width {
			return nil, errors.Errorf("all rows must have the same width (width=%d, expected=%d)", len(line), width)
		}
		row := make([]int, 0, width)
		for _, v := range line {
			risk, err := strconv.Atoi(string(v))
			if err != nil {
				return nil, errors.Wrapf(err, "unexpected character (c=%s)", string(v))
			}
			row = append(row, risk)
		}
		c = append(c, row)
	}

	return c, nil
}

// Enlarge enlarges the cave with the factor specified. It is enlarged by repeating the existing cave
// and adding +1 to the risks for each repetion.
// See puzzle input for more information.
// If the factor specified is 1 or less Enlarge does nothing.
func (c *Cave) Enlarge(factor int) {
	if factor <= 1 {
		return
	}

	incWithWrap := func(n int) int {
		n++
		if n >= 10 {
			n = 1
		}
		return n
	}

	startWidth, startHeight := c.Width(), c.Height()

	for y := 0; y < startHeight*factor; y++ {
		row := make([]int, startWidth*factor)
		if y < startHeight {
			// initialize first values of row using the original cave values
			for x := 0; x < startWidth; x++ {
				row[x] = (*c)[y][x]
			}
		} else {
			// initialize first values of row using the previous repetion + 1
			for x := 0; x < startWidth; x++ {
				row[x] = incWithWrap((*c)[y-startWidth][x])
			}
		}

		// finish the row
		for x := startWidth; x < startWidth*factor; x++ {
			row[x] = incWithWrap(row[x-startWidth])
		}

		// update cave with the new row
		if y < len(*c) {
			(*c)[y] = row
		} else {
			*c = append(*c, row)
		}
	}
}

// Width returns the width of the cave.
func (c Cave) Width() int {
	if len(c) <= 0 {
		return 0
	}
	return len(c[0])
}

// Height returns the height of the cave.
func (c Cave) Height() int {
	return len(c)
}

// TopLeft returns the cave location of the top left position in the cave.
func (c Cave) TopLeft() CaveLoc {
	return CaveLoc{0, 0}
}

// BottomRight returns the cave location of the bottom right position of the cave.
func (c Cave) BottomRight() CaveLoc {
	return CaveLoc{c.Width() - 1, c.Height() - 1}
}

// Print outputs the cave to stdout.
func (c Cave) Print() {
	for y := 0; y < len(c); y++ {
		for x := 0; x < len(c[y]); x++ {
			fmt.Print(c[y][x])
		}
		fmt.Println()
	}
}

// CaveLoc represents a location within the cave.
type CaveLoc struct {
	X, Y int
}

// ShortestRoute finds the lowest risk route through the cave and returns the total sum of risk.
// ShortestRoute will return -1 if either start or end locations are outside the bounds of the cave.
func LowestRiskRoute(cave Cave, start, end CaveLoc) int {
	// make sure both start and end is inside the cave
	if start.X < 0 || start.X >= cave.Width() || start.Y < 0 || start.Y >= cave.Height() {
		return -1
	}
	if end.X < 0 || end.X >= cave.Width() || end.Y < 0 || end.Y >= cave.Height() {
		return -1
	}

	// lets use a-star to find the shortest route
	// See: https://en.wikipedia.org/wiki/A*_search_algorithm

	startLoc := routeLoc{Pos: start, IsOpen: true, BestDistance: 0, EstimatedTotal: estimateDist(start, end)}
	locByCaveLoc := map[CaveLoc]*routeLoc{start: &startLoc}
	openSet := []*routeLoc{&startLoc}

	var cur *routeLoc
	for len(openSet) > 0 {
		cur, openSet = openSet[0], openSet[1:]
		cur.IsOpen = false

		if cur.Pos.X == end.X && cur.Pos.Y == end.Y {
			// lowest risk route found, total risk corresponds to the distance
			return cur.BestDistance
		}

		neighbors := []CaveLoc{
			{cur.Pos.X - 1, cur.Pos.Y},
			{cur.Pos.X, cur.Pos.Y + 1},
			{cur.Pos.X + 1, cur.Pos.Y},
			{cur.Pos.X, cur.Pos.Y - 1},
		}

		for _, n := range neighbors {
			if n.X < 0 || n.X >= cave.Width() || n.Y < 0 || n.Y >= cave.Height() {
				continue
			}

			dist := cur.BestDistance + cave[n.Y][n.X]

			loc, found := locByCaveLoc[n]
			if !found {
				loc = &routeLoc{Pos: n, BestDistance: MaxInt}
				locByCaveLoc[n] = loc
			}

			if dist >= loc.BestDistance {
				// we have already a shorter route to n
				continue
			}

			loc.BestDistance = dist
			loc.EstimatedTotal = dist + estimateDist(n, end)
			if !loc.IsOpen {
				loc.IsOpen = true
				openSet = append(openSet, loc)
			}
		}

		sort.Sort(byEstimatedTotal(openSet))
	}

	// should never happen
	panic(fmt.Sprintf("unable to find a route (start=%v, end=%v)", start, end))
}

// estimateDist provide a distance estimate between two cave locations.
func estimateDist(a, b CaveLoc) int {
	// we are moving in a grid, so manhattan distance works nicely for this
	x := a.X - b.X
	if x < 0 {
		x *= -1
	}
	y := a.Y - b.Y
	if y < 0 {
		y *= -1
	}
	return x + y
}

type routeLoc struct {
	Pos            CaveLoc
	IsOpen         bool
	BestDistance   int
	EstimatedTotal int
}

type byEstimatedTotal []*routeLoc

func (p byEstimatedTotal) Less(i, j int) bool {
	return p[i].EstimatedTotal < p[j].EstimatedTotal
}
func (p byEstimatedTotal) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p byEstimatedTotal) Len() int      { return len(p) }
