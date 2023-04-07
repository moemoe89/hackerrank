// https://www.hackerrank.com/challenges/string-construction/problem?isFullScreen=true
package main

import "fmt"

func stringConstruction(s string) int32 {
	m := make(map[rune]struct{})

	var out int32

	for _, v := range s {
		// if the key doesn't exist, insert it.
		// if exists, skip it.
		if _, ok := m[v]; !ok {
			m[v] = struct{}{}
			out++
		}
	}

	return out
}

func main() {
	fmt.Println(stringConstruction("abcd"))
}
