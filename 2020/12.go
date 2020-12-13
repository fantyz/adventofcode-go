package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

func init() { days["12"] = Day12 }

/*
--- Day 12: Rain Risk ---
Your ferry made decent progress toward the island, but the storm came in faster than anyone expected. The ferry needs to take evasive actions!

Unfortunately, the ship's navigation computer seems to be malfunctioning; rather than giving a route directly to safety, it produced extremely circuitous instructions. When the captain uses the PA system to ask if anyone can help, you quickly volunteer.

The navigation instructions (your puzzle input) consists of a sequence of single-character actions paired with integer input values. After staring at them for a few minutes, you work out what they probably mean:

Action N means to move north by the given value.
Action S means to move south by the given value.
Action E means to move east by the given value.
Action W means to move west by the given value.
Action L means to turn left the given number of degrees.
Action R means to turn right the given number of degrees.
Action F means to move forward by the given value in the direction the ship is currently facing.
The ship starts by facing east. Only the L and R actions change the direction the ship is facing. (That is, if the ship is facing east and the next instruction is N10, the ship would move north 10 units, but would still move east if the following action were F.)

For example:

F10
N3
F7
R90
F11
These instructions would be handled as follows:

F10 would move the ship 10 units east (because the ship starts by facing east) to east 10, north 0.
N3 would move the ship 3 units north to east 10, north 3.
F7 would move the ship another 7 units east (because the ship is still facing east) to east 17, north 3.
R90 would cause the ship to turn right by 90 degrees and face south; it remains at east 17, north 3.
F11 would move the ship 11 units south to east 17, south 8.
At the end of these instructions, the ship's Manhattan distance (sum of the absolute values of its east/west position and its north/south position) from its starting position is 17 + 8 = 25.

Figure out where the navigation instructions lead. What is the Manhattan distance between that location and the ship's starting position?

Your puzzle answer was 962.

--- Part Two ---
Before you can give the destination to the captain, you realize that the actual action meanings were printed on the back of the instructions the whole time.

Almost all of the actions indicate how to move a waypoint which is relative to the ship's position:

Action N means to move the waypoint north by the given value.
Action S means to move the waypoint south by the given value.
Action E means to move the waypoint east by the given value.
Action W means to move the waypoint west by the given value.
Action L means to rotate the waypoint around the ship left (counter-clockwise) the given number of degrees.
Action R means to rotate the waypoint around the ship right (clockwise) the given number of degrees.
Action F means to move forward to the waypoint a number of times equal to the given value.
The waypoint starts 10 units east and 1 unit north relative to the ship. The waypoint is relative to the ship; that is, if the ship moves, the waypoint moves with it.

For example, using the same instructions as above:

F10 moves the ship to the waypoint 10 times (a total of 100 units east and 10 units north), leaving the ship at east 100, north 10. The waypoint stays 10 units east and 1 unit north of the ship.
N3 moves the waypoint 3 units north to 10 units east and 4 units north of the ship. The ship remains at east 100, north 10.
F7 moves the ship to the waypoint 7 times (a total of 70 units east and 28 units north), leaving the ship at east 170, north 38. The waypoint stays 10 units east and 4 units north of the ship.
R90 rotates the waypoint around the ship clockwise 90 degrees, moving it to 4 units east and 10 units south of the ship. The ship remains at east 170, north 38.
F11 moves the ship to the waypoint 11 times (a total of 44 units east and 110 units south), leaving the ship at east 214, south 72. The waypoint stays 4 units east and 10 units south of the ship.
After these operations, the ship's Manhattan distance from its starting position is 214 + 72 = 286.

Figure out where the navigation instructions actually lead. What is the Manhattan distance between that location and the ship's starting position?

Your puzzle answer was 56135.
*/

func Day12() {
	fmt.Println("--- Day 12: Rain Risk ---")
	instructions, err := LoadShipInstructions(day12Input)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Manhattan distance to ship after instructions:", ManhattanDistanceToShipAfterInstructions(instructions))
	fmt.Println("Manhattan distance to ship with waypoint after instructions:", ManhattanDistanceToShipWithWaypointAfterInstructions(10, 1, instructions))
}

// LoadShipInstructions loads a list of ship instructions and return them. LoadShipInstructions
// will return an error if it is unable to do that.
func LoadShipInstructions(in string) ([]ShipInstruction, error) {
	var inst []ShipInstruction
	for _, s := range strings.Split(in, "\n") {
		v, err := strconv.Atoi(s[1:])
		if err != nil {
			return nil, errors.Wrapf(err, "bad ship instruction (inst=%s)", s)
		}
		inst = append(inst, ShipInstruction{s[0], v})
	}
	return inst, nil
}

// ManhattanDistanceToShipAfterInstructions takes a list of ship instructions and returns
// the manhattan distance to the location of the ship after after executing them.
func ManhattanDistanceToShipAfterInstructions(instructions []ShipInstruction) int {
	s := &Ship{}
	s.FollowInstructions(instructions)
	return ManhattanDistance(s.X, s.Y)
}

// ManhattanDistanceToShiWithWaypointpAfterInstructions takes a list of ship instructions
// and returns the manhattan distance to the location of the ship after after executing them.
func ManhattanDistanceToShipWithWaypointAfterInstructions(x, y int, instructions []ShipInstruction) int {
	s := &ShipWithWaypoint{XWay: x, YWay: y}
	s.FollowInstructions(instructions)
	return ManhattanDistance(s.X, s.Y)
}

// ManhattanDistance takes a coordinate pair and returns the manhattan distance from these to 0,0.
func ManhattanDistance(x, y int) int {
	if x < 0 {
		x = -x
	}
	if y < 0 {
		y = -y
	}
	return x + y
}

// ShipInstruction represents a single ship instruction.
type ShipInstruction struct {
	Action byte
	Value  int
}

// Heading is used to represent the headings needed. It use an int and orders the for headings in such
// a way that Heading+1 represents a right turn and Heading-1 represents a left turn. This makes it
// easier to implement left and right turns.
type Heading int

const (
	East Heading = iota
	South
	West
	North
)

// Ship represents a ship that does not use waypoints for navigation.
type Ship struct {
	Heading Heading
	X, Y    int
}

// FollowInstructions takes a list of instructions and execute them one by one.
func (s *Ship) FollowInstructions(instructions []ShipInstruction) {
	for _, inst := range instructions {
		s.Do(inst)
	}
}

// Do takes an instruction and executes it.
func (s *Ship) Do(inst ShipInstruction) {
	switch inst.Action {
	case 'N':
		s.Y += inst.Value
	case 'S':
		s.Y -= inst.Value
	case 'E':
		s.X += inst.Value
	case 'W':
		s.X -= inst.Value
	case 'L':
		if inst.Value%90 != 0 {
			panic("Not able to handle turns in less than 90 degree intervals")
		}

		// left turn is done by subracting from number of 90 degree turns from the ships heading
		s.Heading = Heading(int(s.Heading) - inst.Value/90)
		for s.Heading < 0 {
			// make sure we get a negative heading translated back to the four possible headings (eg. -1 -> 3)
			s.Heading += Heading(4)
		}
	case 'R':
		if inst.Value%90 != 0 {
			panic("Not able to handle turns in less than 90 degree intervals")
		}
		// right turn is done by adding th enumber of 90 degree turns to the ships heading
		// any value higher than 3 is dealt with using %4 (eg. 5 -> 1)
		s.Heading = Heading((inst.Value/90 + int(s.Heading)) % 4)
	case 'F':
		var a byte
		switch s.Heading {
		case East:
			a = 'E'
		case South:
			a = 'S'
		case West:
			a = 'W'
		case North:
			a = 'N'
		default:
			panic(fmt.Sprintf("unhandled heading (heading=%d)", s.Heading))
		}
		s.Do(ShipInstruction{a, inst.Value})
	}
}

// ShipWithWaypoint represents a ship using a waypoint for navigation.
type ShipWithWaypoint struct {
	X, Y       int
	XWay, YWay int
}

// FollowInstructions takes a list of instructions and execute them one by one.
func (s *ShipWithWaypoint) FollowInstructions(instructions []ShipInstruction) {
	for _, inst := range instructions {
		s.Do(inst)
	}
}

// Do takes an instruction and executes it.
func (s *ShipWithWaypoint) Do(inst ShipInstruction) {
	switch inst.Action {
	case 'N':
		s.YWay += inst.Value
	case 'S':
		s.YWay -= inst.Value
	case 'E':
		s.XWay += inst.Value
	case 'W':
		s.XWay -= inst.Value
	case 'L':
		if inst.Value%90 != 0 {
			panic("Not able to handle turns in less than 90 degree intervals")
		}
		for i := 0; i < inst.Value/90; i++ {
			s.XWay, s.YWay = -s.YWay, s.XWay
		}
	case 'R':
		if inst.Value%90 != 0 {
			panic("Not able to handle turns in less than 90 degree intervals")
		}
		for i := 0; i < inst.Value/90; i++ {
			s.XWay, s.YWay = s.YWay, -s.XWay
		}
	case 'F':
		s.X += inst.Value * s.XWay
		s.Y += inst.Value * s.YWay
	}
}
