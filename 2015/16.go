package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

func init() { days["16"] = Day16 }

/*
--- Day 16: Aunt Sue ---
Your Aunt Sue has given you a wonderful gift, and you'd like to send her a thank you card. However, there's a small problem: she signed it "From, Aunt Sue".

You have 500 Aunts named "Sue".

So, to avoid sending the card to the wrong person, you need to figure out which Aunt Sue (which you conveniently number 1 to 500, for sanity) gave you the gift. You open the present and, as luck would have it, good ol' Aunt Sue got you a My First Crime Scene Analysis Machine! Just what you wanted. Or needed, as the case may be.

The My First Crime Scene Analysis Machine (MFCSAM for short) can detect a few specific compounds in a given sample, as well as how many distinct kinds of those compounds there are. According to the instructions, these are what the MFCSAM can detect:

children, by human DNA age analysis.
cats. It doesn't differentiate individual breeds.
Several seemingly random breeds of dog: samoyeds, pomeranians, akitas, and vizslas.
goldfish. No other kinds of fish.
trees, all in one group.
cars, presumably by exhaust or gasoline or something.
perfumes, which is handy, since many of your Aunts Sue wear a few kinds.
In fact, many of your Aunts Sue have many of these. You put the wrapping from the gift into the MFCSAM. It beeps inquisitively at you a few times and then prints out a message on ticker tape:

children: 3
cats: 7
samoyeds: 2
pomeranians: 3
akitas: 0
vizslas: 0
goldfish: 5
trees: 3
cars: 2
perfumes: 1
You make a list of the things you can remember about each Aunt Sue. Things missing from your list aren't zero - you simply don't remember the value.

What is the number of the Sue that got you the gift?

Your puzzle answer was 103.

--- Part Two ---
As you're about to send the thank you note, something in the MFCSAM's instructions catches your eye. Apparently, it has an outdated retroencabulator, and so the output from the machine isn't exact values - some of them indicate ranges.

In particular, the cats and trees readings indicates that there are greater than that many (due to the unpredictable nuclear decay of cat dander and tree pollen), while the pomeranians and goldfish readings indicate that there are fewer than that many (due to the modial interaction of magnetoreluctance).

What is the number of the real Aunt Sue?

Your puzzle answer was 405.
*/

func Day16() {
	fmt.Println("--- Day 16: Aunt Sue ---")
	giftDetails := map[string]int{
		"children":    3,
		"cats":        7,
		"samoyeds":    2,
		"pomeranians": 3,
		"akitas":      0,
		"vizslas":     0,
		"goldfish":    5,
		"trees":       3,
		"cars":        2,
		"perfumes":    1,
	}
	sueDetails, err := NewSueDetails(day16Input)
	if err != nil {
		fmt.Println(errors.Wrap(err, "failed to load sue details"))
		return
	}
	fmt.Println("The gift is from Sue number:", sueDetails.FindFirstMatch(giftDetails, false))
	fmt.Println("The gift is *really* from Sue number:", sueDetails.FindFirstMatch(giftDetails, true))
}

// SueDetails is a slice of map[string]int each containing the details of a specific Sue.
type SueDetails []map[string]int

// NewSueDetails will load a puzzle input containing the remembered details about the
// different Sues and return a SueDetails. An error will be returned if the input is invalid.
func NewSueDetails(input string) (SueDetails, error) {
	var sds SueDetails
	detailsExp := regexp.MustCompile(`(children|cats|samoyeds|pomeranians|akitas|vizslas|goldfish|trees|cars|perfumes): ([0-9]+)`)

	for _, line := range strings.Split(input, "\n") {
		d := map[string]int{}

		n := strings.Index(line, ":")
		if n < 4 || line[:4] != "Sue " {
			return nil, errors.Errorf("unexpected line in input (line=%s)", line)
		}
		num, err := strconv.Atoi(line[4:n])
		if err != nil {
			// should not happen, would be due to bad code
			panic(errors.Wrapf(err, "sue number could not be parsed as a number (num=%s, line=%s)", line[3:n], line))
		}
		d["name"] = num

		m := detailsExp.FindAllStringSubmatch(line, -1)
		for _, detail := range m {
			if len(detail) != 3 {
				// should not happen, would be due to bad detailsExp
				panic(fmt.Sprintf("unexpected number of matches for detail (line=%s)", line))
			}
			num, err := strconv.Atoi(detail[2])
			if err != nil {
				// should not happen, would be due to bad detailsExp
				panic(errors.Wrapf(err, "detail count could not be parsed as a number (detail=%s, num=%s,line=%s)", detail[1], detail[2], line))
			}
			d[detail[1]] = num
		}
		sds = append(sds, d)
	}
	return sds, nil
}

// FindFirstMatch will look through the SueDetails and return the name detail of the first match
// with the given input details. If no match is found the returned value will be -1.
func (sueDetails SueDetails) FindFirstMatch(details map[string]int, useRanges bool) int {
	for _, d := range sueDetails {
		match := true
		for name, desiredCount := range details {
			if count, found := d[name]; found {
				match = desiredCount == count
				if useRanges {
					switch name {
					case "cats", "trees":
						match = desiredCount < count
					case "pomeranians", "goldfish":
						match = desiredCount > count
					}
				}
				if !match {
					break
				}
			}
		}
		if match {
			return d["name"]
		}
	}
	return -1
}
