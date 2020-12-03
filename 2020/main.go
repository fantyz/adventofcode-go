package main

import (
	"fmt"
	"os"
	"strconv"
)

var days = map[string]func(){}

func main() {
	fmt.Println()
	fmt.Println("Advent of Code 2020")
	fmt.Println()

	if len(os.Args) == 2 {
		day, found := days[os.Args[1]]
		if !found {
			fmt.Println("ERROR: No day found for " + os.Args[1])
			return
		}
		day()
	} else {
		// run all days
		for i := 1; i <= 25; i++ {
			if day, found := days[strconv.Itoa(i)]; found {
				day()
				fmt.Println()
			}
		}
	}
}
