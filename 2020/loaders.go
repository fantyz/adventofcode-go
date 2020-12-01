package main

import (
	"fmt"
	"strconv"
	"strings"
)

func LoadInts(in string) []int {
	var v []int
	for _, s := range strings.Split(in, "\n") {
		i, err := strconv.Atoi(s)
		if err != nil {
			panic(fmt.Sprintf("Bad puzzle input - unable to convert %s to int", s))
		}
		v = append(v, i)
	}
	return v
}
