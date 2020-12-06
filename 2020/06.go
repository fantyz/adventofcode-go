package main

import (
	"fmt"
	"strings"
)

func init() { days["6"] = Day6 }

/*
--- Day 6: Custom Customs ---
As your flight approaches the regional airport where you'll switch to a much larger plane, customs declaration forms are distributed to the passengers.

The form asks a series of 26 yes-or-no questions marked a through z. All you need to do is identify the questions for which anyone in your group answers "yes". Since your group is just you, this doesn't take very long.

However, the person sitting next to you seems to be experiencing a language barrier and asks if you can help. For each of the people in their group, you write down the questions for which they answer "yes", one per line. For example:

abcx
abcy
abcz
In this group, there are 6 questions to which anyone answered "yes": a, b, c, x, y, and z. (Duplicate answers to the same question don't count extra; each question counts at most once.)

Another group asks for your help, then another, and eventually you've collected answers from every group on the plane (your puzzle input). Each group's answers are separated by a blank line, and within each group, each person's answers are on a single line. For example:

abc

a
b
c

ab
ac

a
a
a
a

b
This list represents answers from five groups:

The first group contains one person who answered "yes" to 3 questions: a, b, and c.
The second group contains three people; combined, they answered "yes" to 3 questions: a, b, and c.
The third group contains two people; combined, they answered "yes" to 3 questions: a, b, and c.
The fourth group contains four people; combined, they answered "yes" to only 1 question, a.
The last group contains one person who answered "yes" to only 1 question, b.
In this example, the sum of these counts is 3 + 3 + 3 + 1 + 1 = 11.

For each group, count the number of questions to which anyone answered "yes". What is the sum of those counts?

Your puzzle answer was 6351.

--- Part Two ---
As you finish the last group's customs declaration, you notice that you misread one word in the instructions:

You don't need to identify the questions to which anyone answered "yes"; you need to identify the questions to which everyone answered "yes"!

Using the same example as above:

abc

a
b
c

ab
ac

a
a
a
a

b
This list represents answers from five groups:

In the first group, everyone (all 1 person) answered "yes" to 3 questions: a, b, and c.
In the second group, there is no question to which everyone answered "yes".
In the third group, everyone answered yes to only 1 question, a. Since some people did not answer "yes" to b or c, they don't count.
In the fourth group, everyone answered yes to only 1 question, a.
In the fifth group, everyone (all 1 person) answered "yes" to 1 question, b.
In this example, the sum of these counts is 3 + 0 + 1 + 1 + 1 = 6.

For each group, count the number of questions to which everyone answered "yes". What is the sum of those counts?

Your puzzle answer was 3143.
*/

func Day6() {
	fmt.Println("--- Day 6: Custom Customs ---")
	fmt.Println("  Yes count for all groups (any):", ProcessAllForms(day6Input, true))
	fmt.Println("  Yes count for all groups (all):", ProcessAllForms(day6Input, false))
}

// ProcessAllForms takes forms data and counts the amount of yes answers among these. It returns the number
// of questions that each group in the forms data has answered yes to. Either by counting any yes to a
// question if anyYes is set to true or otherwise only count if all in the group has answered yes to it.
func ProcessAllForms(formsData string, anyYes bool) int {
	// make sure formsData ends with an empty line to ease processing
	formsData += "\n\n"

	yesCount := 0
	var group []string
	for _, form := range strings.Split(formsData, "\n") {
		if form == "" && len(group) > 0 {
			groupYeses := ProcessGroupForms(group)
			if anyYes {
				// count a yes for questions anyone answered yes to in the group
				yesCount += len(groupYeses)
			} else {
				// only count a yes for questions where all in the group answered yes
				for _, count := range groupYeses {
					if count == len(group) {
						yesCount++
					}
				}
			}

			group = group[:0]
			continue
		}
		group = append(group, form)
	}
	return yesCount
}

// ProcessGroupForms takes a list of forms and returns a map containing the individual questions
// that someone has answered yes to as well as how many from the group that answered yes.
func ProcessGroupForms(forms []string) map[byte]int {
	yesAnswers := map[byte]int{}
	for _, form := range forms {
		for i := 0; i < len(form); i++ {
			yesAnswers[form[i]]++
		}
	}
	return yesAnswers
}
