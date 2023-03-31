// https://www.hackerrank.com/challenges/2d-array/problem?isFullScreen=true
package main

import "fmt"

func hourglassSum(arr [][]int32) int32 {
	// define the lowest value, it should be -9 * 6
	// -9 -9 -9
	//  0 -9  0
	// -9 -9 -9
	count := int32(-63)

	// we should limit the movement for right & bottom until index: 3
	// because the 2D array dimension is 6 x 6, if we reach index 4,
	// it will be out of range since we need to make 2 steps movement.
	// which index 4 and 2 movement, we will reach index 6 which is unavailable.
	// the max index is 5.
	for i := 0; i <= 3; i++ {
		for j := 0; j <= 3; j++ {
			// sum the position following this pattern:
			// 1 1 1
			//   1
			// 1 1 1
			sum := arr[i][j] + arr[i][j+1] + arr[i][j+2] + arr[i+1][j+1] + arr[i+2][j] + arr[i+2][j+1] + arr[i+2][j+2]

			// if the sum more thant count, replace it.
			if sum > count {
				count = sum
			}
		}
	}

	return count
}

func main() {
	fmt.Println(hourglassSum([][]int32{
		{1, 1, 1, 0, 0, 0},
		{0, 1, 0, 0, 0, 0},
		{1, 1, 1, 0, 0, 0},
		{0, 0, 2, 4, 4, 0},
		{0, 0, 0, 2, 0, 0},
		{0, 0, 1, 2, 4, 0},
	}))
}
