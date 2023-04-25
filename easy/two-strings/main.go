// https://www.hackerrank.com/challenges/two-strings/problem?isFullScreen=true
package main

import "fmt"

func twoStrings(s1 string, s2 string) string {
	charMap := make(map[rune]bool)

	for _, char := range s1 {
		charMap[char] = true
	}

	for _, char := range s2 {
		if charMap[char] {
			return "YES"
		}
	}

	return "NO"
}

func main() {
	fmt.Println(twoStrings("hello", "world"))
}
