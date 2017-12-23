package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewParticle(t *testing.T) {
	testCases := []struct {
		Particle string
	}{
		{`p=<5556,2862,7112>, v=<-6,-118,-35>, a=<-9,2,-10>`},
	}

	for i, testCase := range testCases {
		ps := NewParticles(testCase.Particle)
		if assert.Equal(t, 1, len(ps), "Particle count (case=%d)", i) {
			assert.Equal(t, testCase.Particle, ps[0].String(), "(case=%d)", i)
		}
	}
}

func TestCountCollisions(t *testing.T) {
	testCases := []struct {
		In     string
		Result int
	}{
		{`p=<-6,0,0>, v=<3,0,0>, a=<0,0,0>
p=<-4,0,0>, v=<2,0,0>, a=<0,0,0>
p=<-2,0,0>, v=<1,0,0>, a=<0,0,0>
p=<3,0,0>, v=<-1,0,0>, a=<0,0,0>`, 1},
	}

	for i, testCase := range testCases {
		assert.Equal(t, testCase.Result, CollideParticles(NewParticles(testCase.In)), "(case=%d)", i)
	}
}
