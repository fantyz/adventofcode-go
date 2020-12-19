package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

func init() { days["18"] = Day18 }

/*
--- Day 18: Operation Order ---
As you look out the window and notice a heavily-forested continent slowly appear over the horizon, you are interrupted by the child sitting next to you. They're curious if you could help them with their math homework.

Unfortunately, it seems like this "math" follows different rules than you remember.

The homework (your puzzle input) consists of a series of expressions that consist of addition (+), multiplication (*), and parentheses ((...)). Just like normal math, parentheses indicate that the expression inside must be evaluated before it can be used by the surrounding expression. Addition still finds the sum of the numbers on both sides of the operator, and multiplication still finds the product.

However, the rules of operator precedence have changed. Rather than evaluating multiplication before addition, the operators have the same precedence, and are evaluated left-to-right regardless of the order in which they appear.

For example, the steps to evaluate the expression 1 + 2 * 3 + 4 * 5 + 6 are as follows:

1 + 2 * 3 + 4 * 5 + 6
  3   * 3 + 4 * 5 + 6
      9   + 4 * 5 + 6
         13   * 5 + 6
             65   + 6
                 71
Parentheses can override this order; for example, here is what happens if parentheses are added to form 1 + (2 * 3) + (4 * (5 + 6)):

1 + (2 * 3) + (4 * (5 + 6))
1 +    6    + (4 * (5 + 6))
     7      + (4 * (5 + 6))
     7      + (4 *   11   )
     7      +     44
            51
Here are a few more examples:

2 * 3 + (4 * 5) becomes 26.
5 + (8 * 3 + 9 + 3 * 4 * 3) becomes 437.
5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4)) becomes 12240.
((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2 becomes 13632.
Before you can help with the homework, you need to understand it yourself. Evaluate the expression on each line of the homework; what is the sum of the resulting values?

Your puzzle answer was 1890866893020.

--- Part Two ---
You manage to answer the child's questions and they finish part 1 of their homework, but get stuck when they reach the next section: advanced math.

Now, addition and multiplication have different precedence levels, but they're not the ones you're familiar with. Instead, addition is evaluated before multiplication.

For example, the steps to evaluate the expression 1 + 2 * 3 + 4 * 5 + 6 are now as follows:

1 + 2 * 3 + 4 * 5 + 6
  3   * 3 + 4 * 5 + 6
  3   *   7   * 5 + 6
  3   *   7   *  11
     21       *  11
         231
Here are the other examples from above:

1 + (2 * 3) + (4 * (5 + 6)) still becomes 51.
2 * 3 + (4 * 5) becomes 46.
5 + (8 * 3 + 9 + 3 * 4 * 3) becomes 1445.
5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4)) becomes 669060.
((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2 becomes 23340.
What do you get if you add up the results of evaluating the homework problems using these new rules?

Your puzzle answer was 34646237037193.
*/

// NOTE: This puzzle is simple enough that I can get away with a solution that does
// string manipulation to do things in the right order. A much nicer and more robust
// solution would have been to build an actual abstract syntax tree based on the
// precedence rules.

func Day18() {
	fmt.Println("--- Day 18: Operation Order ---")
	sum, err := SumExpressions(day18Input, true)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Sum of all expressions (equal precedence):", sum)
	sum, err = SumExpressions(day18Input, false)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Sum of all expressions (not equal precedence):", sum)
}

// SumExpressions takes multiple expressions separated by newlines and evaluates them
// one by one summing up the total result. SumExpressions will return an error if any
// of the expressions cannot be evaluated.
func SumExpressions(in string, equalPrecedence bool) (int, error) {
	sum := 0
	for _, exp := range strings.Split(in, "\n") {
		res, err := Evaluate(exp, equalPrecedence)
		if err != nil {
			return 0, errors.Wrapf(err, "bad expression (exp=%s)", exp)
		}
		sum += res
	}
	return sum, nil
}

// Evaluate evaluates the provided expression and return the result. Evaluate returns
// an error if the expression cannot be evaluated.
func Evaluate(exp string, equalPrecedence bool) (int, error) {
	// evaluate paranthesis first
	parExp := regexp.MustCompile(`^(.*)\(([^)]+)\)(.*)$`)
	for {
		m := parExp.FindStringSubmatch(exp)
		if len(m) != 4 {
			// no paranthesis
			break
		}
		res, err := Evaluate(m[2], equalPrecedence)
		if err != nil {
			return 0, err
		}
		replaceStr := fmt.Sprintf("${1}%d${3}", res)
		exp = parExp.ReplaceAllString(exp, replaceStr)
	}

	if !equalPrecedence {
		// evaluate additions next
		addExp := regexp.MustCompile(`^(.* )?([0-9]+) \+ ([0-9]+)( .*)?$`)
		for {
			m := addExp.FindStringSubmatch(exp)
			if len(m) != 5 {
				// no more additions
				break
			}

			i, err := strconv.Atoi(m[2])
			if err != nil {
				return 0, err
			}
			j, err := strconv.Atoi(m[3])
			if err != nil {
				return 0, err
			}

			replaceStr := fmt.Sprintf("${1}%d${4}", i+j)
			exp = addExp.ReplaceAllString(exp, replaceStr)
		}
	}

	// evaluate expression without paranthesis and equal operator precedence
	tokens := strings.Split(exp, " ")
	if len(tokens) < 1 {
		return 0, nil
	}
	if len(tokens)%2 != 1 {
		return 0, errors.Errorf("the number of tokens in the expression must be uneven for the expression to be valid (len=%d)", len(tokens))
	}

	n, err := strconv.Atoi(tokens[0])
	if err != nil {
		return 0, errors.Wrapf(err, "unable to read int from token (token=%s)", tokens[0])
	}

	for i := 1; i < len(tokens); i += 2 {
		// evaluate an operator along with the next number
		m, err := strconv.Atoi(tokens[i+1])
		if err != nil {
			return 0, errors.Wrapf(err, "unable to read int from token (token=%s)", tokens[i+1])
		}
		switch tokens[i] {
		case "+":
			n += m
		case "*":
			n *= m
		default:
			return 0, errors.Errorf("unknown operator (op=%s)", tokens[i])
		}

	}

	return n, nil
}
