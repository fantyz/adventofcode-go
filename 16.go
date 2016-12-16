package main

import (
	"fmt"
)

/*

 */

func main() {
	fmt.Println("Advent of Code 2016 - Day 16")
	fmt.Println("The checksum for size=272 is:", string(checksum(modifiedDragonCurve(272, []byte(puzzleInput)))))
	fmt.Println("The checksum for size=35651584 is:", string(checksum(modifiedDragonCurve(35651584, []byte(puzzleInput)))))
}

func checksum(data []byte) []byte {
	sum := []byte{}
	for i := 0; i < len(data)-1; i += 2 {
		b := byte('0')
		if data[i] == data[i+1] {
			b = '1'
		}
		sum = append(sum, b)
	}

	if len(sum)%2 == 0 {
		return checksum(sum)
	}
	return sum
}

func modifiedDragonCurve(size int, data []byte) []byte {
	if len(data) >= size {
		return data
	}
	data = append(data, '0')
	for i := len(data) - 2; i >= 0; i-- {
		// stop of we have the desired amount of bytes
		if len(data) == size {
			return data
		}
		if data[i] == '0' {
			data = append(data, '1')
		} else {
			data = append(data, '0')
		}
	}
	return modifiedDragonCurve(size, data)
}

const puzzleInput = `00101000101111010`
