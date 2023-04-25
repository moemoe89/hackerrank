// https://www.hackerrank.com/challenges/lonely-integer/problem?isFullScreen=true
package main

import "fmt"

func lonelyinteger(a []int32) int32 {
	m := make(map[int32]int32)

	for _, v := range a {
		m[v]++
	}

	for k, v := range m {
		if v == 1 {
			return k
		}
	}

	return 0
}

func main() {
	fmt.Println(lonelyinteger([]int32{1, 1, 2}))
}
