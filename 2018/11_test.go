package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPowerLevel(t *testing.T) {
	testCases := []struct {
		X, Y, Serial int
		Out          int
	}{
		{3, 5, 8, 4},
		{122, 79, 57, -5},
		{217, 196, 39, 0},
		{101, 153, 71, 4},
	}

	for i, c := range testCases {
		assert.Equal(t, c.Out, PowerLevel(c.Serial, c.X, c.Y), "(case=%d)", i)
	}
}

func TestFindFuelCell(t *testing.T) {
	testCases := []struct {
		Size, Serial     int
		OutX, OutY, OutP int
	}{
		{3, 18, 33, 45, 29},
		{3, 42, 21, 61, 30},
	}

	for i, c := range testCases {
		x, y, p := FindFuelCell(c.Size, c.Serial)
		assert.Equal(t, c.OutX, x, "X (case=%d)", i)
		assert.Equal(t, c.OutY, y, "Y (case=%d)", i)
		assert.Equal(t, c.OutP, p, "P (case=%d)", i)
	}
}

func TestFindAnyFuelCell(t *testing.T) {
	testCases := []struct {
		Serial                 int
		OutX, OutY, OutS, OutP int
	}{
		{18, 90, 269, 16, 113},
		{42, 232, 251, 12, 119},
	}

	for i, c := range testCases {
		x, y, s, p := FindAnyFuelCell(c.Serial)
		assert.Equal(t, c.OutX, x, "X (case=%d)", i)
		assert.Equal(t, c.OutY, y, "Y (case=%d)", i)
		assert.Equal(t, c.OutS, s, "S (case=%d)", i)
		assert.Equal(t, c.OutP, p, "P (case=%d)", i)
	}
}
