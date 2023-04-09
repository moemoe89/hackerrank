// https://www.hackerrank.com/challenges/sherlock-and-squares/problem?isFullScreen=true
package main

import (
	"fmt"
	"math"
)

func squares(a int32, b int32) int32 {
	// Calculate the square roots of the endpoints of the range
	sqrtA := int32(math.Ceil(math.Sqrt(float64(a))))
	sqrtB := int32(math.Floor(math.Sqrt(float64(b))))

	// Calculate the number of integers between the square roots
	return sqrtB - sqrtA + 1
}

func main() {
	fmt.Println(squares(24, 49))
}
