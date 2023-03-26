package main

import "fmt"

func getMoneySpent(keyboards []int32, drives []int32, b int32) int32 {
	max := int32(0)

	// iterates keyboards and drives
	// to find the maximum combination
	for _, k := range keyboards {
		for _, d := range drives {
			// sum keyboard + drive price
			tmp := k + d

			// if the number lower or equal with the money
			// and the number more than max value
			// replace max value
			if tmp <= b && tmp > max {
				max = tmp
			}
		}
	}

	// if max not 0, means the max combination found
	if max != 0 {
		return max
	}

	// if not found return -1
	return -1
}

func main() {
	fmt.Println(getMoneySpent([]int32{3, 1}, []int32{5, 2, 8}, 10))
}
