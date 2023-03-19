package main

import "fmt"

func maximumToys(prices []int32, k int32) int32 {
	left := 0
	right := len(prices) - 1

	quickSort(prices, left, right)

	var output, sum int32

	for i := range prices {
		sum += prices[i]

		if sum >= k {
			break
		}

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
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}

	arr[i], arr[right] = arr[right], arr[i]

	return i
}

func main() {
	fmt.Println(maximumToys([]int32{1, 2, 3, 4}, int32(7)))
}
