package main

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

func init() { days["4"] = Day4 }

/*
--- Day 4: Giant Squid ---
You're already almost 1.5km (almost a mile) below the surface of the ocean, already so deep that you can't see any sunlight. What you can see, however, is a giant squid that has attached itself to the outside of your submarine.

Maybe it wants to play bingo?

Bingo is played on a set of boards each consisting of a 5x5 grid of numbers. Numbers are chosen at random, and the chosen number is marked on all boards on which it appears. (Numbers may not appear on all boards.) If all numbers in any row or any column of a board are marked, that board wins. (Diagonals don't count.)

The submarine has a bingo subsystem to help passengers (currently, you and the giant squid) pass the time. It automatically generates a random order in which to draw numbers and a random set of boards (your puzzle input). For example:

7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1

22 13 17 11  0
 8  2 23  4 24
21  9 14 16  7
 6 10  3 18  5
 1 12 20 15 19

 3 15  0  2 22
 9 18 13 17  5
19  8  7 25 23
20 11 10 24  4
14 21 16 12  6

14 21 17 24  4
10 16 15  9 19
18  8 23 26 20
22 11 13  6  5
 2  0 12  3  7
After the first five numbers are drawn (7, 4, 9, 5, and 11), there are no winners, but the boards are marked as follows (shown here adjacent to each other to save space):

22 13 17 11  0         3 15  0  2 22        14 21 17 24  4
 8  2 23  4 24         9 18 13 17  5        10 16 15  9 19
21  9 14 16  7        19  8  7 25 23        18  8 23 26 20
 6 10  3 18  5        20 11 10 24  4        22 11 13  6  5
 1 12 20 15 19        14 21 16 12  6         2  0 12  3  7
After the next six numbers are drawn (17, 23, 2, 0, 14, and 21), there are still no winners:

22 13 17 11  0         3 15  0  2 22        14 21 17 24  4
 8  2 23  4 24         9 18 13 17  5        10 16 15  9 19
21  9 14 16  7        19  8  7 25 23        18  8 23 26 20
 6 10  3 18  5        20 11 10 24  4        22 11 13  6  5
 1 12 20 15 19        14 21 16 12  6         2  0 12  3  7
Finally, 24 is drawn:

22 13 17 11  0         3 15  0  2 22        14 21 17 24  4
 8  2 23  4 24         9 18 13 17  5        10 16 15  9 19
21  9 14 16  7        19  8  7 25 23        18  8 23 26 20
 6 10  3 18  5        20 11 10 24  4        22 11 13  6  5
 1 12 20 15 19        14 21 16 12  6         2  0 12  3  7
At this point, the third board wins because it has at least one complete row or column of marked numbers (in this case, the entire top row is marked: 14 21 17 24 4).

The score of the winning board can now be calculated. Start by finding the sum of all unmarked numbers on that board; in this case, the sum is 188. Then, multiply that sum by the number that was just called when the board won, 24, to get the final score, 188 * 24 = 4512.

To guarantee victory against the giant squid, figure out which board will win first. What will your final score be if you choose that board?

Your puzzle answer was 58412.

--- Part Two ---
On the other hand, it might be wise to try a different strategy: let the giant squid win.

You aren't sure how many bingo boards a giant squid could play at once, so rather than waste time counting its arms, the safe thing to do is to figure out which board will win last and choose that one. That way, no matter which boards it picks, it will win for sure.

In the above example, the second board is the last to win, which happens after 13 is eventually called and its middle column is completely marked. If you were to keep playing until this point, the second board would have a sum of unmarked numbers equal to 148 for a final score of 148 * 13 = 1924.

Figure out which board will win last. Once it wins, what would its final score be?

Your puzzle answer was 10030.
*/

func Day4() {
	fmt.Println("--- Day 4: Giant Squid ---")
	numbers := LoadInts(day04InputNumbers)
	boards, err := NewBingoBoards(day04InputBoards)
	if err != nil {
		fmt.Println(errors.Wrap(err, "Failed to load bingo boards"))
		return
	}
	fastest, slowest := FindWinningBingoBoards(boards, numbers)

	fmt.Println("Score of the fastest winning board is:", fastest)
	fmt.Println("Score of the slowest winning board is:", slowest)
}

// NewBingoBoards takes the bingo boards part of the puzzle input and returns []BingoBoards that
// corresponds to these.
// NewBingoBoards will return an error if the input either contains anything unexpected or has
// any partial boards in them.
func NewBingoBoards(in string) ([]BingoBoard, error) {
	var boards []BingoBoard

	var b BingoBoard
	row := 0
	for _, line := range strings.Split(in, "\n") {
		if line == "" {
			if row != 0 {
				return nil, errors.New("incomplete board, missing row(s)")
			}
			continue
		}

		values := LoadInts(line)
		if len(values) != BingoBoardDimension {
			return nil, errors.Errorf("row contained too many or too few values (line=%s)", line)
		}
		for i := 0; i < BingoBoardDimension; i++ {
			b[row][i] = BingoNumber{values[i], false}
		}

		row++
		if row == BingoBoardDimension {
			// board finished, add it to boards
			boards = append(boards, b)
			row = 0
		}
	}

	// make sure we don't have a partial board as the last one
	if row != 0 {
		return nil, errors.New("incomplete board, final board is missing row(s)")
	}

	return boards, nil
}

// FindWinningBingoBoard takes a list of one or more boards and a sequence of
// numbers and finds both the first and last board to get bingo. It calculates
// the score of the board by summing the uncalled numbers on the board and
// multiply it by the final number causing the bingo.
// FindWinningBingoBoard will return -1, -1 if no board gets a bingo.
//
// NOTE: FindWinningBingoBoards mutates the boards used as input!
func FindWinningBingoBoards(boards []BingoBoard, numbers []int) (int, int) {
	fastest := -1
	slowest := -1

	for _, n := range numbers {
		i := 0
		for i < len(boards) {
			if boards[i].Call(n) {
				// board got a bingo
				slowest = boards[i].SumUncalled() * n
				if fastest < 0 {
					fastest = slowest
				}
				// remove the board from the list of boards
				boards[i] = boards[len(boards)-1]
				boards = boards[:len(boards)-1]
			} else {
				// only increase i if the board didn't have a bingo
				i++
			}
		}
	}
	return fastest, slowest
}

// BingoBoardDimension defines the square dimensions of the BingoBoard.
const BingoBoardDimension = 5

// BingoNumber is a number for use in a BingoBoard. It contains both the number itself
// and whether this number has been called already or not.
type BingoNumber struct {
	Number int
	Called bool
}

// BingoBoard represents a single bingo board. It contains state on both what numbers
// the board consists of as well as whether this number has been drawn already or not.
type BingoBoard [BingoBoardDimension][BingoBoardDimension]BingoNumber

// Call calls a number on the board, marking it as called. If the number is on the board
// Call checks for bingo (in either the row or column only) and returns true if this is
// the case and false otherwise.
func (b *BingoBoard) Call(n int) bool {
	for y := 0; y < BingoBoardDimension; y++ {
		for x := 0; x < BingoBoardDimension; x++ {
			if b[y][x].Number == n {
				b[y][x].Called = true

				// check for bingo horizontally
				bingoX := true
				for i := 0; i < BingoBoardDimension; i++ {
					if !b[y][i].Called {
						bingoX = false
						break
					}
				}

				// check for bingo vertically
				bingoY := true
				for i := 0; i < BingoBoardDimension; i++ {
					if !b[i][x].Called {
						bingoY = false
						break
					}
				}

				return bingoX || bingoY
			}
		}
	}

	return false
}

// SumUncalled sums all uncalled numbers on the BingoBoard.
func (b *BingoBoard) SumUncalled() int {
	sum := 0
	for y := 0; y < BingoBoardDimension; y++ {
		for x := 0; x < BingoBoardDimension; x++ {
			if !b[y][x].Called {
				sum += b[y][x].Number
			}
		}
	}
	return sum
}
