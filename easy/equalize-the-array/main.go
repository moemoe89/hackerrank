package main

import "fmt"

func equalizeArray(arr []int32) int32 {
	// initialize map
	m := make(map[int32]int, len(arr))

	max := 0

	// put the element as key
	// and increment the value
	// it's mean if the element is more than 1,
	// the value will accumulate
	for _, a := range arr {
		m[a]++

		// check if max value
		// lower than current key value
		// then replace it
		if m[a] > max {
			max = m[a]
		}
	}

	// to get the number of deletion,
	// simple deduct the length with max value
	// e.g. [3, 3, 2, 1, 3], max = 3, length = 5
	// deletion will be 2
	// if the max number duplicated, it should keep as it
	// e.g. [3, 3, 2, 2, 1], max = 2, length = 5
	// deletion will be 1 only
	return int32(len(arr) - max)
}

func main() {
	fmt.Println(equalizeArray([]int32{3, 3, 2, 1, 3}))
}
