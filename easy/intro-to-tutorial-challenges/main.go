// https://www.hackerrank.com/challenges/tutorial-intro/problem?isFullScreen=true
package main

import "fmt"

func introTutorial(V int32, arr []int32) int32 {
	for k, v := range arr {
		if v == V {
			return int32(k)
		}
	}

	return -1
}

func main() {
	fmt.Println(introTutorial(4, []int32{1, 4, 5, 6, 9, 12}))
}
