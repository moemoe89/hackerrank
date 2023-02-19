package main

import "fmt"

func twoStrings(s1 string, s2 string) string {
	nS1 := len(s1)
	nS2 := len(s2)

	splitA := ""
	splitB := ""

	if nS1 > nS2 {
		splitA = s1
		splitB = s2
	} else {
		splitA = s2
		splitB = s1
	}

	mapS := make(map[rune]bool)
	for _, v := range splitA {
		mapS[v] = true
	}

	for _, v := range splitB {
		if _, ok := mapS[v]; ok {
			return "YES"
		}
	}

	return "NO"
}

func main() {
	fmt.Println(twoStrings("hello", "world"))
}
