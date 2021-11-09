package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalibrate(t *testing.T) {
	input := `H => HO
H => OH
O => HH

HOH`
	replacer, molecule := NewMoleculeReplacerAndMolecule(input)
	assert.Equal(t, 4, replacer.Calibrate(molecule))
}

func TestDay19Pt1(t *testing.T) {
	replacer, molecule := NewMoleculeReplacerAndMolecule(day19Input)
	assert.Equal(t, 576, replacer.Calibrate(molecule))
}

func TestDay19Pt2(t *testing.T) {
	replacer, molecule := NewMoleculeReplacerAndMolecule(day19Input)
	assert.Equal(t, 207, replacer.FewestStepsTo(molecule))
}
