// https://www.hackerrank.com/challenges/ctci-making-anagrams/problem?isFullScreen=true
package main

import "fmt"

func makeAnagram(a string, b string) int32 {
	m := make(map[rune]int32)

	for _, r := range a {
		m[r]++
	}

	for _, r := range b {
		m[r]--
	}

	output := int32(0)

	for _, val := range m {
		if val < 0 {
			output += -(val)
		} else {
			output += val
		}
	}

	return output
}

func main() {
	fmt.Println(makeAnagram("cde", "abc"))
}
