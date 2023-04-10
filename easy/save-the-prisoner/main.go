// https://www.hackerrank.com/challenges/save-the-prisoner/problem?isFullScreen=true
package main

import "fmt"

func saveThePrisoner(n int32, m int32, s int32) int32 {
	last := s + m - 1

	last = last % n

	if last == 0 {
		return n
	}

	return last
}

func main() {
	fmt.Println(saveThePrisoner(4, 6, 2))
}
