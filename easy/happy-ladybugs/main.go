// https://www.hackerrank.com/challenges/happy-ladybugs/problem?isFullScreen=true
package main

import "fmt"

func happyLadybugs(b string) string {
	m := make(map[rune]int)

	for _, r := range b {
		m[r]++
	}

	for k, v := range m {
		if k != '_' && v == 1 {
			return "NO"
		}
	}

	if m['_'] == 0 {
		for i := 1; i < len(b)-1; i++ {
			if b[i] != b[i-1] && b[i] != b[i+1] {
				return "NO"
			}
		}

		return "YES"
	}

	return "YES"
}

func main() {
	fmt.Println(happyLadybugs("RRRR"))
}
