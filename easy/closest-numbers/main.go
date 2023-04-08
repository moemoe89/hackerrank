// https://www.hackerrank.com/challenges/closest-numbers/problem?isFullScreen=true
package main

import (
	"fmt"
)

func closestNumbers(arr []int32) []int32 {
	heapsort(arr)

	minDiff := int32(0)

	for i := 0; i < len(arr)-1; i++ {
		d := abs(arr[i+1] - arr[i])

		if i == 0 {
			minDiff = d

			continue
		}

		if d < minDiff {
			minDiff = d
		}
	}

	output := make([]int32, 0)

	for i := 0; i < len(arr)-1; i++ {
		d := abs(arr[i+1] - arr[i])

		if d == minDiff {
			output = append(output, []int32{arr[i], arr[i+1]}...)
		}
	}

	return output
}

func closestNumbers2(arr []int32) []int32 {
	// ascending sort.
	heapsort(arr)

	// Initialize variables to store the minimum absolute difference and the pairs of closest numbers
	minDiff := int32(0)

	var output []int32

	// Find the minimum absolute difference and the pairs of closest numbers
	for i := 0; i < len(arr)-1; i++ {
		diff := arr[i+1] - arr[i]

		if i == 0 {
			minDiff = diff
		}

		if diff < minDiff {
			minDiff = diff
			output = []int32{arr[i], arr[i+1]}
		} else if diff == minDiff {
			output = append(output, arr[i], arr[i+1])
		}
	}

	return output
}

func abs(n int32) int32 {
	if n < 0 {
		return -n
	}

	return n
}

func heapify(arr []int32, n, i int32) {
	largest := i
	l := 2*i + 1
	r := 2*i + 2

	if l < n && arr[l] > arr[largest] {
		largest = l
	}

	if r < n && arr[r] > arr[largest] {
		largest = r
	}

	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		heapify(arr, n, largest)
	}
}

func heapsort(arr []int32) {
	n := int32(len(arr))

	for i := n/2 - 1; i >= 0; i-- {
		heapify(arr, n, i)
	}

	for i := n - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		heapify(arr, i, 0)
	}
}

func main() {
	fmt.Println(closestNumbers([]int32{5, 2, 3, 4, 1}))
}
