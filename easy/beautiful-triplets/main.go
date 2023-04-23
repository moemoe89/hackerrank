// https://www.hackerrank.com/challenges/beautiful-triplets/problem?isFullScreen=true
package main

import "fmt"

func beautifulTriplets(d int32, arr []int32) int32 {
	// initialize map to store the elements.
	m := make(map[int32]bool)

	// store arr element to map.
	for _, a := range arr {
		m[a] = true
	}

	// initialize output var
	output := int32(0)

	// iterates the arr
	for _, i := range arr {
		// check the triplets key exist or not in map
		if m[i+d] && m[i+d*2] {
			output++
		}
	}

	return output
}

func main() {
	fmt.Println(beautifulTriplets(3, []int32{1, 2, 4, 5, 7, 8, 10}))
}
