// https://www.hackerrank.com/challenges/chocolate-feast/problem?isFullScreen=true
package main

import "fmt"

func chocolateFeast(n int32, c int32, m int32) int32 {
	// Count the bar.
	bar := n / c

	// Count the wrapper, basically same with bar.
	wrapper := bar

	// Check if wrapper more than equal with m, which the divider to get a free bar.
	for wrapper >= m {
		// Count free bar from wrapper.
		barFromWrapper := wrapper / m

		// Updates wrapper with the remainder wrapper and wrapper from free bar.
		wrapper = wrapper%m + barFromWrapper

		// Add number of current bar with free bar.
		bar += barFromWrapper
	}

	return bar
}

func main() {
	fmt.Println(chocolateFeast(15, 3, 2))
}
