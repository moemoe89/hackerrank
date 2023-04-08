// https://www.hackerrank.com/challenges/countingsort2/problem?isFullScreen=true
package main

import "fmt"

func countingSort(arr []int32) []int32 {
	quickSort(arr, 0, len(arr)-1)

	return arr
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
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}

	arr[i], arr[right] = arr[right], arr[i]

	return i
}

func main() {
	fmt.Println(countingSort([]int32{1, 1, 3, 2, 1}))
}
