package main

import "fmt"

func getMoneySpent(keyboards []int32, drives []int32, b int32) int32 {
	max := int32(0)

	for _, k := range keyboards {
		for _, d := range drives {
			tmp := k + d

			if tmp <= b && tmp > max {
				max = tmp
			}
		}
	}

	if max != 0 {
		return max
	}

	return -1
}

func main() {
	fmt.Println(getMoneySpent([]int32{3, 1}, []int32{5, 2, 8}, 10))
}
