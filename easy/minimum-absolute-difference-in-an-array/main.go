package main

import "fmt"

func minimumAbsoluteDifference(arr []int32) int32 {
	// initialize the length of arr
	n := len(arr) - 1

	// sort the array in ascending order.
	quickSort(arr, 0, n)

	// initialize the minimum absolute difference as 0
	minDiff := abs(arr[1] - arr[0])

	// find the minimum absolute difference by comparing adjacent elements in the sorted array.
	for i := 2; i <= n; i++ {
		diff := abs(arr[i] - arr[i-1])

		// if it's the first index on iteration, assign the diff to minDiff.
		if i == 1 {
			minDiff = diff
		}

		// check if the diff lower than minDiff, then replace it.
		if diff < minDiff {
			minDiff = diff
		}
	}

	return minDiff
}

func abs(n int32) int32 {
	if n < 0 {
		return -n
	}

	return n
}

func quickSort(arr []int32, left int, right int) {
	if left < right {
		pivotIndex := partition(arr, left, right)
		quickSort(arr, left, pivotIndex-1)
		quickSort(arr, pivotIndex+1, right)
	}
}

func partition(arr []int32, left int, right int) int {
	pivot := arr[right]

	i := left

	for j := left; j < right; j++ {
		// < operator means ascending sort
		// > operator means descending sort
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}

	arr[i], arr[right] = arr[right], arr[i]

	return i
}

func main() {
	fmt.Println(minimumAbsoluteDifference([]int32{-59, -36, -13, 1, -53, -92, -2, -96, -54, 75}))
}
