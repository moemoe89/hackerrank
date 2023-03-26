package main

import (
	"fmt"
	"math/big"
)

// big sorting required big.Int
func bigSorting(unsorted []string) []string {
	arr := make([]*big.Int, len(unsorted))

	for i, s := range unsorted {
		base := 10
		bi := new(big.Int)
		_, _ = bi.SetString(s, base)

		arr[i] = bi
	}

	quickSort(arr, 0, len(arr)-1)

	sorted := make([]string, len(unsorted))

	for i, bi := range arr {
		sorted[i] = bi.String()
	}

	return sorted
}

func quickSort(arr []*big.Int, left int, right int) {
	if left < right {
		pivotIndex := partition(arr, left, right)
		quickSort(arr, left, pivotIndex-1)
		quickSort(arr, pivotIndex+1, right)
	}
}

func partition(arr []*big.Int, left int, right int) int {
	pivot := arr[right]

	i := left

	for j := left; j < right; j++ {
		// < operator means ascending sort
		// > operator means descending sort
		if arr[j].Cmp(pivot) == -1 {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}

	arr[i], arr[right] = arr[right], arr[i]

	return i
}

func main() {
	fmt.Println(bigSorting([]string{"6", "31415926535897932384626433832795", "1", "3", "10", "3", "5"}))
}
