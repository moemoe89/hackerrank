// https://www.hackerrank.com/challenges/circular-array-rotation/problem?isFullScreen=true
package main

import "fmt"

func circularArrayRotation(a []int32, k int32, queries []int32) []int32 {
	// calculate the length of a.
	n := int32(len(a))

	// if k > n, we can simplify the number of rotation
	// by doing module, for the example k = 5, n = 4
	// it not necessary to do sliding 5 times because it's equal
	// with 5 % 4 which is 1,
	// here is the example (5):
	// [1, 2, 3, 4] -> [4, 1, 2, 3] -> [3, 4, 1, 2] -> [2, 3, 4, 1] -> [1, 2, 3, 4] -> [4, 1, 2, 3]
	// equal with (1):
	// [1, 2, 3, 4] -> [4, 1, 2, 3]
	k = k % n

	// sliding the slice to right based on k
	a = append(a[n-k:], a[:n-k]...)

	output := make([]int32, len(queries))

	for i, q := range queries {
		output[i] = a[q]
	}

	return output
}

func main() {
	fmt.Println(circularArrayRotation([]int32{3, 4, 5}, 2, []int32{1, 2}))
}
