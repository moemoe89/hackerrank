// https://www.hackerrank.com/challenges/jumping-on-the-clouds-revisited/problem?isFullScreen=true
package main

import "fmt"

func jumpingOnClouds(c []int32, k int32) int32 {
	// initialize energy value
	energy := int32(100)

	// initialize length of c
	n := int32(len(c))

	// iterates based on k and stop when i = 0
	// e.g. k = 2, then i will be 2, 4, 6, 8, etc
	for i := k; i != 0; i += k {
		// if i more than equal with n, then need to go back
		// the beginning, need to minus it with length
		// e.g. n = 8, k = 2, when i reach 8, it will go back to 0
		// because there's no index 8.
		if i >= n {
			i -= n
		}

		// every jump will decrease energy by 1
		energy -= 1

		// if current position is thunderheads,
		// add more 2 to decrease the energy
		if c[i] == 1 {
			energy -= 2
		}

		// if the position back to 0, finish the jump
		if i == 0 {
			break
		}
	}

	return energy
}

func main() {
	fmt.Println(jumpingOnClouds([]int32{0, 0, 1, 0, 0, 1, 1, 0}, 2))
}
