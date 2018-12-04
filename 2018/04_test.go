package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTimelien(t *testing.T) {
	testCases := []struct {
		In   string
		Out  int
		Out2 int
	}{
		{
			In: `
[1518-11-01 00:00] Guard #10 begins shift
[1518-11-01 00:05] falls asleep
[1518-11-01 00:25] wakes up
[1518-11-01 00:30] falls asleep
[1518-11-01 00:55] wakes up
[1518-11-01 23:58] Guard #99 begins shift
[1518-11-02 00:40] falls asleep
[1518-11-02 00:50] wakes up
[1518-11-03 00:05] Guard #10 begins shift
[1518-11-03 00:24] falls asleep
[1518-11-03 00:29] wakes up
[1518-11-04 00:02] Guard #99 begins shift
[1518-11-04 00:36] falls asleep
[1518-11-04 00:46] wakes up
[1518-11-05 00:03] Guard #99 begins shift
[1518-11-05 00:45] falls asleep
[1518-11-05 00:55] wakes up`,
			Out:  240,
			Out2: 4455,
		},
	}

	for i, c := range testCases {
		tl := NewTimeline(c.In)
		id1 := GuardMostAsleep(tl)
		min1, _ := MinuteMostAsleep(id1, tl)
		id2, min2 := AllMinuteMostAsleep(tl)

		assert.Equal(t, c.Out, id1*min1, "(case=%d)", i)
		assert.Equal(t, c.Out2, id2*min2, "(case=%d)", i)
	}
}
