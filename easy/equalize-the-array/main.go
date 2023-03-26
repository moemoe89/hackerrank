package main

import "fmt"

func equalizeArray(arr []int32) int32 {
	m := make(map[int32]int, len(arr))

	max := 0

	for _, a := range arr {
		m[a]++

		if m[a] > max {
			max = m[a]
		}
	}

	return int32(len(arr) - max)
}

func main() {
	fmt.Println(equalizeArray([]int32{3, 3, 2, 1, 3}))
}
