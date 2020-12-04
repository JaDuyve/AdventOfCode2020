package main

import (
	"AdventOfCode2020/calendar/day01"
	"AdventOfCode2020/calendar/day02"
	"AdventOfCode2020/calendar/day03"
	"AdventOfCode2020/calendar/day04"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Day01_Part1(t *testing.T) {
	solution := 805731
	assert.Equal(t, solution, day01.Part1(), fmt.Sprintf("Solution day 1 should be [%d]", solution))
}

func Test_Day01_Part2(t *testing.T) {
	solution := 192684960
	assert.Equal(t, solution, day01.Part2(), fmt.Sprintf("Solution day 1 should be [%d]", solution))
}

func Test_Day02_Part1(t *testing.T) {
	solution := 538
	assert.Equal(t, solution, day02.Part1(), fmt.Sprintf("Solution day 1 should be [%d]", solution))
}

func Test_Day02_Part2(t *testing.T) {
	solution := 489
	assert.Equal(t, solution, day02.Part2(), fmt.Sprintf("Solution day 1 should be [%d]", solution))
}

func Test_Day03_Part1(t *testing.T) {
	solution := 240
	assert.Equal(t, solution, day03.Part1(), fmt.Sprintf("Solution day 1 should be [%d]", solution))
}

func Test_Day03_Part2(t *testing.T) {
	solution := 2832009600
	assert.Equal(t, solution, day03.Part2(), fmt.Sprintf("Solution day 1 should be [%d]", solution))
}

func Test_Day04_Part1(t *testing.T) {
	solution := 247
	assert.Equal(t, solution, day04.Part1(), fmt.Sprintf("Solution day 1 should be [%d]", solution))
}

func Test_Day04_Part2(t *testing.T) {
	solution := 145
	assert.Equal(t, solution, day04.Part2(), fmt.Sprintf("Solution day 1 should be [%d]", solution))
}

