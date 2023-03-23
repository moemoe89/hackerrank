package main

import "fmt"

func marcsCakewalk(calorie []int32) int64 {
	n := len(calorie) - 1

	quickSort(calorie, 0, n)

	output := int64(0)

	for i := range calorie {
		output += int64(calorie[n-i]) * pow(2, int64(i))
	}

	return output
}

func pow(a, b int64) int64 {
	result := int64(1)

	for i := int64(0); i < b; i++ {
		result *= a
	}

	return result
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
	fmt.Println(marcsCakewalk([]int32{1, 3, 2}))
}
