package main

import (
	"fmt"
	"strings"
)

func init() { days["02"] = Day02 }

/*
--- Day 2: Red-Nosed Reports ---
Fortunately, the first location The Historians want to search isn't a long walk from the Chief Historian's office.

While the Red-Nosed Reindeer nuclear fusion/fission plant appears to contain no sign of the Chief Historian, the engineers there run up to you as soon as they see you. Apparently, they still talk about the time Rudolph was saved through molecular synthesis from a single electron.

They're quick to add that - since you're already here - they'd really appreciate your help analyzing some unusual data from the Red-Nosed reactor. You turn to check if The Historians are waiting for you, but they seem to have already divided into groups that are currently searching every corner of the facility. You offer to help with the unusual data.

The unusual data (your puzzle input) consists of many reports, one report per line. Each report is a list of numbers called levels that are separated by spaces. For example:

7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9
This example data contains six reports each containing five levels.

The engineers are trying to figure out which reports are safe. The Red-Nosed reactor safety systems can only tolerate levels that are either gradually increasing or gradually decreasing. So, a report only counts as safe if both of the following are true:

The levels are either all increasing or all decreasing.
Any two adjacent levels differ by at least one and at most three.
In the example above, the reports can be found safe or unsafe by checking those rules:

7 6 4 2 1: Safe because the levels are all decreasing by 1 or 2.
1 2 7 8 9: Unsafe because 2 7 is an increase of 5.
9 7 6 2 1: Unsafe because 6 2 is a decrease of 4.
1 3 2 4 5: Unsafe because 1 3 is increasing but 3 2 is decreasing.
8 6 4 4 1: Unsafe because 4 4 is neither an increase or a decrease.
1 3 6 7 9: Safe because the levels are all increasing by 1, 2, or 3.
So, in this example, 2 reports are safe.

Analyze the unusual data from the engineers. How many reports are safe?

Your puzzle answer was 490.

--- Part Two ---
The engineers are surprised by the low number of safe reports until they realize they forgot to tell you about the Problem Dampener.

The Problem Dampener is a reactor-mounted module that lets the reactor safety systems tolerate a single bad level in what would otherwise be a safe report. It's like the bad level never happened!

Now, the same rules apply as before, except if removing a single level from an unsafe report would make it safe, the report instead counts as safe.

More of the above example's reports are now safe:

7 6 4 2 1: Safe without removing any level.
1 2 7 8 9: Unsafe regardless of which level is removed.
9 7 6 2 1: Unsafe regardless of which level is removed.
1 3 2 4 5: Safe by removing the second level, 3.
8 6 4 4 1: Safe by removing the third level, 4.
1 3 6 7 9: Safe without removing any level.
Thanks to the Problem Dampener, 4 reports are actually safe!

Update your analysis by handling situations where the Problem Dampener can remove a single level from unsafe reports. How many reports are now safe?

Your puzzle answer was 536.
*/

func Day02() {
	fmt.Println("--- Day 2: Red-Nosed Reports ---")
	reports := LoadReports(day02Input)
	fmt.Println("Part 1: Safe reports (wo. problem dampner):", CountSafeReports(reports, false))
	fmt.Println("Part 2: Safe reports (w. problem dampner):", CountSafeReports(reports, true))
}

func CountSafeReports(reports []Report, useProblemDampner bool) int {
	safe := 0
	for _, r := range reports {
		if r.Safe(useProblemDampner) {
			safe++
		}
	}
	return safe;
}

func LoadReports(in string) []Report {
	var reports []Report
	for _, line := range strings.Split(in, "\n") {
		reports = append(reports, NewReport(line))
	}
	return reports
}

func NewReport(in string) []int {
	return LoadInts(in)
}

type Report []int

func (r Report) Safe(useProblemDampner bool) bool {
	// Given that the problem dampner allows a level to be bad, it is hard to determine whether the levels are
	// increasing or decreasing. To work around this problem, we can try to validate assuming both. If either is
	// valid then we know the report is safe.
	return r.isSafe(useProblemDampner, true) || r.isSafe(useProblemDampner, false)
}

func (r Report) isSafe(useProblemDampner bool, isIncreasing bool) bool {
	// report is by definition safe if only containing one or no levels
	if len(r) <= 1 {
		return true
	}
	// track bad levels in the report
	foundBad := false

	i := 0
	for i<len(r)-1 {
		if !validateReportLevels(r[i], r[i+1], isIncreasing) {
			if !useProblemDampner {
				// no bad levels allowed
				return false
			}

			if foundBad {
				// multiple bad levels now allowed
				return false
			}
			foundBad = true

			// assuming i+1 is the bad level, let us try skipping it
			if i+2 >= len(r) {
				// no additional levels available, all is good
				break
			}

			if !validateReportLevels(r[i], r[i+2], isIncreasing) {
				if i == 0 {
					// let us assume that it is the very first element that is the bad level
					// here rather than it being the two subsequent ones

					i++
					continue
				}

				// multiple bad levels not allowed
				return false
			}

			// subsequent level was good, skip the bad one
			i += 2
			continue
		}

		// level was valid
		i++
	}

	return true
}

// validate takes two sequential levels from a report and validates if they are safe or not.
// validate returns true if the levels are safe.
func validateReportLevels(a, b int, isIncreasing bool) bool {
	// validate it is increasing or not
	if isIncreasing {
		if a >= b {
			return false
		}
	} else {
		if a <= b {
			return false
		}
	}

	// validate increment or decrement is acceptable
	diff := a - b

	return diff <= 3 && diff >= -3
}
