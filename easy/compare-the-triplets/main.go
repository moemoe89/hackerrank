package main

import "fmt"

func compareTriplets(a []int32, b []int32) []int32 {
	output := []int32{0, 0}

	for i := range a {
		if a[i] > b[i] {
			output[0]++
		} else if a[i] < b[i] {
			output[1]++
		}
	}

	return output
}

func main() {
	fmt.Println(compareTriplets([]int32{5, 6, 7}, []int32{3, 6, 10}))
}
