package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

func init() { days["8"] = Day8 }

/*
--- Day 8: Matchsticks ---
Space on the sleigh is limited this year, and so Santa will be bringing his list as a digital copy. He needs to know how much space it will take up when stored.

It is common in many programming languages to provide a way to escape special characters in strings. For example, C, JavaScript, Perl, Python, and even PHP handle special characters in very similar ways.

However, it is important to realize the difference between the number of characters in the code representation of the string literal and the number of characters in the in-memory string itself.

For example:

"" is 2 characters of code (the two double quotes), but the string contains zero characters.
"abc" is 5 characters of code, but 3 characters in the string data.
"aaa\"aaa" is 10 characters of code, but the string itself contains six "a" characters and a single, escaped quote character, for a total of 7 characters in the string data.
"\x27" is 6 characters of code, but the string itself contains just one - an apostrophe ('), escaped using hexadecimal notation.
Santa's list is a file that contains many double-quoted string literals, one on each line. The only escape sequences used are \\ (which represents a single backslash), \" (which represents a lone double-quote character), and \x plus two hexadecimal characters (which represents a single character with that ASCII code).

Disregarding the whitespace in the file, what is the number of characters of code for string literals minus the number of characters in memory for the values of the strings in total for the entire file?

For example, given the four strings above, the total number of characters of string code (2 + 5 + 10 + 6 = 23) minus the total number of characters in memory for string values (0 + 3 + 7 + 1 = 11) is 23 - 11 = 12.

Your puzzle answer was 1342.

--- Part Two ---
Now, let's go the other way. In addition to finding the number of characters of code, you should now encode each code representation as a new string and find the number of characters of the new encoded representation, including the surrounding double quotes.

For example:

"" encodes to "\"\"", an increase from 2 characters to 6.
"abc" encodes to "\"abc\"", an increase from 5 characters to 9.
"aaa\"aaa" encodes to "\"aaa\\\"aaa\"", an increase from 10 characters to 16.
"\x27" encodes to "\"\\x27\"", an increase from 6 characters to 11.
Your task is to find the total number of characters to represent the newly encoded strings minus the number of characters of code in each original string literal. For example, for the strings above, the total encoded length (6 + 9 + 16 + 11 = 42) minus the characters in the original code representation (23, just like in the first part of this puzzle) is 42 - 23 = 19.

Your puzzle answer was 2074.
*/

func Day8() {
	fmt.Println("--- Day 8: Matchsticks ---")
	sumEsc, sumUnesc, err := SumEscapedCharactersMinusUnescapedCharacters(day8Input)
	if err != nil {
		fmt.Println(errors.Wrap(err, "failed to sum"))
		return

	}
	fmt.Println("Sum of characters in escaped strings minus sum of characters in unescaped strings when unescaping:", sumUnesc)
	fmt.Println("Sum of characters in escaped strings minus sum of characters in unescaped strings when escaping:", sumEsc)
}

// SumEscapedCharactersMinusUnescapedCharacters both escapes and unescape each line of the input
// and returns the sum of escaped characters minus the sum of unescaped characters for both
// respectively.
// SumEscapedCharactersMinusUnescapedCharacters returns an error if the unescape of a given line
// fails.
func SumEscapedCharactersMinusUnescapedCharacters(strs string) (int, int, error) {
	sumEsc := 0
	sumUnesc := 0
	for _, str := range strings.Split(strs, "\n") {
		sumEsc += len(Escape(str)) - len(str)

		unescStr, err := Unescape(str)
		if err != nil {
			return 0, 0, errors.Wrapf(err, "failed to unescape string (str=%s)", str)
		}
		sumUnesc += len(str) - len(unescStr)
	}

	return sumEsc, sumUnesc, nil
}

// Unescape takes an escaped string and returns the unescaped version.
// Unescape will return an error if the escaped strig is malformed.
func Unescape(s string) (string, error) {
	// Unescape iterates through the input string and builds up a []byte with the unescaped output
	res := make([]byte, 0, len(s))

	if len(s) < 2 || s[0] != '"' || s[len(s)-1] != '"' {
		return "", errors.Errorf("first and last character in string must be a quote (s=%s)", s)
	}

	i := 1
	for i < len(s)-1 {
		switch s[i] {
		case '"':
			return "", errors.Errorf("unescaped quote (pos=%d, s=%s)", i, s)
		case '\\':
			// i+1 is guaranteed to exist and safe to read due to the surrounding quotation marks
			switch {
			case s[i+1] == '\\' || (s[i+1] == '"' && i+1 != len(s)-1):
				res = append(res, s[i+1])
				i += 2
			case s[i+1] == 'x':
				if len(s) < i+3 {
					return "", errors.Errorf("malformed hex sequence (pos=%d, s=%s)", i, s)
				}
				v, err := strconv.ParseInt(s[i+2:i+4], 16, 16)
				if err != nil {
					return "", errors.Wrapf(err, "malformed hex sequence (pos=%d, s=%s)", i, s)
				}
				res = append(res, byte(v))
				i += 4
			default:
				return "", errors.Errorf("dangling backslash (pos=%d, s=%s)", i, s)
			}
		default:
			res = append(res, s[i])
			i++
		}
	}

	return string(res), nil
}

// Escape will take a string and escape it
func Escape(s string) string {
	res := make([]byte, 0, len(s)+2) // len(s)+2 will be the minimum size of the result
	res = append(res, '"')
	for i := 0; i < len(s); i++ {
		if s[i] == '"' || s[i] == '\\' {
			res = append(res, '\\')
		}
		res = append(res, s[i])
	}
	res = append(res, '"')
	return string(res)
}
