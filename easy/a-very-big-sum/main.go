package main

import "fmt"

func aVeryBigSum(ar []int64) int64 {
	output := int64(0)

	for i := range ar {
		output += ar[i]
	}

	return output
}

func main() {
	fmt.Println(aVeryBigSum([]int64{1000000001, 1000000002, 1000000003, 1000000004, 1000000005}))
}
