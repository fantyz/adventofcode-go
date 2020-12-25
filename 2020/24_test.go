package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimulateDays(t *testing.T) {
	testInput := `sesenwnenenewseeswwswswwnenewsewsw
neeenesenwnwwswnenewnwwsewnenwseswesw
seswneswswsenwwnwse
nwnwneseeswswnenewneswwnewseswneseene
swweswneswnenwsewnwneneseenw
eesenwseswswnenwswnwnwsewwnwsene
sewnenenenesenwsewnenwwwse
wenwwweseeeweswwwnwwe
wsweesenenewnwwnwsenewsenwwsesesenwne
neeswseenwwswnwswswnw
nenwswwsewswnenenewsenwsenwnesesenew
enewnwewneswsewnwswenweswnenwsenwsw
sweneswneswneneenwnewenewwneswswnese
swwesenesewenwneswnwwneseswwne
enesenwswwswneneswsenwnewswseenwsese
wnwnesenesenenwwnenwsewesewsesesew
nenewswnwewswnenesenwnesewesw
eneswnwswnwsenenwnwnwwseeswneewsenese
neswnwewnwnwseenwseesewsenwsweewe
wseweeenwnesenwwwswnew`

	testCases := []struct {
		Days       int
		BlackTiles int
	}{
		{0, 10},
		{1, 15},
		{2, 12},
		{20, 132},
		{50, 566},
		{100, 2208},
	}

	for _, c := range testCases {
		lobby, err := NewLobby(testInput)
		if assert.NoError(t, err) {
			lobby.SimulateDays(c.Days)
			assert.Equal(t, c.BlackTiles, lobby.BlackTiles())
		}
	}
}

func TestDay24P(t *testing.T) {
	lobby, err := NewLobby(day24Input)
	if assert.NoError(t, err) {
		assert.Equal(t, 438, lobby.BlackTiles())
		lobby.SimulateDays(100)
		assert.Equal(t, 4038, lobby.BlackTiles())
	}
}
