package main

import (
	"AdventOfCode2020/utils/conv"
	"AdventOfCode2020/utils/files"
	"fmt"
)

func main() {
	numbers := files.ReadFile("input", "\n")
	expenseReport := conv.ToIntSlice(numbers)

J:
	for indexA, expenseA := range expenseReport {
		for indexB, expenseB := range expenseReport {
			for indexC, expenseC := range expenseReport {
				if indexA == indexB || indexB == indexC || indexA == indexC {
					continue
				}

				if expenseA+expenseB+expenseC == 2020 {
					fmt.Printf("%d + %d + %d = %d\n%d * %d * %d = %d",
						expenseA,
						expenseB,
						expenseC,
						expenseA+expenseB+expenseC,
						expenseA,
						expenseB,
						expenseC,
						expenseA*expenseB*expenseC)
					break J
				}
			}
		}
	}
}
