package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

func init() { days["2"] = Day2 }

/*
--- Day 2: I Was Told There Would Be No Math ---
The elves are running low on wrapping paper, and so they need to submit an order for more. They have a list of the dimensions (length l, width w, and height h) of each present, and only want to order exactly as much as they need.

Fortunately, every present is a box (a perfect right rectangular prism), which makes calculating the required wrapping paper for each gift a little easier: find the surface area of the box, which is 2*l*w + 2*w*h + 2*h*l. The elves also need a little extra paper for each present: the area of the smallest side.

For example:

A present with dimensions 2x3x4 requires 2*6 + 2*12 + 2*8 = 52 square feet of wrapping paper plus 6 square feet of slack, for a total of 58 square feet.
A present with dimensions 1x1x10 requires 2*1 + 2*10 + 2*10 = 42 square feet of wrapping paper plus 1 square foot of slack, for a total of 43 square feet.
All numbers in the elves' list are in feet. How many total square feet of wrapping paper should they order?

Your puzzle answer was 1598415.

--- Part Two ---
The elves are also running low on ribbon. Ribbon is all the same width, so they only have to worry about the length they need to order, which they would again like to be exact.

The ribbon required to wrap a present is the shortest distance around its sides, or the smallest perimeter of any one face. Each present also requires a bow made out of ribbon as well; the feet of ribbon required for the perfect bow is equal to the cubic feet of volume of the present. Don't ask how they tie the bow, though; they'll never tell.

For example:

A present with dimensions 2x3x4 requires 2+2+3+3 = 10 feet of ribbon to wrap the present plus 2*3*4 = 24 feet of ribbon for the bow, for a total of 34 feet.
A present with dimensions 1x1x10 requires 1+1+1+1 = 4 feet of ribbon to wrap the present plus 1*1*10 = 10 feet of ribbon for the bow, for a total of 14 feet.
How many total feet of ribbon should they order?

Your puzzle answer was 3812909.
*/

func Day2() {
	fmt.Println("Day 2: I Was Told There Would Be No Math")
	paperNeeded, ribbonNeeded, err := WrapBoxes(day2Input)
	if err != nil {
		panic(err)
	}
	fmt.Println("  Wrapping paper needed:", paperNeeded, "square feet")
	fmt.Println("          Ribbon needed:", ribbonNeeded, "feet")
}

// WrapBoxes takes a string containing the specifications of a list of boxes and returns the square feets of
// wrapping paper needed to wrap them.
// The input format expects each line in the input to contain the specification of a box in the format
// `<length>x<width>x<height>` all in square feet.
func WrapBoxes(boxes string) (int, int, error) {
	paperNeeded := 0
	ribbonNeeded := 0

	for n, box := range strings.Split(boxes, "\n") {
		dim := strings.Split(box, "x")
		if len(dim) != 3 {
			return 0, 0, fmt.Errorf("invalid dimensions (n=%d, box=%s)", n, box)
		}
		l, err := strconv.Atoi(dim[0])
		if err != nil {
			return 0, 0, errors.Wrapf(err, "invalid length of box (n=%d, length=%s)", n, dim[0])
		}
		w, err := strconv.Atoi(dim[1])
		if err != nil {
			return 0, 0, errors.Wrapf(err, "invalid width of box (n=%d, width=%s)", n, dim[1])
		}
		h, err := strconv.Atoi(dim[2])
		if err != nil {
			return 0, 0, errors.Wrapf(err, "invalid height of box (n=%d, height=%s)", n, dim[2])
		}

		p, r := WrapBox(l, w, h)
		paperNeeded += p
		ribbonNeeded += r
	}

	return paperNeeded, ribbonNeeded, nil
}

// WrapBox takes the dimensions of a box in feet and returns the square feet of wrapping paper and feet of
// ribon needed to wrap it.
func WrapBox(length, width, height int) (int, int) {
	top := length * width
	side := width * height
	end := height * length

	slack := 0
	ribbon := length * width * height
	switch {
	case length >= width && length >= height:
		// avoid length for both slack and shortest distance
		slack = side
		ribbon += 2*width + 2*height
	case width >= length && width >= height:
		// avoid width for both slack and shortest distance
		slack = end
		ribbon += 2*length + 2*height
	default:
		// avoid height for both slack and shortest distance
		slack = top
		ribbon += 2*length + 2*width
	}

	return 2*top + 2*side + 2*end + slack, ribbon
}
