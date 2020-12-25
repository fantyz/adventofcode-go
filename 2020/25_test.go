package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseEngineerLoopSize(t *testing.T) {
	testPubKey := 5764801
	testSubject := 7
	assert.Equal(t, 8, ReverseEngineerLoopSize(testPubKey, testSubject))
}

func TestTransform(t *testing.T) {
	testLoopSize := 8
	testSubject := 17807724
	assert.Equal(t, 14897079, Transform(testLoopSize, testSubject))
}

func TestDay25Pt1(t *testing.T) {
}

func TestDay25Pt2(t *testing.T) {
}
