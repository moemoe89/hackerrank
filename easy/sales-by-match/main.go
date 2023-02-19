package main

import "fmt"

func sockMerchant(n int32, ar []int32) int32 {
	mapPair := make(map[int]int)

	var pairs int32

	for _, v := range ar {
		idx := int(v)

		if _, ok := mapPair[idx]; ok {
			if mapPair[idx] < 2 {
				mapPair[idx]++
				pairs++

				continue
			}
		}

		mapPair[idx] = 1
	}

	return pairs
}

func main() {
	fmt.Println(sockMerchant(9, []int32{10, 20, 20, 10, 10, 30, 50, 10, 20}))
}
