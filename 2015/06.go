package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

func init() { days["6"] = Day6 }

/*
--- Day 6: Probably a Fire Hazard ---
Because your neighbors keep defeating you in the holiday house decorating contest year after year, you've decided to deploy one million lights in a 1000x1000 grid.

Furthermore, because you've been especially nice this year, Santa has mailed you instructions on how to display the ideal lighting configuration.

Lights in your grid are numbered from 0 to 999 in each direction; the lights at each corner are at 0,0, 0,999, 999,999, and 999,0. The instructions include whether to turn on, turn off, or toggle various inclusive ranges given as coordinate pairs. Each coordinate pair represents opposite corners of a rectangle, inclusive; a coordinate pair like 0,0 through 2,2 therefore refers to 9 lights in a 3x3 square. The lights all start turned off.

To defeat your neighbors this year, all you have to do is set up your lights by doing the instructions Santa sent you in order.

For example:

turn on 0,0 through 999,999 would turn on (or leave on) every light.
toggle 0,0 through 999,0 would toggle the first line of 1000 lights, turning off the ones that were on, and turning on the ones that were off.
turn off 499,499 through 500,500 would turn off (or leave off) the middle four lights.
After following the instructions, how many lights are lit?

Your puzzle answer was 543903.

--- Part Two ---
You just finish implementing your winning light pattern when you realize you mistranslated Santa's message from Ancient Nordic Elvish.

The light grid you bought actually has individual brightness controls; each light can have a brightness of zero or more. The lights all start at zero.

The phrase turn on actually means that you should increase the brightness of those lights by 1.

The phrase turn off actually means that you should decrease the brightness of those lights by 1, to a minimum of zero.

The phrase toggle actually means that you should increase the brightness of those lights by 2.

What is the total brightness of all lights combined after following Santa's instructions?

For example:

turn on 0,0 through 0,0 would increase the total brightness by 1.
toggle 0,0 through 999,999 would increase the total brightness by 2000000.
Your puzzle answer was 14687245.
*/

func Day6() {
	fmt.Println("--- Day 6: Probably a Fire Hazard ---")
	tl := NewToggleLights(1000, 1000, false)
	_ = ExecuteInstructions(tl, day6Input)
	fmt.Println("        Lit Lights:", tl.LitLights())
	bl := NewBrightnessLights(1000, 1000, 0)
	_ = ExecuteInstructions(bl, day6Input)
	fmt.Println("  Total brightness:", bl.LitLights())
}

// Given it is not really possible to solve pt1 with the solution for pt2 in the same way as usual,
// the best solution is to have two different implementations of the light display. Using an interface
// allow us to keep the execute logic the same.
type Lights interface {
	TurnOn(int, int, int, int)
	TurnOff(int, int, int, int)
	Toggle(int, int, int, int)
	LitLights() int
}

// ExecuteInstructions takes lights and a set of instructions and executes these on the lights. An
// error is returned if any of the commands are invalid. Note that any instructions up to that point
// will have been executed on the lights.
func ExecuteInstructions(l Lights, in string) error {
	instExp := regexp.MustCompile(`^(turn on|turn off|toggle) ([0-9]+),([0-9]+) through ([0-9]+),([0-9]+)$`)

	for _, inst := range strings.Split(in, "\n") {
		m := instExp.FindStringSubmatch(inst)
		if len(m) != 6 {
			return errors.Errorf("instruction did not match regexp (inst=%s)", inst)
		}
		cmd := m[1]
		xStart, _ := strconv.Atoi(m[2])
		yStart, _ := strconv.Atoi(m[3])
		xEnd, _ := strconv.Atoi(m[4])
		yEnd, _ := strconv.Atoi(m[5])

		switch cmd {
		case "turn on":
			l.TurnOn(xStart, yStart, xEnd, yEnd)
		case "turn off":
			l.TurnOff(xStart, yStart, xEnd, yEnd)
		case "toggle":
			l.Toggle(xStart, yStart, xEnd, yEnd)
		}
	}

	return nil
}

// ToggleLights provide an implementation that simply have the indvidual lights turned on or off.
type ToggleLights [][]bool

// NewToggleLights creates a ToggleLights of the size and initial state specified.
func NewToggleLights(width, height int, lit bool) ToggleLights {
	l := make(ToggleLights, height)
	for i := 0; i < width; i++ {
		l[i] = make([]bool, width)
		if lit {
			for j := 0; j < width; j++ {
				l[i][j] = true
			}
		}
	}
	return l
}

// TurnOn will turn on all lights in the square specified (top left point to bottom right point).
// Any coordinates placed outside the available lights will be ignored.
func (l ToggleLights) TurnOn(xStart, yStart, xEnd, yEnd int) {
	xStart, yStart, xEnd, yEnd = l.clipToLights(xStart, yStart, xEnd, yEnd)
	for y := yStart; y <= yEnd; y++ {
		for x := xStart; x <= xEnd; x++ {
			l[y][x] = true
		}
	}
}

// TurnOff will turn off all lights in the square specified (top left point to bottom right point).
// Any coordinates placed outside the available lights will be ignored.
func (l ToggleLights) TurnOff(xStart, yStart, xEnd, yEnd int) {
	xStart, yStart, xEnd, yEnd = l.clipToLights(xStart, yStart, xEnd, yEnd)
	for y := yStart; y <= yEnd; y++ {
		for x := xStart; x <= xEnd; x++ {
			l[y][x] = false
		}
	}
}

// Toggle will inverse the current state of all lights in the square specified (top left point to
// bottom right point).
// Any coordinates placed outside the available lights will be ignored.
func (l ToggleLights) Toggle(xStart, yStart, xEnd, yEnd int) {
	xStart, yStart, xEnd, yEnd = l.clipToLights(xStart, yStart, xEnd, yEnd)
	for y := yStart; y <= yEnd; y++ {
		for x := xStart; x <= xEnd; x++ {
			l[y][x] = !l[y][x]
		}
	}
}

// LitLights will return the number of lights that are on.
func (l ToggleLights) LitLights() int {
	on := 0
	for y := 0; y < len(l); y++ {
		for x := 0; x < len(l[y]); x++ {
			if l[y][x] {
				on++
			}
		}
	}
	return on
}

func (l ToggleLights) clipToLights(xStart, yStart, xEnd, yEnd int) (int, int, int, int) {
	if xStart < 0 {
		xStart = 0
	}
	if yStart < 0 {
		yStart = 0
	}
	if xEnd > len(l[0]) {
		xEnd = len(l[0]) - 1
	}
	if yEnd > len(l) {
		yEnd = len(l) - 1
	}
	return xStart, yStart, xEnd, yEnd
}

// BrightnessLights provide an implementation of Lights that has brightness.
type BrightnessLights [][]int

// NewBrightnessLights creates a BrightnessLights of the size and initial state specified.
func NewBrightnessLights(width, height int, initialBrightness int) BrightnessLights {
	l := make(BrightnessLights, height)
	for i := 0; i < width; i++ {
		l[i] = make([]int, width)
		if initialBrightness > 0 {
			for j := 0; j < width; j++ {
				l[i][j] = initialBrightness
			}
		}
	}
	return l
}

// TurnOn will increase the brightness all lights in the square specified (top left point to bottom right point)
// by 1. Any coordinates placed outside the available lights will be ignored.
func (l BrightnessLights) TurnOn(xStart, yStart, xEnd, yEnd int) {
	xStart, yStart, xEnd, yEnd = l.clipToLights(xStart, yStart, xEnd, yEnd)
	for y := yStart; y <= yEnd; y++ {
		for x := xStart; x <= xEnd; x++ {
			l[y][x]++
		}
	}
}

// TurnOff will decrease the brightness all lights in the square specified (top left point to bottom right point)
// by 1 down to a minimum of 0. Any coordinates placed outside the available lights will be ignored.
func (l BrightnessLights) TurnOff(xStart, yStart, xEnd, yEnd int) {
	xStart, yStart, xEnd, yEnd = l.clipToLights(xStart, yStart, xEnd, yEnd)
	for y := yStart; y <= yEnd; y++ {
		for x := xStart; x <= xEnd; x++ {
			l[y][x]--
			if l[y][x] < 0 {
				l[y][x] = 0
			}
		}
	}
}

// Toggle will increase the brightness all lights in the square specified (top left point to bottom right point)
// by 2. Any coordinates placed outside the available lights will be ignored.
func (l BrightnessLights) Toggle(xStart, yStart, xEnd, yEnd int) {
	xStart, yStart, xEnd, yEnd = l.clipToLights(xStart, yStart, xEnd, yEnd)
	for y := yStart; y <= yEnd; y++ {
		for x := xStart; x <= xEnd; x++ {
			l[y][x] += 2
		}
	}
}

// LitLights will return the total brightness of all lights.
func (l BrightnessLights) LitLights() int {
	brightness := 0
	for y := 0; y < len(l); y++ {
		for x := 0; x < len(l[y]); x++ {
			brightness += l[y][x]
		}
	}
	return brightness
}

func (l BrightnessLights) clipToLights(xStart, yStart, xEnd, yEnd int) (int, int, int, int) {
	if xStart < 0 {
		xStart = 0
	}
	if yStart < 0 {
		yStart = 0
	}
	if xEnd > len(l[0]) {
		xEnd = len(l[0]) - 1
	}
	if yEnd > len(l) {
		yEnd = len(l) - 1
	}
	return xStart, yStart, xEnd, yEnd
}
