package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)
func TestSumMiddlePageNumberOfValidSections(t *testing.T) {
	assert.Equal(t, 143, SumMiddlePageNumberOfValidSections(`47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47`))
}

func TestSumMiddlePageNumberOfInvalidSectionsAfterReordering(t *testing.T) {
	assert.Equal(t, 123, SumMiddlePageNumberOfInvalidSectionsAfterReordering(`47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47`))
}

func TestDay05Pt1(t *testing.T) {
	assert.Equal(t, 5713, SumMiddlePageNumberOfValidSections(day05Input))
}

func TestDay05Pt2(t *testing.T) {
	assert.Equal(t, 5180, SumMiddlePageNumberOfInvalidSectionsAfterReordering(day05Input))
}

