package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

func init() { days["15"] = Day15 }

/*
--- Day 15: Science for Hungry People ---
Today, you set out on the task of perfecting your milk-dunking cookie recipe. All you have to do is find the right balance of ingredients.

Your recipe leaves room for exactly 100 teaspoons of ingredients. You make a list of the remaining ingredients you could use to finish the recipe (your puzzle input) and their properties per teaspoon:

capacity (how well it helps the cookie absorb milk)
durability (how well it keeps the cookie intact when full of milk)
flavor (how tasty it makes the cookie)
texture (how it improves the feel of the cookie)
calories (how many calories it adds to the cookie)
You can only measure ingredients in whole-teaspoon amounts accurately, and you have to be accurate so you can reproduce your results in the future. The total score of a cookie can be found by adding up each of the properties (negative totals become 0) and then multiplying together everything except calories.

For instance, suppose you have these two ingredients:

Butterscotch: capacity -1, durability -2, flavor 6, texture 3, calories 8
Cinnamon: capacity 2, durability 3, flavor -2, texture -1, calories 3
Then, choosing to use 44 teaspoons of butterscotch and 56 teaspoons of cinnamon (because the amounts of each ingredient must add up to 100) would result in a cookie with the following properties:

A capacity of 44*-1 + 56*2 = 68
A durability of 44*-2 + 56*3 = 80
A flavor of 44*6 + 56*-2 = 152
A texture of 44*3 + 56*-1 = 76
Multiplying these together (68 * 80 * 152 * 76, ignoring calories for now) results in a total score of 62842880, which happens to be the best score possible given these ingredients. If any properties had produced a negative total, it would have instead become zero, causing the whole score to multiply to zero.

Given the ingredients in your kitchen and their properties, what is the total score of the highest-scoring cookie you can make?

Your puzzle answer was 13882464.

--- Part Two ---
Your cookie recipe becomes wildly popular! Someone asks if you can make another recipe that has exactly 500 calories per cookie (so they can use it as a meal replacement). Keep the rest of your award-winning process the same (100 teaspoons, same ingredients, same scoring system).

For example, given the ingredients above, if you had instead selected 40 teaspoons of butterscotch and 60 teaspoons of cinnamon (which still adds to 100), the total calorie count would be 40*8 + 60*3 = 500. The total score would go down, though: only 57600000, the best you can do in such trying circumstances.

Given the ingredients in your kitchen and their properties, what is the total score of the highest-scoring cookie you can make with a calorie total of 500?

Your puzzle answer was 11171160.
*/

func Day15() {
	fmt.Println("--- Day 15: Science for Hungry People ---")
	ings, err := NewIngredients(day15Input)
	if err != nil {
		fmt.Println(errors.Wrap(err, "failed to load ingredients"))
		return
	}
	fmt.Println("Optimal recipe has a property score of:", CreateRecipe(100, ings, -1))
	fmt.Println("Optimal recipe with exactly 500 calories has a property score of:", CreateRecipe(100, ings, 500))
}

// NewIngredients takes the puzzle input and returns a slice containing the ingredients.
// NewIngredients will return an error if the input is malformed.
func NewIngredients(input string) ([]Ingredient, error) {
	var ingredients []Ingredient
	ingExp := regexp.MustCompile(`^([A-Za-z]+): capacity (-?[0-9]+), durability (-?[0-9]+), flavor (-?[0-9]+), texture (-?[0-9]+), calories ([0-9]+)$`)
	for _, line := range strings.Split(input, "\n") {
		m := ingExp.FindStringSubmatch(line)
		if len(m) != 7 {
			return nil, errors.Errorf("line did not match ingredient expression (line=%s)", line)
		}

		toIntOrPanic := func(property, value string) int {
			i, err := strconv.Atoi(value)
			if err != nil {
				// this should never happen
				panic(fmt.Sprintf("unable to convert value from string to int (property=%s, value=%s)", property, value))
			}
			return i
		}

		ingredients = append(ingredients, Ingredient{
			Name:       m[1],
			Capacity:   toIntOrPanic("capacity", m[2]),
			Durability: toIntOrPanic("durability", m[3]),
			Flavor:     toIntOrPanic("flavor", m[4]),
			Texture:    toIntOrPanic("texture", m[5]),
			Calories:   toIntOrPanic("calories", m[6]),
		})
	}
	return ingredients, nil
}

// Ingredient contains the properties of a specific ingredient.
type Ingredient struct {
	Name       string
	Capacity   int
	Durability int
	Flavor     int
	Texture    int
	Calories   int
}

// CreateRecipe takes the number of teaspoons to use for the recipe along with the possible
// ingredients creates a recipe that maximize the combined properties of the ingredients and
// returns the properties (except for calories) multiplied together.
// If exactCalories is set to a non-negative value, only recipes with that exact calory count
// will be considered.
func CreateRecipe(teaspoons int, ings []Ingredient, exactCalories int) int {
	if len(ings) <= 0 {
		// no ingredients means nothing to do
		return 0
	}

	recipe := make([]int, len(ings))
	return createRecipeRecursive(ings, recipe, teaspoons, 0, exactCalories)
}

// createRecipeRecursive is a recursive helper function that allow CreateRecipe to try out all
// possible combinations of the ingredients to find the optimal one.
func createRecipeRecursive(ings []Ingredient, recipe []int, spLeft int, ingIdx int, exactCalories int) int {
	if ingIdx == len(ings)-1 {
		// nothing left to try
		recipe[ingIdx] = spLeft
		return bake(ings, recipe, exactCalories)
	}

	// try all possible number of teaspoons of the current ingredient
	best := 0
	for sp := 0; sp <= spLeft; sp++ {
		recipe[ingIdx] = sp
		res := createRecipeRecursive(ings, recipe, spLeft-sp, ingIdx+1, exactCalories)
		if res > best {
			best = res
		}
	}
	return best
}

// bake is a helper function that takes the ingredients and a recipe and calculates the value of
// the ingredient properties multiplied together. If any of the properties end up being negative
// it is set to zero thus causing the entire multiplication to return as 0 regardles of what the
// other properties contribute with.
// If exactCalories is set to a non negative value and the total calories count must match eactly
// or the result return will be 0 as well.
func bake(ings []Ingredient, recipe []int, exactCalories int) int {
	capacity, durability, flavor, texture, calories := 0, 0, 0, 0, 0
	for i, ts := range recipe {
		capacity += ings[i].Capacity * ts
		durability += ings[i].Durability * ts
		flavor += ings[i].Flavor * ts
		texture += ings[i].Texture * ts
		calories += ings[i].Calories * ts
	}

	if capacity <= 0 || durability <= 0 || flavor <= 0 || texture <= 0 {
		return 0
	}

	if exactCalories >= 0 && calories != exactCalories {
		return 0
	}

	return capacity * durability * flavor * texture
}
