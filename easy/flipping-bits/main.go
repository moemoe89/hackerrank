// https://www.hackerrank.com/challenges/flipping-bits/problem?isFullScreen=true
package main

import "fmt"

func flippingBits(n int64) int64 {
	return int64(uint32(^n))
}

func main() {
	fmt.Println(flippingBits(123456))
}
