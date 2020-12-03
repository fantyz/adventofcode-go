package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidatePassword(t *testing.T) {
	testCases := map[string]struct {
		Min, Max int
		Seq      string
		Pass     string
		ExpValid bool
	}{
		"example 1": {1, 3, "a", "abcde", true},
		"example 2": {1, 3, "b", "cdefg", false},
		"example 3": {2, 9, "c", "ccccccccc", true},
	}

	for n, c := range testCases {
		valid := ValidatePassword(c.Min, c.Max, c.Seq, c.Pass)
		if c.ExpValid {
			assert.True(t, valid, n)
		} else {
			assert.False(t, valid, n)
		}
	}
}

func TestValidatePassword2(t *testing.T) {
	testCases := map[string]struct {
		Min, Max int
		Seq      string
		Pass     string
		ExpValid bool
	}{
		"example 1": {1, 3, "a", "abcde", true},
		"example 2": {1, 3, "b", "cdefg", false},
		"example 3": {2, 9, "c", "ccccccccc", false},
	}

	for n, c := range testCases {
		valid := ValidatePassword2(c.Min, c.Max, c.Seq, c.Pass)
		if c.ExpValid {
			assert.True(t, valid, n)
		} else {
			assert.False(t, valid, n)
		}
	}
}

func TestValidatePasswordFile(t *testing.T) {
	testCase := `1-3 a: abcde
1-3 b: cdefg
2-9 c: ccccccccc`

	assert.Equal(t, 2, ValidatePasswordFile(0, testCase), "Policy 0")
	assert.Equal(t, 1, ValidatePasswordFile(1, testCase), "Policy 1")
}

func TestDay2Pt1(t *testing.T) {
	assert.Equal(t, 416, ValidatePasswordFile(0, day2Input))
}

func TestDay2Pt2(t *testing.T) {
	assert.Equal(t, 688, ValidatePasswordFile(1, day2Input))
}
