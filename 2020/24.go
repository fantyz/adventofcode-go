package main

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

func init() { days["24"] = Day24 }

/*
--- Day 24: Lobby Layout ---
Your raft makes it to the tropical island; it turns out that the small crab was an excellent navigator. You make your way to the resort.

As you enter the lobby, you discover a small problem: the floor is being renovated. You can't even reach the check-in desk until they've finished installing the new tile floor.

The tiles are all hexagonal; they need to be arranged in a hex grid with a very specific color pattern. Not in the mood to wait, you offer to help figure out the pattern.

The tiles are all white on one side and black on the other. They start with the white side facing up. The lobby is large enough to fit whatever pattern might need to appear there.

A member of the renovation crew gives you a list of the tiles that need to be flipped over (your puzzle input). Each line in the list identifies a single tile that needs to be flipped by giving a series of steps starting from a reference tile in the very center of the room. (Every line starts from the same reference tile.)

Because the tiles are hexagonal, every tile has six neighbors: east, southeast, southwest, west, northwest, and northeast. These directions are given in your list, respectively, as e, se, sw, w, nw, and ne. A tile is identified by a series of these directions with no delimiters; for example, esenee identifies the tile you land on if you start at the reference tile and then move one tile east, one tile southeast, one tile northeast, and one tile east.

Each time a tile is identified, it flips from white to black or from black to white. Tiles might be flipped more than once. For example, a line like esew flips a tile immediately adjacent to the reference tile, and a line like nwwswee flips the reference tile itself.

Here is a larger example:

sesenwnenenewseeswwswswwnenewsewsw
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
wseweeenwnesenwwwswnew
In the above example, 10 tiles are flipped once (to black), and 5 more are flipped twice (to black, then back to white). After all of these instructions have been followed, a total of 10 tiles are black.

Go through the renovation crew's list and determine which tiles they need to flip. After all of the instructions have been followed, how many tiles are left with the black side up?

Your puzzle answer was 438.

--- Part Two ---
The tile floor in the lobby is meant to be a living art exhibit. Every day, the tiles are all flipped according to the following rules:

Any black tile with zero or more than 2 black tiles immediately adjacent to it is flipped to white.
Any white tile with exactly 2 black tiles immediately adjacent to it is flipped to black.
Here, tiles immediately adjacent means the six tiles directly touching the tile in question.

The rules are applied simultaneously to every tile; put another way, it is first determined which tiles need to be flipped, then they are all flipped at the same time.

In the above example, the number of black tiles that are facing up after the given number of days has passed is as follows:

Day 1: 15
Day 2: 12
Day 3: 25
Day 4: 14
Day 5: 23
Day 6: 28
Day 7: 41
Day 8: 37
Day 9: 49
Day 10: 37

Day 20: 132
Day 30: 259
Day 40: 406
Day 50: 566
Day 60: 788
Day 70: 1106
Day 80: 1373
Day 90: 1844
Day 100: 2208
After executing this process a total of 100 times, there would be 2208 black tiles facing up.

How many tiles will be black after 100 days?

Your puzzle answer was 4038.
*/

func Day24() {
	fmt.Println("--- Day 24: Lobby Layout ---")
	lobby, err := NewLobby(day24Input)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Black tiles after day 0:", lobby.BlackTiles())
	lobby.SimulateDays(100)
	fmt.Println("Black tiles after day 100:", lobby.BlackTiles())
}

// A hex grid is easily represented with regular x,y coordinates.
//
//           0, 1   1, 1
//        -1,0   0,0   1,0
//          -1,-1   0,-1
//

// NewLobby takes a series of move instructions as input that indicates the location of tiles
// that must be flipped and returns the corresponding lobby after initializing it.
// An error is returned if the input is invalid.
func NewLobby(in string) (*Lobby, error) {
	l := Lobby{}
	for _, moves := range strings.Split(in, "\n") {
		pos := Coord{0, 0}
		idx := 0
		for idx < len(moves) {
			switch {
			case moves[idx] == 'e':
				pos.X++
				idx++
			case moves[idx] == 's' && idx+1 < len(moves) && moves[idx+1] == 'e':
				pos.Y--
				idx += 2
			case moves[idx] == 's' && idx+1 < len(moves) && moves[idx+1] == 'w':
				pos.X--
				pos.Y--
				idx += 2
			case moves[idx] == 'w':
				pos.X--
				idx++
			case moves[idx] == 'n' && idx+1 < len(moves) && moves[idx+1] == 'e':
				pos.X++
				pos.Y++
				idx += 2
			case moves[idx] == 'n' && idx+1 < len(moves) && moves[idx+1] == 'w':
				pos.Y++
				idx += 2
			default:
				return nil, errors.Errorf("unknown direction (idx=%d, moves=%s)", idx, moves)
			}
		}
		l.FlipTile(pos)
	}
	return &l, nil
}

// Lobby represents the tiles in the lobby.
type Lobby map[Coord]struct{}

// FlipTile takes coordinate and flips the tile located at it.
func (l Lobby) FlipTile(pos Coord) {
	if _, isBlack := l[pos]; isBlack {
		// black -> white
		delete(l, pos)
	} else {
		// white -> black
		l[pos] = struct{}{}
	}
}

// BlackTiles returns the number of black tiles in the lobby.
func (l Lobby) BlackTiles() int {
	return len(l)
}

// SimulateDays runs the daily flip the number of days specified.
func (l *Lobby) SimulateDays(days int) {
	for i := 0; i < days; i++ {
		l.DailyFlip()
	}
}

// DailyFlip flipes the tiles using the daily tile flipping rules.
func (l *Lobby) DailyFlip() {
	newL := Lobby{}

	// The daily flipping rules only consider tiles that are already black or is a neighbor
	// to an already black tile. We gather all the coordinates first to avoid evaluating tiles
	// more than once.
	coordsToEvaluate := map[Coord]struct{}{}
	for pos := range *l {
		coordsToEvaluate[pos] = struct{}{}
		for _, neighbor := range HexNeighbors(pos) {
			coordsToEvaluate[neighbor] = struct{}{}
		}
	}

	for pos := range coordsToEvaluate {
		blacks := l.BlackNeighbors(pos)
		if _, isBlack := (*l)[pos]; isBlack {
			// black tile
			if blacks >= 1 && blacks <= 2 {
				// keep black
				newL[pos] = struct{}{}
			}
		} else {
			// white tile
			if blacks == 2 {
				// change to black
				newL[pos] = struct{}{}
			}
		}
	}

	*l = newL
}

// BlackNeighbors returns the number of black neighbors to the specified tile.
func (l *Lobby) BlackNeighbors(pos Coord) int {
	sum := 0
	for _, neighbor := range HexNeighbors(pos) {
		if _, isBlack := (*l)[neighbor]; isBlack {
			sum++
		}
	}
	return sum
}

type Coord struct {
	X, Y int
}

// HexNeighbors returns the coordinates of the directly adjacent tiles in a hex grid.
func HexNeighbors(pos Coord) []Coord {
	return []Coord{
		{pos.X + 1, pos.Y},     // e
		{pos.X, pos.Y - 1},     // se
		{pos.X - 1, pos.Y - 1}, // sw
		{pos.X - 1, pos.Y},     // w
		{pos.X + 1, pos.Y + 1}, // ne
		{pos.X, pos.Y + 1},     // nw
	}
}
