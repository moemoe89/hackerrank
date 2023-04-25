// https://www.hackerrank.com/challenges/missing-numbers/problem?isFullScreen=true
package main

import "fmt"

func missingNumbers(arr []int32, brr []int32) []int32 {
	var out []int32

	m := make(map[int32]int)

	for i := range arr {
		m[arr[i]]++
	}

	for i := range brr {
		m[brr[i]]--
	}

	for k, v := range m {
		if v == 0 {
			continue
		}

		out = append(out, k)
	}

	quickSort(out, 0, len(out)-1)

	return out
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
	fmt.Println(missingNumbers(
		[]int32{11, 4, 11, 7, 13, 4, 12, 11, 10, 14},
		[]int32{11, 4, 11, 7, 3, 7, 10, 13, 4, 8, 12, 11, 10, 14, 12},
	),
	)
}
