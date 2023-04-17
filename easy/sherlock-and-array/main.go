// https://www.hackerrank.com/challenges/sherlock-and-array/problem?isFullScreen=true
package main

import "fmt"

func balancedSums(arr []int32) string {
	sum := int32(0)

	for i := range arr {
		sum += arr[i]
	}

	leftSum := int32(0)

	for i := range arr {
		if leftSum == sum-leftSum-arr[i] {
			return "YES"
		}

		leftSum += arr[i]
	}

	return "NO"
}

func main() {
	fmt.Println(balancedSums([]int32{5, 6, 7, 11}))
}
