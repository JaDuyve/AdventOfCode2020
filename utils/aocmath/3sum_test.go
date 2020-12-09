package aocmath

import (
	"math/rand"
	"testing"
	"time"
)

func BenchmarkTreeSumBruteForce(b *testing.B) {
	size := 10_000
	a := generateRandomArray(size)

	b.ResetTimer()
	b.StartTimer()
	threeSumBruteForce(a, 380)
	b.StopTimer()
}

func BenchmarkTreeSum(b *testing.B) {
	size := 10_000
	a := generateRandomArray(size)

	b.StartTimer()
	result :=ThreeSum(a, 380)
	b.StopTimer()
	print(result)
}

func generateRandomArray(size int) []int {
	array := make([]int, 0, size)

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		array = append(array, rand.Intn(500))
	}

	return array
}

