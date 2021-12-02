package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimpleEndPosition(t *testing.T) {
	c := Course{
		{ForwardAction, 5},
		{DownAction, 5},
		{ForwardAction, 8},
		{UpAction, 3},
		{DownAction, 8},
		{ForwardAction, 2},
	}

	pos, dep := c.SimpleEndPosition()
	assert.Equal(t, 15, pos, "position")
	assert.Equal(t, 10, dep, "depth")
}

func TestRealEndPosition(t *testing.T) {
	c := Course{
		{ForwardAction, 5},
		{DownAction, 5},
		{ForwardAction, 8},
		{UpAction, 3},
		{DownAction, 8},
		{ForwardAction, 2},
	}

	pos, dep := c.RealEndPosition()
	assert.Equal(t, 15, pos, "position")
	assert.Equal(t, 60, dep, "depth")
}

func TestDay02Pt1(t *testing.T) {
	c, err := NewCourse(day02Input)
	if assert.NoError(t, err) {
		pos, dep := c.SimpleEndPosition()
		assert.Equal(t, 1882980, pos*dep)
	}
}

func TestDay02Pt2(t *testing.T) {
	c, err := NewCourse(day02Input)
	if assert.NoError(t, err) {
		pos, dep := c.RealEndPosition()
		assert.Equal(t, 1971232560, pos*dep)
	}
}
