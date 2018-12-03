package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChronalCalibration(t *testing.T) {
	testCases := []struct {
		Claims     string
		ExpOverlap int
		ExpClaims  []Claim
	}{
		{
			Claims: `
#1 @ 1,3: 4x4
#2 @ 3,1: 4x4
#3 @ 5,5: 2x2
`,
			ExpOverlap: 4,
			ExpClaims:  []Claim{{"#3", 5, 5, 2, 2}},
		},
	}

	for i, c := range testCases {
		fabric := NewFabric(10, 10)
		claims := NewClaims(c.Claims)
		fabric.AddClaims(claims)
		fabric.Print()
		assert.Equal(t, c.ExpOverlap, fabric.Overlap(), "(case=%d)", i)
		assert.Equal(t, c.ExpClaims, fabric.FindNonOverlappedClaims(claims), "(case=%d)", i)
	}
}
