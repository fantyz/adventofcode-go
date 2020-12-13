package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadTimestmapAndBusses(t *testing.T) {
	testInput := `939
7,13,x,x,59,x,31,19`

	ts, busses, err := LoadTimestampAndBusses(testInput)
	if assert.NoError(t, err) {
		assert.Equal(t, 939, ts)
		assert.Equal(t, []int{7, 13, -1, -1, 59, -1, 31, 19}, busses)
	}
}

func TestEarliestBus(t *testing.T) {
	testInput := `939
7,13,x,x,59,x,31,19`

	ts, busses, err := LoadTimestampAndBusses(testInput)
	if assert.NoError(t, err) {
		assert.Equal(t, 295, EarliestBus(ts, busses))
	}
}

func TestEarliestTimestampWith1MinuteDelay(t *testing.T) {
	testCases := map[string]struct {
		Busses    []int
		Timestamp int
	}{
		"17,x,13,19":      {[]int{17, -1, 13, 19}, 3417},
		"67,7,59,61":      {[]int{67, 7, 59, 61}, 754018},
		"67,x,7,59,61":    {[]int{67, -1, 7, 59, 61}, 779210},
		"67,7,x,59,61":    {[]int{67, 7, -1, 59, 61}, 1261476},
		"1789,37,47,1889": {[]int{1789, 37, 47, 1889}, 1202161486},
	}

	for n, c := range testCases {
		assert.Equal(t, c.Timestamp, EarliestTimestampWith1MinuteDelay(c.Busses), n)
	}
}

func TestDay13Pt1(t *testing.T) {
	ts, busses, err := LoadTimestampAndBusses(day13Input)
	if assert.NoError(t, err) {
		assert.Equal(t, 6568, EarliestBus(ts, busses))
	}
}

func TestDay13Pt2(t *testing.T) {
	_, busses, err := LoadTimestampAndBusses(day13Input)
	if assert.NoError(t, err) {
		assert.Equal(t, 554865447501099, EarliestTimestampWith1MinuteDelay(busses))
	}
}
