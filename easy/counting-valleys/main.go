// https://www.hackerrank.com/challenges/counting-valleys/problem?isFullScreen=true
package main

import "fmt"

func countingValleys(n int32, s string) int32 {
	var output int32

	position := 0

	for _, r := range s {
		if r == 'U' {
			position += 1
		} else if r == 'D' {
			position -= 1
		}

		if position == 0 {
			if r == 'U' {
				output++
			}
		}
	}

	return output
}

func main() {
	fmt.Println(countingValleys(8, "UDDDUDUU"))
}
