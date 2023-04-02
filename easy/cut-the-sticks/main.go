// https://www.hackerrank.com/challenges/cut-the-sticks/problem?isFullScreen=true
package main

import "fmt"

func cutTheSticks(arr []int32) []int32 {
	// initialize the first ouput with the original length arr.
	output := []int32{int32(len(arr))}

	// iterates until the array become empty.
	for {
		// find the shortest value by iterates and checks arr elements.
		shortest := arr[0]
		for i := 1; i < len(arr); i++ {
			if shortest > arr[i] {
				shortest = arr[i]
			}
		}

		// initialize the tmp array, for the cut number.
		var tmp []int32
		for i := range arr {
			// cut the number with the shortest number.
			n := arr[i] - shortest
			// if the cut number more than 0, put in tmp array.
			if n > 0 {
				tmp = append(tmp, n)
			}
		}

		// if the tmp array is empty, means all numbers already cut,
		// return the output and break the loop.
		if len(tmp) == 0 {
			return output
		}

		// replace arr value with tmp arr.
		arr = tmp

		// push the new arr length to output.
		output = append(output, int32(len(arr)))
	}
}

func main() {
	fmt.Println(cutTheSticks([]int32{8, 8, 14, 10, 3, 5, 14, 12}))
}
