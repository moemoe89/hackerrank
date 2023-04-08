// https://www.hackerrank.com/challenges/one-week-preparation-kit-countingsort1/problem?isFullScreen=true
package main

import "fmt"

func countingSort(arr []int32) []int32 {
	// create map with capacity for preventing memory allocation several times
	mapArr := make(map[int32]int, len(arr))

	for _, a := range arr {
		// if number is newly found, set value to 1,
		// if already exists, adding by 1
		num := 1

		if val, ok := mapArr[a]; ok {
			num = val + 1
		}

		mapArr[a] = num
	}

	// create slice with 100 capacity as question note
	output := make([]int32, 100)

	// assign counted value to sorted index
	for k, v := range mapArr {
		output[int(k)] = int32(v)
	}

	return output
}

func main() {
	fmt.Println(countingSort([]int32{63, 25, 73, 1, 98, 73, 56, 84, 86, 57, 16, 83, 8, 25, 81, 56, 9, 53, 98, 67, 99, 12, 83, 89, 80, 91, 39, 86, 76, 85, 74, 39, 25, 90, 59, 10, 94, 32, 44, 3, 89, 30, 27, 79, 46, 96, 27, 32, 18, 21, 92, 69, 81, 40, 40, 34, 68, 78, 24, 87, 42, 69, 23, 41, 78, 22, 6, 90, 99, 89, 50, 30, 20, 1, 43, 3, 70, 95, 33, 46, 44, 9, 69, 48, 33, 60, 65, 16, 82, 67, 61, 32, 21, 79, 75, 75, 13, 87, 70, 33}))
}
