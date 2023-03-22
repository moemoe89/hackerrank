package main

import "fmt"

func hurdleRace(k int32, height []int32) int32 {
	n := len(height) - 1

	quickSort(height, 0, n)

	maxNumb := height[n]
	if k > maxNumb {
		return 0
	}

	return maxNumb - k
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
	fmt.Println(hurdleRace(1, []int32{1, 2, 3, 3, 2}))
}
