package slices

import (
	"fmt"
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

func ToTwoDimArray(slice []string) [][]rune {
	retSlice := make([][]rune, 0, len(slice))

	for _, str := range slice {
		retSlice = append(retSlice, []rune(str))
	}

	return retSlice
}

func Print2DimRuneArray(slice [][]rune) {
	var result string

	for _, str := range slice {
		result += fmt.Sprintf("%c\n", str)
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