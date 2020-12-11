package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChangeSeats(t *testing.T) {
	testLayout := `#.##.L#.##
#L###LL.L#
L.#.#..#..
#L##.##.L#
#.##.LL.LL
#.###L#.##
..#.#.....
#L######L#
#.LL###L.L
#.#L###.##`

	testCases := []struct {
		Dist    int
		MinAdj  int
		Changes int
		Total   int
	}{
		{1, 4, 21, 30},
	}

	for _, c := range testCases {
		l := NewSeatLayout(testLayout)

		changes, total := l.ChangeSeats(c.Dist, c.MinAdj)
		assert.Equal(t, c.Changes, changes)
		assert.Equal(t, c.Total, total)
	}
}

func TestStabilizeSeats(t *testing.T) {
	testLayout := `L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL`

	testCases := []struct {
		Dist   int
		MinAdj int
		Total  int
	}{
		{1, 4, 37},
	}

	for _, c := range testCases {
		l := NewSeatLayout(testLayout)
		assert.Equal(t, c.Total, l.StabilizeSeats(c.Dist, c.MinAdj))
	}
}

func TestDay11Pt1(t *testing.T) {
	assert.Equal(t, 2406, NewSeatLayout(day11Input).StabilizeSeats(1, 4))
}

func TestDay11Pt2(t *testing.T) {
	assert.Equal(t, 2149, NewSeatLayout(day11Input).StabilizeSeats(-1, 5))
}

func BenchmarkStabilizeSeats(b *testing.B) {
	for n := 0; n < b.N; n++ {
		b.StopTimer()
		l := NewSeatLayout(day11Input)
		b.StartTimer()

		_ = l.StabilizeSeats(-1, 5)
	}
}
