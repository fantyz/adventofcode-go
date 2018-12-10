package main

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
)

/*

--- Day 7: The Sum of Its Parts ---
You find yourself standing on a snow-covered coastline; apparently, you landed a little off course. The region is too hilly to see the North Pole from here, but you do spot some Elves that seem to be trying to unpack something that washed ashore. It's quite cold out, so you decide to risk creating a paradox by asking them for directions.

"Oh, are you the search party?" Somehow, you can understand whatever Elves from the year 1018 speak; you assume it's Ancient Nordic Elvish. Could the device on your wrist also be a translator? "Those clothes don't look very warm; take this." They hand you a heavy coat.

"We do need to find our way back to the North Pole, but we have higher priorities at the moment. You see, believe it or not, this box contains something that will solve all of Santa's transportation problems - at least, that's what it looks like from the pictures in the instructions." It doesn't seem like they can read whatever language it's in, but you can: "Sleigh kit. Some assembly required."

"'Sleigh'? What a wonderful name! You must help us assemble this 'sleigh' at once!" They start excitedly pulling more parts out of the box.

The instructions specify a series of steps and requirements about which steps must be finished before others can begin (your puzzle input). Each step is designated by a single letter. For example, suppose you have the following instructions:

Step C must be finished before step A can begin.
Step C must be finished before step F can begin.
Step A must be finished before step B can begin.
Step A must be finished before step D can begin.
Step B must be finished before step E can begin.
Step D must be finished before step E can begin.
Step F must be finished before step E can begin.
Visually, these requirements look like this:


  -->A--->B--
 /    \      \
C      -->D----->E
 \           /
  ---->F-----
Your first goal is to determine the order in which the steps should be completed. If more than one step is ready, choose the step which is first alphabetically. In this example, the steps would be completed as follows:

Only C is available, and so it is done first.
Next, both A and F are available. A is first alphabetically, so it is done next.
Then, even though F was available earlier, steps B and D are now also available, and B is the first alphabetically of the three.
After that, only D and F are available. E is not available because only some of its prerequisites are complete. Therefore, D is completed next.
F is the only choice, so it is done next.
Finally, E is completed.
So, in this example, the correct order is CABDFE.

In what order should the steps in your instructions be completed?

Your puzzle answer was ABLCFNSXZPRHVEGUYKDIMQTWJO.

--- Part Two ---
As you're about to begin construction, four of the Elves offer to help. "The sun will set soon; it'll go faster if we work together." Now, you need to account for multiple people working on steps simultaneously. If multiple steps are available, workers should still begin them in alphabetical order.

Each step takes 60 seconds plus an amount corresponding to its letter: A=1, B=2, C=3, and so on. So, step A takes 60+1=61 seconds, while step Z takes 60+26=86 seconds. No time is required between steps.

To simplify things for the example, however, suppose you only have help from one Elf (a total of two workers) and that each step takes 60 fewer seconds (so that step A takes 1 second and step Z takes 26 seconds). Then, using the same instructions as above, this is how each second would be spent:

Second   Worker 1   Worker 2   Done
   0        C          .
   1        C          .
   2        C          .
   3        A          F       C
   4        B          F       CA
   5        B          F       CA
   6        D          F       CAB
   7        D          F       CAB
   8        D          F       CAB
   9        D          .       CABF
  10        E          .       CABFD
  11        E          .       CABFD
  12        E          .       CABFD
  13        E          .       CABFD
  14        E          .       CABFD
  15        .          .       CABFDE
Each row represents one second of time. The Second column identifies how many seconds have passed as of the beginning of that second. Each worker column shows the step that worker is currently doing (or . if they are idle). The Done column shows completed steps.

Note that the order of the steps has changed; this is because steps now take time to finish and multiple workers can begin multiple steps simultaneously.

In this example, it would take 15 seconds for two workers to complete these steps.

With 5 workers and the 60+ second step durations described above, how long will it take to complete all of the steps?

*/

func main() {
	fmt.Println("Day 7: The Sum of Its Parts")
	order, time := NewGraph(puzzleInput).Order(1, 0)
	fmt.Println(" > Order (1, 0):", order, time)
	order, time = NewGraph(puzzleInput).Order(5, 60)
	fmt.Println(" > Order (5,60):", order, time)
}

type Step struct {
	Name     string
	parents  []*Step
	children []*Step
}

type Graph []*Step

func NewGraph(in string) Graph {
	steps := map[string]*Step{}
	re := regexp.MustCompile(`^Step ([A-Z]) must be finished before step ([A-Z]) can begin\.$`)

	for _, l := range strings.Split(in, "\n") {
		m := re.FindStringSubmatch(l)
		if len(m) <= 0 {
			panic("line did not match: " + l)
		}

		p := steps[m[1]]
		if p == nil {
			p = &Step{Name: m[1]}
			steps[p.Name] = p
		}
		c := steps[m[2]]
		if c == nil {
			c = &Step{Name: m[2]}
			steps[c.Name] = c
		}
		p.children = append(p.children, c)
		c.parents = append(c.parents, p)
	}

	var res Graph
	for _, step := range steps {
		isChild := false
		for _, s := range steps {
			for _, c := range s.children {
				if c == step {
					isChild = true
				}
			}
		}
		if !isChild {
			res = append(res, step)
		}
	}

	return res
}

type Worker struct {
	step     *Step
	timeLeft int
}

func (g Graph) Order(workerCount int, offset int) (string, int) {
	workers := make([]Worker, workerCount)

	seen := map[string]struct{}{}
	avail := []*Step(g)
	order := []string{}

	sec := -1
	active := 0
	for active > 0 || len(avail) > 0 {
		sec++
		//fmt.Println("Second", sec)

		for i := range workers {
			workers[i].timeLeft--
			if workers[i].timeLeft < -1 {
				workers[i].timeLeft = -1
			}

			if workers[i].timeLeft == 0 {
				active--

				order = append(order, workers[i].step.Name)
				seen[workers[i].step.Name] = struct{}{}

				for _, c := range workers[i].step.children {
					ready := true
					for _, p := range c.parents {
						if _, found := seen[p.Name]; !found {
							ready = false
						}
					}
					if ready {
						avail = append(avail, c)
					}
				}
				workers[i].step = nil
			}
		}

		sort.Slice(avail, func(i, j int) bool { return avail[i].Name < avail[j].Name })

		for i := range workers {
			if workers[i].timeLeft <= 0 && len(avail) > 0 {
				// feed worker
				active++

				workers[i] = Worker{
					step:     avail[0],
					timeLeft: offset + avail[0].Time(),
				}
				avail = avail[1:]
			}
		}

		/*
			for i := range workers {
				step := workers[i].step
				name := "-"
				if step != nil {
					name = step.Name
				}
				fmt.Printf("%d: %s (%d)\n", i, name, workers[i].timeLeft)
			}
			fmt.Println(strings.Join(order, ""))
			fmt.Println()
		*/
	}

	return strings.Join(order, ""), sec
}

func (s *Step) Time() int {
	if len(s.Name) != 1 {
		panic("Bad name")
	}
	return int(s.Name[0]) - 64
}

const puzzleInput = `Step A must be finished before step L can begin.
Step B must be finished before step U can begin.
Step S must be finished before step K can begin.
Step L must be finished before step R can begin.
Step C must be finished before step I can begin.
Step F must be finished before step N can begin.
Step X must be finished before step H can begin.
Step Z must be finished before step U can begin.
Step P must be finished before step T can begin.
Step R must be finished before step U can begin.
Step H must be finished before step T can begin.
Step V must be finished before step G can begin.
Step E must be finished before step D can begin.
Step G must be finished before step W can begin.
Step N must be finished before step J can begin.
Step U must be finished before step D can begin.
Step Y must be finished before step K can begin.
Step K must be finished before step J can begin.
Step D must be finished before step M can begin.
Step I must be finished before step O can begin.
Step M must be finished before step Q can begin.
Step Q must be finished before step J can begin.
Step T must be finished before step J can begin.
Step W must be finished before step O can begin.
Step J must be finished before step O can begin.
Step C must be finished before step F can begin.
Step C must be finished before step J can begin.
Step Z must be finished before step I can begin.
Step K must be finished before step I can begin.
Step L must be finished before step W can begin.
Step I must be finished before step W can begin.
Step N must be finished before step O can begin.
Step B must be finished before step G can begin.
Step S must be finished before step O can begin.
Step P must be finished before step H can begin.
Step R must be finished before step J can begin.
Step N must be finished before step U can begin.
Step U must be finished before step J can begin.
Step E must be finished before step T can begin.
Step T must be finished before step O can begin.
Step L must be finished before step T can begin.
Step P must be finished before step Y can begin.
Step L must be finished before step C can begin.
Step D must be finished before step O can begin.
Step H must be finished before step Y can begin.
Step Q must be finished before step T can begin.
Step P must be finished before step G can begin.
Step G must be finished before step D can begin.
Step F must be finished before step H can begin.
Step G must be finished before step M can begin.
Step F must be finished before step V can begin.
Step X must be finished before step O can begin.
Step V must be finished before step Y can begin.
Step Y must be finished before step D can begin.
Step H must be finished before step G can begin.
Step A must be finished before step S can begin.
Step E must be finished before step U can begin.
Step Y must be finished before step O can begin.
Step C must be finished before step K can begin.
Step R must be finished before step W can begin.
Step G must be finished before step I can begin.
Step V must be finished before step E can begin.
Step V must be finished before step T can begin.
Step E must be finished before step K can begin.
Step X must be finished before step R can begin.
Step Q must be finished before step W can begin.
Step X must be finished before step P can begin.
Step K must be finished before step T can begin.
Step I must be finished before step T can begin.
Step P must be finished before step R can begin.
Step T must be finished before step W can begin.
Step X must be finished before step I can begin.
Step N must be finished before step Q can begin.
Step G must be finished before step Y can begin.
Step Y must be finished before step W can begin.
Step L must be finished before step D can begin.
Step F must be finished before step D can begin.
Step A must be finished before step T can begin.
Step R must be finished before step H can begin.
Step E must be finished before step I can begin.
Step W must be finished before step J can begin.
Step F must be finished before step M can begin.
Step V must be finished before step W can begin.
Step I must be finished before step J can begin.
Step Z must be finished before step P can begin.
Step H must be finished before step U can begin.
Step R must be finished before step V can begin.
Step V must be finished before step M can begin.
Step Y must be finished before step M can begin.
Step P must be finished before step M can begin.
Step K must be finished before step D can begin.
Step C must be finished before step T can begin.
Step Y must be finished before step T can begin.
Step U must be finished before step I can begin.
Step A must be finished before step O can begin.
Step E must be finished before step J can begin.
Step H must be finished before step V can begin.
Step F must be finished before step W can begin.
Step M must be finished before step T can begin.
Step S must be finished before step H can begin.
Step S must be finished before step G can begin.`
