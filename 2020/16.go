package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

func init() { days["16"] = Day16 }

/*
--- Day 16: Ticket Translation ---
As you're walking to yet another connecting flight, you realize that one of the legs of your re-routed trip coming up is on a high-speed train. However, the train ticket you were given is in a language you don't understand. You should probably figure out what it says before you get to the train station after the next flight.

Unfortunately, you can't actually read the words on the ticket. You can, however, read the numbers, and so you figure out the fields these tickets must have and the valid ranges for values in those fields.

You collect the rules for ticket fields, the numbers on your ticket, and the numbers on other nearby tickets for the same train service (via the airport security cameras) together into a single document you can reference (your puzzle input).

The rules for ticket fields specify a list of fields that exist somewhere on the ticket and the valid ranges of values for each field. For example, a rule like class: 1-3 or 5-7 means that one of the fields in every ticket is named class and can be any value in the ranges 1-3 or 5-7 (inclusive, such that 3 and 5 are both valid in this field, but 4 is not).

Each ticket is represented by a single line of comma-separated values. The values are the numbers on the ticket in the order they appear; every ticket has the same format. For example, consider this ticket:

.--------------------------------------------------------.
| ????: 101    ?????: 102   ??????????: 103     ???: 104 |
|                                                        |
| ??: 301  ??: 302             ???????: 303      ??????? |
| ??: 401  ??: 402           ???? ????: 403    ????????? |
'--------------------------------------------------------'
Here, ? represents text in a language you don't understand. This ticket might be represented as 101,102,103,104,301,302,303,401,402,403; of course, the actual train tickets you're looking at are much more complicated. In any case, you've extracted just the numbers in such a way that the first number is always the same specific field, the second number is always a different specific field, and so on - you just don't know what each position actually means!

Start by determining which tickets are completely invalid; these are tickets that contain values which aren't valid for any field. Ignore your ticket for now.

For example, suppose you have the following notes:

class: 1-3 or 5-7
row: 6-11 or 33-44
seat: 13-40 or 45-50

your ticket:
7,1,14

nearby tickets:
7,3,47
40,4,50
55,2,20
38,6,12
It doesn't matter which position corresponds to which field; you can identify invalid nearby tickets by considering only whether tickets contain values that are not valid for any field. In this example, the values on the first nearby ticket are all valid for at least one field. This is not true of the other three nearby tickets: the values 4, 55, and 12 are are not valid for any field. Adding together all of the invalid values produces your ticket scanning error rate: 4 + 55 + 12 = 71.

Consider the validity of the nearby tickets you scanned. What is your ticket scanning error rate?

Your puzzle answer was 24110.

--- Part Two ---
Now that you've identified which tickets contain invalid values, discard those tickets entirely. Use the remaining valid tickets to determine which field is which.

Using the valid ranges for each field, determine what order the fields appear on the tickets. The order is consistent between all tickets: if seat is the third field, it is the third field on every ticket, including your ticket.

For example, suppose you have the following notes:

class: 0-1 or 4-19
row: 0-5 or 8-19
seat: 0-13 or 16-19

your ticket:
11,12,13

nearby tickets:
3,9,18
15,1,5
5,14,9
Based on the nearby tickets in the above example, the first position must be row, the second position must be class, and the third position must be seat; you can conclude that in your ticket, class is 12, row is 11, and seat is 13.

Once you work out which field is which, look for the six fields on your ticket that start with the word departure. What do you get if you multiply those six values together?

Your puzzle answer was 6766503490793.
*/

func Day16() {
	fmt.Println("--- Day 16: Ticket Translation ---")

	rules, err := NewTicketRules(day16InputRules)
	if err != nil {
		fmt.Println(err)
		return
	}
	tickets, err := LoadTickets(day16InputTickets)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Ticket Scanning Error Rate:", TicketScanningErrorRate(rules, tickets[1:]))

	fields, err := TicketFields(rules, tickets[1:])
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Multiplied departure fields:", MultiplyDepartureFields(fields, tickets[0]))
}

// MultiplyDepartureFields multiplies all fields in the provided ticket together that
// starts with the word "departure".
func MultiplyDepartureFields(fields []string, ticket []int) int {
	const fieldPrefix = "departure"
	total := 1
	for idx, field := range fields {
		if !strings.HasPrefix(field, fieldPrefix) {
			continue
		}
		total *= ticket[idx]
	}
	return total
}

// TicketFields takes the rules and tickets and determine based on these what the order
// of fields on the tickets are. TicketFields will return an error if no order can be
// determined.
//
// NOTE: TicketFields only work when there is a deterministic solution for the provided
// tickets (eg. only one choice of column for each field).
func TicketFields(rules TicketRules, tickets [][]int) ([]string, error) {
	// remove invalid tickets
	var validTickets [][]int // use a new slice to avoid causing side effects
	for _, ticket := range tickets {
		valid := true
		for _, field := range ticket {
			if !rules.MatchAny(field) {
				valid = false
				break
			}
		}
		if valid {
			validTickets = append(validTickets, ticket)
		}
	}

	if len(validTickets) <= 0 {
		return nil, errors.New("no valid tickets")
	}

	// figure out what fields are valid for which columns
	fieldValidFor := map[string][]bool{}
	for _, field := range rules.Fields() {
		// find a row in the ticket data that all match the field
		validFor := make([]bool, len(validTickets[0]))
		for i := range validFor {
			validFor[i] = true
		}

		for i := range validTickets {
			for j := range validTickets[i] {
				if !rules.Match(field, validTickets[i][j]) {
					validFor[j] = false
				}
			}
		}
		fieldValidFor[field] = validFor
	}

	// figure out what fields maps to what columns
	fieldsRemaining := len(fieldValidFor)
	fields := make([]string, fieldsRemaining)

	// at each iteration until we are done there should be one field that has exactly one validFor index that is true.
	// this could fairly easily be optimized by sorting instead of brute force iteration
	for i := 0; i < fieldsRemaining; i++ {
		for field, validFor := range fieldValidFor {
			validCount := 0
			idx := -1
			for j, valid := range validFor {
				if valid {
					validCount++
					idx = j
				}
			}
			if validCount == 1 {
				// found column for field
				fields[idx] = field

				// remove field from fieldValidFor and set remaining fields to false for idx
				delete(fieldValidFor, field)
				for _, v := range fieldValidFor {
					v[idx] = false
				}
				break
			}
		}
	}

	if len(fieldValidFor) > 0 {
		return nil, errors.New("unable to match all fields with all columns")
	}

	return fields, nil
}

// TicketScanningErrorRate takes the rules and tickets and return the scanning error rate.
func TicketScanningErrorRate(rules TicketRules, tickets [][]int) int {
	errorSum := 0
	for _, ticket := range tickets {
		for _, field := range ticket {
			if !rules.MatchAny(field) {
				errorSum += field
			}
		}
	}
	return errorSum
}

// NewTicketRules takes the raw rules and return TicketRules. NewTicketRules returns
// an error if the raw rules are invalid.
func NewTicketRules(rules string) (TicketRules, error) {
	r := TicketRules{}
	for _, rule := range strings.Split(rules, "\n") {
		m := strings.SplitN(rule, ": ", 2)
		if len(m) != 2 {
			return nil, errors.Errorf("invalid rule (rule=%s)", rule)
		}

		vr, err := NewValueRange(m[1])
		if err != nil {
			return nil, errors.Wrapf(err, "invalid rule value range (rule=%s)", rule)
		}
		r[m[0]] = vr
	}
	return r, nil
}

type TicketRules map[string]ValueRange

// MatchAny will return true if i is a valid number for any ticket rule.
func (rules TicketRules) MatchAny(i int) bool {
	for _, r := range rules {
		if r.IsWithin(i) {
			return true
		}
	}
	return false
}

// Fields will return the list of fields that has rules associated with them.
func (rules TicketRules) Fields() []string {
	fields := make([]string, 0, len(rules))
	for f, _ := range rules {
		fields = append(fields, f)
	}
	return fields
}

// Match will return true if i is a valid nuber for the specified field rule. Match will
// return false if the specified field does not have a rule.
func (rules TicketRules) Match(field string, i int) bool {
	r, found := rules[field]
	if !found {
		return false
	}
	return r.IsWithin(i)
}

// NewValueRange takes a raw value range and returns a ValueRange representing it.
// NewValueRange returns an error if the raw value is invalid.
func NewValueRange(raw string) (ValueRange, error) {
	rangeExp := regexp.MustCompile(`^([0-9]+)-([0-9]+) or ([0-9]+)-([0-9]+)$`)

	m := rangeExp.FindStringSubmatch(raw)
	if len(m) != 5 {
		return ValueRange{}, errors.Errorf("value range did not match as expected (raw=%s, len(m)=%d)", raw, len(m))
	}

	var vals [4]int
	for i := 0; i < 4; i++ {
		n, err := strconv.Atoi(m[i+1])
		if err != nil {
			return ValueRange{}, errors.Wrapf(err, "bad range value (val=%s, raw=%s)", m[i+1], raw)
		}
		vals[i] = n
	}

	return ValueRange{
		LowStart:  vals[0],
		LowEnd:    vals[1],
		HighStart: vals[2],
		HighEnd:   vals[3],
	}, nil
}

type ValueRange struct {
	LowStart, LowEnd   int
	HighStart, HighEnd int
}

// IsWithin returns true if i falls within the value range.
func (r ValueRange) IsWithin(i int) bool {
	return (i >= r.LowStart && i <= r.LowEnd) || (i >= r.HighStart && i <= r.HighEnd)
}

// LoadTickets loads the raw ticket data and returns it. An error is returned
// if the ticket data is invalid.
func LoadTickets(raw string) ([][]int, error) {
	var tickets [][]int

	for _, ticket := range strings.Split(raw, "\n") {
		values, err := LoadInts(ticket, ",")
		if err != nil {
			return nil, errors.Wrapf(err, "bad ticket (ticket=%s)", ticket)
		}
		tickets = append(tickets, values)
	}

	return tickets, nil
}
