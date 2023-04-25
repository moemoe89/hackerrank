// https://www.hackerrank.com/challenges/mini-max-sum/problem?isFullScreen=true
package main

import "fmt"

func miniMaxSum(arr []int64) {
	var total, minInt, maxInt int64

	for k, v := range arr {
		total += v

		if k == 0 {
			minInt = v
			maxInt = v

			continue
		}

		if minInt > v {
			minInt = v
		}

		if maxInt < v {
			maxInt = v
		}
	}

	fmt.Printf("%d %d\n", total-maxInt, total-minInt)
}

func main() {
	miniMaxSum([]int64{1, 2, 3, 4, 5})
}
