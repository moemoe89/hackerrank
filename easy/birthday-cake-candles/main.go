// https://www.hackerrank.com/challenges/birthday-cake-candles/problem?isFullScreen=true
package main

import "fmt"

func birthdayCakeCandles(candles []int32) int32 {
	// initialize the highest variable.
	highest := int32(0)

	// initialize the map for the element and the number of element appeared.
	m := make(map[int32]int32)

	// iterates the candles.
	for _, c := range candles {
		// store the candle element to map and increment the numbers.
		m[c]++

		// if current candle value higher than highest, replace it.
		if c > highest {
			highest = c
		}
	}

	// get the number appeared for the highest element.
	return m[highest]
}

func main() {
	fmt.Println(birthdayCakeCandles([]int32{3, 2, 1, 3}))
}
