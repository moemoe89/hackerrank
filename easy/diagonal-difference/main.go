package main

import "fmt"

func diagonalDifference(arr [][]int32) int32 {
	left := int32(0)
	right := int32(0)

	for i := range arr {
		left += arr[i][i]
		right += arr[i][len(arr)-1-i]
	}

	if left < right {
		return right - left
	}

	return left - right
}

func main() {
	fmt.Println(diagonalDifference([][]int32{
		{11, 2, 4},
		{4, 5, 6},
		{10, 8, -12},
	}))
}
