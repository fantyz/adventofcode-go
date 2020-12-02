package main

import (
	"fmt"
	"os"
)

var days = map[string]func(){}

func main() {
	fmt.Println("Advent of Code 2015")
	if len(os.Args) != 2 {
		fmt.Println("ERROR: Specify the day to run... (eg \"1\")")
		return
	}

	day, found := days[os.Args[1]]
	if !found {
		fmt.Println("ERROR: No day found for " + os.Args[1])
		return
	}

	day()
}
