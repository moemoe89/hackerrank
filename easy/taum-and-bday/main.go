// https://www.hackerrank.com/challenges/taum-and-bday/problem?isFullScreen=true
package main

import "fmt"

func taumBday(b int32, w int32, bc int32, wc int32, z int32) int64 {
	// convert input parameters to int64.
	var b64, w64, bc64, wc64, z64 = int64(b), int64(w), int64(bc), int64(wc), int64(z)

	// initialize the min variable as original cost.
	// (black gift * black cost) + (white gift * white cost)
	min := (b64 * bc64) + (w64 * wc64)

	// calculate cost use the white price, but it's mean
	// need to add cost convert to white price when calculating black gift cost.
	wcPrice := wc64 + z64
	minWC := (b64 * wcPrice) + (w64 * wc64)

	// if it's lower than current min value, replace it.
	if minWC < min {
		min = minWC
	}

	// calculate cost use the black price, but it's mean
	// need to add cost convert to black price when calculating black white cost.
	bcPrice := bc64 + z64
	minBC := (b64 * bc64) + (w64 * bcPrice)

	// if it's lower than current min value, replace it.
	if minBC < min {
		min = minBC
	}

	return min
}

func main() {
	fmt.Println(taumBday(3, 5, 3, 4, 1))
}
