package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadFoods(t *testing.T) {
	testFoods := `mxmxvkd kfcds sqjhc nhms (contains dairy, fish)
trh fvjkl sbzzf mxmxvkd (contains dairy)
sqjhc fvjkl (contains soy)
sqjhc mxmxvkd sbzzf (contains fish)`

	f, err := LoadFoods(testFoods)
	if assert.NoError(t, err) {
		assert.Equal(t, []Food{
			{[]string{"mxmxvkd", "kfcds", "sqjhc", "nhms"}, []string{"dairy", "fish"}},
			{[]string{"trh", "fvjkl", "sbzzf", "mxmxvkd"}, []string{"dairy"}},
			{[]string{"sqjhc", "fvjkl"}, []string{"soy"}},
			{[]string{"sqjhc", "mxmxvkd", "sbzzf"}, []string{"fish"}},
		}, f)
	}
}

func TestAnalyzeFoods(t *testing.T) {
	testFoods := []Food{
		{[]string{"mxmxvkd", "kfcds", "sqjhc", "nhms"}, []string{"dairy", "fish"}},
		{[]string{"trh", "fvjkl", "sbzzf", "mxmxvkd"}, []string{"dairy"}},
		{[]string{"sqjhc", "fvjkl"}, []string{"soy"}},
		{[]string{"sqjhc", "mxmxvkd", "sbzzf"}, []string{"fish"}},
	}

	appearances, allergens := AnalyzeAllergens(testFoods)
	assert.Equal(t, 5, appearances)
	assert.Equal(t, map[string]string{
		"dairy": "mxmxvkd",
		"fish":  "sqjhc",
		"soy":   "fvjkl",
	}, allergens)
}

func TestDay21(t *testing.T) {
	foods, err := LoadFoods(day21Input)
	if assert.NoError(t, err) {
		appearances, allergens := AnalyzeAllergens(foods)
		assert.Equal(t, 2635, appearances)
		assert.Equal(t, "xncgqbcp,frkmp,qhqs,qnhjhn,dhsnxr,rzrktx,ntflq,lgnhmx", CanonicalDangerousIngredientList(allergens))
	}
}
