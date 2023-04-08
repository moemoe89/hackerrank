// https://www.hackerrank.com/challenges/beautiful-triplets/problem?isFullScreen=true
package main

import "fmt"

func beautifulTriplets(d int32, arr []int32) int32 {
	// initialize map to store the elements.
	m := make(map[int32]int32)

	// store arr element to map.
	for _, a := range arr {
		m[a]++
	}

	// initialize output var
	output := int32(0)

	// iterates arr backward.
	for i := len(arr) - 1; i >= 0; i-- {
		// if the element lower than d,
		// it's impossible to have the beauty triplets,
		// then stop it.
		if arr[i] <= d {
			break
		}

		// each pair diff should equal with d,
		// then we can find the map key based on that each pair diff.
		p1 := arr[i]
		p2 := p1 - d
		p3 := p2 - d

		// if one of map doesn't exists, it will return 0,
		// then output will added by 0 which not changes anything.
		output += m[p1] * m[p2] * m[p3]
	}

	return output
}

func main() {
	fmt.Println(beautifulTriplets(3, []int32{1, 2, 4, 5, 7, 8, 10}))
}
