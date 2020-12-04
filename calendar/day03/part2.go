package day03

import (
	"AdventOfCode2020/utils/files"
	"AdventOfCode2020/utils/slices"
)

func Part2() int {
	input := files.ReadFile("calendar/day03/input", "\n")
	area := slices.ToTwoDimArray(input, "")
	slopes := [5][2]int{ {1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2}}

	multipliedNumberOfTrees := calcNumberOfTreesForSlope(area, slopes[0][0], slopes[0][1])
	for _, slope := range slopes[1:] {
		multipliedNumberOfTrees *= calcNumberOfTreesForSlope(area, slope[0], slope[1])
	}

	return multipliedNumberOfTrees
}


func calcNumberOfTreesForSlope(area [][]string, right int, down int) (numberOfTrees int) {

	areaWidth := len(area[0])

	for x, y := right, down; y < len(area); x, y= (x+right) % areaWidth, y + down {
		if area[y][x] == "#" {
			numberOfTrees++
		}
	}

	return
}



