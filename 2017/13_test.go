package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTraverseFirewall(t *testing.T) {
	testCases := []struct {
		In     string
		Result int
	}{
		{`0: 3
1: 2
4: 4
6: 4`, 24},
	}

	for i, testCase := range testCases {
		fw := NewFirewall(testCase.In)
		assert.Equal(t, testCase.Result, TraverseFirewall(fw), "(case=%d)", i)
	}
}

func TestTraverseFirewallWithoutGettingCaught(t *testing.T) {
	testCases := []struct {
		In     string
		Result int
	}{
		{`0: 3
1: 2
4: 4
6: 4`, 10},
	}

	for i, testCase := range testCases {
		fw := NewFirewall(testCase.In)
		assert.Equal(t, testCase.Result, TraverseFirewallWithoutGettingCaught(fw), "(case=%d)", i)
	}
}
