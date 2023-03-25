package main

import "fmt"

func icecreamParlor(m int32, arr []int32) []int32 {
	n := len(arr) - 1

	mArr := make(map[int32][]int, n)

	for i, a := range arr {
		mArr[a] = append(mArr[a], i)
	}

	quickSort(arr, 0, n)

	for i, a := range arr {
		target := m - a

		res := binarySearch(arr[i+1:], target)
		if res != -1 {
			var first, second int32

			first = int32(mArr[a][0] + 1)

			if a == target {
				second = int32(mArr[a][1] + 1)
			} else {
				second = int32(mArr[target][0] + 1)
			}

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
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}

	arr[i], arr[right] = arr[right], arr[i]

	return i
}

func icecreamParlor2(m int32, arr []int32) []int32 {
	mArr := make(map[int32]int, len(arr))

	for i, a := range arr {
		if val, ok := mArr[m-a]; ok {
			return []int32{int32(val + 1), int32(i + 1)}
		}

		mArr[a] = i
	}

	return []int32{}
}

func icecreamParlor3(m int32, arr []int32) []int32 {
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j <= len(arr)-1; j++ {
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
