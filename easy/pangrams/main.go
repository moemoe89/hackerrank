package main

import "fmt"

func pangrams(s string) string {
	mapS := make(map[rune]struct{}, len(s))

	for _, r := range s {
		if r >= 'A' && r <= 'Z' {
			r = r - 'A' + 'a'
		}

		if r >= 'a' && r <= 'z' {
			mapS[r] = struct{}{}
		}
	}

	if len(mapS) == 26 {
		return "pangram"
	}

	return "not pangram"
}

func main() {
	fmt.Println(pangrams("We promptly judged antique ivory buckles for the prize"))
}
