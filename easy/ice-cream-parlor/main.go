// https://www.hackerrank.com/challenges/icecream-parlor/problem?isFullScreen=true
package main

import "fmt"

func icecreamParlor(m int32, arr []int32) []int32 {
	// initialize the index length.
	n := len(arr) - 1

	// initialize map to store the element and indices.
	mArr := make(map[int32][]int, n)

	// put all element to map with the indices.
	// for same element value, the indices will be more than 1.
	for i, a := range arr {
		mArr[a] = append(mArr[a], i)
	}

	// sort the arr due to we want to find by element by binary search.
	quickSort(arr, 0, n)

	// iterates the array.
	for i, a := range arr {
		// find the target combination.
		target := m - a

		// find the target using binary search,
		// for reduce the array size, always cut
		// from right current position.
		res := binarySearch(arr[i+1:], target)
		if res != -1 {
			// prepare the index output.
			var first, second int32

			// assign the first index value
			first = int32(mArr[a][0] + 1)

			// if the element value is same,
			// it's mean the number is duplicated.
			if a == target {
				// get by the 2nd index from current / target key.
				second = int32(mArr[a][1] + 1)
			} else {
				// get by 1st index from target key.
				second = int32(mArr[target][0] + 1)
			}

			// swap the value from lower to higher.
			if first > second {
				first, second = second, first
			}

			return []int32{first, second}
		}
	}

	return []int32{}
}

func binarySearch(arr []int32, x int32) int32 {
	left := 0
	right := len(arr) - 1

	for left <= right {
		mid := (left + right) / 2

		if arr[mid] == x {
			return int32(mid)
		} else if arr[mid] > x {
			right = mid - 1
		} else if arr[mid] < x {
			left = mid + 1
		}
	}

	return -1
}

func quickSort(arr []int32, left int, right int) {
	if left < right {
		pivotIndex := partition(arr, left, right)
		quickSort(arr, left, pivotIndex-1)
		quickSort(arr, pivotIndex+1, right)
	}
}

func partition(arr []int32, left int, right int) int {
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

func icecreamParlor2(m int32, arr []int32) []int32 {
	// initialize map for store the element and index.
	mArr := make(map[int32]int, len(arr))

	// iterates the array
	for i, a := range arr {
		// find the target by checking the map key represent as element
		// e.g. m = 6, a = 5 then find map with key 1
		if val, ok := mArr[m-a]; ok {
			// if found, return the both index and plus with 1
			return []int32{int32(val + 1), int32(i + 1)}
		}

		// if not found, always store the element and index to map.
		mArr[a] = i
	}

	// if the combination not found, return empty arr.
	return []int32{}
}

func icecreamParlor3(m int32, arr []int32) []int32 {
	// initialize the index length.
	n := len(arr) - 1

	// iterates the array and doing nested iterates to find the combination
	// that match with m value.
	for i := 0; i < n; i++ {
		for j := i + 1; j <= n; j++ {
			if arr[i]+arr[j] == m {
				return []int32{int32(i + 1), int32(j + 1)}
			}
		}
	}

	return []int32{}
}

func main() {
	fmt.Println(icecreamParlor(6, []int32{1, 3, 4, 5, 6}))
}
