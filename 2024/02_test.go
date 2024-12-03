package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReportSafe(t *testing.T) {
	testCases := []struct{
		R Report
		UseProblemDampner bool
		ExpSafe bool
	}{
		{Report{7,6,4,2,1}, false, true},
		{Report{1,2,7,8,9}, false, false},
		{Report{9,7,6,2,1}, false, false},
		{Report{1,3,2,4,5}, false, false},
		{Report{8,6,4,4,1}, false, false},
		{Report{1,3,6,7,9}, false, true},
		{Report{7,6,4,2,1}, true, true},
		{Report{1,2,7,8,9}, true, false},
		{Report{9,7,6,2,1}, true, false},
		{Report{1,3,2,4,5}, true, true},
		{Report{8,6,4,4,1}, true, true},
		{Report{1,3,6,7,9}, true, true},
		{Report{9,1,2,3}, false, false},
		{Report{9,1,2,3}, true, true},
	}

	for _, c := range testCases {
		assert.Equal(t, c.ExpSafe, c.R.Safe(c.UseProblemDampner), "in=%v, problemDampner=%v", c.R, c.UseProblemDampner)
	}
}

func TestDay02Pt1(t *testing.T) {
	reports := LoadReports(day02Input)
	assert.Equal(t, 490, CountSafeReports(reports, false))
}

func TestDay02Pt2(t *testing.T) {
	reports := LoadReports(day02Input)
	assert.Equal(t, 536, CountSafeReports(reports, true))
}
