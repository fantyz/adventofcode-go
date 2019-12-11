package main

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

/*


 */

func main() {
	fmt.Println("Day 10: Monitoring Station")
	BestMonitoringStation(Load(puzzleInput))
	fmt.Println("Detected asteroids from best position:")
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

	cur := 0
	for i := 0; i < n; i++ {
		asteroids[cur].dead = true
		for j := cur + 1; j < len(asteroids); j++ {

		}

		lastAngle = asteroids[next].angle
		asteroids[next] = asteroids[len(asteroids)-1]
		asteroids = asteroids[:len(asteroids)-1]

		for j := 0; j < len(asteroids); j++ {
			if asteroids[j].angle == lastAngle {
				continue
			}
			next = j
		}

	}
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
