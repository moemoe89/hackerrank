package main

import "fmt"

func migratoryBirds(arr []int32) int32 {
	max := 0

	m := make(map[int32]int)

	for _, a := range arr {
		m[a]++

		if m[a] > max {
			max = m[a]
		}
	}

	var mostNums []int32

	for k, v := range m {
		if v == max {
			mostNums = append(mostNums, k)
		}
	}

	quickSort(mostNums, 0, len(mostNums)-1)

	return mostNums[0]
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
	fmt.Println(migratoryBirds([]int32{1, 4, 4, 4, 5, 3}))
}
