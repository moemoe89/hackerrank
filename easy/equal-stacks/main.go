// https://www.hackerrank.com/challenges/equal-stacks/problem?isFullScreen=true
package main

import "fmt"

func equalStacks(h1 []int32, h2 []int32, h3 []int32) int32 {
	// get sum of all stacks.
	total1, total2, total3 := sum(h1), sum(h2), sum(h3)

	// loop until the all total equal.
	for {
		// break if all total is equal.
		if total1 == total2 && total1 == total3 {
			return total1
		}

		// check if total1 is the highest, then remove the first stack
		// then going back to the beginning to check the all total is equal or not.
		// do the same for total2 and total 3.
		if total1 >= total2 && total1 >= total3 {
			total1 -= h1[0]
			h1 = h1[1:]
		} else if total2 >= total1 && total2 >= total3 {
			total2 -= h2[0]
			h2 = h2[1:]
		} else if total3 >= total1 && total3 >= total2 {
			total3 -= h3[0]
			h3 = h3[1:]
		}
	}
}

func sum(h []int32) int32 {
	total := int32(0)

	for i := range h {
		total += h[i]
	}

	return total
}

func main() {
	fmt.Println(equalStacks([]int32{3, 2, 1, 1, 1}, []int32{4, 3, 2}, []int32{1, 1, 4, 1}))
}
