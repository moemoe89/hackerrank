// https://www.hackerrank.com/challenges/maximizing-xor/problem?isFullScreen=true
package main

import "fmt"

func maximizingXor(l int32, r int32) int32 {
	// initialize max variable.
	max := int32(0)

	// because we need each pair, do nested loop.
	for i := l; i <= r; i++ {
		for j := l; j <= r; j++ {
			// checks the ⊕ or xor by using ^
			if i^j > max {
				max = i ^ j
			}
		}
	}

	return max
}

func main() {
	fmt.Println(maximizingXor(11, 100))
}
