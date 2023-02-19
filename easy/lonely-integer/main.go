package main

import "fmt"

func lonelyinteger(a []int32) int32 {
	mapA := make(map[int32]int)

	for _, v := range a {
		mapA[v] = mapA[v] + 1
	}

	var out int32

	for k, v := range mapA {
		if v == 1 {
			out = k
			
			break
		}
	}

	return out
}

func main() {
	fmt.Println(lonelyinteger([]int32{1, 1, 2}))
}
