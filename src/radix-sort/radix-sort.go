package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(radixSort([]int{1, 3, 77, 23, 55, 93, 208, 702, 304, 105, 1025, 5269, 4231}, 10))
}

func getDigit(num, i int) int {
	return int(math.Floor(math.Abs(float64(num))/math.Pow(10, float64(i)))) % 10
}

func digitCount(num int) int {
	return len(fmt.Sprintf("%d", int(math.Abs(float64(num)))))
}

func mostDigits(nums []int) (max int) {
	for _, n := range nums {
		max = int(math.Max(float64(digitCount(n)), float64(max)))
	}
	return
}

func radixSort(nums []int, base int) []int {
	maxDigit := mostDigits(nums)
	sortedNums := nums

	for i := 0; i < maxDigit; i++ {
		buckets := make([][]int, base)
		for _, n := range sortedNums {
			buckets[getDigit(n, i)] = append(buckets[getDigit(n, i)], n)
		}
		sortedNums = []int{}
		for i := 0; i < base; i++ {
			sortedNums = append(sortedNums, buckets[i]...)
		}
	}
	return sortedNums
}
