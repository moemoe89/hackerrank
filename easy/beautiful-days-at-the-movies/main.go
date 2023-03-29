package main

import "fmt"

func beautifulDays(i int32, j int32, k int32) int32 {
	var out int32

	// find the day between number i and j
	for ; i <= j; i++ {
		// get the diff from i and reserved i
		subs := i - reverseInt32(i)

		// if the modulo is 0, it's beautiful day.
		if subs%k == 0 {
			out++
		}
	}

	return out
}

// just like checking palindrom,
// this function is to reverse number.
// e.g. 2000 -> 0002 (2), 1234 -> 4321
func reverseInt32(i int32) int32 {
	reversed := int32(0)

	for i > 0 {
		// reverse the input digit by digit from backward
		// for the example input 123
		// first  -> reserve*10 : 0,   x%10 : 3, reserve : 3,   x divide 10 : 12
		// second -> reserve*10 : 30,  x%10 : 2, reserve : 32,  x divide 10 : 1
		// third  -> reserve*10 : 320, x%10 : 1, reserve : 321, x divide 10 : 0
		// break the loop since x = 0
		reversed = reversed*10 + i%10

		// divide i by 10
		i /= 10
	}

	return reversed
}

func main() {
	fmt.Println(beautifulDays(20, 23, 6))
}
