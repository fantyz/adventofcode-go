package main

import (
	"fmt"
)

/*

--- Day 9: Sensor Boost ---
You've just said goodbye to the rebooted rover and left Mars when you receive a faint distress signal coming from the asteroid belt. It must be the Ceres monitoring station!

In order to lock on to the signal, you'll need to boost your sensors. The Elves send up the latest BOOST program - Basic Operation Of System Test.

While BOOST (your puzzle input) is capable of boosting your sensors, for tenuous safety reasons, it refuses to do so until the computer it runs on passes some checks to demonstrate it is a complete Intcode computer.

Your existing Intcode computer is missing one key feature: it needs support for parameters in relative mode.

Parameters in mode 2, relative mode, behave very similarly to parameters in position mode: the parameter is interpreted as a position. Like position mode, parameters in relative mode can be read from or written to.

The important difference is that relative mode parameters don't count from address 0. Instead, they count from a value called the relative base. The relative base starts at 0.

The address a relative mode parameter refers to is itself plus the current relative base. When the relative base is 0, relative mode parameters and position mode parameters with the same value refer to the same address.

For example, given a relative base of 50, a relative mode parameter of -7 refers to memory address 50 + -7 = 43.

The relative base is modified with the relative base offset instruction:

Opcode 9 adjusts the relative base by the value of its only parameter. The relative base increases (or decreases, if the value is negative) by the value of the parameter.
For example, if the relative base is 2000, then after the instruction 109,19, the relative base would be 2019. If the next instruction were 204,-34, then the value at address 1985 would be output.

Your Intcode computer will also need a few other capabilities:

The computer's available memory should be much larger than the initial program. Memory beyond the initial program starts with the value 0 and can be read or written like any other memory. (It is invalid to try to access memory at a negative address, though.)
The computer should have support for large numbers. Some instructions near the beginning of the BOOST program will verify this capability.
Here are some example programs that use these features:

109,1,204,-1,1001,100,1,100,1008,100,16,101,1006,101,0,99 takes no input and produces a copy of itself as output.
1102,34915192,34915192,7,4,7,99,0 should output a 16-digit number.
104,1125899906842624,99 should output the large number in the middle.
The BOOST program will ask for a single input; run it in test mode by providing it the value 1. It will perform a series of checks on each opcode, output any opcodes (and the associated parameter modes) that seem to be functioning incorrectly, and finally output a BOOST keycode.

Once your Intcode computer is fully functional, the BOOST program should report no malfunctioning opcodes when run in test mode; it should only output a single value, the BOOST keycode. What BOOST keycode does it produce?

Your puzzle answer was 3839402290.

--- Part Two ---
You now have a complete Intcode computer.

Finally, you can lock on to the Ceres distress signal! You just need to boost your sensors using the BOOST program.

The program runs in sensor boost mode by providing the input instruction the value 2. Once run, it will boost the sensors automatically, but it might take a few seconds to complete the operation on slower hardware. In sensor boost mode, the program will output a single value: the coordinates of the distress signal.

Run the BOOST program in sensor boost mode. What are the coordinates of the distress signal?

Your puzzle answer was 35734.

*/

func main() {
	fmt.Println("Day 9: Sensor Boost")
	_, out := ExecuteOpcode(Load(puzzleInput), Inputter([]int{1}))
	fmt.Println("Keycode:", out)
	_, out = ExecuteOpcode(Load(puzzleInput), Inputter([]int{2}))
	fmt.Println("Coord:", out)
}

const puzzleInput = `1102,34463338,34463338,63,1007,63,34463338,63,1005,63,53,1102,3,1,1000,109,988,209,12,9,1000,209,6,209,3,203,0,1008,1000,1,63,1005,63,65,1008,1000,2,63,1005,63,904,1008,1000,0,63,1005,63,58,4,25,104,0,99,4,0,104,0,99,4,17,104,0,99,0,0,1101,0,26,1014,1102,1,30,1013,1101,22,0,1000,1101,0,35,1015,1101,0,34,1011,1102,0,1,1020,1102,1,481,1022,1101,0,36,1003,1102,1,28,1005,1101,857,0,1024,1101,20,0,1008,1101,0,385,1026,1102,37,1,1006,1101,33,0,1017,1101,0,38,1002,1102,23,1,1007,1102,32,1,1010,1101,29,0,1016,1102,1,25,1009,1102,1,27,1012,1101,24,0,1018,1101,474,0,1023,1102,1,39,1004,1101,0,31,1001,1102,378,1,1027,1101,0,848,1025,1102,21,1,1019,1102,760,1,1029,1102,1,1,1021,1101,769,0,1028,109,-6,2107,21,6,63,1005,63,199,4,187,1106,0,203,1001,64,1,64,1002,64,2,64,109,16,2101,0,-6,63,1008,63,39,63,1005,63,225,4,209,1106,0,229,1001,64,1,64,1002,64,2,64,109,5,2108,20,-7,63,1005,63,247,4,235,1105,1,251,1001,64,1,64,1002,64,2,64,109,-1,2108,36,-8,63,1005,63,267,1106,0,273,4,257,1001,64,1,64,1002,64,2,64,109,-13,1201,-1,0,63,1008,63,22,63,1005,63,299,4,279,1001,64,1,64,1106,0,299,1002,64,2,64,109,15,2102,1,-8,63,1008,63,20,63,1005,63,321,4,305,1106,0,325,1001,64,1,64,1002,64,2,64,109,-13,21108,40,40,8,1005,1011,347,4,331,1001,64,1,64,1105,1,347,1002,64,2,64,109,-2,1207,8,24,63,1005,63,363,1105,1,369,4,353,1001,64,1,64,1002,64,2,64,109,35,2106,0,-9,1001,64,1,64,1106,0,387,4,375,1002,64,2,64,109,-26,21102,41,1,3,1008,1013,41,63,1005,63,409,4,393,1106,0,413,1001,64,1,64,1002,64,2,64,109,2,1202,-6,1,63,1008,63,36,63,1005,63,433,1106,0,439,4,419,1001,64,1,64,1002,64,2,64,109,-3,21102,42,1,10,1008,1019,40,63,1005,63,463,1001,64,1,64,1106,0,465,4,445,1002,64,2,64,109,15,2105,1,-1,1001,64,1,64,1106,0,483,4,471,1002,64,2,64,109,-27,1207,3,23,63,1005,63,505,4,489,1001,64,1,64,1105,1,505,1002,64,2,64,109,13,2102,1,-9,63,1008,63,28,63,1005,63,525,1105,1,531,4,511,1001,64,1,64,1002,64,2,64,109,1,2101,0,-8,63,1008,63,35,63,1005,63,551,1105,1,557,4,537,1001,64,1,64,1002,64,2,64,109,6,21107,43,44,-4,1005,1013,575,4,563,1106,0,579,1001,64,1,64,1002,64,2,64,109,-9,1201,-4,0,63,1008,63,40,63,1005,63,599,1105,1,605,4,585,1001,64,1,64,1002,64,2,64,109,12,1206,1,621,1001,64,1,64,1106,0,623,4,611,1002,64,2,64,109,-22,1202,9,1,63,1008,63,23,63,1005,63,649,4,629,1001,64,1,64,1105,1,649,1002,64,2,64,109,17,1206,5,667,4,655,1001,64,1,64,1106,0,667,1002,64,2,64,109,-3,1205,9,685,4,673,1001,64,1,64,1106,0,685,1002,64,2,64,109,3,1208,-9,37,63,1005,63,707,4,691,1001,64,1,64,1105,1,707,1002,64,2,64,109,7,1205,-2,723,1001,64,1,64,1106,0,725,4,713,1002,64,2,64,109,-15,21101,44,0,8,1008,1015,45,63,1005,63,745,1105,1,751,4,731,1001,64,1,64,1002,64,2,64,109,28,2106,0,-7,4,757,1001,64,1,64,1106,0,769,1002,64,2,64,109,-12,21101,45,0,-5,1008,1018,45,63,1005,63,791,4,775,1105,1,795,1001,64,1,64,1002,64,2,64,109,-9,2107,26,-5,63,1005,63,815,1001,64,1,64,1106,0,817,4,801,1002,64,2,64,109,-1,21107,46,45,-3,1005,1010,833,1105,1,839,4,823,1001,64,1,64,1002,64,2,64,109,3,2105,1,8,4,845,1001,64,1,64,1106,0,857,1002,64,2,64,109,-9,1208,-4,37,63,1005,63,877,1001,64,1,64,1105,1,879,4,863,1002,64,2,64,109,8,21108,47,46,2,1005,1017,895,1106,0,901,4,885,1001,64,1,64,4,64,99,21102,1,27,1,21102,1,915,0,1106,0,922,21201,1,14429,1,204,1,99,109,3,1207,-2,3,63,1005,63,964,21201,-2,-1,1,21102,1,942,0,1105,1,922,21202,1,1,-1,21201,-2,-3,1,21101,957,0,0,1106,0,922,22201,1,-1,-2,1105,1,968,21201,-2,0,-2,109,-3,2105,1,0`
