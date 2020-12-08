package main

import (
	"AdventOfCode2020/calendar/day01"
	"AdventOfCode2020/calendar/day02"
	"AdventOfCode2020/calendar/day03"
	"AdventOfCode2020/calendar/day04"
	"AdventOfCode2020/calendar/day05"
	"AdventOfCode2020/calendar/day06"
	"AdventOfCode2020/calendar/day07"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Day01_Part1(t *testing.T) {
	solution := 805731
	assert.Equalf(t, solution, day01.Part1(), "Solution day 1 should be [%d]", solution)
}

func Test_Day01_Part2(t *testing.T) {
	solution := 192684960
	assert.Equalf(t, solution, day01.Part2(), "Solution day 1 should be [%d]", solution)
}

func Test_Day02_Part1(t *testing.T) {
	solution := 538
	assert.Equalf(t, solution, day02.Part1(), "Solution day 2 should be [%d]", solution)
}

func Test_Day02_Part2(t *testing.T) {
	solution := 489
	assert.Equalf(t, solution, day02.Part2(), "Solution day 2 should be [%d]", solution)
}

func Test_Day03_Part1(t *testing.T) {
	solution := 240
	assert.Equalf(t, solution, day03.Part1(), "Solution day 3 should be [%d]", solution)
}

func Test_Day03_Part2(t *testing.T) {
	solution := 2832009600
	assert.Equalf(t, solution, day03.Part2(), "Solution day 3 should be [%d]", solution)
}

func Test_Day04_Part1(t *testing.T) {
	solution := 247
	assert.Equalf(t, solution, day04.Part1(), "Solution day 3 should be [%d]", solution)
}

func Test_Day04_Part2(t *testing.T) {
	solution := 145
	assert.Equalf(t, solution, day04.Part2(), "Solution day 4 should be [%d]", solution)
}

func Test_Day05_Part1(t *testing.T) {
	solution := 951
	assert.Equalf(t, solution, day05.Part1(), "Solution day 4 should be [%d]", solution)
}

func Test_Day05_Part2(t *testing.T) {
	solution := 653
	assert.Equalf(t, solution, day05.Part2(), "Solution day 5 should be [%d]", solution)
}

func Test_Day06_Part1(t *testing.T) {
	solution := 6585
	assert.Equalf(t, solution, day06.Part1(), "Solution day 7 should be [%d]", solution)
}

func Test_Day06_Part2(t *testing.T) {
	solution := 3276
	assert.Equalf(t, solution, day06.Part2(), "Solution day 6 should be [%d]", solution)
}

func Test_Day07_Part1(t *testing.T) {
	solution := 142
	assert.Equalf(t, solution, day07.Part1(), "Solution day 7 should be [%d]", solution)
}

func Test_Day07_Part2(t *testing.T) {
	solution := 10219
	assert.Equalf(t, solution, day07.Part2(), "Solution day 7 should be [%d]", solution)
}
