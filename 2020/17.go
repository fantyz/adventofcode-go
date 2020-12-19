package main

import (
	"fmt"
	"strings"
)

func init() { days["17"] = Day17 }

/*
--- Day 17: Conway Cubes ---
As your flight slowly drifts through the sky, the Elves at the Mythical Information Bureau at the North Pole contact you. They'd like some help debugging a malfunctioning experimental energy source aboard one of their super-secret imaging satellites.

The experimental energy source is based on cutting-edge technology: a set of Conway Cubes contained in a pocket dimension! When you hear it's having problems, you can't help but agree to take a look.

The pocket dimension contains an infinite 3-dimensional grid. At every integer 3-dimensional coordinate (x,y,z), there exists a single cube which is either active or inactive.

In the initial state of the pocket dimension, almost all cubes start inactive. The only exception to this is a small flat region of cubes (your puzzle input); the cubes in this region start in the specified active (#) or inactive (.) state.

The energy source then proceeds to boot up by executing six cycles.

Each cube only ever considers its neighbors: any of the 26 other cubes where any of their coordinates differ by at most 1. For example, given the cube at x=1,y=2,z=3, its neighbors include the cube at x=2,y=2,z=2, the cube at x=0,y=2,z=3, and so on.

During a cycle, all cubes simultaneously change their state according to the following rules:

If a cube is active and exactly 2 or 3 of its neighbors are also active, the cube remains active. Otherwise, the cube becomes inactive.
If a cube is inactive but exactly 3 of its neighbors are active, the cube becomes active. Otherwise, the cube remains inactive.
The engineers responsible for this experimental energy source would like you to simulate the pocket dimension and determine what the configuration of cubes should be at the end of the six-cycle boot process.

For example, consider the following initial state:

.#.
..#
###
Even though the pocket dimension is 3-dimensional, this initial state represents a small 2-dimensional slice of it. (In particular, this initial state defines a 3x3x1 region of the 3-dimensional space.)

Simulating a few cycles from this initial state produces the following configurations, where the result of each cycle is shown layer-by-layer at each given z coordinate (and the frame of view follows the active cells in each cycle):

Before any cycles:

z=0
.#.
..#
###


After 1 cycle:

z=-1
#..
..#
.#.

z=0
#.#
.##
.#.

z=1
#..
..#
.#.


After 2 cycles:

z=-2
.....
.....
..#..
.....
.....

z=-1
..#..
.#..#
....#
.#...
.....

z=0
##...
##...
#....
....#
.###.

z=1
..#..
.#..#
....#
.#...
.....

z=2
.....
.....
..#..
.....
.....


After 3 cycles:

z=-2
.......
.......
..##...
..###..
.......
.......
.......

z=-1
..#....
...#...
#......
.....##
.#...#.
..#.#..
...#...

z=0
...#...
.......
#......
.......
.....##
.##.#..
...#...

z=1
..#....
...#...
#......
.....##
.#...#.
..#.#..
...#...

z=2
.......
.......
..##...
..###..
.......
.......
.......
After the full six-cycle boot process completes, 112 cubes are left in the active state.

Starting with your given initial configuration, simulate six cycles. How many cubes are left in the active state after the sixth cycle?

Your puzzle answer was 372.

--- Part Two ---
For some reason, your simulated results don't match what the experimental energy source engineers expected. Apparently, the pocket dimension actually has four spatial dimensions, not three.

The pocket dimension contains an infinite 4-dimensional grid. At every integer 4-dimensional coordinate (x,y,z,w), there exists a single cube (really, a hypercube) which is still either active or inactive.

Each cube only ever considers its neighbors: any of the 80 other cubes where any of their coordinates differ by at most 1. For example, given the cube at x=1,y=2,z=3,w=4, its neighbors include the cube at x=2,y=2,z=3,w=3, the cube at x=0,y=2,z=3,w=4, and so on.

The initial state of the pocket dimension still consists of a small flat region of cubes. Furthermore, the same rules for cycle updating still apply: during each cycle, consider the number of active neighbors of each cube.

For example, consider the same initial state as in the example above. Even though the pocket dimension is 4-dimensional, this initial state represents a small 2-dimensional slice of it. (In particular, this initial state defines a 3x3x1x1 region of the 4-dimensional space.)

Simulating a few cycles from this initial state produces the following configurations, where the result of each cycle is shown layer-by-layer at each given z and w coordinate:

Before any cycles:

z=0, w=0
.#.
..#
###


After 1 cycle:

z=-1, w=-1
#..
..#
.#.

z=0, w=-1
#..
..#
.#.

z=1, w=-1
#..
..#
.#.

z=-1, w=0
#..
..#
.#.

z=0, w=0
#.#
.##
.#.

z=1, w=0
#..
..#
.#.

z=-1, w=1
#..
..#
.#.

z=0, w=1
#..
..#
.#.

z=1, w=1
#..
..#
.#.


After 2 cycles:

z=-2, w=-2
.....
.....
..#..
.....
.....

z=-1, w=-2
.....
.....
.....
.....
.....

z=0, w=-2
###..
##.##
#...#
.#..#
.###.

z=1, w=-2
.....
.....
.....
.....
.....

z=2, w=-2
.....
.....
..#..
.....
.....

z=-2, w=-1
.....
.....
.....
.....
.....

z=-1, w=-1
.....
.....
.....
.....
.....

z=0, w=-1
.....
.....
.....
.....
.....

z=1, w=-1
.....
.....
.....
.....
.....

z=2, w=-1
.....
.....
.....
.....
.....

z=-2, w=0
###..
##.##
#...#
.#..#
.###.

z=-1, w=0
.....
.....
.....
.....
.....

z=0, w=0
.....
.....
.....
.....
.....

z=1, w=0
.....
.....
.....
.....
.....

z=2, w=0
###..
##.##
#...#
.#..#
.###.

z=-2, w=1
.....
.....
.....
.....
.....

z=-1, w=1
.....
.....
.....
.....
.....

z=0, w=1
.....
.....
.....
.....
.....

z=1, w=1
.....
.....
.....
.....
.....

z=2, w=1
.....
.....
.....
.....
.....

z=-2, w=2
.....
.....
..#..
.....
.....

z=-1, w=2
.....
.....
.....
.....
.....

z=0, w=2
###..
##.##
#...#
.#..#
.###.

z=1, w=2
.....
.....
.....
.....
.....

z=2, w=2
.....
.....
..#..
.....
.....
After the full six-cycle boot process completes, 848 cubes are left in the active state.

Starting with your given initial configuration, simulate six cycles in a 4-dimensional space. How many cubes are left in the active state after the sixth cycle?

Your puzzle answer was 1896.
*/

func Day17() {
	fmt.Println("--- Day 17: Conway Cubes ---")
	p := NewPocketDimension3D(day17Input)
	p.RunCycles(6)
	fmt.Println("Active cubes in 3D after 6 cycles:", p.ActiveCubes())
	p2 := NewPocketDimension4D(day17Input)
	p2.RunCycles(6)
	fmt.Println("Active cubes in 4D after 6 cycles:", p2.ActiveCubes())

}

// NewPocketDimension3D creates a new pocket dimension in 3D initialized with the intial layout provided.
func NewPocketDimension3D(initialLayout string) *PocketDimension3D {
	p := &PocketDimension3D{
		cubes: make(map[Coord3D]struct{}),
	}
	for y, line := range strings.Split(initialLayout, "\n") {
		for x := 0; x < len(line); x++ {
			if line[x] == '#' {
				p.Activate(Coord3D{x, y, 0})
			}
		}
	}
	return p
}

// PocketDimension is modeled as a map containing the 3 dimensional coordinates of all activated
// cubes along with the size of the universe represented by the minimum and maximum coordinates
// used in each dimension.
//
// This is easier to implement this way than using arrays given the dynamic growth of the universe.
// Using maps will be slightly slower than arrays but still running in constant time.
type PocketDimension3D struct {
	cubes      map[Coord3D]struct{}
	minV, maxV [3]int
}

// Coord3D represents a coordinate in 3D space.
type Coord3D [3]int

// Activate will activate the cube at the specified coordinate and grow the pocket
// universe if needed.
func (p *PocketDimension3D) Activate(coord Coord3D) {
	// expand the size of the pocket universe if needed
	for dim, val := range coord {
		if p.minV[dim] > val {
			p.minV[dim] = val
		}
		if p.maxV[dim] < val {
			p.maxV[dim] = val
		}
	}

	// activate cube
	p.cubes[coord] = struct{}{}
}

// Neighbors returns the number of neighboring cubes that are active.
func (p *PocketDimension3D) Neighbors(pos Coord3D) int {
	count := 0
	for z := pos[2] - 1; z <= pos[2]+1; z++ {
		for y := pos[1] - 1; y <= pos[1]+1; y++ {
			for x := pos[0] - 1; x <= pos[0]+1; x++ {
				coord := Coord3D{x, y, z}
				if pos == coord {
					// ignore the center
					continue
				}

				if _, active := p.cubes[coord]; active {
					count++
				}
			}
		}
	}
	return count
}

// ActiveCubes returns the number of active cubes in the pocket universe.
func (p *PocketDimension3D) ActiveCubes() int {
	return len(p.cubes)
}

// RunCycles performs the specified number of cycles
func (p *PocketDimension3D) RunCycles(cycles int) {
	for i := 0; i < cycles; i++ {
		p.Cycle()
	}
}

// Cycle will perform a cycle of the pocket universe.
func (p *PocketDimension3D) Cycle() {
	// create a new empty pocket universe
	pNew := NewPocketDimension3D("")

	// evaluate every coordinate in the universe that potentially could actiavet
	// a cycle can at most grow the universe by one in each direction
	for z := p.minV[2] - 1; z <= p.maxV[2]+1; z++ {
		for y := p.minV[1] - 1; y <= p.maxV[1]+1; y++ {
			for x := p.minV[0] - 1; x <= p.maxV[0]+1; x++ {
				coord := Coord3D{x, y, z}
				neighbors := p.Neighbors(coord)
				if _, isActive := p.cubes[coord]; isActive {
					// current cube is active
					if neighbors == 2 || neighbors == 3 {
						pNew.Activate(coord)
					}
				} else {
					// current cube is inactive
					if neighbors == 3 {
						pNew.Activate(coord)
					}
				}
			}
		}
	}
	*p = *pNew
}

// NewPocketDimension4D creates a new pocket dimension in 4D initialized with the intial layout provided.
func NewPocketDimension4D(initialLayout string) *PocketDimension4D {
	p := &PocketDimension4D{
		cubes: make(map[Coord4D]struct{}),
	}
	for y, line := range strings.Split(initialLayout, "\n") {
		for x := 0; x < len(line); x++ {
			if line[x] == '#' {
				p.Activate(Coord4D{x, y, 0, 0})
			}
		}
	}
	return p
}

// PocketDimension is modeled as a map containing the 4 dimensional coordinates of all activated
// cubes along with the size of the universe represented by the minimum and maximum coordinates
// used in each dimension.
//
// The lack of generics in Go is painful here. It is not possible to easily make a n-dimensional
// solution easily based on maps.
//
// Instead this solution is a copy-paste of the 3D solution and only adding the extra dimension.
type PocketDimension4D struct {
	cubes      map[Coord4D]struct{}
	minV, maxV [4]int
}

// Coord4D represents a coordinate in 4D space.
type Coord4D [4]int

// Activate will activate the cube at the specified coordinate and grow the pocket
// universe if needed.
func (p *PocketDimension4D) Activate(coord Coord4D) {
	// expand the size of the pocket universe if needed
	for dim, val := range coord {
		if p.minV[dim] > val {
			p.minV[dim] = val
		}
		if p.maxV[dim] < val {
			p.maxV[dim] = val
		}
	}

	// activate cube
	p.cubes[coord] = struct{}{}
}

// Neighbors returns the number of neighboring cubes that are active.
func (p *PocketDimension4D) Neighbors(pos Coord4D) int {
	count := 0
	for w := pos[3] - 1; w <= pos[3]+1; w++ {
		for z := pos[2] - 1; z <= pos[2]+1; z++ {
			for y := pos[1] - 1; y <= pos[1]+1; y++ {
				for x := pos[0] - 1; x <= pos[0]+1; x++ {
					coord := Coord4D{x, y, z, w}
					if pos == coord {
						// ignore the center
						continue
					}

					if _, active := p.cubes[coord]; active {
						count++
					}
				}
			}
		}
	}
	return count
}

// ActiveCubes returns the number of active cubes in the pocket universe.
func (p *PocketDimension4D) ActiveCubes() int {
	return len(p.cubes)
}

// RunCycles performs the specified number of cycles
func (p *PocketDimension4D) RunCycles(cycles int) {
	for i := 0; i < cycles; i++ {
		p.Cycle()
	}
}

// Cycle will perform a cycle of the pocket universe.
func (p *PocketDimension4D) Cycle() {
	// create a new empty pocket universe
	pNew := NewPocketDimension4D("")

	// evaluate every coordinate in the universe that potentially could actiavet
	// a cycle can at most grow the universe by one in each direction
	for w := p.minV[3] - 1; w <= p.maxV[3]+1; w++ {
		for z := p.minV[2] - 1; z <= p.maxV[2]+1; z++ {
			for y := p.minV[1] - 1; y <= p.maxV[1]+1; y++ {
				for x := p.minV[0] - 1; x <= p.maxV[0]+1; x++ {
					coord := Coord4D{x, y, z, w}
					neighbors := p.Neighbors(coord)
					if _, isActive := p.cubes[coord]; isActive {
						// current cube is active
						if neighbors == 2 || neighbors == 3 {
							pNew.Activate(coord)
						}
					} else {
						// current cube is inactive
						if neighbors == 3 {
							pNew.Activate(coord)
						}
					}
				}
			}
		}
	}
	*p = *pNew
}
