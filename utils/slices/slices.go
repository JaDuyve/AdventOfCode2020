package slices

import (
	"fmt"
	"strings"
)

func StringFilter(slice []string, filter func(string) bool) []string {
	resSlice := make([]string, 0)

	for _, element := range slice {
		if filter(element) {
			resSlice = append(resSlice, element)
		}
	}

	return resSlice
}

func IntFilter(slice []int, filter func(int) bool) []int {
	resSlice := make([]int, 0)

	for _, element := range slice {
		if filter(element) {
			resSlice = append(resSlice, element)
		}
	}

	return resSlice
}

func ToTwoDimArray(slice []string, delimiter string) [][]string {
	retSlice := make([][]string, 0, len(slice))

	for _, str := range slice {
		retSlice = append(retSlice, strings.Split(str, delimiter))
	}

	return retSlice
}

func Print2DimStringArray(slice [][]string) {
	var result string

	for _, str := range slice {
		result += fmt.Sprintf("%v\n", str)
	}

	print(result)
}

func Print2DimIntArray(slice [][]int) {
	var result string

	for _, str := range slice {
		result += fmt.Sprintf("%v\n", str)
	}

	print(result)
}

func SumSlice(slice []int) int {
	sum := 0

	for _, number := range slice {
		sum += number
	}

	return sum
}

func Min(slice []int) int {
	min := slice[0]

	for _, number := range slice {
		if number < min {
			min = number
		}
	}

	return min
}

func Max(slice []int) int {
	max := slice[0]

	for _, number := range slice {
		if number > max {
			max = number
		}
	}

	return max
}