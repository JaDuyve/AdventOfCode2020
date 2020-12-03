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

func Print2DimArray(slice [][]string) {
	var result string

	for _, str := range slice {
		result += fmt.Sprintf("%v\n", str)
	}

	print(result)
}