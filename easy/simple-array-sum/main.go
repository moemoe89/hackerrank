package main

import "fmt"

func simpleArraySum(ar []int32) int32 {
	output := int32(0)

	for i := range ar {
		output += ar[i]
	}

	return output
}

func main() {
	fmt.Println(simpleArraySum([]int32{1, 2, 3, 4, 10, 11}))
}
