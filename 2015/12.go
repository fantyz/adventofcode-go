package main

import (
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
)

func init() { days["12"] = Day12 }

/*
--- Day 12: JSAbacusFramework.io ---
Santa's Accounting-Elves need help balancing the books after a recent order. Unfortunately, their accounting software uses a peculiar storage format. That's where you come in.

They have a JSON document which contains a variety of things: arrays ([1,2,3]), objects ({"a":1, "b":2}), numbers, and strings. Your first job is to simply find all of the numbers throughout the document and add them together.

For example:

[1,2,3] and {"a":2,"b":4} both have a sum of 6.
[[[3]]] and {"a":{"b":4},"c":-1} both have a sum of 3.
{"a":[-1,1]} and [-1,{"a":1}] both have a sum of 0.
[] and {} both have a sum of 0.
You will not encounter any strings containing numbers.

What is the sum of all numbers in the document?

Your puzzle answer was 111754.

--- Part Two ---
Uh oh - the Accounting-Elves have realized that they double-counted everything red.

Ignore any object (and all of its children) which has any property with the value "red". Do this only for objects ({...}), not arrays ([...]).

[1,2,3] still has a sum of 6.
[1,{"c":"red","b":2},3] now has a sum of 4, because the middle object is ignored.
{"d":"red","e":[1,2,3,4],"f":5} now has a sum of 0, because the entire structure is ignored.
[1,"red",5] has a sum of 6, because "red" in an array has no effect.
Your puzzle answer was 65402.
*/

func Day12() {
	fmt.Println("--- Day 12: JSAbacusFramework.io ---")
	sum, err := SumJSONDocumentNumbers(day12Input, false)
	if err != nil {
		fmt.Println(errors.Wrap(err, "failed to sum document"))
		return
	}
	fmt.Println("The sum of all numbers in the document is:", sum)
	sum2, err := SumJSONDocumentNumbers(day12Input, true)
	if err != nil {
		fmt.Println(errors.Wrap(err, "failed to sum document ignoring red"))
		return
	}
	fmt.Println("The sum of all numbers in the document is (ignoring anything with red properties):", sum2)
}

// SumJSONDocumentNumbers takes a generic json document and sums all numbers within it.
// SumJSONDocumentNumbers returns an error if the input document is not valid json.
func SumJSONDocumentNumbers(doc string, ignoreRed bool) (int, error) {
	var d interface{}
	if err := json.Unmarshal([]byte(doc), &d); err != nil {
		return 0, errors.Wrap(err, "failed to json umarshal input document")
	}
	sum, _ := sumJSONDocumentNumbersRecursively(d, ignoreRed)

	return sum, nil
}

// sumJSONDocumentNumbersRecursively is a recursive helper function used to sum up all values in a
// generic JSON document.
func sumJSONDocumentNumbersRecursively(i interface{}, ignoreRed bool) (int, bool) {
	switch v := i.(type) {
	case float64:
		// number value, return it
		return int(v), false
	case []interface{}:
		// array, sum any values in it
		sum := 0
		for i := range v {
			v, _ := sumJSONDocumentNumbersRecursively(v[i], ignoreRed)
			sum += v
		}
		return sum, false
	case map[string]interface{}:
		// object, sum any values in it
		sum := 0
		for i := range v {
			v, ignore := sumJSONDocumentNumbersRecursively(v[i], ignoreRed)
			if ignore {
				return 0, false
			}
			sum += v
		}
		return sum, false
	case string:
		if ignoreRed && v == "red" {
			return 0, true
		}
	default:
		// ignore any other types
	}
	return 0, false
}
