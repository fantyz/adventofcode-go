package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTicketScanningErrorRate(t *testing.T) {
	testRules := `class: 1-3 or 5-7
row: 6-11 or 33-44
seat: 13-40 or 45-50`
	testTickets := `7,1,14
7,3,47
40,4,50
55,2,20
38,6,12`

	rules, rulesErr := NewTicketRules(testRules)
	tickets, ticketsErr := LoadTickets(testTickets)
	if assert.NoError(t, rulesErr, "rules") && assert.NoError(t, ticketsErr, "tickets") {
		assert.Equal(t, 71, TicketScanningErrorRate(rules, tickets[1:]))
	}
}

func TestTicketFields(t *testing.T) {
	testRules := `class: 0-1 or 4-19
row: 0-5 or 8-19
seat: 0-13 or 16-19`
	testTickets := `11,12,13
3,9,18
15,1,5
5,14,9`

	rules, rulesErr := NewTicketRules(testRules)
	tickets, ticketsErr := LoadTickets(testTickets)
	if assert.NoError(t, rulesErr, "rules") && assert.NoError(t, ticketsErr, "tickets") {
		fields, err := TicketFields(rules, tickets[1:])
		if assert.NoError(t, err, "fields") {
			assert.Equal(t, []string{"row", "class", "seat"}, fields)
		}
	}
}

func TestDay16Pt1(t *testing.T) {
	rules, rulesErr := NewTicketRules(day16InputRules)
	tickets, ticketsErr := LoadTickets(day16InputTickets)
	if assert.NoError(t, rulesErr, "rules") && assert.NoError(t, ticketsErr, "tickets") {
		assert.Equal(t, 24110, TicketScanningErrorRate(rules, tickets))
	}
}

func TestDay16Pt2(t *testing.T) {
	rules, rulesErr := NewTicketRules(day16InputRules)
	tickets, ticketsErr := LoadTickets(day16InputTickets)
	if assert.NoError(t, rulesErr, "rules") && assert.NoError(t, ticketsErr, "tickets") {
		fields, err := TicketFields(rules, tickets[1:])
		if assert.NoError(t, err, "fields") {
			assert.Equal(t, 6766503490793, MultiplyDepartureFields(fields, tickets[0]))
		}
	}
}
