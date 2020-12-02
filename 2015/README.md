# Advent of Code 2015

All puzzles are part of the same application. This is primarily done to avoid having a ton of modules and ease re-use of code between days where possible.

The `main.go` contains the basic logic needed to simply run the puzzle of one day.

The puzzles are organized by day and consist of:

* `<day>.go` containing a description of the puzzle of the day itself along with the code needed to solve it.
* `<day>_input.go` containing the puzzle input unique to me needed to solve the puzzle.
* `<day>_test.go` containing unit tests primarily based on examples provided in the description of the puzzle.

Any code that can be re-used in muliple puzzles is located in its own files.

## Running the puzzles

Specify what day you want to run as an argument (eg. `go run . <day>`) or run all days by not specifying a day (eg. `go run .`).

## Whats with `old/`?

Old contains the quick and dirty solution that I originally made. I am in the process of rewriting these into a much nicer and easier to read solution. This might take a while (years?) as I'm doing this as time permits. Once I'm done I'll remove `old/` entirely.


## Why the `days` global variable and use of `init()` functions?

I have found this to be the easiest way to separate the `main()` function from the individual puzzles thus allow me to add days without having to touch `main()`.
