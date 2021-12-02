package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

func init() { days["2"] = Day2 }

/*
--- Day 2: Dive! ---
Now, you need to figure out how to pilot this thing.

It seems like the submarine can take a series of commands like forward 1, down 2, or up 3:

forward X increases the horizontal position by X units.
down X increases the depth by X units.
up X decreases the depth by X units.
Note that since you're on a submarine, down and up affect your depth, and so they have the opposite result of what you might expect.

The submarine seems to already have a planned course (your puzzle input). You should probably figure out where it's going. For example:

forward 5
down 5
forward 8
up 3
down 8
forward 2
Your horizontal position and depth both start at 0. The steps above would then modify them as follows:

forward 5 adds 5 to your horizontal position, a total of 5.
down 5 adds 5 to your depth, resulting in a value of 5.
forward 8 adds 8 to your horizontal position, a total of 13.
up 3 decreases your depth by 3, resulting in a value of 2.
down 8 adds 8 to your depth, resulting in a value of 10.
forward 2 adds 2 to your horizontal position, a total of 15.
After following these instructions, you would have a horizontal position of 15 and a depth of 10. (Multiplying these together produces 150.)

Calculate the horizontal position and depth you would have after following the planned course. What do you get if you multiply your final horizontal position by your final depth?

Your puzzle answer was 1882980.

--- Part Two ---
Based on your calculations, the planned course doesn't seem to make any sense. You find the submarine manual and discover that the process is actually slightly more complicated.

In addition to horizontal position and depth, you'll also need to track a third value, aim, which also starts at 0. The commands also mean something entirely different than you first thought:

down X increases your aim by X units.
up X decreases your aim by X units.
forward X does two things:
It increases your horizontal position by X units.
It increases your depth by your aim multiplied by X.
Again note that since you're on a submarine, down and up do the opposite of what you might expect: "down" means aiming in the positive direction.

Now, the above example does something different:

forward 5 adds 5 to your horizontal position, a total of 5. Because your aim is 0, your depth does not change.
down 5 adds 5 to your aim, resulting in a value of 5.
forward 8 adds 8 to your horizontal position, a total of 13. Because your aim is 5, your depth increases by 8*5=40.
up 3 decreases your aim by 3, resulting in a value of 2.
down 8 adds 8 to your aim, resulting in a value of 10.
forward 2 adds 2 to your horizontal position, a total of 15. Because your aim is 10, your depth increases by 2*10=20 to a total of 60.
After following these new instructions, you would have a horizontal position of 15 and a depth of 60. (Multiplying these produces 900.)

Using this new interpretation of the commands, calculate the horizontal position and depth you would have after following the planned course. What do you get if you multiply your final horizontal position by your final depth?

Your puzzle answer was 1971232560.
*/

func Day2() {
	fmt.Println("--- Day 2: Dive! ---")
	c, err := NewCourse(day02Input)
	if err != nil {
		fmt.Println(errors.Wrap(err, "Unable to create new course"))
		return
	}

	pos, dep := c.SimpleEndPosition()
	fmt.Println("Position multiplied by depth after planned course (without aim):", pos*dep)
	pos, dep = c.RealEndPosition()
	fmt.Println("Position multiplied by depth after planned course (with aim):", pos*dep)
}

// Course is a sequence of commands for the submarine to follow to get to its destination.
type Course []Command

// Command is a single instruction to move the submarine where Action specifies how to move and
// N specifies by what amount.
type Command struct {
	Action Action
	N      int
}

// Action defines the different ways the submarine can move.
type Action uint8

const (
	ForwardAction Action = iota
	DownAction
	UpAction
)

// NewCourse takes the puzzle input and returns a Course.
// NewCourse will return an error if the input contains anything unexpected.
func NewCourse(in string) (Course, error) {
	var course Course
	cmdExp := regexp.MustCompile(`^(forward|down|up) ([0-9]+)$`)
	for _, line := range strings.Split(in, "\n") {
		m := cmdExp.FindStringSubmatch(line)
		if len(m) != 3 {
			return nil, errors.Errorf("unknown command (command=%s)", line)
		}
		n, err := strconv.Atoi(m[2])
		if err != nil {
			// should never happen
			panic(fmt.Sprintf("unable to parse n from command (n=%s, line=%s)", m[2], line))
		}
		var action Action
		switch m[1] {
		case "forward":
			action = ForwardAction
		case "down":
			action = DownAction
		case "up":
			action = UpAction
		default:
			// should never happen
			panic(fmt.Sprintf("unknown action (action=%s, line=%s)", m[1], line))
		}

		course = append(course, Command{Action: action, N: n})
	}
	return course, nil
}

// SimpleEndPosition calculate the final position of the submarine after having followed
// the course and returns the horizontal position, depth.
// SimpleEndPosition does not use aim to calculate depth.
func (c Course) SimpleEndPosition() (int, int) {
	pos, dep := 0, 0

	for _, cmd := range c {
		switch cmd.Action {
		case ForwardAction:
			pos += cmd.N
		case DownAction:
			dep += cmd.N
		case UpAction:
			dep -= cmd.N
		default:
			panic(fmt.Sprintf("unknown action (action=%v)", cmd.Action))
		}
	}

	return pos, dep
}

// RealEndPosition calculates the final position of the submarine after having followed
// the course and returns the horizontal position, depth.
// RealEndPosition correctly use aim to calculate the depth.
func (c Course) RealEndPosition() (int, int) {
	pos, dep, aim := 0, 0, 0

	for _, cmd := range c {
		switch cmd.Action {
		case ForwardAction:
			pos += cmd.N
			dep += aim * cmd.N
		case DownAction:
			aim += cmd.N
		case UpAction:
			aim -= cmd.N
		default:
			panic(fmt.Sprintf("unknown action (action=%v)", cmd.Action))
		}
	}

	return pos, dep
}
