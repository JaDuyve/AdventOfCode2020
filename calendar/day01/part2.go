package day01

import (
	"AdventOfCode2020/utils/conv"
	"AdventOfCode2020/utils/files"
)

func Part2() int {
	numbers := files.ReadFile("calendar/day01/input", "\n")
	expenseReport := conv.ToIntSlice(numbers)

	for indexA, expenseA := range expenseReport {
		for indexB, expenseB := range expenseReport {
			for indexC, expenseC := range expenseReport {
				if indexA == indexB || indexB == indexC || indexA == indexC {
					continue
				}

				if expenseA+expenseB+expenseC == 2020 {
					return expenseA * expenseB * expenseC
				}
			}
		}
	}

	return -1
}
