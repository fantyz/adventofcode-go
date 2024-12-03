package main

import (
	"fmt"
	"regexp"
	"strings"
)

func init() { days["03"] = Day03 }

/*
 */

func Day03() {
	fmt.Println("Day 3: Mull It Over")
	fmt.Println("Part 1: Sum of multiplied numbers found in corrupted memory:", ScanCorruptedMemoryForMuls([]string{day03Input}))
	fmt.Println("Part 2: Sum of multiplied numbers found in corrupted memory considering do() and don't() regions:", ScanCorruptedMemoryForMuls(ScanCorruptedMemoryForDoRegions(day03Input)))
}

// ScanCurruptedMemoryForDoRegions takes a strong containing corrupted memory and scans it
// for do() statements. It returns all memory segments that follows a do() statement until a
// don't() statement is encountered.
// ScanCurruptedMemoryForDoRegions assumes the memory starts with a do() statement implicitly.
func ScanCorruptedMemoryForDoRegions(mem string) []string {
	// break the memory in to segments that are preceeded with a do() statement
	segments := strings.Split(mem, "do()")

	// for each segment remove any memory that follows a don't() statement
	for i := range segments {
		idx := strings.Index(segments[i], "don't()")
		if idx > 0 {
			// remove the don't() segment
			segments[i] = segments[i][:idx]
		}
	}
	return segments
}

// ScanCorruptedMemoryForMuls takes a string continaing corrupted memory and scans it for
// mul(<int>,<int>) expressions. It performs the multiplication of the two integers and adds
// all the multiplied numbers together and returns them.
func ScanCorruptedMemoryForMuls(mem []string) int {
	r := regexp.MustCompile(`mul\(([0-9]+),([0-9]+)\)`)

	sum := 0
	for _, segment := range mem {
		for _, m := range r.FindAllStringSubmatch(segment, -1) {
			sum += ToIntOrPanic(m[1]) * ToIntOrPanic(m[2])
		}
	}
	return sum
}
