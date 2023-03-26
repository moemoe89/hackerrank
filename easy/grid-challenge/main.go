package main

import (
	"fmt"
)

func gridChallenge(grid []string) string {
	// initialize length index of grid
	n := len(grid) - 1

	// iterates grid array
	for i := 0; i <= n; i++ {
		// transform to byte, in order to
		// support sorting.
		g := []byte(grid[i])

		// ascending sort the string in grid.
		quickSort(g, 0, len(g)-1)

		// put back the sorted value to grid.
		grid[i] = string(g)
	}

	// check the column order
	// first do iterates for string length (right),
	// and then iterates the grid (down)
	for i := 0; i < len(grid[0]); i++ {
		// start from 1 because we want to compare
		// with the previous row, which is minus 1.
		for j := 1; j <= n; j++ {
			// if not sorted return NO
			if grid[j][i] < grid[j-1][i] {
				return "NO"
			}
		}
	}

	return "YES"
}

func quickSort(arr []byte, left int, right int) {
	if left < right {
		pivotIndex := partition(arr, left, right)
		quickSort(arr, left, pivotIndex-1)
		quickSort(arr, pivotIndex+1, right)
	}
}

func partition(arr []byte, left int, right int) int {
	pivot := arr[right]

	i := left

	for j := left; j < right; j++ {
		// < operator means ascending sort
		// > operator means descending sort
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}

	arr[i], arr[right] = arr[right], arr[i]

	return i
}

func main() {
	fmt.Println(gridChallenge([]string{"ebacd", "fghij", "olmkn", "trpqs", "xywuv"}))
}
