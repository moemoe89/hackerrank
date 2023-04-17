// https://www.hackerrank.com/challenges/jim-and-the-orders/problem?isFullScreen=true
package main

import "fmt"

func jimOrders(orders [][]int32) []int32 {
	var serves [][]int32

	for i := range orders {
		serves = append(serves, []int32{orders[i][0] + orders[i][1], int32(i) + 1})
	}

	quickSort(serves, 0, len(serves)-1)

	var o []int32

	for i := range serves {
		o = append(o, serves[i][1])
	}

	return o
}

func quickSort(arr [][]int32, left int, right int) {
	if left < right {
		pivotIndex := partition(arr, left, right)
		quickSort(arr, left, pivotIndex-1)
		quickSort(arr, pivotIndex+1, right)
	}
}

func partition(arr [][]int32, left int, right int) int {
	pivot := arr[right][0]

	i := left

	for j := left; j < right; j++ {
		// < operator means ascending sort
		// > operator means descending sort
		if arr[j][0] < pivot || (arr[j][0] == pivot && arr[j][1] < arr[right][1]) {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}

	arr[i], arr[right] = arr[right], arr[i]

	return i
}

func main() {
	fmt.Println(jimOrders([][]int32{
		{1, 1},
		{1, 1},
	}))
}
