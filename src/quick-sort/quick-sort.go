package main

import "fmt"

func main() {
	arr := []int{4, 8, 2, 1, 5, 7, 6, 3}
	fmt.Println(quickSort(arr, 0, len(arr)-1))
}

func pivot(arr []int, start, end int) int {
	swapIdx := start
	pivot := arr[start]

	for i := start + 1; i <= end; i++ {
		if arr[i] < pivot {
			swapIdx++
			arr[swapIdx], arr[i] = arr[i], arr[swapIdx]
		}
	}

	arr[start], arr[swapIdx] = arr[swapIdx], arr[start]

	return swapIdx
}

func quickSort(arr []int, left, right int) []int {
	if left < right {
		pivotIdx := pivot(arr, left, right)
		quickSort(arr, left, pivotIdx-1)
		quickSort(arr, pivotIdx+1, right)
	}

	return arr
}
