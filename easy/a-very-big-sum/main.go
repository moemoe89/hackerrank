// https://www.hackerrank.com/challenges/a-very-big-sum/problem?isFullScreen=true
package main

import "fmt"

// the test case using a big number as an input, we need to choose
// the right data type to handle the number. here are some list about the data type:
// uint8  : 0 to 255
// uint16 : 0 to 65535
// uint32 : 0 to 4294967295
// uint64 : 0 to 18446744073709551615
// int8   : -128 to 127
// int16  : -32768 to 32767
// int32  : -2147483648 to 2147483647
// int64  : -9223372036854775808 to 9223372036854775807
func aVeryBigSum(ar []int64) int64 {
	output := int64(0)

	for i := range ar {
		output += ar[i]
	}

	return output
}

func main() {
	fmt.Println(aVeryBigSum([]int64{1000000001, 1000000002, 1000000003, 1000000004, 1000000005}))
}
