// https://www.hackerrank.com/challenges/minimum-swaps-2/problem?isFullScreen=true
package main

import "fmt"

func minimumSwaps(arr []int32) int32 {
	var swap, i, n int32

	n = int32(len(arr))

	for i = 0; i < n; i++ {
		for i+1 != arr[i] {
			tmp := arr[arr[i]-1]
			arr[arr[i]-1] = arr[i]
			arr[i] = tmp
			swap++
		}
	}

	return swap
}

func main() {
	fmt.Println(minimumSwaps([]int32{7, 1, 3, 2, 4, 5, 6}))
}
