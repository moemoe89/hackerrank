// https://www.hackerrank.com/challenges/anagram/problem?isFullScreen=true
package main

import "fmt"

func anagram(s string) int32 {
	n := len(s)

	if n%2 != 0 {
		return -1
	}

	m := make(map[rune]int32)

	for _, r := range s[:n/2] {
		m[r]++
	}

	for _, r := range s[n/2:] {
		m[r]--
	}

	output := int32(0)

	for _, v := range m {
		if v > 0 {
			output += v
		}
	}

	return output
}

func main() {
	fmt.Println(anagram("xaxbbbxx"))
}
