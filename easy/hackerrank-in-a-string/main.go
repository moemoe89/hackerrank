// https://www.hackerrank.com/challenges/hackerrank-in-a-string/problem?isFullScreen=true
package main

import "fmt"

func hackerrankInString(s string) string {
	mapS := make(map[int]rune)

	str := "hackerrank"

	for i, r := range str {
		mapS[i] = r
	}

	i := 0

	for _, v := range s {
		if v == mapS[i] {
			i++

			continue
		}
	}

	if i == len(str) {
		return "YES"
	}

	return "NO"
}

func hackerrankInString2(s string) string {
	target := "hackerrank"

	j := 0
	for i := 0; i < len(s) && j < len(target); i++ {
		if s[i] == target[j] {
			j++
		}
	}

	if j == len(target) {
		return "YES"
	}

	return "NO"
}

func main() {
	fmt.Println(hackerrankInString("hereiamstackerrank"))
}
