// https://www.hackerrank.com/challenges/correctness-invariant/problem?isFullScreen=true
package main

import "fmt"

func main() {
	// read input, first is the total number
	// of the numbers will as an input.
	var n int
	_, _ = fmt.Scan(&n)

	// initialize array with the total numbers.
	arr := make([]int, n)
	// iterates to catch each input numbers.
	for i := 0; i < n; i++ {
		_, _ = fmt.Scan(&arr[i])
	}

	// insertion sort
	// how this algorithm works can be seen in this gif
	// https://commons.wikimedia.org/wiki/File:Insertion-sort-example.gif
	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}

	// print output
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", arr[i])
	}
}
