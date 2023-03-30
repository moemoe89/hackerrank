package main

import "fmt"

func reverseArray(a []int32) []int32 {
	// initialize left and right index.
	left := 0
	right := len(a) - 1

	// loop until left index with right index
	for left <= right {
		// swap the first and last index
		// and then moving one by one
		// e.g. 1, 2, 3, 4, 5
		// 5, 2, 3, 4, 1
		// 5, 4, 3, 2, 1
		a[left], a[right] = a[right], a[left]

		// increment and decrement to moving the index position
		left++
		right--
	}

	return a
}

func main() {
	fmt.Println(reverseArray([]int32{1, 2, 3, 4, 5}))
}
