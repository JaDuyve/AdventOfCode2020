package main

import (
	"AdventOfCode2020/utils/files"
	"AdventOfCode2020/utils/slices"
	"fmt"
)

func main() {
	input := files.ReadFile("input", "\n")
	validPasswords := slices.StringFilter(input, isPasswordValid)

	fmt.Printf("There are [%d] valid passwords.\n", len(validPasswords))
}

func isPasswordValid(entry string) bool {
	var min, max int
	var character uint8
	var password string

	_, err := fmt.Sscanf(entry, "%d-%d %c: %s", &min, &max, &character, &password)
	if err != nil {
		panic(err)
	}

	return (password[min - 1] == character) != (password[max - 1] == character)
}