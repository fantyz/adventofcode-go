package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*


 */

func main() {
	fmt.Println("Day X")
}

func a(in string) int {
	for _, l := range strings.Split(in, "\n") {
		i, err := strconv.Atoi(l)
		if err != nil {
			panic(err)
		}

		fmt.Println(i)
		return i
	}
}

const puzzleInput = ``
