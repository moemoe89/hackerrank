// https://www.hackerrank.com/challenges/two-characters/problem?isFullScreen=true
package main

import "fmt"

func alternate(s string) int32 {
	m := make(map[rune]struct{})

	var chars []rune

	for _, r := range s {
		if _, ok := m[r]; ok {
			continue
		}

		m[r] = struct{}{}

		chars = append(chars, r)
	}

	var max int32

	for i := 0; i < len(chars)-1; i++ {
		for j := i + 1; j < len(chars); j++ {
			strs := []rune{}
			for _, r := range s {
				if r == chars[i] || r == chars[j] {
					strs = append(strs, r)
				}
			}

			isDup := false
			for k := 1; k < len(strs); k++ {
				if strs[k] == strs[k-1] {
					isDup = true
					break
				}
			}

			if !isDup && int32(len(strs)) > max {
				max = int32(len(strs))
			}
		}
	}

	return max
}

func main() {
	fmt.Println(alternate("asdcbsdcagfsdbgdfanfghbsfdab"))
}
