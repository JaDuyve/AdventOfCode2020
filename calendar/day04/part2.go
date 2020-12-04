package day04

import (
	"AdventOfCode2020/utils/aocstring"
	"AdventOfCode2020/utils/conv"
	"AdventOfCode2020/utils/files"
	"AdventOfCode2020/utils/slices"
	"fmt"
	"regexp"
	"strings"
)

func Part2() int {
	input := files.ReadFile("calendar/day04/input", "\n\n")

	validPassports := slices.StringFilter(input, isPassportValidWithConstraints)

	return len(validPassports)
}

func isPassportValidWithConstraints(passportString string) bool {
	fields := map[string]func(string) bool{
		"byr": isValidBirthYear,
		"iyr": isValidIssueYear,
		"eyr": isValidExpirationYear,
		"hgt": isValidHeight,
		"hcl": isValidHairColor,
		"ecl": isValidEyeColor,
		"pid": isValidPassportId,
	}

	passportString = strings.ReplaceAll(passportString, "\n", " ")
	passport := aocstring.ExtractFieldOfString(strings.Split(passportString, " "), ":")

	for field, validationFunction := range fields {
		value, ok := passport[field]

		if ok == false {
			return false
		}

		if !validationFunction(value) {
			return false
		}
	}

	return true
}

func isValidBirthYear(year string) bool {
	return isValidYear(year, 1920, 2002)
}

func isValidIssueYear(year string) bool {
	return isValidYear(year, 2010, 2020)
}

func isValidExpirationYear(year string) bool {
	return isValidYear(year, 2020, 2030)
}

func isValidYear(yearString string, start int, end int) bool {
	year := conv.ToInt(yearString)
	return year >= start && year <= end
}

func isValidEyeColor(color string) bool {
	match, _ := regexp.MatchString("^(amb|blu|brn|gry|grn|hzl|oth)$", color)
	return match
}

func isValidHairColor(color string) bool {
	match, _ := regexp.MatchString("^#[0-9a-f]{6}$", color)
	return match
}

func isValidPassportId(str string) bool {
	match, _ := regexp.MatchString("^[0-9]{9}$", str)
	return match
}

func isValidHeight(str string) bool {

	switch {
	case strings.Contains(str, "cm"):
		var height int

		_, err := fmt.Sscanf(str, "%dcm", &height)

		if err != nil {
			panic(err)
		}

		if height >= 150 && height <= 193 {
			return true
		}
	case strings.Contains(str, "in"):
		var height int
		_, err := fmt.Sscanf(str, "%din", &height)

		if err != nil {
			panic(err)
		}

		if height >= 59 && height <= 76 {
			return true
		}
	default:
		return false
	}

	return false
}
