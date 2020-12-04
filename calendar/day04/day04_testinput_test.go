package day04

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_SolvePart2_ValidInput(t *testing.T) {
	input := "pid:087499704 hgt:74in ecl:grn iyr:2012 eyr:2030 byr:1980\nhcl:#623a2f\n\n" +
		"eyr:2029 ecl:blu cid:129 byr:1989\niyr:2014 pid:896056539 hcl:#a97842 hgt:165cm\n\n" +
		"hcl:#888785\nhgt:164cm byr:2001 iyr:2015 cid:88\npid:545766238 ecl:hzl\neyr:2022\n\n" +
		"iyr:2010 hgt:158cm hcl:#b6652a ecl:blu byr:1944 eyr:2021 pid:093154719"
	assert.Equal(t, 4, solvePart2(input), "Answer to the puzzle should be 4.")
}

func Test_SolvePart2_InvalidInput(t *testing.T) {
	input := "eyr:1972 cid:100\nhcl:#18171d ecl:amb hgt:170 pid:186cm iyr:2018 byr:1926\n\n" +
		"iyr:2019\nhcl:#602927 eyr:1967 hgt:170cm\necl:grn pid:012533040 byr:1946\n\n" +
		"hcl:dab227 iyr:2012\necl:brn hgt:182cm pid:021572410 eyr:2020 byr:1992 cid:277\n\n" +
		"hgt:59cm ecl:zzz\neyr:2038 hcl:74454a iyr:2023\npid:3556412378 byr:2007"
	assert.Zerof(t, solvePart2(input), "Answer to the puzzle should be 0.")
}
