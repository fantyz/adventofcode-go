package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

var days = map[string]func(){}

func main() {
	fmt.Println()
	fmt.Println("Advent of Code 2024")
	fmt.Println()

	timedDay := func(day func()) {
		ts := time.Now()
		day()
		fmt.Printf("[Time taken: %v]\n", time.Since(ts))
	}

	if len(os.Args) == 2 {
		num := os.Args[1]
		if len(num) == 1 {
			num = "0"+num
		}
		day, found := days[num]
		if !found {
			fmt.Println("ERROR: No day found for " + os.Args[1])
			return
		}
		timedDay(day)
	} else {
		// run all days
		for i := 1; i <= 25; i++ {
			num := strconv.Itoa(i)
			if i < 10 {
				num = "0"+num
			}

			if day, found := days[num]; found {
				timedDay(day)
				fmt.Println()
			}
		}
	}
}
