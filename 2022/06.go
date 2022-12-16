package main

import (
	"fmt"
)

func init() { days["6"] = Day6 }

/*
 */

func Day6() {
	fmt.Println("--- Day 6: Tuning Trouble ---")
	fmt.Println("First start-of-packet-marker is located at:", FindStartOfPacketMarker(day06Input))
}

// FindStartOfPacket returns the index of the start-of-packet-marker.
// If no start-of-packet-marker exist -1 is returned.
func FindStartOfPacketMarker(in string) int {
	lastDuplicate := 0
	for i := 0; i < len(in); i++ {
		// compare in[i] with the previous 3 characters looking for duplicates
		for j := i - 1; j >= i-3 && j >= 0; j-- {
			if in[i] == in[j] {
				// duplicate found, earliest marker can start at j+1
				lastDuplicate = j
			}
		}

		// if lastDuplicate is more than 4 characters away we have our marker
		if i-lastDuplicate > 4 {
			// marker found!
			return i
		}
	}
	return -1
}
