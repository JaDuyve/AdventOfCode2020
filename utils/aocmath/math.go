package aocmath

import "sort"

func SumSlice(slice []int) int {
	sum := 0

	for _, number := range slice {
		sum += number
	}

	return sum
}

func Min(slice []int) int {
	min := slice[0]

	for _, number := range slice {
		if number < min {
			min = number
		}
	}

	return min
}

func Max(slice []int) int {
	max := slice[0]

	for _, number := range slice {
		if number > max {
			max = number
		}
	}

	return max
}

func ThreeSum(slice []int, target int) (result [][]int) {
	sort.Ints(slice)

	for indexA := 0; indexA < len(slice); indexA++ {
		a := slice[indexA]

		if a > target {
			return
		}
		if indexA > 0 && a == slice[indexA-1] {
			continue
		}

		indexB := indexA + 1
		indexC := len(slice) - 1

		for indexB < indexC {
			b := slice[indexB]
			c := slice[indexC]
			sum := a + b + c

			switch {
			case sum == target:
				result = append(result, []int{a, b, c})

				// ignore duplicates
				for slice[indexB] == slice[indexB-1] {
					indexB++
				}

				for slice[indexC] == slice[indexC+1] {
					indexC--
				}

				indexB++
				indexC--
			case sum < target:
				indexB++
			default:
				indexC--
			}
		}
	}

	return
}

func threeSumBruteForce(slice []int, target int) (result [][]int) {
	sort.Ints(slice)
	sliceLength := len(slice)

	for indexA := 0; indexA < sliceLength; indexA++ {
		a := slice[indexA]

		if indexA > 0 && a == slice[indexA-1] {
			continue
		}

		for indexB := indexA + 1; indexB < sliceLength; indexB++ {
			b := slice[indexB]

			if indexB > 0 && b == slice[indexB-1] {
				continue
			}

			for indexC := indexB + 1; indexC < sliceLength; indexC++ {
				c := slice[indexC]

				if indexC > 0 && c == slice[indexC-1] {
					continue
				}

				if a+b+c == target {
					result = append(result, []int{a, b, c})
				}
			}
		}
	}

	return
}
