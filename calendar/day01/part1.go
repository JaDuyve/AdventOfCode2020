package main

import (
	"AdventOfCode2020/utils/conv"
	"AdventOfCode2020/utils/files"
	"fmt"
)



func main() {

	numbers := files.ReadFile("input", "\n")
	expenseReport := conv.ToIntSlice(numbers)

J: 	for indexA, expenseA := range expenseReport {
		for indexB, expenseB := range expenseReport {
			if indexA == indexB {
				continue
			}

			if expenseA + expenseB == 2020 {
				fmt.Printf("%d + %d = %d\n%d * %d = %d",
					expenseA,
					expenseB,
					expenseA+expenseB,
					expenseA,
					expenseB,
					expenseA*expenseB)
				break J
			}
		}
	}
}
