// https://www.hackerrank.com/challenges/manasa-and-stones/problem?isFullScreen=true
package main

import "fmt"

func stones(n int32, a int32, b int32) []int32 {
	if a > b {
		a, b = b, a
	}

	m := make(map[int32]struct{})

	for i := int32(0); i < n; i++ {
		max := b * (n - 1 - i)
		min := a * i
		m[max+min] = struct{}{}
	}

	var output []int32
	for key := range m {
		output = append(output, key)
	}

	quickSort(output, 0, len(output)-1)

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
	fmt.Println(stones(73, 25, 25))
}
