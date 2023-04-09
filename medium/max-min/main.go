package main

import "fmt"

func maxMin(k int32, arr []int32) int32 {
	// if the length is under or equal 1, return 0.
	if len(arr) <= 1 {
		return int32(0)
	}

	// initialize the index length
	n := len(arr) - 1

	// sort the array.
	quickSort(arr, 0, n)

	// identify the maximum and minimum values in first subarray.
	output := arr[k-1] - arr[0]

	// iterates to find the other maximum and minimum values.
	// if arr is [1, 2, 3], k = 2, i should stop at i = 1.
	for i := 1; i <= n-int(k)+1; i++ {
		tmp := arr[int32(i)+k-1] - arr[i]

		// if the value lower that output, replace it.
		if tmp < output {
			output = tmp
		}
	}

	return output
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
	fmt.Println(maxMin(4, []int32{1, 2, 3, 4, 10, 20, 30, 40, 100, 200}))
}
