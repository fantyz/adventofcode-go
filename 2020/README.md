# Advent of Code 2020

The puzzles have been solved with focus on writing easy to read, clean and simple solutions that runs the individual puzzles in less than a second when possible. That means I actively favor a simple brute-force solution over a faster and more complex solution if the simple solution is fast enough to solve the puzzle in less than a second.

## Favorite Puzzles

Some puzzles are just more fun to solve than others. These are the most memorable ones:

* [Day 8: Handheld Halting](08.go): Debugging a BoodCode program.


## File Structure

All puzzles are part of the same application. This is primarily done to avoid having a ton of modules and ease re-use of code between days where possible.

The `main.go` contains the basic logic needed to simply run the puzzle of one day.

The puzzles are organized by day and consist of:

* `<day>.go` containing a description of the puzzle of the day itself along with the code needed to solve it.
* `<day>_input.go` containing the puzzle input unique to me needed to solve the puzzle.
* `<day>_test.go` containing unit tests primarily based on examples provided in the description of the puzzle.

Any code that can be re-used in muliple puzzles is located in its own files.

## Running the puzzles

Specify what day you want to run as an argument (eg. `go run . <day>`) or run all days by not specifying a day (eg. `go run .`).

## Why the `days` global variable and use of `init()` functions?

I have found this to be the easiest way to separate the `main()` function from the individual puzzles thus allow me to add days without having to touch `main()`.
