package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateRecipe(t *testing.T) {
	ings, err := NewIngredients(`Butterscotch: capacity -1, durability -2, flavor 6, texture 3, calories 8
Cinnamon: capacity 2, durability 3, flavor -2, texture -1, calories 3`)
	if assert.NoError(t, err) {
		assert.Equal(t, 62842880, CreateRecipe(100, ings, -1))
	}
}

func TestCreateRecipeMax500Calories(t *testing.T) {
	ings, err := NewIngredients(`Butterscotch: capacity -1, durability -2, flavor 6, texture 3, calories 8
Cinnamon: capacity 2, durability 3, flavor -2, texture -1, calories 3`)
	if assert.NoError(t, err) {
		assert.Equal(t, 57600000, CreateRecipe(100, ings, 500))
	}
}

func TestDay15Pt1(t *testing.T) {
	ings, err := NewIngredients(day15Input)
	if assert.NoError(t, err) {
		assert.Equal(t, 13882464, CreateRecipe(100, ings, -1))
	}
}

func TestDay15Pt2(t *testing.T) {
	ings, err := NewIngredients(day15Input)
	if assert.NoError(t, err) {
		assert.Equal(t, 11171160, CreateRecipe(100, ings, 500))
	}
}
