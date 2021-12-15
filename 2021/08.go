package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/pkg/errors"
)

func init() { days["8"] = Day8 }

/*
--- Day 8: Seven Segment Search ---
You barely reach the safety of the cave when the whale smashes into the cave mouth, collapsing it. Sensors indicate another exit to this cave at a much greater depth, so you have no choice but to press on.

As your submarine slowly makes its way through the cave system, you notice that the four-digit seven-segment displays in your submarine are malfunctioning; they must have been damaged during the escape. You'll be in a lot of trouble without them, so you'd better figure out what's wrong.

Each digit of a seven-segment display is rendered by turning on or off any of seven segments named a through g:

  0:      1:      2:      3:      4:
 aaaa    ....    aaaa    aaaa    ....
b    c  .    c  .    c  .    c  b    c
b    c  .    c  .    c  .    c  b    c
 ....    ....    dddd    dddd    dddd
e    f  .    f  e    .  .    f  .    f
e    f  .    f  e    .  .    f  .    f
 gggg    ....    gggg    gggg    ....

  5:      6:      7:      8:      9:
 aaaa    aaaa    aaaa    aaaa    aaaa
b    .  b    .  .    c  b    c  b    c
b    .  b    .  .    c  b    c  b    c
 dddd    dddd    ....    dddd    dddd
.    f  e    f  .    f  e    f  .    f
.    f  e    f  .    f  e    f  .    f
 gggg    gggg    ....    gggg    gggg
So, to render a 1, only segments c and f would be turned on; the rest would be off. To render a 7, only segments a, c, and f would be turned on.

The problem is that the signals which control the segments have been mixed up on each display. The submarine is still trying to display numbers by producing output on signal wires a through g, but those wires are connected to segments randomly. Worse, the wire/segment connections are mixed up separately for each four-digit display! (All of the digits within a display use the same connections, though.)

So, you might know that only signal wires b and g are turned on, but that doesn't mean segments b and g are turned on: the only digit that uses two segments is 1, so it must mean segments c and f are meant to be on. With just that information, you still can't tell which wire (b/g) goes to which segment (c/f). For that, you'll need to collect more information.

For each display, you watch the changing signals for a while, make a note of all ten unique signal patterns you see, and then write down a single four digit output value (your puzzle input). Using the signal patterns, you should be able to work out which pattern corresponds to which digit.

For example, here is what you might see in a single entry in your notes:

acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab |
cdfeb fcadb cdfeb cdbaf
(The entry is wrapped here to two lines so it fits; in your notes, it will all be on a single line.)

Each entry consists of ten unique signal patterns, a | delimiter, and finally the four digit output value. Within an entry, the same wire/segment connections are used (but you don't know what the connections actually are). The unique signal patterns correspond to the ten different ways the submarine tries to render a digit using the current wire/segment connections. Because 7 is the only digit that uses three segments, dab in the above example means that to render a 7, signal lines d, a, and b are on. Because 4 is the only digit that uses four segments, eafb means that to render a 4, signal lines e, a, f, and b are on.

Using this information, you should be able to work out which combination of signal wires corresponds to each of the ten digits. Then, you can decode the four digit output value. Unfortunately, in the above example, all of the digits in the output value (cdfeb fcadb cdfeb cdbaf) use five segments and are more difficult to deduce.

For now, focus on the easy digits. Consider this larger example:

be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb |
fdgacbe cefdb cefbgd gcbe
edbfga begcd cbg gc gcadebf fbgde acbgfd abcde gfcbed gfec |
fcgedb cgb dgebacf gc
fgaebd cg bdaec gdafb agbcfd gdcbef bgcad gfac gcb cdgabef |
cg cg fdcagb cbg
fbegcd cbd adcefb dageb afcb bc aefdc ecdab fgdeca fcdbega |
efabcd cedba gadfec cb
aecbfdg fbg gf bafeg dbefa fcge gcbea fcaegb dgceab fcbdga |
gecf egdcabf bgf bfgea
fgeab ca afcebg bdacfeg cfaedg gcfdb baec bfadeg bafgc acf |
gebdcfa ecba ca fadegcb
dbcfg fgd bdegcaf fgec aegbdf ecdfab fbedc dacgb gdcebf gf |
cefg dcbef fcge gbcadfe
bdfegc cbegaf gecbf dfcage bdacg ed bedf ced adcbefg gebcd |
ed bcgafe cdgba cbgef
egadfb cdbfeg cegd fecab cgb gbdefca cg fgcdab egfdb bfceg |
gbdfcae bgc cg cgb
gcafb gcf dcaebfg ecagb gf abcdeg gaef cafbge fdbac fegbdc |
fgae cfgab fg bagce
Because the digits 1, 4, 7, and 8 each use a unique number of segments, you should be able to tell which combinations of signals correspond to those digits. Counting only digits in the output values (the part after | on each line), in the above example, there are 26 instances of digits that use a unique number of segments (highlighted above).

In the output values, how many times do digits 1, 4, 7, or 8 appear?

Your puzzle answer was 412.

--- Part Two ---
Through a little deduction, you should now be able to determine the remaining digits. Consider again the first example above:

acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab |
cdfeb fcadb cdfeb cdbaf
After some careful analysis, the mapping between signal wires and segments only make sense in the following configuration:

 dddd
e    a
e    a
 ffff
g    b
g    b
 cccc
So, the unique signal patterns would correspond to the following digits:

acedgfb: 8
cdfbe: 5
gcdfa: 2
fbcad: 3
dab: 7
cefabd: 9
cdfgeb: 6
eafb: 4
cagedb: 0
ab: 1
Then, the four digits of the output value can be decoded:

cdfeb: 5
fcadb: 3
cdfeb: 5
cdbaf: 3
Therefore, the output value for this entry is 5353.

Following this same process for each entry in the second, larger example above, the output value of each entry can be determined:

fdgacbe cefdb cefbgd gcbe: 8394
fcgedb cgb dgebacf gc: 9781
cg cg fdcagb cbg: 1197
efabcd cedba gadfec cb: 9361
gecf egdcabf bgf bfgea: 4873
gebdcfa ecba ca fadegcb: 8418
cefg dcbef fcge gbcadfe: 4548
ed bcgafe cdgba cbgef: 1625
gbdfcae bgc cg cgb: 8717
fgae cfgab fg bagce: 4315
Adding all of the output values in this larger example produces 61229.

For each entry, determine all of the wire/segment connections and decode the four-digit output values. What do you get if you add up all of the output values?

Your puzzle answer was 978171.
*/

func Day8() {
	fmt.Println("--- Day 8: Seven Segment Search ---")
	displays, err := ReadAllDisplays(day08Input)
	if err != nil {
		fmt.Println(errors.Wrap(err, "Failed to create displays"))
		return
	}
	fmt.Println("Number of times that 1, 4, 7 or 8 appears in the displays:", CountOneFourSevenAndEightDigits(displays))
	fmt.Println("Sum of all displays:", SumDisplays(displays))
}

// ReadAllDisplays takes the puzzle input and returns a list of the actual integer value of all
// the displays.
// ReadAllDisplays returns an error if anything unexpected is found in the input string.
func ReadAllDisplays(in string) ([]int, error) {
	var displays []int
	for _, line := range strings.Split(in, "\n") {
		d, err := ReadDisplay(line)
		if err != nil {
			return nil, errors.Wrapf(err, "unable to read display (line=%s)", line)
		}
		displays = append(displays, d)
	}
	return displays, nil
}

// ReadDisplay takes a display input line containing 10 unqiue signal digits along with the
// actual four digit value and returns the real integer value represened by the display.
// See the puzzle input for more information on the syntax of the display.
// ReadDisplay returns an error if it us unable to resolve the integer value of the display.
func ReadDisplay(in string) (int, error) {
	unique, value, err := loadDisplayDigits(in)
	if err != nil {
		return 0, errors.Wrapf(err, "failed to load display digits (input=%s)", in)
	}

	// To identify what digit sequence corresponds to what interger value we need to do
	// some deduction. As the puzzle description outlines, we know the digit sequence
	// of 1, 4, 7 and 8 alone by looking at the number of charaters in the sequence.
	//
	// There is a number of ways we can deduce the remaining numbers. This approach
	// tries to rely on only on whether a number is the only one remaining that contains
	// a certain number of segments (eg. 1 is the only one that contains 2 segments) and
	// whether a sequence is a subset of another sequence.
	//
	// Each number can be identified uniquely in this way:
	//   0: Only 6-segment number which 5's segments is not a complete subset of
	//   1: Only unidentified number with 2 segments left
	//   2: Only unidentified number with 5 segments left (3 and 5 can be identified)
	//   3: Only 5-segment number which 1's segments is a complete subset of
	//   4: Only unidentified number with 4 segments left
	//   5: Only 5-segment number which is a complete subset of 6's segments
	//   6: Only 6-segment number which 1's segments is not a complete subset of
	//   7: Only unidentified number with 3 segments left
	//   8: Only unidentified number with 7 segments left
	//   9: Only 6-segment number which 3's segments is a complete subset of
	//

	numToDigit := map[int]Digit{}
	digitToNum := map[Digit]int{}

	// we might need several passes on unique to identify all the digits given many of
	// of the digits depends on other digits having been identified first
	infLoopCheckIn := len(unique)
	infLoopCheckDigitsLeft := 0
	var digit Digit
	for len(unique) > 0 {
		// make sure we catch it if we end up in an infite loop
		if infLoopCheckIn == 0 {
			if infLoopCheckDigitsLeft == len(unique) {
				return 0, errors.New("input does not allow determining the value in the display")
			}
			infLoopCheckIn = len(unique)
			infLoopCheckDigitsLeft = len(unique)
		}
		infLoopCheckIn--

		// pop the first element
		digit, unique = unique[0], unique[1:]

		mustBe := -1
		switch len(digit) {
		case 2:
			mustBe = 1
		case 3:
			mustBe = 7
		case 4:
			mustBe = 4
		case 5:
			// 2
			_, found3 := numToDigit[3]
			_, found5 := numToDigit[5]
			if found3 && found5 {
				mustBe = 2
				break
			}

			// 3
			digit1, found1 := numToDigit[1]
			if found1 && digit1.IsSubsetOf(digit) {
				mustBe = 3
				break
			}

			// 5
			digit6, found6 := numToDigit[6]
			if found6 && digit.IsSubsetOf(digit6) {
				mustBe = 5
				break
			}
		case 6:
			// 0
			digit5, found5 := numToDigit[5]
			if found5 && !digit5.IsSubsetOf(digit) {
				mustBe = 0
				break
			}

			// 6
			digit1, found1 := numToDigit[1]
			if found1 && !digit1.IsSubsetOf(digit) {
				mustBe = 6
				break
			}

			// 9
			digit3, found3 := numToDigit[3]
			if found3 && digit3.IsSubsetOf(digit) {
				mustBe = 9
			}
		case 7:
			mustBe = 8
		default:
			// should never happen
			panic(fmt.Sprintf("bad digit length (digit=%s)", digit))
		}

		if mustBe < 0 {
			// unable to identify the digit yet, re-add it to unique
			unique = append(unique, digit)
			continue
		}

		numToDigit[mustBe] = digit
		digitToNum[digit] = mustBe
	}

	// convert the value digits into a number
	val := 0
	for _, digit := range value {
		// add the digit as the rightmost digit in val
		val = val*10 + digitToNum[digit]
	}

	return val, nil
}

// loadDisplayDigits takes a display input line and returns the corresponding 10 unique digits
// and 4 value digits.
// loadDisplayDigits will return an error if anything unexpected is found in the input string.
func loadDisplayDigits(in string) ([]Digit, []Digit, error) {
	idx := strings.Index(in, " | ")
	if idx < 0 {
		return nil, nil, errors.New("no separator found")
	}
	uniqueInput := strings.Split(in[:idx], " ")
	if len(uniqueInput) != 10 {
		return nil, nil, errors.Errorf("unexpected number of unique signals (num=%d)", len(uniqueInput))
	}

	valueInput := strings.Split(in[idx+3:], " ")
	if len(valueInput) != 4 {
		return nil, nil, errors.Errorf("unexpected number of value displays (num=%d)", len(valueInput))
	}

	var err error
	unique := make([]Digit, 10)
	value := make([]Digit, 4)
	for i := 0; i < 10; i++ {
		if unique[i], err = NewDigit(uniqueInput[i]); err != nil {
			return nil, nil, errors.Wrapf(err, "bad digit (digit=%s)", uniqueInput[i])
		}
	}
	for i := 0; i < 4; i++ {
		if value[i], err = NewDigit(valueInput[i]); err != nil {
			return nil, nil, errors.Wrapf(err, "bad digit (digit=%s)", valueInput[i])
		}
	}

	return unique, value, nil
}

// Digit represents a single seven-segment display containing between 2 to 7 sorted segments
// valued from 'a' to 'g'.
type Digit string

// NewDigit takes a string input and converts this to Digit while ensuring that it contains
// 2-7 segments with valid values and the bytes are sorted.
func NewDigit(in string) (Digit, error) {
	d := []byte(in)
	if len(d) < 2 || len(d) > 7 {
		return "", errors.New("digit must be between 2 and 7 characters long")
	}
	for _, seg := range d {
		switch seg {
		case 'a', 'b', 'c', 'd', 'e', 'f', 'g':
			// all good
		default:
			return "", errors.New("digit can only contain characters from a-g")
		}
	}
	sort.Sort(bySegments(d))

	return Digit(d), nil
}

// IsSubsetOf returns true of the digit iself is a subset of the input digit meaning all
// segments in the digit also exist in the input digit.
func (d1 Digit) IsSubsetOf(d2 Digit) bool {
	d1Idx := 0
	for d2Idx := 0; d2Idx < len(d2); d2Idx++ {
		if d1Idx < len(d1) && d1[d1Idx] == d2[d2Idx] {
			d1Idx++
		}
	}
	return d1Idx == len(d1)
}

// CountOneFourSevenAndEightDigits counts the number of times the digits 1, 4, 7 or 8 occur in
// the displays.
func CountOneFourSevenAndEightDigits(displays []int) int {
	count := 0

	for _, v := range displays {
		for v > 0 {
			// check the right-most digit
			switch v % 10 {
			case 1, 4, 7, 8:
				count++
			}

			// discard the right-most digit
			v = v / 10
		}
	}

	return count
}

// SumDisplays sums the value of each display.
func SumDisplays(displays []int) int {
	sum := 0
	for _, v := range displays {
		sum += v
	}
	return sum
}

//  bySegments allows sorting a slice of bytes representing segments
type bySegments []byte

func (s bySegments) Less(i, j int) bool { return s[i] < s[j] }
func (s bySegments) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s bySegments) Len() int           { return len(s) }
