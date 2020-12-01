package main

import (
	"fmt"
	"os"
)

var puzzles = map[string]func(){}

func main() {
	fmt.Println("Advent of Code 2020")
	if len(os.Args) != 2 {
		fmt.Println("ERROR: Specify the day to run... (eg \"1\")")
		return
	}

	puzzle, found := puzzles[os.Args[1]]
	if !found {
		fmt.Println("ERROR: No puzzle found for " + os.Args[1])
		return
	}

	puzzle()
}
