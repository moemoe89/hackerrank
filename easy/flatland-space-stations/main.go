// https://www.hackerrank.com/challenges/flatland-space-stations/problem?isFullScreen=true
package main

import "fmt"

func flatlandSpaceStations(n int32, c []int32) int32 {
	x := len(c) - 1

	quickSort(c, 0, x)

	max := int32(0)

	for i := 0; i < x; i++ {
		diff := (c[i+1] - c[i]) / 2
		if diff > max {
			max = diff
		}
	}

	if c[0] > max {
		max = c[0]
	}

	if n-1-c[x] > max {
		max = n - 1 - c[x]
	}

	return max
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
	fmt.Println(flatlandSpaceStations(5, []int32{0, 4}))
}
