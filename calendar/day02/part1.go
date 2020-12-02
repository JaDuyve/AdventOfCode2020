package main

import (
	"AdventOfCode2020/utils/files"
	"AdventOfCode2020/utils/slices"
	"fmt"
	"strings"
)

func main() {
	input := files.ReadFile("input", "\n")
	validPasswords := slices.StringFilter(input, isPasswordValidCountChar)

	fmt.Printf("There are [%d] valid passwords.\n", len(validPasswords))
}

func isPasswordValidCountChar(entry string) bool {
	var min, max int
	var character, password string

	_, err := fmt.Sscanf(entry, "%d-%d %1s: %s", &min, &max, &character, &password)
	if err != nil {
		panic(err)
	}

	count := strings.Count(password, character)
	return count >= min && count <= max
}
