package main

import (
	"fmt"
)

func init() { puzzles["1"] = func() { Day1(day1Input) } }

func Day1(in string) {
	fmt.Println("Day 1: Report Repair")
	vals := LoadInts(in)
	fmt.Println("Expense 2 reports that sum to 2020 multiplied:", Find2ExpenseReportEntriesMultiplied(vals))
	fmt.Println("Expense 3 reports that sum to 2020 multiplied:", Find3ExpenseReportEntriesMultiplied(vals))
}

func Find2ExpenseReportEntriesMultiplied(values []int) int {
	for n, i := range values {
		for m, j := range values {
			if n == m {
				continue
			}

			if i+j == 2020 {
				return i * j
			}
		}
	}
	panic("No expense report entries found")
}

func Find3ExpenseReportEntriesMultiplied(values []int) int {
	for n, i := range values {
		for m, j := range values {
			if n == m {
				continue
			}
			for o, k := range values {
				if n == o || m == o {
					continue
				}
				if i+j+k == 2020 {
					return i * j * k
				}
			}
		}
	}
	panic("No expense report entries found")
}
