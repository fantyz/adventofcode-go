package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScanCorruptedMemoryForMuls(t *testing.T) {
	assert.Equal(t, 161, ScanCorruptedMemoryForMuls([]string{`xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))`}))
}

func TestScanCorruptedMemoryForMulsUsingDoRegions(t *testing.T) {
	assert.Equal(t, 48, ScanCorruptedMemoryForMuls(ScanCorruptedMemoryForDoRegions(`xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))`)))
}

func TestDay03Pt1(t *testing.T) {
	assert.Equal(t, 159892596, ScanCorruptedMemoryForMuls([]string{day03Input}))
}

func TestDay03Pt2(t *testing.T) {
	assert.Equal(t, 92626942, ScanCorruptedMemoryForMuls(ScanCorruptedMemoryForDoRegions(day03Input)))
}
