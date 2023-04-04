// https://www.hackerrank.com/challenges/making-anagrams/problem?isFullScreen=true
package main

import "fmt"

func makingAnagrams(s1 string, s2 string) int32 {
	m := make(map[rune]int32)

	for _, r := range s1 {
		m[r]++
	}

	for _, r := range s2 {
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
	fmt.Println(makingAnagrams("cde", "abc"))
}
