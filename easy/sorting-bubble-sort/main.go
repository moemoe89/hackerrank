package main

import "fmt"

func countSwaps(a []int32) {
	n := len(a) - 1

	swaps := 0

	for i := 0; i < n; i++ {

		for j := 0; j < n-i; j++ {
			if a[j] > a[j+1] {
				swaps++
				a[j], a[j+1] = a[j+1], a[j]
			}
		}

	}

	fmt.Printf("Array is sorted in %v swaps.\n", swaps)
	fmt.Printf("First Element: %v\n", a[0])
	fmt.Printf("Last Element: %v\n", a[n])
}

func main() {
	countSwaps([]int32{6, 4, 1})
}
