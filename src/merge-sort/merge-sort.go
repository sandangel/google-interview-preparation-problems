package main

import (
	"fmt"
	"math"
)

func main() {
	arr := []int{10, 24, 76, 73, 72, 1, 9}

	fmt.Println(mergeSort(arr))
}

func merge(arr1, arr2 []int) []int {
	var results []int
	i, j := 0, 0
	for i < len(arr1) && j < len(arr2) {
		if arr2[j] > arr1[i] {
			results = append(results, arr1[i])
			i++
		} else {
			results = append(results, arr2[j])
			j++
		}
	}

	for i < len(arr1) {
		results = append(results, arr1[i])
		i++
	}

	for j < len(arr2) {
		results = append(results, arr2[j])
		j++
	}

	return results
}

func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	mid := math.Floor(float64(len(arr)) / 2)
	left := mergeSort(arr[:int(mid)])
	right := mergeSort(arr[int(mid):])
	return merge(left, right)
}
