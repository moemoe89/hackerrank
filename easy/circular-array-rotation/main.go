package main

import "fmt"

func circularArrayRotation(a []int32, k int32, queries []int32) []int32 {
	n := int32(len(a))

	k = k % int32(len(a))

	a = append(a[n-k:], a[:n-k]...)

	output := make([]int32, len(queries))

	for i, q := range queries {
		output[i] = a[q]
	}

	return output
}

func main() {
	fmt.Println(circularArrayRotation([]int32{3, 4, 5}, 2, []int32{1, 2}))
}
