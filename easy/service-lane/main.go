// https://www.hackerrank.com/challenges/service-lane/problem?isFullScreen=true
package main

import "fmt"

var Width []int32

func serviceLane(n int32, cases [][]int32) []int32 {
	output := make([]int32, len(cases))

	for i := range cases {

		min := int32(3)

		for j := cases[i][0]; j <= cases[i][1]; j++ {
			if min > Width[j] {
				min = Width[j]
			}
		}

		output[i] = min

	}

	return output
}

func main() {
	Width = []int32{2, 3, 1, 2, 3, 2, 3, 3}

	fmt.Println(serviceLane(8, [][]int32{
		{0, 3},
		{4, 6},
		{6, 7},
		{3, 5},
		{0, 7},
	}))
}
