package main

import (
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

func LoadInts(in string) ([]int, error) {
	var v []int
	for _, s := range strings.Split(in, "\n") {
		i, err := strconv.Atoi(s)
		if err != nil {
			return nil, errors.Wrapf(err, "unable to convert %s to int", s)
		}
		v = append(v, i)
	}
	return v, nil
}
