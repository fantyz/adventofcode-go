package main

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"
)

func init() { days["19"] = Day19 }

/*
--- Day 19: Medicine for Rudolph ---
Rudolph the Red-Nosed Reindeer is sick! His nose isn't shining very brightly, and he needs medicine.

Red-Nosed Reindeer biology isn't similar to regular reindeer biology; Rudolph is going to need custom-made medicine. Unfortunately, Red-Nosed Reindeer chemistry isn't similar to regular reindeer chemistry, either.

The North Pole is equipped with a Red-Nosed Reindeer nuclear fusion/fission plant, capable of constructing any Red-Nosed Reindeer molecule you need. It works by starting with some input molecule and then doing a series of replacements, one per step, until it has the right molecule.

However, the machine has to be calibrated before it can be used. Calibration involves determining the number of molecules that can be generated in one step from a given starting point.

For example, imagine a simpler machine that supports only the following replacements:

H => HO
H => OH
O => HH
Given the replacements above and starting with HOH, the following molecules could be generated:

HOOH (via H => HO on the first H).
HOHO (via H => HO on the second H).
OHOH (via H => OH on the first H).
HOOH (via H => OH on the second H).
HHHH (via O => HH).
So, in the example above, there are 4 distinct molecules (not five, because HOOH appears twice) after one replacement from HOH. Santa's favorite molecule, HOHOHO, can become 7 distinct molecules (over nine replacements: six from H, and three from O).

The machine replaces without regard for the surrounding characters. For example, given the string H2O, the transition H => OO would result in OO2O.

Your puzzle input describes all of the possible replacements and, at the bottom, the medicine molecule for which you need to calibrate the machine. How many distinct molecules can be created after all the different ways you can do one replacement on the medicine molecule?

Your puzzle answer was 576.

--- Part Two ---
Now that the machine is calibrated, you're ready to begin molecule fabrication.

Molecule fabrication always begins with just a single electron, e, and applying replacements one at a time, just like the ones during calibration.

For example, suppose you have the following replacements:

e => H
e => O
H => HO
H => OH
O => HH
If you'd like to make HOH, you start with e, and then make the following replacements:

e => O to get O
O => HH to get HH
H => OH (on the second H) to get HOH
So, you could make HOH after 3 steps. Santa's favorite molecule, HOHOHO, can be made in 6 steps.

How long will it take to make the medicine? Given the available replacements and the medicine molecule in your puzzle input, what is the fewest number of steps to go from e to the medicine molecule?

Your puzzle answer was 207.--- Day 19: Medicine for Rudolph ---
Rudolph the Red-Nosed Reindeer is sick! His nose isn't shining very brightly, and he needs medicine.

Red-Nosed Reindeer biology isn't similar to regular reindeer biology; Rudolph is going to need custom-made medicine. Unfortunately, Red-Nosed Reindeer chemistry isn't similar to regular reindeer chemistry, either.

The North Pole is equipped with a Red-Nosed Reindeer nuclear fusion/fission plant, capable of constructing any Red-Nosed Reindeer molecule you need. It works by starting with some input molecule and then doing a series of replacements, one per step, until it has the right molecule.

However, the machine has to be calibrated before it can be used. Calibration involves determining the number of molecules that can be generated in one step from a given starting point.

For example, imagine a simpler machine that supports only the following replacements:

H => HO
H => OH
O => HH
Given the replacements above and starting with HOH, the following molecules could be generated:

HOOH (via H => HO on the first H).
HOHO (via H => HO on the second H).
OHOH (via H => OH on the first H).
HOOH (via H => OH on the second H).
HHHH (via O => HH).
So, in the example above, there are 4 distinct molecules (not five, because HOOH appears twice) after one replacement from HOH. Santa's favorite molecule, HOHOHO, can become 7 distinct molecules (over nine replacements: six from H, and three from O).

The machine replaces without regard for the surrounding characters. For example, given the string H2O, the transition H => OO would result in OO2O.

Your puzzle input describes all of the possible replacements and, at the bottom, the medicine molecule for which you need to calibrate the machine. How many distinct molecules can be created after all the different ways you can do one replacement on the medicine molecule?

Your puzzle answer was 576.

--- Part Two ---
Now that the machine is calibrated, you're ready to begin molecule fabrication.

Molecule fabrication always begins with just a single electron, e, and applying replacements one at a time, just like the ones during calibration.

For example, suppose you have the following replacements:

e => H
e => O
H => HO
H => OH
O => HH
If you'd like to make HOH, you start with e, and then make the following replacements:

e => O to get O
O => HH to get HH
H => OH (on the second H) to get HOH
So, you could make HOH after 3 steps. Santa's favorite molecule, HOHOHO, can be made in 6 steps.

How long will it take to make the medicine? Given the available replacements and the medicine molecule in your puzzle input, what is the fewest number of steps to go from e to the medicine molecule?

Your puzzle answer was 207.
*/

func Day19() {
	fmt.Println("--- Day 19: Medicine for Rudolph ---")
	replacer, molecule := NewMoleculeReplacerAndMolecule(day19Input)
	fmt.Println("Distinct molecules possible to make by the molecule replacement machine:", replacer.Calibrate(molecule))
	fmt.Println("Fewest steps needed to generate the medice for Rudolph:", replacer.FewestStepsTo(molecule))
}

// MoleculeReplacer can mutate a molecule using a number of replacement rules.
// The map[string][]string represents these replacement rules. The map index is the left
// side of the replacement and the slice of strings is all the possible replacements.
type MoleculeReplacer map[string][]string

// NewReplacementsAndMolecule takes the puzzle input and returns a MoleculeReplacer based
// off the possible replacements found in the input as well as the input molecule itself.
// NOTE: NewReplacementsAndMolecule does not perform any error checking on the input.
func NewMoleculeReplacerAndMolecule(input string) (MoleculeReplacer, string) {
	var molecule string
	replacer := MoleculeReplacer{}

	repExp := regexp.MustCompile(`^([A-Za-z]+) => ([A-Za-z]+)$`)

	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}
		m := repExp.FindStringSubmatch(line)
		if len(m) != 3 {
			// not a replacement or empty line, must be the molecule
			molecule = line
			continue
		}
		replacer[m[1]] = append(replacer[m[1]], m[2])
	}

	return replacer, molecule
}

// Calibrate returns the number of molecules that the MoleculeReplacer could generate after
// making a single replacement.
func (r MoleculeReplacer) Calibrate(molecule string) int {
	// multiple replacements could potentially create the same molecule
	// lets use a map[string]struct{} to get rid of duplicates.
	outputs := map[string]struct{}{}

	for rep, vals := range r {
		for _, val := range vals {
			remainderIdx := 0
			for {
				idx := strings.Index(molecule[remainderIdx:], rep)
				if idx < 0 {
					break
				}
				remainderIdx += idx + len(rep)

				// build new molecule
				outputs[molecule[:remainderIdx-len(rep)]+val+molecule[remainderIdx:]] = struct{}{}
			}
		}
	}

	return len(outputs)
}

// FewestStepsTo returns the fewest number of steps needed to generate the desired molecule
// starting from just a single electron ("e").
func (r MoleculeReplacer) FewestStepsTo(molecule string) int {
	// I'm not too happy about this one. It turns out that you need to analyze the input
	// to find a solution specific to it.
	//
	// Since I've made it a goal of its own to make solutions that work independently of
	// the input, I've struck out on making a fast solution for this part.
	//
	// ...or rather I hit the solution randomly due to the non-deterministic nature of
	// maps in Go while trying out different things. It seems, from reading the thread
	// on solutions linked below, that most people actually got lucky on inputs rather
	// than actually making a proper solution.
	//
	// The following implementation is not based on my own work - it is based on:
	// https://www.reddit.com/r/adventofcode/comments/3xflz8/day_19_solutions/cy4etju
	//
	//
	// First insight
	//
	// There are only two types of productions:
	//   1. e => XX and X => XX (X is not Rn, Y, or Ar)
	//   2. X => X Rn X Ar | X Rn X Y X Ar | X Rn X Y X Y X Ar
	//
	//
	// Second insight
	//
	// You can think of Rn Y Ar as the characters ( , ):
	//   X => X(X) | X(X,X) | X(X,X,X)
	//
	// Whenever there are two adjacent "elements" in your "molecule", you apply the first
	// production. This reduces your molecule length by 1 each time.
	//
	// And whenever you have T(T) T(T,T) or T(T,T,T) (T is a literal token such as "Mg",
	// i.e. not a nonterminal like "TiTiCaCa"), you apply the second production. This
	// reduces your molecule length by 3, 5, or 7.
	//
	//
	// Third insight
	//
	// Repeatedly applying X => XX until you arrive at a single token takes
	// count(tokens) - 1 steps:
	//   ABCDE => XCDE => XDE => XE => X
	//   count("ABCDE") = 5
	//   5 - 1 = 4 steps
	//
	// Applying X => X(X) is similar to X => XX, except you get the () for free. This can
	// be expressed as count(tokens) - count("(" or ")") - 1.
	//
	//   A(B(C(D(E)))) => A(B(C(X))) => A(B(X)) => A(X) => X
	//   count("A(B(C(D(E))))") = 13
	//   count("(((())))") = 8
	//   13 - 8 - 1 = 4 steps
	//
	// You can generalize to X => X(X,X) by noting that each , reduces the length by two
	// (,X). The new formula is count(tokens) - count("(" or ")") - 2*count(",") - 1.
	//
	//   A(B(C,D),E(F,G)) => A(B(C,D),X) => A(X,X) => X
	//   count("A(B(C,D),E(F,G))") = 16
	//   count("(()())") = 6
	//   count(",,,") = 3
	//   16 - 6 - 2*3 - 1 = 3 steps
	//
	// This final formula works for all of the production types (for X => XX, the (,)
	// counts are zero by definition.)

	elementsTotal := 0
	RnArs := 0
	Ys := 0

	i := 0
	for i < len(molecule) {
		elm := string(molecule[i])
		i++
		for i < len(molecule) && unicode.IsLower(rune(molecule[i])) {
			elm += string(molecule[i])
			i++
		}
		elementsTotal++
		switch elm {
		case "Rn", "Ar":
			RnArs++
		case "Y":
			Ys++
		}
	}

	return elementsTotal - RnArs - 2*Ys - 1
}
