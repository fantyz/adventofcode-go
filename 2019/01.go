package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
--- Day 1: The Tyranny of the Rocket Equation ---
Santa has become stranded at the edge of the Solar System while delivering presents to other planets! To accurately calculate his position in space, safely align his warp drive, and return to Earth in time to save Christmas, he needs you to bring him measurements from fifty stars.

Collect stars by solving puzzles. Two puzzles will be made available on each day in the Advent calendar; the second puzzle is unlocked when you complete the first. Each puzzle grants one star. Good luck!

The Elves quickly load you into a spacecraft and prepare to launch.

At the first Go / No Go poll, every Elf is Go until the Fuel Counter-Upper. They haven't determined the amount of fuel required yet.

Fuel required to launch a given module is based on its mass. Specifically, to find the fuel required for a module, take its mass, divide by three, round down, and subtract 2.

For example:

For a mass of 12, divide by 3 and round down to get 4, then subtract 2 to get 2.
For a mass of 14, dividing by 3 and rounding down still yields 4, so the fuel required is also 2.
For a mass of 1969, the fuel required is 654.
For a mass of 100756, the fuel required is 33583.
The Fuel Counter-Upper needs to know the total fuel requirement. To find it, individually calculate the fuel needed for the mass of each module (your puzzle input), then add together all the fuel values.

What is the sum of the fuel requirements for all of the modules on your spacecraft?

Your puzzle answer was 3266288.

--- Part Two ---
During the second Go / No Go poll, the Elf in charge of the Rocket Equation Double-Checker stops the launch sequence. Apparently, you forgot to include additional fuel for the fuel you just added.

Fuel itself requires fuel just like a module - take its mass, divide by three, round down, and subtract 2. However, that fuel also requires fuel, and that fuel requires fuel, and so on. Any mass that would require negative fuel should instead be treated as if it requires zero fuel; the remaining mass, if any, is instead handled by wishing really hard, which has no mass and is outside the scope of this calculation.

So, for each module mass, calculate its fuel and add it to the total. Then, treat the fuel amount you just calculated as the input mass and repeat the process, continuing until a fuel requirement is zero or negative. For example:

A module of mass 14 requires 2 fuel. This fuel requires no further fuel (2 divided by 3 and rounded down is 0, which would call for a negative fuel), so the total fuel required is still just 2.
At first, a module of mass 1969 requires 654 fuel. Then, this fuel requires 216 more fuel (654 / 3 - 2). 216 then requires 70 more fuel, which requires 21 fuel, which requires 5 fuel, which requires no further fuel. So, the total fuel required for a module of mass 1969 is 654 + 216 + 70 + 21 + 5 = 966.
The fuel required by a module of mass 100756 and its fuel is: 33583 + 11192 + 3728 + 1240 + 411 + 135 + 43 + 12 + 2 = 50346.
What is the sum of the fuel requirements for all of the modules on your spacecraft when also taking into account the mass of the added fuel? (Calculate the fuel requirements for each module separately, then add them all up at the end.)

Your puzzle answer was 4896582.
*/

func main() {
	fmt.Println("Day 1 - The Tyranny of the Rocket Equation")
	naive, real := RequiredFuel(puzzleInput)
	fmt.Println("Required fuel (naive): ", naive)
	fmt.Println("Required fuel (real):  ", real)
}

func RequiredFuel(modules string) (naive, real int) {
	for _, l := range strings.Split(modules, "\n") {
		i, err := strconv.Atoi(l)
		if err != nil {
			panic(err)
		}

		f := FuelUsage(i)
		naive += f
		for f > 0 {
			real += f
			f = FuelUsage(f)
		}
	}
	return
}

func FuelUsage(weight int) int {
	fuel := weight/3 - 2
	if fuel < 0 {
		return 0
	}
	return fuel
}

const puzzleInput = `128398
118177
139790
84818
75859
139920
90212
74975
120844
85533
77851
127044
128094
77724
81951
115804
60506
65055
52549
108749
92367
53974
52896
66403
93539
118392
78768
128172
85643
109508
104742
71305
84558
68640
58328
58404
70131
73745
149553
57511
119045
90210
129537
114869
113353
114181
130737
134877
90983
84361
62750
114532
139233
139804
130391
144731
84309
137050
79866
121266
93502
132060
109190
61326
58826
129305
141059
143017
56552
102142
110604
136052
93872
71951
72954
70701
137381
76580
62535
62666
126366
66361
109076
126230
73367
94459
126314
133327
143771
50752
75607
117606
142366
59068
75574
149836
57058
77622
83276
82734`
