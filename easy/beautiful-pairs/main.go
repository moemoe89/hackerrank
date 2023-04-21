// https://www.hackerrank.com/challenges/beautiful-pairs/problem?isFullScreen=true
package main

import "fmt"

func beautifulPairs(A []int32, B []int32) int32 {
	freq := make(map[int32]int)
	var count int32

	// Count the frequency of each element in A
	for _, a := range A {
		freq[a]++
	}

	// Try to form beautiful pairs by iterating through B
	for _, b := range B {
		if freq[b] > 0 {
			freq[b]--
			count++
		}
	}

	// If we were able to form a beautiful pair for every element in B,
	// we can increase the count by 1 to account for the swapping of an element in A
	if count == int32(len(A)) {
		count--
	} else {
		count++
	}

	return count
}

func main() {
	fmt.Println(beautifulPairs([]int32{1, 2, 4, 3}, []int32{1, 2, 3, 4}))
}
