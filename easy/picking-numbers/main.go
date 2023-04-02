// https://www.hackerrank.com/challenges/picking-numbers/problem?isFullScreen=true
package main

import "fmt"

func pickingNumbers(a []int32) int32 {
	// initialize the map, key for the a element,
	// value for the frequency the a element appeared.
	m := make(map[int32]int32, len(a))
	for i := range a {
		m[a[i]]++
	}

	// initialize the output value.
	output := int32(0)

	// iterates the array.
	for i := range a {
		// sum the element and next (+1) element.
		// because the difference should be 0 or 1.
		// if the next element doesn't exists,
		// it simply just calculates the current element.
		n := m[a[i]] + m[a[i]+1]

		// check if the number higher than output
		if n > output {
			output = n
		}
	}

	return output
}

func main() {
	fmt.Println(pickingNumbers([]int32{4, 97, 5, 97, 97, 4, 97, 4, 97, 97, 97, 97, 4, 4, 5, 5, 97, 5, 97, 99, 4, 97, 5, 97, 97, 97, 5, 5, 97, 4, 5, 97, 97, 5, 97, 4, 97, 5, 4, 4, 97, 5, 5, 5, 4, 97, 97, 4, 97, 5, 4, 4, 97, 97, 97, 5, 5, 97, 4, 97, 97, 5, 4, 97, 97, 4, 97, 97, 97, 5, 4, 4, 97, 4, 4, 97, 5, 97, 97, 97, 97, 4, 97, 5, 97, 5, 4, 97, 4, 5, 97, 97, 5, 97, 5, 97, 5, 97, 97, 97}))
}
