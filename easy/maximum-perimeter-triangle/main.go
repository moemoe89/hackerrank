// https://www.hackerrank.com/challenges/maximum-perimeter-triangle/problem?isFullScreen=true
package main

import "fmt"

// need to use int64
func maximumPerimeterTriangle(sticks []int64) []int64 {
	m := make(map[int64][]int64)

	m[0] = []int64{-1}
	max := int64(0)

	quickSort(sticks, 0, len(sticks)-1)

	for i := 0; i <= len(sticks)-3; i++ {
		if sticks[i]+sticks[i+1] <= sticks[i+2] {
			continue
		}

		sum := sticks[i] + sticks[i+1] + sticks[i+2]

		if sum > max {
			m[sum] = sticks[i : i+3]
			max = sum
		}
	}

	return m[max]
}

func quickSort(arr []int64, left int, right int) {
	if left < right {
		pivotIndex := partition(arr, left, right)
		quickSort(arr, left, pivotIndex-1)
		quickSort(arr, pivotIndex+1, right)
	}
}

func partition(arr []int64, left int, right int) int {
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
	fmt.Println(maximumPerimeterTriangle([]int64{1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000}))
}
