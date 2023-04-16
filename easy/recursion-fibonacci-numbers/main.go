// https://www.hackerrank.com/challenges/ctci-fibonacci-numbers/problem?isFullScreen=true
package main

import "fmt"

// The Fibonacci sequence is the series of numbers where each number is the sum of the two preceding numbers.
// For example: 0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610,
func fibonacci(n int) int {
	m := map[int]int{
		0: 0,
		1: 1,
	}

	return fibonacciMap(n, m)
}

func fibonacciMap(n int, m map[int]int) int {
	if val, ok := m[n]; ok {
		return val
	}

	m[n] = fibonacciMap(n-2, m) + fibonacciMap(n-1, m)

	return m[n]
}

func main() {
	fmt.Println(fibonacci(3))
}
