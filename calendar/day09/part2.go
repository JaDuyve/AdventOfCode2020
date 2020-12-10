package day09

import (
	"AdventOfCode2020/utils/aocmath"
	"AdventOfCode2020/utils/conv"
	"AdventOfCode2020/utils/files"
	"strings"
)

func Part2() int {
	input := files.ReadFile("calendar/day09/input")
	return solvePart2(input, 25)

}

func solvePart2(puzzleInput string, preamble int) int {
	cypher := conv.ToIntSlice(strings.Split(puzzleInput, "\n"))
	invalidNumber := findInvalidNumber(cypher, preamble)

    return findWeakness(invalidNumber, cypher, preamble)
}

func findWeakness(keyNumber int, cypher []int, preamble int) int {
	for lengthContinuousSet := 2; lengthContinuousSet <= preamble; lengthContinuousSet++ {
		for i := 0; i < len(cypher)-lengthContinuousSet; i++ {
			continuousSet := cypher[i : i+lengthContinuousSet]
			if aocmath.SumSlice(continuousSet) == keyNumber {
				return aocmath.Min(continuousSet) + aocmath.Max(continuousSet)
			}
		}
	}

	return -1
}