package day04

import (
	"AdventOfCode2020/utils/aocstring"
	"AdventOfCode2020/utils/files"
	"AdventOfCode2020/utils/slices"
	"strings"
)

func Part1() int {
	input := files.ReadFile("calendar/day04/input", "\n\n")

	validPassports := slices.StringFilter(input, isPassportValid)

	return len(validPassports)
}

func isPassportValid(passportString string) bool {
	fields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

	passportString = strings.ReplaceAll(passportString, "\n", " ")
	passport := aocstring.ExtractFieldOfString(strings.Split(passportString, " "), ":")

	for _, field := range fields {
		if _, ok := passport[field]; ok == false {
			return false
		}
	}

	return true
}

