package day21

import (
	"AdventOfCode2020/utils/files"
	"fmt"
	"sort"
)

func Part2() int {
	input := files.ReadFile("calendar/day21/input")
	return solvePart2(input)
}

func solvePart2(puzzleInput string) int {
	foods, allergens := parseInput(puzzleInput)

	countIngredientsWithoutAllergens(foods, &allergens)
	fmt.Printf("Result day 21 part 2: [%s]\n", generateCanonicalDangerousList(allergens))
	return -1
}

func generateCanonicalDangerousList(allergens map[string]string) string {

	keys := make([]string, 0, len(allergens))
	for key, _ := range allergens {
		keys = append(keys, key)
	}

	sort.Strings(keys)
	canonicalDangerousList := allergens[keys[0]]
	for _, key := range keys[1:] {
		canonicalDangerousList += "," + allergens[key]
	}

	return canonicalDangerousList
}
