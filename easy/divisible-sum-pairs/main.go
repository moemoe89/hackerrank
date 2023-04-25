// https://www.hackerrank.com/challenges/divisible-sum-pairs/problem?isFullScreen=true
package main

import "fmt"

func divisibleSumPairs(n int32, k int32, ar []int32) int32 {
	output := int32(0)

	for i := 0; i < int(n); i++ {
		for j := i + 1; j < int(n); j++ {
			if (ar[i]+ar[j])%k == 0 {
				output++
			}
		}
	}

	return output
}

func main() {
	fmt.Println(divisibleSumPairs(6, 3, []int32{1, 3, 2, 6, 1, 2}))
}
