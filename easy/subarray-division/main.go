// https://www.hackerrank.com/challenges/the-birthday-bar/problem?isFullScreen=true
package main

import "fmt"

func birthday(s []int32, d int32, m int32) int32 {
	var out int32

	for k, _ := range s {
		var i, sum int32

		for i = 0; i < m; i++ {
			j := k + int(i)

			if j > len(s)-1 {
				break
			}

			sum += s[j]
		}

		if sum == d {
			out++
		}
	}

	return out
}

func main() {
	fmt.Println(birthday([]int32{1, 2, 1, 3, 2}, 3, 2))
}
