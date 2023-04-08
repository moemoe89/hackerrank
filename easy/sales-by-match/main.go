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

func sockMerchant2(n int32, ar []int32) int32 {
	sockCounts := make(map[int32]int32)

	for _, sock := range ar {
		sockCounts[sock]++
	}

	pairCount := int32(0)

	for _, count := range sockCounts {
		pairCount += count / 2
	}

	return pairCount
}

func main() {
	fmt.Println(sockMerchant(9, []int32{10, 20, 20, 10, 10, 30, 50, 10, 20}))
}
