// https://www.hackerrank.com/challenges/sum-vs-xor/problem?isFullScreen=true
package main

import "fmt"

func sumXor(n int64) int64 {
	count := int64(0)

	for n > 0 {
		if n%2 == 0 {
			count++
		}

		n /= 2
	}

	return int64(1) << count
}

func main() {
	fmt.Println(sumXor(10))
}
