package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

func init() { days["13"] = Day13 }

/*
--- Day 13: Shuttle Search ---
Your ferry can make it safely to a nearby port, but it won't get much further. When you call to book another ship, you discover that no ships embark from that port to your vacation island. You'll need to get from the port to the nearest airport.

Fortunately, a shuttle bus service is available to bring you from the sea port to the airport! Each bus has an ID number that also indicates how often the bus leaves for the airport.

Bus schedules are defined based on a timestamp that measures the number of minutes since some fixed reference point in the past. At timestamp 0, every bus simultaneously departed from the sea port. After that, each bus travels to the airport, then various other locations, and finally returns to the sea port to repeat its journey forever.

The time this loop takes a particular bus is also its ID number: the bus with ID 5 departs from the sea port at timestamps 0, 5, 10, 15, and so on. The bus with ID 11 departs at 0, 11, 22, 33, and so on. If you are there when the bus departs, you can ride that bus to the airport!

Your notes (your puzzle input) consist of two lines. The first line is your estimate of the earliest timestamp you could depart on a bus. The second line lists the bus IDs that are in service according to the shuttle company; entries that show x must be out of service, so you decide to ignore them.

To save time once you arrive, your goal is to figure out the earliest bus you can take to the airport. (There will be exactly one such bus.)

For example, suppose you have the following notes:

939
7,13,x,x,59,x,31,19
Here, the earliest timestamp you could depart is 939, and the bus IDs in service are 7, 13, 59, 31, and 19. Near timestamp 939, these bus IDs depart at the times marked D:

time   bus 7   bus 13  bus 59  bus 31  bus 19
929      .       .       .       .       .
930      .       .       .       D       .
931      D       .       .       .       D
932      .       .       .       .       .
933      .       .       .       .       .
934      .       .       .       .       .
935      .       .       .       .       .
936      .       D       .       .       .
937      .       .       .       .       .
938      D       .       .       .       .
939      .       .       .       .       .
940      .       .       .       .       .
941      .       .       .       .       .
942      .       .       .       .       .
943      .       .       .       .       .
944      .       .       D       .       .
945      D       .       .       .       .
946      .       .       .       .       .
947      .       .       .       .       .
948      .       .       .       .       .
949      .       D       .       .       .
The earliest bus you could take is bus ID 59. It doesn't depart until timestamp 944, so you would need to wait 944 - 939 = 5 minutes before it departs. Multiplying the bus ID by the number of minutes you'd need to wait gives 295.

What is the ID of the earliest bus you can take to the airport multiplied by the number of minutes you'll need to wait for that bus?

Your puzzle answer was 6568.

--- Part Two ---
The shuttle company is running a contest: one gold coin for anyone that can find the earliest timestamp such that the first bus ID departs at that time and each subsequent listed bus ID departs at that subsequent minute. (The first line in your input is no longer relevant.)

For example, suppose you have the same list of bus IDs as above:

7,13,x,x,59,x,31,19
An x in the schedule means there are no constraints on what bus IDs must depart at that time.

This means you are looking for the earliest timestamp (called t) such that:

Bus ID 7 departs at timestamp t.
Bus ID 13 departs one minute after timestamp t.
There are no requirements or restrictions on departures at two or three minutes after timestamp t.
Bus ID 59 departs four minutes after timestamp t.
There are no requirements or restrictions on departures at five minutes after timestamp t.
Bus ID 31 departs six minutes after timestamp t.
Bus ID 19 departs seven minutes after timestamp t.
The only bus departures that matter are the listed bus IDs at their specific offsets from t. Those bus IDs can depart at other times, and other bus IDs can depart at those times. For example, in the list above, because bus ID 19 must depart seven minutes after the timestamp at which bus ID 7 departs, bus ID 7 will always also be departing with bus ID 19 at seven minutes after timestamp t.

In this example, the earliest timestamp at which this occurs is 1068781:

time     bus 7   bus 13  bus 59  bus 31  bus 19
1068773    .       .       .       .       .
1068774    D       .       .       .       .
1068775    .       .       .       .       .
1068776    .       .       .       .       .
1068777    .       .       .       .       .
1068778    .       .       .       .       .
1068779    .       .       .       .       .
1068780    .       .       .       .       .
1068781    D       .       .       .       .
1068782    .       D       .       .       .
1068783    .       .       .       .       .
1068784    .       .       .       .       .
1068785    .       .       D       .       .
1068786    .       .       .       .       .
1068787    .       .       .       D       .
1068788    D       .       .       .       D
1068789    .       .       .       .       .
1068790    .       .       .       .       .
1068791    .       .       .       .       .
1068792    .       .       .       .       .
1068793    .       .       .       .       .
1068794    .       .       .       .       .
1068795    D       D       .       .       .
1068796    .       .       .       .       .
1068797    .       .       .       .       .
In the above example, bus ID 7 departs at timestamp 1068788 (seven minutes after t). This is fine; the only requirement on that minute is that bus ID 19 departs then, and it does.

Here are some other examples:

The earliest timestamp that matches the list 17,x,13,19 is 3417.
67,7,59,61 first occurs at timestamp 754018.
67,x,7,59,61 first occurs at timestamp 779210.
67,7,x,59,61 first occurs at timestamp 1261476.
1789,37,47,1889 first occurs at timestamp 1202161486.
However, with so many bus IDs in your list, surely the actual earliest timestamp will be larger than 100000000000000!

What is the earliest timestamp such that all of the listed bus IDs depart at offsets matching their positions in the list?

Your puzzle answer was 554865447501099.
*/

func Day13() {
	fmt.Println("--- Day 13: Shuttle Search ---")
	ts, busses, err := LoadTimestampAndBusses(day13Input)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("  The earliest bus:", EarliestBus(ts, busses))
	fmt.Println("  The earliest timestamp where all busses leaves within a minute of each other:", EarliestTimestampWith1MinuteDelay(busses))
}

// LoadTimestampAndBusses takes the raw input containing a timestamp and a list of busses and
// returns an integer representation of these. Missing busses (with the x value) are represented
// as -1. LoadTimestampAndBusses will return an error if it cannot be loaded.
func LoadTimestampAndBusses(in string) (int, []int, error) {
	nlPos := strings.Index(in, "\n")
	ts, err := strconv.Atoi(in[:nlPos])
	if err != nil {
		return 0, nil, errors.Wrapf(err, "bad timestamp (ts=%s)", in[:nlPos])
	}

	var busses []int
	for _, bus := range strings.Split(in[nlPos+1:], ",") {
		if bus == "x" {
			busses = append(busses, -1)
			continue
		}

		i, err := strconv.Atoi(bus)
		if err != nil {
			return 0, nil, errors.Wrapf(err, "invalid bus (bus=%s)", bus)
		}
		busses = append(busses, i)
	}

	return ts, busses, nil
}

// EarliestBus takes a timestamp and a list of busses and returns the earliest bus that will leave
// multiplied by the amount of minutes the wait is long for it.
func EarliestBus(ts int, busses []int) int {
	// the earliest bus is the one with the departure closest to ts
	shortestWait := int(^uint(0) >> 1) // max value of an int
	earliestBus := -1
	for _, bus := range busses {
		if bus == -1 {
			continue
		}
		wait := bus - ts%bus
		if wait < shortestWait {
			shortestWait = wait
			earliestBus = bus
		}
	}

	return earliestBus * shortestWait
}

// EarliestTimestampWith1MinuteDelay takes a list of busses and returns the timestamp of the first
// point in time where all busses arrive/depart exatly at the timestamp offset by their index in the
// busses input slice.
func EarliestTimestampWith1MinuteDelay(busses []int) int {
	if len(busses) < 1 {
		return 0
	}

	// We need to look for the "perfect order" (bus 1 departs at ts, bus 2 departs at ts+1, etc).
	//
	// If only looking at two busses, you can reason that these two will be in the prefect order
	// with a specific exact interval between. (eg. the busses 2,3 would be in the perfect order
	// at ts=2, ts=8, ts=14, etc).
	//
	// Given a 3rd bus would have to align perfectly with the first two, it would have to be one
	// of the timestamps from above. And it would again repeat itself with a specific and exact
	// interval.
	//
	// The idea here is that we levrage this knowledge to take increasing larger steps in order
	// to find the timestamp where all busses are in the perfect order.

	step := busses[0] // step guarantees all previous busses to be in the perfect order
	ts := 0

	for n := 1; n < len(busses); n++ {
		if busses[n] == -1 {
			continue
		}

		// we need to determine the steps needed between the first and second occurance
		// of the perfect order
		firstOccurance := -1

		for {
			ts += step
			if (ts+n)%busses[n] == 0 {
				if firstOccurance < 0 {
					firstOccurance = ts

					// we might be done now if this is the last bus - or all other remaining busses are aligned already
					done := true
					for m := n + 1; m < len(busses); m++ {
						if busses[m] == -1 {
							continue
						}
						if (ts+m)%busses[m] != 0 {
							// nope, not quite done after all
							done = false
							break
						}
					}

					if done {
						return ts
					}
				} else {
					step = ts - firstOccurance
					break
				}
			}
		}

	}

	panic("the bus loop exited without returning - this should never happen")
}
