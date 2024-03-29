// https://www.hackerrank.com/challenges/jumping-on-the-clouds/problem?isFullScreen=true
package main

import "fmt"

func jumpingOnClouds(c []int32) int32 {
	n := len(c) - 1

	var jump int32

	for i := 0; i <= n; i++ {
		if i == n {
			break
		}

		twoJump := i + 2
		if twoJump <= n && c[twoJump] == 0 {
			i++

			jump++

			continue
		}

		jump++
	}

	return jump
}

func main() {
	fmt.Println(jumpingOnClouds([]int32{0, 0, 1, 0, 0, 1, 0}))
}
