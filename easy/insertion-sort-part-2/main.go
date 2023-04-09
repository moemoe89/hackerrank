// https://www.hackerrank.com/challenges/insertionsort2/problem?isFullScreen=true
package main

import "fmt"

// https://commons.wikimedia.org/wiki/File:Insertion-sort-example.gif
func insertionSort2(n int32, arr []int32) {
	for i := 1; i < len(arr); i++ {
		key := arr[i]

		j := i - 1

		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			
			j = j - 1
		}

		arr[j+1] = key

		for i := range arr {
			fmt.Print(arr[i], " ")
		}

		fmt.Println()
	}
}

func main() {
	insertionSort2(6, []int32{1, 4, 3, 5, 6, 2})
}
