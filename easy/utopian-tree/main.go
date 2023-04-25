// https://www.hackerrank.com/challenges/utopian-tree/problem?isFullScreen=true
package main

import "fmt"

func utopianTree(n int32) int32 {
	var h int32

	// The Utopian Tree goes through 2 cycles of growth every year.
	// Each spring, it doubles in height. Each summer, its height increases by 1 meter.
	// In the example, spring represent as odd number, other than that means summer (even number).
	// Period  Height
	// 0          1  -> even, + 1
	// 1          2  -> odd, * 2
	// 2          3  -> even, + 1
	// 3          6  -> odd, * 2
	// 4          7  -> even, + 1
	// 5          14 -> odd, * 2
	for i := int32(0); i <= n; i++ {
		if i%2 == 0 {
			h = h + 1
		} else {
			h = h * 2
		}
	}

	return h
}

func main() {
	fmt.Println(utopianTree(5))
}
