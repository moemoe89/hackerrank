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

func main() {
	fmt.Println(hackerrankInString("hereiamstackerrank"))
}
