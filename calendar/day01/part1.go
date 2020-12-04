package day01

import (
	"AdventOfCode2020/utils/conv"
	"AdventOfCode2020/utils/files"
)

func Part1() int {

	numbers := files.ReadFile("calendar/day01/input", "\n")
	expenseReport := conv.ToIntSlice(numbers)

	for indexA, expenseA := range expenseReport {
		for indexB, expenseB := range expenseReport {
			if indexA == indexB {
				continue
			}

			if expenseA+expenseB == 2020 {
				return expenseA * expenseB
			}
		}
	}

	return -1
}
