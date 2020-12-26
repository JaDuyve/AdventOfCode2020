package day21

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_SolvePart1_TestInput(t *testing.T) {
	input := "mxmxvkd kfcds sqjhc nhms (contains dairy, fish)\n" +
		"trh fvjkl sbzzf mxmxvkd (contains dairy)\n" +
		"sqjhc fvjkl (contains soy)\n" +
		"sqjhc mxmxvkd sbzzf (contains fish)"
	expected := 5
	assert.Equal(t, expected, solvePart1(input))
}

func Test_SolvePart2_TestInput(t *testing.T) {
	input := "mxmxvkd kfcds sqjhc nhms (contains dairy, fish)\n" +
		"trh fvjkl sbzzf mxmxvkd (contains dairy)\n" +
		"sqjhc fvjkl (contains soy)\n" +
		"sqjhc mxmxvkd sbzzf (contains fish)"
	expected := "mxmxvkd,sqjhc,fvjkl"

	foods, allergens := parseInput(input)
	countIngredientsWithoutAllergens(foods, &allergens)
	assert.Equal(t, expected, generateCanonicalDangerousList(allergens))
}