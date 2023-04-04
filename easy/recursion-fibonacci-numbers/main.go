// https://www.hackerrank.com/challenges/ctci-fibonacci-numbers/problem?isFullScreen=true
package main

import "fmt"

// The Fibonacci sequence is the series of numbers where each number is the sum of the two preceding numbers.
// For example: 0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610,
func fibonacci(n int) int {
	// if n = 0 or n = 1, return n
	if n <= 1 {
		return n
	}

	// do recursive
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	fmt.Println(fibonacci(3))
}
