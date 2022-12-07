package main

import (
	"fmt"
	"strconv"

	"github.com/pkg/errors"
)

// LoadInts takes an input string and reads out any integers found in it in the order they
// appear ignoring all other text in string.
func LoadInts(in string) []int {
	in += "\n" // make sure in has a non-number at the end of it

	var v []int

	start := -1
	for idx, r := range in {
		switch r {
		case '-', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			if start >= 0 {
				continue
			}
			start = idx
		default:
			if start >= 0 {
				n, err := strconv.Atoi(in[start:idx])
				if err != nil {
					// should never happen
					panic(fmt.Sprintf("unable to convert number to int (n=%s)", in[start:idx]))
				}
				v = append(v, n)
				start = -1
			}
		}
	}
	return v
}

// ToIntOrPanic takes an input string and converts it to an integer. If unable it will
// panic. You should only use this if you are guaranteed that the input indeed is an integer
// like is the case where you matched the string with a regexp already.
func ToIntOrPanic(in string) int {
	n, err := strconv.Atoi(in)
	if err != nil {
		panic(errors.Wrapf(err, "Unable to convert value to int (in=%s)", in))
	}
	return n
}
