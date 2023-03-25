package main

import "fmt"

func jumpingOnClouds(c []int32, k int32) int32 {
	energy := int32(100)

	n := int32(len(c))

	for i := k; i != 0; i += k {
		if i >= n {
			i -= n
		}

		energy -= 1

		if c[i] == 1 {
			energy -= 2
		}

		if i == 0 {
			break
		}
	}

	return energy
}

func main() {
	fmt.Println(jumpingOnClouds([]int32{0, 0, 1, 0, 0, 1, 1, 0}, 2))
}
