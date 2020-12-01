package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func sliceAtoi(sa []string) ([]int, error) {
	si := make([]int, 0, len(sa))
	for _, a := range sa {
		i, err := strconv.Atoi(a)
		if err != nil {
			return si, err
		}
		si = append(si, i)
	}
	return si, nil
}

func main() {
	data, err := ioutil.ReadFile("input")
	if err != nil {
		panic(err)
	}

	numbers := strings.Split(string(data), "\n")
	expenseReport, err := sliceAtoi(numbers)

	if err != nil {
		panic(err)
	}

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
