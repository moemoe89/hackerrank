// https://www.hackerrank.com/challenges/sparse-arrays/problem?isFullScreen=true
package main

import "fmt"

func matchingStrings(stringList []string, queries []string) []int32 {
	m := make(map[string]int32, len(stringList))

	for _, s := range stringList {
		m[s]++
	}

	o := make([]int32, len(queries))

	for i, q := range queries {
		o[i] = m[q]
	}

	return o
}

func main() {
	fmt.Println(matchingStrings([]string{"aba", "baba", "aba", "xzxb"}, []string{"aba", "xzxb", "ab"}))
}
