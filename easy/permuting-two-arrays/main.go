// https://www.hackerrank.com/challenges/two-arrays/problem?isFullScreen=true
package main

import "fmt"

func twoArrays(k int32, A []int32, B []int32) string {
	n := len(A) - 1

	quickSort(A, 0, n)
	quickSort(B, 0, n)

	for i := range A {
		if A[i]+B[n-i] < k {
			return "NO"
		}
	}

	return "YES"
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
	fmt.Println(twoArrays(10, []int32{2, 1, 3}, []int32{7, 8, 9}))
}
