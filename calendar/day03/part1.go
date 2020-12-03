package main

import (
	"AdventOfCode2020/utils/files"
	"AdventOfCode2020/utils/slices"
	"fmt"
)

func main() {
    input := files.ReadFile("input", "\n")
    area := slices.ToTwoDimArray(input, "")

    numberOfTrees := calcNumberOfTrees(&area)

    fmt.Printf("came across [%d] trees during the journey.\n", numberOfTrees)
    slices.Print2DimArray(area)
}

func calcNumberOfTrees(area *[][]string) (numberOfTrees int) {

	areaWidth := len((*area)[0])

	for x, y := 3, 1; y < len(*area); x, y= (x+3) % areaWidth, y + 1{
		if (*area)[y][x] == "#" {
			numberOfTrees++
			(*area)[y][x] = "X"
		} else {
			(*area)[y][x] = "O"
		}
	}

	return
}



