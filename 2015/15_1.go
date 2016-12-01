package main

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
*/

import (
	"fmt"
	"regexp"
	"strconv"
)

type Ingredient struct {
	Name string

	Capacity, Durability, Flavor, Texture, Calories int
}

type Recipie struct {
	Measurements []int
}

func NewRecipie() Recipie {
	return Recipie{
		Measurements: make([]int, len(ingredients)),
	}
}

func (r *Recipie) AddIngredient(ingredientNum int, teaspoons int) {
	r.Measurements[ingredientNum] = teaspoons
}

func (r Recipie) Copy() Recipie {
	new := Recipie{
		Measurements: make([]int, len(ingredients)),
	}
	copy(new.Measurements, r.Measurements)
	return new
}

func (r *Recipie) Bake() int {
	capacity := 0
	durability := 0
	flavor := 0
	texture := 0
	for i := range r.Measurements {
		capacity += ingredients[i].Capacity * r.Measurements[i]
		durability += ingredients[i].Durability * r.Measurements[i]
		flavor += ingredients[i].Flavor * r.Measurements[i]
		texture += ingredients[i].Texture * r.Measurements[i]
	}

	if capacity < 0 || durability < 0 || flavor < 0 || texture < 0 {
		return 0
	}

	return capacity * durability * flavor * texture
}

var ingredients []*Ingredient

func main() {
	re := regexp.MustCompile("^([a-zA-Z]+): capacity (-?[0-9]+), durability (-?[0-9]+), flavor (-?[0-9]+), texture (-?[0-9]+), calories (-?[0-9]+)$")

	ingredients = make([]*Ingredient, 0)

	for {
		line, found := readNextInputLine()
		if !found {
			break
		}
		matches := re.FindStringSubmatch(line)
		if len(matches) <= 0 {
			panic("input did not match expected format")
		}

		ingredient := &Ingredient{}

		var err error

		ingredient.Name = matches[1]
		ingredient.Capacity, err = strconv.Atoi(matches[2])
		if err != nil {
			panic("could not convert capacity to int")
		}
		ingredient.Durability, err = strconv.Atoi(matches[3])
		if err != nil {
			panic("could not convert durability to int")
		}
		ingredient.Flavor, err = strconv.Atoi(matches[4])
		if err != nil {
			panic("could not convert flavor to int")
		}
		ingredient.Texture, err = strconv.Atoi(matches[5])
		if err != nil {
			panic("could not convert texture to int")
		}
		ingredient.Calories, err = strconv.Atoi(matches[6])
		if err != nil {
			panic("could not convert calories to int")
		}
		ingredients = append(ingredients, ingredient)
	}

	optimalScore := putIngredients(NewRecipie(), 100, 0)
	fmt.Println("Best score:", optimalScore)
}

func putIngredients(recipie Recipie, spoonsLeft int, ingredientNum int) int {
	if ingredientNum >= len(ingredients) {
		panic("invalid ingredientNum")
	}
	if ingredientNum == len(ingredients)-1 {
		// only one combination possible
		recipie.AddIngredient(ingredientNum, spoonsLeft)
		return recipie.Bake()
	}

	best := 0
	for i := 0; i <= spoonsLeft; i++ {
		recipieCopy := recipie.Copy()
		recipieCopy.AddIngredient(ingredientNum, i)
		score := putIngredients(recipieCopy, spoonsLeft-i, ingredientNum+1)
		if score > best {
			best = score
		}
	}
	return best
}

var pos = 0

func readNextInputLine() (string, bool) {
	start := pos
	for i := pos + 1; i <= len(input); i++ {
		if i == len(input) || input[i] == '\n' {
			pos = i + 1
			return input[start : pos-1], true
		}
	}
	return "", false
}

var input_example = `Butterscotch: capacity -1, durability -2, flavor 6, texture 3, calories 8
Cinnamon: capacity 2, durability 3, flavor -2, texture -1, calories 3`

var input = `Sprinkles: capacity 2, durability 0, flavor -2, texture 0, calories 3
Butterscotch: capacity 0, durability 5, flavor -3, texture 0, calories 3
Chocolate: capacity 0, durability 0, flavor 5, texture -1, calories 8
Candy: capacity 0, durability -1, flavor 0, texture 5, calories 8`
