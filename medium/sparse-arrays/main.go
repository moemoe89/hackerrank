package main

import "fmt"

func matchingStrings(stringList []string, queries []string) []int32 {
	m := make(map[string]int32, 0)

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
