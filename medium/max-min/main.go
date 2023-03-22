package main

import "fmt"

func maxMin(k int32, arr []int32) int32 {
	if len(arr) == 0 || len(arr) == 1 {
		return int32(0)
	}

	n := len(arr) - 1

	quickSort(arr, 0, n)

	output := arr[k-1] - arr[0]

	for i := 1; i <= n-int(k)+1; i++ {
		tmp := arr[int32(i)+k-1] - arr[i]

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
