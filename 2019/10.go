package main

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

/*

--- Day 10: Monitoring Station ---
You fly into the asteroid belt and reach the Ceres monitoring station. The Elves here have an emergency: they're having trouble tracking all of the asteroids and can't be sure they're safe.

The Elves would like to build a new monitoring station in a nearby area of space; they hand you a map of all of the asteroids in that region (your puzzle input).

The map indicates whether each position is empty (.) or contains an asteroid (#). The asteroids are much smaller than they appear on the map, and every asteroid is exactly in the center of its marked position. The asteroids can be described with X,Y coordinates where X is the distance from the left edge and Y is the distance from the top edge (so the top-left corner is 0,0 and the position immediately to its right is 1,0).

Your job is to figure out which asteroid would be the best place to build a new monitoring station. A monitoring station can detect any asteroid to which it has direct line of sight - that is, there cannot be another asteroid exactly between them. This line of sight can be at any angle, not just lines aligned to the grid or diagonally. The best location is the asteroid that can detect the largest number of other asteroids.

For example, consider the following map:

.#..#
.....
#####
....#
...##
The best location for a new monitoring station on this map is the highlighted asteroid at 3,4 because it can detect 8 asteroids, more than any other location. (The only asteroid it cannot detect is the one at 1,0; its view of this asteroid is blocked by the asteroid at 2,2.) All other asteroids are worse locations; they can detect 7 or fewer other asteroids. Here is the number of other asteroids a monitoring station on each asteroid could detect:

.7..7
.....
67775
....7
...87
Here is an asteroid (#) and some examples of the ways its line of sight might be blocked. If there were another asteroid at the location of a capital letter, the locations marked with the corresponding lowercase letter would be blocked and could not be detected:

#.........
...A......
...B..a...
.EDCG....a
..F.c.b...
.....c....
..efd.c.gb
.......c..
....f...c.
...e..d..c
Here are some larger examples:

Best is 5,8 with 33 other asteroids detected:

......#.#.
#..#.#....
..#######.
.#.#.###..
.#..#.....
..#....#.#
#..#....#.
.##.#..###
##...#..#.
.#....####
Best is 1,2 with 35 other asteroids detected:

#.#...#.#.
.###....#.
.#....#...
##.#.#.#.#
....#.#.#.
.##..###.#
..#...##..
..##....##
......#...
.####.###.
Best is 6,3 with 41 other asteroids detected:

.#..#..###
####.###.#
....###.#.
..###.##.#
##.##.#.#.
....###..#
..#.#..#.#
#..#.#.###
.##...##.#
.....#.#..
Best is 11,13 with 210 other asteroids detected:

.#..##.###...#######
##.############..##.
.#.######.########.#
.###.#######.####.#.
#####.##.#.##.###.##
..#####..#.#########
####################
#.####....###.#.#.##
##.#################
#####.##.###..####..
..######..##.#######
####.##.####...##..#
.#####..#.######.###
##...#.##########...
#.##########.#######
.####.#.###.###.#.##
....##.##.###..#####
.#.#.###########.###
#.#.#.#####.####.###
###.##.####.##.#..##
Find the best location for a new monitoring station. How many other asteroids can be detected from that location?

Your puzzle answer was 278.

--- Part Two ---
Once you give them the coordinates, the Elves quickly deploy an Instant Monitoring Station to the location and discover the worst: there are simply too many asteroids.

The only solution is complete vaporization by giant laser.

Fortunately, in addition to an asteroid scanner, the new monitoring station also comes equipped with a giant rotating laser perfect for vaporizing asteroids. The laser starts by pointing up and always rotates clockwise, vaporizing any asteroid it hits.

If multiple asteroids are exactly in line with the station, the laser only has enough power to vaporize one of them before continuing its rotation. In other words, the same asteroids that can be detected can be vaporized, but if vaporizing one asteroid makes another one detectable, the newly-detected asteroid won't be vaporized until the laser has returned to the same position by rotating a full 360 degrees.

For example, consider the following map, where the asteroid with the new monitoring station (and laser) is marked X:

.#....#####...#..
##...##.#####..##
##...#...#.#####.
..#.....X...###..
..#.#.....#....##
The first nine asteroids to get vaporized, in order, would be:

.#....###24...#..
##...##.13#67..9#
##...#...5.8####.
..#.....X...###..
..#.#.....#....##
Note that some asteroids (the ones behind the asteroids marked 1, 5, and 7) won't have a chance to be vaporized until the next full rotation. The laser continues rotating; the next nine to be vaporized are:

.#....###.....#..
##...##...#.....#
##...#......1234.
..#.....X...5##..
..#.9.....8....76
The next nine to be vaporized are then:

.8....###.....#..
56...9#...#.....#
34...7...........
..2.....X....##..
..1..............
Finally, the laser completes its first full rotation (1 through 3), a second rotation (4 through 8), and vaporizes the last asteroid (9) partway through its third rotation:

......234.....6..
......1...5.....7
.................
........X....89..
.................
In the large example above (the one with the best monitoring station location at 11,13):

The 1st asteroid to be vaporized is at 11,12.
The 2nd asteroid to be vaporized is at 12,1.
The 3rd asteroid to be vaporized is at 12,2.
The 10th asteroid to be vaporized is at 12,8.
The 20th asteroid to be vaporized is at 16,0.
The 50th asteroid to be vaporized is at 16,9.
The 100th asteroid to be vaporized is at 10,16.
The 199th asteroid to be vaporized is at 9,6.
The 200th asteroid to be vaporized is at 8,2.
The 201st asteroid to be vaporized is at 10,9.
The 299th and final asteroid to be vaporized is at 11,1.
The Elves are placing bets on which will be the 200th asteroid to be vaporized. Win the bet by determining which asteroid that will be; what do you get if you multiply its X coordinate by 100 and then add its Y coordinate? (For example, 8,2 becomes 802.)

Your puzzle answer was 1417.

*/

func main() {
	fmt.Println("Day 10: Monitoring Station")
	x, y, count := BestMonitoringStation(Load(puzzleInput))
	fmt.Println("Detected asteroids from best position:", count)
	xshot, yshot := ShootAsteroid(x, y, Load(puzzleInput), 200)
	fmt.Println("200th asteroid to be shot would be ", xshot, yshot)
}

type Asteroid struct {
	x, y        int
	angle, dist float64
	dead        bool
}

func ShootAsteroid(xbase, ybase int, field Field, n int) (int, int) {
	f := field.Copy()

	f[ybase][xbase] = false

	var asteroids []Asteroid
	for y := 0; y < len(f); y++ {
		for x := 0; x < len(f[y]); x++ {
			if !f[y][x] {
				continue
			}
			asteroids = append(asteroids, Asteroid{x, y, AngleToY(x-xbase, ybase-y), Dist(x-xbase, ybase-y), false})
		}
	}

	sort.Slice(asteroids, func(i, j int) bool {
		if asteroids[i].angle == asteroids[j].angle {
			return asteroids[i].dist < asteroids[j].dist
		}
		return asteroids[i].angle < asteroids[j].angle
	})
	/*
	 */
	cur := 0
	for i := 0; i < n-1; i++ {
		asteroids[cur].dead = true
		angle := asteroids[cur].angle

		// find next one
		for {
			cur++
			if cur >= len(asteroids) {
				cur = 0
				break
			}
			if !asteroids[cur].dead && angle+0.0001 < asteroids[cur].angle {
				break
			}
		}
	}

	for i := cur; i < len(asteroids); i++ {
		if !asteroids[i].dead {
			/*
				for i, a := range asteroids {
					if i == cur {
						fmt.Println("====================")
					}
					fmt.Println(a)
					if i == cur {
						fmt.Println("====================")
					}
				}
			*/

			return asteroids[i].x, asteroids[i].y
		}
	}

	panic("no living asteroids left..?")
}

func BestMonitoringStation(field Field) (int, int, int) {
	best := 0
	bestX, bestY := 0, 0
	for y := 0; y < len(field); y++ {
		for x := 0; x < len(field[y]); x++ {
			if !field[y][x] {
				continue
			}
			los := LineOfSight(x, y, field)
			if los > best {
				best = los
				bestX = x
				bestY = y
			}
		}
	}

	return bestX, bestY, best
}

type Coord struct {
	X, Y int
}

func LineOfSight(xfrom, yfrom int, field Field) int {
	counted := map[Coord]bool{}
	counted[Coord{xfrom, yfrom}] = true

	count := 0
	for y := 0; y < len(field); y++ {
		for x := 0; x < len(field[y]); x++ {
			if _, found := counted[Coord{x, y}]; !field[y][x] || found {
				continue
			}

			count++

			// remove any that are blocked by this one
			xstep := x - xfrom
			ystep := y - yfrom
			gcd := GCD(xstep, ystep)
			xstep = xstep / gcd
			ystep = ystep / gcd

			//fmt.Println("from", xfrom, yfrom, " - x,y", x, y, " step", xstep, ystep)
			//field.Print()

			for i := 1; xstep*i+xfrom >= 0 && xstep*i+xfrom < len(field[0]) && ystep*i+yfrom >= 0 && ystep*i+yfrom < len(field); i++ {
				if field[ystep*i+yfrom][xstep*i+xfrom] {
					counted[Coord{xstep*i + xfrom, ystep*i + yfrom}] = true
				}
			}
		}
	}
	return count
}

func GCD(a, b int) int {
	if a < 0 {
		a *= -1
	}
	if b < 0 {
		b *= -1
	}
	if a < b {
		return GCD(b, a)
	}
	if b == 0 {
		return a
	}
	return GCD(b, a%b)
}

func AngleToY(x, y int) float64 {
	// angle = cos^-1(y / sqrt(x^2+y^2))
	if x >= 0 {
		return math.Acos(float64(y)/Dist(x, y)) * (180 / math.Pi)
	} else {
		return 360 - math.Acos(float64(y)/Dist(x, y))*(180/math.Pi)
	}
}

func Dist(x, y int) float64 {
	return math.Sqrt(math.Pow(float64(x), 2) + math.Pow(float64(y), 2))
}

type Field [][]bool

func (f Field) Print() {
	for y := 0; y < len(f); y++ {
		for x := 0; x < len(f[y]); x++ {
			if f[y][x] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Print("\n")
	}
}

func (f Field) Copy() Field {
	new := make(Field, len(f))
	for i := 0; i < len(f); i++ {
		new[i] = make([]bool, len(f[i]))
		copy(new[i], f[i])
	}
	return new
}

func Load(in string) Field {
	var field Field
	for _, l := range strings.Split(in, "\n") {
		row := make([]bool, 0, len(l))
		for _, c := range l {
			switch c {
			case '.':
				row = append(row, false)
			case '#':
				row = append(row, true)
			default:
				panic(fmt.Sprintf("unknown field char %s", string(c)))
			}
		}
		field = append(field, row)
	}
	return field
}

const puzzleInput = `.#......#...#.....#..#......#..##..#
..#.......#..........#..##.##.......
##......#.#..#..#..##...#.##.###....
..#........#...........#.......##...
.##.....#.......#........#..#.#.....
.#...#...#.....#.##.......#...#....#
#...#..##....#....#......#..........
....#......#.#.....#..#...#......#..
......###.......#..........#.##.#...
#......#..#.....#..#......#..#..####
.##...##......##..#####.......##....
.....#...#.........#........#....#..
....##.....#...#........#.##..#....#
....#........#.###.#........#...#..#
....#..#.#.##....#.........#.....#.#
##....###....##..#..#........#......
.....#.#.........#.......#....#....#
.###.....#....#.#......#...##.##....
...##...##....##.........#...#......
.....#....##....#..#.#.#...##.#...#.
#...#.#.#.#..##.#...#..#..#..#......
......#...#...#.#.....#.#.....#.####
..........#..................#.#.##.
....#....#....#...#..#....#.....#...
.#####..####........#...............
#....#.#..#..#....##......#...#.....
...####....#..#......#.#...##.....#.
..##....#.###.##.#.##.#.....#......#
....#.####...#......###.....##......
.#.....#....#......#..#..#.#..#.....
..#.......#...#........#.##...#.....
#.....####.#..........#.#.......#...
..##..#..#.....#.#.........#..#.#.##
.........#..........##.#.##.......##
#..#.....#....#....#.#.......####..#
..............#.#...........##.#.#..`
