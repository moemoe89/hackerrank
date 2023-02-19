package main

import "fmt"

func breakingRecords(scores []int32) []int32 {
	var out []int32
	var max, min int32

	lastMax := scores[0]
	lastMin := scores[0]

	for i := range scores {
		if scores[i] > lastMax {
			lastMax = scores[i]
			max++
		}

		if scores[i] < lastMin {
			lastMin = scores[i]
			min++
		}
	}

	return append(out, max, min)
}

func main() {
	fmt.Println(breakingRecords([]int32{10, 5, 20, 20, 4, 5, 2, 25, 1}))
}
