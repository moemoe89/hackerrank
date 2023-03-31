// https://www.hackerrank.com/challenges/array-left-rotation/problem?isFullScreen=true
package main

import "fmt"

func rotateLeft(d int32, arr []int32) []int32 {
	// calculate the length of arr.
	n := int32(len(arr))

	// if d > n, we can simplify the number of rotation
	// by doing module, for the example d = 6, n = 6
	// it not necessary to do sliding 5 times because it's equal
	// with 6 % 5 which is 1,
	// here is the example (5):
	// [1, 2, 3, 4, 5] -> [2, 3, 4, 5, 1] -> [3, 4, 5, 1, 2] ->
	// [4, 5, 1, 2, 3] -> [5, 1, 2, 3, 4] -> [1, 2, 3, 4, 5] -> [2, 3, 4, 5, 1]
	// equal with (1):
	// [1, 2, 3, 4, 5] -> [2, 3, 4, 5, 1]
	d = d % n

	// sliding the slice to left based on d.
	// e.g. [1, 2, 3, 4, 5]
	// a[1:] -> [2, 3, 4, 5]
	// a[:1] -> [1]
	// then combine the array become [2, 3, 4, 5, 1]
	return append(arr[d:], arr[:d]...)
}

func main() {
	fmt.Println(rotateLeft(2, []int32{1, 2, 3, 4, 5}))
}
