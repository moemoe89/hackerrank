// https://www.hackerrank.com/challenges/priyanka-and-toys/problem?isFullScreen=true
package main

import "fmt"

func toys(w []int32) int32 {
	output := int32(1)

	quickSort(w, 0, len(w)-1)

	max := w[0] + 4

	for i := 1; i < len(w); i++ {
		if w[i] <= max {
			continue
		}

		max = w[i] + 4

		output++
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
	fmt.Println(toys([]int32{1, 2, 3, 4, 5, 10, 11, 12, 13}))
}
