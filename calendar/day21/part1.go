package day21

import (
	"AdventOfCode2020/utils/files"
	"AdventOfCode2020/utils/slices"
	"go/types"
	"strings"
)

type Food struct {
	ingredients []string
	allergens   map[string]types.Nil
}

func Part1() int {
	input := files.ReadFile("calendar/day21/input")
	return solvePart1(input)

}

func solvePart1(puzzleInput string) int {
	foods, allergens := parseInput(puzzleInput)

	return countIngredientsWithoutAllergens(foods, &allergens)
}

func parseInput(input string) ([]Food, map[string]string) {
	foods := make([]Food, 0)

	//var foodRegex = regexp.MustCompile(`(?m)([\W\S]+) \(contains ([\W,\S]+)\)`)

	for _, str := range strings.Split(input, "\n") {
		split := strings.Split(str, " (contains ")
		allergens := make(map[string]types.Nil)
		for _, allergen := range strings.Split(split[1][:len(split[1])-1], ", ") {
			allergens[allergen] = types.Nil{}
		}

		foods = append(foods, Food{
			ingredients: strings.Split(split[0], " "),
			allergens:   allergens,
		})
	}

	return foods, extractAllAllergens(foods)
}

func extractAllAllergens(foods []Food) map[string]string {
	allergens := make(map[string]string)

	for _, food := range foods {
		for allergen, _ := range food.allergens {
			if _, exist := allergens[allergen]; !exist {
				allergens[allergen] = ""
			}
		}
	}

	return allergens
}

func countIngredientsWithoutAllergens(foods []Food, allergens *map[string]string) int {
	ingredients := getAllIngredients(foods)
	connectedIngredients := make([]string, 0)

	for !foundAllTuples(*allergens) {
		for allergen, _ := range *allergens {
			if len((*allergens)[allergen]) != 0 {
				continue
			}

			filteredFood := getAllFoodsForAllergen(allergen, foods)
			intersection := findIntersection(filteredFood)

			if len(intersection) != 1 {

				for _, ingredient := range connectedIngredients {
					for i := len(intersection) - 1; i >= 0; i-- {
						if intersection[i] == ingredient {
							intersection = slices.RemoveChangeOrder(intersection, i)
							continue
						}
					}
				}
			}

			if len(intersection) == 1 {
				foundIngredient := intersection[0]

				connectedIngredients = append(connectedIngredients, foundIngredient)
				delete(ingredients, foundIngredient)
				(*allergens)[allergen] = foundIngredient
			}

		}
	}

	return countIngredientEncounters(ingredients, foods)
}

func getAllFoodsForAllergen(allergen string, foods []Food) []Food {
	res := make([]Food, 0)

	for _, food := range foods {
		if _, exist := food.allergens[allergen]; exist {
			res = append(res, food)
		}
	}

	return res
}

func getAllIngredients(foods []Food) map[string]types.Nil {
	ingredients := make(map[string]types.Nil)

	for _, food := range foods {
		for _, ingredient := range food.ingredients {
			ingredients[ingredient] = types.Nil{}
		}
	}

	return ingredients
}

func findIntersection(foods []Food) []string {
	freq := make(map[string]int)

	for _, food := range foods {
		for _, ingredient := range food.ingredients {
			freq[ingredient]++
		}
	}

	res := make([]string, 0)

	for key, value := range freq {
		if value == len(foods) {
			res = append(res, key)
		}
	}

	return res
}

func foundAllTuples(allergens map[string]string) bool {

	for _, value := range allergens {
		if value == "" {
			return false
		}
	}

	return true
}

func countIngredientEncounters(ingredients map[string]types.Nil, foods []Food) int {
	counter := 0

	for ingredientKey, _ := range ingredients {
		for _, food := range foods {
			for _, ingredient := range food.ingredients {
				if ingredientKey == ingredient {
					counter++
				}
			}
		}
	}

	return counter
}