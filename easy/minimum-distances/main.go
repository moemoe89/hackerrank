// https://www.hackerrank.com/challenges/minimum-distances/problem?isFullScreen=true
package main

import "fmt"

func minimumDistances(a []int32) int32 {
	// initialize the length of arr.
	n := len(a)

	// initialize the map for the element and index.
	m := make(map[int32]int, n)

	// set the minDiff same as length, because if
	// there's a pair element exist, the value should
	// lower than the length.
	// e.g. [7, 1, 3, 4, 1, 7] the pair is 7 [1, 4] and 1 [0, 5]
	// value for pair 7 is 0 - 5 => 5 and 1 is 1 - 4 => 3
	// it will lower than the length = 6
	minDiff := n

	// iterates the array.
	for i, x := range a {
		// if the key of element found, it's mean there's a pair exist
		if j, ok := m[x]; ok {
			// calculate the diff based on both index
			diff := i - j

			// if diff lower than minDiff, replace it.
			if diff < minDiff {
				minDiff = diff
			}
		}

		// always store the element and index to map.
		m[x] = i
	}

	// if minDiff value equal with the length,
	// it's mean no pair element exists.
	if minDiff == n {
		return -1
	}

	return int32(minDiff)
}

func main() {
	fmt.Println(minimumDistances([]int32{7, 1, 3, 4, 1, 7}))
}
