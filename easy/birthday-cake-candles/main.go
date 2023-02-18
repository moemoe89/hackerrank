package main

import "fmt"

func birthdayCakeCandles(candles []int32) int32 {
	highest := int32(0)

	m := make(map[int32]int32)

	for i := range candles {
		m[candles[i]]++

		if candles[i] > highest {
			highest = candles[i]
		}
	}

	return m[highest]
}

func main() {
	fmt.Println(birthdayCakeCandles([]int32{3, 2, 1, 3}))
}
